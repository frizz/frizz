package messages

import (
	"github.com/twinj/uuid"
	"kego.io/system"
)

var Path = "kego.io/editor/shared/messages"
var Aliases = map[string]string{}

type Message interface {
	GetMessageBase() *MessageBase
}

func (b *MessageBase) GetMessageBase() *MessageBase {
	return b
}

func NewMessageBase() *MessageBase {
	return &MessageBase{
		Guid: system.NewString(uuid.NewV4().String()),
	}
}

func NewGlobalResponse(name string, ok bool, data string) *GlobalResponse {
	return &GlobalResponse{
		Base: &system.Base{
			Type: system.NewReference("kego.io/editor/shared/messages", "globalResponse"),
		},
		MessageBase: NewMessageBase(),
		Name:        system.NewString(name),
		Found:       system.NewBool(ok),
		Data:        system.NewString(data),
	}
}

func NewGlobalRequest(name string) *GlobalRequest {
	return &GlobalRequest{
		Base: &system.Base{
			Type: system.NewReference("kego.io/editor/shared/messages", "globalRequest"),
		},
		MessageBase: NewMessageBase(),
		Name:        system.NewString(name),
	}
}
