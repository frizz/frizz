package cmdctx // import "kego.io/context/cmdctx"

// ke: {"package": {"notest": true}}

import (
	"fmt"

	"context"

	"github.com/dave/kerr"
)

// Env is the type of value stored in the Contexts.
type Cmd struct {
	Edit     bool
	Validate bool
	Update   bool
	Log      bool
	Debug    bool
	Port     int
}

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// envKey is the key for cmdctx.Env values in Contexts.  It is
// unexported; clients use cmdctx.NewContext and cmdctx.FromContext
// instead of using this key directly.
var cmdKey key = 0

// NewContext returns a new Context that carries value u.
func NewContext(ctx context.Context, e *Cmd) context.Context {
	return context.WithValue(ctx, cmdKey, e)
}

// FromContext returns the User value stored in ctx, if any.
func FromContext(ctx context.Context) *Cmd {
	e, ok := ctx.Value(cmdKey).(*Cmd)
	if !ok {
		panic(kerr.New("OQVLBQFQJW", "No cmd in ctx").Error())
	}
	return e
}

func FromContextOrNil(ctx context.Context) *Cmd {
	e, ok := ctx.Value(cmdKey).(*Cmd)
	if ok {
		return e
	}
	return nil
}

func (c *Cmd) Printf(format string, a ...interface{}) (n int, err error) {
	if c.Log {
		return fmt.Printf(format, a...)
	}
	return 0, nil
}

func (c *Cmd) Println(a ...interface{}) (n int, err error) {
	if c.Log {
		return fmt.Println(a...)
	}
	return 0, nil
}

func (c *Cmd) Print(a ...interface{}) (n int, err error) {
	if c.Log {
		return fmt.Print(a...)
	}
	return 0, nil
}
