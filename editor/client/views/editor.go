package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"golang.org/x/net/context"
	"kego.io/editor"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type Editor struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	node   *node.Node
	editor editor.EditorInterface
}

func NewEditor(ctx context.Context, node *node.Node) *Editor {

	e := &Editor{
		ctx:  ctx,
		app:  stores.FromContext(ctx),
		node: node,
	}

	go func() {
		for _ = range e.app.Editors.Changed() {
			e.editor = e.app.Editors.Get(e.node)
			e.ReconcileBody()
		}
	}()

	return e
}

// Apply implements the vecty.Markup interface.
func (e *Editor) Apply(element *vecty.Element) {
	element.AddChild(e)
}

func (e *Editor) Reconcile(old vecty.Component) {
	if old, ok := old.(*Editor); ok {
		e.Body = old.Body
		e.node = old.node
		e.editor = old.editor
	}
	e.RenderFunc = e.render
	e.ReconcileBody()
}

func (e *Editor) render() vecty.Component {
	return elem.Div(
		vecty.Text("foo"),
	)
}
