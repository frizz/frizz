package views

import (
	"context"

	"frizz.io/editor/client/actions"
	"frizz.io/editor/client/stores"
	"frizz.io/editor/shared"
	"frizz.io/frizz"
	"github.com/dave/kerr"
	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/event"
	"github.com/dave/vecty/prop"
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

func (v *SaveView) Render() *vecty.HTML {
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
						b, err := frizz.MarshalIndent(v.Ctx, n.Value, "", "\t")
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
