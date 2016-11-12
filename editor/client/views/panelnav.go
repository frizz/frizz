package views // import "kego.io/editor/client/views"

import (
	"context"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"kego.io/editor/client/models"
)

type PanelNavView struct {
	*View

	branch   *models.BranchModel
	contents []vecty.MarkupOrComponentOrHTML
}

func NewPanelNavView(ctx context.Context, branch *models.BranchModel) *PanelNavView {
	v := &PanelNavView{}
	v.View = New(ctx, v)
	v.branch = branch
	return v
}

func (v *PanelNavView) Contents(markup ...vecty.MarkupOrComponentOrHTML) *PanelNavView {
	v.contents = markup
	return v
}

func (v *PanelNavView) Render() *vecty.HTML {
	return elem.Navigation(
		prop.Class("navbar navbar-default navbar-static-top nagative-margin"),
		elem.Div(
			prop.Class("container-fluid"),
			elem.Div(
				prop.Class("navbar-header"),
				NewBreadcrumbsView(v.Ctx, v.branch),
			),
			vecty.List(v.contents),
		),
	)
}
