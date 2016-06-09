package actions

import (
	"kego.io/editor/client/models"
	"kego.io/editor/shared"
	"kego.io/system"
	"kego.io/system/node"
)

type NewMessage struct {
	Message string
}

type InitialState struct {
	Info *shared.Info
}

type BranchDescendantSelect struct {
	Branch *models.BranchModel
	Op     models.BranchOps
}

type BranchSelecting struct {
	Branch *models.BranchModel
	Op     models.BranchOps
}
type BranchSelected struct {
	Branch *models.BranchModel
	Loaded bool
}

type BranchOpening struct {
	Branch *models.BranchModel
}
type BranchOpened struct {
	Branch *models.BranchModel
	Loaded bool
}

type BranchClose struct {
	Branch *models.BranchModel
}

type LoadSource struct {
	Contents models.BranchContentsInterface
	Signal   chan struct{}
}

type KeyboardEvent struct {
	KeyCode int
}

type LoadSourceSent struct {
	Branch *models.BranchModel
}
type LoadSourceCancelled struct {
	Branch *models.BranchModel
}
type LoadSourceSuccess struct {
	Branch *models.BranchModel
}
type LoadSourceError struct {
	Branch *models.BranchModel
}

type InitializeNode struct {
	Node   *node.Node
	Parent *node.Node
	New    bool
	Key    string
	Index  int
	Type   *system.Type
}

type OpenAddPop struct {
	Parent *node.Node
	Node   *node.Node
	Types  []*system.Type
}
type CloseAddPop struct{}

type FocusNode struct {
	Node *node.Node
}
