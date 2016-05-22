package ctests

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/connection"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/process/tests"
)

type ClientContextBuilder struct {
	Base *tests.ContextBuilder
}

func New() *ClientContextBuilder {
	return &ClientContextBuilder{tests.New()}
}

func (cb *ClientContextBuilder) Ctx() context.Context {
	return cb.Base.Ctx()
}

func (cb *ClientContextBuilder) SetApp() *ClientContextBuilder {
	app := &stores.App{
		Conn: &connection.Mock{},
		Fail: make(chan error, 1),
	}
	ctx := stores.NewContext(cb.Ctx(), app)
	cb.Base.SetCtx(ctx)

	app.Init(cb.Ctx())
	app.Dispatcher = &flux.DispatcherMock{}
	return cb
}

func (cb *ClientContextBuilder) GetApp() *stores.App {
	return stores.FromContext(cb.Ctx())
}

func (cb *ClientContextBuilder) GetDispatcher() *flux.DispatcherMock {
	app := stores.FromContext(cb.Ctx())
	return app.Dispatcher.(*flux.DispatcherMock)
}

func (cb *ClientContextBuilder) SetConnReply(reply interface{}) *ClientContextBuilder {
	cb.GetConn().Reply = reply
	return cb
}

func (cb *ClientContextBuilder) GetConn() *connection.Mock {
	app := stores.FromContext(cb.Ctx())
	return app.Conn.(*connection.Mock)
}

func (cb *ClientContextBuilder) ResetAppMocks() *ClientContextBuilder {
	app := stores.FromContext(cb.Ctx())
	app.Dispatcher = &flux.DispatcherMock{}
	app.Conn = &connection.Mock{}
	return cb
}
