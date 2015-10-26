package editor

import (
	"fmt"

	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
	"kego.io/kerr"
	"kego.io/system"
	"kego.io/system/node"
)

type NodeArrayEditor struct {
	*node.Node
	*Common
}

var _ Editor = (*NodeArrayEditor)(nil)

func (e *NodeArrayEditor) Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error {

	e.Panel = panel
	e.Path = path
	e.Aliases = aliases

	table := mdl.Table()

	items, err := system.WrapRule(e.Rule.Interface.(system.CollectionRule).GetItemsRule())
	if err != nil {
		return kerr.New("GQROTGVBXS", err, "NewRuleHolder")
	}
	hold, err := items.HoldsDisplayType(e.Path, e.Aliases)
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
			val, err := item.Type.Id.ValueContext(e.Path, e.Aliases)
			if err != nil {
				return kerr.New("RWHEKAOPHQ", err, "ValueContext")
			}
			values.Cell(val)
		}

	}
	e.Panel.AppendChild(table.Build())

	e.Initialized = true
	return nil
}
