package store

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func СreateDB(connectStr string, connLife time.Duration, maxIdle, poolSize int) (db *sql.DB, err error) {
	db, err = sql.Open("postgres", connectStr)
	if err != nil {
		log.Fatalf("opening DB connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("establishing DB connection; %v", err)
	}

	db.SetConnMaxLifetime(connLife)
	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(poolSize)

	return
}
