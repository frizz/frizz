package items

import (
	"fmt"

	"honnef.co/go/js/dom"

	"kego.io/editor/client/tree"
	"kego.io/system/node"
)

// Entry items are nodes. Each branch inside a source branch are entry.
type entry struct {
	*item
	name  string
	index int
	node  *node.Node
	label *dom.HTMLDivElement
}

var _ tree.Item = (*entry)(nil)
var _ tree.Noder = (*entry)(nil)

func (e *entry) Node() *node.Node {
	return e.node
}

func (e *entry) Initialise(div *dom.HTMLDivElement) {
	label := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)

	name := ""
	if e.index > -1 {
		name = fmt.Sprint("[", e.index, "]")
	} else {
		name = e.name
	}

	label.SetTextContent(name)
	e.label = label
	div.AppendChild(label)
}

func shortenString(in string) string {
	var runes = 0
	for i, _ := range in {
		runes++
		if runes > 25 {
			return fmt.Sprint(in[:i], "...")
		}
	}
	return in
}

func addEntry(name string, index int, node *node.Node, parentBranch *tree.Branch) *entry {

	if node.Parent != nil && node.Parent.Type.Native.Value() == "object" {
		// Don't display "type" or "id" nodes if the parent is an object (maps are ok!)
		if name == "type" || name == "id" {
			return nil
		}
	}

	newEntry := &entry{item: &item{tree: parentBranch.Tree}, name: name, index: index, node: node}
	newBranch := parentBranch.Tree.Branch(newEntry)
	newEntry.branch = newBranch

	parentBranch.Append(newBranch)

	addNodeChildren(node, newBranch)

	newBranch.Close()

	return newEntry
}

func addNodeChildren(n *node.Node, b *tree.Branch) {
	if n == nil {
		return
	}
	switch n.Type.Native.Value() {
	case "array":
		for i, childNode := range n.Array {
			addEntry("", i, childNode, b)
		}
	case "map":
		for name, childNode := range n.Map {
			addEntry(name, -1, childNode, b)
		}
	case "object":
		for name, childNode := range n.Fields {
			if childNode.Missing {
				continue
			}
			addEntry(name, -1, childNode, b)
		}
	}
}
