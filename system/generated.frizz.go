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
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Type":
		return p.UnpackType(root, stack, in)
	case "Field":
		return p.UnpackField(root, stack, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) UnpackType(root *frizz.Root, stack frizz.Stack, in interface{}) (value Type, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		For    string
		Fields map[string]Field
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.New("unpacking into struct, value should be a map")
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out struct {
			For    string
			Fields map[string]Field
		}
		if v, ok := m["For"]; ok {
			stack := stack.Append(frizz.FieldItem("For"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackString(stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.For = u
			}
		}
		if v, ok := m["Fields"]; ok {
			stack := stack.Append(frizz.FieldItem("Fields"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]Field, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// mapUnpacker
				m, ok := in.(map[string]interface{})
				if !ok {
					return value, false, errors.New("unpacking into map, value should be a map")
				}
				if len(m) == 0 {
					return value, true, nil
				}
				var out = make(map[string]Field, len(m))
				for k, v := range m {
					stack := stack.Append(frizz.MapItem(k))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Field, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// localUnpacker
						out, null, err := p.UnpackField(root, stack, in)
						if err != nil {
							return value, false, err
						}
						if null {
							return value, true, nil
						}
						return out, false, nil
					}(root, stack, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out[k] = u
					}
				}
				return out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Fields = u
			}
		}
		return out, false, nil
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return Type(out), false, nil
}
func (p packer) UnpackField(root *frizz.Root, stack frizz.Stack, in interface{}) (value Field, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Validators []common.Validator
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.New("unpacking into struct, value should be a map")
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out struct {
			Validators []common.Validator
		}
		if v, ok := m["Validators"]; ok {
			stack := stack.Append(frizz.FieldItem("Validators"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []common.Validator, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// sliceUnpacker
				a, ok := in.([]interface{})
				if !ok {
					return value, false, errors.New("unpacking into slice, value should be an array")
				}
				if len(a) == 0 {
					return value, true, nil
				}
				var out = make([]common.Validator, len(a))
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value common.Validator, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// selectorUnpacker
						out, null, err := common.Packer.UnpackValidator(root, stack, in)
						if err != nil {
							return value, false, err
						}
						if null {
							return value, true, nil
						}
						return out, false, nil
					}(root, stack, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out[i] = u
					}
				}
				return out[:], false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Validators = u
			}
		}
		return out, false, nil
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return Field(out), false, nil
}
func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Type":
		return p.RepackType(root, stack, in.(Type))
	case "Field":
		return p.RepackField(root, stack, in.(Field))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) RepackType(root *frizz.Root, stack frizz.Stack, in Type) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		For    string
		Fields map[string]Field
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackString(in)
		}(root, stack, in.For); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["For"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in map[string]Field) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(root *frizz.Root, stack frizz.Stack, in Field) (value interface{}, dict bool, null bool, err error) {
					// localRepacker
					out, dict, null, err := p.RepackField(root, stack, in)
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, null, nil
				}(root, stack, item)
				if err != nil {
					return nil, false, false, err
				}
				out[k] = v
			}
			return out, true, len(in) == 0, nil
		}(root, stack, in.Fields); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Fields"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		For    string
		Fields map[string]Field
	})(in))
}
func (p packer) RepackField(root *frizz.Root, stack frizz.Stack, in Field) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		Validators []common.Validator
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in common.Validator) (value interface{}, dict bool, null bool, err error) {
					// selectorRepacker
					out, dict, null, err := common.Packer.RepackValidator(root, stack, in)
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, null, nil
				}(root, stack, item)
				if err != nil {
					return nil, false, false, err
				}
				if !null {
					empty = false
				}
				out[i] = v
			}
			return out, false, empty, nil
		}(root, stack, in.Validators); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Validators"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		Validators []common.Validator
	})(in))
}

const Types types = 0

type types int

func (t types) Path() string {
	return "frizz.io/system"
}
func (t types) Get(name string) string {
	return nil
}
