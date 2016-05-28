package editors

import (
	"fmt"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system/node"
)

type NumberEditor struct{}

func (s *NumberEditor) GetEditorView(ctx context.Context, node *node.Node) vecty.Component {
	return NewNumberEditorView(ctx, node)
}

type NumberEditorView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model *models.EditorModel
}

func NewNumberEditorView(ctx context.Context, node *node.Node) *NumberEditorView {
	v := &NumberEditorView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.model = v.app.Editors.Get(node)
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
	v.notifs = v.app.Editors.Watch(v.model,
		stores.EditorLoaded,
		stores.EditorInitialStateLoaded,
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
}

func (v *NumberEditorView) Unmount() {
	if v.notifs != nil {
		v.app.Editors.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *NumberEditorView) render() vecty.Component {
	if v.model == nil {
		return elem.Div()
	}
	return elem.Div(
		elem.Input(
			prop.Type(prop.TypeNumber),
			prop.Value(fmt.Sprintf("%v", v.model.Node.ValueNumber)),
		),
	)
}
