package flux

import (
	"context"

	"github.com/dave/vecty"
)

type ViewInterface interface {
	vecty.Component
	Receive(notif NotifPayload)
}

type View struct {
	vecty.Core
	Ctx       context.Context
	App       AppInterface
	Notifs    []chan NotifPayload
	Self      ViewInterface
	unmounted bool
	//	index     int
}

//var index int

func NewView(ctx context.Context, self ViewInterface, app AppInterface) *View {
	v := &View{
		Ctx:  ctx,
		App:  app,
		Self: self,
	}
	//	index++
	//	v.index = index
	return v
}

//func (v *View) Desc() string {
//	return fmt.Sprintf("%T %d", v.Self, v.index)
//}

func (v *View) Render() *vecty.HTML {
	//if v.unmounted {
	//	return nil
	//}
	return v.Self.Render()
}

func (v *View) Unmount() {
	v.unmounted = true
	for _, n := range v.Notifs {
		v.App.Delete(n)
	}
	v.Notifs = nil
	v.Core.Unmount()
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
	// notest
	// Default receive function for when view doesn't override it
	defer close(notif.Done)
	vecty.Rerender(v.Self)
}
