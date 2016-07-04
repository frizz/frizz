package views

import (
	"fmt"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"github.com/davelondon/vecty/style"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/system/node"
)

type ObjectView struct {
	*View

	model *models.EditorModel
}

func NewObjectView(ctx context.Context, node *node.Node) *ObjectView {
	v := &ObjectView{}
	v.View = New(ctx, v)
	v.model = v.App.Editors.Get(node)
	return v
}

func (v *ObjectView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ObjectView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *ObjectView) Render() vecty.Component {
	if v.model == nil {
		return elem.Div(vecty.Text("Object (nil)"))
	}

	tabs := vecty.List{}
	panes := vecty.List{}
	for i, o := range v.model.Node.Type.FieldOrigins() {

		tabs = append(tabs, elem.ListItem(
			vecty.ClassMap{
				"active": i == 0,
			},
			elem.Anchor(
				prop.Href(fmt.Sprintf("#pane%d", i)),
				vecty.Data("toggle", "tab"),
				vecty.Text(o.Name),
			),
		))

		panes = append(panes, elem.Div(
			vecty.ClassMap{
				"tab-pane": true,
				"active":   i == 0,
			},
			prop.ID(fmt.Sprintf("pane%d", i)),
			NewEditorListView(v.Ctx, v.model, o),
			NewObjectTableView(v.Ctx, v.model, o),
		))
	}

	return elem.Div(
		elem.UnorderedList(
			prop.ID("tabs"),
			prop.Class("nav nav-tabs"),
			style.Margin(style.Size("0 0 15px 0")),
			vecty.Data("tabs", "tabs"),
			tabs,
		),
		elem.Div(
			prop.Class("tab-content"),
			panes,
		),
	)

}
