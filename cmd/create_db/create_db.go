package main

import (
	"log"

	"github.com/YulyaY/go_final_project.git/internal/config"
	"github.com/YulyaY/go_final_project.git/internal/db"
)

func main() {
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		log.Fatalf("Can not set config: '%s'", err.Error())
	}

	db, err := db.CreateDb(appConfig.DbFilePath)
	if err != nil {
		log.Fatalf("Can not create datebase: '%s'", err.Error())
	}
	defer db.Close()
}
