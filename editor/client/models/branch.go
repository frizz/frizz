package models

import "kego.io/system/node"

type BranchModel struct {
	Node *node.Node
	Root bool
	Open bool
}

func (b *BranchModel) CanOpen() bool {
	if b.Root {
		return true
	}
	if b.IsAsyncAndNotLoaded() {
		return true
	}
	if len(b.Node.Array) == 0 && len(b.Node.Map) == 0 {
		return false
	}
	return true
}

func (b *BranchModel) IsAsyncAndNotLoaded() bool {
	return false
}

func (b *BranchModel) Icon() string {
	if !b.CanOpen() {
		return "empty"
	}
	if b.IsAsyncAndNotLoaded() {
		return "unknown"
	}
	if b.Open {
		return "minus"
	}
	return "plus"
}
