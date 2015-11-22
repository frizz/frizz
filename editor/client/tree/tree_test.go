package tree

import (
	"testing"

	"fmt"

	"kego.io/kerr/assert"
)

type dummy struct {
	*Branch
	id string
}

func add(id string, parent BranchInterface) *dummy {
	d := &dummy{id: id}
	d.Branch = &Branch{tree: parent.GetTree(), parent: parent, self: d}
	parent.Append(d)
	return d
}

func TestTree(t *testing.T) {

	tr := &Tree{Path: "", Aliases: map[string]string{}}
	r := newRoot(tr, nil, nil)
	tr.Root = r

	a := add("a", r)
	b := add("b", r)
	test(t, "1a1b", tr)

	add("c", a)
	add("d", b)

	test(t, "1a1b", tr)

	a.Open()
	test(t, "1a2c1b", tr)

	b.Open()
	test(t, "1a2c1b2d", tr)

	// close root should have no effect
	tr.Root.Close()
	test(t, "1a2c1b2d", tr)
}

func TestTree2(t *testing.T) {
	tr := &Tree{Path: "", Aliases: map[string]string{}}
	r := newRoot(tr, nil, nil)
	tr.Root = r

	a := add("a", r)
	a.Open()

	test(t, "1a", tr)

	b := add("b", a)
	b.Open()

	test(t, "1a2b", tr)

	c := add("c", b)
	c.Open()

	test(t, "1a2b3c", tr)

	d := add("d", c)
	d.Open()

	test(t, "1a2b3c4d", tr)

	e := add("e", r)
	e.Open()

	test(t, "1a2b3c4d1e", tr)
	a.Close()
	test(t, "1a1e", tr)
	a.Toggle()
	test(t, "1a2b1e", tr)
	a.Toggle()
	test(t, "1a1e", tr)
}

func test(t *testing.T, expected string, tree *Tree) {
	f, b := testString(tree)
	assert.Equal(t, expected, f)
	assert.Equal(t, expected, b)
}

func testString(tree *Tree) (forwards string, backwards string) {
	n := tree.Root
	for n.NextVisible(true) != nil {
		n = n.NextVisible(true)
		forwards = fmt.Sprint(forwards, n.GetLevel(), n.(*dummy).id)
	}
	backwards = fmt.Sprint(n.GetLevel(), n.(*dummy).id)
	for n.PrevVisible() != nil && !n.PrevVisible().IsRoot() {
		n = n.PrevVisible()
		backwards = fmt.Sprint(n.GetLevel(), n.(*dummy).id, backwards)
	}
	return
}
