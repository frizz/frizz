package unpacker // import "kego.io/tests/unpacker"

// ke: {"package": {"notest": true}}

import (
	"bytes"

	"github.com/davelondon/kerr"
	"golang.org/x/net/context"
	"kego.io/json"
)

type Interface interface {
	Process(ctx context.Context, data []byte, i *interface{}) error
}

var Unmarshal = &Unmarshaler{}
var Unpack = &Unpacker{}
var Decode = &Decoder{}

type Unmarshaler struct{}
type Unpacker struct{}
type Decoder struct{}

func (*Decoder) Process(ctx context.Context, data []byte, i *interface{}) error {
	buf := bytes.NewBuffer(data)
	return json.NewDecoder(ctx, buf).Decode(i)
}
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
