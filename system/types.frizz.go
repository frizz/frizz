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
	Field: unpack_Field,
	Type:  unpack_Type,
}

func unpack_Type(root *frizz.Root, stack frizz.Stack, in interface{}) (value Type, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Type
	if v, ok := m["Fields"]; ok {
		stack := stack.Append(frizz.FieldItem("Fields"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]Field, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]Field, len(m))
			for k, v := range m {
				stack := stack.Append(frizz.MapItem(k))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Field, err error) {
					// localUnpacker
					out, err := unpack_Field(root, stack, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(root, stack, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Fields = u
	}
	return out, nil
}
func unpack_Field(root *frizz.Root, stack frizz.Stack, in interface{}) (value Field, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Field
	if v, ok := m["Validators"]; ok {
		stack := stack.Append(frizz.FieldItem("Validators"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []common.Validator, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]common.Validator, len(a))
			for i, v := range a {
				stack := stack.Append(frizz.ArrayItem(i))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value common.Validator, err error) {
					// selectorUnpacker
					out, err := common.Unpackers.Validator(root, stack, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(root, stack, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Validators = u
	}
	return out, nil
}
func init() {
	frizz.DefaultRegistry.Set("frizz.io/system", "Type", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Type(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/system", "Field", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Field(root, stack, in)
	})
}
