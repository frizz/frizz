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
	*View

	node *node.Node
}

func NewArrayRowView(ctx context.Context, node *node.Node) *ArrayRowView {
	v := &ArrayRowView{}
	v.View = New(ctx, v)
	v.node = node
	v.Mount()
	return v
}

func (v *ArrayRowView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ArrayRowView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *ArrayRowView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *ArrayRowView) Mount() {
	v.Notifs = v.App.Watch(v.node,
		stores.NodeValueChanged,
	)

	go func() {
		for notif := range v.Notifs {
			v.reaction(notif)
		}
	}()
}

func (v *ArrayRowView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *ArrayRowView) Unmount() {
	if v.Notifs != nil {
		v.App.Delete(v.Notifs)
		v.Notifs = nil
	}
	v.Body.Unmount()
}

func (v *ArrayRowView) Render() vecty.Component {

	val, err := v.node.DisplayType(v.Ctx)
	if err != nil {
		v.App.Fail <- kerr.Wrap("MOECUHNHPC", err)
		return nil
	}

	return elem.TableRow(
		prop.Class("clickable"),
		event.Click(func(ev *vecty.Event) {
			clickSummaryRow(v.App, v.node)
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
		elem.TableData(vecty.Text(v.node.Label(v.Ctx))),
		elem.TableData(vecty.Text(val)),
		elem.TableData(elem.Anchor(
			event.Click(func(e *vecty.Event) {
				v.App.Dispatch(&actions.DeleteNode{
					Node: v.node,
				})
			}).PreventDefault().StopPropagation(),
			prop.Href("#"),
			vecty.Text("delete"),
		)),
	)

}
