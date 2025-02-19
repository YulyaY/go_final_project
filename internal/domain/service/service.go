package service

import (
	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

type IRepository interface {
	AddTask(model.Task) (int, error)
	DeleteTask(int) error
	GetTask(int) (model.Task, error)
	GetTasks(model.GetTaskFilter, int) ([]model.Task, error)
	UpdateTask(model.Task) error
}
type Service struct {
	repo IRepository
}

func New(repo IRepository) *Service {
	service := Service{repo: repo}
	return &service
}
