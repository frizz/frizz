package system

import (
	"fmt"
	"regexp"
	"strconv"

	"context"

	"github.com/davelondon/kerr"
	"kego.io/json"
)

type String string

func NewString(s string) *String {
	out := String(s)
	return &out
}

func (s *String) Value() string {
	return string(*s)
}

func (s *String) Set(in string) {
	*s = String(in)
}

func (r *StringRule) Validate(ctx context.Context) (fail bool, messages []string, err error) {
	if r.MaxLength != nil && r.MinLength != nil {
		if r.MaxLength.Value() < r.MinLength.Value() {
			fail = true
			messages = append(messages, fmt.Sprintf("MaxLength %d must not be less than MinLength %d", r.MaxLength.Value(), r.MinLength.Value()))
		}
	}
	if r.Pattern != nil {
		if _, err := regexp.Compile(r.Pattern.Value()); err != nil {
			fail = true
			messages = append(messages, fmt.Sprintf("Pattern: regex does not compile: %s", r.Pattern.Value()))
		}
	}
	return
}

var _ Validator = (*StringRule)(nil)

func (r *StringRule) Enforce(ctx context.Context, data interface{}) (fail bool, messages []string, err error) {

	if i, ok := data.(StringInterface); ok && i != nil {
		data = i.GetString(ctx)
	}

	s, ok := data.(*String)
	if !ok && data != nil {
		return true, nil, kerr.New("SXFBXGQSEA", "String rule: value %T should be *system.String", data)
	}

	// TODO: This restricts the value to one of several built-in formats
	// TODO: Format String
	if r.Format != nil {
		panic("TODO: Format rule implementation")
	}

	// This is a regex to match the value to
	// Pattern String
	if r.Pattern != nil {
		if s == nil && !r.Optional {
			fail = true
			messages = append(messages, "Pattern: value must exist")
		}
		if s != nil {
			reg, err := regexp.Compile(r.Pattern.Value())
			if err != nil {
				fail = true
				messages = append(messages, fmt.Sprintf("Pattern: regex does not compile: %s", r.Pattern.Value()))
			} else if !reg.Match([]byte(s.Value())) {
				fail = true
				messages = append(messages, fmt.Sprintf("Pattern: value must match %s", r.Pattern.Value()))
			}
		}
	}

	// This is a string that the value must match
	// Equal String
	if r.Equal != nil {
		if s == nil && !r.Optional {
			fail = true
			messages = append(messages, "Equal: value must exist")
		}
		if s != nil && *s != *r.Equal {
			fail = true
			messages = append(messages, fmt.Sprintf("Equal: value must equal '%s'", r.Equal.Value()))
		}
	}

	// The value must be longer or equal to the provided minimum length
	// MinLength Int
	if r.MinLength != nil {
		if s == nil && !r.Optional {
			fail = true
			messages = append(messages, "MinLength: value must exist")
		}
		if s != nil && len(s.Value()) < r.MinLength.Value() {
			fail = true
			messages = append(messages, fmt.Sprintf("MinLength: length must not be less than %d", r.MinLength.Value()))
		}
	}

	// The value must be shorter or equal to the provided maximum length
	// MaxLength Int
	if r.MaxLength != nil {
		if s == nil && !r.Optional {
			fail = true
			messages = append(messages, "MaxLength: value must exist")
		}
		if s != nil && len(s.Value()) > r.MaxLength.Value() {
			fail = true
			messages = append(messages, fmt.Sprintf("MaxLength: length must not be greater than %d", r.MaxLength.Value()))
		}
	}

	// The value of this string is restricted to one of the provided values
	// Enum []string
	if len(r.Enum) > 0 {
		if s == nil && !r.Optional {
			fail = true
			messages = append(messages, "Enum: value must exist")
		}
		if s != nil {
			found := false
			for _, e := range r.Enum {
				if e == s.Value() {
					found = true
				}
			}
			if !found {
				fail = true
				messages = append(messages, fmt.Sprintf("Enum: value must be one of: %v", r.Enum))
			}
		}
	}

	return
}

var _ Enforcer = (*StringRule)(nil)

func (out *String) Unpack(ctx context.Context, in json.Packed) error {
	if in == nil || in.Type() == json.J_NULL {
		return kerr.New("PWTAHLCCWR", "Called String.Unpack with nil value")
	}
	if in.Type() != json.J_STRING {
		return kerr.New("IXASCXOPMG", "Can't unpack %s into *system.String", in.Type())
	}
	s := NewString(in.String())
	*out = *s
	return nil
}

var _ json.Unpacker = (*String)(nil)

func (s *String) MarshalJSON(ctx context.Context) ([]byte, error) {
	if s == nil {
		return []byte("null"), nil
	}
	return []byte(strconv.Quote(s.Value())), nil
}

var _ json.Marshaler = (*String)(nil)

func (s *String) String() string {
	if s == nil {
		return ""
	}
	return s.Value()
}

type NativeString interface {
	NativeString() string
}

func (s String) NativeString() string {
	return string(s)
}

var _ NativeString = (*String)(nil)

func (r *StringRule) GetDefault() interface{} {
	return r.Default
}

var _ DefaultRule = (*StringRule)(nil)
