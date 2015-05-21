package system

import (
	"strconv"

	"math"

	"kego.io/json"
	"kego.io/uerr"
)

type Int struct {
	Value  int
	Exists bool
}

func NewInt(n int) Int {
	return Int{Value: n, Exists: true}
}

func (out *Int) UnmarshalJSON(in []byte, path string, imports map[string]string) error {
	var f *float64
	if err := json.UnmarshalPlain(in, &f, path, imports); err != nil {
		return uerr.New("WCXYWVMOTT", err, "Int.UnmarshalJSON", "json.UnmarshalPlain")
	}
	if f == nil {
		out.Exists = false
		out.Value = 0
	} else {
		i := math.Floor(*f)
		if i != *f {
			return uerr.New("KVEOETSIJY", nil, "Int.UnmarshalJSON", "%v is not an integer", *f)
		}
		out.Exists = true
		out.Value = int(i)
	}
	return nil
}

func (n *Int) MarshalJSON() ([]byte, error) {
	if !n.Exists {
		return []byte("null"), nil
	}
	return []byte(formatInt(n.Value)), nil
}

func (n *Int) String() string {
	if !n.Exists {
		return ""
	}
	return formatInt(n.Value)
}

func formatInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}
