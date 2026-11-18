// Package option provides options for parsing environment variables.
package option

import (
	"github.com/obalunenko/getenv/internal"
)

// Option is a contract for EnvParser parameters options.
type Option interface {
	// Apply applies the option to the parameters.
	Apply(params *internal.Parameters)
}

type withSeparator string

func (w withSeparator) Apply(p *internal.Parameters) {
	p.Separator = string(w)
}

// WithSeparator sets the separator of slice values. The default is ",".
func WithSeparator(separator string) Option {
	return withSeparator(separator)
}

type withTimeLayout string

func (w withTimeLayout) Apply(p *internal.Parameters) {
	p.Layout = string(w)
}

// WithTimeLayout sets the layout of time.Time values. The default is time.RFC3339.
func WithTimeLayout(layout string) Option {
	return withTimeLayout(layout)
}
