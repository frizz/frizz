package stores

import (
	"golang.org/x/net/context"
	"kego.io/context/envctx"
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

type branchKey string

const BranchSelectedChanged branchKey = "BranchSelectedChanged"

func NewBranchStore(ctx context.Context) *BranchStore {
	s := &BranchStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
	s.Init(s)
	return s
}

func (s *BranchStore) NotifyOne(notificationType interface{}, changed ...interface{}) {

	if notificationType == BranchSelectedChanged {
		s.Store.NotifyOne(notificationType, changed...)
		return
	}

	// eliminate descendants...
	changedBranches := []*models.BranchModel{}
	for _, c := range changed {
		br := c.(*models.BranchModel)
		if br != nil {
			changedBranches = append(changedBranches, br)
		}
	}
	deleted := map[interface{}]bool{}
	for _, b := range changedBranches {
		for _, b1 := range changedBranches {
			if b.IsDescendantOf(b1) {
				deleted[b] = true
			}
		}
	}
	out := []interface{}{}
	for _, b := range changedBranches {
		if _, ok := deleted[b]; !ok {
			out = append(out, b)
		}
	}
	s.Store.NotifyOne(notificationType, out...)
}

func (s *BranchStore) Handle(payload *flux.Payload) bool {
	previous := s.selected
	switch action := payload.Action.(type) {
	case *actions.KeyboardEvent:
		switch action.KeyCode {
		case 38: // up
			if s.selected == nil {
				if b := s.app.Branches.Root().LastVisible(); b != nil {
					s.selected = b
					s.NotifyOne(BranchSelectedChanged, previous, s.selected)
				}
				return true
			}
			if b := s.selected.PrevVisible(); b != nil {
				s.selected = b
				s.NotifyOne(BranchSelectedChanged, previous, s.selected)
				return true
			}
		case 40: // down
			if s.selected == nil {
				s.selected = s.app.Branches.Root()
				s.NotifyOne(BranchSelectedChanged, s.selected)
				return true
			}
			if b := s.selected.NextVisible(true); b != nil {
				s.selected = b
				s.NotifyOne(BranchSelectedChanged, previous, s.selected)
				return true
			}
		case 37: // left
			if s.selected == nil {
				return true
			}
			if s.selected.CanOpen() && s.selected.Open {
				// if the branch is open, left arrow should close it.
				s.selected.Open = false
				s.NotifyAll(s.selected)
				return true
			} else {
				if b := s.selected.Parent; b != nil {
					s.selected = b
					s.NotifyOne(BranchSelectedChanged, previous, s.selected)
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
				s.NotifyAll(s.selected)
				return true
			} else {
				if b := s.selected.FirstChild(); b != nil {
					s.selected = b
					s.NotifyOne(BranchSelectedChanged, previous, s.selected)
					return true
				}
			}
		}
	case *actions.ToggleBranch:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			break
		}
		action.Branch.Open = !action.Branch.Open
		s.NotifyAll(action.Branch)
		return true
	case *actions.SelectBranch:
		s.selected = action.Branch
		s.NotifyOne(BranchSelectedChanged, previous, s.selected)
		return true
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
			Contents: &models.RootContents{
				Path: envctx.FromContext(s.ctx).Path,
			},
		}
		s.root.Append(s.pkg, s.types, s.data)
		s.NotifyAll()
		return true
	}

	return true
}
