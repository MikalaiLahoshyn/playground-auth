package models

import "errors"

var (
	ErrForbidden   = errors.New("forbidden")
	ErrWrongAction = errors.New("wrong action")
	ErrNotFound    = errors.New("not found")
	ErrInternal    = errors.New("internal server error")
	ErrBadRequest  = errors.New("bad request")
)
