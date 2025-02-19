package handler

import "errors"

var errIdIsEmpty error = errors.New("id is empty")
var errAuthRequired error = errors.New("authentification required")
var errInvalidPassword error = errors.New("invalid password")
