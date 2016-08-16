package repacker // import "kego.io/tests/repacker"

// ke: {"package": {"notest": true}}

import (
	"context"

	"github.com/davelondon/kerr"
	"kego.io/json"
	"kego.io/system/node"
)

var Repack = &Repacker{}

type Repacker struct{}

func (*Repacker) Process(ctx context.Context, data []byte, i *interface{}) error {
	n, err := node.Unmarshal(ctx, data)
	if err != nil {
		return kerr.Wrap("RXAARYNHLP", err)
	}
	return json.Unpack(ctx, node.Pack(n), i)
}
