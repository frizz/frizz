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

	node  *node.Node
	model *models.EditorModel
}

func NewEditorView(ctx context.Context, node *node.Node) *EditorView {

	e := &EditorView{
		ctx:  ctx,
		app:  stores.FromContext(ctx),
		node: node,
	}

	go func() {
		for range e.app.Editors.Changed() {
			e.model = e.app.Editors.Get(e.node)
			e.ReconcileBody()
		}
	}()

	return e
}

// Apply implements the vecty.Markup interface.
func (e *EditorView) Apply(element *vecty.Element) {
	element.AddChild(e)
}

func (e *EditorView) Reconcile(old vecty.Component) {
	if old, ok := old.(*EditorView); ok {
		e.Body = old.Body
		e.node = old.node
		e.model = old.model
	}
	e.RenderFunc = e.render
	e.ReconcileBody()
}

func (e *EditorView) render() vecty.Component {
	label := "no editor selected"
	if e.node != nil {
		label = e.node.Key + " editor"
	}
	return elem.Div(
		vecty.Text(label),
	)
}
