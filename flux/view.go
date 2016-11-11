package flux

import (
	"context"

	"github.com/davelondon/vecty"
)

type ViewInterface interface {
	vecty.Component
	Render() vecty.Component
	Receive(notif NotifPayload)
}

type View struct {
	vecty.Composite
	Ctx       context.Context
	App       AppInterface
	Notifs    []chan NotifPayload
	Self      ViewInterface
	unmounted bool
}

func NewView(ctx context.Context, self ViewInterface, app AppInterface) *View {
	v := &View{
		Ctx:  ctx,
		App:  app,
		Self: self,
	}
	v.Composite = vecty.Composite{
		RenderFunc: v.render,
	}
	return v
}

func (v *View) render() vecty.Component {
	// TODO: was this needed?
	//if v.unmounted {
	//	return elem.Nil()
	//}
	return v.Self.Render()
}

func (v *View) Unmount() {
	v.unmounted = true
	for _, n := range v.Notifs {
		v.App.Delete(n)
	}
	v.Notifs = nil
	v.Body.Unmount()
}

// Apply implements the vecty.Markup interface.
func (v *View) Apply(element *vecty.Element) {
	// ke: {"block": {"notest": true}}
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
	// ke: {"block": {"notest": true}}
	// Default receive function for when view doesn't override it
	defer close(notif.Done)
	v.ReconcileBody()
}
