package system

import (
	"fmt"

	"kego.io/json"
)

type Bool struct {
	Value  bool
	Exists bool
}

func (out *Bool) UnmarshalJSON(in []byte, context *json.Context) error {
	var b *bool
	if err := json.Unmarshal(in, &b, context); err != nil {
		return err
	}
	if b == nil {
		out.Exists = false
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
	return []byte(fmt.Sprintf("%v", b.Value)), nil
}

func (r *Bool_rule) HasDefault() bool {
	return r.Default.Exists
}

func (r *Bool_rule) GetDefault() ([]byte, error) {
	return json.Marshal(&r.Default)
}
