package ctests

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/mock_vecty"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"kego.io/editor/client/connection/mock_connection"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/flux/mock_flux"
	"kego.io/tests"
)

type ClientContextBuilder struct {
	Base *tests.ContextBuilder
	mock *gomock.Controller
	t    *testing.T
}

func New(t *testing.T) *ClientContextBuilder {
	cb := &ClientContextBuilder{Base: tests.New()}
	cb.t = t
	cb.mock = gomock.NewController(t)
	return cb
}

func (cb *ClientContextBuilder) Finish() {
	cb.mock.Finish()
}

func (cb *ClientContextBuilder) GetMock() *gomock.Controller {
	return cb.mock
}

func (cb *ClientContextBuilder) Ctx() context.Context {
	return cb.Base.Ctx()
}

func (cb *ClientContextBuilder) SetApp() *ClientContextBuilder {
	app := &stores.App{
		Conn: mock_connection.NewMockInterface(cb.mock),
		Fail: make(chan error, 1),
	}
	ctx := stores.NewContext(cb.Ctx(), app)
	cb.Base.SetCtx(ctx)

	app.Init(cb.Ctx())
	app.Dispatcher = mock_flux.NewMockDispatcherInterface(cb.mock)
	return cb
}

func (cb *ClientContextBuilder) GetApp() *stores.App {
	return stores.FromContext(cb.Ctx())
}

func (cb *ClientContextBuilder) GetDispatcher() *mock_flux.MockDispatcherInterface {
	app := stores.FromContext(cb.Ctx())
	return app.Dispatcher.(*mock_flux.MockDispatcherInterface)
}

func (cb *ClientContextBuilder) AssertAppFail(errorId string) *ClientContextBuilder {
	assert.IsError(cb.t, readErrorOrNil(cb.GetApp().Fail), errorId)
	return cb
}

func (cb *ClientContextBuilder) AssertAppSuccess() *ClientContextBuilder {
	assert.NoError(cb.t, readErrorOrNil(cb.GetApp().Fail))
	return cb
}

func readErrorOrNil(c chan error) error {
	select {
	case err := <-c:
		return err
	default:
		return nil
	}
}

func (cb *ClientContextBuilder) ExpectReconcile(comp *vecty.Composite) *ClientContextBuilder {
	c := mock_vecty.NewMockComponent(cb.GetMock())
	c.EXPECT().Reconcile(nil)
	comp.RenderFunc = func() vecty.Component { return c }
	return cb
}

func (cb *ClientContextBuilder) ExpectNoneDispatched(actions ...flux.ActionInterface) *ClientContextBuilder {
	cb.GetDispatcher().EXPECT().Dispatch(gomock.Any()).Times(0)
	return cb
}

func (cb *ClientContextBuilder) ExpectDispatched(actions ...flux.ActionInterface) *ClientContextBuilder {
	closedChannel := func() chan struct{} {
		c := make(chan struct{}, 1)
		close(c)
		return c
	}
	expects := make([]*gomock.Call, len(actions))
	for i, action := range actions {
		expects[i] = cb.GetDispatcher().EXPECT().Dispatch(action).Return(closedChannel())
	}
	gomock.InOrder(expects...)
	return cb
}

func (cb *ClientContextBuilder) GetConnection() *mock_connection.MockInterface {
	app := stores.FromContext(cb.Ctx())
	return app.Conn.(*mock_connection.MockInterface)
}
