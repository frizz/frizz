package flux_test

import (
	"testing"

	"kego.io/flux"
	"kego.io/kerr/assert"
)

type App struct {
	Dispatcher *flux.Dispatcher
	Messages   *MessageStore
	Topics     *TopicStore
}

func TestDispatcher(t *testing.T) {

	a := &App{}
	a.Messages = &MessageStore{Store: &flux.Store{}, app: a}
	a.Topics = &TopicStore{Store: &flux.Store{}, app: a}
	a.Dispatcher = flux.NewDispatcher(a.Messages, a.Topics)

	a.Dispatcher.Dispatch(&AddMessage{Message: "a"})
	a.Dispatcher.Dispatch(&AddTopic{Topic: "b", Message: "c"})
	msg := a.Messages.GetMessages()
	assert.Equal(t, []string{"a", "c"}, msg)
}

type AddMessage struct {
	Message string
}

type MessageStore struct {
	*flux.Store
	app      *App
	messages []string
}

func (m *MessageStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *AddTopic:
		payload.WaitFor(m.app.Topics)
		m.messages = append(m.messages, action.Message)
	case *AddMessage:
		m.messages = append(m.messages, action.Message)
	}
	return true
}

func (m *MessageStore) GetMessages() []string {
	return m.messages
}

type AddTopic struct {
	Topic   string
	Message string
}

type TopicStore struct {
	*flux.Store
	app    *App
	topics []string
}

func (m *TopicStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *AddTopic:
		m.topics = append(m.topics, action.Topic)
	}
	return true
}

func (m *TopicStore) GetTopics() []string {
	return m.topics
}
