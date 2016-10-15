package system

import (
	"strconv"

	"context"

	"fmt"

	"math"

	"github.com/davelondon/kerr"
	"kego.io/packer"
)

func NewNumber(n float64) *Number {
	out := Number(n)
	return &out
}

func (n *Number) Value() float64 {
	return float64(*n)
}

func (r *NumberRule) Enforce(ctx context.Context, data interface{}) (fail bool, messages []string, err error) {

	if i, ok := data.(NumberInterface); ok {
		data = i.GetNumber(ctx)
	}

	n, ok := data.(*Number)
	if !ok && data != nil {
		return true, nil, kerr.New("FUGYGJVHYS", "Data %T should be *system.Number", data)
	}

	// Maximum provides an upper bound for the restriction If ExclusiveMaximum is true, the value
	// must be less than maximum. If false or not provided, the value must be less than or equal
	// to the maximum.
	if r.Maximum != nil {
		if n == nil && !r.Optional {
			fail = true
			messages = append(messages, "Maximum: value must exist")
		}
		if n != nil {
			if r.ExclusiveMaximum {
				if n.Value() >= r.Maximum.Value() {
					fail = true
					messages = append(messages, fmt.Sprintf("Maximum (exclusive): value %v must be less than %v", n, r.Maximum.Value()))
				}
			} else {
				if n.Value() > r.Maximum.Value() {
					fail = true
					messages = append(messages, fmt.Sprintf("Maximum: value %v must not be greater than %v", n, r.Maximum.Value()))
				}
			}
		}
	}

	// Minimum provides a lower bound for the restriction. If ExclusiveMinimum is true, the value
	// must be greater than minimum. If false or not provided, the value must be greater than or
	// equal to the minimum.
	if r.Minimum != nil {
		if n == nil && !r.Optional {
			fail = true
			messages = append(messages, "Minimum: value must exist")
		}
		if n != nil {
			if r.ExclusiveMinimum {
				if n.Value() <= r.Minimum.Value() {
					fail = true
					messages = append(messages, fmt.Sprintf("Minimum (exclusive): value %v must be greater than %v", n, r.Minimum.Value()))
				}
			} else {
				if n.Value() < r.Minimum.Value() {
					fail = true
					messages = append(messages, fmt.Sprintf("Minimum: value %v must not be less than %v", n, r.Minimum.Value()))
				}
			}
		}
	}

	// This restricts the number to be a multiple of the given number
	// MultipleOf Number
	if r.MultipleOf != nil {
		if n == nil && !r.Optional {
			fail = true
			messages = append(messages, "MultipleOf: value must exist")
		}
		if n != nil {
			_, frac := math.Modf(n.Value() / r.MultipleOf.Value())
			if frac != 0 {
				fail = true
				messages = append(messages, fmt.Sprintf("MultipleOf: value %v must be a multiple of %v", n, r.MultipleOf.Value()))
			}
		}
	}

	return
}

var _ Enforcer = (*NumberRule)(nil)

var _ packer.Unpacker = (*Number)(nil)

//var _ packer.Repacker = (*Number)(nil)

func (n *Number) String() string {
	if n == nil {
		return ""
	}
	return formatFloat(n)
}

func formatFloat(f *Number) string {
	return strconv.FormatFloat(f.Value(), 'f', -1, 64)
}

type NativeNumber interface {
	NativeNumber() float64
}

func (n Number) NativeNumber() float64 {
	return float64(n)
}

var _ NativeNumber = (*Number)(nil)

func (r *NumberRule) GetDefault() interface{} {
	if r.Default == nil {
		return nil
	}
	return r.Default
}

var _ DefaultRule = (*NumberRule)(nil)
