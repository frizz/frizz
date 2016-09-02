package views

import (
	"context"

	"github.com/davelondon/kerr"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/system/node"
)

type ArrayView struct {
	*View

	model  *models.EditorModel
	branch *models.BranchModel
	node   *models.NodeModel
}

func NewArrayView(ctx context.Context, node *node.Node) *ArrayView {
	v := &ArrayView{}
	v.View = New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.branch = v.App.Branches.Get(node)
	v.node = v.App.Nodes.Get(node)
	return v
}

func (v *ArrayView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ArrayView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *ArrayView) Render() vecty.Component {
	if v.model == nil {
		return elem.Div(vecty.Text("Array (nil)"))
	}

	ir, err := v.node.Node.Rule.ItemsRule()
	if err != nil {
		v.App.Fail <- kerr.Wrap("WTLDRIWHQL", err)
		return nil
	}
	dt, err := ir.DisplayType()
	if err != nil {
		v.App.Fail <- kerr.Wrap("XJVBKRSABX", err)
		return nil
	}

	return elem.Div(
		NewPanelNavView(v.Ctx, v.branch).Contents(
			elem.UnorderedList(
				prop.Class("nav navbar-nav navbar-right"),
				elem.ListItem(
					elem.Anchor(
						vecty.Text("Add"),
						prop.Href("#"),
						event.Click(func(ev *vecty.Event) {
							addCollectionItem(v.App, v.model.Node)
						}).PreventDefault(),
					),
				),
				elem.ListItem(
					prop.Class("dropdown"),
					elem.Anchor(
						prop.Href("#"),
						prop.Class("dropdown-toggle"),
						vecty.Data("toggle", "dropdown"),
						vecty.Property("role", "button"),
						vecty.Property("aria-haspopup", "true"),
						vecty.Property("aria-expanded", "false"),
						vecty.Text("Options"),
						elem.Span(
							prop.Class("caret"),
						),
					),
					elem.UnorderedList(
						prop.Class("dropdown-menu"),
						elem.ListItem(
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("Type: system:array"),
							),
						),
						elem.ListItem(
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("Items: "+dt),
							),
						),
						elem.ListItem(
							prop.Class("divider"),
							vecty.Property("role", "separator"),
						),
						elem.ListItem(
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("Delete"),
								event.Click(func(e *vecty.Event) {
									v.App.Dispatch(&actions.Delete{
										Undoer: &actions.Undoer{},
										Node:   v.model.Node,
										Parent: v.model.Node.Parent,
									})
								}).PreventDefault(),
							),
						),
					),
				),
			),
		),
		NewEditorListView(v.Ctx, v.model, nil),
	)

}
