package editor

import (
	"golang.org/x/net/context"
	"honnef.co/go/js/dom"
	"kego.io/editor/client/mdl"
	"kego.io/kerr"
	"kego.io/system"
)

type MapEditor struct {
	*Node
	*Editor
}

var _ EditorInterface = (*MapEditor)(nil)

func NewMapEditor(n *Node) *MapEditor {
	return &MapEditor{Node: n, Editor: &Editor{}}
}

func (e *MapEditor) Layout() Layout {
	return Page
}

func (e *MapEditor) Initialize(ctx context.Context, holder BranchInterface, layout Layout, fail chan error) error {

	e.Editor.Initialize(ctx, holder, layout, fail)

	table := mdl.Table()

	items, err := system.WrapRule(e.Rule.Interface.(system.CollectionRule).GetItemsRule())
	if err != nil {
		return kerr.New("GQROTGVBXS", err, "NewRuleHolder")
	}
	hold, err := items.HoldsDisplayType(ctx)
	if err != nil {
		return kerr.New("XDKOSFJVQV", err, "ValueContext")
	}

	table.Head("name", "holds", "value")

	for name, item := range e.Map {

		r := table.Row()

		if !item.Null {
			ed := item.Editor()
			r.Click(func(e dom.Event) {
				e.(*dom.MouseEvent).PreventDefault()
				ed.Select()
				ed.Focus()
			})
		}

		r.Cell().Text(name)
		r.Cell().Text(hold)

		if item.Null {
			r.Cell().Text("")
		} else {
			val, err := item.Type.Id.ValueContext(ctx)
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

func (e *MapEditor) Value() interface{} {
	return e.Node.Value
}
