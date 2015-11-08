package editor

import (
	"kego.io/editor/mdl"
	"kego.io/kerr"
)

type ObjectEditor struct {
	*Node
	*Common
}

func NewObjectEditor(n *Node) *ObjectEditor {
	return &ObjectEditor{Node: n, Common: &Common{}}
}

func (e *ObjectEditor) Layout() Layout {
	return Page
}

var _ Editor = (*ObjectEditor)(nil)

func (e *ObjectEditor) Initialize(holder Holder, layout Layout, path string, aliases map[string]string) error {

	e.Common.Initialize(holder, layout, path, aliases)

	e.initializeBlockEditors()

	if err := e.initializeTable(); err != nil {
		return kerr.New("KAVTMDDFYW", err, "initializeTable")
	}

	return nil
}

func (e *ObjectEditor) initializeBlockEditors() {
	for _, field := range e.Fields {
		node := Node{field}
		ed := node.Editor()
		if ed == nil || ed.Layout() == Page {
			continue
		}
		e.Editors = append(e.Editors, ed)
		ed.Initialize(e.holder, Block, e.Path, e.Aliases)
		e.AppendChild(ed)
	}
}

func (e *ObjectEditor) initializeTable() error {

	table := mdl.Table()

	table.Head("name", "origin", "holds", "value")

	for name, field := range e.Fields {

		r := table.Row()

		r.Cell().Text(name)

		origin, err := field.Origin.ValueContext(e.Path, e.Aliases)
		if err != nil {
			return kerr.New("ACQLJXWYQX", err, "ValueContext")
		}
		r.Cell().Text(origin)

		hold, err := field.Rule.HoldsDisplayType(e.Path, e.Aliases)
		if err != nil {
			return kerr.New("OYMARPFDGA", err, "ValueContext")
		}
		r.Cell().Text(hold)

		if field.Missing || field.Null {
			r.Cell().Text("")
		} else {
			value, err := field.Type.Id.ValueContext(e.Path, e.Aliases)
			if err != nil {
				return kerr.New("RWHEKAOPHQ", err, "ValueContext")
			}
			r.Cell().Text(value)
		}

	}
	table.Upgrade()
	e.AppendChild(table)
	return nil
}

func (e *ObjectEditor) AddChildTreeEntry(child Editor) bool {
	return child.Layout() == Page
}
