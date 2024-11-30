package main

import (
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
	dbName      = "scheduler.db"
	webDir      = "./web"
	portDefault = "7540"
	envVarPort  = "TODO_PORT"
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
	r.Get("/api/task", handler.GetTask)
	r.Put("/api/task", handler.UpdateTask)
	r.Post("/api/task/done", handler.DoneTask)
	r.Delete("/api/task", handler.DeleteTask)

	port := os.Getenv(envVarPort)
	if port != "" {
		log.Printf("Server is going to start at http://localhost:%s\n", port)
		if err := http.ListenAndServe(":"+port, r); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("Server is going to start at http://localhost:%s\n", portDefault)
		if err := http.ListenAndServe(":"+portDefault, r); err != nil {
			log.Fatal(err)
		}
	}

}

func migration(rep *repository.Repository) {
	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), dbName)
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
