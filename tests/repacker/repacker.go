package repacker // import "frizz.io/tests/repacker"

// notest

import (
	"context"

	"frizz.io/system"
	"frizz.io/system/node"
	"github.com/dave/kerr"
)

var Repack = &Repacker{}

type Repacker struct{}

func (*Repacker) Process(ctx context.Context, data []byte, i *interface{}) error {
	n, err := node.Unmarshal(ctx, data)
	if err != nil {
		return kerr.Wrap("RXAARYNHLP", err)
	}
	return system.Unpack(ctx, node.Pack(n), i)
}
