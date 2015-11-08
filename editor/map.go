package editor

import (
	"kego.io/editor/mdl"
	"kego.io/kerr"
	"kego.io/system"
)

type MapEditor struct {
	*Node
	*Common
}

func NewMapEditor(n *Node) *MapEditor {
	return &MapEditor{Node: n, Common: &Common{}}
}

func (e *MapEditor) Layout() Layout {
	return Page
}

var _ Editor = (*MapEditor)(nil)

func (e *MapEditor) Initialize(holder Holder, layout Layout, path string, aliases map[string]string) error {

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

	table.Head("name", "holds", "value")

	for name, item := range e.Map {

		r := table.Row()
		r.Cell().Text(name)
		r.Cell().Text(hold)

		if item.Null {
			r.Cell().Text("")
		} else {
			val, err := item.Type.Id.ValueContext(e.Path, e.Aliases)
			if err != nil {
				return kerr.New("RWHEKAOPHQ", err, "ValueContext")
			}
			r.Cell().Text(val)
		}

	}
	table.Upgrade()
	e.AppendChild(table)

	return nil
}
