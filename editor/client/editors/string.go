package editors

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system"
	"kego.io/system/node"
)

type StringEditor struct{}

func (s *StringEditor) Format(rule *system.RuleWrapper) editable.Format {
	if rule != nil {
		if r, ok := rule.Interface.(*system.StringRule); ok && r.Long {
			return editable.Block
		}
	}
	return editable.Inline
}

func (s *StringEditor) EditorView(ctx context.Context, node *node.Node) vecty.Component {
	return NewStringEditorView(ctx, node)
}

type StringEditorView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model *models.EditorModel
	input *vecty.Element
}

func NewStringEditorView(ctx context.Context, node *node.Node) *StringEditorView {
	v := &StringEditorView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.model = v.app.Editors.Get(node)
	v.Mount()
	return v
}

func (v *StringEditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*StringEditorView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *StringEditorView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *StringEditorView) Mount() {
	v.notifs = v.app.Watch(v.model,
		stores.EditorFocus,
	)
	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *StringEditorView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
	if notif.Type == stores.EditorFocus {
		v.Focus()
	}
}

func (v *StringEditorView) Focus() {
	v.input.Node().Call("focus")
}

func (v *StringEditorView) Unmount() {
	if v.notifs != nil {
		v.app.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *StringEditorView) render() vecty.Component {
	id := randomId()

	sr, ok := v.model.Node.Rule.Interface.(*system.StringRule)
	if ok && sr.Long {
		v.input = elem.TextArea(
			prop.Value(v.model.Node.ValueString),
			prop.Class("form-control"),
			prop.ID(id),
			event.KeyUp(func(e *vecty.Event) {
				v.app.Dispatch(&actions.EditorValueChange{
					Editor: v.model,
					Value:  e.Target.Get("value").String(),
				})
			}),
		)
	} else {
		v.input = elem.Input(
			prop.Value(v.model.Node.ValueString),
			prop.Class("form-control"),
			prop.ID(id),
			event.KeyUp(func(e *vecty.Event) {
				v.app.Dispatch(&actions.EditorValueChange{
					Editor: v.model,
					Value:  e.Target.Get("value").String(),
				})
			}),
		)
	}

	return elem.Div(
		prop.Class("form-group"),
		elem.Label(
			prop.For(id),
			vecty.Text(v.model.Node.Label(v.ctx)),
		),
		v.input,
		helpBlock(v.ctx, v.model.Node),
	)

}
