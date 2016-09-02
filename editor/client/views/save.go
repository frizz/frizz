package views

import (
	"context"

	"github.com/davelondon/kerr"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/stores"
	"kego.io/editor/shared"
	"kego.io/ke"
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
					Request: shared.Request{
						Path: v.App.Env.Path(),
						Hash: v.App.Env.Hash(),
					},
				}
				for n, f := range v.App.Files.All() {
					if f.Changed() {
						b, err := ke.MarshalIndentContext(v.Ctx, n.Value, "", "\t")
						if err != nil {
							v.App.Fail <- kerr.Wrap("MKHONBVODK", err)
						}
						request.Files = append(request.Files, shared.SaveRequestFile{
							File:  f.Filename,
							Bytes: b,
							Hash:  f.Hash,
						})
					}
				}
				response := &shared.SaveResponse{}
				done := v.App.Conn.Go(shared.Save, request, response, v.App.Fail)

				go func() {
					if err := <-done; err != nil {
						v.App.Fail <- kerr.Wrap("PVLCRRYIUD", err)
						return
					}
					v.App.Dispatcher.Dispatch(&actions.SaveFileSuccess{Response: response})
				}()

			}).PreventDefault(),
			prop.Href("#"),
			vecty.Text("Save"),
		),
	)
}
