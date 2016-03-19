package views // import "kego.io/editor/client/views"

import (
	"code.google.com/p/go.net/context"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"kego.io/editor/client/flux"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type TreeView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	Root     *node.Node
	Selected *node.Node
}

func NewTreeView(ctx context.Context) *TreeView {
	p := &TreeView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}

	go func() {
		for range flux.Changed(p.app.Branches, p.app.Selected) {
			p.Root = p.app.Root.Get()
			p.ReconcileBody()
		}
	}()

	return p
}

// Apply implements the vecty.Markup interface.
func (p *TreeView) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (v *TreeView) Reconcile(old vecty.Component) {
	if old, ok := old.(*TreeView); ok {
		v.Body = old.Body
		v.Selected = old.Selected
		v.Root = old.Root
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

func (p *TreeView) render() vecty.Component {
	return elem.Div(
		prop.Class("content"),
		vecty.If(
			p.Root != nil,
			NewBranchView(p.ctx, p.Root),
		),
	)
}
