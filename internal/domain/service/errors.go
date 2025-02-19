package service

import "errors"

var ErrRepeatIsEmpty error = errors.New("repeat is empty")
var ErrWrongFormat error = errors.New("repeat has wrong format")
var ErrIdIsNotFound error = errors.New("select by id: id is not found")
var ErrRecordDoesNotExists error = errors.New("record does not exists")

var errIsExceed error = errors.New("repeat is maximum permissible interval has been exceeded")
var errIsInvalidValueDayOfWeek error = errors.New("invalid value for day of the week")
var errIsInvalidValueDayOfMonth error = errors.New("invalid value for day of the month")
var errIsInvalidValueMonth error = errors.New("invalid value for month")
var errTitleIsEmpty error = errors.New("title is empty")
var errIdIsEmpty error = errors.New("id is empty")
var errIdHasWrongFormat error = errors.New("id has wrong format")
