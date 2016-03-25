package actions

import (
	"kego.io/editor/client/models"
	"kego.io/editor/shared"
)

type NewMessage struct {
	Message string
}

type InitialState struct {
	Info *shared.Info
}

type ToggleBranch struct {
	Branch *models.BranchModel
}

type SelectBranch struct {
	Branch   *models.BranchModel
	Keyboard bool
}

type LoadSource struct {
	Contents models.BranchContentsInterface
	Signal   chan struct{}
}
