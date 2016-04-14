package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type EditorView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App
	c   chan struct{}

	node  *node.Node
	model *models.EditorModel
}

func NewEditorView(ctx context.Context, node *node.Node) *EditorView {
	v := &EditorView{
		ctx:  ctx,
		app:  stores.FromContext(ctx),
		node: node,
	}
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
	v.c = v.app.Editors.Watch(stores.EditorChanged, v.c)
	go func() {
		for range v.c {
			v.model = v.app.Editors.Get(v.node)
			v.ReconcileBody()
		}
	}()
}

func (v *EditorView) Unmount() {
	if v.c != nil {
		v.app.Editors.Delete(v.c)
		v.c = nil
	}
	v.Body.Unmount()
}

func (v *EditorView) render() vecty.Component {
	if v.node == nil {
		return elem.Div()
	}
	return elem.Div(
		vecty.Text(v.node.Key + " editor"),
	)
}
