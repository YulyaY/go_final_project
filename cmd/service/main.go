package main

import (
	"log"
	"net/http"

	"github.com/YulyaY/go_final_project.git/internal/config"
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
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		log.Fatalf("Can not set config: '%s'", err.Error())
	}

	db, err := db.New(appConfig.DbFilePath)
	if err != nil {
		log.Fatalf("Can not init db connect or create datebase: '%s'", err.Error())
	}
	defer db.Close()

	repo := repository.New(db)
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

	log.Printf("Server is going to start at 0.0.0.0:%s\n", appConfig.Port)
	if err := http.ListenAndServe(":"+appConfig.Port, r); err != nil {
		log.Fatal(err)
	}

}
