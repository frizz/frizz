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
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model  *models.EditorModel
	input  *vecty.Element
	format editable.Format
}

func NewBoolEditorView(ctx context.Context, node *node.Node, format editable.Format) *BoolEditorView {
	v := &BoolEditorView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.model = v.app.Editors.Get(node)
	v.format = format
	v.Mount()
	return v
}

func (v *BoolEditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BoolEditorView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *BoolEditorView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *BoolEditorView) Mount() {
	v.notifs = v.app.Watch(v.model,
		stores.EditorFocus,
		stores.EditorValueChanged,
	)
	go func() {
		for notif := range v.notifs {
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

func (v *BoolEditorView) Unmount() {
	if v.notifs != nil {
		v.app.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *BoolEditorView) render() vecty.Component {

	v.input = elem.Input(
		prop.Type(prop.TypeCheckbox),
		prop.Checked(v.model.Node.ValueBool),
		event.Change(func(e *vecty.Event) {
			change(v.app, v.model, func() interface{} {
				return e.Target.Get("checked").Bool()
			})
		}),
	)

	group := elem.Div(
		vecty.ClassMap{
			"form-group": true,
			"has-error":  v.model.Invalid,
		},
		elem.Div(
			prop.Class("checkbox"),
			elem.Label(
				v.input,
				vecty.Text(v.model.Node.Label(v.ctx)),
			),
		),
	)

	if v.format == editable.Inline {
		return group
	}

	helpBlock(v.ctx, v.model.Node).Apply(group)
	return group
}
