package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func OpenConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=root password=root dbname=godb sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	return db, err
}
