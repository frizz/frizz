package actions

import (
	"kego.io/editor/client/models"
	"kego.io/editor/shared"
	"kego.io/system/node"
)

type NewMessage struct {
	Message string
}

type InitialState struct {
	Info *shared.Info
}

type BranchSelect struct {
	Branch *models.BranchModel
	Op     models.BranchOps
}
type BranchSelectPostLoad struct {
	Branch *models.BranchModel
	Loaded bool
}
type BranchOpen struct {
	Branch *models.BranchModel
}
type BranchClose struct {
	Branch *models.BranchModel
}
type BranchOpenPostLoad struct {
	Branch *models.BranchModel
	Loaded bool
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
type AddNodeClick struct {
	Node *node.Node
}
type AddCollectionItemClick struct {
	Parent *node.Node
}
type AddModalCloseClick struct{}
type AddModalSaveClick struct{}
