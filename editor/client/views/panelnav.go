package views // import "frizz.io/editor/client/views"

import (
	"context"

	"frizz.io/editor/client/models"
	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/prop"
)

type PanelNavView struct {
	*View

	branch   *models.BranchModel
	contents func() vecty.MarkupOrComponentOrHTML
}

func NewPanelNavView(ctx context.Context, branch *models.BranchModel) *PanelNavView {
	v := &PanelNavView{}
	v.View = New(ctx, v)
	v.branch = branch
	v.contents = func() vecty.MarkupOrComponentOrHTML { return nil }
	return v
}

func (v *PanelNavView) Contents(markup func() vecty.MarkupOrComponentOrHTML) *PanelNavView {
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
			v.contents(),
		),
	)
}
