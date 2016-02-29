package stores

import (
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
)

type MessageStore struct {
	messages []string
}

func (m *MessageStore) Handle(in flux.ActionInterface) {
	switch action := in.(type) {
	case *actions.NewMessageAction:
		m.messages = append(m.messages, action.Message)
	}
}
