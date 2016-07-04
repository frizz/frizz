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
	input  *vecty.Element
	format editable.Format
}

func NewBoolEditorView(ctx context.Context, node *node.Node, format editable.Format) *BoolEditorView {
	v := &BoolEditorView{}
	v.View = views.New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.node = v.App.Nodes.Get(node)
	v.format = format
	v.Mount()
	return v
}

func (v *BoolEditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BoolEditorView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *BoolEditorView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *BoolEditorView) Mount() {
	v.Notifs = v.App.Watch(v.model,
		stores.EditorFocus,
		stores.EditorValueChanged,
		stores.EditorErrorsChanged,
	)
	go func() {
		for notif := range v.Notifs {
			v.reaction(notif)
		}
	}()
}

func (v *BoolEditorView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
	if notif.Type == stores.EditorFocus {
		v.Focus()
	}
}

func (v *BoolEditorView) Focus() {
	v.input.Node().Call("focus")
}

func (v *BoolEditorView) Render() vecty.Component {

	v.input = elem.Input(
		prop.Type(prop.TypeCheckbox),
		prop.Checked(v.model.Node.ValueBool),
		event.Change(func(e *vecty.Event) {
			change(v.App, v.model, func() interface{} {
				return e.Target.Get("checked").Bool()
			})
		}),
	)

	group := elem.Div(
		vecty.ClassMap{
			"form-group": true,
			"has-error":  v.node.Invalid,
		},
		elem.Div(
			prop.Class("checkbox"),
			elem.Label(
				v.input,
				vecty.Text(v.model.Node.Label(v.Ctx)),
			),
		),
	)

	if v.format == editable.Inline {
		return group
	}

	helpBlock(v.Ctx, v.model.Node).Apply(group)
	return group
}
