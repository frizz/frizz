package editors

import (
	"context"

	"github.com/davelondon/kerr"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/editor/client/views"
	"kego.io/system"
	"kego.io/system/node"
)

var _ editable.Editable = (*ObjectEditor)(nil)

type ObjectEditor struct{}

func (s *ObjectEditor) Format(rule *system.RuleWrapper) editable.Format {
	return editable.Block
}

func (s *ObjectEditor) EditorView(ctx context.Context, node *node.Node, format editable.Format) vecty.Component {
	return NewObjectEditorView(ctx, node)
}

type ObjectEditorView struct {
	*views.View

	model  *models.EditorModel
	node   *models.NodeModel
	branch *models.BranchModel
	object *system.Object
}

func NewObjectEditorView(ctx context.Context, node *node.Node) *ObjectEditorView {
	v := &ObjectEditorView{}
	v.View = views.New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.node = v.App.Nodes.Get(node)
	v.branch = v.App.Branches.Get(node)
	v.object = node.Value.(system.ObjectInterface).GetObject(v.Ctx)
	v.Watch(v.model.Node,
		stores.NodeValueChanged,
		stores.NodeErrorsChanged,
	)
	return v
}

func (v *ObjectEditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ObjectEditorView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *ObjectEditorView) Render() vecty.Component {

	dt, err := v.node.Node.DisplayType(v.Ctx)
	if err != nil {
		v.App.Fail <- kerr.Wrap("KFKGCGFULR", err)
		return nil
	}

	sections := vecty.List{}
	sections = append(sections,
		views.NewPanelNavView(v.Ctx, v.branch).Contents(
			elem.UnorderedList(
				prop.Class("nav navbar-nav navbar-right"),
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
								vecty.Text("Type: "+dt),
							),
						),
						elem.ListItem(
							prop.Class("divider"),
							vecty.Property("role", "separator"),
						),
						elem.ListItem(
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("Description..."),
							),
						),
						elem.ListItem(
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("Tags..."),
							),
						),
						elem.ListItem(
							elem.Anchor(
								prop.Href("#"),
								vecty.Text("Rules..."),
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
	)

	if v.object.Description != "" {
		sections = append(sections,
			elem.Paragraph(
				prop.Class("lead"),
				vecty.Text(v.object.Description),
			),
		)
	}

	return elem.Div(
		sections,
	)

}
