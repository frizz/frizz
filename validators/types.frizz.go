package validators

import (
	common "frizz.io/common"
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

type Packer struct{}

func (Packer) Path() string {
	return "frizz.io/validators"
}
func (p Packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (interface{}, error) {
	switch name {
	case "Keys":
		return p.UnpackKeys(root, stack, in)
	case "Items":
		return p.UnpackItems(root, stack, in)
	case "Regex":
		return p.UnpackRegex(root, stack, in)
	}
	return nil, errors.Errorf("%s: type %s not found", stack, name)
}
func (p Packer) UnpackKeys(root *frizz.Root, stack frizz.Stack, in interface{}) (value Keys, err error) {
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
					out, err := common.Packer{}.UnpackValidator(root, stack, in)
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
func (p Packer) UnpackItems(root *frizz.Root, stack frizz.Stack, in interface{}) (value Items, err error) {
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
					out, err := common.Packer{}.UnpackValidator(root, stack, in)
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
func (p Packer) UnpackRegex(root *frizz.Root, stack frizz.Stack, in interface{}) (value Regex, err error) {
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
