package stores

import (
	"sync"

	"github.com/davelondon/kerr"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/flux"
	"kego.io/json"
	"kego.io/system"
	"kego.io/system/node"
)

type NodeStore struct {
	*flux.Store
	ctx context.Context
	app *App

	selected *node.Node
	m        *sync.Mutex
	models   map[*node.Node]*models.NodeModel
}

func (s *NodeStore) Selected() *node.Node {
	return s.selected
}

func (s *NodeStore) Get(n *node.Node) *models.NodeModel {
	s.m.Lock()
	defer s.m.Unlock()
	m, ok := s.models[n]
	if !ok {
		m = &models.NodeModel{Node: n}
		s.models[n] = m
	}
	return m
}

type nodeNotif string

func (b nodeNotif) IsNotif() {}

const (
	NodeInitialised            nodeNotif = "NodeInitialised"
	NodeValueChanged           nodeNotif = "NodeValueChanged"
	NodeDescendantValueChanged nodeNotif = "NodeDescendantValueChanged"
	NodeFocus                  nodeNotif = "NodeFocus"
	NodeErrorsChanged          nodeNotif = "NodeErrorsChanged"
)

func NewNodeStore(ctx context.Context) *NodeStore {
	s := &NodeStore{
		Store:  &flux.Store{},
		ctx:    ctx,
		app:    FromContext(ctx),
		m:      &sync.Mutex{},
		models: map[*node.Node]*models.NodeModel{},
	}
	s.Init(s)
	return s
}

func (s *NodeStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.Add:
		payload.Wait(s.app.Actions)
		if action.Forward() {
			if err := mutateAddNode(s, action.Node, action.Parent, action.Key, len(action.Parent.Array), action.Type); err != nil {
				s.app.Fail <- kerr.Wrap("HUOGBUQCAO", err)
				break
			}
		} else {
			if err := mutateUndoAddNode(s, action.Node, action.Key, len(action.Parent.Array)); err != nil {
				s.app.Fail <- kerr.Wrap("RTAGMUIKMD", err)
				break
			}
		}
	case *actions.Delete:
		payload.Wait(s.app.Actions)
		if action.Forward() {
			if err := mutateDeleteNode(s, action.Node, action.Node.Parent, action.Backup); err != nil {
				s.app.Fail <- kerr.Wrap("DFHTKJRLQC", err)
				break
			}
		} else {
			if err := mutateUndoDeleteNode(s, action.Node, action.Node.Parent, action.Backup); err != nil {
				s.app.Fail <- kerr.Wrap("HAPWUOPBTW", err)
				break
			}
		}
	case *actions.Reorder:
		payload.Wait(s.app.Actions)
		from := action.Before
		to := action.After
		if action.Backward() {
			from = action.After
			to = action.Before
		}
		if err := action.Model.Node.ReorderArrayChild(from, to); err != nil {
			s.app.Fail <- kerr.Wrap("HMIFPKVJCN", err)
			break
		}

	case *actions.Modify:
		payload.Wait(s.app.Actions)
		n := action.Editor.Node
		val := action.After
		if action.Backward() {
			val = action.Before
		}
		if err := n.SetValue(s.ctx, val); err != nil {
			s.app.Fail <- kerr.Wrap("VIMXVIHPFY", err)
			break
		}

		model := s.app.Nodes.Get(n)
		changed, err := model.Validate(s.ctx, s.app.Rule.Get(n.Root(), n))
		if err != nil {
			s.app.Fail <- kerr.Wrap("EEIYMGQCCA", err)
			break
		}

		if changed {
			s.app.Notify(n, NodeErrorsChanged)
			s.app.Notify(action.Editor, EditorErrorsChanged)
		}

		s.app.Notify(n, NodeValueChanged)
		s.app.Notify(action.Editor, EditorValueChanged)

		c := n.Parent
		for c != nil {
			s.app.Notify(c, NodeDescendantValueChanged)
			c = c.Parent
		}
	case *actions.BranchSelecting:
		if ni, ok := action.Branch.Contents.(models.NodeContentsInterface); ok {
			s.selected = ni.GetNode()
		} else {
			s.selected = nil
		}
	case *actions.BranchSelected:
		if ni, ok := action.Branch.Contents.(models.NodeContentsInterface); ok {
			s.selected = ni.GetNode()
		} else {
			s.selected = nil
		}
	case *actions.EditorFocus:
		s.app.Notify(action.Editor.Node, NodeFocus)
	}
	return true
}

func mutateDeleteNode(s *NodeStore, n *node.Node, p *node.Node, b *node.Node) error {
	switch p.Type.NativeJsonType() {
	case json.J_MAP:
		if err := p.DeleteMapChild(n.Key); err != nil {
			return kerr.Wrap("BUUOWYSJNG", err)
		}
	case json.J_ARRAY:
		if err := p.DeleteArrayChild(n.Index); err != nil {
			return kerr.Wrap("RWFQSINACH", err)
		}
	case json.J_OBJECT:
		*b = *n.Backup()
		if err := p.DeleteObjectChild(s.ctx, n.Key); err != nil {
			return kerr.Wrap("XGVEXEOBUP", err)
		}
	}
	return nil
}

func mutateUndoDeleteNode(s *NodeStore, n *node.Node, p *node.Node, b *node.Node) error {
	switch p.Type.NativeJsonType() {
	case json.J_MAP:
		if err := n.AddToMap(s.ctx, p, n.Key, false); err != nil {
			return kerr.Wrap("TOPLOONYCL", err)
		}
	case json.J_ARRAY:
		if err := n.AddToArray(s.ctx, p, n.Index, false); err != nil {
			return kerr.Wrap("WFXSQYOEAY", err)
		}
	case json.J_OBJECT:
		if err := n.Restore(s.ctx, b); err != nil {
			return kerr.Wrap("EVSGQSPUPT", err)
		}
	}
	return nil
}

func mutateAddNode(s *NodeStore, n *node.Node, p *node.Node, key string, index int, t *system.Type) error {
	switch p.Type.NativeJsonType() {
	case json.J_ARRAY:
		if err := n.AddToArray(s.ctx, p, index, false); err != nil {
			return kerr.Wrap("PLEJOTCSGH", err)
		}
	case json.J_MAP:
		if err := n.AddToMap(s.ctx, p, key, false); err != nil {
			return kerr.Wrap("UEPLLMTLDB", err)
		}
	}
	if err := n.SetValueZero(s.ctx, false, t); err != nil {
		return kerr.Wrap("NLSRNQGLLW", err)
	}
	return nil
}

func mutateUndoAddNode(s *NodeStore, p *node.Node, key string, index int) error {
	switch p.Type.NativeJsonType() {
	case json.J_ARRAY:
		if err := p.DeleteArrayChild(index); err != nil {
			return kerr.Wrap("NJLSXKSGMX", err)
		}
	case json.J_MAP:
		if err := p.DeleteMapChild(key); err != nil {
			return kerr.Wrap("UQVCSSNLUO", err)
		}
	case json.J_OBJECT:
		if err := p.DeleteObjectChild(s.ctx, key); err != nil {
			return kerr.Wrap("MXVPORUPJS", err)
		}
	}
	return nil
}
