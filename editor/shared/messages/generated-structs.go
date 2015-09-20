package messages

import (
	"reflect"

	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for globalRequest
type GlobalRequest_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for globalResponse
type GlobalResponse_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for message
type Message_rule struct {
	*system.Base
	*system.RuleBase
}

// Automatically created basic rule for messageBase
type MessageBase_rule struct {
	*system.Base
	*system.RuleBase
}

// This is sent by the client to request a global.
type GlobalRequest struct {
	*system.Base
	*MessageBase
	Name system.String `json:"name"`
}

// This is returned when the client requests a global.
type GlobalResponse struct {
	*system.Base
	*MessageBase
	Data  system.String `json:"data"`
	Found system.Bool   `json:"found"`
	Name  system.String `json:"name"`
}

// This is a base class embedded in each message
type MessageBase struct {
	// A globally unique string identifying this message
	Guid system.String `json:"guid"`
	// If this is a response, this is the message guid that we are responding to
	Request system.String `json:"request"`
}

func init() {
	json.Register("kego.io/editor/shared/messages", "@globalRequest", reflect.TypeOf(&GlobalRequest_rule{}), 0xb72faee9d9111cb1)
	json.Register("kego.io/editor/shared/messages", "@globalResponse", reflect.TypeOf(&GlobalResponse_rule{}), 0x12e106b47dc01e79)
	json.Register("kego.io/editor/shared/messages", "@message", reflect.TypeOf(&Message_rule{}), 0x6cbf132e7d248f63)
	json.Register("kego.io/editor/shared/messages", "@messageBase", reflect.TypeOf(&MessageBase_rule{}), 0x5a67e563d231a890)
	json.Register("kego.io/editor/shared/messages", "globalRequest", reflect.TypeOf(&GlobalRequest{}), 0xb72faee9d9111cb1)
	json.Register("kego.io/editor/shared/messages", "globalResponse", reflect.TypeOf(&GlobalResponse{}), 0x12e106b47dc01e79)
	json.Register("kego.io/editor/shared/messages", "messageBase", reflect.TypeOf(&MessageBase{}), 0x5a67e563d231a890)
}
