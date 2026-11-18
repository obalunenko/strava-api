// Package getenv provides a simple way to get environment variables.
// It's type-safe and supports built-in types and slices of them.
// Note: empty environment variable values are treated as "not set".
//
// Slices are split by "," and time.Time values are parsed with the time.RFC3339 layout,
// unless option.WithSeparator or option.WithTimeLayout sets another one.
//
// Types supported:
// - string
// - []string
// - int
// - []int
// - int8
// - []int8
// - int16
// - []int16
// - int32
// - []int32
// - int64
// - []int64
// - uint8
// - []uint8
// - uint16
// - []uint16
// - uint64
// - []uint64
// - uint
// - []uint
// - uintptr
// - []uintptr
// - uint32
// - []uint32
// - float32
// - []float32
// - float64
// - []float64
// - time.Time
// - []time.Time
// - time.Duration
// - []time.Duration
// - bool
// - []bool
// - url.URL
// - []url.URL
// - net.IP
// - []net.IP
// - netip.Addr
// - []netip.Addr
// - netip.Prefix
// - []netip.Prefix
// - net.HardwareAddr
// - []net.HardwareAddr
// - complex64
// - []complex64
// - complex128
// - []complex128
package getenv

import (
	"fmt"
	"os"
	"time"

	"github.com/obalunenko/getenv/internal"
	"github.com/obalunenko/getenv/option"
)

var (
	// ErrNotSet is an error that is returned when the environment variable is not set.
	ErrNotSet = internal.ErrNotSet
	// ErrInvalidValue is an error that is returned when the environment variable is not valid.
	ErrInvalidValue = internal.ErrInvalidValue
)

// Defaults for the options that were not passed.
const (
	defaultSeparator  = ","
	defaultTimeLayout = time.RFC3339
)

// Parsable is a constraint for the types that Env and EnvOrDefault can parse.
// The list of supported types is in the package documentation.
type Parsable = internal.EnvParsable

// Env retrieves the value of the environment variable named by the key.
// If the variable is present in the environment, the value will be parsed and returned.
// Otherwise, an error will be returned.
func Env[T Parsable](key string, options ...option.Option) (T, error) {
	var zero T

	raw, ok := os.LookupEnv(key)
	if !ok || raw == "" {
		return zero, fmt.Errorf("failed to get environment variable[%s]: %w", key, ErrNotSet)
	}

	val, err := internal.Parse[T](raw, newParseParams(options))
	if err != nil {
		return zero, fmt.Errorf("failed to parse environment variable[%s]: %w", key, err)
	}

	return val, nil
}

// EnvOrDefault retrieves the value of the environment variable named by the key.
// If the variable is set and its value can be parsed, the parsed value is returned.
// Otherwise, defaultVal is returned: both when the variable is not set (or empty)
// and when its value is invalid. Use Env to tell these cases apart.
func EnvOrDefault[T Parsable](key string, defaultVal T, options ...option.Option) T {
	val, err := Env[T](key, options...)
	if err != nil {
		return defaultVal
	}

	return val
}

// newParseParams creates new parameters from the defaults and options.
func newParseParams(opts []option.Option) internal.Parameters {
	p := internal.Parameters{
		Separator: defaultSeparator,
		Layout:    defaultTimeLayout,
	}

	for _, opt := range opts {
		opt.Apply(&p)
	}

	return p
}
