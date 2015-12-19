package messages // import "kego.io/editor/shared/messages"

import (
	"github.com/twinj/uuid"
	"kego.io/context/envctx"
	"kego.io/system"
)

var Ctx = envctx.Dummy("kego.io/editor/shared/messages", map[string]string{})

func NewMessage() *Message {
	return &Message{
		Guid: system.NewString(uuid.NewV4().String()),
	}
}

func NewSourceResponse(name string, ok bool, data string) *SourceResponse {
	return &SourceResponse{
		Object: &system.Object{
			Type: system.NewReference("kego.io/editor/shared/messages", "sourceResponse"),
		},
		Message: NewMessage(),
		Name:    system.NewString(name),
		Found:   system.NewBool(ok),
		Data:    system.NewString(data),
	}
}

func NewSourceRequest(name string) *SourceRequest {
	return &SourceRequest{
		Object: &system.Object{
			Type: system.NewReference("kego.io/editor/shared/messages", "sourceRequest"),
		},
		Message: NewMessage(),
		Name:    system.NewString(name),
	}
}
