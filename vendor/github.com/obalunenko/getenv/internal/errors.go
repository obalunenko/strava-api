package internal

import (
	"errors"
)

var (
	// ErrNotSet is an error that is returned when the environment variable is not set.
	ErrNotSet = errors.New("not set")
	// ErrInvalidValue is an error that is returned when the environment variable is not valid.
	ErrInvalidValue = errors.New("invalid value")

	errEmptySeparator = errors.New("empty separator")
)
