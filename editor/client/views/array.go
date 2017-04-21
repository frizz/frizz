package views

import (
	"context"

	"frizz.io/editor/client/actions"
	"frizz.io/editor/client/models"
	"frizz.io/editor/client/stores"
	"frizz.io/system/node"
	"github.com/dave/kerr"
	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/event"
	"github.com/dave/vecty/prop"
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
	v.Watch(nil,
		stores.InfoStateChange,
	)
	v.Watch(node,
		stores.NodeErrorsChanged,
	)
	return v
}

func (v *ArrayView) Render() *vecty.HTML {
	if v.model == nil {
		return elem.Div(vecty.Text("Array (nil)"))
	}

	var info vecty.List
	if v.App.Misc.Info() {
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
		info = append(info,
			elem.Paragraph(
				prop.Class("lead"),
				vecty.Text("type: array of "+dt),
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

func (v *ArrayView) errorBlock() *vecty.HTML {
	if !v.node.Invalid {
		return elem.Div()
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
