package models

import "kego.io/system/node"

type BranchModel struct {
	Children []*BranchModel
	Root     bool
	Open     bool
	Contents BranchContentsInterface
}

func (b *BranchModel) CanOpen() bool {
	if b.Root {
		return true
	}
	if !b.Contents.Loaded() {
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
	if !b.Contents.Loaded() {
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
	for _, c := range n.Array {
		b.Children = append(b.Children, NewNodeBranch(c.GetNode(), ""))
	}
	for _, c := range n.Map {
		b.Children = append(b.Children, NewNodeBranch(c.GetNode(), ""))
	}
	return b
}
