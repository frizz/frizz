package editor

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
	"kego.io/kerr"
)

type NodeObjectEditor struct {
	*Node
	*Common
	editors []Editor
}

func (e *NodeObjectEditor) Layout() Layout {
	return Page
}

var _ Editor = (*NodeObjectEditor)(nil)

func (e *NodeObjectEditor) Initialize(panel *dom.HTMLDivElement, holder Holder, layout Layout, path string, aliases map[string]string) error {

	e.Common.Initialize(panel, holder, layout, path, aliases)

	e.initializeBlockEditors()

	if err := e.initializeTable(); err != nil {
		return kerr.New("KAVTMDDFYW", err, "initializeTable")
	}

	return nil
}

func (e *NodeObjectEditor) initializeBlockEditors() {
	for _, field := range e.Fields {
		node := Node{field}
		ed := node.Editor()
		if ed == nil || ed.Layout() == Page {
			continue
		}
		e.editors = append(e.editors, ed)
		ed.Initialize(e.Panel, e.holder, Block, e.Path, e.Aliases)
	}
}

func (e *NodeObjectEditor) initializeTable() error {

	table := mdl.Table()

	names := table.Column("name")
	origins := table.Column("origin")
	holds := table.Column("holds")
	values := table.Column("value")

	for name, field := range e.Fields {

		names.Cell(name)

		origin, err := field.Origin.ValueContext(e.Path, e.Aliases)
		if err != nil {
			return kerr.New("ACQLJXWYQX", err, "ValueContext")
		}
		origins.Cell(origin)

		hold, err := field.Rule.HoldsDisplayType(e.Path, e.Aliases)
		if err != nil {
			return kerr.New("OYMARPFDGA", err, "ValueContext")
		}
		holds.Cell(hold)

		if field.Missing || field.Null {
			values.Cell("")
		} else {
			value, err := field.Type.Id.ValueContext(e.Path, e.Aliases)
			if err != nil {
				return kerr.New("RWHEKAOPHQ", err, "ValueContext")
			}
			values.Cell(value)
		}

	}
	e.Panel.AppendChild(table.Build())
	return nil
}

func (e *NodeObjectEditor) AddChildTreeEntry(child Editor) bool {
	return child.Layout() == Page
}
