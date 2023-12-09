package storage

import "errors"

// Error types
var (
	ErrNotFound   = errors.New("NotFoundError")
	ErrEmptyInput = errors.New("EmptyInputError")
	ErrBadInput   = errors.New("BadInputError")
)
