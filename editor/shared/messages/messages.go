package messages // import "kego.io/editor/shared/messages"

import (
	"github.com/twinj/uuid"
	"kego.io/system"
)

// ke: {"package": {"notest": true}}

func NewMessage() *Message {
	return &Message{
		Guid: system.NewString(uuid.NewV4().String()),
	}
}

func NewDataResponse(path string, name string, ok bool, data string) *DataResponse {
	return &DataResponse{
		Object: &system.Object{
			Type: system.NewReference("kego.io/editor/shared/messages", "dataResponse"),
		},
		Message: NewMessage(),
		Package: system.NewString(path),
		Name:    system.NewString(name),
		Found:   system.NewBool(ok),
		Data:    system.NewString(data),
	}
}

func NewDataRequest(path string, name string, file string) *DataRequest {
	return &DataRequest{
		Object: &system.Object{
			Type: system.NewReference("kego.io/editor/shared/messages", "dataRequest"),
		},
		Message: NewMessage(),
		Package: system.NewString(path),
		Name:    system.NewString(name),
		File:    system.NewString(file),
	}
}
