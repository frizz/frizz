package views

import (
	"fmt"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system/node"
)

type CompositeView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model *models.EditorModel
}

func NewCompositeView(ctx context.Context, node *node.Node) *CompositeView {
	v := &CompositeView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.model = v.app.Editors.Get(node)
	v.Mount()
	return v
}

func (v *CompositeView) Reconcile(old vecty.Component) {
	if old, ok := old.(*CompositeView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *CompositeView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *CompositeView) Mount() {
	v.notifs = v.app.Nodes.Watch(v.model.Node,
		stores.NodeInitialised,
	)
	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *CompositeView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *CompositeView) Unmount() {
	if v.notifs != nil {
		v.app.Nodes.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *CompositeView) render() vecty.Component {
	if v.model == nil {
		return elem.Div(vecty.Text("Composite (nil)"))
	}

	tabs := vecty.List{}
	panes := vecty.List{}
	for i, o := range v.model.Node.Type.FieldOrigins() {

		tabs = append(tabs, elem.ListItem(
			vecty.ClassMap{
				"active": i == 0,
			},
			elem.Anchor(
				prop.Href(fmt.Sprintf("#pane%d", i)),
				vecty.Data("toggle", "tab"),
				vecty.Text(o.Name),
			),
		))

		var editor vecty.Component
		if !v.model.Node.Missing && !v.model.Node.Null {
			editor = NewObjectView(v.ctx, v.model.Node, o)
		} else {
			editor = elem.Div(vecty.Text("null"))
		}

		panes = append(panes, elem.Div(
			vecty.ClassMap{
				"tab-pane": true,
				"active":   i == 0,
			},
			prop.ID(fmt.Sprintf("pane%d", i)),
			elem.Header1(
				vecty.Text(o.Name),
			),
			editor,
			NewSummaryView(v.ctx, v.model.Node, o),
		))
	}

	return elem.Div(
		elem.UnorderedList(
			prop.ID("tabs"),
			prop.Class("nav nav-tabs"),
			vecty.Data("tabs", "tabs"),
			tabs,
		),
		elem.Div(
			prop.Class("tab-content"),
			panes,
		),
	)

}
