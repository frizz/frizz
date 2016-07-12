package editors

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/editor/client/views"
	"kego.io/flux"
	"kego.io/system"
	"kego.io/system/node"
)

var _ editable.Editable = (*StringEditor)(nil)

type StringEditor struct{}

func (s *StringEditor) Format(rule *system.RuleWrapper) editable.Format {
	if rule != nil {
		if r, ok := rule.Interface.(*system.StringRule); ok && r.Long {
			return editable.Block
		}
	}
	return editable.Inline
}

func (s *StringEditor) EditorView(ctx context.Context, node *node.Node, format editable.Format) vecty.Component {
	return NewStringEditorView(ctx, node, format)
}

type StringEditorView struct {
	*views.View

	model  *models.EditorModel
	node   *models.NodeModel
	input  *vecty.Element
	format editable.Format
}

func NewStringEditorView(ctx context.Context, node *node.Node, format editable.Format) *StringEditorView {
	v := &StringEditorView{}
	v.View = views.New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.node = v.App.Nodes.Get(node)
	v.format = format
	v.Watch(v.model.Node,
		stores.NodeFocus,
		stores.NodeValueChanged,
		stores.NodeErrorsChanged,
	)
	return v
}

func (v *StringEditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*StringEditorView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *StringEditorView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
	if notif.Type == stores.NodeFocus {
		v.Focus()
	}
}

func (v *StringEditorView) Focus() {
	v.input.Node().Call("focus")
}

func (v *StringEditorView) Render() vecty.Component {
	id := randomId()

	sr, ok := v.model.Node.Rule.Interface.(*system.StringRule)

	if ok && sr.Long {
		v.input = elem.TextArea()
	} else {
		v.input = elem.Input(
			prop.Type(prop.TypeText),
		)
	}

	vecty.List{
		prop.Value(v.model.Node.ValueString),
		prop.Class("form-control"),
		prop.ID(id),
		event.KeyUp(func(e *vecty.Event) {
			change(v.App, v.model, func() interface{} {
				return e.Target.Get("value").String()
			})
		}),
	}.Apply(v.input)

	group := elem.Div(
		vecty.ClassMap{
			"form-group": true,
			"has-error":  v.node.Invalid,
		},
		elem.Label(
			prop.For(id),
			vecty.Text(v.model.Node.Label(v.Ctx)),
		),
		v.input,
	)

	if v.format == editable.Inline {
		return group
	}

	helpBlock(v.Ctx, v.model.Node).Apply(group)
	errorBlock(v.Ctx, v.node).Apply(group)
	return group
}
