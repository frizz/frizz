package system

import (
	"fmt"
	"strconv"

	"context"

	"math"

	"github.com/davelondon/kerr"
	"kego.io/json"
)

type Int int

func NewInt(i int) *Int {
	out := Int(i)
	return &out
}

func (i *Int) GetString(ctx context.Context) *String {
	return NewString(i.String())
}

func (i *Int) Value() int {
	return int(*i)
}

func (r *IntRule) Enforce(ctx context.Context, data interface{}) (fail bool, messages []string, err error) {

	if i, ok := data.(IntInterface); ok {
		data = i.GetInt(ctx)
	}

	i, ok := data.(*Int)
	if !ok && data != nil {
		return true, nil, kerr.New("AISBHNCJXJ", "Data %T should be *system.Int", data)
	}

	// This provides an upper bound for the restriction
	if r.Maximum != nil {
		if i == nil && !r.Optional {
			fail = true
			messages = append(messages, "Maximum: value must exist")
		}
		if i != nil && i.Value() > r.Maximum.Value() {
			fail = true
			messages = append(messages, fmt.Sprintf("Maximum: value %v must not be greater than %v", i, r.Maximum.Value()))
		}
	}

	// This provides a lower bound for the restriction
	if r.Minimum != nil {
		if i == nil && !r.Optional {
			fail = true
			messages = append(messages, "Minimum: value must exist")
		}
		if i != nil && i.Value() < r.Minimum.Value() {
			fail = true
			messages = append(messages, fmt.Sprintf("Minimum: value %v must not be less than %v", i, r.Minimum.Value()))
		}
	}

	// This restricts the number to be a multiple of the given number
	if r.MultipleOf != nil {
		if i == nil && !r.Optional {
			fail = true
			messages = append(messages, "MultipleOf: value must exist")
		}
		if i != nil && i.Value()%r.MultipleOf.Value() != 0 {
			fail = true
			messages = append(messages, fmt.Sprintf("MultipleOf: value %v must be a multiple of %v", i, r.MultipleOf.Value()))
		}
	}
	return
}

var _ Enforcer = (*IntRule)(nil)

func (out *Int) Unpack(ctx context.Context, in json.Packed) error {
	if in == nil || in.Type() == json.J_NULL {
		return kerr.New("JEJANRWFMH", "Called Int.Unpack with nil value")
	}
	if in.Type() == json.J_MAP {
		in = in.Map()["value"]
	}
	if in.Type() != json.J_NUMBER {
		return kerr.New("UJUBDGVYGF", "Can't unpack %s into *system.Int", in.Type())
	}
	i := math.Floor(in.Number())
	if i != in.Number() {
		return kerr.New("KVEOETSIJY", "%v is not an integer", in.Number())
	}
	*out = Int(int(i))
	return nil
}

var _ json.Unpacker = (*Int)(nil)

func (n *Int) MarshalJSON(ctx context.Context) ([]byte, error) {
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
