package validators

import (
	json "encoding/json"
	common "frizz.io/common"
	global "frizz.io/global"
	pack "frizz.io/pack"
	system "frizz.io/system"
	errors "github.com/pkg/errors"
)

const Package packageType = 0

type packageType int

func (p packageType) Path() string {
	return "frizz.io/validators"
}
func (p packageType) Unpack(context global.Context, root global.Root, stack global.Stack, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Equal":
		return p.UnpackEqual(context, root, stack, in)
	case "GreaterThan":
		return p.UnpackGreaterThan(context, root, stack, in)
	case "GreaterThanOrEqual":
		return p.UnpackGreaterThanOrEqual(context, root, stack, in)
	case "IsNull":
		return p.UnpackIsNull(context, root, stack, in)
	case "Items":
		return p.UnpackItems(context, root, stack, in)
	case "Keys":
		return p.UnpackKeys(context, root, stack, in)
	case "Length":
		return p.UnpackLength(context, root, stack, in)
	case "LessThan":
		return p.UnpackLessThan(context, root, stack, in)
	case "LessThanOrEqual":
		return p.UnpackLessThanOrEqual(context, root, stack, in)
	case "NotNull":
		return p.UnpackNotNull(context, root, stack, in)
	case "Regex":
		return p.UnpackRegex(context, root, stack, in)
	case "Struct":
		return p.UnpackStruct(context, root, stack, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packageType) UnpackEqual(context global.Context, root global.Root, stack global.Stack, in interface{}) (value Equal, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value system.Raw, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker
		out, null, err := system.Package.UnpackRaw(context, root, stack, in)
		if err != nil || null {
			return value, null, err
		}
		return out, false, nil
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return Equal(out), false, nil
}
func (p packageType) UnpackGreaterThan(context global.Context, root global.Root, stack global.Stack, in interface{}) (value GreaterThan, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value json.Number, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker (json.Number)
		out, ok := in.(json.Number)
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into json.Number, found %T", stack, in)
		}
		if out == "" {
			return value, true, nil
		}
		return out, false, nil
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return GreaterThan(out), false, nil
}
func (p packageType) UnpackGreaterThanOrEqual(context global.Context, root global.Root, stack global.Stack, in interface{}) (value GreaterThanOrEqual, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value json.Number, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker (json.Number)
		out, ok := in.(json.Number)
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into json.Number, found %T", stack, in)
		}
		if out == "" {
			return value, true, nil
		}
		return out, false, nil
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return GreaterThanOrEqual(out), false, nil
}
func (p packageType) UnpackIsNull(context global.Context, root global.Root, stack global.Stack, in interface{}) (value IsNull, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value struct{}, null bool, err error) {
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
		var out struct{}
		return out, false, nil
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return IsNull(out), false, nil
}
func (p packageType) UnpackItems(context global.Context, root global.Root, stack global.Stack, in interface{}) (value Items, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value []common.Validator, null bool, err error) {
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
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return Items(out), false, nil
}
func (p packageType) UnpackKeys(context global.Context, root global.Root, stack global.Stack, in interface{}) (value Keys, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value []common.Validator, null bool, err error) {
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
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return Keys(out), false, nil
}
func (p packageType) UnpackLength(context global.Context, root global.Root, stack global.Stack, in interface{}) (value Length, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value []common.Validator, null bool, err error) {
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
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return Length(out), false, nil
}
func (p packageType) UnpackLessThan(context global.Context, root global.Root, stack global.Stack, in interface{}) (value LessThan, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value json.Number, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker (json.Number)
		out, ok := in.(json.Number)
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into json.Number, found %T", stack, in)
		}
		if out == "" {
			return value, true, nil
		}
		return out, false, nil
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return LessThan(out), false, nil
}
func (p packageType) UnpackLessThanOrEqual(context global.Context, root global.Root, stack global.Stack, in interface{}) (value LessThanOrEqual, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value json.Number, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker (json.Number)
		out, ok := in.(json.Number)
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into json.Number, found %T", stack, in)
		}
		if out == "" {
			return value, true, nil
		}
		return out, false, nil
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return LessThanOrEqual(out), false, nil
}
func (p packageType) UnpackNotNull(context global.Context, root global.Root, stack global.Stack, in interface{}) (value NotNull, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value struct{}, null bool, err error) {
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
		var out struct{}
		return out, false, nil
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return NotNull(out), false, nil
}
func (p packageType) UnpackRegex(context global.Context, root global.Root, stack global.Stack, in interface{}) (value Regex, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value struct {
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
			stack := stack.Append(global.FieldItem("Regex"))
			u, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackString(stack, in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Regex = u
			}
		}
		if v, ok := m["Invert"]; ok {
			stack := stack.Append(global.FieldItem("Invert"))
			u, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value bool, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackBool(stack, in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Invert = u
			}
		}
		return out, false, nil
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return Regex(out), false, nil
}
func (p packageType) UnpackStruct(context global.Context, root global.Root, stack global.Stack, in interface{}) (value Struct, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.Context, root global.Root, stack global.Stack, in interface{}) (value map[string][]common.Validator, null bool, err error) {
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
			stack := stack.Append(global.MapItem(k))
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
				out[k] = u
			}
		}
		return out, false, nil
	}(context, root, stack, in)
	if err != nil || null {
		return value, null, err
	}
	return Struct(out), false, nil
}
func (p packageType) Repack(context global.Context, root global.Root, stack global.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Equal":
		return p.RepackEqual(context, root, stack, in.(Equal))
	case "GreaterThan":
		return p.RepackGreaterThan(context, root, stack, in.(GreaterThan))
	case "GreaterThanOrEqual":
		return p.RepackGreaterThanOrEqual(context, root, stack, in.(GreaterThanOrEqual))
	case "IsNull":
		return p.RepackIsNull(context, root, stack, in.(IsNull))
	case "Items":
		return p.RepackItems(context, root, stack, in.(Items))
	case "Keys":
		return p.RepackKeys(context, root, stack, in.(Keys))
	case "Length":
		return p.RepackLength(context, root, stack, in.(Length))
	case "LessThan":
		return p.RepackLessThan(context, root, stack, in.(LessThan))
	case "LessThanOrEqual":
		return p.RepackLessThanOrEqual(context, root, stack, in.(LessThanOrEqual))
	case "NotNull":
		return p.RepackNotNull(context, root, stack, in.(NotNull))
	case "Regex":
		return p.RepackRegex(context, root, stack, in.(Regex))
	case "Struct":
		return p.RepackStruct(context, root, stack, in.(Struct))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packageType) RepackEqual(context global.Context, root global.Root, stack global.Stack, in Equal) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in system.Raw) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker
		out, dict, null, err := system.Package.RepackRaw(context, root, stack, in)
		if err != nil {
			return nil, false, false, err
		}
		return out, dict, null, nil
	}(context, root, stack, (system.Raw)(in))
}
func (p packageType) RepackGreaterThan(context global.Context, root global.Root, stack global.Stack, in GreaterThan) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in json.Number) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker (json.Number)
		if in == "" {
			return value, false, true, nil
		}
		return in, false, false, nil
	}(context, root, stack, (json.Number)(in))
}
func (p packageType) RepackGreaterThanOrEqual(context global.Context, root global.Root, stack global.Stack, in GreaterThanOrEqual) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in json.Number) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker (json.Number)
		if in == "" {
			return value, false, true, nil
		}
		return in, false, false, nil
	}(context, root, stack, (json.Number)(in))
}
func (p packageType) RepackIsNull(context global.Context, root global.Root, stack global.Stack, in IsNull) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in struct{}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 1)
		empty := true
		return out, false, empty, nil
	}(context, root, stack, (struct{})(in))
}
func (p packageType) RepackItems(context global.Context, root global.Root, stack global.Stack, in Items) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
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
	}(context, root, stack, ([]common.Validator)(in))
}
func (p packageType) RepackKeys(context global.Context, root global.Root, stack global.Stack, in Keys) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
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
	}(context, root, stack, ([]common.Validator)(in))
}
func (p packageType) RepackLength(context global.Context, root global.Root, stack global.Stack, in Length) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
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
	}(context, root, stack, ([]common.Validator)(in))
}
func (p packageType) RepackLessThan(context global.Context, root global.Root, stack global.Stack, in LessThan) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in json.Number) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker (json.Number)
		if in == "" {
			return value, false, true, nil
		}
		return in, false, false, nil
	}(context, root, stack, (json.Number)(in))
}
func (p packageType) RepackLessThanOrEqual(context global.Context, root global.Root, stack global.Stack, in LessThanOrEqual) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in json.Number) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker (json.Number)
		if in == "" {
			return value, false, true, nil
		}
		return in, false, false, nil
	}(context, root, stack, (json.Number)(in))
}
func (p packageType) RepackNotNull(context global.Context, root global.Root, stack global.Stack, in NotNull) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in struct{}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 1)
		empty := true
		return out, false, empty, nil
	}(context, root, stack, (struct{})(in))
}
func (p packageType) RepackRegex(context global.Context, root global.Root, stack global.Stack, in Regex) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in struct {
		Regex  string
		Invert bool
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(context global.Context, root global.Root, stack global.Stack, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackString(in)
		}(context, root, stack, in.Regex); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Regex"] = v
			}
		}
		if v, _, null, err := func(context global.Context, root global.Root, stack global.Stack, in bool) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackBool(in)
		}(context, root, stack, in.Invert); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Invert"] = v
			}
		}
		return out, false, empty, nil
	}(context, root, stack, (struct {
		Regex  string
		Invert bool
	})(in))
}
func (p packageType) RepackStruct(context global.Context, root global.Root, stack global.Stack, in Struct) (value interface{}, dict bool, null bool, err error) {
	return func(context global.Context, root global.Root, stack global.Stack, in map[string][]common.Validator) (value interface{}, dict bool, null bool, err error) {
		// mapRepacker
		out := make(map[string]interface{}, len(in))
		for k, item := range in {
			v, _, _, err := func(context global.Context, root global.Root, stack global.Stack, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
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
			}(context, root, stack, item)
			if err != nil {
				return nil, false, false, err
			}
			out[k] = v
		}
		return out, true, len(in) == 0, nil
	}(context, root, stack, (map[string][]common.Validator)(in))
}
func (p packageType) Get(name string) string {
	return ""
}
func (p packageType) Add(packages map[string]global.Package) {
	packages["frizz.io/validators"] = Package
}
