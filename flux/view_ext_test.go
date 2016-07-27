package flux_test

import (
	"testing"

	"github.com/davelondon/ktest/require"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"kego.io/flux"
	"kego.io/flux/mock_flux"
)

func TestNewView(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	ctx := context.Background()
	app := mock_flux.NewMockAppInterface(m)
	v := NewTestView(ctx, app, false)
	require.NotNil(t, v.Composite)
	v.Body = vecty.Nil()
	require.IsType(t, elem.Div(), v.RenderFunc())
	v.Unmount()
	require.IsType(t, vecty.Nil(), v.RenderFunc())
}

func TestView_Unmount(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	ctx := context.Background()
	app := mock_flux.NewMockAppInterface(m)
	app.EXPECT().Watch("a", notif1{})
	v := NewTestView(ctx, app, true)
	v.Body = vecty.Nil()
	app.EXPECT().Delete(nil)
	v.Unmount()
}

func TestView_Receive(t *testing.T) {
	m := gomock.NewController(t)
	defer m.Finish()
	ctx := context.Background()
	app := mock_flux.NewMockAppInterface(m)
	c := make(chan flux.NotifPayload)
	app.EXPECT().Watch("a", notif1{}).Return(c)
	v := NewTestView(ctx, app, true)
	v.Body = vecty.Nil()
	np := flux.NotifPayload{Done: make(chan struct{}, 1)}
	v.Notifs[0] <- np
	<-np.Done
	require.True(t, v.received)
}

type TestView struct {
	*flux.View
	received bool
}

func NewTestView(ctx context.Context, app flux.AppInterface, watch bool) *TestView {
	v := &TestView{}
	v.View = flux.NewView(ctx, v, app)
	if watch {
		v.Watch("a", notif1{})
	}
	return v
}

func (v *TestView) Receive(notif flux.NotifPayload) {
	v.received = true
	close(notif.Done)
}

func (v *TestView) Reconcile(old vecty.Component) {
	if old, ok := old.(*TestView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *TestView) Render() vecty.Component {
	return elem.Div()
}
