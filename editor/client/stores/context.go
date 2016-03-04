package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/connection"
	"kego.io/editor/client/flux"
	"kego.io/kerr"
)

type App struct {
	Dispatcher *flux.Dispatcher
	Messages   *MessageStore
	Fail       chan error
	Conn       *connection.Conn
}

type ctxKeyType int

var ctxKey ctxKeyType = 0

func NewContext(ctx context.Context, app *App) context.Context {
	return context.WithValue(ctx, ctxKey, app)
}

func FromContext(ctx context.Context) *App {
	app, ok := ctx.Value(ctxKey).(*App)
	if !ok {
		panic(kerr.New("BIUVXISEMA", "No app in ctx").Error())
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
