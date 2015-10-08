package items // import "kego.io/editor/client/items"

import "kego.io/editor/client/tree"

type item struct {
	tree   *tree.Tree
	branch *tree.Branch
}
