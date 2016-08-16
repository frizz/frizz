package stores

import (
	"context"

	"kego.io/editor/client/actions"
	"kego.io/flux"
)

type ActionStore struct {
	*flux.Store
	ctx context.Context
	app *App

	actions []actions.Undoable
	index   int
}

type actionNotif string

func (actionNotif) IsNotif() {}

const (
	ActionsChanged actionNotif = "ActionsChanged"
)

func NewActionStore(ctx context.Context) *ActionStore {
	s := &ActionStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
	s.Init(s)
	return s
}

func (s *ActionStore) Count() int {
	return len(s.actions)
}

func (s *ActionStore) Index() int {
	return s.index
}

func (s *ActionStore) Redo() actions.Undoable {
	if s.index >= len(s.actions) {
		return nil
	}
	a := s.actions[s.index]
	a.SetRedo()
	return a
}

func (s *ActionStore) Undo() actions.Undoable {
	if s.index <= 0 {
		return nil
	}
	a := s.actions[s.index-1]
	a.SetUndo()
	return a
}

func (s *ActionStore) Handle(payload *flux.Payload) bool {

	u, ok := payload.Action.(actions.Undoable)
	if !ok {
		return true
	}
	switch u.Direction() {
	case actions.New:
		s.actions = append(s.actions[:s.index], u)
		s.index = len(s.actions)
	case actions.Undo:
		if s.index <= 0 {
			break
		}
		s.index--
	case actions.Redo:
		if s.index >= len(s.actions) {
			break
		}
		s.index++
	}
	payload.Notify(nil, ActionsChanged)

	/*
		switch action := payload.Action.(type) {
		case *actions.Add:
			action.Action = &Add{
				Node:   action.Node,
				Parent: action.Parent,
				Key:    action.Key,
				Index:  len(action.Parent.Array),
			}
			s.Add(action.Action)
		case *actions.Delete:
			action.Action = &Delete{
				Parent: action.Node.Parent,
				Node:   action.Node,
			}
			s.Add(action.Action)
		case *actions.Modify:
			action.Action = &Modify{
				Node:   action.Editor.Node,
				Before: action.Editor.Node.Value,
				After:  action.Value,
			}
			s.Add(action.Action)
		case *actions.Reorder:
			if !action.Model.Node.Type.IsNativeArray() {
				s.app.Fail <- kerr.New("EPBQVIICFM", "Must be array")
				break
			}
			action.Action = &Reorder{
				Parent: action.Model.Node,
				Before: action.OldIndex,
				After:  action.NewIndex,
			}
			s.Add(action.Action)
		case *actions.Undo:
			if s.index == 0 {
				break
			}
			s.index--
			action.Action = s.actions[s.index]
		case *actions.Redo:
			if s.index == len(s.actions) {
				break
			}
			s.index++
			action.Action = s.actions[s.index]
		}*/
	return true
}
