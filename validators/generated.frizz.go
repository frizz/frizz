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
func (p packageType) Unpack(context global.DataContext, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Equal":
		return p.UnpackEqual(context, in)
	case "GreaterThan":
		return p.UnpackGreaterThan(context, in)
	case "GreaterThanOrEqual":
		return p.UnpackGreaterThanOrEqual(context, in)
	case "IsNull":
		return p.UnpackIsNull(context, in)
	case "Items":
		return p.UnpackItems(context, in)
	case "Keys":
		return p.UnpackKeys(context, in)
	case "Length":
		return p.UnpackLength(context, in)
	case "LessThan":
		return p.UnpackLessThan(context, in)
	case "LessThanOrEqual":
		return p.UnpackLessThanOrEqual(context, in)
	case "NotNull":
		return p.UnpackNotNull(context, in)
	case "Regex":
		return p.UnpackRegex(context, in)
	case "Struct":
		return p.UnpackStruct(context, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) UnpackEqual(context global.DataContext, in interface{}) (value Equal, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value system.Raw, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker
		p := context.Package().Get("frizz.io/system")
		if p != nil {
			out, null, err := p.Unpack(context, in, "Raw")
			if err != nil || null {
				return value, null, err
			}
			return out.(system.Raw), false, nil
		}
		var vi interface{} = &value
		if vu, ok := vi.(json.Unmarshaler); ok {
			b, err := json.Marshal(in)
			if err != nil {
				return value, false, err
			}
			if err := vu.UnmarshalJSON(b); err != nil {
				return value, false, err
			}
			return value, false, nil
		}
		return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/system", "Raw")
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Equal(out), false, nil
}
func (p packageType) UnpackGreaterThan(context global.DataContext, in interface{}) (value GreaterThan, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value json.Number, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker (json.Number)
		out, ok := in.(json.Number)
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into json.Number, found %T", context.Location(), in)
		}
		if out == "" {
			return value, true, nil
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return GreaterThan(out), false, nil
}
func (p packageType) UnpackGreaterThanOrEqual(context global.DataContext, in interface{}) (value GreaterThanOrEqual, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value json.Number, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker (json.Number)
		out, ok := in.(json.Number)
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into json.Number, found %T", context.Location(), in)
		}
		if out == "" {
			return value, true, nil
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return GreaterThanOrEqual(out), false, nil
}
func (p packageType) UnpackIsNull(context global.DataContext, in interface{}) (value IsNull, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct{}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into struct, value should be a map, found: %#v", context.Location(), in)
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out struct{}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return IsNull(out), false, nil
}
func (p packageType) UnpackItems(context global.DataContext, in interface{}) (value Items, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value []common.Validator, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// sliceUnpacker
		a, ok := in.([]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into slice, value should be an array, found: %#v", context.Location(), in)
		}
		if len(a) == 0 {
			return value, true, nil
		}
		var out = make([]common.Validator, len(a))
		for i, v := range a {
			context := context.New(context.Location().Child(global.ArrayItem(i)))
			u, null, err := func(context global.DataContext, in interface{}) (value common.Validator, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// selectorUnpacker
				p := context.Package().Get("frizz.io/common")
				if p != nil {
					out, null, err := p.Unpack(context, in, "Validator")
					if err != nil || null {
						return value, null, err
					}
					return out.(common.Validator), false, nil
				}
				var vi interface{} = &value
				if vu, ok := vi.(json.Unmarshaler); ok {
					b, err := json.Marshal(in)
					if err != nil {
						return value, false, err
					}
					if err := vu.UnmarshalJSON(b); err != nil {
						return value, false, err
					}
					return value, false, nil
				}
				return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/common", "Validator")
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out[i] = u
			}
		}
		return out[:], false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Items(out), false, nil
}
func (p packageType) UnpackKeys(context global.DataContext, in interface{}) (value Keys, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value []common.Validator, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// sliceUnpacker
		a, ok := in.([]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into slice, value should be an array, found: %#v", context.Location(), in)
		}
		if len(a) == 0 {
			return value, true, nil
		}
		var out = make([]common.Validator, len(a))
		for i, v := range a {
			context := context.New(context.Location().Child(global.ArrayItem(i)))
			u, null, err := func(context global.DataContext, in interface{}) (value common.Validator, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// selectorUnpacker
				p := context.Package().Get("frizz.io/common")
				if p != nil {
					out, null, err := p.Unpack(context, in, "Validator")
					if err != nil || null {
						return value, null, err
					}
					return out.(common.Validator), false, nil
				}
				var vi interface{} = &value
				if vu, ok := vi.(json.Unmarshaler); ok {
					b, err := json.Marshal(in)
					if err != nil {
						return value, false, err
					}
					if err := vu.UnmarshalJSON(b); err != nil {
						return value, false, err
					}
					return value, false, nil
				}
				return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/common", "Validator")
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out[i] = u
			}
		}
		return out[:], false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Keys(out), false, nil
}
func (p packageType) UnpackLength(context global.DataContext, in interface{}) (value Length, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value []common.Validator, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// sliceUnpacker
		a, ok := in.([]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into slice, value should be an array, found: %#v", context.Location(), in)
		}
		if len(a) == 0 {
			return value, true, nil
		}
		var out = make([]common.Validator, len(a))
		for i, v := range a {
			context := context.New(context.Location().Child(global.ArrayItem(i)))
			u, null, err := func(context global.DataContext, in interface{}) (value common.Validator, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// selectorUnpacker
				p := context.Package().Get("frizz.io/common")
				if p != nil {
					out, null, err := p.Unpack(context, in, "Validator")
					if err != nil || null {
						return value, null, err
					}
					return out.(common.Validator), false, nil
				}
				var vi interface{} = &value
				if vu, ok := vi.(json.Unmarshaler); ok {
					b, err := json.Marshal(in)
					if err != nil {
						return value, false, err
					}
					if err := vu.UnmarshalJSON(b); err != nil {
						return value, false, err
					}
					return value, false, nil
				}
				return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/common", "Validator")
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out[i] = u
			}
		}
		return out[:], false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Length(out), false, nil
}
func (p packageType) UnpackLessThan(context global.DataContext, in interface{}) (value LessThan, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value json.Number, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker (json.Number)
		out, ok := in.(json.Number)
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into json.Number, found %T", context.Location(), in)
		}
		if out == "" {
			return value, true, nil
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return LessThan(out), false, nil
}
func (p packageType) UnpackLessThanOrEqual(context global.DataContext, in interface{}) (value LessThanOrEqual, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value json.Number, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker (json.Number)
		out, ok := in.(json.Number)
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into json.Number, found %T", context.Location(), in)
		}
		if out == "" {
			return value, true, nil
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return LessThanOrEqual(out), false, nil
}
func (p packageType) UnpackNotNull(context global.DataContext, in interface{}) (value NotNull, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct{}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into struct, value should be a map, found: %#v", context.Location(), in)
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out struct{}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return NotNull(out), false, nil
}
func (p packageType) UnpackRegex(context global.DataContext, in interface{}) (value Regex, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Regex  string
		Invert bool
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into struct, value should be a map, found: %#v", context.Location(), in)
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out struct {
			Regex  string
			Invert bool
		}
		if v, ok := m["Regex"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Regex")))
			u, null, err := func(context global.DataContext, in interface{}) (value string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackString(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Regex = u
			}
		}
		if v, ok := m["Invert"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Invert")))
			u, null, err := func(context global.DataContext, in interface{}) (value bool, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackBool(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Invert = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Regex(out), false, nil
}
func (p packageType) UnpackStruct(context global.DataContext, in interface{}) (value Struct, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value map[string][]common.Validator, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// mapUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.Errorf("unpacking into map, value should be a map, found: %#v", context.Location(), in)
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out = make(map[string][]common.Validator, len(m))
		for k, v := range m {
			context := context.New(context.Location().Child(global.MapItem(k)))
			u, null, err := func(context global.DataContext, in interface{}) (value []common.Validator, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// sliceUnpacker
				a, ok := in.([]interface{})
				if !ok {
					return value, false, errors.Errorf("%s: unpacking into slice, value should be an array, found: %#v", context.Location(), in)
				}
				if len(a) == 0 {
					return value, true, nil
				}
				var out = make([]common.Validator, len(a))
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
					u, null, err := func(context global.DataContext, in interface{}) (value common.Validator, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// selectorUnpacker
						p := context.Package().Get("frizz.io/common")
						if p != nil {
							out, null, err := p.Unpack(context, in, "Validator")
							if err != nil || null {
								return value, null, err
							}
							return out.(common.Validator), false, nil
						}
						var vi interface{} = &value
						if vu, ok := vi.(json.Unmarshaler); ok {
							b, err := json.Marshal(in)
							if err != nil {
								return value, false, err
							}
							if err := vu.UnmarshalJSON(b); err != nil {
								return value, false, err
							}
							return value, false, nil
						}
						return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/common", "Validator")
					}(context, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out[i] = u
					}
				}
				return out[:], false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out[k] = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Struct(out), false, nil
}
func (p packageType) Repack(context global.DataContext, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Equal":
		return p.RepackEqual(context, in.(Equal))
	case "GreaterThan":
		return p.RepackGreaterThan(context, in.(GreaterThan))
	case "GreaterThanOrEqual":
		return p.RepackGreaterThanOrEqual(context, in.(GreaterThanOrEqual))
	case "IsNull":
		return p.RepackIsNull(context, in.(IsNull))
	case "Items":
		return p.RepackItems(context, in.(Items))
	case "Keys":
		return p.RepackKeys(context, in.(Keys))
	case "Length":
		return p.RepackLength(context, in.(Length))
	case "LessThan":
		return p.RepackLessThan(context, in.(LessThan))
	case "LessThanOrEqual":
		return p.RepackLessThanOrEqual(context, in.(LessThanOrEqual))
	case "NotNull":
		return p.RepackNotNull(context, in.(NotNull))
	case "Regex":
		return p.RepackRegex(context, in.(Regex))
	case "Struct":
		return p.RepackStruct(context, in.(Struct))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) RepackEqual(context global.DataContext, in Equal) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in system.Raw) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker
		p := context.Package().Get("frizz.io/system")
		if p != nil {
			out, dict, null, err := p.Repack(context, in, "Raw")
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}
		var vi interface{} = &in
		if vu, ok := vi.(json.Marshaler); ok {
			b, err := vu.MarshalJSON()
			if err != nil {
				return nil, false, false, err
			}
			if err := json.Unmarshal(b, &value); err != nil {
				return nil, false, false, err
			}
			return value, b[0] == '{', len(b) == 4 && string(b) == "null", nil
		}
		return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/system", "Raw")
	}(context, (system.Raw)(in))
}
func (p packageType) RepackGreaterThan(context global.DataContext, in GreaterThan) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in json.Number) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker (json.Number)
		if in == "" {
			return value, false, true, nil
		}
		return in, false, false, nil
	}(context, (json.Number)(in))
}
func (p packageType) RepackGreaterThanOrEqual(context global.DataContext, in GreaterThanOrEqual) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in json.Number) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker (json.Number)
		if in == "" {
			return value, false, true, nil
		}
		return in, false, false, nil
	}(context, (json.Number)(in))
}
func (p packageType) RepackIsNull(context global.DataContext, in IsNull) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct{}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 1)
		empty := true
		return out, false, empty, nil
	}(context, (struct{})(in))
}
func (p packageType) RepackItems(context global.DataContext, in Items) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
		// sliceRepacker
		out := make([]interface{}, len(in))
		empty := true
		for i, item := range in {
			v, _, null, err := func(context global.DataContext, in common.Validator) (value interface{}, dict bool, null bool, err error) {
				// selectorRepacker
				p := context.Package().Get("frizz.io/common")
				if p != nil {
					out, dict, null, err := p.Repack(context, in, "Validator")
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, null, nil
				}
				var vi interface{} = &in
				if vu, ok := vi.(json.Marshaler); ok {
					b, err := vu.MarshalJSON()
					if err != nil {
						return nil, false, false, err
					}
					if err := json.Unmarshal(b, &value); err != nil {
						return nil, false, false, err
					}
					return value, b[0] == '{', len(b) == 4 && string(b) == "null", nil
				}
				return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/common", "Validator")
			}(context, item)
			if err != nil {
				return nil, false, false, err
			}
			if !null {
				empty = false
			}
			out[i] = v
		}
		return out, false, empty, nil
	}(context, ([]common.Validator)(in))
}
func (p packageType) RepackKeys(context global.DataContext, in Keys) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
		// sliceRepacker
		out := make([]interface{}, len(in))
		empty := true
		for i, item := range in {
			v, _, null, err := func(context global.DataContext, in common.Validator) (value interface{}, dict bool, null bool, err error) {
				// selectorRepacker
				p := context.Package().Get("frizz.io/common")
				if p != nil {
					out, dict, null, err := p.Repack(context, in, "Validator")
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, null, nil
				}
				var vi interface{} = &in
				if vu, ok := vi.(json.Marshaler); ok {
					b, err := vu.MarshalJSON()
					if err != nil {
						return nil, false, false, err
					}
					if err := json.Unmarshal(b, &value); err != nil {
						return nil, false, false, err
					}
					return value, b[0] == '{', len(b) == 4 && string(b) == "null", nil
				}
				return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/common", "Validator")
			}(context, item)
			if err != nil {
				return nil, false, false, err
			}
			if !null {
				empty = false
			}
			out[i] = v
		}
		return out, false, empty, nil
	}(context, ([]common.Validator)(in))
}
func (p packageType) RepackLength(context global.DataContext, in Length) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
		// sliceRepacker
		out := make([]interface{}, len(in))
		empty := true
		for i, item := range in {
			v, _, null, err := func(context global.DataContext, in common.Validator) (value interface{}, dict bool, null bool, err error) {
				// selectorRepacker
				p := context.Package().Get("frizz.io/common")
				if p != nil {
					out, dict, null, err := p.Repack(context, in, "Validator")
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, null, nil
				}
				var vi interface{} = &in
				if vu, ok := vi.(json.Marshaler); ok {
					b, err := vu.MarshalJSON()
					if err != nil {
						return nil, false, false, err
					}
					if err := json.Unmarshal(b, &value); err != nil {
						return nil, false, false, err
					}
					return value, b[0] == '{', len(b) == 4 && string(b) == "null", nil
				}
				return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/common", "Validator")
			}(context, item)
			if err != nil {
				return nil, false, false, err
			}
			if !null {
				empty = false
			}
			out[i] = v
		}
		return out, false, empty, nil
	}(context, ([]common.Validator)(in))
}
func (p packageType) RepackLessThan(context global.DataContext, in LessThan) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in json.Number) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker (json.Number)
		if in == "" {
			return value, false, true, nil
		}
		return in, false, false, nil
	}(context, (json.Number)(in))
}
func (p packageType) RepackLessThanOrEqual(context global.DataContext, in LessThanOrEqual) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in json.Number) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker (json.Number)
		if in == "" {
			return value, false, true, nil
		}
		return in, false, false, nil
	}(context, (json.Number)(in))
}
func (p packageType) RepackNotNull(context global.DataContext, in NotNull) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct{}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 1)
		empty := true
		return out, false, empty, nil
	}(context, (struct{})(in))
}
func (p packageType) RepackRegex(context global.DataContext, in Regex) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Regex  string
		Invert bool
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackString(in)
		}(context, in.Regex); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Regex"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in bool) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackBool(in)
		}(context, in.Invert); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Invert"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Regex  string
		Invert bool
	})(in))
}
func (p packageType) RepackStruct(context global.DataContext, in Struct) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in map[string][]common.Validator) (value interface{}, dict bool, null bool, err error) {
		// mapRepacker
		out := make(map[string]interface{}, len(in))
		for k, item := range in {
			v, _, _, err := func(context global.DataContext, in []common.Validator) (value interface{}, dict bool, null bool, err error) {
				// sliceRepacker
				out := make([]interface{}, len(in))
				empty := true
				for i, item := range in {
					v, _, null, err := func(context global.DataContext, in common.Validator) (value interface{}, dict bool, null bool, err error) {
						// selectorRepacker
						p := context.Package().Get("frizz.io/common")
						if p != nil {
							out, dict, null, err := p.Repack(context, in, "Validator")
							if err != nil {
								return nil, false, false, err
							}
							return out, dict, null, nil
						}
						var vi interface{} = &in
						if vu, ok := vi.(json.Marshaler); ok {
							b, err := vu.MarshalJSON()
							if err != nil {
								return nil, false, false, err
							}
							if err := json.Unmarshal(b, &value); err != nil {
								return nil, false, false, err
							}
							return value, b[0] == '{', len(b) == 4 && string(b) == "null", nil
						}
						return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/common", "Validator")
					}(context, item)
					if err != nil {
						return nil, false, false, err
					}
					if !null {
						empty = false
					}
					out[i] = v
				}
				return out, false, empty, nil
			}(context, item)
			if err != nil {
				return nil, false, false, err
			}
			out[k] = v
		}
		return out, true, len(in) == 0, nil
	}(context, (map[string][]common.Validator)(in))
}
func (p packageType) GetData(filename string) string {
	return ""
}
func (p packageType) GetType(name string) string {
	return ""
}
func (p packageType) GetImportedPackages(packages map[string]global.Package) {
	packages["frizz.io/validators"] = Package
	common.Package.GetImportedPackages(packages)
	system.Package.GetImportedPackages(packages)
}
