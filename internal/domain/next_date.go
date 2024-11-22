package domain

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	formatDate = "20060102"
)

var errRepeatIsEmpty error = errors.New("repeat is empty")
var errWrongFormat error = errors.New("repeat has wrong format")
var errIsExceed error = errors.New("repeat is maximum permissible interval has been exceeded")

var nextDate time.Time

func NextDate(now time.Time, date string, repeat string) (string, error) {
	dateParse, err := time.Parse(formatDate, date)
	if err != nil {
		return "", err
	}

	repeatSlice := strings.Split(repeat, " ")
	if len(repeatSlice) == 0 || repeat == "" {
		return "", errRepeatIsEmpty
	}

	if repeatSlice[0] == "d" && len(repeatSlice) == 2 {
		repeatDays, err := strconv.Atoi(repeatSlice[1])
		if err != nil {
			return "", errWrongFormat
		} else if repeatDays > 400 {
			return "", errIsExceed
		}
		return nextDateForOptionDay(now, dateParse, repeatDays)
	}

	if repeatSlice[0] == "y" && len(repeatSlice) == 1 {
		return nextDateForOptionYear(now, dateParse)
	}

	return "", errWrongFormat
}

func nextDateForOptionYear(now, dateParse time.Time) (string, error) {
	if dateParse.After(now) || dateParse.Equal(now) {
		nextDate = dateParse.AddDate(1, 0, 0)
	} else if dateParse.YearDay() >= now.YearDay() {
		nextDate = dateParse.AddDate(now.Year()-dateParse.Year(), 0, 0)
	} else {
		nextDate = dateParse.AddDate(now.Year()-dateParse.Year()+1, 0, 0)
	}
	return format(nextDate), nil
}

func nextDateForOptionDay(now, dateParse time.Time, repeatDays int) (string, error) {
	if dateParse.After(now) || dateParse.Equal(now) {
		nextDate = dateParse.AddDate(0, 0, repeatDays)
	} else {
		nextDate = dateParse
		for nextDate.After(now) {
			nextDate = nextDate.AddDate(0, 0, repeatDays)
		}
	}
	return format(nextDate), nil
}

func format(nextDate time.Time) string {
	return nextDate.Format(formatDate)
}
