package database

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	if err != nil {
		log.Panic(err)
	}
	return db
}
