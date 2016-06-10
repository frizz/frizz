package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"github.com/gopherjs/gopherjs/js"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

type ArrayTableView struct {
	vecty.Composite
	ctx context.Context
	app *stores.App

	node  *node.Node
	tbody *vecty.Element
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
	if v.tbody != nil {
		js.Global.Call("$", v.tbody.Node()).Call("sortable", js.M{
			"handle":               ".handle",
			"axis":                 "y",
			"forcePlaceholderSize": true,
			"placeholder":          "drag-placeholder",
			"start": func(event *js.Object, ui *js.Object) {
				ui.Get("item").Call("data", "start_pos", ui.Get("item").Call("index"))
			},
			"beforeStop": func(event *js.Object, ui *js.Object) {
				oldIndex := ui.Get("item").Call("data", "start_pos").Int()
				newIndex := ui.Get("placeholder").Call("index").Int() - 1
				if oldIndex == newIndex {
					return
				}
				v.app.Dispatch(&actions.ArrayOrder{
					Parent:   v.node,
					OldIndex: oldIndex,
					NewIndex: newIndex,
				})
			},
		})
	}
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
		v.tbody = nil
		return elem.Div()
	}

	head := elem.TableRow(
		elem.TableHeader(
			prop.Class("handle-head"),
			vecty.Text(""),
		),
		elem.TableHeader(vecty.Text("value")),
		elem.TableHeader(vecty.Text("options")),
	)

	rows := vecty.List{}
	for _, c := range v.node.Array {
		rows = append(rows, NewArrayRowView(v.ctx, c))
	}

	v.tbody = elem.TableBody(
		rows,
	)

	return elem.Div(
		elem.Table(
			prop.Class("table"),
			elem.TableHead(
				head,
			),
			v.tbody,
		),
	)
}
