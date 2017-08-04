package system

import (
	common "frizz.io/common"
	global "frizz.io/global"
	errors "github.com/pkg/errors"
)

const Package packageType = 0

type packageType int

func (p packageType) Path() string {
	return "frizz.io/system"
}
func (p packageType) Unpack(context global.Context, root global.Root, stack global.Stack, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Raw":
		return p.UnpackRaw(context, root, stack, in)
	case "Type":
		return p.UnpackType(context, root, stack, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packageType) UnpackRaw(context global.Context, root global.Root, stack global.Stack, in interface{}) (value Raw, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// customUnpacker
	out := new(Raw)
	null, err = out.Unpack(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return *out, false, nil
}
func (p packageType) UnpackType(context global.Context, root global.Root, stack global.Stack, in interface{}) (value Type, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value struct {
		Validators []common.Validator
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
			Validators []common.Validator
		}
		if v, ok := m["Validators"]; ok {
			stack := stack.Append(global.FieldItem("Validators"))
			u, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value []common.Validator, null bool, err error) {
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
					stack := stack.Append(global.ArrayItem(i))
					u, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value common.Validator, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// selectorUnpacker
						out, null, err := common.Package.UnpackValidator(context, root, stack, in)
						if err != nil || null {
							return value, null, err
						}
						return out, false, nil
					}(context, root, stack, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out[i] = u
					}
				}
				return out[:], false, nil
			}(context, root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Validators = u
			}
		}
		return out, false, nil
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return Type(out), false, nil
}
func (p packageType) Repack(context global.Context, root global.Root, stack global.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Raw":
		return p.RepackRaw(context, root, stack, in.(Raw))
	case "Type":
		return p.RepackType(context, root, stack, in.(Type))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packageType) RepackRaw(context global.Context, root global.Root, stack global.Stack, in Raw) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(context, root, stack)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
}
func (p packageType) RepackType(context global.Context, root global.Root, stack global.Stack, in Type) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in struct {
		Validators []common.Validator
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(context global.Context, root global.Root, stack global.Stack, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.Context, root global.Root, stack global.Stack, in common.Validator) (value interface{}, dict bool, null bool, err error) {
					// selectorRepacker
					out, dict, null, err := common.Package.RepackValidator(context, root, stack, in)
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, null, nil
				}(context, root, stack, item)
				if err != nil {
					return nil, false, false, err
				}
				if !null {
					empty = false
				}
				out[i] = v
			}
			return out, false, empty, nil
		}(context, root, stack, in.Validators); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Validators"] = v
			}
		}
		return out, false, empty, nil
	}(context, root, stack, (struct {
		Validators []common.Validator
	})(in))
}
func (p packageType) Get(name string) string {
	return ""
}
func (p packageType) Add(packages map[string]global.Package) {
	packages["frizz.io/system"] = Package
}
