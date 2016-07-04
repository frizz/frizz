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

type ArrayView struct {
	*View

	model *models.EditorModel
}

func NewArrayView(ctx context.Context, node *node.Node) *ArrayView {
	v := &ArrayView{}
	v.View = New(ctx, v)
	v.model = v.App.Editors.Get(node)
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

func (v *ArrayView) Render() vecty.Component {
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
					addCollectionItem(v.App, v.model.Node)
				}).PreventDefault(),
			),
		),
		NewEditorListView(v.Ctx, v.model, nil),
		NewArrayTableView(v.Ctx, v.model),
	)

}
