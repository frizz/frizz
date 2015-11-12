package editor

import (
	"honnef.co/go/js/dom"
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
	for _, node := range e.Map {
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

	for name, node := range e.Map {

		ed := node.Editor()

		r := table.Row()

		r.Cell().Text(name)

		origin, err := node.Origin.ValueContext(e.Path, e.Aliases)
		if err != nil {
			return kerr.New("ACQLJXWYQX", err, "ValueContext")
		}
		r.Cell().Text(origin)

		hold, err := node.Rule.HoldsDisplayType(e.Path, e.Aliases)
		if err != nil {
			return kerr.New("OYMARPFDGA", err, "ValueContext")
		}
		r.Cell().Text(hold)

		if node.Missing || node.Null {
			r.Cell().Text("")
		} else {
			value, err := node.Type.Id.ValueContext(e.Path, e.Aliases)
			if err != nil {
				return kerr.New("RWHEKAOPHQ", err, "ValueContext")
			}
			cell := r.Cell()
			a := mdl.Anchor().Text(value).Click(func(e dom.Event) {
				e.(*dom.MouseEvent).PreventDefault()
				ed.Select()
				ed.Focus()
			})
			cell.AppendChild(a)
		}
	}
	table.Upgrade()
	e.AppendChild(table)

	return nil
}

func (e *ObjectEditor) AddChildTreeEntry(child Editor) bool {
	return child.Layout() == Page
}
