package views // import "kego.io/editor/client/views"

import (
	"code.google.com/p/go.net/context"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"kego.io/context/envctx"
	"kego.io/editor/client/stores"
)

type PageView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	Environment *envctx.Env
}

func NewPage(ctx context.Context, env *envctx.Env) *PageView {
	p := &PageView{
		ctx:         ctx,
		app:         stores.FromContext(ctx),
		Environment: env,
	}

	return p
}

// Apply implements the vecty.Markup interface.
func (p *PageView) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (p *PageView) Reconcile(old vecty.Component) {
	if old, ok := old.(*PageView); ok {
		p.Body = old.Body
		p.Environment = old.Environment
	}
	p.RenderFunc = p.render
	p.ReconcileBody()
}

func (p *PageView) render() vecty.Component {
	return elem.Div(
		prop.ID("wrapper"),
		NewHeader(p.ctx, p.Environment),
		elem.Div(
			prop.Class("wrapper"),
			elem.Div(
				prop.ID("tree"),
				prop.Class("split split-horizontal"),
				NewTreeView(p.ctx),
			),
			elem.Div(
				prop.ID("main"),
				prop.Class("split split-horizontal"),
				NewPanelView(p.ctx),
			),
		),
	)
}
