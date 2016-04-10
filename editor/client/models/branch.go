package models

import "kego.io/system/node"

type BranchModel struct {
	Children []*BranchModel
	Root     bool
	Open     bool
	Contents BranchContentsInterface
	Parent   *BranchModel
	index    int
}

func (b *BranchModel) CanOpen() bool {
	if b.Root {
		return true
	}
	if async, ok := b.Contents.(AsyncInterface); ok && !async.Loaded() {
		return true
	}
	if len(b.Children) == 0 {
		return false
	}
	return true
}

func (b *BranchModel) Icon() string {
	if !b.CanOpen() {
		return "empty"
	}
	if async, ok := b.Contents.(AsyncInterface); ok && !async.Loaded() {
		return "unknown"
	}
	if b.Open {
		return "minus"
	}
	return "plus"
}

func NewNodeBranch(n *node.Node, name string) *BranchModel {
	b := &BranchModel{
		Contents: &NodeContents{
			Node: n,
			Name: name,
		},
	}
	AppendNodeChildren(b, n)
	return b
}

func AppendNodeChildren(b *BranchModel, n *node.Node) {
	for _, c := range n.Array {
		b.Append(NewNodeBranch(c.GetNode(), ""))
	}
	for _, c := range n.Map {
		b.Append(NewNodeBranch(c.GetNode(), ""))
	}
}

func (b *BranchModel) Append(children ...*BranchModel) *BranchModel {
	for i, child := range children {
		child.index = len(b.Children) + i
		child.Parent = b
	}
	b.Children = append(b.Children, children...)
	return b
}
