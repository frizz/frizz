package editor

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/mdl"
	"kego.io/kerr"
)

type ObjectEditor struct {
	*Node
	*Editor
}

var _ EditorInterface = (*ObjectEditor)(nil)

func NewObjectEditor(n *Node) *ObjectEditor {
	return &ObjectEditor{Node: n, Editor: &Editor{}}
}

func (e *ObjectEditor) Layout() Layout {
	return Page
}

func (e *ObjectEditor) Initialize(holder BranchInterface, layout Layout, path string, aliases map[string]string) error {

	e.Editor.Initialize(holder, layout, path, aliases)

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
		ed.Initialize(e.branch, Block, e.Path, e.Aliases)
		e.branch.ListenForEditorChanges(ed.Listen().Ch)
		e.AppendChild(ed)
	}
}

func (e *ObjectEditor) initializeTable() error {

	table, err := NewObjectSummary(e.Node, e.Path, e.Aliases)
	if err != nil {
		return kerr.New("VYJPBBMEBS", err, "NewObjectSummary")
	}
	e.AppendChild(table)

	return nil
}

func (e *ObjectEditor) AddChildTreeEntry(child EditorInterface) bool {
	return child.Layout() == Page
}

func (e *ObjectEditor) Value() interface{} {
	return e.Node.Value
}

func (e *ObjectEditor) AddField(node *Node) {
	node.Missing = false
	node.Null = false
	// TODO: update e.Node.Value and other Node things
	// TODO: add block editor, or somehow get the branch to add a child :/
	node.changes.Send(node)
}

type objectSummary struct {
	*summary
}

type objectSummaryRow struct {
	*summaryRow
	name    *summaryCell
	origin  *summaryCell
	holds   *summaryCell
	value   *summaryCell
	options *summaryCell
}

func NewObjectSummary(node *Node, path string, aliases map[string]string) (*objectSummary, error) {
	s := &objectSummary{summary: &summary{TableStruct: mdl.Table(), path: path, aliases: aliases}}
	s.Head("name", "origin", "holds", "value", "options")
	for _, child := range node.Map {
		r := s.newRow(child)
		if err := r.update(); err != nil {
			return nil, kerr.New("UNVCXHWGRF", err, "update")
		}
	}
	s.Upgrade()
	return s, nil
}

func (s *objectSummary) newRow(node *Node) *objectSummaryRow {
	r := &objectSummaryRow{summaryRow: &summaryRow{TableRowStruct: s.Row(), node: node, table: s.summary}}
	r.name = &summaryCell{TableCellStruct: r.Cell()}
	r.origin = &summaryCell{TableCellStruct: r.Cell()}
	r.holds = &summaryCell{TableCellStruct: r.Cell()}
	r.value = &summaryCell{TableCellStruct: r.Cell()}
	r.options = &summaryCell{TableCellStruct: r.Cell()}

	changes := node.changes.Listen().Ch
	go func() {
		for e := range changes {
			_ = e.(*Node)
			r.update()
		}
	}()

	return r
}

func (r *objectSummaryRow) update() error {
	node := r.node

	if !node.Missing && !node.Null {
		ed := node.Editor()
		r.Click(func(e dom.Event) {
			e.(*dom.MouseEvent).PreventDefault()
			ed.Select()
			ed.Focus()
		})
	}
	r.name.Text(node.Key)

	origin, err := node.Origin.ValueContext(r.table.path, r.table.aliases)
	if err != nil {
		return kerr.New("ACQLJXWYQX", err, "ValueContext")
	}
	r.origin.Text(origin)

	hold, err := node.Rule.HoldsDisplayType(r.table.path, r.table.aliases)
	if err != nil {
		return kerr.New("OYMARPFDGA", err, "ValueContext")
	}
	r.holds.Text(hold)

	if node.Missing || node.Null {
		r.value.Text("")
	} else {
		value, err := node.Type.Id.ValueContext(r.table.path, r.table.aliases)
		if err != nil {
			return kerr.New("LDAMPELUCM", err, "ValueContext")
		}
		r.value.Text(value)
	}

	// TODO: optimize this with state
	for _, child := range r.options.ChildNodes() {
		r.options.RemoveChild(child)
	}
	if node.Missing || node.Null {
		a := mdl.Anchor().Text("add").Click(func(e dom.Event) {
			e.(*dom.MouseEvent).PreventDefault()
			node.Parent.editor.(*ObjectEditor).AddField(node)

		})
		r.options.AppendChild(a)
	}
	return nil
}
