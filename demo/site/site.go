//go:generate ke
package site

import "context"

// ke: {"package": {"notest": true}}

type Section interface {
	Template() string
}

func (v *Body) Template() string {
	return "body"
}

func (v *Body) Label(ctx context.Context) string {
	return "body"
}

func (v *Columns) Template() string {
	return "columns"
}

func (v *Columns) Label(ctx context.Context) string {
	return "columns"
}

func (v *Hero) Template() string {
	return "hero"
}

func (v *Hero) Label(ctx context.Context) string {
	return "hero"
}
