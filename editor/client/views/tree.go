package views // import "kego.io/editor/client/views"

import (
	"code.google.com/p/go.net/context"
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/prop"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type TreeView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	root     *models.BranchModel
	Package  *node.Node
	Types    map[string]*node.Node
	Data     map[string]*models.DataModel
	Selected *models.BranchModel
}

func (v *TreeView) Reconcile(old vecty.Component) {
	if old, ok := old.(*TreeView); ok {
		v.Body = old.Body
		v.Selected = old.Selected
		v.Package = old.Package
		v.Types = old.Types
		v.Data = old.Data
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

func NewTreeView(ctx context.Context) *TreeView {
	p := &TreeView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}

	go func() {
		branches := p.app.Branches.Changed()
		selected := p.app.Selected.Changed()
		types := p.app.Types.Changed()
		data := p.app.Data.Changed()
		pkg := p.app.Package.Changed()
		for {
			select {
			case <-selected:
				p.Selected = p.app.Selected.Get()
				p.ReconcileBody()
			case <-branches:
				p.root = p.app.Branches.Root()
				p.ReconcileBody()
			case <-types:
				p.Types = p.app.Types.All()
				p.ReconcileBody()
			case <-data:
				p.Data = p.app.Data.All()
				p.ReconcileBody()
			case <-pkg:
				p.Package = p.app.Package.Get()
				p.ReconcileBody()
			}
		}
	}()

	return p
}

// Apply implements the vecty.Markup interface.
func (p *TreeView) Apply(element *vecty.Element) {
	element.AddChild(p)
}

func (p *TreeView) render() vecty.Component {
	if p.root == nil {
		return elem.Div()
	}
	return elem.Div(
		prop.Class("content"),
		vecty.If(
			p.Package != nil,
			NewBranchView(p.ctx, p.root),
		),
	)
}
