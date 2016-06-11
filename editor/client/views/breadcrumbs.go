package views

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
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
				current.Contents.Label(v.ctx),
			)
		} else {
			content = elem.Anchor(
				prop.Href("#"),
				vecty.Text(
					current.Contents.Label(v.ctx),
				),
				event.Click(func(ev *vecty.Event) {
					v.app.Dispatch(&actions.BranchSelecting{Branch: current, Op: models.BranchOpClickBreadcrumb})
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
