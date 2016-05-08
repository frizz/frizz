package ctests

import (
	"kego.io/editor/client/stores"
	"kego.io/process/tests"
)

type ContextBuilder struct {
	*tests.ContextBuilder
}

func New() *ContextBuilder {
	return Wrap(tests.New())
}

func Wrap(cb *tests.ContextBuilder) *ContextBuilder {
	return &ContextBuilder{cb}
}

func (cb *ContextBuilder) App() *ContextBuilder {
	app := &stores.App{
		Fail: make(chan error),
	}
	ctx := stores.NewContext(cb.Ctx(), app)
	cb.SetCtx(ctx)

	app.Init(cb.Ctx())
	return cb
}
