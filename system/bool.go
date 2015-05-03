package system

import (
	"fmt"

	"kego.io/json"
)

type Bool struct {
	Value  bool
	Exists bool
}

func NewBool(b bool) Bool {
	return Bool{Value: b, Exists: true}
}

func (out *Bool) UnmarshalJSON(in []byte, path string, imports map[string]string) error {
	var b *bool
	if err := json.UnmarshalPlain(in, &b, path, imports); err != nil {
		return err
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

func (b *Bool) MarshalJSON() ([]byte, error) {
	if !b.Exists {
		return []byte("null"), nil
	}
	return []byte(formatBool(b.Value)), nil
}

func (b *Bool) String() string {
	if !b.Exists {
		return ""
	}
	return formatBool(b.Value)
}

func formatBool(b bool) string {
	return fmt.Sprintf("%v", b)
}
