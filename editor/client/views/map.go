package views

import (
	"context"

	"github.com/dave/kerr"
	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/event"
	"github.com/dave/vecty/prop"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type MapView struct {
	*View

	model  *models.EditorModel
	branch *models.BranchModel
	node   *models.NodeModel
}

func NewMapView(ctx context.Context, node *node.Node) *MapView {
	v := &MapView{}
	v.View = New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.branch = v.App.Branches.Get(node)
	v.node = v.App.Nodes.Get(node)
	v.Watch(nil,
		stores.InfoStateChange,
	)
	v.Watch(node,
		stores.NodeErrorsChanged,
	)
	return v
}

func (v *MapView) Render() *vecty.HTML {
	if v.model == nil {
		return elem.Div(vecty.Text("Map (nil)"))
	}

	var info vecty.List
	if v.App.Misc.Info() {
		ir, err := v.node.Node.Rule.ItemsRule()
		if err != nil {
			v.App.Fail <- kerr.Wrap("KYTRRFBKGP", err)
			return nil
		}
		dt, err := ir.DisplayType()
		if err != nil {
			v.App.Fail <- kerr.Wrap("RRJAVDKLSI", err)
			return nil
		}
		info = append(info,
			elem.Paragraph(
				prop.Class("lead"),
				vecty.Text("type: map of "+dt),
			),
		)
	}

	return elem.Div(
		NewPanelNavView(v.Ctx, v.branch).Contents(
			func() vecty.MarkupOrComponentOrHTML {
				return elem.UnorderedList(
					prop.Class("nav navbar-nav navbar-right"),
					elem.ListItem(
						elem.Anchor(
							vecty.Text("Add"),
							prop.Href("#"),
							event.Click(func(ev *vecty.Event) {
								addCollectionItem(v.Ctx, v.App, v.model.Node)
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
				)
			},
		),
		info,
		NewEditorListView(v.Ctx, v.model, nil, nil),
		v.errorBlock(),
	)

}

func (v *MapView) errorBlock() *vecty.HTML {
	if !v.node.Invalid {
		return nil
	}

	errors := vecty.List{}
	for _, e := range v.node.Errors {
		errors = append(errors, elem.ListItem(vecty.Text(e.Description)))
	}
	return elem.Div(
		prop.Class("has-error"),
		elem.Paragraph(
			prop.Class("help-block text-danger"),
			elem.UnorderedList(errors),
		),
	)
}
