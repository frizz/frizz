package system

import (
	"fmt"

	"kego.io/json"
	"kego.io/kerr"
)

type Bool struct {
	Value  bool
	Exists bool
}

func NewBool(b bool) Bool {
	return Bool{Value: b, Exists: true}
}

func (out *Bool) UnmarshalJSON(in []byte, path string, aliases map[string]string) error {
	var b *bool
	if err := json.UnmarshalPlain(in, &b); err != nil {
		return kerr.New("CJMOICJGJG", err, "json.UnmarshalPlain")
	}
	if b == nil {
		out.Exists = false
		out.Value = false
	} else {
		out.Exists = true
		out.Value = *b
	}
	return nil
}

var _ json.Unmarshaler = (*Bool)(nil)

func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.Exists {
		return []byte("null"), nil
	}
	return []byte(formatBool(b.Value)), nil
}

var _ json.Marshaler = (*Bool)(nil)

func (b *Bool) String() string {
	if !b.Exists {
		return ""
	}
	return formatBool(b.Value)
}

func formatBool(b bool) string {
	return fmt.Sprintf("%v", b)
}

type NativeBool interface {
	NativeBool() (bool, bool)
}

func (b Bool) NativeBool() (value bool, exists bool) {
	return b.Value, b.Exists
}

// We satisfy the json.EmptyAware interface to allow intelligent omission of
// empty values when marshalling
func (b Bool) Empty() bool {
	return !b.Exists
}

var _ json.EmptyAware = (*Bool)(nil)
