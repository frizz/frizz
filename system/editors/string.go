package editors

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system/node"
)

type StringEditor struct{}

func (s *StringEditor) Format() editable.Format {
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
	v.notifs = v.app.Nodes.Watch(v.model.Node,
		stores.NodeInitialised,
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
}

func (v *StringEditorView) Unmount() {
	if v.notifs != nil {
		v.app.Nodes.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *StringEditorView) render() vecty.Component {
	id := randomId()

	return elem.Div(
		prop.Class("form-group"),
		elem.Label(
			prop.For(id),
			vecty.Text(v.model.Node.Label()),
		),
		elem.Input(
			prop.Value(v.model.Node.ValueString),
			prop.Class("form-control"),
			prop.ID(id),
			prop.Placeholder(v.model.Node.Label()),
		),
		helpBlock(v.ctx, v.model.Node),
	)

}
