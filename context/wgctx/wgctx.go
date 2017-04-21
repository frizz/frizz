package wgctx // import "frizz.io/context/wgctx"

// notest

import (
	"os"
	"time"

	"fmt"

	"sync"

	"context"

	"github.com/dave/kerr"
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
		panic(kerr.New("UNYGADKEGY", "No wg in ctx").Error())
	}
	return wg
}

func FromContextOrNil(ctx context.Context) *sync.WaitGroup {
	wg, ok := ctx.Value(wgKey).(*sync.WaitGroup)
	if !ok {
		return nil
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

	go func() {
		// only show the waiting message after 250ms
		<-time.After(time.Millisecond * 250)
		fmt.Println("Waiting for long-running processes to finish...")
	}()

	select {
	case <-c:
		os.Exit(code)
	case <-time.After(time.Second * 5):
		fmt.Println("Timed out waiting for long-running processes to finish... Exiting now.")
		os.Exit(code)
	}

}

func Add(ctx context.Context, ref string) {
	wg := FromContext(ctx)
	wg.Add(1)
	//fmt.Println("Add:", ref)
}

func Done(ctx context.Context, ref string) {
	wg := FromContext(ctx)
	wg.Done()
	//fmt.Println("Done:", ref)
}
