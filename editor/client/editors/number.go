package editors

import (
	"fmt"

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

var _ editable.Editable = (*NumberEditor)(nil)

type NumberEditor struct{}

func (s *NumberEditor) Format(rule *system.RuleWrapper) editable.Format {
	return editable.Inline
}

func (s *NumberEditor) EditorView(ctx context.Context, node *node.Node, format editable.Format) vecty.Component {
	return NewNumberEditorView(ctx, node, format)
}

type NumberEditorView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model  *models.EditorModel
	input  *vecty.Element
	format editable.Format
}

func NewNumberEditorView(ctx context.Context, node *node.Node, format editable.Format) *NumberEditorView {
	v := &NumberEditorView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.model = v.app.Editors.Get(node)
	v.format = format
	v.Mount()
	return v
}

func (v *NumberEditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*NumberEditorView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *NumberEditorView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *NumberEditorView) Mount() {
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

func (v *NumberEditorView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
	if notif.Type == stores.EditorFocus {
		v.Focus()
	}
}

func (v *NumberEditorView) Focus() {
	v.input.Node().Call("focus")
}

func (v *NumberEditorView) Unmount() {
	if v.notifs != nil {
		v.app.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *NumberEditorView) render() vecty.Component {
	id := randomId()

	v.input = elem.Input(
		prop.Type(prop.TypeNumber),
		prop.Value(fmt.Sprintf("%v", v.model.Node.ValueNumber)),
		prop.Class("form-control"),
		prop.ID(id),
		event.KeyUp(func(e *vecty.Event) {
			change(v.app, v.model, func() interface{} {
				return e.Target.Get("value").String()
			})
		}),
	)

	group := elem.Div(
		vecty.ClassMap{
			"form-group": true,
			"has-error":  v.model.Invalid,
		},
		elem.Label(
			prop.For(id),
			vecty.Text(v.model.Node.Label(v.ctx)),
		),
		v.input,
	)

	if v.format == editable.Inline {
		return group
	}

	helpBlock(v.ctx, v.model.Node).Apply(group)
	return group
}
