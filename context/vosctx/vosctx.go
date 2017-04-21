package vosctx // import "frizz.io/context/vosctx"

// notest

import (
	"os"

	"context"
)

type Vos interface {
	Getenv(key string) string
	Getwd() (string, error)
	Environ() []string
}

type key int

var fsKey key = 0

// NewContext returns a new Context containing files
func NewContext(ctx context.Context, vos Vos) context.Context {
	return context.WithValue(ctx, fsKey, vos)
}

// FromContext: If there's a fsctx in the ctx, this returns it. If not, we return vfs.OS()
func FromContext(ctx context.Context) Vos {
	if o, ok := ctx.Value(fsKey).(Vos); ok {
		return o
	}
	return &realOs{}
}

func FromContextOrNil(ctx context.Context) Vos {
	if o, ok := ctx.Value(fsKey).(Vos); ok {
		return o
	}
	return nil
}

var _ Vos = (*realOs)(nil)

type realOs struct{}

func (*realOs) Getenv(key string) string {
	return os.Getenv(key)
}

func (*realOs) Getwd() (string, error) {
	return os.Getwd()
}

// Environ returns a copy of strings representing the environment, in the form "key=value".
func (*realOs) Environ() []string {
	return os.Environ()
}
