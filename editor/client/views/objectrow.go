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

type ObjectRowView struct {
	*View

	node *node.Node
}

func NewObjectRowView(ctx context.Context, node *node.Node) *ObjectRowView {
	v := &ObjectRowView{}
	v.View = New(ctx, v)
	v.node = node
	v.Watch(v.node,
		stores.NodeInitialised,
		stores.NodeValueChanged,
		stores.NodeDescendantValueChanged,
	)
	return v
}

func (v *ObjectRowView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ObjectRowView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *ObjectRowView) Render() vecty.Component {

	name := v.node.Key

	hold, err := v.node.Rule.DisplayType()
	if err != nil {
		v.App.Fail <- kerr.Wrap("CETBRLENSP", err)
		return nil
	}

	val, err := v.node.DisplayType(v.Ctx)
	if err != nil {
		v.App.Fail <- kerr.Wrap("AWLAMTFJSO", err)
		return nil
	}

	var add, delete vecty.Component
	if v.node.Missing || v.node.Null {
		add = elem.Anchor(
			event.Click(func(e *vecty.Event) {
				types := v.node.Rule.PermittedTypes()
				if len(types) == 1 {
					// if only one type is compatible, don't show the popup, just add it.
					v.App.Dispatch(&actions.Add{
						Undoer: &actions.Undoer{},
						Parent: v.node.Parent,
						Node:   v.node,
						Key:    v.node.Key,
						Type:   types[0],
					})
					return
				}
				v.App.Dispatch(&actions.OpenAddPopup{
					Parent: v.node.Parent,
					Node:   v.node,
					Types:  types,
				})
			}).PreventDefault().StopPropagation(),
			prop.Href("#"),
			vecty.Text("add"),
		)
	} else {
		delete = elem.Anchor(
			event.Click(func(e *vecty.Event) {
				v.App.Dispatch(&actions.Delete{
					Undoer: &actions.Undoer{},
					Node:   v.node,
					Parent: v.node.Parent,
				})
			}).PreventDefault().StopPropagation(),
			prop.Href("#"),
			vecty.Text("delete"),
		)
	}

	return elem.TableRow(
		vecty.ClassMap{
			"clickable": !v.node.Missing && !v.node.Null,
		},
		event.Click(func(ev *vecty.Event) {
			if !v.node.Missing && !v.node.Null {
				clickSummaryRow(v.App, v.node)
			}
		}),
		elem.TableData(vecty.Text(name)),
		elem.TableData(vecty.Text(hold)),
		elem.TableData(vecty.Text(val)),
		elem.TableData(vecty.List{add, delete}),
	)

}
