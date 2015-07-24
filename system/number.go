package system

import (
	"strconv"

	"fmt"

	"math"

	"kego.io/json"
	"kego.io/kerr"
)

type Number struct {
	Value  float64
	Exists bool
}

func NewNumber(n float64) Number {
	return Number{Value: n, Exists: true}
}

func (r *Number_rule) Enforce(data interface{}, path string, aliases map[string]string) (bool, string, error) {

	n, ok := data.(Number)
	if !ok {
		return false, "", kerr.New("AISBHNCJXJ", nil, "Data %T should be Number", data)
	}

	// If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.
	// ExclusiveMaximum bool
	// This provides an upper bound for the restriction
	// Maximum Number
	if r.Maximum.Exists {
		if !n.Exists && !r.Optional {
			return false, "Maximum: value must exist", nil
		}
		if n.Exists {
			if r.ExclusiveMaximum {
				if n.Value >= r.Maximum.Value {
					return false, fmt.Sprintf("Maximum (exclusive): value %v must be less than %v", n.Value, r.Maximum.Value), nil
				}
			} else {
				if n.Value > r.Maximum.Value {
					return false, fmt.Sprintf("Maximum: value %v must not be greater than %v", n.Value, r.Maximum.Value), nil
				}
			}
		}
	}

	// If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.
	// ExclusiveMinimum bool
	// This provides a lower bound for the restriction
	// Minimum Number
	if r.Minimum.Exists {
		if !n.Exists && !r.Optional {
			return false, "Minimum: value must exist", nil
		}
		if n.Exists {
			if r.ExclusiveMinimum {
				if n.Value <= r.Minimum.Value {
					return false, fmt.Sprintf("Minimum (exclusive): value %v must be greater than %v", n.Value, r.Minimum.Value), nil
				}
			} else {
				if n.Value < r.Minimum.Value {
					return false, fmt.Sprintf("Minimum: value %v must not be less than %v", n.Value, r.Minimum.Value), nil
				}
			}
		}
	}

	// This restricts the number to be a multiple of the given number
	// MultipleOf Number
	if r.MultipleOf.Exists {
		if !n.Exists && !r.Optional {
			return false, "MultipleOf: value must exist", nil
		}
		if n.Exists {
			_, frac := math.Modf(n.Value / r.MultipleOf.Value)
			if frac != 0 {
				return false, fmt.Sprintf("MultipleOf: value %v must be a multiple of %v", n.Value, r.MultipleOf.Value), nil
			}
		}
	}

	return true, "", nil
}

func (out *Number) UnmarshalJSON(in []byte, path string, aliases map[string]string) error {
	var f *float64
	if err := json.UnmarshalPlain(in, &f, path, aliases); err != nil {
		return kerr.New("GXNBRBELFA", err, "json.UnmarshalPlain")
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

type NativeNumber interface {
	NativeNumber() (float64, bool)
}

func (n Number) NativeNumber() (value float64, exists bool) {
	return n.Value, n.Exists
}
