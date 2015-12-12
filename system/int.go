package system

import (
	"fmt"
	"strconv"

	"math"

	"kego.io/json"
	"kego.io/kerr"
)

type Int int

func NewInt(i int) *Int {
	out := Int(i)
	return &out
}

func (i *Int) GetString() *String {
	return NewString(i.String())
}

func (i *Int) Value() int {
	return int(*i)
}

func (i *Int) Set(in int) {
	*i = Int(in)
}

func (r *IntRule) Enforce(data interface{}, path string, aliases map[string]string) (bool, string, error) {

	i, ok := data.(*Int)
	if !ok && data != nil {
		return false, "", kerr.New("AISBHNCJXJ", nil, "Data %T should be *system.Int", data)
	}

	// If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.
	// ExclusiveMaximum bool
	// This provides an upper bound for the restriction
	// Maximum Number
	if r.Maximum != nil {
		if i == nil && !r.Optional {
			return false, "Maximum: value must exist", nil
		}
		if i != nil && i.Value() > r.Maximum.Value() {
			return false, fmt.Sprintf("Maximum: value %v must not be greater than %v", i, r.Maximum.Value()), nil
		}
	}

	// If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.
	// ExclusiveMinimum bool
	// This provides a lower bound for the restriction
	// Minimum Number
	if r.Minimum != nil {
		if i == nil && !r.Optional {
			return false, "Minimum: value must exist", nil
		}
		if i != nil && i.Value() < r.Minimum.Value() {
			return false, fmt.Sprintf("Minimum: value %v must not be less than %v", i, r.Minimum.Value()), nil
		}
	}

	// This restricts the number to be a multiple of the given number
	// MultipleOf Number
	if r.MultipleOf != nil {
		if i == nil && !r.Optional {
			return false, "MultipleOf: value must exist", nil
		}
		if i != nil && i.Value()%r.MultipleOf.Value() != 0 {
			return false, fmt.Sprintf("MultipleOf: value %v must be a multiple of %v", i, r.MultipleOf.Value()), nil
		}
	}

	return true, "", nil
}

var _ Enforcer = (*IntRule)(nil)

func (out *Int) Unpack(in json.Packed) error {
	if in == nil || in.Type() == json.J_NULL {
		return kerr.New("JEJANRWFMH", nil, "Called Int.Unpack with nil value")
	}
	if in.Type() != json.J_NUMBER {
		return kerr.New("UJUBDGVYGF", nil, "Can't unpack %s into *system.Int", in.Type())
	}
	i := math.Floor(in.Number())
	if i != in.Number() {
		return kerr.New("KVEOETSIJY", nil, "%v is not an integer", in.Number())
	}
	*out = Int(int(i))
	return nil
}

var _ json.Unpacker = (*Int)(nil)

func (n *Int) MarshalJSON() ([]byte, error) {
	if n == nil {
		return []byte("null"), nil
	}
	return []byte(formatInt(n)), nil
}

var _ json.Marshaler = (*Int)(nil)

func (n *Int) String() string {
	if n == nil {
		return ""
	}
	return formatInt(n)
}

func formatInt(i *Int) string {
	return strconv.FormatInt(int64(*i), 10)
}

func (i Int) NativeNumber() float64 {
	return float64(i)
}

var _ NativeNumber = (*Int)(nil)

func (r *IntRule) GetDefault() interface{} {
	return r.Default
}

var _ DefaultRule = (*IntRule)(nil)
