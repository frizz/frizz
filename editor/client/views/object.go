package views

import (
	"fmt"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"github.com/davelondon/vecty/style"
	"golang.org/x/net/context"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system/node"
)

type ObjectView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model *models.EditorModel
}

func NewObjectView(ctx context.Context, node *node.Node) *ObjectView {
	v := &ObjectView{
		ctx: ctx,
		app: stores.FromContext(ctx),
	}
	v.model = v.app.Editors.Get(node)
	v.Mount()
	return v
}

func (v *ObjectView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ObjectView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *ObjectView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *ObjectView) Mount() {
	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *ObjectView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *ObjectView) Unmount() {
	if v.notifs != nil {
		v.app.Editors.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *ObjectView) render() vecty.Component {
	if v.model == nil {
		return elem.Div(vecty.Text("Object (nil)"))
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

		panes = append(panes, elem.Div(
			vecty.ClassMap{
				"tab-pane": true,
				"active":   i == 0,
			},
			prop.ID(fmt.Sprintf("pane%d", i)),
			NewEditorListView(v.ctx, v.model, o),
			NewObjectTableView(v.ctx, v.model, o),
		))
	}

	return elem.Div(
		elem.UnorderedList(
			prop.ID("tabs"),
			prop.Class("nav nav-tabs"),
			style.Margin(style.Size("0 0 15px 0")),
			vecty.Data("tabs", "tabs"),
			tabs,
		),
		elem.Div(
			prop.Class("tab-content"),
			panes,
		),
	)

}
