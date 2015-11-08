package editor

import (
	"fmt"

	"kego.io/editor/mdl"
	"kego.io/kerr"
	"kego.io/system"
)

type ArrayEditor struct {
	*Node
	*Common
}

func NewArrayEditor(n *Node) *ArrayEditor {
	return &ArrayEditor{Node: n, Common: &Common{}}
}

func (e *ArrayEditor) Layout() Layout {
	return Page
}

var _ Editor = (*ArrayEditor)(nil)

func (e *ArrayEditor) Initialize(holder Holder, layout Layout, path string, aliases map[string]string) error {

	e.Common.Initialize(holder, layout, path, aliases)

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

	for i, item := range e.Array {

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
	e.AppendChild(table.Build())

	return nil
}
