package system

import (
	"strconv"

	"kego.io/json"
)

type Number struct {
	Value  float64
	Exists bool
}

func NewNumber(n float64) Number {
	return Number{Value: n, Exists: true}
}

func (out *Number) UnmarshalJSON(in []byte, path string, imports map[string]string) error {
	var f *float64
	if err := json.UnmarshalPlain(in, &f, path, imports); err != nil {
		return err
	}
	if f == nil {
		out.Exists = false
		out.Value = 0.0
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
	return []byte(formatFloat(n.Value)), nil
}

func (n *Number) String() string {
	if !n.Exists {
		return ""
	}
	return formatFloat(n.Value)
}

func formatFloat(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
