package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type MapTableView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	node *node.Node
}

func NewMapTableView(ctx context.Context, node *node.Node) *MapTableView {
	v := &MapTableView{
		ctx:  ctx,
		app:  stores.FromContext(ctx),
		node: node,
	}
	v.Mount()
	return v
}

func (v *MapTableView) Reconcile(old vecty.Component) {
	if old, ok := old.(*MapTableView); ok {
		v.Body = old.Body
	}
	v.RenderFunc = v.render
	v.ReconcileBody()
}

// Apply implements the vecty.Markup interface.
func (v *MapTableView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *MapTableView) Mount() {}

func (v *MapTableView) Unmount() {
	v.Body.Unmount()
}

func (v *MapTableView) render() vecty.Component {

	if v.node == nil || len(v.node.Map) == 0 {
		return elem.Div()
	}

	head := elem.TableRow(
		elem.TableHeader(vecty.Text("name")),
		elem.TableHeader(vecty.Text("value")),
		elem.TableHeader(vecty.Text("options")),
	)

	rows := vecty.List{}
	for _, c := range v.node.Map {
		rows = append(rows, NewMapRowView(v.ctx, c))
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
