package system

import (
	"context"

	"kego.io/packer"
)

func NewBool(b bool) *Bool {
	out := Bool(b)
	return &out
}

func (b *Bool) Value() bool {
	return bool(*b)
}

var _ packer.Unpacker = (*Bool)(nil)

func (b *Bool) Repack(ctx context.Context) (interface{}, error) {
	if b == nil {
		return nil, nil
	}
	return formatBool(b), nil
}

var _ packer.Repacker = (*Bool)(nil)

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
