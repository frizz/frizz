package items

import (
	"testing"

	"fmt"

	"kego.io/editor/client/tree"
	"kego.io/kerr/assert"
)

type dummy struct {
	branch *tree.Branch
	id     string
}

func add(id string, parent *tree.Branch) *dummy {
	d := &dummy{id: id}
	d.branch = &tree.Branch{Tree: parent.Tree, Parent: parent, Item: d}
	parent.Append(d.branch)
	return d
}

func TestTree(t *testing.T) {

	tr := &tree.Tree{Path: "", Aliases: map[string]string{}}
	r := &Root{branch: tree.NewRootBranch(tr, nil, nil)}
	tr.Root = r.branch

	a := add("a", r.branch)
	b := add("b", r.branch)
	test(t, "1a1b", tr)

	add("c", a.branch)
	add("d", b.branch)

	test(t, "1a1b", tr)

	a.branch.Open()
	test(t, "1a2c1b", tr)

	b.branch.Open()
	test(t, "1a2c1b2d", tr)

	// close root should have no effect
	tr.Root.Close()
	test(t, "1a2c1b2d", tr)
}

func TestTree2(t *testing.T) {
	tr := &tree.Tree{Path: "", Aliases: map[string]string{}}
	r := &Root{branch: tree.NewRootBranch(tr, nil, nil)}
	tr.Root = r.branch

	a := add("a", r.branch)
	a.branch.Open()

	b := add("b", a.branch)
	b.branch.Open()

	c := add("c", b.branch)
	c.branch.Open()

	d := add("d", c.branch)
	d.branch.Open()

	e := add("e", r.branch)
	e.branch.Open()

	test(t, "1a2b3c4d1e", tr)
	a.branch.Close()
	test(t, "1a1e", tr)
	a.branch.Toggle()
	test(t, "1a2b3c4d1e", tr)
	a.branch.Toggle()
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
		forwards = fmt.Sprint(forwards, n.Level(), n.Item.(*dummy).id)
	}
	backwards = fmt.Sprint(n.Level(), n.Item.(*dummy).id)
	for n.PrevVisible() != nil && !n.PrevVisible().IsRoot() {
		n = n.PrevVisible()
		backwards = fmt.Sprint(n.Level(), n.Item.(*dummy).id, backwards)
	}
	return
}
