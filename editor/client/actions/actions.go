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

type BranchToggle struct {
	Branch *models.BranchModel
}

type BranchSelectClick struct {
	Branch *models.BranchModel
}
type BranchSelectKeyboard struct {
	Branch *models.BranchModel
}
type BranchOpen struct {
	Branch *models.BranchModel
}
type BranchOpenPostLoad struct {
	Branch *models.BranchModel
	Loaded bool
}
type BranchClose struct {
	Branch *models.BranchModel
}
type BranchSelectPostLoad struct {
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
