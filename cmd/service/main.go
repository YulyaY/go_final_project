package main

import (
	"log"
	"net/http"

	"github.com/YulyaY/go_final_project.git/internal/app"
	"github.com/YulyaY/go_final_project.git/internal/config"
	"github.com/YulyaY/go_final_project.git/internal/db"
	"github.com/YulyaY/go_final_project.git/internal/domain/service"

	"github.com/YulyaY/go_final_project.git/internal/handler"
	"github.com/YulyaY/go_final_project.git/internal/repository"
	"github.com/go-chi/chi"
	_ "modernc.org/sqlite"
)

const (
	webDir = "./web"
)

func main() {

	appConfig, err := config.LoadAppConfig()
	if err != nil {
		log.Fatalf("Can not set config: '%s'", err.Error())
	}

	appSetting := app.AppSettings{
		IsAuthentificationControlSwitchedOn: appConfig.IsPasswordSet(),
	}

	//db, err := db.New(appConfig.DbFilePath)
	db, err := db.NewPosgres(appConfig)
	if err != nil {
		log.Fatalf("Can not init db connect or create datebase: '%s'", err.Error())
	}
	defer db.Close()

	repo := repository.New(db)
	svc := service.New(repo)
	handlers := handler.New(svc, appConfig)

	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir(webDir)))

	r.Post("/api/signin", handlers.Signin)
	r.Get("/api/nextdate", handlers.NextDate)
	r.Group(func(r chi.Router) {
		r.Use(handler.BuildAuthMiddleware(appConfig, appSetting))

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
