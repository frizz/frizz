package views // import "kego.io/editor/client/views"

import (
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"github.com/gopherjs/gopherjs/js"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
)

type ArrayTableView struct {
	*View

	model *models.EditorModel
	tbody *vecty.Element
}

func NewArrayTableView(ctx context.Context, model *models.EditorModel) *ArrayTableView {
	v := &ArrayTableView{}
	v.View = New(ctx, v)
	v.model = model
	v.Watch(v.model.Node,
		stores.NodeArrayReorder,
		stores.NodeChildAdded,
		stores.NodeChildDeleted,
	)
	return v
}

func (v *ArrayTableView) Reconcile(old vecty.Component) {
	if old, ok := old.(*ArrayTableView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
	v.sortable()
}

func (v *ArrayTableView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	if notif.Type == stores.NodeArrayReorder {
		js.Global.Call("$", v.tbody.Node()).Call("sortable", "cancel")
	}
	v.ReconcileBody()
	v.sortable()
}

func (v *ArrayTableView) sortable() {
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
				ui.Get("item").Call("data", "end_pos", ui.Get("placeholder").Call("index"))
			},
			"update": func(event *js.Object, ui *js.Object) {
				oldIndex := ui.Get("item").Call("data", "start_pos").Int()
				newIndex := ui.Get("item").Call("data", "end_pos").Int() - 1
				if oldIndex == newIndex {
					return
				}
				v.App.Dispatch(&actions.Reorder{
					Undoer: &actions.Undoer{},
					Model:  v.model,
					Before: oldIndex,
					After:  newIndex,
				})
			},
		})
	}
}

func (v *ArrayTableView) Render() vecty.Component {

	if v.model.Node == nil || len(v.model.Node.Array) == 0 {
		v.tbody = nil
		return elem.Div()
	}

	head := elem.TableRow(
		elem.TableHeader(
			prop.Class("handle-head"),
			vecty.Text(""),
		),
		elem.TableHeader(vecty.Text("label")),
		elem.TableHeader(vecty.Text("value")),
		elem.TableHeader(vecty.Text("options")),
	)

	rows := vecty.List{}
	for _, c := range v.model.Node.Array {
		rows = append(rows, NewArrayRowView(v.Ctx, c))
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
