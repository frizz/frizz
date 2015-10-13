package messages

import (
	"reflect"

	"kego.io/json"
	"kego.io/system"
)

// This is a base class embedded in each message
type Message_base struct {
	// A globally unique string identifying this message
	Guid system.String `json:"guid"`
	// If this is a response, this is the message guid that we are responding to
	Request system.String `json:"request"`
}

// Automatically created basic rule for message
type Message_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for sourceRequest
type SourceRequest_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for sourceResponse
type SourceResponse_rule struct {
	*system.Object_base
	*system.Rule_base
}

// This is sent by the client to request a source file.
type SourceRequest struct {
	*system.Object_base
	*Message_base
	Name system.String `json:"name"`
}

// This is returned when the client requests a source.
type SourceResponse struct {
	*system.Object_base
	*Message_base
	Data  system.String `json:"data"`
	Found system.Bool   `json:"found"`
	Name  system.String `json:"name"`
}

func init() {
	json.Register("kego.io/editor/shared/messages", "$message", reflect.TypeOf(&Message_base{}), 0xaabb15d6475cfec1)
	json.Register("kego.io/editor/shared/messages", "@message", reflect.TypeOf(&Message_rule{}), 0xaabb15d6475cfec1)
	json.Register("kego.io/editor/shared/messages", "@sourceRequest", reflect.TypeOf(&SourceRequest_rule{}), 0xe7f2816c1a88470b)
	json.Register("kego.io/editor/shared/messages", "@sourceResponse", reflect.TypeOf(&SourceResponse_rule{}), 0x4421b9f7cf76b1b0)
	json.Register("kego.io/editor/shared/messages", "message", reflect.TypeOf((*Message)(nil)).Elem(), 0xaabb15d6475cfec1)
	json.Register("kego.io/editor/shared/messages", "sourceRequest", reflect.TypeOf(&SourceRequest{}), 0xe7f2816c1a88470b)
	json.Register("kego.io/editor/shared/messages", "sourceResponse", reflect.TypeOf(&SourceResponse{}), 0x4421b9f7cf76b1b0)
}
