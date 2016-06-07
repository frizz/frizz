package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/kerr"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"golang.org/x/net/context"
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
	v.notifs = v.app.Nodes.Watch(v.node,
		stores.NodeInitialised,
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
		v.app.Nodes.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *ArrayRowView) render() vecty.Component {

	var err error
	val := ""
	if !v.node.Missing && !v.node.Null {
		val, err = v.node.Type.Id.ValueContext(v.ctx)
		if err != nil {
			v.app.Fail <- kerr.Wrap("MOECUHNHPC", err)
			return nil
		}
	}

	return elem.TableRow(
		elem.TableData(vecty.Text(val)),
		elem.TableData(vecty.Text("")),
	)

}
