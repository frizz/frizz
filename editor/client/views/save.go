package views

import (
	"net/rpc"

	"fmt"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/stores"
	"kego.io/editor/shared"
)

type SaveView struct {
	*View
}

func NewSaveView(ctx context.Context) *SaveView {
	v := &SaveView{}
	v.View = New(ctx, v)
	v.Watch(nil, stores.FileChangedStateChange)
	return v
}

func (v *SaveView) Reconcile(old vecty.Component) {
	if old, ok := old.(*SaveView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *SaveView) Render() vecty.Component {
	return elem.ListItem(
		vecty.ClassMap{
			"disabled": !v.App.Files.Changed(),
		},
		elem.Anchor(

			event.Click(func(e *vecty.Event) {
				request := &shared.SaveRequest{
					File:  "foo",
					Bytes: []byte("bar"),
				}
				response := shared.SaveResponse{}
				done := make(chan *rpc.Call, 1)
				call := v.App.Conn.Go(shared.Save, request, &response, done, v.App.Fail)
				//v.App.Dispatcher.Dispatch(&actions.SaveSourceSent{Branch: b}))

				go func() {
					<-call.Done
					fmt.Println("Save", response.Error)
				}()

			}).PreventDefault(),
			prop.Href("#"),
			vecty.Text("Save"),
		),
	)
}
