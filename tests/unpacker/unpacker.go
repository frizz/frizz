package unpacker // import "kego.io/tests/unpacker"

// ke: {"package": {"notest": true}}

import (
	"context"

	"kego.io/system"
)

type Interface interface {
	Process(ctx context.Context, data []byte, i *interface{}) error
}

var Unpack = &Unpacker{}

type Unpacker struct{}

func (*Unpacker) Process(ctx context.Context, data []byte, i *interface{}) error {
	return system.Unmarshal(ctx, data, i)
}
