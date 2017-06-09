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

func unpacker_Keys(root *frizz.Root, stack frizz.Stack, in interface{}) (value Keys, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Keys
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
func unpacker_Items(root *frizz.Root, stack frizz.Stack, in interface{}) (value Items, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Items
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
func unpacker_Regex(root *frizz.Root, stack frizz.Stack, in interface{}) (value Regex, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Regex
	if v, ok := m["Regex"]; ok {
		stack := stack.Append(frizz.FieldItem("Regex"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackString(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Regex = u
	}
	return out, nil
}
func init() {
	frizz.DefaultRegistry.Set("frizz.io/validators", "Keys", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Keys(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/validators", "Items", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Items(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/validators", "Regex", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Regex(root, stack, in)
	})
}
