package editor

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/mdl"
	"kego.io/kerr"
	"kego.io/system/node"
)

type NodeObjectEditor struct {
	*node.Node
	*Common
}

var _ Editor = (*NodeObjectEditor)(nil)

func (e *NodeObjectEditor) Initialize(panel *dom.HTMLDivElement, path string, aliases map[string]string) error {

	e.Panel = panel
	e.Path = path
	e.Aliases = aliases

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

	e.Initialized = true
	return nil
}
