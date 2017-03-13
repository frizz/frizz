package clientctx

import (
	"sync"

	"context"

	"github.com/dave/kerr"
	"kego.io/editor/client/editable"
)

type ctxKeyType int

var ctxKey ctxKeyType = 0

func NewContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey, &EditorCache{m: map[string]editable.Editable{}})
}

func FromContext(ctx context.Context) *EditorCache {
	ec, ok := ctx.Value(ctxKey).(*EditorCache)
	if !ok {
		panic(kerr.New("BIUVXISEMA", "No editors in ctx").Error())
	}
	return ec
}

func FromContextOrNil(ctx context.Context) *EditorCache {
	ec, ok := ctx.Value(ctxKey).(*EditorCache)
	if ok {
		return ec
	}
	return nil
}

type EditorCache struct {
	sync.RWMutex
	m map[string]editable.Editable
}

func (c *EditorCache) Get(path string) (editable.Editable, bool) {
	c.RLock()
	defer c.RUnlock()
	e, ok := c.m[path]
	return e, ok
}

func (c *EditorCache) Set(path string, e editable.Editable) {
	c.Lock()
	defer c.Unlock()
	c.m[path] = e
}
