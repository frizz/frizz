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

type BoolEditor struct{}

func (s *BoolEditor) Format() editable.Format {
	return editable.Inline
}

func (s *BoolEditor) EditorView(ctx context.Context, node *node.Node) vecty.Component {
	return NewBoolEditorView(ctx, node)
}

type BoolEditorView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model *models.EditorModel
}

func NewBoolEditorView(ctx context.Context, node *node.Node) *BoolEditorView {
	v := &BoolEditorView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.model = v.app.Editors.Get(node)
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
	v.notifs = v.app.Nodes.Watch(v.model.Node,
		stores.NodeInitialised,
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
}

func (v *BoolEditorView) Unmount() {
	if v.notifs != nil {
		v.app.Nodes.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *BoolEditorView) render() vecty.Component {
	return elem.Div(
		prop.Class("checkbox"),
		elem.Label(
			elem.Input(
				prop.Type(prop.TypeCheckbox),
				prop.Checked(v.model.Node.ValueBool),
			),
			vecty.Text(v.model.Node.Label()),
		),
	)
}
