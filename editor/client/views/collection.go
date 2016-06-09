package views

import (
	"github.com/davelondon/kerr"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/system/node"
)

func clickSummaryRow(app *stores.App, n *node.Node) {

	app.Dispatch(&actions.FocusNode{
		Node: n,
	})

	b := app.Branches.Get(n)
	if b == nil {
		return
	}
	selectBranch(app, b, models.BranchOpClickSummaryRow, nil)
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
		app.Dispatch(&actions.InitializeNode{
			Node:   node.NewNode(),
			New:    true,
			Parent: parent,
			Index:  len(parent.Array),
			Type:   types[0],
		})
		return
	}

	app.Dispatch(&actions.OpenAddPop{
		Parent: parent,
		Types:  types,
	})

}
