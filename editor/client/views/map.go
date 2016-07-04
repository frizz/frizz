package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/system/node"
)

type MapView struct {
	*View

	model *models.EditorModel
}

func NewMapView(ctx context.Context, node *node.Node) *MapView {
	v := &MapView{}
	v.View = New(ctx, v)
	v.model = v.App.Editors.Get(node)
	return v
}

func (v *MapView) Reconcile(old vecty.Component) {
	if old, ok := old.(*MapView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *MapView) Render() vecty.Component {
	if v.model == nil {
		return elem.Div(vecty.Text("Map (nil)"))
	}

	return elem.Div(
		elem.Div(
			elem.Button(
				prop.Type(prop.TypeButton),
				prop.Class("btn btn-primary"),
				vecty.Text("Add"),
				event.Click(func(ev *vecty.Event) {
					addCollectionItem(v.App, v.model.Node)
				}).PreventDefault(),
			),
		),
		NewEditorListView(v.Ctx, v.model, nil),
		NewMapTableView(v.Ctx, v.model),
	)

}
