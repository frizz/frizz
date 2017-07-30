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
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Items":
		return p.UnpackItems(root, stack, in)
	case "Keys":
		return p.UnpackKeys(root, stack, in)
	case "Regex":
		return p.UnpackRegex(root, stack, in)
	case "Struct":
		return p.UnpackStruct(root, stack, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) UnpackItems(root *frizz.Root, stack frizz.Stack, in interface{}) (value Items, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []common.Validator, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// sliceUnpacker
		a, ok := in.([]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into slice, value should be an array", stack)
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
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return Items(out), false, nil
}
func (p packer) UnpackKeys(root *frizz.Root, stack frizz.Stack, in interface{}) (value Keys, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []common.Validator, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// sliceUnpacker
		a, ok := in.([]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into slice, value should be an array", stack)
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
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return Keys(out), false, nil
}
func (p packer) UnpackRegex(root *frizz.Root, stack frizz.Stack, in interface{}) (value Regex, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Regex  string
		Invert bool
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into struct, value should be a map", stack)
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out struct {
			Regex  string
			Invert bool
		}
		if v, ok := m["Regex"]; ok {
			stack := stack.Append(frizz.FieldItem("Regex"))
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
				out.Regex = u
			}
		}
		if v, ok := m["Invert"]; ok {
			stack := stack.Append(frizz.FieldItem("Invert"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value bool, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackBool(stack, in)
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
				out.Invert = u
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
	return Regex(out), false, nil
}
func (p packer) UnpackStruct(root *frizz.Root, stack frizz.Stack, in interface{}) (value Struct, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string][]common.Validator, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// mapUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.Errorf("unpacking into map, value should be a map", stack)
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out = make(map[string][]common.Validator, len(m))
		for k, v := range m {
			stack := stack.Append(frizz.MapItem(k))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []common.Validator, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// sliceUnpacker
				a, ok := in.([]interface{})
				if !ok {
					return value, false, errors.Errorf("%s: unpacking into slice, value should be an array", stack)
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
				out[k] = u
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
	return Struct(out), false, nil
}
func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Items":
		return p.RepackItems(root, stack, in.(Items))
	case "Keys":
		return p.RepackKeys(root, stack, in.(Keys))
	case "Regex":
		return p.RepackRegex(root, stack, in.(Regex))
	case "Struct":
		return p.RepackStruct(root, stack, in.(Struct))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) RepackItems(root *frizz.Root, stack frizz.Stack, in Items) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
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
	}(root, stack, ([]common.Validator)(in))
}
func (p packer) RepackKeys(root *frizz.Root, stack frizz.Stack, in Keys) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
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
	}(root, stack, ([]common.Validator)(in))
}
func (p packer) RepackRegex(root *frizz.Root, stack frizz.Stack, in Regex) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		Regex  string
		Invert bool
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackString(in)
		}(root, stack, in.Regex); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Regex"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in bool) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackBool(in)
		}(root, stack, in.Invert); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Invert"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		Regex  string
		Invert bool
	})(in))
}
func (p packer) RepackStruct(root *frizz.Root, stack frizz.Stack, in Struct) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in map[string][]common.Validator) (value interface{}, dict bool, null bool, err error) {
		// mapRepacker
		out := make(map[string]interface{}, len(in))
		for k, item := range in {
			v, _, _, err := func(root *frizz.Root, stack frizz.Stack, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
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
			}(root, stack, item)
			if err != nil {
				return nil, false, false, err
			}
			out[k] = v
		}
		return out, true, len(in) == 0, nil
	}(root, stack, (map[string][]common.Validator)(in))
}

const Imports imports = 0

type imports int

func (i imports) Path() string {
	return "frizz.io/validators"
}
func (i imports) Add(packers map[string]frizz.Packer, types map[string]frizz.Typer) {
	if packers != nil {
		packers["frizz.io/validators"] = Packer
	}
}
