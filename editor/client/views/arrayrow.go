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
	v.Watch(v.node,
		stores.NodeValueChanged,
	)
	return v
}

func (v *ArrayRowView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ArrayRowView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
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
