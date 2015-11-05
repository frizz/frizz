package system

import (
	"kego.io/json"
	"kego.io/kerr"
)

type Bool bool

func NewBool(b bool) *Bool {
	out := Bool(b)
	return &out
}

func (b *Bool) Value() bool {
	return bool(*b)
}

func (b *Bool) Set(in bool) {
	*b = Bool(in)
}

func (out *Bool) Unpack(in json.Unpackable) error {
	if in == nil || in.UpType() == json.J_NULL {
		return kerr.New("FXCQGNYKIJ", nil, "Called Bool.Unpack with nil value")
	}
	if in.UpType() != json.J_BOOL {
		return kerr.New("GXQGNEPJYS", nil, "Can't unpack %s into *system.Bool", in.UpType())
	}
	*out = Bool(in.UpBool())
	return nil
}

var _ json.Unpacker = (*Bool)(nil)

func (b *Bool) MarshalJSON() ([]byte, error) {
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
