package stores

import (
	"sync"

	"time"

	"context"

	"github.com/davelondon/kerr"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/common"
	"kego.io/editor/client/models"
	"kego.io/flux"
	"kego.io/process/validate"
	"kego.io/system"
	"kego.io/system/node"
)

type RuleStore struct {
	*flux.Store
	ctx context.Context
	app *App

	m     *sync.RWMutex
	rules map[*node.Node]map[*node.Node][]system.RuleInterface
}

type ruleNotif string

func (b ruleNotif) IsNotif() {}

//const (
//RuleChanged ruleNotif = "RuleChanged"
//)

func NewRuleStore(ctx context.Context) *RuleStore {
	s := &RuleStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
		m:     &sync.RWMutex{},
		rules: map[*node.Node]map[*node.Node][]system.RuleInterface{},
	}
	s.Init(s)
	return s
}

func (s *RuleStore) Get(r *node.Node, n *node.Node) []system.RuleInterface {
	s.m.RLock()
	defer s.m.RUnlock()
	root, ok := s.rules[r]
	if !ok {
		return nil
	}
	rules, ok := root[n]
	if !ok {
		return nil
	}
	return rules
}

func (s *RuleStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.InitialState:
		payload.Wait(s.app.Package, s.app.Types)
		s.validateNodes(payload, s.build(s.app.Package.Node()))
		for _, t := range s.app.Types.All() {
			s.validateNodes(payload, s.build(t.Node))
		}
	case *actions.LoadFileSuccess:
		payload.Wait(s.app.Branches)
		ni, ok := action.Branch.Contents.(models.NodeContentsInterface)
		if !ok {
			break
		}
		n := ni.GetNode()
		s.validateNodes(payload, s.build(n))
	case *actions.Modify:
		payload.Wait(s.app.Nodes)
		// Node mutations that are in response to keyboard input should wait an
		// additional time before doing a full validation rebuild. This
		// operation may be expensive, so we shouldn't risk interupting the
		// user input.
		if action.Direction() != actions.New || action.Immediate {
			// Undo and Redo actions are from mouse input, so always immediate.
			// New actions are immediate if action.Immediate == true
			s.validateNodes(payload, s.build(action.Editor.Node.Root()))
			break
		}
		// Delayed actions wait an extra time before the validation rebuild.
		go func() {
			// We have already waited for EditorKeyboardDebounceShort, so we
			// reduce the duration by this.
			<-time.After(common.EditorKeyboardDebounceLong - common.EditorKeyboardDebounceShort)
			if action.Changed() {
				return
			}
			s.validateNodes(payload, s.build(action.Editor.Node.Root()))
		}()
	case *actions.Reorder:
		payload.Wait(s.app.Nodes)
		s.validateNodes(payload, s.build(action.Model.Node.Root()))
	case *actions.Delete:
		payload.Wait(s.app.Nodes)
		s.validateNodes(payload, s.build(action.Node.Root()))
	case *actions.Add:
		payload.Wait(s.app.Nodes)
		s.validateNodes(payload, s.build(action.Node.Root()))
	}
	return true
}

func (s *RuleStore) validateNodes(payload *flux.Payload, changes []*node.Node) {
	for _, n := range changes {
		m := s.app.Nodes.Get(n)
		changed, err := m.Validate(s.ctx, s.app.Rule.Get(n.Root(), n))
		if err != nil {
			s.app.Fail <- kerr.Wrap("BYQOBLPRDP", err)
			break
		}
		if changed {
			payload.Notify(n, NodeErrorsChanged)
		}
	}
}

func (s *RuleStore) build(n *node.Node) (changes []*node.Node) {
	rules := map[*node.Node][]system.RuleInterface{}
	if err := validate.BuildRulesNode(s.ctx, n, rules); err != nil {
		s.app.Fail <- kerr.Wrap("BRRRGDBXMR", err)
	}
	changes = s.compare(rules, s.rules[n])
	s.rules[n] = rules
	return changes
}

func (s *RuleStore) compare(a, b map[*node.Node][]system.RuleInterface) (changes []*node.Node) {

	in := func(n system.RuleInterface, h []system.RuleInterface) bool {
		for _, v := range h {
			if n == v {
				return true
			}
		}
		return false
	}

	eq := func(a, b []system.RuleInterface) bool {

		if a == nil && b == nil {
			return true
		}

		if a == nil || b == nil {
			return false
		}

		if len(a) != len(b) {
			return false
		}

		for _, v := range a {
			if !in(v, b) {
				return false
			}
		}

		return true
	}

	for k, av := range a {
		if bv, ok := b[k]; !ok {
			changes = append(changes, k)
		} else {
			if !eq(av, bv) {
				changes = append(changes, k)
			}
		}
	}

	for k := range b {
		if _, ok := a[k]; !ok {
			changes = append(changes, k)
		}
	}

	return changes
}
