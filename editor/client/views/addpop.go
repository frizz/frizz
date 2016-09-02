package views

import (
	"context"

	"github.com/davelondon/kerr"
	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/event"
	"github.com/davelondon/vecty/prop"
	"github.com/gopherjs/gopherjs/js"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/flux"
	"kego.io/system"
	"kego.io/system/node"
)

type AddPopupView struct {
	*View

	model      *models.AddPopupModel
	nameInput  *vecty.Element
	typeSelect *vecty.Element
}

func NewAddPopupView(ctx context.Context) *AddPopupView {
	v := &AddPopupView{}
	v.View = New(ctx, v)
	v.Watch(nil,
		stores.AddPopupChange,
	)
	return v
}

func (v *AddPopupView) Reconcile(old vecty.Component) {
	if old, ok := old.(*AddPopupView); ok {
		v.Body = old.Body
	}
	v.ReconcileBody()
}

func (v *AddPopupView) Receive(notif flux.NotifPayload) {
	defer close(notif.Done)
	v.model = v.App.Misc.AddPopup()
	v.ReconcileBody()
	if v.model.Visible {
		js.Global.Call("$", "#add-modal").Call("modal", "show")
		if v.model.HasName() {
			js.Global.Call("$", "#add-modal-name").Call("focus")
		} else {
			js.Global.Call("$", "#add-modal-type").Call("focus")
		}
	} else {
		js.Global.Call("$", "#add-modal").Call("modal", "hide")
	}
}

func (v *AddPopupView) Render() vecty.Component {
	if v.model == nil || !v.model.Visible {
		return v.modal()
	}

	v.nameInput = nil
	v.typeSelect = nil

	var nameControl, typeControl vecty.Markup
	if v.model.HasName() {
		v.nameInput = elem.Input(
			prop.Class("form-control"),
			prop.ID("add-modal-name"),
		)
		nameControl = elem.Div(
			prop.Class("form-group"),
			elem.Label(
				prop.For("add-modal-name"),
				vecty.Text("name"),
			),
			v.nameInput,
			elem.Paragraph(
				prop.Class("help-block"),
				vecty.Text("Enter a name for the new item"),
			),
		)
	}
	if len(v.model.Types) > 1 {

		var options vecty.List
		for _, t := range v.model.Types {
			displayName, err := t.Id.ValueContext(v.Ctx)
			if err != nil {
				// we shouldn't be able to get here
				v.App.Fail <- kerr.Wrap("NTSWPIEAHC", err)
			}
			options = append(options, elem.Option(
				prop.ID(t.Id.String()),
				vecty.Text(displayName),
			))
		}

		v.typeSelect = elem.Select(
			prop.Class("form-control"),
			prop.ID("add-modal-type"),
			elem.Option(
				prop.ID(""),
				vecty.Text(""),
			),
			options,
		)

		typeControl = elem.Div(
			prop.Class("form-group"),
			elem.Label(
				prop.For("add-modal-type"),
				vecty.Text("type"),
			),
			v.typeSelect,
			elem.Paragraph(
				prop.Class("help-block"),
				vecty.Text("Select a type for the new object"),
			),
		)
	}

	return v.modal(
		elem.Div(
			prop.Class("modal-body"),
			elem.Form(
				event.Submit(func(ev *vecty.Event) {
					v.save()
				}).PreventDefault(),
				nameControl,
				typeControl,
			),
		),
	)
}

func (v *AddPopupView) modal(markup ...vecty.Markup) *vecty.Element {

	return elem.Div(
		prop.ID("add-modal"),
		prop.Class("modal"),
		vecty.Data("backdrop", "static"),
		vecty.Data("keyboard", "false"),
		elem.Div(
			prop.Class("modal-dialog"),
			elem.Div(
				prop.Class("modal-content"),
				elem.Div(
					prop.Class("modal-header"),
					elem.Button(
						prop.Type(prop.TypeButton),
						prop.Class("close"),
						elem.Span(
							vecty.Text("×"),
						),
						event.Click(func(ev *vecty.Event) {
							v.App.Dispatch(&actions.CloseAddPopup{})
						}).PreventDefault(),
					),
					elem.Header4(
						prop.Class("modal-title"),
						vecty.Text("Add item"),
					),
				),
				vecty.List(markup),
				elem.Div(
					prop.Class("modal-footer"),
					elem.Button(
						prop.Type(prop.TypeButton),
						prop.Class("btn btn-default"),
						elem.Span(
							vecty.Text("Close"),
						),
						event.Click(func(ev *vecty.Event) {
							v.App.Dispatch(&actions.CloseAddPopup{})
						}).PreventDefault(),
					),
					elem.Button(
						prop.Type(prop.TypeButton),
						prop.Class("btn btn-primary"),
						elem.Span(
							vecty.Text("Save"),
						),
						event.Click(func(ev *vecty.Event) {
							v.save()
						}).PreventDefault(),
					),
				),
			),
		),
	)
}

func (v *AddPopupView) save() {

	var t *system.Type
	if len(v.model.Types) == 1 {
		t = v.model.Types[0]
	} else {
		options := v.typeSelect.Node().Get("options")
		selectedIndex := v.typeSelect.Node().Get("selectedIndex").Int()
		value := options.Index(selectedIndex).Get("id").String()
		if value == "" {
			return
		}
		r, err := system.NewReferenceFromString(v.Ctx, value)
		if err != nil {
			v.App.Fail <- kerr.Wrap("SEMCIELKRN", err)
			return
		}
		ty, ok := system.GetTypeFromCache(v.Ctx, r.Package, r.Name)
		if !ok {
			v.App.Fail <- kerr.New("RWHSCOFNQM", "Type %s not found in cache", r.Value())
			return
		}
		t = ty
	}

	switch {
	case v.model.IsMap():
		key := v.nameInput.Node().Get("value").String()
		if key == "" {
			// TODO: show an error
			return
		}
		if _, duplicate := v.model.Parent.Map[key]; duplicate {
			// TODO: show an error
			return
		}
		v.App.Dispatch(&actions.Add{
			Undoer: &actions.Undoer{},
			Node:   node.NewNode(),
			Parent: v.model.Parent,
			Key:    key,
			Type:   t,
		})
	case v.model.IsArray():
		v.App.Dispatch(&actions.Add{
			Undoer: &actions.Undoer{},
			Node:   node.NewNode(),
			Parent: v.model.Parent,
			Index:  len(v.model.Parent.Array),
			Type:   t,
		})
	case v.model.IsField():
		v.App.Dispatch(&actions.Add{
			Undoer: &actions.Undoer{},
			Parent: v.model.Parent,
			Node:   v.model.Node,
			Key:    v.model.Node.Key,
			Type:   t,
		})
	case v.model.IsFile():
		name := v.nameInput.Node().Get("value").String()
		if name == "" {
			// TODO: show an error
			return
		}
		// TODO: show an error if a duplicate file exists
		v.App.Dispatch(&actions.Add{
			Undoer:     &actions.Undoer{},
			Node:       node.NewNode(),
			Type:       t,
			BranchName: name,
			BranchFile: name + ".json", // TODO: choose filename based on package
		})
	}

	v.App.Dispatch(&actions.CloseAddPopup{})

}
