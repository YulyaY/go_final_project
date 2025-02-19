package service

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	FormatDate = "20060102"
)

var nextDate time.Time

func NextDate(now time.Time, date string, repeat string) (string, error) {
	dateParse, err := time.Parse(FormatDate, date)
	if err != nil {
		return "", err
	}

	repeatSlice := strings.Split(repeat, " ")
	if len(repeatSlice) == 0 || repeat == "" {
		return "", ErrRepeatIsEmpty
	}

	if repeatSlice[0] == "d" && len(repeatSlice) == 2 {
		repeatDays, err := strconv.Atoi(repeatSlice[1])
		if err != nil {
			return "", ErrWrongFormat
		}
		if repeatDays > 400 {
			return "", errIsExceed
		}
		return nextDateForOptionDay(now, dateParse, repeatDays)
	}

	if repeatSlice[0] == "y" && len(repeatSlice) == 1 {
		return nextDateForOptionYear(now, dateParse)
	}

	if repeatSlice[0] == "w" && len(repeatSlice) == 2 {
		repeatDaySlice := strings.Split(repeatSlice[1], ",")
		if len(repeatDaySlice) == 0 || repeatSlice[1] == "" {
			return "", ErrWrongFormat
		}

		repeatDaySliceInt := make([]int, 0)
		for _, day := range repeatDaySlice {
			weekDay, err := strconv.Atoi(day)
			if err != nil {
				return "", ErrWrongFormat
			}
			if weekDay < 1 || weekDay > 7 {
				return "", errIsInvalidValueDayOfWeek
			}
			repeatDaySliceInt = append(repeatDaySliceInt, weekDay)
		}
		return nextDateForOptionWeek(now, dateParse, repeatDaySliceInt)
	}

	if repeatSlice[0] == "m" && (len(repeatSlice) == 2 || len(repeatSlice) == 3) {
		repeatDaySlice := strings.Split(repeatSlice[1], ",")
		if len(repeatDaySlice) == 0 || repeatSlice[1] == "" {
			return "", ErrWrongFormat
		}
		repeatDaySliceInt := make([]int, 0)
		for _, day := range repeatDaySlice {
			monthDay, err := strconv.Atoi(day)
			if monthDay < 1 || monthDay > 31 {
				if !(monthDay == -1 || monthDay == -2) {
					return "", errIsInvalidValueDayOfMonth
				}
			}
			if err != nil {
				return "", ErrWrongFormat
			}
			repeatDaySliceInt = append(repeatDaySliceInt, monthDay)
		}

		if len(repeatSlice) == 3 {
			repeatMonthSliceInt := make([]int, 0)
			repeatMonthSlice := strings.Split(repeatSlice[2], ",")
			for _, month := range repeatMonthSlice {
				monthInt, err := strconv.Atoi(month)
				if monthInt < 1 || monthInt > 12 {
					return "", errIsInvalidValueMonth
				}
				if err != nil {
					return "", ErrWrongFormat
				}
				repeatMonthSliceInt = append(repeatMonthSliceInt, monthInt)
			}
			return nextDateForOptionMonth(now, dateParse, repeatDaySliceInt, repeatMonthSliceInt)
		}
		return nextDateForOptionMonthOnlyDay(now, dateParse, repeatDaySliceInt)

	}

	return "", ErrWrongFormat
}

func nextDateForOptionYear(now, dateParse time.Time) (string, error) {
	if dateParse.After(now) || !IsDateNotTheSameDayAsNow(now, dateParse) {
		nextDate = dateParse.AddDate(1, 0, 0)
	} else if dateParse.YearDay() >= now.YearDay() {
		nextDate = dateParse.AddDate(now.Year()-dateParse.Year(), 0, 0)
	} else {
		nextDate = dateParse.AddDate(now.Year()-dateParse.Year()+1, 0, 0)
	}
	return Format(nextDate), nil
}

func nextDateForOptionDay(now, dateParse time.Time, repeatDays int) (string, error) {
	if dateParse.After(now) || !IsDateNotTheSameDayAsNow(now, dateParse) {
		nextDate = dateParse.AddDate(0, 0, repeatDays)
	} else {
		nextDate = dateParse
		for nextDate.Before(now) {
			nextDate = nextDate.AddDate(0, 0, repeatDays)
		}
	}
	return Format(nextDate), nil
}

func Format(nextDate time.Time) string {
	return nextDate.Format(FormatDate)
}

func IsDateNotTheSameDayAsNow(now, dateParse time.Time) bool {
	yearNow, monthNow, dayNow := now.Date()
	yearDate, monthDate, dayDate := dateParse.Date()
	if yearNow == yearDate && monthNow == monthDate && dayNow == dayDate {
		return false
	}
	return true
}

func nextDateForOptionWeek(now, dateParse time.Time, repeatDays []int) (string, error) {
	sort.Ints(repeatDays)
	var nextDate time.Time
	if dateParse.After(now) || !IsDateNotTheSameDayAsNow(now, dateParse) {
		dateParseWeekDay := int(dateParse.Weekday())
		if dateParseWeekDay == 0 {
			nextDate = dateParse.AddDate(0, 0, repeatDays[0])
			return Format(nextDate), nil
		}
		for _, day := range repeatDays {
			if dateParseWeekDay < day {
				nextDate = dateParse.AddDate(0, 0, day-dateParseWeekDay)
				return Format(nextDate), nil
			}
		}
		nextDate = dateParse.AddDate(0, 0, 7-dateParseWeekDay+repeatDays[0])
		return Format(nextDate), nil
	} else {
		nextDate = now
		nextDateWeekDay := int(nextDate.Weekday())
		if nextDateWeekDay == 0 {
			nextDate = nextDate.AddDate(0, 0, repeatDays[0])
			return Format(nextDate), nil
		}
		nextDateWeekDayBeforeDay := false
		var dayOne int
		for _, day := range repeatDays {
			if nextDateWeekDay < day {
				nextDateWeekDayBeforeDay = true
				dayOne = day
				break
			}
		}
		if nextDateWeekDayBeforeDay {
			nextDate = nextDate.AddDate(0, 0, dayOne-nextDateWeekDay)
		} else {
			nextDate = nextDate.AddDate(0, 0, 7-nextDateWeekDay+repeatDays[0])
		}
		return Format(nextDate), nil
	}
}

func nextDateForOptionMonth(now, dateParse time.Time, repeatDays, repeatMonth []int) (string, error) {
	sort.Ints(repeatDays)
	sort.Ints(repeatMonth)
	dateDay := dateParse.Day()
	dateMonth := dateParse.Month()
	dateYear := dateParse.Year()
	var nextDate time.Time
	if dateParse.After(now) || !IsDateNotTheSameDayAsNow(now, dateParse) {
		for _, month := range repeatMonth {
			if dateMonth < time.Month(month) {
				for _, day := range repeatDays {
					if day > 0 {
						nextDate = time.Date(dateYear, time.Month(month), repeatDays[0], 0, 0, 0, 0, time.UTC)
						if nextDate.Day() == repeatDays[0] {
							return Format(nextDate), nil
						}
					}
				}
				nextDate = time.Date(dateYear, time.Month(month)+1, repeatDays[0]+1, 0, 0, 0, 0, time.UTC)
				return Format(nextDate), nil
			}
			if dateMonth == time.Month(month) {
				for _, day := range repeatDays {
					if dateDay < day {
						nextDate = time.Date(dateYear, time.Month(month), day, 0, 0, 0, 0, time.UTC)
						if nextDate.Day() == day {
							return Format(nextDate), nil
						}
					}
				}
				if repeatDays[0] < 0 {
					nextDate = time.Date(dateYear, time.Month(month)+1, repeatDays[0]+1, 0, 0, 0, 0, time.UTC)
					return Format(nextDate), nil
				}
			}
		}
		if repeatDays[0] < 0 {
			nextDate = time.Date(dateYear+1, time.Month(repeatMonth[0])+1, repeatDays[0]+1, 0, 0, 0, 0, time.UTC)
			return Format(nextDate), nil
		}
		nextDate = time.Date(dateYear+1, time.Month(repeatMonth[0]), repeatDays[0], 0, 0, 0, 0, time.UTC)
		if nextDate.Day() == repeatDays[0] {
			return Format(nextDate), nil
		} else {
			return "", ErrWrongFormat
		}

	}
	nextDate = now.AddDate(0, 0, -1)
	nextDateDay := nextDate.Day()
	nextDateMonth := nextDate.Month()
	nextDateYear := nextDate.Year()
	for _, month := range repeatMonth {
		if nextDateMonth < time.Month(month) {
			for _, day := range repeatDays {
				if day > 0 {
					nextDate = time.Date(nextDateYear, time.Month(month), repeatDays[0], 0, 0, 0, 0, time.UTC)
					if nextDate.Day() == repeatDays[0] {
						return Format(nextDate), nil
					}
				}
			}
			nextDate = time.Date(nextDateYear, time.Month(month)+1, repeatDays[0]+1, 0, 0, 0, 0, time.UTC)
			return Format(nextDate), nil
		}
		if nextDateMonth == time.Month(month) {
			for _, day := range repeatDays {
				if nextDateDay < day {
					nextDate = time.Date(nextDateYear, time.Month(month), day, 0, 0, 0, 0, time.UTC)
					if nextDate.Day() == day {
						return Format(nextDate), nil
					}
				}
			}
			if repeatDays[0] < 0 {
				nextDate = time.Date(nextDateYear, time.Month(month)+1, repeatDays[0]+1, 0, 0, 0, 0, time.UTC)
				return Format(nextDate), nil
			}
		}
	}
	if repeatDays[0] < 0 {
		nextDate = time.Date(nextDateYear+1, time.Month(repeatMonth[0])+1, repeatDays[0]+1, 0, 0, 0, 0, time.UTC)
		return Format(nextDate), nil
	}
	nextDate = time.Date(nextDateYear+1, time.Month(repeatMonth[0]), repeatDays[0], 0, 0, 0, 0, nil)
	if nextDate.Day() == repeatDays[0] {
		return Format(nextDate), nil
	} else {
		return "", ErrWrongFormat
	}

}

func nextDateForOptionMonthOnlyDay(now, dateParse time.Time, repeatDays []int) (string, error) {
	sort.Ints(repeatDays)
	dateDay := dateParse.Day()
	dateMonth := dateParse.Month()
	dateYear := dateParse.Year()
	var nextDate time.Time
	if dateParse.After(now) || !IsDateNotTheSameDayAsNow(now, dateParse) {
		for _, day := range repeatDays {
			if dateDay < day {
				nextDate = time.Date(dateYear, dateMonth, day, 0, 0, 0, 0, time.UTC)
				if nextDate.Day() == day {
					return Format(nextDate), nil
				}
			}
		}
		if repeatDays[0] < 0 {
			nextDate = time.Date(dateYear, dateMonth+1, repeatDays[0]+1, 0, 0, 0, 0, time.UTC)
			return Format(nextDate), nil
		} else {
			nextDate = time.Date(dateYear, dateMonth+1, repeatDays[0], 0, 0, 0, 0, time.UTC)
			if nextDate.Day() == repeatDays[0] {
				return Format(nextDate), nil
			} else {
				return "", ErrWrongFormat
			}
		}

	}
	nextDate = now.AddDate(0, 0, -1)
	nextDateDay := nextDate.Day()
	nextDateMonth := nextDate.Month()
	nextDateYear := nextDate.Year()
	for _, day := range repeatDays {
		if nextDateDay < day {
			nextDate = time.Date(nextDateYear, nextDateMonth, repeatDays[0], 0, 0, 0, 0, time.UTC)
			if nextDate.Day() == repeatDays[0] {
				return Format(nextDate), nil
			}
		}
	}
	if repeatDays[0] < 0 {
		nextDate = time.Date(nextDateYear, nextDateMonth+1, repeatDays[0]+1, 0, 0, 0, 0, time.UTC)
		return Format(nextDate), nil
	} else {
		nextDate = time.Date(nextDateYear, nextDateMonth+1, repeatDays[0], 0, 0, 0, 0, time.UTC)
		if nextDate.Day() == repeatDays[0] {
			return Format(nextDate), nil
		} else {
			return "", ErrWrongFormat
		}
	}
}
