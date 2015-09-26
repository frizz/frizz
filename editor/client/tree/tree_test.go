package tree

import (
	"testing"

	"honnef.co/go/js/dom"

	"fmt"

	"kego.io/kerr/assert"
)

type item struct {
	id string
}

func (i *item) Initialise(div dom.Element) {
	div.SetTextContent(i.id)
}

var _ = Item(&item{})

func TestTree(t *testing.T) {
	r := New(nil)
	a := r.AppendItem(&item{"a"})
	b := r.AppendItem(&item{"b"})
	test(t, "1a1b", r)
	a.AppendItem(&item{"c"})
	b.AppendItem(&item{"d"})
	test(t, "1a1b", r)
	a.Open()
	test(t, "1a2c1b", r)
	b.Open()
	test(t, "1a2c1b2d", r)
	// close root should have no effect
	r.Close()
	test(t, "1a2c1b2d", r)
}

func TestTree2(t *testing.T) {
	r := New(nil)
	a := r.AppendItem(&item{"a"}).Open()
	a.AppendItem(&item{"b"}).Open().
		AppendItem(&item{"c"}).Open().
		AppendItem(&item{"d"}).Open()

	r.AppendItem(&item{"e"}).Open()

	test(t, "1a2b3c4d1e", r)
	a.Close()
	test(t, "1a1e", r)
	a.Toggle()
	test(t, "1a2b3c4d1e", r)
	a.Toggle()
	test(t, "1a1e", r)
}

func test(t *testing.T, expected string, root *Node) {
	f, b := testString(root)
	assert.Equal(t, expected, f)
	assert.Equal(t, expected, b)
}

func testString(root *Node) (forwards string, backwards string) {
	n := root
	for n.next != nil {
		n = n.next
		forwards = fmt.Sprint(forwards, n.level, n.item.(*item).id)
	}
	backwards = fmt.Sprint(n.level, n.item.(*item).id)
	for n.prev != nil && !n.prev.root {
		n = n.prev
		backwards = fmt.Sprint(n.level, n.item.(*item).id, backwards)
	}
	return
}
