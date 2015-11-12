package node

// Flatten returns a flat array of this node and all descendants
func (n *Node) Flatten(removeMissing bool) []*Node {
	if removeMissing && n.Missing {
		return []*Node{}
	}
	out := []*Node{n}
	for _, child := range n.Array {
		out = append(out, child.Flatten(removeMissing)...)
	}
	for _, child := range SortNodeMap(n.Map) {
		out = append(out, child.Node.Flatten(removeMissing)...)
	}
	return out
}
