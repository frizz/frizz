// info:{"Path":"kego.io/editor/shared/messages","Hash":9398310788482718939}
package messages

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/json"
	"kego.io/system"
)

// Automatically created basic rule for dataRequest
type DataRequestRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for dataResponse
type DataResponseRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for message
type MessageRule struct {
	*system.Object
	*system.Rule
}

// This is sent by the client to request a source file.
type DataRequest struct {
	*system.Object
	*Message
	File    *system.String `json:"file"`
	Name    *system.String `json:"name"`
	Package *system.String `json:"package"`
}
type DataRequestInterface interface {
	GetDataRequest(ctx context.Context) *DataRequest
}

func (o *DataRequest) GetDataRequest(ctx context.Context) *DataRequest {
	return o
}

// This is returned when the client requests a source.
type DataResponse struct {
	*system.Object
	*Message
	Data    *system.String `json:"data"`
	Found   *system.Bool   `json:"found"`
	Name    *system.String `json:"name"`
	Package *system.String `json:"package"`
}
type DataResponseInterface interface {
	GetDataResponse(ctx context.Context) *DataResponse
}

func (o *DataResponse) GetDataResponse(ctx context.Context) *DataResponse {
	return o
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
func init() {
	json.RegisterPackage("kego.io/editor/shared/messages", 9398310788482718939)
	json.RegisterType("kego.io/editor/shared/messages", "dataRequest", reflect.TypeOf((*DataRequest)(nil)), reflect.TypeOf((*DataRequestRule)(nil)), reflect.TypeOf((*DataRequestInterface)(nil)).Elem())
	json.RegisterType("kego.io/editor/shared/messages", "dataResponse", reflect.TypeOf((*DataResponse)(nil)), reflect.TypeOf((*DataResponseRule)(nil)), reflect.TypeOf((*DataResponseInterface)(nil)).Elem())
	json.RegisterType("kego.io/editor/shared/messages", "message", reflect.TypeOf((*Message)(nil)), reflect.TypeOf((*MessageRule)(nil)), reflect.TypeOf((*MessageInterface)(nil)).Elem())
}
