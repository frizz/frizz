package unpacker // import "kego.io/process/tests/unpacker"

// ke: {"package": {"notest": true}}

import (
	"golang.org/x/net/context"
	"kego.io/json"
	"kego.io/kerr"
)

type Interface interface {
	Process(ctx context.Context, data []byte, i *interface{}) error
}

var Unmarshal = &Unmarshaler{}
var Unpack = &Unpacker{}

type Unmarshaler struct{}
type Unpacker struct{}

func (*Unmarshaler) Process(ctx context.Context, data []byte, i *interface{}) error {
	return json.Unmarshal(ctx, data, i)
}
func (*Unpacker) Process(ctx context.Context, data []byte, i *interface{}) error {
	var j interface{}
	if err := json.UnmarshalPlain(data, &j); err != nil {
		return kerr.Wrap("PJJYUVNXED", err)
	}
	return json.Unpack(ctx, json.Pack(j), i)
}
