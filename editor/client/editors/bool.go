package editors

import (
	"context"

	"frizz.io/editor/client/actions"
	"frizz.io/editor/client/editable"
	"frizz.io/editor/client/models"
	"frizz.io/editor/client/stores"
	"frizz.io/editor/client/views"
	"frizz.io/flux"
	"frizz.io/system"
	"frizz.io/system/node"
	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/event"
	"github.com/dave/vecty/prop"
)

var _ editable.Editable = (*BoolEditor)(nil)

type BoolEditor struct{}

func (s *BoolEditor) Format(rule *system.RuleWrapper) editable.Format {
	return editable.Inline
}

func (s *BoolEditor) EditorView(ctx context.Context, node *node.Node, format editable.Format) vecty.Component {
	return NewBoolEditorView(ctx, node, format)
}

type BoolEditorView struct {
	*views.View

	model  *models.EditorModel
	node   *models.NodeModel
	input  *vecty.HTML
	format editable.Format
}

func NewBoolEditorView(ctx context.Context, node *node.Node, format editable.Format) *BoolEditorView {
	v := &BoolEditorView{}
	v.View = views.New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.node = v.App.Nodes.Get(node)
	v.format = format
	v.Watch(v.model.Node,
		stores.NodeFocus,
		stores.NodeValueChanged,
		stores.NodeErrorsChanged)
	return v
}

func (v *BoolEditorView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	vecty.Rerender(v)
	if notif.Type == stores.NodeFocus {
		v.Focus()
	}
}

func (v *BoolEditorView) Focus() {
	v.input.Node.Call("focus")
}

func (v *BoolEditorView) Render() *vecty.HTML {

	v.input = elem.Input(
		prop.Type(prop.TypeCheckbox),
		prop.Checked(v.model.Node.ValueBool),
		event.Change(func(e *vecty.Event) {
			v.App.Dispatch(&actions.Modify{
				Undoer:    &actions.Undoer{},
				Editor:    v.model,
				Before:    v.model.Node.NativeValue(),
				After:     e.Target.Get("checked").Bool(),
				Immediate: true,
			})
		}),
	)

	return elem.Div(
		prop.Class("checkbox"),
		elem.Label(
			v.input,
			vecty.Text(v.model.Node.Label(v.Ctx)),
		),
	)

}
