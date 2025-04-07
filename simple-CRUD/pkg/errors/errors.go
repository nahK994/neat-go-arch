package errors

import "errors"

var ErrEmailAlreadyExists = errors.New("email already exists")
var ErrEmailNotExists = errors.New("email not exists")
var ErrUnauthorized = errors.New("unauthorized")
