package main

import (
	"flag"
	"log"

	"github.com/YulyaY/go_final_project.git/internal/config"
	"github.com/YulyaY/go_final_project.git/internal/db"
	"github.com/pressly/goose"
)

const (
	defaultDirectory = "./migrations"
	usage            = "the path of directory migrations"
)

func main() {
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		log.Fatalf("Can not set config: '%s'", err.Error())
	}

	//db, err := db.New(appConfig.DbFilePath)
	db, err := db.NewPosgres(appConfig)
	if err != nil {
		log.Fatalf("Can not init db connect to datebase: '%s'", err.Error())
	}
	defer db.Close()

	// if err := goose.SetDialect("sqlite3"); err != nil {
	// 	log.Fatalf("Goose can not init connect to datebase: '%s'", err.Error())
	// }

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Goose can not init connect to datebase: '%s'", err.Error())
	}

	var dirMigrations string
	flag.StringVar(&dirMigrations, "dir", defaultDirectory, usage)
	log.Println(dirMigrations)

	if err := goose.Up(db, dirMigrations); err != nil {
		log.Fatalf("Goose can not migrate datebase: '%s'", err.Error())
	}
}
