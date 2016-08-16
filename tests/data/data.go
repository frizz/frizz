package data

import (
	"context"

	"kego.io/system"
)

type Face interface {
	Face() string
}

func (f *Facea) Face() string {
	return f.A.String()
}

func (f *Faceb) Face() string {
	return f.B.String()
}

func (f *Facea) GetString(ctx context.Context) *system.String {
	return f.A
}

func (m *Multi) GetString(ctx context.Context) *system.String {
	return m.Ss
}
