package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/kerr"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system/node"
)

type ArrayRowView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	node *node.Node
}

func NewArrayRowView(ctx context.Context, node *node.Node) *ArrayRowView {
	v := &ArrayRowView{
		ctx:  ctx,
		app:  stores.FromContext(ctx),
		node: node,
	}
	v.Mount()
	return v
}

func (v *ArrayRowView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ArrayRowView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *ArrayRowView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *ArrayRowView) Mount() {
	v.notifs = v.app.Watch(v.node,
		stores.NodeInitialised,
		stores.NodeValueChanged,
	)

	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *ArrayRowView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *ArrayRowView) Unmount() {
	if v.notifs != nil {
		v.app.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *ArrayRowView) render() vecty.Component {

	val, err := v.node.DisplayType(v.ctx)
	if err != nil {
		v.app.Fail <- kerr.Wrap("MOECUHNHPC", err)
		return nil
	}

	return elem.TableRow(
		prop.Class("clickable"),
		event.Click(func(ev *vecty.Event) {
			clickSummaryRow(v.app, v.node)
		}),
		elem.TableData(
			prop.Class("handle"),
			elem.Span(
				prop.Class("glyphicon glyphicon-option-vertical"),
			),
			event.Click(func(e *vecty.Event) {
				// nothing
			}).PreventDefault().StopPropagation(),
		),
		elem.TableData(vecty.Text(v.node.Label(v.ctx))),
		elem.TableData(vecty.Text(val)),
		elem.TableData(elem.Anchor(
			event.Click(func(e *vecty.Event) {
				v.app.Dispatch(&actions.DeleteNode{
					Node: v.node,
				})
			}).PreventDefault().StopPropagation(),
			prop.Href("#"),
			vecty.Text("delete"),
		)),
	)

}
