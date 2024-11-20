package domain

import (
	"errors"
	"time"
)

var errRepeatIsEmpty error = errors.New("Repeat is empty")

func NextDate(now time.Time, date string, repeat string) (string, error) {
	if dateParse, err := time.Parse("20060102", date); err != nil {
		return "", err
	}
	if repeat == "" {
		return "", errRepeatIsEmpty
	}

	var nextDate string
	return nextDate, nil
}
