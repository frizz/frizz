package system

import "sort"

func SortNodeMap(nodes map[string]*Node) []NamedNode {
	sortable := SortableNamedNodes{}
	for name, node := range nodes {
		sortable = append(sortable, NamedNode{name, node})
	}
	sort.Sort(sortable)
	return []NamedNode(sortable)
}

type NamedNode struct {
	Name string
	Node *Node
}

type SortableNamedNodes []NamedNode

func (s SortableNamedNodes) Len() int {
	return len(s)
}
func (s SortableNamedNodes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s SortableNamedNodes) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

var _ sort.Interface = (*SortableNamedNodes)(nil)
