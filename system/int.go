package system

import (
	"fmt"
	"strconv"

	"math"

	"kego.io/json"
	"kego.io/kerr"
)

type Int struct {
	Value  int
	Exists bool
}

func NewInt(n int) Int {
	return Int{Value: n, Exists: true}
}

func (r *Int_rule) Enforce(data interface{}, path string, aliases map[string]string) (bool, string, error) {

	i, ok := data.(Int)
	if !ok {
		return false, "", kerr.New("AISBHNCJXJ", nil, "Data %T should be Int", data)
	}

	// If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.
	// ExclusiveMaximum bool
	// This provides an upper bound for the restriction
	// Maximum Number
	if r.Maximum.Exists {
		if !i.Exists && !r.Optional {
			return false, "Maximum: value must exist", nil
		}
		if i.Exists && i.Value > r.Maximum.Value {
			return false, fmt.Sprintf("Maximum: value %v must not be greater than %v", i.Value, r.Maximum.Value), nil
		}
	}

	// If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.
	// ExclusiveMinimum bool
	// This provides a lower bound for the restriction
	// Minimum Number
	if r.Minimum.Exists {
		if !i.Exists && !r.Optional {
			return false, "Minimum: value must exist", nil
		}
		if i.Exists && i.Value < r.Minimum.Value {
			return false, fmt.Sprintf("Minimum: value %v must not be less than %v", i.Value, r.Minimum.Value), nil
		}
	}

	// This restricts the number to be a multiple of the given number
	// MultipleOf Number
	if r.MultipleOf.Exists {
		if !i.Exists && !r.Optional {
			return false, "MultipleOf: value must exist", nil
		}
		if i.Exists && i.Value%r.MultipleOf.Value != 0 {
			return false, fmt.Sprintf("MultipleOf: value %v must be a multiple of %v", i.Value, r.MultipleOf.Value), nil
		}
	}

	return true, "", nil
}

func (out *Int) Unpack(in json.Unpackable) error {
	if in == nil || in.UpType() == json.J_NULL {
		out.Exists = false
		out.Value = 0
		return nil
	}
	if in.UpType() != json.J_NUMBER {
		return kerr.New("UJUBDGVYGF", nil, "Can't unpack %s into system.Int", in.UpType())
	}
	i := math.Floor(in.UpNumber())
	if i != in.UpNumber() {
		return kerr.New("KVEOETSIJY", nil, "%v is not an integer", in.UpNumber())
	}
	out.Exists = true
	out.Value = int(i)
	return nil
}

var _ json.Unpacker = (*Int)(nil)

func (n Int) MarshalJSON() ([]byte, error) {
	if !n.Exists {
		return []byte("null"), nil
	}
	return []byte(formatInt(n.Value)), nil
}

var _ json.Marshaler = (*Int)(nil)

func (n *Int) String() string {
	if !n.Exists {
		return ""
	}
	return formatInt(n.Value)
}

func formatInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int) NativeNumber() (value float64, exists bool) {
	return float64(i.Value), i.Exists
}

// We satisfy the json.EmptyAware interface to allow intelligent omission of
// empty values when marshalling
func (i Int) Empty() bool {
	return !i.Exists
}

var _ json.EmptyAware = (*Int)(nil)
