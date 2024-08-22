package database

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	log.Print(os.Getenv("POSTGRES_URI"), "\n")
	if err != nil {
		log.Panic("Can not open db")
	}
	err = db.Ping()
	if err != nil {
		log.Panic("Can not ping db")
	}
	return db
}
