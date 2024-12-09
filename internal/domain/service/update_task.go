package service

import (
	"strconv"
	"time"

	"github.com/YulyaY/go_final_project.git/internal/domain/model"
)

func (s *Service) UpdateTask(t model.Task) error {
	now := time.Now()

	if t.Id == "" {
		return errIdIsEmpty
	}
	if _, err := strconv.Atoi(t.Id); err != nil {
		return errIdHasWrongFormat
	}
	if t.Title == "" {
		return errTitleIsEmpty
	}

	dateParse, err := time.Parse(FormatDate, t.Date)
	if err != nil {
		return err
	}

	if !dateParse.After(now) && IsDateNotTheSameDayAsNow(now, dateParse) {
		var assignDateBuf string
		if t.Repeat == "" {
			assignDateBuf = now.Format(FormatDate)
		} else {
			nextDate, err := NextDate(now, t.Date, t.Repeat)
			if err != nil {
				return err
			}
			assignDateBuf = nextDate
		}
		t.Date = assignDateBuf
	}
	errOfUpdate := s.repo.UpdateTask(t)
	if errOfUpdate != nil {
		return errOfUpdate
	}
	return nil
}
