package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
)

type BreadcrumbsView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App
	c   chan struct{}

	branch *models.BranchModel
}

func NewBreadcrumbsView(ctx context.Context, b *models.BranchModel) *BreadcrumbsView {
	v := &BreadcrumbsView{
		ctx:    ctx,
		app:    stores.FromContext(ctx),
		branch: b,
	}
	return v
}

func (v *BreadcrumbsView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BreadcrumbsView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *BreadcrumbsView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *BreadcrumbsView) render() vecty.Component {
	if v.branch == nil {
		return elem.Div()
	}
	return elem.Div(
		vecty.Text("Breadcrumbs " + v.branch.Contents.Label()),
	)
}
