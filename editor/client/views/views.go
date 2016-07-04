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
	Notifs []chan flux.NotifPayload
}

func New(ctx context.Context, self vecty.Renderer) *View {
	v := &View{
		Composite: vecty.Composite{
			Self: self,
		},
		Ctx: ctx,
		App: stores.FromContext(ctx),
	}
	return v
}

func (v *View) Unmount() {
	for _, n := range v.Notifs {
		v.App.Delete(n)
	}
	v.Notifs = nil
	v.Body.Unmount()
}

// Apply implements the vecty.Markup interface.
func (v *View) Apply(element *vecty.Element) {
	element.AddChild(v.Self)
}

func (v *View) Watch(reaction func(notif flux.NotifPayload), object interface{}, notifs ...flux.Notif) {
	n := v.App.Watch(object, notifs...)
	go func() {
		for notif := range n {
			reaction(notif)
		}
	}()
	v.Notifs = append(v.Notifs, n)
}
