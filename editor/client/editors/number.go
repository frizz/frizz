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
	"kego.io/editor/client/views"
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
	*views.View

	model  *models.EditorModel
	node   *models.NodeModel
	input  *vecty.Element
	format editable.Format
}

func NewNumberEditorView(ctx context.Context, node *node.Node, format editable.Format) *NumberEditorView {
	v := &NumberEditorView{}
	v.View = views.New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.node = v.App.Nodes.Get(node)
	v.format = format
	v.Mount()
	return v
}

func (v *NumberEditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*NumberEditorView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *NumberEditorView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *NumberEditorView) Mount() {
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
	if v.Notifs != nil {
		v.App.Delete(v.Notifs)
		v.Notifs = nil
	}
	v.Body.Unmount()
}

func (v *NumberEditorView) Render() vecty.Component {
	id := randomId()

	v.input = elem.Input(
		prop.Type(prop.TypeNumber),
		prop.Value(fmt.Sprintf("%v", v.model.Node.ValueNumber)),
		prop.Class("form-control"),
		prop.ID(id),
		event.KeyUp(func(e *vecty.Event) {
			change(v.App, v.model, func() interface{} {
				return e.Target.Get("value").String()
			})
		}),
	)

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
	return group
}
