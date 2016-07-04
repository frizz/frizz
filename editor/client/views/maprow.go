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

type MapRowView struct {
	*View

	node *node.Node
}

func NewMapRowView(ctx context.Context, node *node.Node) *MapRowView {
	v := &MapRowView{}
	v.View = New(ctx, v)
	v.node = node
	v.Watch(v.node,
		stores.NodeValueChanged,
	)
	return v
}

func (v *MapRowView) Reconcile(old vecty.Component) {
	if old, ok := old.(*MapRowView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *MapRowView) Render() vecty.Component {

	name := v.node.Key

	val, err := v.node.DisplayType(v.Ctx)
	if err != nil {
		v.App.Fail <- kerr.Wrap("NPJIEIKJVK", err)
		return nil
	}

	return elem.TableRow(
		prop.Class("clickable"),
		event.Click(func(ev *vecty.Event) {
			clickSummaryRow(v.App, v.node)
		}),
		elem.TableData(vecty.Text(name)),
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
