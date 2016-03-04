package stores

import (
	"code.google.com/p/go.net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/flux"
)

type MessageStore struct {
	*flux.Store
	ctx      context.Context
	app      *App
	messages []string
}

func NewMessageStore(ctx context.Context) *MessageStore {
	return &MessageStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
}

func (m *MessageStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.NewMessage:
		m.messages = append(m.messages, action.Message)
	}
	return true
}
