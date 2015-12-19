package wgctx // import "kego.io/context/wgctx"

import (
	"os"
	"time"

	"fmt"

	"sync"

	"golang.org/x/net/context"
)

// key is an unexported type for keys defined in this package.
// This prevents collisions with keys defined in other packages.
type key int

// wgKey is the key for wgctx values in Contexts.  It is unexported;
// clients use wgctx.NewContext and wgctx.FromContext instead of
// using this key directly.
var wgKey key = 0

// NewContext returns a new Context that carries value e.
func NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, wgKey, &sync.WaitGroup{})
}

// FromContext returns the User value stored in ctx, if any.
func FromContext(ctx context.Context) *sync.WaitGroup {
	wg, ok := ctx.Value(wgKey).(*sync.WaitGroup)
	if !ok {
		panic("No wg in ctx")
	}
	return wg
}

func WaitAndExit(ctx context.Context, code int) {
	wg := FromContext(ctx)

	c := make(chan struct{}, 1)
	go func() {
		defer close(c)
		wg.Wait()
	}()
	quartersecond := time.After(time.Millisecond * 250)
	twoseconds := time.After(time.Second * 5)
	for {
		select {
		case <-c:
			os.Exit(code)
		case <-quartersecond:
			fmt.Println("Waiting for long-running processes to finish...")
		case <-twoseconds:
			fmt.Println("Timed out waiting for long-running processes to finish... Exiting now.")
			os.Exit(code)
		}
	}

}
