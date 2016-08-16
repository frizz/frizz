package stores

import (
	"context"

	"github.com/davelondon/kerr"
)

type ctxKeyType int

var ctxKey ctxKeyType = 0

func NewContext(ctx context.Context, app *App) context.Context {
	return context.WithValue(ctx, ctxKey, app)
}

func FromContext(ctx context.Context) *App {
	app, ok := ctx.Value(ctxKey).(*App)
	if !ok {
		panic(kerr.New("EJRTLPWCKH", "No app in ctx").Error())
	}
	return app
}

func FromContextOrNil(ctx context.Context) *App {
	e, ok := ctx.Value(ctxKey).(*App)
	if ok {
		return e
	}
	return nil
}
