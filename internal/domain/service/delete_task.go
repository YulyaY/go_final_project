package service

func (s *Service) DeleteTask(id int) error {
	err := s.repo.DeleteTask(id)
	return err
}
