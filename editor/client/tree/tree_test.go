package tree

import (
	"testing"

	"honnef.co/go/js/dom"

	"fmt"

	"kego.io/kerr/assert"
)

type dummy struct {
	id string
}

func (i *dummy) Initialise(div *dom.HTMLDivElement) {
	div.SetTextContent(i.id)
}

var _ = Item(&dummy{})

func TestTree(t *testing.T) {
	tree := New(nil, nil, nil, nil, nil, "", map[string]string{})
	a := tree.Root.Append(tree.Branch(&dummy{"a"}))
	b := tree.Root.Append(tree.Branch(&dummy{"b"}))
	test(t, "1a1b", tree)
	a.Append(tree.Branch(&dummy{"c"}))
	b.Append(tree.Branch(&dummy{"d"}))
	test(t, "1a1b", tree)
	a.Open()
	test(t, "1a2c1b", tree)
	b.Open()
	test(t, "1a2c1b2d", tree)
	// close root should have no effect
	tree.Root.Close()
	test(t, "1a2c1b2d", tree)
}

func TestTree2(t *testing.T) {
	tree := New(nil, nil, nil, nil, nil, "", map[string]string{})

	a := tree.Root.Append(tree.Branch(&dummy{"a"}))
	a.Open()

	b := a.Append(tree.Branch(&dummy{"b"}))
	b.Open()

	c := b.Append(tree.Branch(&dummy{"c"}))
	c.Open()

	d := c.Append(tree.Branch(&dummy{"d"}))
	d.Open()

	e := tree.Root.Append(tree.Branch(&dummy{"e"}))
	e.Open()

	test(t, "1a2b3c4d1e", tree)
	a.Close()
	test(t, "1a1e", tree)
	a.Toggle()
	test(t, "1a2b3c4d1e", tree)
	a.Toggle()
	test(t, "1a1e", tree)
}

func test(t *testing.T, expected string, tree *Tree) {
	f, b := testString(tree)
	assert.Equal(t, expected, f)
	assert.Equal(t, expected, b)
}

func testString(tree *Tree) (forwards string, backwards string) {
	n := tree.Root
	for n.next != nil {
		n = n.next
		forwards = fmt.Sprint(forwards, n.level, n.item.(*dummy).id)
	}
	backwards = fmt.Sprint(n.level, n.item.(*dummy).id)
	for n.prev != nil && !n.prev.root {
		n = n.prev
		backwards = fmt.Sprint(n.level, n.item.(*dummy).id, backwards)
	}
	return
}
