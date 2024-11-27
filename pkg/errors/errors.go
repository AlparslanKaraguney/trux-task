package errors

import "errors"

var (
	ErrNotFound       = errors.New("record not found")
	ErrAlreadyExists  = errors.New("record already exists")
	ErrInvalidOrderBy = errors.New("invalid order by")
	ErrInvalidFilter  = errors.New("invalid filter")
)
