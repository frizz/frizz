package stores

import (
	"strings"

	"time"

	"github.com/davelondon/kerr"
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
	BranchClose              branchNotif = "BranchClose"
	BranchSelecting          branchNotif = "BranchSelecting"
	BranchSelected           branchNotif = "BranchSelected"
	BranchChildAdded         branchNotif = "BranchChildAdded"
	BranchChildDeleted       branchNotif = "BranchChildDeleted"
	BranchChildrenReordered  branchNotif = "BranchChildrenReordered"
	BranchSelectControl      branchNotif = "BranchSelectControl"
	BranchUnselectControl    branchNotif = "BranchUnselectControl"
)

// BranchDescendantSelect is passed as the notif data when the specified descendant should be
// selected directly after rendering the branch.
type BranchDescendantSelectData struct {
	Branch *models.BranchModel
	Op     models.BranchOps
}

// BranchSelectOperation is passed as the notif data to modify the function of BranchSelecting
type BranchSelectOperationData struct {
	Op models.BranchOps
}

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
	case *actions.Add:
		payload.Wait(s.app.Nodes)
		switch action.Direction() {
		case actions.New, actions.Redo:
			parent, err := mutateAppendBranch(s, action.Parent, action.Node)
			if err != nil {
				s.app.Fail <- kerr.Wrap("LDBMBRHWHB", err)
				break
			}
			if parent != nil {
				payload.Notify(parent, BranchChildAdded)
			}
		case actions.Undo:
			_, parent, err := mutateDeleteBranch(s, action.Node)
			if err != nil {
				s.app.Fail <- kerr.Wrap("NLFWVSNNTY", err)
				break
			}
			if parent != nil {
				payload.Notify(parent, BranchChildDeleted)
			}
		}
	case *actions.Delete:
		payload.Wait(s.app.Nodes)
		switch action.Direction() {
		case actions.New, actions.Redo:
			branch, parent, err := mutateDeleteBranch(s, action.Node)
			if err != nil {
				s.app.Fail <- kerr.Wrap("QTXPXAKXHH", err)
				break
			}
			if branch != nil {
				action.BranchIndex = branch.Index
			}
			if parent != nil {
				payload.Notify(parent, BranchChildDeleted)
			}
		case actions.Undo:
			parent, err := mutateInsertBranch(s, action.Parent, action.Node, action.BranchIndex)
			if err != nil {
				s.app.Fail <- kerr.Wrap("OOGOEWKPIL", err)
				break
			}
			if parent != nil {
				payload.Notify(parent, BranchChildAdded)
			}
		}
	case *actions.Reorder:
		payload.Wait(s.app.Nodes)
		parent, err := mutateReorderBranch(s, action.Model.Node)
		if err != nil {
			s.app.Fail <- kerr.Wrap("NUQOPWWXHA", err)
			break
		}
		if parent != nil {
			payload.Notify(parent, BranchChildrenReordered)
		}
	case *actions.BranchClose:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			break
		}
		action.Branch.RecursiveClose()
		payload.Notify(action.Branch, BranchClose)
	case *actions.BranchOpening:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			break
		}
		// The branch may not be loaded, so we don't open the branch until the BranchOpenPostLoad
		// action is received. This will happen immediately if the branch is loaded or not async.
		payload.Notify(action.Branch, BranchOpening)
	case *actions.BranchOpened:
		if !action.Branch.CanOpen() {
			// branch can't open - ignore
			break
		}
		action.Branch.Open = true
		payload.Notify(action.Branch, BranchOpened)

	case *actions.BranchSelecting:

		if ancestor := action.Branch.EnsureVisible(); ancestor != nil {
			payload.NotifyWithData(ancestor, BranchOpened, &BranchDescendantSelectData{
				Branch: action.Branch,
				Op:     action.Op,
			})
			break
		}

		s.selected = action.Branch

		payload.Notify(previous, BranchUnselectControl)
		payload.Notify(s.selected, BranchSelectControl)

		if action.Op == models.BranchOpKeyboard {
			go func() {
				<-time.After(time.Millisecond * 50)
				if s.selected == action.Branch {
					payload.NotifyWithData(
						s.selected,
						BranchSelecting,
						&BranchSelectOperationData{Op: action.Op},
					)
				}
			}()
		} else {
			payload.NotifyWithData(
				s.selected,
				BranchSelecting,
				&BranchSelectOperationData{Op: action.Op},
			)
		}

	case *actions.BranchSelected:
		payload.Wait(s.app.Nodes)
		payload.Notify(s.selected, BranchSelected)
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
		payload.Notify(nil, BranchInitialStateLoaded)
	case *actions.LoadSourceSuccess:
		ni, ok := action.Branch.Contents.(models.NodeContentsInterface)
		if !ok {
			break
		}
		n := ni.GetNode()
		s.AppendNodeBranchModelChildren(action.Branch, n)
		s.nodeBranches[n] = action.Branch
		payload.Notify(action.Branch, BranchLoaded)
	}
	return true
}

func mutateDeleteBranch(s *BranchStore, n *node.Node) (branch *models.BranchModel, parent *models.BranchModel, err error) {
	branch, ok := s.nodeBranches[n]
	if !ok {
		return nil, nil, nil
	}
	delete(s.nodeBranches, n)
	for i, c := range branch.Parent.Children {
		if c == branch {
			branch.Parent.DeleteChild(i)
			break
		}
	}
	return branch, branch.Parent, nil
}

func mutateInsertBranch(s *BranchStore, p *node.Node, n *node.Node, i int) (parent *models.BranchModel, err error) {
	if p == nil {
		return nil, nil
	}
	parentBranch, ok := s.nodeBranches[p]
	if !ok {
		return nil, nil
	}
	s.InsertNodeBranchModelChild(parentBranch, n, i)
	return parentBranch, nil
}

func mutateAppendBranch(s *BranchStore, p *node.Node, n *node.Node) (parent *models.BranchModel, err error) {
	if p == nil {
		return nil, nil
	}
	parentBranch, ok := s.nodeBranches[p]
	if !ok {
		return nil, nil
	}
	s.AppendNodeBranchModelChild(parentBranch, n)
	return parentBranch, nil
}

func mutateReorderBranch(s *BranchStore, p *node.Node) (parent *models.BranchModel, err error) {
	parentBranch, ok := s.nodeBranches[p]
	if !ok {
		return nil, nil
	}
	parentBranch.Children = []*models.BranchModel{}
	for _, n := range p.Array {
		b, ok := s.nodeBranches[n]
		if ok {
			parentBranch.Append(b)
		}
	}
	return parentBranch, nil
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

func (s *BranchStore) InsertNodeBranchModelChild(b *models.BranchModel, n *node.Node, index int) {
	if n.Missing || n.Null {
		return
	}
	e := models.GetEditable(s.ctx, n)
	f := e.Format(n.Rule)
	if f == editable.Branch {
		b.Insert(index, s.NewNodeBranchModel(s.ctx, n, ""))
	}
}

func (s *BranchStore) AppendNodeBranchModelChild(b *models.BranchModel, n *node.Node) {
	s.InsertNodeBranchModelChild(b, n, len(b.Children))
}
