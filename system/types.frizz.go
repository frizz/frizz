package system

import (
	common "frizz.io/common"
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

var Unpackers = struct {
	Type  func(*frizz.Root, frizz.Stack, interface{}) (Type, error)
	Field func(*frizz.Root, frizz.Stack, interface{}) (Field, error)
}{
	Field: unpacker_Field,
	Type:  unpacker_Type,
}

func unpacker_Type(r *frizz.Root, s frizz.Stack, in interface{}) (value Type, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Type
	if v, ok := m["Fields"]; ok {
		s := s.Append(frizz.FieldItem("Fields"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value map[string]Field, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]Field, len(m))
			for k, v := range m {
				s := s.Append(frizz.MapItem(k))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value Field, err error) {
					// localUnpacker
					out, err := unpacker_Field(r, s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Fields = u
	}
	return out, nil
}
func unpacker_Field(r *frizz.Root, s frizz.Stack, in interface{}) (value Field, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Field
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
func init() {
	frizz.DefaultRegistry.Set("frizz.io/system", "Type", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Type(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/system", "Field", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Field(r, s, in)
	})
}
