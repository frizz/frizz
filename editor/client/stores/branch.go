package stores

import (
	"strings"

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

func (s *BranchStore) NotifySingle(notificationType interface{}, changed ...interface{}) {

	if notificationType == BranchSelectedChanged {
		s.Store.NotifySingle(notificationType, changed...)
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
	s.Store.NotifySingle(notificationType, out...)
}

func (s *BranchStore) Handle(payload *flux.Payload) bool {
	previous := s.selected
	switch action := payload.Action.(type) {
	case *actions.ToggleBranch:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			return true
		}
		action.Branch.Open = !action.Branch.Open
		if !action.Branch.Open {
			action.Branch.CloseAllChildren()
		}
		s.Notify(action.Branch)
		return true
	case *actions.CloseBranch:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			return true
		}
		action.Branch.Open = false
		action.Branch.CloseAllChildren()
		s.Notify(action.Branch)
		return true
	case *actions.OpenBranch:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			return true
		}
		action.Branch.Open = true
		s.Notify(action.Branch)
		return true
	case *actions.SelectBranch:
		s.selected = action.Branch
		s.NotifySingle(BranchSelectedChanged, previous, s.selected)
		return true
	case *actions.InitialState:
		payload.WaitFor(s.app.Package, s.app.Types, s.app.Data)
		s.pkg = models.NewNodeBranch(s.app.Package.Node(), "package")

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

		path := envctx.FromContext(s.ctx).Path
		name := path[strings.LastIndex(path, "/")+1:]

		s.root = &models.BranchModel{
			Root: true,
			Open: true,
			Contents: &models.RootContents{
				Name: name,
			},
		}
		s.root.Append(s.pkg, s.types, s.data)
		s.Notify()
		return true
	case *actions.LoadSourceSuccess:
		n, ok := action.Branch.Contents.(models.NodeContentsInterface)
		if !ok {
			return true
		}
		models.AppendNodeChildren(action.Branch, n.GetNode())
		s.Notify(action.Branch)
		return true
	}

	return true
}
