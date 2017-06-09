package validators

import (
	common "frizz.io/common"
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

var Unpackers = struct {
	Keys  func(*frizz.Root, frizz.Stack, interface{}) (Keys, error)
	Items func(*frizz.Root, frizz.Stack, interface{}) (Items, error)
	Regex func(*frizz.Root, frizz.Stack, interface{}) (Regex, error)
}{
	Items: unpacker_Items,
	Keys:  unpacker_Keys,
	Regex: unpacker_Regex,
}

func unpacker_Keys(r *frizz.Root, s frizz.Stack, in interface{}) (value Keys, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Keys
	if v, ok := m["Validators"]; ok {
		s := s.Append(frizz.FieldItem("Validators"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value []common.Validator, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]common.Validator, len(a))
			for i, v := range a {
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value common.Validator, err error) {
					// selectorUnpacker
					out, err := common.Unpackers.Validator(r, s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Validators = u
	}
	return out, nil
}
func unpacker_Items(r *frizz.Root, s frizz.Stack, in interface{}) (value Items, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Items
	if v, ok := m["Validators"]; ok {
		s := s.Append(frizz.FieldItem("Validators"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value []common.Validator, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]common.Validator, len(a))
			for i, v := range a {
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value common.Validator, err error) {
					// selectorUnpacker
					out, err := common.Unpackers.Validator(r, s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Validators = u
	}
	return out, nil
}
func unpacker_Regex(r *frizz.Root, s frizz.Stack, in interface{}) (value Regex, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Regex
	if v, ok := m["Regex"]; ok {
		s := s.Append(frizz.FieldItem("Regex"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackString(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Regex = u
	}
	return out, nil
}
func init() {
	frizz.DefaultRegistry.Set("frizz.io/validators", "Keys", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Keys(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/validators", "Items", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Items(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/validators", "Regex", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Regex(r, s, in)
	})
}
