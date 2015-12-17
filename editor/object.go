package editor

import (
	"honnef.co/go/js/dom"
	"kego.io/editor/client/mdl"
	"kego.io/kerr"
	"kego.io/system"
)

type ObjectEditor struct {
	*Node
	*Editor
	editors *dom.HTMLDivElement
}

var _ EditorInterface = (*ObjectEditor)(nil)

func NewObjectEditor(n *Node) *ObjectEditor {
	return &ObjectEditor{Node: n, Editor: &Editor{}}
}

func (e *ObjectEditor) Layout() Layout {
	return Page
}

func (e *ObjectEditor) Initialize(holder BranchInterface, layout Layout, fail chan error, path string, aliases map[string]string) error {

	e.Editor.Initialize(holder, layout, fail, path, aliases)

	e.editors = dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	e.AppendChild(e.editors)

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
		e.initializeBlockEditor(ed)
	}
}

func (e *ObjectEditor) initializeBlockEditor(ed EditorInterface) {
	e.Editors = append(e.Editors, ed)
	ed.Initialize(e.branch, Block, e.fail, e.Path, e.Aliases)
	e.branch.ListenForEditorChanges(ed.Listen().Ch)
	e.editors.AppendChild(ed)
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

	if !node.Rule.Parent.Interface && !node.Rule.Struct.Interface {
		// if the rule only permits a single concrete type, we initialise immediately with that type
		e.InitialiseChildWithConcreteType(node, node.Rule.Parent)
		return
	}

	// If the rule is an interface, so we must pop up UX to choose the concrete type.
	hashedTypes := system.GetAllTypesThatImplementInterface(node.Rule.Parent, node.Rule.Struct.Interface)

	if len(hashedTypes) == 1 {
		// if only one type is compatible, don't show the popup, just add it.
		e.InitialiseChildWithConcreteType(node, hashedTypes[0].Object.(*system.Type))
		return
	}

	card := mdl.Card("Choose a type", "Add")
	options := map[string]string{}

	for _, h := range hashedTypes {
		displayName, err := h.Object.GetObject().Id.ValueContext(e.Path, e.Aliases)
		if err != nil {
			// we shouldn't be able to get here
			e.fail <- kerr.New("IPLHSXDWQK", err, "ValueContext")
		}
		options[h.Object.GetObject().Id.String()] = displayName
	}

	dropdown := mdl.Select(options)
	card.Content.AppendChild(dropdown)

	go func() {
		result := <-card.Result
		if result {
			r, err := system.NewReferenceFromString(dropdown.Value, e.Path, e.Aliases)
			if err != nil {
				e.fail <- kerr.New("KHJGQXORPD", err, "NewReferenceFromString")
			}
			h, ok := system.GetGlobal(r.Package, r.Name)
			if !ok {
				e.fail <- kerr.New("WEADSXTPYC", nil, "Global %s not found", r.Value())
			}
			t, ok := h.Object.(*system.Type)
			if !ok {
				e.fail <- kerr.New("FPOYMFLKVE", nil, "Global not a *Type")
			}
			e.InitialiseChildWithConcreteType(node, t)
		}
		card.Hide()
	}()

}

func (e *ObjectEditor) InitialiseChildWithConcreteType(node *Node, t *system.Type) {

	if err := node.InitialiseWithConcreteType(t, e.Path, e.Aliases); err != nil {
		e.fail <- kerr.New("KHDLYHXOWS", err, "InitialiseWithConcreteType")
	}

	node.UpdateFromInnerNode()

	// Updates the object summary table with the new node info
	node.changes.Send(node)

	ed := node.Editor()

	if ed == nil {
		return
	}

	if ed.Layout() == Page {
		e.branch.AddFieldNode(node)
	} else {
		e.initializeBlockEditor(ed)
	}

	ed.Focus()

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
