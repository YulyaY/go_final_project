package service

import (
	"time"
)

func (s *Service) DoneTask(id int) error {
	now := time.Now()
	t, err := s.repo.GetTask(id)
	if err != nil {
		return err
	}

	if t.Repeat == "" {
		err := s.repo.DeleteTask(id)
		if err != nil {
			return err
		}
		return nil
	}

	nextDate, err := NextDate(now, t.Date, t.Repeat)
	if err == ErrWrongFormat {
		err := s.repo.DeleteTask(id)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}
	t.Date = nextDate
	err = s.repo.UpdateTask(t)
	if err != nil {
		return err
	}
	return nil
}
