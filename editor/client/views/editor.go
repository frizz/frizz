package views

import (
	"context"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"github.com/gopherjs/gopherjs/js"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system"
	"kego.io/system/node"
)

type EditorView struct {
	*View

	model    *models.EditorModel
	node     *models.NodeModel
	focus    *js.Object
	controls vecty.List
	icons    vecty.List
	dropdown vecty.List
}

func NewEditorView(ctx context.Context, node *node.Node) *EditorView {
	v := &EditorView{}
	v.View = New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.node = v.App.Nodes.Get(node)
	v.Watch(v.model.Node,
		stores.NodeValueChanged,
	)
	return v

}

func (v *EditorView) Dropdown(markup ...vecty.Markup) *EditorView {
	v.dropdown = markup
	return v
}

func (v *EditorView) Icons(markup ...vecty.Markup) *EditorView {
	v.icons = markup
	return v
}

func (v *EditorView) Controls(markup ...vecty.Markup) *EditorView {
	v.controls = markup
	return v
}

func (v *EditorView) FocusElement(o *js.Object) *EditorView {
	v.Watch(v.model.Node,
		stores.NodeFocus,
	)
	return v
}

func (v *EditorView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
	if v.focus != nil && notif.Type == stores.NodeFocus {
		v.focus.Call("focus")
	}
}

func (v *EditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*EditorView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *EditorView) Render() vecty.Component {

	if !v.model.Node.Missing && !v.model.Node.Null {
		v.dropdown = append(v.dropdown, elem.ListItem(
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
		))
	}

	var dropdown vecty.Markup
	if v.dropdown != nil {
		dropdown = elem.Span(
			prop.Class("dropdown"),
			elem.Anchor(
				prop.Href("#"),
				prop.Class("dropdown-toggle"),
				vecty.Data("toggle", "dropdown"),
				vecty.Property("aria-haspopup", "true"),
				vecty.Property("aria-expanded", "true"),
				elem.Italic(
					prop.Class("editor-icon editor-icon-before glyphicon glyphicon-collapse-down"),
				),
			),
			elem.UnorderedList(
				prop.Class("dropdown-menu"),
				v.dropdown,
			),
		)
	}

	var handle vecty.List
	if v.model.Node.Index != -1 {
		handle = append(handle, elem.Italic(
			prop.Class("handle"),
			elem.Span(
				prop.Class("glyphicon glyphicon-option-vertical"),
			),
		))
	}

	label := elem.Label(
		prop.Class("control-label"),
		vecty.Text(
			v.model.Node.Label(v.Ctx),
		),
	)

	group := elem.Div(
		vecty.ClassMap{
			"form-group": true,
			"has-error":  v.node.Invalid,
		},
		handle,
		dropdown,
		label,
		v.icons,
		v.controls,
	)
	v.helpBlock().Apply(group)
	v.errorBlock().Apply(group)

	return group
}

func (v *EditorView) helpBlock() vecty.Markup {
	if v.model.Node.Rule == nil || v.model.Node.Rule.Interface == nil {
		return vecty.List{}
	}
	description := v.model.Node.Rule.Interface.(system.ObjectInterface).GetObject(v.Ctx).Description
	if description == "" {
		return vecty.List{}
	}
	return elem.Paragraph(
		prop.Class("help-block"),
		vecty.Text(description),
	)
}

func (v *EditorView) errorBlock() vecty.Markup {
	if !v.node.Invalid {
		return vecty.List{}
	}

	errors := vecty.List{}
	for _, e := range v.node.Errors {
		errors = append(errors, elem.ListItem(vecty.Text(e.Description)))
	}
	return elem.Paragraph(
		prop.Class("help-block text-danger"),
		elem.UnorderedList(errors),
	)
}
