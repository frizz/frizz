package system

import (
	"strconv"

	"kego.io/json"
)

type String struct {
	Value  string
	Exists bool
}

func NewString(s string) String {
	return String{Value: s, Exists: true}
}
func NewStringMap(m map[string]string) map[string]String {
	var o = make(map[string]String, len(m))
	for name, value := range m {
		o[name] = NewString(value)
	}
	return o
}

func (out *String) UnmarshalJSON(in []byte, path string, imports map[string]string) error {
	var s *string
	if err := json.UnmarshalPlain(in, &s, path, imports); err != nil {
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
