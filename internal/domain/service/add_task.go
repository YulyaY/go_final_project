package service

import (
	"errors"
	"log"
	"time"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

func (s *Service) AddTask(t model.Task) (int, error) {
	now := time.Now()
	if t.Title == "" {
		return 0, errTitleIsEmpty
	}

	dateParse, err := time.Parse(FormatDate, t.Date)
	if err != nil {
		return 0, err
	}

	if !dateParse.After(now) && IsDateNotTheSameDayAsNow(now, dateParse) {
		var assignDateBuf string
		if t.Repeat == "" {
			assignDateBuf = now.Format(FormatDate)
		} else {
			nextDate, err := NextDate(now, t.Date, t.Repeat)
			if err != nil {
				return 0, err
			}
			assignDateBuf = nextDate
		}
		t.Date = assignDateBuf
	}

	resultId, err := s.repo.AddTask(t)

	if err != nil {
		log.Printf("service.AddTask db operation error: %v", err)
		return 0, errors.New("AddTask db operation error")
	}

	return resultId, nil
}
