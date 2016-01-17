package editor

import (
	"golang.org/x/net/context"
	"honnef.co/go/js/dom"
	"kego.io/editor/client/mdl"
	"kego.io/kerr"
	"kego.io/system"
)

type ObjectEditor struct {
	*Node
	*Editor
	editors	*dom.HTMLDivElement
}

var _ EditorInterface = (*ObjectEditor)(nil)

func NewObjectEditor(n *Node) *ObjectEditor {
	return &ObjectEditor{Node: n, Editor: &Editor{}}
}

func (e *ObjectEditor) Layout() Layout {
	return Page
}

func (e *ObjectEditor) Initialize(ctx context.Context, holder BranchInterface, layout Layout, fail chan error) error {

	e.Editor.Initialize(ctx, holder, layout, fail)

	e.editors = dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	e.AppendChild(e.editors)

	e.initializeBlockEditors()

	if err := e.initializeTable(); err != nil {
		return kerr.Wrap("KAVTMDDFYW", err)
	}

	return nil
}

func (e *ObjectEditor) initializeBlockEditors() {
	for _, v := range e.Map {
		node := v.(*Node)
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
	ed.Initialize(e.ctx, e.branch, Block, e.fail)
	e.branch.ListenForEditorChanges(ed.Listen().Ch)
	e.editors.AppendChild(ed)
}

func (e *ObjectEditor) initializeTable() error {

	table, err := NewObjectSummary(e.ctx, e.Node)
	if err != nil {
		return kerr.Wrap("VYJPBBMEBS", err)
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
	types := system.GetAllTypesThatImplementInterface(e.ctx, node.Rule.Parent)

	if len(types) == 1 {
		// if only one type is compatible, don't show the popup, just add it.
		e.InitialiseChildWithConcreteType(node, types[0])
		return
	}

	card := mdl.Card("Choose a type", "Add")
	options := map[string]string{}

	for _, t := range types {
		displayName, err := t.Id.ValueContext(e.ctx)
		if err != nil {
			// we shouldn't be able to get here
			e.fail <- kerr.Wrap("IPLHSXDWQK", err)
		}
		options[t.Id.String()] = displayName
	}

	dropdown := mdl.Select(options)
	card.Content.AppendChild(dropdown)

	go func() {
		result := <-card.Result
		if result {
			r, err := system.NewReferenceFromString(e.ctx, dropdown.Value)
			if err != nil {
				e.fail <- kerr.Wrap("KHJGQXORPD", err)
			}
			t, ok := system.GetTypeFromCache(e.ctx, r.Package, r.Name)
			if !ok {
				e.fail <- kerr.New("WEADSXTPYC", "Type %s not found in cache", r.Value())
			}
			e.InitialiseChildWithConcreteType(node, t)
		}
		card.Hide()
	}()

}

func (e *ObjectEditor) InitialiseChildWithConcreteType(node *Node, t *system.Type) {

	if err := node.InitialiseWithConcreteType(e.ctx, t); err != nil {
		e.fail <- kerr.Wrap("KHDLYHXOWS", err)
	}

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

}

type objectSummary struct {
	*summary
}

type objectSummaryRow struct {
	*summaryRow
	name	*summaryCell
	origin	*summaryCell
	holds	*summaryCell
	value	*summaryCell
	options	*summaryCell
}

func NewObjectSummary(ctx context.Context, node *Node) (*objectSummary, error) {
	s := &objectSummary{summary: &summary{TableStruct: mdl.Table(), ctx: ctx}}
	s.Head("name", "origin", "holds", "value", "options")
	for _, v := range node.Map {
		child := v.(*Node)
		r := s.newRow(child)
		if err := r.update(); err != nil {
			return nil, kerr.Wrap("UNVCXHWGRF", err)
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

	origin, err := node.Origin.ValueContext(r.table.ctx)
	if err != nil {
		return kerr.Wrap("ACQLJXWYQX", err)
	}
	r.origin.Text(origin)

	hold, err := node.Rule.HoldsDisplayType(r.table.ctx)
	if err != nil {
		return kerr.Wrap("OYMARPFDGA", err)
	}
	r.holds.Text(hold)

	if node.Missing || node.Null {
		r.value.Text("")
	} else {
		value, err := node.Type.Id.ValueContext(r.table.ctx)
		if err != nil {
			return kerr.Wrap("LDAMPELUCM", err)
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
			node.Parent.(*Node).editor.(*ObjectEditor).AddField(node)
		})
		r.options.AppendChild(a)
	}
	return nil
}
