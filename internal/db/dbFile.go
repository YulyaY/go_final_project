package db

import (
	"log"
	"os"
)

const (
	dbNameDefault = "scheduler.db"

	envVarDbFile = "TODO_DBFILE"
)

func GetDbFile() string {
	dbFileEnv := os.Getenv(envVarDbFile)
	if dbFileEnv != "" {
		_, err := os.Stat(dbFileEnv)
		if err != nil {
			log.Fatal(err)
		}
		return dbFileEnv
	}

	return dbNameDefault
}
