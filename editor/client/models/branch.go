package models

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
