package tree

import (
	"testing"

	"fmt"

	"kego.io/context/envctx"
	"kego.io/kerr/assert"
)

type dummy struct {
	*Branch
	id string
}

func add(id string, parent BranchInterface) *dummy {
	d := &dummy{id: id}
	d.Branch = &Branch{tree: parent.branch().tree, parent: parent.branch(), self: d}
	parent.Append(d)
	return d
}

func (b *Branch) getLevel() int {
	current := b
	i := 0
	for current.parent != nil {
		current = current.parent
		i++
	}
	return i
}

func TestTree(t *testing.T) {

	tr := &Tree{ctx: envctx.Empty}
	r := newRoot(tr, nil, nil)
	tr.Root = r

	a := add("a", r)
	b := add("b", r)
	test(t, "1a1b", tr)

	add("c", a)
	add("d", b)

	test(t, "1a1b", tr)

	a.open()
	test(t, "1a2c1b", tr)

	b.open()
	test(t, "1a2c1b2d", tr)

	// close root should have no effect
	tr.Root.close()
	test(t, "1a2c1b2d**", tr)
}

func TestTree2(t *testing.T) {
	tr := &Tree{ctx: envctx.Empty}
	r := newRoot(tr, nil, nil)
	tr.Root = r

	a := add("a", r)
	a.open()

	test(t, "1a", tr)

	b := add("b", a)
	b.open()

	test(t, "1a2b", tr)

	c := add("c", b)
	c.open()

	test(t, "1a2b3c", tr)

	d := add("d", c)
	d.open()

	test(t, "1a2b3c4d", tr)

	e := add("e", r)
	e.open()

	test(t, "1a2b3c4d1e", tr)
	a.close()
	test(t, "1a1e", tr)
	a.toggle()
	test(t, "1a2b1e", tr)
	a.toggle()
	test(t, "1a1e", tr)
}

func test(t *testing.T, expected string, tree *Tree) {
	f, b := testString(tree)
	assert.Equal(t, expected, f)
	assert.Equal(t, expected, b)
}

func testString(tree *Tree) (forwards string, backwards string) {
	n := tree.Root.Branch
	for n.nextVisible(true) != nil {
		n = n.nextVisible(true)
		forwards = fmt.Sprint(forwards, n.getLevel(), n.self.(*dummy).id)
	}
	backwards = fmt.Sprint(n.getLevel(), n.self.(*dummy).id)
	for n.prevVisible() != nil && n.prevVisible().parent != nil {
		n = n.prevVisible()
		backwards = fmt.Sprint(n.getLevel(), n.self.(*dummy).id, backwards)
	}
	return
}
