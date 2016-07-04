package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
)

type BreadcrumbsView struct {
	*View

	c      chan struct{}
	branch *models.BranchModel
}

func NewBreadcrumbsView(ctx context.Context, b *models.BranchModel) *BreadcrumbsView {
	v := &BreadcrumbsView{}
	v.View = New(ctx, v)
	v.branch = b
	return v
}

func (v *BreadcrumbsView) Reconcile(old vecty.Component) {
	if old, ok := old.(*BreadcrumbsView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *BreadcrumbsView) Render() vecty.Component {
	if v.branch == nil {
		return elem.Div()
	}

	b := v.branch
	crumbs := vecty.List{}

	for b != nil {

		// copy value of b into new var because it will be used in the click handler
		current := b

		if v.branch.Parent != nil && current.Parent == nil {
			break
		}

		var content vecty.Markup
		if current == v.branch {
			content = vecty.Text(
				current.Contents.Label(v.Ctx),
			)
		} else {
			content = elem.Anchor(
				prop.Href("#"),
				vecty.Text(
					current.Contents.Label(v.Ctx),
				),
				event.Click(func(ev *vecty.Event) {
					v.App.Dispatch(&actions.BranchSelecting{Branch: current, Op: models.BranchOpClickBreadcrumb})
				}).PreventDefault(),
			)
		}

		crumbs = append(
			vecty.List{
				elem.ListItem(
					vecty.ClassMap{
						"active": current == v.branch,
					},
					content,
				),
			},
			crumbs...,
		)
		b = b.Parent
	}
	return elem.OrderedList(
		prop.Class("breadcrumb"),
		crumbs,
	)
}
