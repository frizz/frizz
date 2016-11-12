package views

import (
	"context"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"github.com/gopherjs/gopherjs/js"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system"
	"kego.io/system/node"
)

type EditorListView struct {
	*View

	model     *models.EditorModel
	filter    *system.Reference
	container *vecty.HTML
	sort      bool
	items     int
	exclude   []string
}

func (v *EditorListView) Items() int {
	return v.items
}

func NewEditorListView(ctx context.Context, model *models.EditorModel, filter *system.Reference, exclude []string) *EditorListView {
	v := &EditorListView{}
	v.View = New(ctx, v)
	v.model = model
	v.filter = filter
	v.exclude = exclude
	v.sort = model.Node.JsonType == system.J_ARRAY
	v.Watch(v.model.Node,
		stores.NodeArrayReorder,
		stores.NodeChildAdded,
		stores.NodeChildDeleted,
	)
	return v
}

func (v *EditorListView) Restore(prev vecty.Component) bool {
	if v.sort {
		v.sortable()
	}
	return false
}

func (v *EditorListView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	if v.sort && notif.Type == stores.NodeArrayReorder {
		js.Global.Call("$", v.container.Node).Call("sortable", "cancel")
	}
	vecty.Rerender(v)
	if v.sort {
		v.sortable()
	}
}

func (v *EditorListView) sortable() {
	if v.container == nil {
		return
	}
	js.Global.Call("$", v.container.Node).Call("sortable", js.M{
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

func (v *EditorListView) Render() *vecty.HTML {

	if v.model == nil || v.model.Node.Missing || v.model.Node.Null {
		return elem.Div(vecty.Text("editor (nil)"))
	}

	children := vecty.List{}

	add := func(n *node.Node) {

		for _, v := range v.exclude {
			if n.Key == v {
				return
			}
		}

		if n.Null || n.Missing {
			children = append(children, nullEditor(v.Ctx, n, v.App))
			return
		}

		f := editable.Branch
		if e := models.GetEditable(v.Ctx, n); e != nil {
			f = e.Format(n.Rule)
			if f == editable.Block || f == editable.Inline {
				children = append(children, NewEditorView(v.Ctx, n).Controls(e.EditorView(v.Ctx, n, editable.Block)))
				return
			}
		}
		if f == editable.Branch {
			b := v.App.Branches.Get(n)
			children = append(children, NewEditorView(v.Ctx, n).Icons(
				elem.Anchor(
					prop.Href("#"),
					event.Click(func(e *vecty.Event) {
						v.App.Dispatch(&actions.BranchSelecting{Branch: b, Op: models.BranchOpClickEditorLink})
					}).PreventDefault(),
					elem.Italic(
						prop.Class("editor-icon editor-icon-after glyphicon glyphicon-share-alt"),
					),
				),
			))
		}
	}

	for _, n := range v.model.Node.Map {
		// TODO: hide optional fields
		//if n.Missing || n.Null {
		//	continue
		//}
		if v.filter != nil && *n.Origin != *v.filter {
			continue
		}
		add(n)
	}

	for _, n := range v.model.Node.Array {
		add(n)
	}

	v.items = len(children)

	if len(children) == 0 {
		v.container = elem.Div()
	} else {
		v.container = elem.Form(
			event.Submit(func(*vecty.Event) {}).PreventDefault(),
			children,
		)
	}

	return v.container
}
