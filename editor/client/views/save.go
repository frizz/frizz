package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/stores"
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
			prop.Href("#"),
			vecty.Text("Save"),
		),
	)
}
