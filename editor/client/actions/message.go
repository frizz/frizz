package actions

import "kego.io/editor/client/flux"

type NewMessageAction struct {
	*flux.Action
	Message string
}

func NewMessage(message string) *NewMessageAction {
	return &NewMessageAction{
		Action:  &flux.Action{},
		Message: message,
	}
}
