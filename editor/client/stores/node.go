package stores

import (
	"github.com/davelondon/kerr"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/flux"
	"kego.io/system/node"
)

type NodeStore struct {
	*flux.Store
	ctx context.Context
	app *App

	selected *node.Node
}

func (b *NodeStore) Selected() *node.Node {
	return b.selected
}

type nodeNotif string

func (b nodeNotif) IsNotif() {}

const (
	NodeInitialised nodeNotif = "NodeInitialised"
)

func NewNodeStore(ctx context.Context) *NodeStore {
	s := &NodeStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
	s.Init(s)
	return s
}

func (s *NodeStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.BranchSelect:
		s.selected = nil
		if ni, ok := action.Branch.Contents.(models.NodeContentsInterface); ok {
			s.selected = ni.GetNode()
		}
	case *actions.BranchSelectPostLoad:
		s.selected = nil
		if ni, ok := action.Branch.Contents.(models.NodeContentsInterface); ok {
			s.selected = ni.GetNode()
		}
	case *actions.AddNodeClick:

		node := action.Node

		if !node.Rule.Parent.Interface && !node.Rule.Struct.Interface {
			// if the rule only permits a single concrete type, we initialise immediately with that type
			if err := node.InitialiseWithConcreteType(s.ctx, nil); err != nil {
				s.app.Fail <- kerr.Wrap("KBDDDIVSWK", err)
			}
			return true
		}

		/*
			// If the rule is an interface, so we must pop up UX to choose the concrete type.
			types := system.GetAllTypesThatImplementInterface(e.ctx, node.Rule.Parent)

			if len(types) == 1 {
				// if only one type is compatible, don't show the popup, just add it.
				e.InitialiseChildWithConcreteType(node, types[0])
				return
			}

			card := mdl.Card("Choose a type", "Add")
			options := map[string]string{}

			for _, t := range types {
				displayName, err := t.Id.ValueContext(e.ctx)
				if err != nil {
					// we shouldn't be able to get here
					e.fail <- kerr.Wrap("IPLHSXDWQK", err)
				}
				options[t.Id.String()] = displayName
			}

			dropdown := mdl.Select(options)
			card.Content.AppendChild(dropdown)

			go func() {
				result := <-card.Result
				if result {
					r, err := system.NewReferenceFromString(e.ctx, dropdown.Value)
					if err != nil {
						e.fail <- kerr.Wrap("KHJGQXORPD", err)
					}
					t, ok := system.GetTypeFromCache(e.ctx, r.Package, r.Name)
					if !ok {
						e.fail <- kerr.New("WEADSXTPYC", "Type %s not found in cache", r.Value())
					}
					e.InitialiseChildWithConcreteType(node, t)
				}
				card.Hide()
			}()
		*/

	}
	return true
}
