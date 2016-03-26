package stores

import (
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
	"kego.io/editor/client/models"
)

type BranchStore struct {
	*flux.Store
	ctx context.Context
	app *App

	selected *models.BranchModel
	root     *models.BranchModel
	pkg      *models.BranchModel
	types    *models.BranchModel
	data     *models.BranchModel
}

func (b *BranchStore) Selected() *models.BranchModel {
	return b.selected
}

func (b *BranchStore) Root() *models.BranchModel {
	return b.root
}

func (b *BranchStore) Package() *models.BranchModel {
	return b.pkg
}

func (b *BranchStore) Types() *models.BranchModel {
	return b.types
}

func (b *BranchStore) Data() *models.BranchModel {
	return b.data
}

func NewBranchStore(ctx context.Context) *BranchStore {
	return &BranchStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
}

func (s *BranchStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.KeyboardEvent:
		switch action.KeyCode {
		case 38: // up
			if s.selected == nil {
				if b := s.app.Branches.Root().LastVisible(); b != nil {
					s.selected = b
					s.Notify()
				}
				return true
			}
			if b := s.selected.PrevVisible(); b != nil {
				s.selected = b
				s.Notify()
				return true
			}
		case 40: // down
			if s.selected == nil {
				if b := s.app.Branches.Root().FirstChild(); b != nil {
					s.selected = b
					s.Notify()
				}
				return true
			}
			if b := s.selected.NextVisible(true); b != nil {
				s.selected = b
				s.Notify()
				return true
			}
		case 37: // left
			if s.selected == nil {
				return true
			}
			if s.selected.CanOpen() && s.selected.Open {
				// if the branch is open, left arrow should close it.
				s.selected.Open = false
				s.Notify()
				return true
			} else {
				if b := s.selected.Parent; b != nil {
					s.selected = b
					s.Notify()
					return true
				}
			}
		case 39: // right
			if s.selected == nil {
				return true
			}
			if s.selected.CanOpen() && !s.selected.Open {
				// if the branch is closed, right arrow should open it
				s.selected.Open = true
				s.Notify()
				return true
			} else {
				if b := s.selected.FirstChild(); b != nil {
					s.selected = b
					s.Notify()
					return true
				}
			}
		}
	case *actions.SelectBranch:
		s.selected = action.Branch
		s.Notify()
	case *actions.InitialState:
		payload.WaitFor(s.app.Package, s.app.Types, s.app.Data)
		s.pkg = models.NewNodeBranch(s.app.Package.Get(), "package")

		s.types = &models.BranchModel{
			Contents: &models.TypesContents{},
		}
		for name, n := range s.app.Types.All() {
			s.types.Append(models.NewNodeBranch(n, name))
		}

		s.data = &models.BranchModel{
			Contents: &models.DataContents{},
			Open:     true,
		}
		for name, d := range s.app.Data.All() {
			s.data.Append(&models.BranchModel{
				Contents: &models.SourceContents{
					Name:     name,
					Filename: d.File,
				},
			})
		}

		s.root = &models.BranchModel{
			Root: true,
			Open: true,
		}
		s.root.Append(s.pkg, s.types, s.data)

		s.Notify()
	case *actions.ToggleBranch:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			break
		}
		action.Branch.Open = !action.Branch.Open
		s.Notify()
	}

	return true
}
