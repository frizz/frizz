package models

import "context"

type BranchModel struct {
	ctx      context.Context
	Children []*BranchModel
	Root     bool
	Open     bool
	Contents BranchContentsInterface
	Parent   *BranchModel
	Index    int
}

func NewBranchModel(ctx context.Context, contents BranchContentsInterface) *BranchModel {
	return &BranchModel{ctx: ctx, Contents: contents}
}

type BranchOps string

const (
	BranchOpKeyboard        BranchOps = "BranchOpKeyboard"
	BranchOpClickLabel      BranchOps = "BranchOpClickLabel"
	BranchOpClickSummaryRow BranchOps = "BranchOpClickSummaryRow"
	BranchOpClickEditorLink BranchOps = "BranchOpClickEditorLink"
	BranchOpClickBreadcrumb BranchOps = "BranchOpClickBreadcrumb"
	BranchOpClickToggle     BranchOps = "BranchOpClickToggle"
	BranchOpChildAdded      BranchOps = "BranchOpAddNewNode"
)

func (b *BranchModel) CanOpen() bool {
	if async, ok := b.Contents.(AsyncContentsInterface); ok && !async.Loaded() {
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
	if async, ok := b.Contents.(AsyncContentsInterface); ok && !async.Loaded() {
		return "unknown"
	}
	if b.Open {
		return "minus"
	}
	return "plus"
}

func (b *BranchModel) Insert(index int, children ...*BranchModel) *BranchModel {
	ca := []*BranchModel{}
	for i := 0; i < len(children)+len(b.Children); i++ {
		var c *BranchModel
		if i < index {
			c = b.Children[i]
		} else if i < index+len(children) {
			c = children[i-index]
		} else {
			c = b.Children[i-len(children)]
		}
		c.Index = i
		c.Parent = b
		ca = append(ca, c)
	}
	b.Children = ca
	return b
}

func (b *BranchModel) Append(children ...*BranchModel) *BranchModel {
	return b.Insert(len(b.Children), children...)
}

func (b *BranchModel) DeleteChild(index int) {
	b.Children = append(b.Children[0:index], b.Children[index+1:]...)
	for i := index; i < len(b.Children); i++ {
		// re-index only the children that have changed index
		b.Children[i].Index = i
	}
}

// IsVisible checks if all ancestors are open.
func (b *BranchModel) IsVisible() bool {
	current := b.Parent
	for current != nil {
		if !current.Open {
			return false
		}
		current = current.Parent
	}
	return true
}

// EnsureVisible ensures that all ancestors are open. If any are open, we return
// the oldest closed ancestor. If all are open we return nil.
func (b *BranchModel) EnsureVisible() *BranchModel {
	current := b.Parent
	var closed *BranchModel
	for current != nil {
		if !current.Open {
			closed = current
			current.Open = true
		}
		current = current.Parent
	}
	return closed
}
