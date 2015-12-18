package editor

import (
	"fmt"

	"honnef.co/go/js/dom"

	"golang.org/x/net/context"
	"kego.io/editor/client/mdl"
	"kego.io/kerr"
	"kego.io/system"
)

type ArrayEditor struct {
	*Node
	*Editor
}

var _ EditorInterface = (*ArrayEditor)(nil)

func NewArrayEditor(n *Node) *ArrayEditor {
	return &ArrayEditor{Node: n, Editor: &Editor{}}
}

func (e *ArrayEditor) Layout() Layout {
	return Page
}

func (e *ArrayEditor) Initialize(ctx context.Context, holder BranchInterface, layout Layout, fail chan error) error {

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

	table.Head("index", "holds", "value")

	for i, item := range e.Array {

		r := table.Row()

		if !item.Null {
			ed := item.Editor()
			r.Click(func(e dom.Event) {
				e.(*dom.MouseEvent).PreventDefault()
				ed.Select()
				ed.Focus()
			})
		}

		r.Cell().Text(fmt.Sprintf("%d", i))
		r.Cell().Text(hold)

		if item.Null {
			r.Cell().Text("")
		} else {
			val, err := item.Type.Id.ValueContext(ctx)
			if err != nil {
				return kerr.New("PYNPNMICPQ", err, "ValueContext")
			}
			r.Cell().Text(val)
		}

	}
	table.Upgrade()
	e.AppendChild(table)

	return nil
}

func (e *ArrayEditor) Value() interface{} {
	return e.Node.Value
}
