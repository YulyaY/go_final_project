package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/YulyaY/go_final_project.git/internal/db"
	"github.com/YulyaY/go_final_project.git/internal/handler"
	"github.com/YulyaY/go_final_project.git/internal/repository"
	"github.com/go-chi/chi"
	_ "modernc.org/sqlite"
)

const (
	dbName = "scheduler.db"
	webDir = "./web"
)

func main() {
	db := db.New()
	defer db.Close()

	repo := repository.New(db)
	migration(repo)

	handler := handler.New(repo)

	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir(webDir)))

	r.Get("/api/nextdate", handler.NextDate)
	r.Post("/api/task", handler.AddTask)
	r.Get("/api/tasks", handler.GetTasks)

	fmt.Println("Server is going to start")
	port := os.Getenv("TODO_PORT")
	if port != "" {
		if err := http.ListenAndServe(":"+port, r); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := http.ListenAndServe(":7540", r); err != nil {
			log.Fatal(err)
		}
	}
}

func migration(rep *repository.Repository) {
	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), "scheduler.db")
	_, err = os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	if install {
		if err := rep.CreateScheduler(); err != nil {
			log.Fatal(err)
		}
	}
}
