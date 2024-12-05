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
	envPort     = "TODO_PORT"
)

func main() {
	db := db.New()
	defer db.Close()

	repo := repository.New(db)
	migration(repo)

	handlers := handler.New(repo)

	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir(webDir)))

	r.Post("/api/signin", handlers.Signin)
	r.Get("/api/nextdate", handlers.NextDate)
	r.Group(func(r chi.Router) {
		r.Use(handler.AuthMiddleware)

		r.Post("/api/task", handlers.AddTask)
		r.Get("/api/tasks", handlers.GetTasks)
		r.Get("/api/task", handlers.GetTask)
		r.Put("/api/task", handlers.UpdateTask)
		r.Post("/api/task/done", handlers.DoneTask)
		r.Delete("/api/task", handlers.DeleteTask)
	})

	port := os.Getenv(envPort)
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
