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

type ObjectRowView struct {
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	node *node.Node
}

func NewObjectRowView(ctx context.Context, node *node.Node) *ObjectRowView {
	v := &ObjectRowView{
		ctx:  ctx,
		app:  stores.FromContext(ctx),
		node: node,
	}
	v.Mount()
	return v
}

func (v *ObjectRowView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ObjectRowView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *ObjectRowView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *ObjectRowView) Mount() {
	v.notifs = v.app.Nodes.Watch(v.node,
		stores.NodeInitialised,
	)

	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *ObjectRowView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.ReconcileBody()
}

func (v *ObjectRowView) Unmount() {
	if v.notifs != nil {
		v.app.Nodes.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *ObjectRowView) render() vecty.Component {

	name := v.node.Key

	hold, err := v.node.Rule.DisplayType()
	if err != nil {
		v.app.Fail <- kerr.Wrap("CETBRLENSP", err)
		return nil
	}

	val, err := v.node.DisplayType(v.ctx)
	if err != nil {
		v.app.Fail <- kerr.Wrap("AWLAMTFJSO", err)
		return nil
	}

	var add vecty.Component
	if v.node.Missing || v.node.Null {
		add = elem.Anchor(
			event.Click(func(e *vecty.Event) {
				v.app.Dispatch(&actions.AddNodeClick{Node: v.node})
			}).PreventDefault(),
			prop.Href("#"),
			vecty.Text("add"),
		)
	}

	return elem.TableRow(
		elem.TableData(vecty.Text(name)),
		elem.TableData(vecty.Text(hold)),
		elem.TableData(vecty.Text(val)),
		elem.TableData(add),
	)

}
