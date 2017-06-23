package validators

import (
	common "frizz.io/common"
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

const Packer packer = 0

type packer int

func (p packer) Path() string {
	return "frizz.io/validators"
}
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (interface{}, error) {
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
func (p packer) UnpackKeys(root *frizz.Root, stack frizz.Stack, in interface{}) (value Keys, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Validators []common.Validator
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			Validators []common.Validator
		}
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
						out, err := common.Packer.UnpackValidator(root, stack, in)
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Keys(out), nil
}
func (p packer) UnpackItems(root *frizz.Root, stack frizz.Stack, in interface{}) (value Items, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Validators []common.Validator
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			Validators []common.Validator
		}
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
						out, err := common.Packer.UnpackValidator(root, stack, in)
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Items(out), nil
}
func (p packer) UnpackRegex(root *frizz.Root, stack frizz.Stack, in interface{}) (value Regex, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Regex string
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			Regex string
		}
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Regex(out), nil
}
func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (interface{}, bool, error) {
	switch name {
	case "Keys":
		return p.RepackKeys(root, stack, in.(Keys))
	case "Items":
		return p.RepackItems(root, stack, in.(Items))
	case "Regex":
		return p.RepackRegex(root, stack, in.(Regex))
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) RepackKeys(root *frizz.Root, stack frizz.Stack, in Keys) (interface{}, bool, error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		Validators []common.Validator
	}) (interface{}, bool, error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		if v, _, err := func(root *frizz.Root, stack frizz.Stack, in []common.Validator) (interface{}, bool, error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			for i, item := range in {
				v, _, err := func(root *frizz.Root, stack frizz.Stack, in common.Validator) (interface{}, bool, error) {
					// selectorRepacker
					out, dict, err := common.Packer.RepackValidator(root, stack, in)
					if err != nil {
						return nil, false, err
					}
					return out, dict, nil
				}(root, stack, item)
				if err != nil {
					return nil, false, err
				}
				out[i] = v
			}
			return out, false, nil
		}(root, stack, in.Validators); err != nil {
			return nil, false, err
		} else {
			out["Validators"] = v
		}
		return out, false, nil
	}(root, stack, (struct {
		Validators []common.Validator
	})(in))
}
func (p packer) RepackItems(root *frizz.Root, stack frizz.Stack, in Items) (interface{}, bool, error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		Validators []common.Validator
	}) (interface{}, bool, error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		if v, _, err := func(root *frizz.Root, stack frizz.Stack, in []common.Validator) (interface{}, bool, error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			for i, item := range in {
				v, _, err := func(root *frizz.Root, stack frizz.Stack, in common.Validator) (interface{}, bool, error) {
					// selectorRepacker
					out, dict, err := common.Packer.RepackValidator(root, stack, in)
					if err != nil {
						return nil, false, err
					}
					return out, dict, nil
				}(root, stack, item)
				if err != nil {
					return nil, false, err
				}
				out[i] = v
			}
			return out, false, nil
		}(root, stack, in.Validators); err != nil {
			return nil, false, err
		} else {
			out["Validators"] = v
		}
		return out, false, nil
	}(root, stack, (struct {
		Validators []common.Validator
	})(in))
}
func (p packer) RepackRegex(root *frizz.Root, stack frizz.Stack, in Regex) (interface{}, bool, error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		Regex string
	}) (interface{}, bool, error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		if v, _, err := func(root *frizz.Root, stack frizz.Stack, in string) (interface{}, bool, error) {
			// nativeRepacker
			return frizz.RepackString(in), false, nil
		}(root, stack, in.Regex); err != nil {
			return nil, false, err
		} else {
			out["Regex"] = v
		}
		return out, false, nil
	}(root, stack, (struct {
		Regex string
	})(in))
}
