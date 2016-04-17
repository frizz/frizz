package stores

import (
	"golang.org/x/net/context"
	"kego.io/flux"
	"kego.io/kerr"
)

type App struct {
	Dispatcher *flux.Dispatcher
	Todos      *TodoStore
}

type ctxKeyType int

var ctxKey ctxKeyType = 0

func NewApp(ctx context.Context) (context.Context, *App) {
	a := &App{}
	return NewContext(ctx, a), a
}

func NewContext(ctx context.Context, app *App) context.Context {
	return context.WithValue(ctx, ctxKey, app)
}

func FromContext(ctx context.Context) *App {
	app, ok := ctx.Value(ctxKey).(*App)
	if !ok {
		panic(kerr.New("WTHRXAIJRP", "No app in ctx").Error())
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
