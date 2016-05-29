package stores

import (
	"strings"

	"time"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/flux"
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

type branchNotif string

func (b branchNotif) IsNotif() {}

const (
	BranchInitialStateLoaded branchNotif = "BranchInitialStateLoaded"
	BranchLoaded             branchNotif = "BranchLoaded"
	BranchOpen               branchNotif = "BranchOpen"
	BranchOpenPostLoad       branchNotif = "BranchOpenPostLoad"
	BranchClose              branchNotif = "BranchClose"
	BranchPreSelect          branchNotif = "BranchPreSelect"
	BranchSelect             branchNotif = "BranchSelect"
	BranchSelectPostLoad     branchNotif = "BranchSelectPostLoad"
	BranchUnselect           branchNotif = "BranchUnselect"
)

func NewBranchStore(ctx context.Context) *BranchStore {
	s := &BranchStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
	s.Init(s)
	return s
}

func (s *BranchStore) Handle(payload *flux.Payload) bool {
	previous := s.selected
	switch action := payload.Action.(type) {
	case *actions.BranchClose:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			return true
		}
		action.Branch.RecursiveClose()
		s.Notify(action.Branch, BranchClose)
	case *actions.BranchOpen:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			return true
		}
		// The branch may not be loaded, so we don't open the branch until the BranchOpenPostLoad
		// action is received. This will happen immediately if the branch is loaded or not async.
		s.Notify(action.Branch, BranchOpen)
	case *actions.BranchOpenPostLoad:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			return true
		}
		action.Branch.Open = true
		s.Notify(action.Branch, BranchOpenPostLoad)
	case *actions.BranchSelect:
		payload.Wait(s.app.Nodes)
		s.selected = action.Branch
		s.selected.LastOp = action.Op
		s.Notify(previous, BranchUnselect)
		s.Notify(s.selected, BranchPreSelect)
		if action.Op == models.BranchOpKeyboard {
			go func() {
				<-time.After(time.Millisecond * 50)
				if s.selected == action.Branch {
					s.Notify(s.selected, BranchSelect)
				}
			}()
		} else {
			s.Notify(s.selected, BranchSelect)
		}
	case *actions.BranchSelectPostLoad:
		payload.Wait(s.app.Nodes)
		s.Notify(nil, BranchSelectPostLoad)
	case *actions.InitialState:
		payload.Wait(s.app.Package, s.app.Types, s.app.Data)
		s.pkg = models.NewNodeBranch(s.ctx, s.app.Package.Node(), "package")

		s.types = models.NewBranchModel(s.ctx, &models.TypesContents{})
		for name, n := range s.app.Types.All() {
			s.types.Append(models.NewNodeBranch(s.ctx, n, name))
		}

		s.data = models.NewBranchModel(s.ctx, &models.DataContents{})
		s.data.Open = true

		for name, d := range s.app.Data.All() {
			s.data.Append(models.NewBranchModel(s.ctx, &models.SourceContents{
				Name:     name,
				Filename: d.File,
			}))
		}

		path := envctx.FromContext(s.ctx).Path
		name := path[strings.LastIndex(path, "/")+1:]

		s.root = models.NewBranchModel(s.ctx, &models.RootContents{
			Name: name,
		})
		s.root.Root = true
		s.root.Open = true

		s.root.Append(s.pkg, s.types, s.data)
		s.Notify(nil, BranchInitialStateLoaded)
	case *actions.LoadSourceSuccess:
		ni, ok := action.Branch.Contents.(models.NodeContentsInterface)
		if !ok {
			return true
		}
		models.AppendNodeChildren(action.Branch, ni.GetNode())
		s.Notify(action.Branch, BranchLoaded)
	}

	return true
}

/*
// We don't currently need to eliminate descendants because we never have
// to update a descendant at the same time it's ancestor. I'll leave the
// code in here in case we need it.
func (s *BranchStore) NotifySingle(notificationType interface{}, changed ...interface{}) {

	if notificationType == BranchSelectedImmediate {
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

*/
