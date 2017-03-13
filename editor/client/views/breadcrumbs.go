package views

import (
	"context"

	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/event"
	"github.com/dave/vecty/prop"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
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
	if nci, ok := b.Contents.(models.NodeContentsInterface); ok {
		v.Watch(nci.GetNode(),
			stores.NodeDescendantChanged,
		)
	}
	return v
}

func (v *BreadcrumbsView) Render() *vecty.HTML {
	if v.branch == nil {
		return elem.Div()
	}

	b := v.branch
	crumbs := vecty.List{}

	for b != nil {

		// copy value of b into new var because it will be used in the click
		// handler
		current := b

		if v.branch.Parent != nil && current.Parent == nil {
			break
		}

		var content vecty.List
		if current == v.branch {
			content = append(content,
				vecty.Text(
					current.Contents.Label(v.Ctx),
				),
			)
		} else {
			content = append(content,
				elem.Anchor(
					prop.Href("#"),
					vecty.Text(
						current.Contents.Label(v.Ctx),
					),
					event.Click(func(ev *vecty.Event) {
						v.App.Dispatch(&actions.BranchSelecting{Branch: current, Op: models.BranchOpClickBreadcrumb})
					}).PreventDefault(),
				),
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
