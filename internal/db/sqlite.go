package db

import (
	"database/sql"

	"os"


	_ "github.com/mattn/go-sqlite3"
)


const (
	dbAdapterName = "sqlite3"
)

// db.Ping() will create dbFile if it does not exist.
func New(dbFilePath string) (*sql.DB, error) {
	_, err := os.Stat(dbFilePath)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(dbAdapterName, dbFilePath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CreateDb(dbFilePath string) (*sql.DB, error) {
	db, err := sql.Open(dbAdapterName, dbFilePath)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

}
