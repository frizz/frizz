package views

import (
	"fmt"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type BranchView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	node   *node.Node
	branch *models.BranchModel
}

func (b *BranchView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BranchView); ok {
		b.Body = old.Body
		b.node = old.node
		b.branch = old.branch
	}
	b.RenderFunc = b.render
	b.ReconcileBody()
}

func NewBranchView(ctx context.Context, node *node.Node) *BranchView {
	if node == nil {
		return nil
	}
	app := stores.FromContext(ctx)
	b := &BranchView{
		ctx:    ctx,
		app:    app,
		node:   node,
		branch: app.Branches.Get(node),
	}
	return b
}

// Apply implements the vecty.Markup interface.
func (b *BranchView) Apply(element *vecty.Element) {
	element.AddChild(b)
}

func (b *BranchView) toggleClick(*vecty.Event) {
	go func() {
		if b.branch.CanOpen() {
			b.app.Dispatcher.Dispatch(&actions.ToggleNode{Node: b.node})
		} else {
			b.app.Dispatcher.Dispatch(&actions.SelectNode{Node: b.node})
		}
	}()
	/*
		if b.canOpen() {
			b.toggle()
		} else {
			b.Select(false)
		}
	*/
}

func (b *BranchView) selectClick(*vecty.Event) {
	go func() {
		b.app.Dispatcher.Dispatch(&actions.SelectNode{Node: b.node})
	}()
	/*
		b.Select(false)
	*/
}

func (b *BranchView) render() vecty.Component {

	if b.branch == nil {
		return elem.Div()
	}

	var childern vecty.List
	for _, c := range b.node.Map {
		childern = append(childern, NewBranchView(b.ctx, c.(*node.Node)))
	}
	for _, c := range b.node.Array {
		childern = append(childern, NewBranchView(b.ctx, c.(*node.Node)))
	}

	if b.branch.Root {
		return elem.Div(
			prop.Class("node root"),
			elem.Div(
				prop.Class("children"),
				childern,
			),
		)
	}

	return elem.Div(
		prop.Class("node"),
		elem.Anchor(
			vecty.ClassMap{
				"toggle": true,
				"plus":   !b.branch.Open,
				"minus":  b.branch.Open,
			},
			event.Click(b.toggleClick),
		),
		elem.Div(
			prop.Class("node-content"),
			elem.Span(
				prop.Class("node-label"),
				event.Click(b.selectClick),
				vecty.If(b.node.Index > -1, vecty.Text(fmt.Sprintf("[%d]", b.node.Index))),
				vecty.If(b.node.Index == -1, vecty.Text(b.node.Key)),
			),
			elem.Span(
				prop.Class("badge"),
				vecty.Style("display", "none"),
			),
		),
		elem.Div(
			prop.Class("children"),
			vecty.If(!b.branch.Open, vecty.Style("display", "none")),
			childern,
		),
	)
}
