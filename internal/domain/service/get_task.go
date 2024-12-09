package service

import "github.com/YulyaY/go_final_project.git/internal/domain/model"

func (s *Service) GetTask(id int) (model.Task, error) {
	t, err := s.repo.GetTask(id)
	return t, err
}
