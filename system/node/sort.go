package node

import (
	"sort"

	"github.com/davelondon/sorter"
)

func SortNodeMap(nodes map[string]*Node) []NamedNode {
	var s []NamedNode
	for name, node := range nodes {
		s = append(s, NamedNode{Name: name, Node: node})
	}
	sort.Sort(sorter.New(
		len(s),
		func(i, j int) { s[i], s[j] = s[j], s[i] },
		func(i, j int) bool { return s[i].Name < s[j].Name },
	))
	return s
}

type NamedNode struct {
	Name string
	Node *Node
}
