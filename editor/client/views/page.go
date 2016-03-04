package views // import "kego.io/editor/client/views"

import (
	"code.google.com/p/go.net/context"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"kego.io/context/envctx"
	"kego.io/editor"
	"kego.io/editor/client/stores"
)

type Page struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	Environment *envctx.Env
	Package     *editor.Node
}

func NewPage(ctx context.Context, env *envctx.Env, pkg *editor.Node) *Page {
	return &Page{
		ctx:         ctx,
		app:         stores.FromContext(ctx),
		Environment: env,
		Package:     pkg,
	}
}

// Apply implements the vecty.Markup interface.
func (p *Page) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (p *Page) Reconcile(old vecty.Component) {
	if old, ok := old.(*Page); ok {
		p.Body = old.Body
		p.Environment = old.Environment
		p.Package = old.Package
	}
	p.RenderFunc = p.render
	p.ReconcileBody()
}

func (p *Page) render() vecty.Component {
	return elem.Div(
		prop.ID("wrapper"),
		NewHeader(p.ctx, p.Environment),
		elem.Div(
			prop.Class("wrapper"),
			elem.Div(
				prop.ID("tree"),
				prop.Class("split split-horizontal"),
				elem.Div(
					prop.Class("content"),
					vecty.Text("Sidebar"),
				),
			),
			elem.Div(
				prop.ID("main"),
				prop.Class("split split-horizontal"),
				elem.Div(
					prop.Class("content"),
					vecty.Text("Content"),
				),
			),
		),
	)
}
