package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func New() *sql.DB {
	dbFile := GetDbFile()
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal("init db", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("ping db", err)
	}

	return db

}
