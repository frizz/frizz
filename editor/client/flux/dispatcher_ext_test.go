package flux_test

import (
	"testing"

	"kego.io/editor/client/flux"
	"kego.io/kerr/assert"
)

var store struct {
	message *MessageStore
	topic   *TopicStore
}

func TestDispatcher(t *testing.T) {
	d := flux.NewDispatcher()
	store.message = &MessageStore{}
	store.topic = &TopicStore{}
	d.Register(store.message)
	d.Register(store.topic)
	d.Dispatch(AddMessage("a"))
	d.Dispatch(AddTopic("b", "c"))
	m := store.message.GetMessages()
	assert.Equal(t, []string{"a", "c"}, m)
}

type AddMessageAction struct {
	*flux.Action
	Message string
}

func AddMessage(message string) *AddMessageAction {
	return &AddMessageAction{
		Action:  &flux.Action{},
		Message: message,
	}
}

type MessageStore struct {
	messages []string
}

func (m *MessageStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *AddTopicAction:
		payload.WaitFor(store.topic)
		m.messages = append(m.messages, action.Message)
	case *AddMessageAction:
		m.messages = append(m.messages, action.Message)
	}
	return true
}

func (m *MessageStore) GetMessages() []string {
	return m.messages
}

type AddTopicAction struct {
	*flux.Action
	Topic   string
	Message string
}

func AddTopic(topic string, message string) *AddTopicAction {
	return &AddTopicAction{
		Action:  &flux.Action{},
		Topic:   topic,
		Message: message,
	}
}

type TopicStore struct {
	topics []string
}

func (m *TopicStore) Handle(payload *flux.Payload) (finished bool) {
	switch action := payload.Action.(type) {
	case *AddTopicAction:
		m.topics = append(m.topics, action.Topic)
	}
	return true
}

func (m *TopicStore) GetTopics() []string {
	return m.topics
}
