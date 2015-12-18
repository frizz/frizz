package envctx // import "kego.io/context/envctx"

import "golang.org/x/net/context"

// Env is the type of value stored in the Contexts.
type Env struct {
	Path    string
	Aliases map[string]string
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
func FromContext(ctx context.Context) (*Env, bool) {
	e, ok := ctx.Value(envKey).(*Env)
	return e, ok
}

var Empty = NewContext(context.Background(), &Env{})
