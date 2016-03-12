package models

type BranchModel struct {
	Root bool
	Open bool
}

func (b *BranchModel) CanOpen() bool {
	return true
}
