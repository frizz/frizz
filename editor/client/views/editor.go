package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"golang.org/x/net/context"
	"kego.io/editor/client/clientctx"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system/node"
)

func RegisterDefaultEditor(ctx context.Context) {
	editors := clientctx.FromContext(ctx)
	editors.Set("", new(DefaultEditor))
}

type DefaultEditor struct{}

func (s *DefaultEditor) Format() editable.Format {
	return editable.Branch
}

func (s *DefaultEditor) EditorView(ctx context.Context, node *node.Node) vecty.Component {
	return NewEditorView(ctx, node)
}

type EditorView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model *models.EditorModel
}

func NewEditorView(ctx context.Context, node *node.Node) *EditorView {
	v := &EditorView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.model = v.app.Editors.Get(node)
	v.Mount()
	return v
}

func (v *EditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*EditorView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *EditorView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *EditorView) Mount() {
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

func (v *EditorView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *EditorView) Unmount() {
	if v.notifs != nil {
		v.app.Editors.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *EditorView) render() vecty.Component {
	if v.model == nil {
		return vecty.Text("default editor (nil)")
	}

	children := vecty.List{}
	for label, n := range v.model.Node.Map {
		e := models.GetEditable(v.ctx, n)
		if e.Format() == editable.Block || e.Format() == editable.Inline {
			children = append(children, elem.Div(vecty.Text(label)))
			children = append(children, elem.Div(e.EditorView(v.ctx, n)))
		}
	}
	if len(children) > 0 {
		return elem.Div(children)
	}

	return vecty.Text("default editor (" + v.model.Node.Type.Id.String() + ")")
}
