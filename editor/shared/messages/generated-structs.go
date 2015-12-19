package messages

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for message
type MessageRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for sourceRequest
type SourceRequestRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for sourceResponse
type SourceResponseRule struct {
	*system.Object
	*system.Rule
}

// This is a base class embedded in each message
type Message struct {
	// A globally unique string identifying this message
	Guid *system.String `json:"guid"`
	// If this is a response, this is the message guid that we are responding to
	Request *system.String `json:"request"`
}
type MessageInterface interface {
	GetMessage(ctx context.Context) *Message
}

func (o *Message) GetMessage(ctx context.Context) *Message {
	return o
}

// This is sent by the client to request a source file.
type SourceRequest struct {
	*system.Object
	*Message
	Name *system.String `json:"name"`
}
type SourceRequestInterface interface {
	GetSourceRequest(ctx context.Context) *SourceRequest
}

func (o *SourceRequest) GetSourceRequest(ctx context.Context) *SourceRequest {
	return o
}

// This is returned when the client requests a source.
type SourceResponse struct {
	*system.Object
	*Message
	Data  *system.String `json:"data"`
	Found *system.Bool   `json:"found"`
	Name  *system.String `json:"name"`
}
type SourceResponseInterface interface {
	GetSourceResponse(ctx context.Context) *SourceResponse
}

func (o *SourceResponse) GetSourceResponse(ctx context.Context) *SourceResponse {
	return o
}
func init() {
	json.Register("kego.io/editor/shared/messages", "@message", reflect.TypeOf((*MessageRule)(nil)), nil, 4785498943088766147)
	json.Register("kego.io/editor/shared/messages", "@sourceRequest", reflect.TypeOf((*SourceRequestRule)(nil)), nil, 6522969464666442267)
	json.Register("kego.io/editor/shared/messages", "@sourceResponse", reflect.TypeOf((*SourceResponseRule)(nil)), nil, 14174040136885668420)
	json.Register("kego.io/editor/shared/messages", "message", reflect.TypeOf((*Message)(nil)), reflect.TypeOf((*MessageInterface)(nil)).Elem(), 4785498943088766147)
	json.Register("kego.io/editor/shared/messages", "sourceRequest", reflect.TypeOf((*SourceRequest)(nil)), reflect.TypeOf((*SourceRequestInterface)(nil)).Elem(), 6522969464666442267)
	json.Register("kego.io/editor/shared/messages", "sourceResponse", reflect.TypeOf((*SourceResponse)(nil)), reflect.TypeOf((*SourceResponseInterface)(nil)).Elem(), 14174040136885668420)
}
