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

// Automatically created basic rule for globalRequest
type GlobalRequest_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for globalResponse
type GlobalResponse_rule struct {
	*system.Object_base
	*system.Rule_base
}

// Automatically created basic rule for message
type Message_rule struct {
	*system.Object_base
	*system.Rule_base
}

// This is sent by the client to request a global.
type GlobalRequest struct {
	*system.Object_base
	*Message_base
	Name system.String `json:"name"`
}

// This is returned when the client requests a global.
type GlobalResponse struct {
	*system.Object_base
	*Message_base
	Data  system.String `json:"data"`
	Found system.Bool   `json:"found"`
	Name  system.String `json:"name"`
}

func init() {
	json.Register("kego.io/editor/shared/messages", "$message", reflect.TypeOf(&Message_base{}), 0xaabb15d6475cfec1)
	json.Register("kego.io/editor/shared/messages", "@globalRequest", reflect.TypeOf(&GlobalRequest_rule{}), 0x136af35684ddacf7)
	json.Register("kego.io/editor/shared/messages", "@globalResponse", reflect.TypeOf(&GlobalResponse_rule{}), 0x6b2590de8a9ff1d6)
	json.Register("kego.io/editor/shared/messages", "@message", reflect.TypeOf(&Message_rule{}), 0xaabb15d6475cfec1)
	json.Register("kego.io/editor/shared/messages", "globalRequest", reflect.TypeOf(&GlobalRequest{}), 0x136af35684ddacf7)
	json.Register("kego.io/editor/shared/messages", "globalResponse", reflect.TypeOf(&GlobalResponse{}), 0x6b2590de8a9ff1d6)
}
