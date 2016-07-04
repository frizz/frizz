package views

import (
	"github.com/davelondon/vecty"
	"golang.org/x/net/context"
	"kego.io/editor/client/stores"
	"kego.io/flux"
)

type Receiver interface {
	vecty.Renderer
	Receive(notif flux.NotifPayload)
}

type View struct {
	vecty.Composite
	Ctx      context.Context
	App      *stores.App
	Notifs   []chan flux.NotifPayload
	Receiver Receiver
}

func New(ctx context.Context, self Receiver) *View {
	v := &View{
		Composite: vecty.Composite{
			Self: self,
		},
		Ctx:      ctx,
		App:      stores.FromContext(ctx),
		Receiver: self,
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

func (v *View) Watch(object interface{}, notifs ...flux.Notif) {
	n := v.App.Watch(object, notifs...)
	go func() {
		for notif := range n {
			v.Receiver.Receive(notif)
		}
	}()
	v.Notifs = append(v.Notifs, n)
}

func (v *View) Receive(notif flux.NotifPayload) {
	// Default receive function for when view doesn't override it
	defer close(notif.Done)
	v.ReconcileBody()
}
