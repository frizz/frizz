package system

import (
	"fmt"

	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
	"kego.io/kerr"
)

type NodeArrayEditor struct {
	*Node
	*editorCommon
}

var _ Editor = (*NodeArrayEditor)(nil)

func (e *NodeArrayEditor) Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error {

	e.panel = panel
	e.path = path
	e.aliases = aliases

	table := mdl.Table()

	items, err := WrapRule(e.Rule.Interface.(CollectionRule).GetItemsRule())
	if err != nil {
		return kerr.New("GQROTGVBXS", err, "NewRuleHolder")
	}
	hold, err := items.HoldsDisplayType(e.path, e.aliases)
	if err != nil {
		return kerr.New("XDKOSFJVQV", err, "ValueContext")
	}

	index := table.Column("index")
	holds := table.Column("holds")
	values := table.Column("value")

	for i, item := range e.Map {

		index.Cell(fmt.Sprintf("%d", i))
		holds.Cell(hold)

		if item.Null {
			values.Cell("")
		} else {
			val, err := item.Type.Id.ValueContext(e.path, e.aliases)
			if err != nil {
				return kerr.New("RWHEKAOPHQ", err, "ValueContext")
			}
			values.Cell(val)
		}

	}
	e.panel.AppendChild(table.Build())

	e.initialized = true
	return nil
}

func (e *NodeArrayEditor) Update() {
	//if e.Exists {
	//	e.node.ValueString = e.input.Value
	//}
}
