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

type Directions int

const (
	New  Directions = 0
	Undo Directions = 1
	Redo Directions = 2
)

type Undoable interface {
	Direction() Directions
	SetUndo()
	SetRedo()
}

type Undoer struct {
	direction Directions
}

func (u *Undoer) Direction() Directions { return u.direction }
func (u *Undoer) SetUndo()              { u.direction = Undo }
func (u *Undoer) SetRedo()              { u.direction = Redo }

type Add struct {
	*Undoer
	Node   *node.Node
	Parent *node.Node
	Key    string
	Index  int
	Type   *system.Type
}

type Delete struct {
	*Undoer
	Node        *node.Node
	Parent      *node.Node
	Backup      *node.Node
	BranchIndex int
}

type Reorder struct {
	*Undoer
	Model  *models.EditorModel
	Before int
	After  int
}

type Modify struct {
	*Undoer
	Editor *models.EditorModel
	After  interface{}
	Before interface{}
}

type OpenAddPopup struct {
	Parent *node.Node
	Node   *node.Node
	Types  []*system.Type
}
type CloseAddPopup struct{}

type EditorFocus struct {
	Editor *models.EditorModel
}

type EditorValueChange500ms struct {
	Editor *models.EditorModel
	Value  interface{}
}
