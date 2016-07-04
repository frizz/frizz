package views

import (
	"github.com/davelondon/vecty"
	"golang.org/x/net/context"
	"kego.io/editor/client/stores"
	"kego.io/flux"
)

type View struct {
	vecty.Composite
	Ctx    context.Context
	App    *stores.App
	Notifs chan flux.NotifPayload
}

func New(ctx context.Context, renderer vecty.Renderer) *View {
	v := &View{
		Composite: vecty.Composite{
			Renderer: renderer,
		},
		Ctx: ctx,
		App: stores.FromContext(ctx),
	}
	return v
}

func (v *View) Unmount() {
	if v.Notifs != nil {
		v.App.Delete(v.Notifs)
		v.Notifs = nil
	}
	v.Body.Unmount()
}
