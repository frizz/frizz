package system

import (
	"context"

	"github.com/davelondon/kerr"
	"kego.io/json"
)

func NewBool(b bool) *Bool {
	out := Bool(b)
	return &out
}

func (b *Bool) Value() bool {
	return bool(*b)
}

func (out *Bool) Unpack(ctx context.Context, in json.Packed) error {
	if in == nil || in.Type() == json.J_NULL {
		return kerr.New("FXCQGNYKIJ", "Called Bool.Unpack with nil value")
	}
	if in.Type() == json.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != json.J_BOOL {
		return kerr.New("GXQGNEPJYS", "Can't unpack %s into *system.Bool", in.Type())
	}
	*out = Bool(in.Bool())
	return nil
}

var _ json.Unpacker = (*Bool)(nil)

func (b *Bool) MarshalJSON(ctx context.Context) ([]byte, error) {
	if b == nil {
		return []byte("null"), nil
	}
	return []byte(formatBool(b)), nil
}

var _ json.Marshaler = (*Bool)(nil)

func (b *Bool) String() string {
	if b == nil {
		return ""
	}
	return formatBool(b)
}

func formatBool(b *Bool) string {
	if *b {
		return "true"
	}
	return "false"
}

type NativeBool interface {
	NativeBool() bool
}

func (b Bool) NativeBool() bool {
	return bool(b)
}

var _ NativeBool = (*Bool)(nil)

func (r *BoolRule) GetDefault() interface{} {
	return r.Default
}

var _ DefaultRule = (*BoolRule)(nil)
