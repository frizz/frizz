package views

import (
	"context"

	"github.com/dave/kerr"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/system"
	"kego.io/system/node"
)

func clickSummaryRow(app *stores.App, n *node.Node) {

	if e := app.Editors.Get(n); e != nil {
		app.Dispatch(&actions.EditorFocus{
			Editor: e,
		})
	}

	if b := app.Branches.Get(n); b != nil {
		app.Dispatch(&actions.BranchSelecting{
			Branch: b,
			Op:     models.BranchOpClickSummaryRow,
		})
	}
}

func addCollectionItem(ctx context.Context, app *stores.App, parent *node.Node) {

	if parent == nil {
		return
	}

	if !parent.Type.IsNativeCollection() {
		return
	}

	var rw *system.RuleWrapper

	var err error
	if parent.Type.IsAliasCollection() {
		rw = system.WrapRule(ctx, parent.Type.Alias.(system.CollectionRule).GetItemsRule())
	} else {
		if rw, err = parent.Rule.ItemsRule(); err != nil {
			app.Fail <- kerr.Wrap("EWYOMNAQMU", err)
		}
	}

	types := rw.PermittedTypes()

	if len(types) == 1 && parent.Type.IsNativeArray() {
		// if only one type is compatible and adding to an array, don't show the popup, just
		// add it.
		app.Dispatch(&actions.Add{
			Undoer: &actions.Undoer{},
			Node:   node.NewNode(),
			Parent: parent,
			Index:  len(parent.Array),
			Type:   types[0],
		})
		return
	}

	app.Dispatch(&actions.OpenAddPopup{
		Parent: parent,
		Types:  types,
	})

}
