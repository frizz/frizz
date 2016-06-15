package views

import (
	"github.com/davelondon/kerr"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
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

func addCollectionItem(app *stores.App, parent *node.Node) {

	if parent == nil {
		return
	}

	if !parent.Type.IsNativeCollection() {
		return
	}

	rw, err := parent.Rule.ItemsRule()
	if err != nil {
		app.Fail <- kerr.Wrap("EWYOMNAQMU", err)
	}

	types := rw.PermittedTypes()

	if len(types) == 1 && parent.Type.IsNativeArray() {
		// if only one type is compatible and adding to an array, don't show the popup, just
		// add it.
		rule, err := parent.Rule.ItemsRule()
		if err != nil {
			app.Fail <- kerr.Wrap("ROVRPRAMFF", err)
		}
		app.Dispatch(&actions.InitializeNode{
			Node:   node.NewNode(),
			Parent: parent,
			Rule:   rule,
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
