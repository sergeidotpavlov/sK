package mware

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func initPool() error {
	dbpath, exists := os.LookupEnv("DATABASE_URL")
	// dbpath = "postgresql:///systemK?host=localhost&port=5432&user=user1&password=1user"
	if exists {
		config, err := pgxpool.ParseConfig(dbpath)
		if err != nil {
			log.Fatal("DATABASE_URL is nil")
		}
		log.Println("Started!")
		config.MinConns = 1
		pool, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			return err
		}
		log.Println("pgxpool configured!")
	}
	return nil
}

func SetPool1(p *pgxpool.Pool) {
	pool = p
	log.Println(pool.Config().ConnString())
}

func GetConn() *pgxpool.Conn {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("Unable to acquire a database connection: %v\n", err)
	}
	return conn
}
