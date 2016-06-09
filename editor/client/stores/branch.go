package stores

import (
	"strings"

	"time"

	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/editable"
	"kego.io/editor/client/models"
	"kego.io/flux"
	"kego.io/system/node"
)

type BranchStore struct {
	*flux.Store
	ctx context.Context
	app *App

	selected     *models.BranchModel
	root         *models.BranchModel
	pkg          *models.BranchModel
	types        *models.BranchModel
	data         *models.BranchModel
	nodeBranches map[*node.Node]*models.BranchModel
}

func (b *BranchStore) Get(n *node.Node) *models.BranchModel {
	return b.nodeBranches[n]
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
	BranchOpening            branchNotif = "BranchOpening"
	BranchOpened             branchNotif = "BranchOpened"
	BranchDescendantSelected branchNotif = "BranchDescendantSelected"
	BranchClose              branchNotif = "BranchClose"
	BranchSelecting          branchNotif = "BranchSelecting"
	BranchSelected           branchNotif = "BranchSelected"
	BranchChildAdded         branchNotif = "BranchChildAdded"
	BranchSelectControl      branchNotif = "BranchSelectControl"
	BranchUnselectControl    branchNotif = "BranchUnselectControl"
)

func NewBranchStore(ctx context.Context) *BranchStore {
	s := &BranchStore{
		Store:        &flux.Store{},
		ctx:          ctx,
		app:          FromContext(ctx),
		nodeBranches: map[*node.Node]*models.BranchModel{},
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
	case *actions.BranchOpening:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			return true
		}
		// The branch may not be loaded, so we don't open the branch until the BranchOpenPostLoad
		// action is received. This will happen immediately if the branch is loaded or not async.
		s.Notify(action.Branch, BranchOpening)
	case *actions.BranchOpened:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			return true
		}
		action.Branch.Open = true
		s.Notify(action.Branch, BranchOpened)
	case *actions.BranchDescendantSelect:

		s.selected = action.Branch
		s.selected.LastOp = action.Op

		ancestor := action.Branch.EnsureVisible()
		if ancestor == nil {
			break
		}

		s.Notify(ancestor, BranchDescendantSelected)

	case *actions.BranchSelecting:

		s.selected = action.Branch
		s.selected.LastOp = action.Op

		if oldestClosedAncestor := action.Branch.EnsureVisible(); oldestClosedAncestor != nil {
			s.Notify(oldestClosedAncestor, BranchDescendantSelected)
			break
		}

		s.Notify(previous, BranchUnselectControl)
		s.Notify(s.selected, BranchSelectControl)

		if action.Op == models.BranchOpKeyboard {
			go func() {
				<-time.After(time.Millisecond * 50)
				if s.selected == action.Branch {
					s.Notify(s.selected, BranchSelecting)
				}
			}()
		} else {
			s.Notify(s.selected, BranchSelecting)
		}

	case *actions.BranchSelected:
		payload.Wait(s.app.Nodes)
		s.Notify(nil, BranchSelected)
	case *actions.InitialState:
		payload.Wait(s.app.Package, s.app.Types, s.app.Data)

		pkgNode := s.app.Package.Node()
		pkgBranch := s.NewNodeBranchModel(s.ctx, pkgNode, "package")
		s.pkg = pkgBranch

		s.types = models.NewBranchModel(s.ctx, &models.TypesContents{})
		for _, name := range s.app.Types.Names() {
			typeBranch := s.NewNodeBranchModel(s.ctx, s.app.Types.Get(name), name)
			s.types.Append(typeBranch)
		}

		s.data = models.NewBranchModel(s.ctx, &models.DataContents{})
		s.data.Open = true

		for _, name := range s.app.Data.Names() {
			s.data.Append(models.NewBranchModel(s.ctx, &models.SourceContents{
				Name:     name,
				Filename: s.app.Data.Get(name).File,
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
		n := ni.GetNode()
		s.AppendNodeBranchModelChildren(action.Branch, n)
		s.nodeBranches[n] = action.Branch
		s.Notify(action.Branch, BranchLoaded)
	case *actions.InitializeNode:
		payload.Wait(s.app.Nodes)
		parent := action.Node.Parent
		if parent == nil {
			break
		}
		parentBranch, ok := s.nodeBranches[parent]
		if !ok {
			break
		}
		s.AppendNodeBranchModelChild(parentBranch, action.Node)
		s.Notify(parentBranch, BranchChildAdded)
	}

	return true
}

func (s *BranchStore) NewNodeBranchModel(ctx context.Context, n *node.Node, name string) *models.BranchModel {
	b := models.NewBranchModel(ctx, &models.NodeContents{
		Node: n,
		Name: name,
	})
	s.AppendNodeBranchModelChildren(b, n)
	s.nodeBranches[n] = b
	return b
}

func (s *BranchStore) AppendNodeBranchModelChildren(b *models.BranchModel, n *node.Node) {
	for _, c := range n.Array {
		s.AppendNodeBranchModelChild(b, c)
	}
	for _, c := range n.Map {
		s.AppendNodeBranchModelChild(b, c)
	}
}

func (s *BranchStore) AppendNodeBranchModelChild(b *models.BranchModel, n *node.Node) {
	if n.Missing || n.Null {
		return
	}
	f := models.GetEditable(s.ctx, n).Format()
	if f == editable.Branch {
		b.Append(s.NewNodeBranchModel(s.ctx, n, ""))
	}
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
