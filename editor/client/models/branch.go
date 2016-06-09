package models

import "golang.org/x/net/context"

type BranchModel struct {
	ctx      context.Context
	Children []*BranchModel
	Root     bool
	Open     bool
	Contents BranchContentsInterface
	Parent   *BranchModel
	index    int
	LastOp   BranchOps
}

func NewBranchModel(ctx context.Context, contents BranchContentsInterface) *BranchModel {
	return &BranchModel{ctx: ctx, Contents: contents}
}

type BranchOps string

const BranchOpKeyboard BranchOps = "BranchOpKeyboard"
const BranchOpClickLabel BranchOps = "BranchOpClickLabel"
const BranchOpClickBreadcrumb BranchOps = "BranchOpClickBreadcrumb"
const BranchOpClickToggle BranchOps = "BranchOpClickToggle"

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

func (b *BranchModel) RecursiveClose() {
	if b.Open {
		b.Open = false
		for _, c := range b.Children {
			c.RecursiveClose()
		}
	}
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

func (b *BranchModel) Append(children ...*BranchModel) *BranchModel {
	for i, child := range children {
		child.index = len(b.Children) + i
		child.Parent = b
	}
	b.Children = append(b.Children, children...)
	return b
}
