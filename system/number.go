package system

import (
	"strconv"

	"kego.io/json"
)

type Number struct {
	Value  float64
	Exists bool
}

func (out *Number) UnmarshalJSON(in []byte, path string, imports map[string]string) error {
	var f *float64
	if err := json.Unmarshal(in, &f, path, imports); err != nil {
		return err
	}
	if f == nil {
		out.Exists = false
	} else {
		out.Exists = true
		out.Value = *f
	}
	return nil
}

func (n *Number) MarshalJSON() ([]byte, error) {
	if !n.Exists {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatFloat(n.Value, 'f', -1, 64)), nil
}
