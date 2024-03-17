/*
 * SystemK
 *
 * API <br/>  Documentation:    <ul>      <li><a href=\"https://atlassian.net/wiki/</a></li>      <li><a href=\"https://git.net/f\">Git</a></li>      <li><a href=\"https://humansinc.atlassian.net/browse/ORM\">Jira</a></li>    </ul>  <br/>  <a href=\"git .yaml\">API Artifact</a><br/>
 *
 * API version: 1.0.0
 */
package mware

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	//"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUserV1(w http.ResponseWriter, r *http.Request) {
	var rec UserParams
	operation := "CreateUserV1"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn := GetConn()
	var userid string
	var sql = "INSERT INTO public.\"Users\"( \"firstName\", \"lastName\", \"secondName\", \"birhDate\", email, phone, username, password, \"statusId\", \"subDivisionId\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9,$10) RETURNING id"
	err = conn.QueryRow(context.Background(), sql, rec.FirstName, rec.LastName, rec.SecondName, rec.BirhDate, rec.Email, rec.Phone, rec.Username, rec.Password, rec.UserStatus, rec.SubDivision).Scan(&userid)
	defer conn.Release()
	if err != nil {

		ErrResponse(w, err, operation)
		return
	}
	sql = "SELECT u.id, u.\"firstName\", u.\"lastName\", u.\"secondName\", u.\"birhDate\", u.email, u.phone, u.username, u.password, u.\"statusId\", s.\"SubDivision\" FROM public.\"Users\" u, public.\"SubDivisions\" s where u.\"subDivisionId\"=s.\"id\" and u.id=$1"
	rows, _ := conn.Query(context.Background(), sql, userid)
	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[CreateUserResponse])
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	f, errF := json.Marshal(users)
	if errF != nil {
		log.Fatal(errF)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func DeleteUserV1(w http.ResponseWriter, r *http.Request) {
	operation := "DeleteUserV1"
	s := r.URL.Query()["id"]
	var userid uuid.UUID
	userid, err := uuid.FromString(strings.Join(s, ""))
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	conn := GetConn()
	sql := "Update public.\"Users\" set \"statusId\"='dropped' where id=$1 RETURNING id"
	err1 := conn.QueryRow(context.Background(), sql, userid).Scan(&userid)
	defer conn.Release()
	if err1 != nil {
		ErrResponse(w, err1, operation)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUserV1(w http.ResponseWriter, r *http.Request) {
	operation := "GetUserV1"
	s := r.URL.Query()["id"]
	var userid uuid.UUID
	userid, err := uuid.FromString(strings.Join(s, ""))
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	conn := GetConn()
	sql := "SELECT u.id, u.\"firstName\", u.\"lastName\", u.\"secondName\", u.\"birhDate\", u.email, u.phone, u.username, u.password, u.\"statusId\", s.\"SubDivision\" FROM public.\"Users\" u, public.\"SubDivisions\" s where u.\"subDivisionId\"=s.\"id\" and u.id=$1"
	rows, _ := conn.Query(context.Background(), sql, userid)
	defer conn.Release()
	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[CreateUserResponse])
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	f, errF := json.Marshal(users)
	if errF != nil {
		ErrResponse(w, errF, operation)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func UpdateUserV1(w http.ResponseWriter, r *http.Request) {
	operation := "UpdateUserV1"
	var rec UserParams
	s := r.URL.Query()["id"]
	var userid uuid.UUID
	userid, err := uuid.FromString(strings.Join(s, ""))
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err = json.NewDecoder(r.Body).Decode(&rec)
	if err != nil {
		ErrResponse(w, errors.New("bad request"), operation)
		return // bad request
	}
	conn := GetConn()
	var s_userid string
	var sql = "Update public.\"Users\" set \"firstName\"=$1, \"lastName\"=$2, \"secondName\"=$3, \"birhDate\"=$4, email=$5, phone=$6, username=$7, \"statusId\"=$8, \"subDivisionId\"=$9 where id=$10 RETURNING id"
	err = conn.QueryRow(context.Background(), sql, rec.FirstName, rec.LastName, rec.SecondName, rec.BirhDate, rec.Email, rec.Phone, rec.Username, rec.UserStatus, rec.SubDivision, userid).Scan(&s_userid)
	defer conn.Release()
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	sql = "SELECT u.id, u.\"firstName\", u.\"lastName\", u.\"secondName\", u.\"birhDate\", u.email, u.phone, u.username, u.password, u.\"statusId\", s.\"SubDivision\" FROM public.\"Users\" u, public.\"SubDivisions\" s where u.deleted=false and u.\"subDivisionId\"=s.\"id\" and u.id=$1"
	rows, _ := conn.Query(context.Background(), sql, s_userid)
	users, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[CreateUserResponse])
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	f, errF := json.Marshal(users)
	if errF != nil {
		log.Fatal(errF)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func ErrResponse(w http.ResponseWriter, err error, operation string) {
	var fault CommonFault
	fault.Operation = operation
	log.Print(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusBadRequest)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		fault.Code = pgErr.Code
		fault.Cause = pgErr.Message
	} else {
		fault.Code = "-1"
		fault.Cause = err.Error()
	}
	f, errF := json.Marshal(fault)
	if errF != nil {
		log.Fatal(errF)
	}
	w.Write(f)
}

func CreateSubDivisionV1(w http.ResponseWriter, r *http.Request) {
	operation := "CreateSubDivisionV1"
	var rec SubDivisions
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn := GetConn()
	var id string
	//var companyId=strconv.Atoi(rec.CompanyId)
	var sql = "INSERT INTO public.\"SubDivisions\"( \"SubDivision\",\"CompanyId\") VALUES ($1,$2) RETURNING id"
	err = conn.QueryRow(context.Background(), sql, rec.SubDivision, rec.CompanyId).Scan(&id)
	defer conn.Release()
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	sql = "SELECT id,\"SubDivision\",\"CompanyId\" FROM public.\"SubDivisions\"  where id=$1 and deleted=false"
	rows, _ := conn.Query(context.Background(), sql, id)
	s, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[CreateSubDivisionResponse])
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	f, errF := json.Marshal(s)
	if errF != nil {
		log.Fatal(errF)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func DeleteSubDivisionV1(w http.ResponseWriter, r *http.Request) {
	operation := "DeleteSubDivisionV1"
	s := r.URL.Query()["id"]
	var sql string
	var rows pgx.Rows
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	conn := GetConn()
	defer conn.Release()
	id, err := strconv.Atoi(s[0])
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	ts := time.Now()
	sql = "Update public.\"SubDivisions\"  set deleted=true, ts=$1 where id=$2"
	rows, _ = conn.Query(context.Background(), sql, ts, id)
	sd, err := pgx.CollectRows(rows, pgx.RowToStructByName[CreateSubDivisionResponse])
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	f, errF := json.Marshal(sd)
	if errF != nil {
		ErrResponse(w, errF, operation)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func GetSubDivisionV1(w http.ResponseWriter, r *http.Request) {
	operation := "GetSubDivisionV1"
	s := r.URL.Query()["id"]
	var id int
	var sql string
	var rows pgx.Rows
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	conn := GetConn()
	defer conn.Release()
	id, err := strconv.Atoi(strings.Join(s, ""))
	if err != nil {
		sql = "SELECT id, \"SubDivision\",\"CompanyId\" FROM public.\"SubDivisions\" where deleted=false"
		rows, _ = conn.Query(context.Background(), sql)
	} else {
		sql = "SELECT id, \"SubDivision\",\"CompanyId\" FROM public.\"SubDivisions\"  where deleted=false and id=$1"
		rows, _ = conn.Query(context.Background(), sql, id)
	}
	sd, err := pgx.CollectRows(rows, pgx.RowToStructByName[CreateSubDivisionResponse])
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	f, errF := json.Marshal(sd)
	if errF != nil {
		ErrResponse(w, errF, operation)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func UpdateSubDivisionV1(w http.ResponseWriter, r *http.Request) {
	//operation:="UpdateSubDivisionV1"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func CreateWorkPositionV1(w http.ResponseWriter, r *http.Request) {
	operation := "CreateWorkPositionV1"
	var rec WorkPosition
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn := GetConn()
	var id string
	var sql = "INSERT INTO public.\"WorkPosition\"(name, \"subDivisionId\") VALUES ($1,$2) RETURNING id"
	err = conn.QueryRow(context.Background(), sql, rec.WorkPosition, rec.SubDivisionId).Scan(&id)
	defer conn.Release()
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	sql = "SELECT id,\"name\",\"subDivisionId\" FROM public.\"WorkPosition\"  where id=$1"
	rows, _ := conn.Query(context.Background(), sql, id)
	s, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[CreateWorkPositionResponse])
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	f, errF := json.Marshal(s)
	if errF != nil {
		log.Fatal(errF)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func GetWorkPositionV1(w http.ResponseWriter, r *http.Request) {
	operation := "GetWorkPositionV1"
	s := r.URL.Query()["id"]
	var id int
	var sql string
	var rows pgx.Rows
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	conn := GetConn()
	/*
		pool := r.Context().Value("pgxpool").(*pgxpool.Pool)
		conn, err := pool.Acquire(context.Background())
		if err != nil {
			log.Fatalf("Unable to acquire a database connection: %v\n", err)
		}
	*/
	defer conn.Release()
	id, err1 := strconv.Atoi(strings.Join(s, ""))
	if err1 != nil {
		sql = "SELECT w.\"id\", w.\"name\",  w.\"subDivisionId\" FROM public.\"WorkPosition\" as w inner join public.\"SubDivisions\" as s on w.\"subDivisionId\"=s.\"id\""
		rows, _ = conn.Query(context.Background(), sql)
	} else {
		sql = "SELECT w.\"id\", w.\"name\",  w.\"subDivisionId\" FROM public.\"WorkPosition\" as w inner join public.\"SubDivisions\" as s on w.\"subDivisionId\"=s.\"id\" where w.id=$1"
		rows, _ = conn.Query(context.Background(), sql, id)
	}

	//defer conn.Release()
	wp, err := pgx.CollectRows(rows, pgx.RowToStructByName[CreateWorkPositionResponse])
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	f, errF := json.Marshal(wp)
	if errF != nil {
		ErrResponse(w, errF, operation)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func DeleteWorkPositionV1(w http.ResponseWriter, r *http.Request) {
	//operation:="DeleteWorkPositionV1"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func UpdateWorkPositionV1(w http.ResponseWriter, r *http.Request) {
	//operation:="UpdateWorkPositionV1"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

func CreateWorkerV1(w http.ResponseWriter, r *http.Request) {
	operation := "CreateWorkerV1"
	var rec Worker
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	conn := GetConn()
	var id string
	var sql = "INSERT INTO public.\"Worker\"(\"user_id\", \"workposition_id\") VALUES ($1,$2) RETURNING id"
	err = conn.QueryRow(context.Background(), sql, rec.UserId, rec.WorkPositionId).Scan(&id)
	defer conn.Release()
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	sql = "SELECT \"id\",\"user_id\", \"workposition_id\" FROM public.\"Worker\"  where id=$1"
	rows, _ := conn.Query(context.Background(), sql, id)
	s, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[CreateWorkerResponse])
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	f, errF := json.Marshal(s)
	if errF != nil {
		log.Fatal(errF)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func GetWorkerV1(w http.ResponseWriter, r *http.Request) {
	operation := "GetWorkerV1"
	s := r.URL.Query()["id"]
	var id int
	var sql string
	var rows pgx.Rows
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	conn := GetConn()
	defer conn.Release()
	id, err := strconv.Atoi(strings.Join(s, ""))
	if err != nil {
		sql = "SELECT \"id\",\"user_id\", \"workposition_id\" FROM public.\"Worker\" "
		rows, _ = conn.Query(context.Background(), sql)
	} else {
		sql = "SELECT \"id\",\"user_id\", \"workposition_id\" FROM public.\"Worker\"  where id=$1"
		rows, _ = conn.Query(context.Background(), sql, id)
	}
	defer conn.Release()
	wp, err := pgx.CollectRows(rows, pgx.RowToStructByName[CreateWorkerResponse])
	if err != nil {
		ErrResponse(w, err, operation)
		return
	}
	f, errF := json.Marshal(wp)
	if errF != nil {
		ErrResponse(w, errF, operation)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(f)
}

func UpdateWorkerV1(w http.ResponseWriter, r *http.Request) {
	//operation:="UpdateWorkerV1"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotImplemented)
}

//добавить ролевую модель
