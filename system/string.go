package system

import (
	"strconv"

	"kego.io/json"
	"kego.io/kerr"
)

type String struct {
	Value  string
	Exists bool
}

func NewString(s string) String {
	return String{Value: s, Exists: true}
}

func (out *String) UnmarshalJSON(in []byte, path string, imports map[string]string) error {
	var s *string
	if err := json.UnmarshalPlain(in, &s, path, imports); err != nil {
		return kerr.New("ACHMRKVFAB", err, "String.UnmarshalJSON", "json.UnmarshalPlain")
	}
	if s == nil {
		out.Exists = false
		out.Value = ""
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

func (s *String) String() string {
	if !s.Exists {
		return ""
	}
	return s.Value
}
