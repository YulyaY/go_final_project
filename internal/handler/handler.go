package handler

import "github.com/YulyaY/go_final_project.git/internal/repository"

type Handler struct {
	repo *repository.Repository
}

func New(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}
