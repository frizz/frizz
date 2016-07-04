package flux

import (
	"github.com/davelondon/vecty"
	"golang.org/x/net/context"
)

type ViewInterface interface {
	vecty.Component
	Render() vecty.Component
	Receive(notif NotifPayload)
}

type View struct {
	vecty.Composite
	Ctx    context.Context
	App    AppInterface
	Notifs []chan NotifPayload
	Self   ViewInterface
}

func NewView(ctx context.Context, self ViewInterface, app AppInterface) *View {
	v := &View{
		Composite: vecty.Composite{
			RenderFunc: self.Render,
		},
		Ctx:  ctx,
		App:  app,
		Self: self,
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

func (v *View) Watch(object interface{}, notifs ...Notif) {
	n := v.App.Watch(object, notifs...)
	go func() {
		for notif := range n {
			v.Self.Receive(notif)
		}
	}()
	v.Notifs = append(v.Notifs, n)
}

func (v *View) Receive(notif NotifPayload) {
	// Default receive function for when view doesn't override it
	defer close(notif.Done)
	v.ReconcileBody()
}
