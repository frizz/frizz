package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type ArrayView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	model *models.EditorModel
}

func NewArrayView(ctx context.Context, node *node.Node) *ArrayView {
	v := &ArrayView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.RenderFunc = v.render
	v.model = v.app.Editors.Get(node)
	return v
}

func (v *ArrayView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ArrayView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *ArrayView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *ArrayView) render() vecty.Component {
	if v.model == nil {
		return elem.Div(vecty.Text("Array (nil)"))
	}

	return elem.Div(
		elem.Div(
			elem.Button(
				prop.Type(prop.TypeButton),
				prop.Class("btn btn-primary"),
				vecty.Text("Add"),
				event.Click(func(ev *vecty.Event) {
					addCollectionItem(v.app, v.model.Node)
				}).PreventDefault(),
			),
		),
		NewEditorListView(v.ctx, v.model, nil),
		NewArrayTableView(v.ctx, v.model),
	)

}
