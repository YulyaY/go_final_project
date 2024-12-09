package service

import (
	"github.com/YulyaY/go_final_project.git/internal/domain"
	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

func (s *Service) GetTasks(taskFilter model.GetTaskFilter) ([]model.Task, error) {
	tasks, err := s.repo.GetTasks(taskFilter, domain.Limit)
	return tasks, err
}
