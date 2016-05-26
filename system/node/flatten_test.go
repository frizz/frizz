package node

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"kego.io/json"
)

func TestFlatten(t *testing.T) {
	n := NewNode()
	n.Missing = true
	a := n.Flatten(true)
	assert.Equal(t, 0, len(a))

	n = getArrayNode(
		getStringNode("c"),
		getStringNode("a"),
		getStringNode("b"),
	)
	assert.Equal(t, "c,a,b", flatten(n))

	n = getMapNode(
		"123",
		getStringNode("a"),
		getStringNode("b"),
		getStringNode("c"),
	)
	assert.Equal(t, "a,b,c", flatten(n))

	// map nodes should be ordered by key
	n = getMapNode(
		"312",
		getStringNode("a"),
		getStringNode("b"),
		getStringNode("c"),
	)
	assert.Equal(t, "b,c,a", flatten(n))

	n = getArrayNode(
		getStringNode("a"),
		getStringNode("b"),
		getMapNode(
			"12",
			getArrayNode(
				getStringNode("f"),
				getStringNode("g"),
			),
			getMapNode(
				"21",
				getStringNode("h"),
				getStringNode("i"),
			),
		),
		getStringNode("c"),
	)
	assert.Equal(t, "a,b,f,g,i,h,c", flatten(n))
}

func getArrayNode(a ...*Node) *Node {
	n := NewNode()
	n.JsonType = json.J_ARRAY
	for _, c := range a {
		n.Array = append(n.Array, c)
	}
	return n
}

func getMapNode(names string, a ...*Node) *Node {
	n := NewNode()
	n.JsonType = json.J_MAP
	n.Map = map[string]*Node{}
	for i, c := range a {
		n.Map[string(names[i])] = c
	}
	return n
}

func getStringNode(val string) *Node {
	n := NewNode()
	n.JsonType = json.J_STRING
	n.ValueString = val
	return n
}

func flatten(n *Node) string {
	a := n.Flatten(false)
	out := ""
	for _, c := range a {
		if c.JsonType == json.J_STRING {
			if out != "" {
				out += ","
			}
			out += c.ValueString
		}
	}
	return out
}
