package messages // import "kego.io/editor/shared/messages"

import (
	"github.com/twinj/uuid"
	"kego.io/system"
)

var Path = "kego.io/editor/shared/messages"
var Aliases = map[string]string{}

type Message interface {
	Message() *Message_base
}

func (b *Message_base) Message() *Message_base {
	return b
}

func NewMessageBase() *Message_base {
	return &Message_base{
		Guid: system.NewString(uuid.NewV4().String()),
	}
}

func NewSourceResponse(name string, ok bool, data string) *SourceResponse {
	return &SourceResponse{
		Object_base: &system.Object_base{
			Type: system.NewReference("kego.io/editor/shared/messages", "sourceResponse"),
		},
		Message_base: NewMessageBase(),
		Name:         system.NewString(name),
		Found:        system.NewBool(ok),
		Data:         system.NewString(data),
	}
}

func NewSourceRequest(name string) *SourceRequest {
	return &SourceRequest{
		Object_base: &system.Object_base{
			Type: system.NewReference("kego.io/editor/shared/messages", "sourceRequest"),
		},
		Message_base: NewMessageBase(),
		Name:         system.NewString(name),
	}
}
