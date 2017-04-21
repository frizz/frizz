package envctx // import "frizz.io/context/envctx"

// notest

import (
	"context"

	"github.com/dave/kerr"
)

// Env is the type of value stored in the Contexts.
type Env struct {
	Path      string
	Aliases   map[string]string
	Recursive bool
	Hash      uint64
	Dir       string
}

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// envKey is the key for envctx.Env values in Contexts.  It is
// unexported; clients use envctx.NewContext and envctx.FromContext
// instead of using this key directly.
var envKey key = 0

// NewContext returns a new Context that carries value e.
func NewContext(ctx context.Context, e *Env) context.Context {
	if e.Aliases == nil {
		e.Aliases = map[string]string{}
	}
	return context.WithValue(ctx, envKey, e)
}

// FromContext returns the User value stored in ctx, if any.
func FromContext(ctx context.Context) *Env {
	e, ok := ctx.Value(envKey).(*Env)
	if !ok {
		panic(kerr.New("WYDYAGVLCR", "No env in ctx").Error())
	}
	return e
}

func FromContextOrNil(ctx context.Context) *Env {
	e, ok := ctx.Value(envKey).(*Env)
	if ok {
		return e
	}
	return nil
}

var Empty = NewContext(context.Background(), &Env{})

func Dummy(ctx context.Context, path string, aliases map[string]string) context.Context {
	return NewContext(ctx, &Env{Path: path, Aliases: aliases})
}
