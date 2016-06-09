package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type ArrayTableView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	node *node.Node
}

func NewArrayTableView(ctx context.Context, node *node.Node) *ArrayTableView {
	v := &ArrayTableView{
		ctx:  ctx,
		app:  stores.FromContext(ctx),
		node: node,
	}
	v.Mount()
	return v
}

func (v *ArrayTableView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ArrayTableView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *ArrayTableView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *ArrayTableView) Mount() {}

func (v *ArrayTableView) Unmount() {
	v.Body.Unmount()
}

func (v *ArrayTableView) render() vecty.Component {

	if v.node == nil || len(v.node.Array) == 0 {
		return elem.Div()
	}

	head := elem.TableRow(
		elem.TableHeader(vecty.Text("value")),
		elem.TableHeader(vecty.Text("options")),
	)

	rows := vecty.List{}
	for _, c := range v.node.Array {
		rows = append(rows, NewArrayRowView(v.ctx, c))
	}

	return elem.Div(
		elem.Table(
			prop.Class("table table-hover"),
			elem.TableHead(
				head,
			),
			elem.TableBody(
				rows,
			),
		),
	)
}
