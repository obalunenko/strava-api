package client

import "errors"

// ErrTooManyOptions is returned when too many options are passed to a function
var ErrTooManyOptions = errors.New("too many options")
