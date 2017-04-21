package actions

import (
	"frizz.io/editor/client/models"
	"frizz.io/editor/shared"
	"frizz.io/flux"
	"frizz.io/system"
	"frizz.io/system/node"
)

type ToggleSystemControls struct {
	Node *node.Node
}

type ToggleInfoState struct{}

type ChangeView struct {
	View models.Views
}

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

type LoadFile struct {
	Contents models.BranchContentsInterface
	Signal   chan struct{}
}

type KeyboardEvent struct {
	KeyCode int
}

type LoadFileSent struct {
	Branch *models.BranchModel
}
type LoadFileCancelled struct {
	Branch *models.BranchModel
}
type LoadFileSuccess struct {
	Branch *models.BranchModel
}
type LoadFileError struct {
	Branch *models.BranchModel
}

type Directions int

const (
	New  Directions = 0
	Undo Directions = 1
	Redo Directions = 2
)

type Mutator interface {
	CommonAncestor() *node.Node
}

type Undoable interface {
	Direction() Directions
	SetUndo()
	SetRedo()
	Description() string
}

type Undoer struct {
	direction Directions
}

func (u *Undoer) Direction() Directions { return u.direction }
func (u *Undoer) SetUndo()              { u.direction = Undo }
func (u *Undoer) SetRedo()              { u.direction = Redo }

type Add struct {
	*Undoer
	Node       *node.Node
	Parent     *node.Node
	Backup     *node.Node
	Key        string
	Index      int
	Type       *system.Type
	BranchName string
	BranchFile string
}

func (a *Add) Description() string {
	return "add"
}

func (a *Add) CommonAncestor() *node.Node {
	return a.Parent
}

type Delete struct {
	*Undoer
	Node        *node.Node
	Parent      *node.Node
	Backup      *node.Node
	BranchIndex int
	BranchName  string
	BranchFile  string
}

func (a *Delete) Description() string {
	return "delete"
}

func (a *Delete) CommonAncestor() *node.Node {
	return a.Parent
}

type Reorder struct {
	*Undoer
	Model  *models.EditorModel
	Before int
	After  int
}

func (a *Reorder) Description() string {
	return "reorder"
}

func (a *Reorder) CommonAncestor() *node.Node {
	return a.Model.Node
}

type EditorData struct {
	Func func(payload *flux.Payload)
}

type Modify struct {
	*Undoer
	Editor    *models.EditorModel
	After     interface{}
	Before    interface{}
	Changed   func() bool
	Immediate bool
}

func (a *Modify) Description() string {
	return "modify"
}

func (a *Modify) CommonAncestor() *node.Node {
	return a.Editor.Node
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

type SaveFileSuccess struct {
	Response *shared.SaveResponse
}
