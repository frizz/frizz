package views

import (
	"context"

	"frizz.io/context/sysctx"
	"frizz.io/editor/client/actions"
	"frizz.io/editor/client/editable"
	"frizz.io/editor/client/models"
	"frizz.io/editor/client/stores"
	"frizz.io/system"
	"frizz.io/system/node"
	"github.com/dave/kerr"
	"github.com/dave/vecty"
	"github.com/dave/vecty/elem"
	"github.com/dave/vecty/event"
	"github.com/dave/vecty/prop"
)

type StructFragmentView struct {
	*View

	model      *models.EditorModel
	origin     *system.Reference
	originType *system.Type
}

func NewStructFragmentView(ctx context.Context, node *node.Node, origin *system.Reference) *StructFragmentView {
	v := &StructFragmentView{}
	v.View = New(ctx, v)
	v.model = v.App.Editors.Get(node)
	v.origin = origin
	if ti, ok := sysctx.FromContext(ctx).GetType(origin.Package, origin.Name); ok {
		v.originType = ti.(*system.Type)
	} else {
		v.App.Fail <- kerr.New("DNMXCJXNVU", "Type %s not found in system context", origin.String())
		return nil
	}
	return v
}

func (v *StructFragmentView) Render() *vecty.HTML {
	if v.model == nil {
		return elem.Div(vecty.Text("StructFragment (nil)"))
	}

	if v.model.Node == nil {
		return elem.Div(vecty.Text("StructFragment (node == nil)"))
	}

	if v.model.Node.Missing || v.model.Node.Null {
		return elem.Div(nullEditor(v.Ctx, v.model.Node, v.App))
	}

	// Get the custom editor for the main type. If it wants to edit more of the
	// embedded structs, it will implement editable.EditsExtraEmbeddedTypes.
	var embeddedTypes []*system.Reference
	mainEditor, err := models.GetEmbedEditable(v.Ctx, v.model.Node, v.origin)
	if err != nil {
		v.App.Fail <- kerr.Wrap("NCBVRLKXED", err)
		return nil
	}
	if mainEditor != nil {
		embeddedTypes = append(embeddedTypes, v.origin)
		if emo, ok := mainEditor.(editable.EditsExtraEmbeddedTypes); ok {
			embeddedTypes = append(embeddedTypes, emo.ExtraEmbeddedTypes()...)
		}
	}

	isEditedByMainEditor := func(o *system.Reference) bool {
		for _, et := range embeddedTypes {
			if *o == *et {
				return true
			}
		}
		return false
	}

	out := vecty.List{}

	for _, embed := range v.originType.Embed {
		if isEditedByMainEditor(embed) {
			continue
		}
		out = append(out, NewStructFragmentView(v.Ctx, v.model.Node, embed))
	}

	if mainEditor != nil {
		out = append(out, mainEditor.EditorView(v.Ctx, v.model.Node, editable.Block))
	} else {
		out = append(out, NewEditorListView(v.Ctx, v.model, v.origin, nil))
	}

	return elem.Div(
		out,
	)

}

func nullEditor(ctx context.Context, n *node.Node, app *stores.App) *EditorView {
	add := func(e *vecty.Event) {
		types := n.Rule.PermittedTypes()
		if len(types) == 1 {
			// if only one type is compatible, don't show the
			// popup, just add it.
			app.Dispatch(&actions.Add{
				Undoer: &actions.Undoer{},
				Parent: n.Parent,
				Node:   n,
				Key:    n.Key,
				Type:   types[0],
			})
			return
		}
		app.Dispatch(&actions.OpenAddPopup{
			Parent: n.Parent,
			Node:   n,
			Types:  types,
		})
	}
	return NewEditorView(ctx, n).Icons(
		func() vecty.MarkupOrComponentOrHTML {
			return elem.Anchor(
				prop.Href("#"),
				event.Click(add).PreventDefault(),
				elem.Italic(
					prop.Class("editor-icon editor-icon-after glyphicon glyphicon-plus-sign"),
				),
			)
		},
	).Dropdown(
		func() vecty.MarkupOrComponentOrHTML {
			return elem.ListItem(
				elem.Anchor(
					prop.Href("#"),
					vecty.Text("Add"),
					event.Click(add).PreventDefault(),
				),
			)
		},
	)
}
