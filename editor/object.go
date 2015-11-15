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
		if node.Missing || node.Null {
			continue
		}
		ed := node.Editor()
		if ed == nil || ed.Layout() == Page {
			continue
		}
		e.Editors = append(e.Editors, ed)
		ed.Initialize(e.holder, Block, e.Path, e.Aliases)
		e.holder.Listen(ed.Listen().Ch)
		e.AppendChild(ed)
	}
}

func (e *ObjectEditor) initializeTable() error {

	table := mdl.Table()
	table.Head("name", "origin", "holds", "value", "options")

	for name, node := range e.Map {

		r := table.Row()

		if !node.Missing && !node.Null {
			ed := node.Editor()
			r.Click(func(e dom.Event) {
				e.(*dom.MouseEvent).PreventDefault()
				ed.Select()
				ed.Focus()
			})
		}
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
				return kerr.New("LDAMPELUCM", err, "ValueContext")
			}
			r.Cell().Text(value)
		}

		cell := r.Cell()
		if node.Missing || node.Null {
			a := mdl.Anchor().Text("add").Click(func(e dom.Event) {

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

func (e *ObjectEditor) Value() interface{} {
	return e.Node.Value
}
