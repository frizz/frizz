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

	root  *models.BranchModel
	pkg   *models.BranchModel
	types *models.BranchModel
	data  *models.BranchModel
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
	case *actions.InitialState:
		payload.WaitFor(s.app.Package, s.app.Types, s.app.Data)
		s.pkg = models.NewNodeBranch(s.app.Package.Get(), "package")

		types := []*models.BranchModel{}
		for name, n := range s.app.Types.All() {
			types = append(types, models.NewNodeBranch(n, name))
		}
		s.types = &models.BranchModel{
			Contents: &models.TypesContents{},
			Children: types,
		}

		data := []*models.BranchModel{}
		for name, d := range s.app.Data.All() {
			data = append(data, &models.BranchModel{
				Contents: &models.SourceContents{
					Name:     name,
					Filename: d.File,
				},
			})
		}
		s.data = &models.BranchModel{
			Contents: &models.DataContents{},
			Open:     true,
			Children: data,
		}

		s.root = &models.BranchModel{
			Root: true,
			Open: true,
			Children: []*models.BranchModel{
				s.pkg,
				s.types,
				s.data,
			},
		}
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
