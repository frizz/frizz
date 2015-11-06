package branch

import (
	"testing"

	"fmt"

	"kego.io/editor/client/tree"
	"kego.io/kerr/assert"
)

type dummy struct {
	*Common
	id string
}

func (d *dummy) Initialise(parent tree.Branch) {
	d.Common = &Common{this: d, tree: parent.Tree()}
	//d.Common.Initialise(parent)
	//d.Common.SetLabel(d.id)
}

func add(id string, parent tree.Branch) *dummy {
	d := &dummy{id: id}
	d.Initialise(parent)
	parent.Append(d)
	return d
}

func TestTree(t *testing.T) {

	tr := &tree.Tree{Path: "", Aliases: map[string]string{}}
	r := &Root{}
	tr.Root = r
	r.Common = &Common{this: r, tree: tr, root: true, open: true}

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
	tr := &tree.Tree{Path: "", Aliases: map[string]string{}}
	r := &Root{}
	tr.Root = r
	r.Common = &Common{this: r, tree: tr, root: true, open: true}

	a := add("a", r)
	a.Open()

	b := add("b", a)
	b.Open()

	c := add("c", b)
	c.Open()

	d := add("d", c)
	d.Open()

	e := add("e", r)
	e.Open()

	test(t, "1a2b3c4d1e", tr)
	a.Close()
	test(t, "1a1e", tr)
	a.toggle()
	test(t, "1a2b3c4d1e", tr)
	a.toggle()
	test(t, "1a1e", tr)
}

func test(t *testing.T, expected string, tree *tree.Tree) {
	f, b := testString(tree)
	assert.Equal(t, expected, f)
	assert.Equal(t, expected, b)
}

func testString(tree *tree.Tree) (forwards string, backwards string) {
	n := tree.Root
	for n.NextVisible(true) != nil {
		n = n.NextVisible(true)
		forwards = fmt.Sprint(forwards, n.Level(), n.(*dummy).id)
	}
	backwards = fmt.Sprint(n.Level(), n.(*dummy).id)
	for n.PrevVisible() != nil && !n.PrevVisible().Root() {
		n = n.PrevVisible()
		backwards = fmt.Sprint(n.Level(), n.(*dummy).id, backwards)
	}
	return
}
