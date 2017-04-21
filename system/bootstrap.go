package system

import (
	"context"

	"github.com/dave/kerr"
)

type Mappy map[string]string

type MappyInterface interface {
	GetMappy(context.Context) error
}

func UnpackMappyInterface(ctx context.Context, in Packed) (MappyInterface, error) {
	// Mappy is an map type, so we only accept a typed map
	switch in.Type() {
	case J_MAP:
		ob, err := UnpackUnknownType(ctx, in, true, "frizz.io/system", "mappy")
		if err != nil {
			return nil, kerr.Wrap("QSTEIBUNWO", err)
		}
		i, ok := ob.(MappyInterface)
		if !ok {
			return nil, kerr.New("SMJXFKMKSP", "%T does not implement system.MappyInterface", ob)
		}
		return i, nil
	default:
		return nil, kerr.New("VYHBOJFSIM", "Unpacking into a MappyInterface, so input must be a map. Found %s", in.Type())
	}
}

func (v *Mappy) Unpack(ctx context.Context, in Packed, iface bool) error {

	if in == nil || in.Type() == J_NULL {
		return nil
	}

	if iface {
		if in.Type() != J_MAP {
			return kerr.New("DTLQPHOJDT", "Mappy.Unpack: %s must by J_MAP", in.Type())
		}
		in = in.Map()["value"]
	}

	if in.Type() != J_MAP {
		return kerr.New("KYPSXYBLNC", "Mappy.Unpack: %s must by J_MAP", in.Type())
	}

	for key, value := range in.Map() {
		ob, err := UnpackString(ctx, value)
		if err != nil {
			return kerr.Wrap("HOLNALWYBA", err)
		}
		(*v)[key] = ob
	}

	return nil
}
