package system

import (
	"strconv"

	"kego.io/json"
)

type String struct {
	Value  string
	Exists bool
}

func (out *String) UnmarshalJSON(in []byte, context *json.Context) error {
	var s *string
	if err := json.Unmarshal(in, &s, context); err != nil {
		return err
	}
	if s == nil {
		out.Exists = false
	} else {
		out.Exists = true
		out.Value = *s
	}
	return nil
}

func (s *String) MarshalJSON() ([]byte, error) {
	if !s.Exists {
		return []byte("null"), nil
	}
	return []byte(strconv.Quote(s.Value)), nil
}

func (r *String_rule) HasDefault() bool {
	return r.Default.Exists
}

func (r *String_rule) GetDefault() ([]byte, error) {
	return json.Marshal(&r.Default)
}
