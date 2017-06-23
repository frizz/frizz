package system

import (
	common "frizz.io/common"
	frizz "frizz.io/frizz"
	errors "github.com/pkg/errors"
)

const Packer packer = 0

type packer int

func (p packer) Path() string {
	return "frizz.io/system"
}
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (interface{}, error) {
	switch name {
	case "Type":
		return p.UnpackType(root, stack, in)
	case "Field":
		return p.UnpackField(root, stack, in)
	}
	return nil, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) UnpackType(root *frizz.Root, stack frizz.Stack, in interface{}) (value Type, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Fields map[string]Field
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			Fields map[string]Field
		}
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
						out, err := p.UnpackField(root, stack, in)
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Type(out), nil
}
func (p packer) UnpackField(root *frizz.Root, stack frizz.Stack, in interface{}) (value Field, err error) {
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
	return Field(out), nil
}
func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (interface{}, bool, error) {
	switch name {
	case "Type":
		return p.RepackType(root, stack, in.(Type))
	case "Field":
		return p.RepackField(root, stack, in.(Field))
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) RepackType(root *frizz.Root, stack frizz.Stack, in Type) (interface{}, bool, error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		Fields map[string]Field
	}) (interface{}, bool, error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		if v, _, err := func(root *frizz.Root, stack frizz.Stack, in map[string]Field) (interface{}, bool, error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, err := func(root *frizz.Root, stack frizz.Stack, in Field) (interface{}, bool, error) {
					// localRepacker
					out, dict, err := p.RepackField(root, stack, in)
					if err != nil {
						return nil, false, err
					}
					return out, dict, nil
				}(root, stack, item)
				if err != nil {
					return nil, false, err
				}
				out[k] = v
			}
			return out, true, nil
		}(root, stack, in.Fields); err != nil {
			return nil, false, err
		} else {
			out["Fields"] = v
		}
		return out, false, nil
	}(root, stack, (struct {
		Fields map[string]Field
	})(in))
}
func (p packer) RepackField(root *frizz.Root, stack frizz.Stack, in Field) (interface{}, bool, error) {
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
