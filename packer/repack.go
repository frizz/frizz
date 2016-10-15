package packer

import (
	"encoding/json"

	"context"

	"github.com/davelondon/kerr"
)

type repackStruct struct{}

type Repacker interface {
	// Repack packs the object into json data ready for marshaling to a string.
	// If iface is true, we are packing into an interface type, so map types are
	// always specified in specific typed form: {"type": "mytype", "value":
	// {"foo": "bar"}}
	// Repack(ctx context.Context, iface bool) (interface{}, error)
	Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, err error)
}

func Marshal(ctx context.Context, in interface{}) ([]byte, error) {
	i, err := Repack(ctx, in)
	if err != nil {
		return nil, kerr.Wrap("SBWVUFLWMV", err)
	}
	b, err := json.Marshal(i)
	if err != nil {
		return nil, kerr.Wrap("NNBPQBDWEW", err)
	}
	return b, nil
}

func Repack(ctx context.Context, in interface{}) (interface{}, error) {
	// TODO: this!
	return nil, nil
}
