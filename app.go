/*
* SystemK
*
* API <br/>  Documentation:    <ul>      <li><a href=\"https://atlassian.net/wiki/</a></li>      <li><a href=\"https://git.net/f\">Git</a></li>      <li><a href=\"https://humansinc.atlassian.net/browse/ORM\">Jira</a></li>    </ul>  <br/>  <a href=\"git .yaml\">API Artifact</a><br/>
*
* API version: 1.0.0
 */
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	sw "github.com/sergeidotpavlov/systemK/user-srv/go"
)

func main() {
	dbpath, exists := os.LookupEnv("DATABASE_URL")
	//	dbpath = "postgresql:///systemK?host=localhost&port=5432&user=user1&password=1user"
	if exists {
		config, err := pgxpool.ParseConfig(dbpath)
		if err != nil {
			log.Fatal("DATABASE_URL is nil")
		}
		log.Println("Started!")
		config.MinConns = 1
		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("pgxpool configured!")
		// Передаем пул подключений через контекст.
		//ctx := context.WithValue(context.Background(), "pgxpool", pool)
		if err := sw.initPool(); err != nil {
			log.Println(err)
		}
		defer pool.Close()
		router := sw.NewRouter()
		srv := &http.Server{
			Addr:         "0.0.0.0:8080",
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Second * 60,
			Handler:      router,
		}
		//log.Fatal(http.ListenAndServe(":8080", router))
		go func() {
			if err := srv.ListenAndServe(); err != nil {
				log.Println(err)
			}
		}()
		// shutdowns when quit via SIGINT (Ctrl+C)
		// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/)
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		srv.Shutdown(context.Background())
		log.Println("shutting down")
		os.Exit(0)
	}
}
