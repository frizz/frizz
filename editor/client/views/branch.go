package views

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor"
	"kego.io/editor/client/stores"
)

type Branch struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	node *editor.Node
}

func NewBranch(ctx context.Context, node *editor.Node) *Branch {
	if node == nil {
		return nil
	}
	return &Branch{
		ctx:  ctx,
		app:  stores.FromContext(ctx),
		node: node,
	}
}

// Apply implements the vecty.Markup interface.
func (b *Branch) Apply(element *vecty.Element) {
	element.AddChild(b)
}

func (b *Branch) Reconcile(old vecty.Component) {
	if old, ok := old.(*Branch); ok {
		b.Body = old.Body
		b.node = old.node
	}
	b.RenderFunc = b.render
	b.ReconcileBody()
}

func (b *Branch) toggleClick(*vecty.Event) {
	// dispatch
	/*
		if b.canOpen() {
			b.toggle()
		} else {
			b.Select(false)
		}
	*/
}

func (b *Branch) selectClick(*vecty.Event) {
	// dispatch
	/*
		b.Select(false)
	*/
}

func (b *Branch) render() vecty.Component {

	var childern vecty.List
	for _, c := range b.node.Map {
		childern = append(childern, NewBranch(b.ctx, c.(*editor.Node)))
	}
	for _, c := range b.node.Array {
		childern = append(childern, NewBranch(b.ctx, c.(*editor.Node)))
	}

	return elem.Div(
		prop.Class("node"),
		elem.Anchor(
			prop.Class("toggle plus"),
			event.Click(b.toggleClick),
		),
		elem.Div(
			prop.Class("node-content"),
			elem.Span(
				prop.Class("node-label"),
				event.Click(b.selectClick),
				vecty.Text("foo"),
			),
			elem.Span(
				prop.Class("badge"),
				vecty.Style("display", "none"),
			),
		),
		elem.Div(
			prop.Class("children"),
			//vecty.Style("display", "none"),
			childern,
		),
	)
}
