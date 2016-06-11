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
	vecty.Composite
	ctx    context.Context
	app    *stores.App
	notifs chan flux.NotifPayload

	model *models.EditorModel
	tbody *vecty.Element
}

func NewArrayTableView(ctx context.Context, model *models.EditorModel) *ArrayTableView {
	v := &ArrayTableView{
		ctx:   ctx,
		app:   stores.FromContext(ctx),
		model: model,
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
	v.sortable()
}

// Apply implements the vecty.Markup interface.
func (v *ArrayTableView) Apply(element *vecty.Element) {
	element.AddChild(v)
}

func (v *ArrayTableView) Mount() {
	v.notifs = v.app.Watch(v.model,
		stores.EditorArrayOrderChanged,
		stores.EditorChildAdded,
		stores.EditorChildDeleted,
	)
	go func() {
		for notif := range v.notifs {
			v.reaction(notif)
		}
	}()
}

func (v *ArrayTableView) reaction(notif flux.NotifPayload) {
	defer close(notif.Done)
	if notif.Type == stores.EditorArrayOrderChanged {
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
				v.app.Dispatch(&actions.ArrayOrder{
					Model:    v.model,
					OldIndex: oldIndex,
					NewIndex: newIndex,
				})
			},
		})
	}
}

func (v *ArrayTableView) Unmount() {
	if v.notifs != nil {
		v.app.Delete(v.notifs)
		v.notifs = nil
	}
	v.Body.Unmount()
}

func (v *ArrayTableView) render() vecty.Component {

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
