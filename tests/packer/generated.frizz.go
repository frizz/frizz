package packer

import (
	json "encoding/json"
	global "frizz.io/global"
	pack "frizz.io/pack"
	silent "frizz.io/tests/packer/silent"
	sub "frizz.io/tests/packer/sub"
	errors "github.com/pkg/errors"
	time "time"
)

const Package packageType = 0

type packageType int

func (p packageType) Path() string {
	return "frizz.io/tests/packer"
}
func (p packageType) Unpack(context global.DataContext, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Ages":
		return p.UnpackAges(context, in)
	case "Alias":
		return p.UnpackAlias(context, in)
	case "AliasArray":
		return p.UnpackAliasArray(context, in)
	case "AliasMap":
		return p.UnpackAliasMap(context, in)
	case "AliasPointer":
		return p.UnpackAliasPointer(context, in)
	case "AliasSlice":
		return p.UnpackAliasSlice(context, in)
	case "AliasSub":
		return p.UnpackAliasSub(context, in)
	case "Csv":
		return p.UnpackCsv(context, in)
	case "Custom":
		return p.UnpackCustom(context, in)
	case "CustomSub":
		return p.UnpackCustomSub(context, in)
	case "EmbedNatives":
		return p.UnpackEmbedNatives(context, in)
	case "EmbedPointer":
		return p.UnpackEmbedPointer(context, in)
	case "EmbedQual":
		return p.UnpackEmbedQual(context, in)
	case "EmbedQualPointer":
		return p.UnpackEmbedQualPointer(context, in)
	case "Float64":
		return p.UnpackFloat64(context, in)
	case "HasTime":
		return p.UnpackHasTime(context, in)
	case "ImpSilent":
		return p.UnpackImpSilent(context, in)
	case "Impi":
		return p.UnpackImpi(context, in)
	case "Imps":
		return p.UnpackImps(context, in)
	case "Int":
		return p.UnpackInt(context, in)
	case "Interface":
		return p.UnpackInterface(context, in)
	case "InterfaceField":
		return p.UnpackInterfaceField(context, in)
	case "InterfaceSlice":
		return p.UnpackInterfaceSlice(context, in)
	case "InterfaceValidator":
		return p.UnpackInterfaceValidator(context, in)
	case "Maps":
		return p.UnpackMaps(context, in)
	case "Natives":
		return p.UnpackNatives(context, in)
	case "Pointers":
		return p.UnpackPointers(context, in)
	case "Private":
		return p.UnpackPrivate(context, in)
	case "Qual":
		return p.UnpackQual(context, in)
	case "Silent":
		return p.UnpackSilent(context, in)
	case "Slices":
		return p.UnpackSlices(context, in)
	case "String":
		return p.UnpackString(context, in)
	case "Structs":
		return p.UnpackStructs(context, in)
	case "SubInterface":
		return p.UnpackSubInterface(context, in)
	case "SubInterfaceSlice":
		return p.UnpackSubInterfaceSlice(context, in)
	case "SubMap":
		return p.UnpackSubMap(context, in)
	case "SubSlice":
		return p.UnpackSubSlice(context, in)
	case "Type":
		return p.UnpackType(context, in)
	case "Uint":
		return p.UnpackUint(context, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) UnpackAges(context global.DataContext, in interface{}) (value Ages, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// customUnpacker
	out := new(Ages)
	null, err = out.Unpack(context, in)
	if err != nil || null {
		return value, null, err
	}
	return *out, false, nil
}
func (p packageType) UnpackAlias(context global.DataContext, in interface{}) (value Alias, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value Int, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// localUnpacker
		out, null, err := p.UnpackInt(context, in)
		if err != nil || null {
			return value, null, err
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Alias(out), false, nil
}
func (p packageType) UnpackAliasArray(context global.DataContext, in interface{}) (value AliasArray, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value [3]string, null bool, err error) {
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
		var out [3]string
		if len(a) > 3 {
			return value, false, errors.Errorf("%s: data length %d does not fit in array of length %d", context.Location(), len(a), 3)
		}
		for i, v := range a {
			context := context.New(context.Location().Child(global.ArrayItem(i)))
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
				out[i] = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return AliasArray(out), false, nil
}
func (p packageType) UnpackAliasMap(context global.DataContext, in interface{}) (value AliasMap, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value map[string]*Qual, null bool, err error) {
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
		var out = make(map[string]*Qual, len(m))
		for k, v := range m {
			context := context.New(context.Location().Child(global.MapItem(k)))
			u, null, err := func(context global.DataContext, in interface{}) (value *Qual, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(context global.DataContext, in interface{}) (value Qual, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// localUnpacker
					out, null, err := p.UnpackQual(context, in)
					if err != nil || null {
						return value, null, err
					}
					return out, false, nil
				}(context, in)
				if err != nil || null {
					return value, null, err
				}
				return &out, false, nil
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
	return AliasMap(out), false, nil
}
func (p packageType) UnpackAliasPointer(context global.DataContext, in interface{}) (value AliasPointer, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value *Int, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// pointerUnpacker
		out, null, err := func(context global.DataContext, in interface{}) (value Int, null bool, err error) {
			if in == nil {
				return value, true, nil
			}
			// localUnpacker
			out, null, err := p.UnpackInt(context, in)
			if err != nil || null {
				return value, null, err
			}
			return out, false, nil
		}(context, in)
		if err != nil || null {
			return value, null, err
		}
		return &out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return AliasPointer(out), false, nil
}
func (p packageType) UnpackAliasSlice(context global.DataContext, in interface{}) (value AliasSlice, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value []Int, null bool, err error) {
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
		var out = make([]Int, len(a))
		for i, v := range a {
			context := context.New(context.Location().Child(global.ArrayItem(i)))
			u, null, err := func(context global.DataContext, in interface{}) (value Int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// localUnpacker
				out, null, err := p.UnpackInt(context, in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
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
	return AliasSlice(out), false, nil
}
func (p packageType) UnpackAliasSub(context global.DataContext, in interface{}) (value AliasSub, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value sub.Sub, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker
		p := context.Package().Get("frizz.io/tests/packer/sub")
		if p != nil {
			out, null, err := p.Unpack(context, in, "Sub")
			if err != nil || null {
				return value, null, err
			}
			return out.(sub.Sub), false, nil
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
		return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return AliasSub(out), false, nil
}
func (p packageType) UnpackCsv(context global.DataContext, in interface{}) (value Csv, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// customUnpacker
	out := new(Csv)
	null, err = out.Unpack(context, in)
	if err != nil || null {
		return value, null, err
	}
	return *out, false, nil
}
func (p packageType) UnpackCustom(context global.DataContext, in interface{}) (value Custom, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// customUnpacker
	out := new(Custom)
	null, err = out.Unpack(context, in)
	if err != nil || null {
		return value, null, err
	}
	return *out, false, nil
}
func (p packageType) UnpackCustomSub(context global.DataContext, in interface{}) (value CustomSub, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// customUnpacker
	out := new(CustomSub)
	null, err = out.Unpack(context, in)
	if err != nil || null {
		return value, null, err
	}
	return *out, false, nil
}
func (p packageType) UnpackEmbedNatives(context global.DataContext, in interface{}) (value EmbedNatives, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Natives
		Int int
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
			Natives
			Int int
		}
		if v, ok := m["Natives"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Natives")))
			u, null, err := func(context global.DataContext, in interface{}) (value Natives, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// localUnpacker
				out, null, err := p.UnpackNatives(context, in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Natives = u
			}
		}
		if v, ok := m["Int"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Int")))
			u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackInt(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return EmbedNatives(out), false, nil
}
func (p packageType) UnpackEmbedPointer(context global.DataContext, in interface{}) (value EmbedPointer, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		*Natives
		Int int
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
			*Natives
			Int int
		}
		if v, ok := m["Natives"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Natives")))
			u, null, err := func(context global.DataContext, in interface{}) (value *Natives, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(context global.DataContext, in interface{}) (value Natives, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// localUnpacker
					out, null, err := p.UnpackNatives(context, in)
					if err != nil || null {
						return value, null, err
					}
					return out, false, nil
				}(context, in)
				if err != nil || null {
					return value, null, err
				}
				return &out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Natives = u
			}
		}
		if v, ok := m["Int"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Int")))
			u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackInt(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return EmbedPointer(out), false, nil
}
func (p packageType) UnpackEmbedQual(context global.DataContext, in interface{}) (value EmbedQual, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		sub.Sub
		Int int
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
			sub.Sub
			Int int
		}
		if v, ok := m["Sub"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Sub")))
			u, null, err := func(context global.DataContext, in interface{}) (value sub.Sub, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// selectorUnpacker
				p := context.Package().Get("frizz.io/tests/packer/sub")
				if p != nil {
					out, null, err := p.Unpack(context, in, "Sub")
					if err != nil || null {
						return value, null, err
					}
					return out.(sub.Sub), false, nil
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
				return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Sub = u
			}
		}
		if v, ok := m["Int"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Int")))
			u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackInt(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return EmbedQual(out), false, nil
}
func (p packageType) UnpackEmbedQualPointer(context global.DataContext, in interface{}) (value EmbedQualPointer, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		*sub.Sub
		Int int
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
			*sub.Sub
			Int int
		}
		if v, ok := m["Sub"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Sub")))
			u, null, err := func(context global.DataContext, in interface{}) (value *sub.Sub, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(context global.DataContext, in interface{}) (value sub.Sub, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// selectorUnpacker
					p := context.Package().Get("frizz.io/tests/packer/sub")
					if p != nil {
						out, null, err := p.Unpack(context, in, "Sub")
						if err != nil || null {
							return value, null, err
						}
						return out.(sub.Sub), false, nil
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
					return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
				}(context, in)
				if err != nil || null {
					return value, null, err
				}
				return &out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Sub = u
			}
		}
		if v, ok := m["Int"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Int")))
			u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackInt(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return EmbedQualPointer(out), false, nil
}
func (p packageType) UnpackFloat64(context global.DataContext, in interface{}) (value Float64, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value float64, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// nativeUnpacker
		out, null, err := pack.UnpackFloat64(context.Location(), in)
		if err != nil || null {
			return value, null, err
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Float64(out), false, nil
}
func (p packageType) UnpackHasTime(context global.DataContext, in interface{}) (value HasTime, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Time time.Time
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
			Time time.Time
		}
		if v, ok := m["Time"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Time")))
			u, null, err := func(context global.DataContext, in interface{}) (value time.Time, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// selectorUnpacker
				p := context.Package().Get("time")
				if p != nil {
					out, null, err := p.Unpack(context, in, "Time")
					if err != nil || null {
						return value, null, err
					}
					return out.(time.Time), false, nil
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
				return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "time", "Time")
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Time = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return HasTime(out), false, nil
}
func (p packageType) UnpackImpSilent(context global.DataContext, in interface{}) (value ImpSilent, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// nativeUnpacker
		out, null, err := pack.UnpackInt(context.Location(), in)
		if err != nil || null {
			return value, null, err
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return ImpSilent(out), false, nil
}
func (p packageType) UnpackImpi(context global.DataContext, in interface{}) (value Impi, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Int int
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
			Int int
		}
		if v, ok := m["Int"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Int")))
			u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackInt(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Impi(out), false, nil
}
func (p packageType) UnpackImps(context global.DataContext, in interface{}) (value Imps, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		String string
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
			String string
		}
		if v, ok := m["String"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("String")))
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
				out.String = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Imps(out), false, nil
}
func (p packageType) UnpackInt(context global.DataContext, in interface{}) (value Int, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// nativeUnpacker
		out, null, err := pack.UnpackInt(context.Location(), in)
		if err != nil || null {
			return value, null, err
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Int(out), false, nil
}
func (p packageType) UnpackInterface(context global.DataContext, in interface{}) (value Interface, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value interface {
		Foo() string
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// interfaceUnpacker
		out, null, err := pack.UnpackInterface(context, in)
		if err != nil {
			return value, false, err
		}
		iface, ok := out.(interface {
			Foo() string
		})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into interface, type %T does not implement interface", context.Location(), out)
		}
		if null {
			return value, true, nil
		}
		return iface, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Interface(out), false, nil
}
func (p packageType) UnpackInterfaceField(context global.DataContext, in interface{}) (value InterfaceField, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Iface Interface
		Slice []Interface
		Array [3]Interface
		Map   map[string]Interface
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
			Iface Interface
			Slice []Interface
			Array [3]Interface
			Map   map[string]Interface
		}
		if v, ok := m["Iface"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Iface")))
			u, null, err := func(context global.DataContext, in interface{}) (value Interface, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// localUnpacker
				out, null, err := p.UnpackInterface(context, in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Iface = u
			}
		}
		if v, ok := m["Slice"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Slice")))
			u, null, err := func(context global.DataContext, in interface{}) (value []Interface, null bool, err error) {
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
				var out = make([]Interface, len(a))
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
					u, null, err := func(context global.DataContext, in interface{}) (value Interface, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// localUnpacker
						out, null, err := p.UnpackInterface(context, in)
						if err != nil || null {
							return value, null, err
						}
						return out, false, nil
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
				out.Slice = u
			}
		}
		if v, ok := m["Array"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Array")))
			u, null, err := func(context global.DataContext, in interface{}) (value [3]Interface, null bool, err error) {
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
				var out [3]Interface
				if len(a) > 3 {
					return value, false, errors.Errorf("%s: data length %d does not fit in array of length %d", context.Location(), len(a), 3)
				}
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
					u, null, err := func(context global.DataContext, in interface{}) (value Interface, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// localUnpacker
						out, null, err := p.UnpackInterface(context, in)
						if err != nil || null {
							return value, null, err
						}
						return out, false, nil
					}(context, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out[i] = u
					}
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Array = u
			}
		}
		if v, ok := m["Map"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Map")))
			u, null, err := func(context global.DataContext, in interface{}) (value map[string]Interface, null bool, err error) {
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
				var out = make(map[string]Interface, len(m))
				for k, v := range m {
					context := context.New(context.Location().Child(global.MapItem(k)))
					u, null, err := func(context global.DataContext, in interface{}) (value Interface, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// localUnpacker
						out, null, err := p.UnpackInterface(context, in)
						if err != nil || null {
							return value, null, err
						}
						return out, false, nil
					}(context, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out[k] = u
					}
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Map = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return InterfaceField(out), false, nil
}
func (p packageType) UnpackInterfaceSlice(context global.DataContext, in interface{}) (value InterfaceSlice, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value []Interface, null bool, err error) {
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
		var out = make([]Interface, len(a))
		for i, v := range a {
			context := context.New(context.Location().Child(global.ArrayItem(i)))
			u, null, err := func(context global.DataContext, in interface{}) (value Interface, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// localUnpacker
				out, null, err := p.UnpackInterface(context, in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
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
	return InterfaceSlice(out), false, nil
}
func (p packageType) UnpackInterfaceValidator(context global.DataContext, in interface{}) (value InterfaceValidator, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value string, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// nativeUnpacker
		out, null, err := pack.UnpackString(context.Location(), in)
		if err != nil || null {
			return value, null, err
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return InterfaceValidator(out), false, nil
}
func (p packageType) UnpackMaps(context global.DataContext, in interface{}) (value Maps, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Ints    map[string]int
		Strings map[string]string
		Slices  map[string][]string
		Arrays  map[string][2]int
		Maps    map[string]map[string]string
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
			Ints    map[string]int
			Strings map[string]string
			Slices  map[string][]string
			Arrays  map[string][2]int
			Maps    map[string]map[string]string
		}
		if v, ok := m["Ints"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Ints")))
			u, null, err := func(context global.DataContext, in interface{}) (value map[string]int, null bool, err error) {
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
				var out = make(map[string]int, len(m))
				for k, v := range m {
					context := context.New(context.Location().Child(global.MapItem(k)))
					u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// nativeUnpacker
						out, null, err := pack.UnpackInt(context.Location(), in)
						if err != nil || null {
							return value, null, err
						}
						return out, false, nil
					}(context, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out[k] = u
					}
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Ints = u
			}
		}
		if v, ok := m["Strings"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Strings")))
			u, null, err := func(context global.DataContext, in interface{}) (value map[string]string, null bool, err error) {
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
				var out = make(map[string]string, len(m))
				for k, v := range m {
					context := context.New(context.Location().Child(global.MapItem(k)))
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
						out[k] = u
					}
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Strings = u
			}
		}
		if v, ok := m["Slices"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Slices")))
			u, null, err := func(context global.DataContext, in interface{}) (value map[string][]string, null bool, err error) {
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
				var out = make(map[string][]string, len(m))
				for k, v := range m {
					context := context.New(context.Location().Child(global.MapItem(k)))
					u, null, err := func(context global.DataContext, in interface{}) (value []string, null bool, err error) {
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
						var out = make([]string, len(a))
						for i, v := range a {
							context := context.New(context.Location().Child(global.ArrayItem(i)))
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
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Slices = u
			}
		}
		if v, ok := m["Arrays"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Arrays")))
			u, null, err := func(context global.DataContext, in interface{}) (value map[string][2]int, null bool, err error) {
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
				var out = make(map[string][2]int, len(m))
				for k, v := range m {
					context := context.New(context.Location().Child(global.MapItem(k)))
					u, null, err := func(context global.DataContext, in interface{}) (value [2]int, null bool, err error) {
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
						var out [2]int
						if len(a) > 2 {
							return value, false, errors.Errorf("%s: data length %d does not fit in array of length %d", context.Location(), len(a), 2)
						}
						for i, v := range a {
							context := context.New(context.Location().Child(global.ArrayItem(i)))
							u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
								if in == nil {
									return value, true, nil
								}
								// nativeUnpacker
								out, null, err := pack.UnpackInt(context.Location(), in)
								if err != nil || null {
									return value, null, err
								}
								return out, false, nil
							}(context, v)
							if err != nil {
								return value, false, err
							}
							if !null {
								out[i] = u
							}
						}
						return out, false, nil
					}(context, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out[k] = u
					}
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Arrays = u
			}
		}
		if v, ok := m["Maps"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Maps")))
			u, null, err := func(context global.DataContext, in interface{}) (value map[string]map[string]string, null bool, err error) {
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
				var out = make(map[string]map[string]string, len(m))
				for k, v := range m {
					context := context.New(context.Location().Child(global.MapItem(k)))
					u, null, err := func(context global.DataContext, in interface{}) (value map[string]string, null bool, err error) {
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
						var out = make(map[string]string, len(m))
						for k, v := range m {
							context := context.New(context.Location().Child(global.MapItem(k)))
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
								out[k] = u
							}
						}
						return out, false, nil
					}(context, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out[k] = u
					}
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Maps = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Maps(out), false, nil
}
func (p packageType) UnpackNatives(context global.DataContext, in interface{}) (value Natives, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Bool      bool
		Byte      byte
		Float32   float32
		Float64   float64
		Int       int
		Int8      int8
		Int16     int16
		Int32     int32
		Int64     int64
		Uint      uint
		Uint8     uint8
		Uint16    uint16
		Uint32    uint32
		Uint64    uint64
		Rune      rune
		String    string
		PtrString *string
		PtrInt    *int
		PtrBool   *bool
		Number    json.Number
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
			Bool      bool
			Byte      byte
			Float32   float32
			Float64   float64
			Int       int
			Int8      int8
			Int16     int16
			Int32     int32
			Int64     int64
			Uint      uint
			Uint8     uint8
			Uint16    uint16
			Uint32    uint32
			Uint64    uint64
			Rune      rune
			String    string
			PtrString *string
			PtrInt    *int
			PtrBool   *bool
			Number    json.Number
		}
		if v, ok := m["Bool"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Bool")))
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
				out.Bool = u
			}
		}
		if v, ok := m["Byte"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Byte")))
			u, null, err := func(context global.DataContext, in interface{}) (value byte, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackByte(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Byte = u
			}
		}
		if v, ok := m["Float32"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Float32")))
			u, null, err := func(context global.DataContext, in interface{}) (value float32, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackFloat32(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Float32 = u
			}
		}
		if v, ok := m["Float64"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Float64")))
			u, null, err := func(context global.DataContext, in interface{}) (value float64, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackFloat64(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Float64 = u
			}
		}
		if v, ok := m["Int"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Int")))
			u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackInt(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int = u
			}
		}
		if v, ok := m["Int8"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Int8")))
			u, null, err := func(context global.DataContext, in interface{}) (value int8, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackInt8(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int8 = u
			}
		}
		if v, ok := m["Int16"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Int16")))
			u, null, err := func(context global.DataContext, in interface{}) (value int16, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackInt16(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int16 = u
			}
		}
		if v, ok := m["Int32"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Int32")))
			u, null, err := func(context global.DataContext, in interface{}) (value int32, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackInt32(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int32 = u
			}
		}
		if v, ok := m["Int64"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Int64")))
			u, null, err := func(context global.DataContext, in interface{}) (value int64, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackInt64(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int64 = u
			}
		}
		if v, ok := m["Uint"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Uint")))
			u, null, err := func(context global.DataContext, in interface{}) (value uint, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackUint(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Uint = u
			}
		}
		if v, ok := m["Uint8"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Uint8")))
			u, null, err := func(context global.DataContext, in interface{}) (value uint8, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackUint8(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Uint8 = u
			}
		}
		if v, ok := m["Uint16"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Uint16")))
			u, null, err := func(context global.DataContext, in interface{}) (value uint16, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackUint16(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Uint16 = u
			}
		}
		if v, ok := m["Uint32"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Uint32")))
			u, null, err := func(context global.DataContext, in interface{}) (value uint32, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackUint32(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Uint32 = u
			}
		}
		if v, ok := m["Uint64"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Uint64")))
			u, null, err := func(context global.DataContext, in interface{}) (value uint64, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackUint64(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Uint64 = u
			}
		}
		if v, ok := m["Rune"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Rune")))
			u, null, err := func(context global.DataContext, in interface{}) (value rune, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackRune(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Rune = u
			}
		}
		if v, ok := m["String"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("String")))
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
				out.String = u
			}
		}
		if v, ok := m["PtrString"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("PtrString")))
			u, null, err := func(context global.DataContext, in interface{}) (value *string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(context global.DataContext, in interface{}) (value string, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// nativeUnpacker
					out, null, err := pack.UnpackString(context.Location(), in)
					if err != nil || null {
						return value, null, err
					}
					return out, false, nil
				}(context, in)
				if err != nil || null {
					return value, null, err
				}
				return &out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.PtrString = u
			}
		}
		if v, ok := m["PtrInt"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("PtrInt")))
			u, null, err := func(context global.DataContext, in interface{}) (value *int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// nativeUnpacker
					out, null, err := pack.UnpackInt(context.Location(), in)
					if err != nil || null {
						return value, null, err
					}
					return out, false, nil
				}(context, in)
				if err != nil || null {
					return value, null, err
				}
				return &out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.PtrInt = u
			}
		}
		if v, ok := m["PtrBool"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("PtrBool")))
			u, null, err := func(context global.DataContext, in interface{}) (value *bool, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(context global.DataContext, in interface{}) (value bool, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// nativeUnpacker
					out, null, err := pack.UnpackBool(context.Location(), in)
					if err != nil || null {
						return value, null, err
					}
					return out, false, nil
				}(context, in)
				if err != nil || null {
					return value, null, err
				}
				return &out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.PtrBool = u
			}
		}
		if v, ok := m["Number"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Number")))
			u, null, err := func(context global.DataContext, in interface{}) (value json.Number, null bool, err error) {
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
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Number = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Natives(out), false, nil
}
func (p packageType) UnpackPointers(context global.DataContext, in interface{}) (value Pointers, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		String      *string
		Int         *Int
		Sub         *sub.Sub
		Array       *[3]int
		Slice       *[]string
		Map         *map[string]int
		SliceString []*string
		SliceInt    []*Int
		SliceSub    []*sub.Sub
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
			String      *string
			Int         *Int
			Sub         *sub.Sub
			Array       *[3]int
			Slice       *[]string
			Map         *map[string]int
			SliceString []*string
			SliceInt    []*Int
			SliceSub    []*sub.Sub
		}
		if v, ok := m["String"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("String")))
			u, null, err := func(context global.DataContext, in interface{}) (value *string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(context global.DataContext, in interface{}) (value string, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// nativeUnpacker
					out, null, err := pack.UnpackString(context.Location(), in)
					if err != nil || null {
						return value, null, err
					}
					return out, false, nil
				}(context, in)
				if err != nil || null {
					return value, null, err
				}
				return &out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.String = u
			}
		}
		if v, ok := m["Int"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Int")))
			u, null, err := func(context global.DataContext, in interface{}) (value *Int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(context global.DataContext, in interface{}) (value Int, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// localUnpacker
					out, null, err := p.UnpackInt(context, in)
					if err != nil || null {
						return value, null, err
					}
					return out, false, nil
				}(context, in)
				if err != nil || null {
					return value, null, err
				}
				return &out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int = u
			}
		}
		if v, ok := m["Sub"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Sub")))
			u, null, err := func(context global.DataContext, in interface{}) (value *sub.Sub, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(context global.DataContext, in interface{}) (value sub.Sub, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// selectorUnpacker
					p := context.Package().Get("frizz.io/tests/packer/sub")
					if p != nil {
						out, null, err := p.Unpack(context, in, "Sub")
						if err != nil || null {
							return value, null, err
						}
						return out.(sub.Sub), false, nil
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
					return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
				}(context, in)
				if err != nil || null {
					return value, null, err
				}
				return &out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Sub = u
			}
		}
		if v, ok := m["Array"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Array")))
			u, null, err := func(context global.DataContext, in interface{}) (value *[3]int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(context global.DataContext, in interface{}) (value [3]int, null bool, err error) {
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
					var out [3]int
					if len(a) > 3 {
						return value, false, errors.Errorf("%s: data length %d does not fit in array of length %d", context.Location(), len(a), 3)
					}
					for i, v := range a {
						context := context.New(context.Location().Child(global.ArrayItem(i)))
						u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
							if in == nil {
								return value, true, nil
							}
							// nativeUnpacker
							out, null, err := pack.UnpackInt(context.Location(), in)
							if err != nil || null {
								return value, null, err
							}
							return out, false, nil
						}(context, v)
						if err != nil {
							return value, false, err
						}
						if !null {
							out[i] = u
						}
					}
					return out, false, nil
				}(context, in)
				if err != nil || null {
					return value, null, err
				}
				return &out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Array = u
			}
		}
		if v, ok := m["Slice"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Slice")))
			u, null, err := func(context global.DataContext, in interface{}) (value *[]string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(context global.DataContext, in interface{}) (value []string, null bool, err error) {
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
					var out = make([]string, len(a))
					for i, v := range a {
						context := context.New(context.Location().Child(global.ArrayItem(i)))
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
							out[i] = u
						}
					}
					return out[:], false, nil
				}(context, in)
				if err != nil || null {
					return value, null, err
				}
				return &out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Slice = u
			}
		}
		if v, ok := m["Map"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Map")))
			u, null, err := func(context global.DataContext, in interface{}) (value *map[string]int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(context global.DataContext, in interface{}) (value map[string]int, null bool, err error) {
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
					var out = make(map[string]int, len(m))
					for k, v := range m {
						context := context.New(context.Location().Child(global.MapItem(k)))
						u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
							if in == nil {
								return value, true, nil
							}
							// nativeUnpacker
							out, null, err := pack.UnpackInt(context.Location(), in)
							if err != nil || null {
								return value, null, err
							}
							return out, false, nil
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
				return &out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Map = u
			}
		}
		if v, ok := m["SliceString"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("SliceString")))
			u, null, err := func(context global.DataContext, in interface{}) (value []*string, null bool, err error) {
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
				var out = make([]*string, len(a))
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
					u, null, err := func(context global.DataContext, in interface{}) (value *string, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// pointerUnpacker
						out, null, err := func(context global.DataContext, in interface{}) (value string, null bool, err error) {
							if in == nil {
								return value, true, nil
							}
							// nativeUnpacker
							out, null, err := pack.UnpackString(context.Location(), in)
							if err != nil || null {
								return value, null, err
							}
							return out, false, nil
						}(context, in)
						if err != nil || null {
							return value, null, err
						}
						return &out, false, nil
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
				out.SliceString = u
			}
		}
		if v, ok := m["SliceInt"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("SliceInt")))
			u, null, err := func(context global.DataContext, in interface{}) (value []*Int, null bool, err error) {
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
				var out = make([]*Int, len(a))
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
					u, null, err := func(context global.DataContext, in interface{}) (value *Int, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// pointerUnpacker
						out, null, err := func(context global.DataContext, in interface{}) (value Int, null bool, err error) {
							if in == nil {
								return value, true, nil
							}
							// localUnpacker
							out, null, err := p.UnpackInt(context, in)
							if err != nil || null {
								return value, null, err
							}
							return out, false, nil
						}(context, in)
						if err != nil || null {
							return value, null, err
						}
						return &out, false, nil
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
				out.SliceInt = u
			}
		}
		if v, ok := m["SliceSub"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("SliceSub")))
			u, null, err := func(context global.DataContext, in interface{}) (value []*sub.Sub, null bool, err error) {
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
				var out = make([]*sub.Sub, len(a))
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
					u, null, err := func(context global.DataContext, in interface{}) (value *sub.Sub, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// pointerUnpacker
						out, null, err := func(context global.DataContext, in interface{}) (value sub.Sub, null bool, err error) {
							if in == nil {
								return value, true, nil
							}
							// selectorUnpacker
							p := context.Package().Get("frizz.io/tests/packer/sub")
							if p != nil {
								out, null, err := p.Unpack(context, in, "Sub")
								if err != nil || null {
									return value, null, err
								}
								return out.(sub.Sub), false, nil
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
							return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
						}(context, in)
						if err != nil || null {
							return value, null, err
						}
						return &out, false, nil
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
				out.SliceSub = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Pointers(out), false, nil
}
func (p packageType) UnpackPrivate(context global.DataContext, in interface{}) (value Private, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		i int
		s string
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
			i int
			s string
		}
		if v, ok := m["i"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("i")))
			u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackInt(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.i = u
			}
		}
		if v, ok := m["s"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("s")))
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
				out.s = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Private(out), false, nil
}
func (p packageType) UnpackQual(context global.DataContext, in interface{}) (value Qual, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Sub sub.Sub
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
			Sub sub.Sub
		}
		if v, ok := m["Sub"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Sub")))
			u, null, err := func(context global.DataContext, in interface{}) (value sub.Sub, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// selectorUnpacker
				p := context.Package().Get("frizz.io/tests/packer/sub")
				if p != nil {
					out, null, err := p.Unpack(context, in, "Sub")
					if err != nil || null {
						return value, null, err
					}
					return out.(sub.Sub), false, nil
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
				return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Sub = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Qual(out), false, nil
}
func (p packageType) UnpackSilent(context global.DataContext, in interface{}) (value Silent, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Silent silent.Silent
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
			Silent silent.Silent
		}
		if v, ok := m["Silent"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Silent")))
			u, null, err := func(context global.DataContext, in interface{}) (value silent.Silent, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// selectorUnpacker
				p := context.Package().Get("frizz.io/tests/packer/silent")
				if p != nil {
					out, null, err := p.Unpack(context, in, "Silent")
					if err != nil || null {
						return value, null, err
					}
					return out.(silent.Silent), false, nil
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
				return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/tests/packer/silent", "Silent")
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Silent = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Silent(out), false, nil
}
func (p packageType) UnpackSlices(context global.DataContext, in interface{}) (value Slices, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Ints      []int
		Strings   []string
		ArrayLit  [5]string
		ArrayExpr [2 * N]int
		Structs   []struct {
			Int int
		}
		Arrays [][]string
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
			Ints      []int
			Strings   []string
			ArrayLit  [5]string
			ArrayExpr [2 * N]int
			Structs   []struct {
				Int int
			}
			Arrays [][]string
		}
		if v, ok := m["Ints"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Ints")))
			u, null, err := func(context global.DataContext, in interface{}) (value []int, null bool, err error) {
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
				var out = make([]int, len(a))
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
					u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// nativeUnpacker
						out, null, err := pack.UnpackInt(context.Location(), in)
						if err != nil || null {
							return value, null, err
						}
						return out, false, nil
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
				out.Ints = u
			}
		}
		if v, ok := m["Strings"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Strings")))
			u, null, err := func(context global.DataContext, in interface{}) (value []string, null bool, err error) {
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
				var out = make([]string, len(a))
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
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
						out[i] = u
					}
				}
				return out[:], false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Strings = u
			}
		}
		if v, ok := m["ArrayLit"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("ArrayLit")))
			u, null, err := func(context global.DataContext, in interface{}) (value [5]string, null bool, err error) {
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
				var out [5]string
				if len(a) > 5 {
					return value, false, errors.Errorf("%s: data length %d does not fit in array of length %d", context.Location(), len(a), 5)
				}
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
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
						out[i] = u
					}
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.ArrayLit = u
			}
		}
		if v, ok := m["ArrayExpr"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("ArrayExpr")))
			u, null, err := func(context global.DataContext, in interface{}) (value [2 * N]int, null bool, err error) {
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
				var out [2 * N]int
				if len(a) > 2*N {
					return value, false, errors.Errorf("%s: data length %d does not fit in array of length %d", context.Location(), len(a), 2*N)
				}
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
					u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// nativeUnpacker
						out, null, err := pack.UnpackInt(context.Location(), in)
						if err != nil || null {
							return value, null, err
						}
						return out, false, nil
					}(context, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out[i] = u
					}
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.ArrayExpr = u
			}
		}
		if v, ok := m["Structs"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Structs")))
			u, null, err := func(context global.DataContext, in interface{}) (value []struct {
				Int int
			}, null bool, err error) {
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
				var out = make([]struct {
					Int int
				}, len(a))
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
					u, null, err := func(context global.DataContext, in interface{}) (value struct {
						Int int
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
							Int int
						}
						if v, ok := m["Int"]; ok {
							context := context.New(context.Location().Child(global.FieldItem("Int")))
							u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
								if in == nil {
									return value, true, nil
								}
								// nativeUnpacker
								out, null, err := pack.UnpackInt(context.Location(), in)
								if err != nil || null {
									return value, null, err
								}
								return out, false, nil
							}(context, v)
							if err != nil {
								return value, false, err
							}
							if !null {
								out.Int = u
							}
						}
						return out, false, nil
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
				out.Structs = u
			}
		}
		if v, ok := m["Arrays"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Arrays")))
			u, null, err := func(context global.DataContext, in interface{}) (value [][]string, null bool, err error) {
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
				var out = make([][]string, len(a))
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
					u, null, err := func(context global.DataContext, in interface{}) (value []string, null bool, err error) {
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
						var out = make([]string, len(a))
						for i, v := range a {
							context := context.New(context.Location().Child(global.ArrayItem(i)))
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
								out[i] = u
							}
						}
						return out[:], false, nil
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
				out.Arrays = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Slices(out), false, nil
}
func (p packageType) UnpackString(context global.DataContext, in interface{}) (value String, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value string, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// nativeUnpacker
		out, null, err := pack.UnpackString(context.Location(), in)
		if err != nil || null {
			return value, null, err
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return String(out), false, nil
}
func (p packageType) UnpackStructs(context global.DataContext, in interface{}) (value Structs, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Simple struct {
			Int  int
			Bool bool
		}
		Complex struct {
			String string
			Inner  struct {
				Float32 float32
			}
		}
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
			Simple struct {
				Int  int
				Bool bool
			}
			Complex struct {
				String string
				Inner  struct {
					Float32 float32
				}
			}
		}
		if v, ok := m["Simple"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Simple")))
			u, null, err := func(context global.DataContext, in interface{}) (value struct {
				Int  int
				Bool bool
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
					Int  int
					Bool bool
				}
				if v, ok := m["Int"]; ok {
					context := context.New(context.Location().Child(global.FieldItem("Int")))
					u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// nativeUnpacker
						out, null, err := pack.UnpackInt(context.Location(), in)
						if err != nil || null {
							return value, null, err
						}
						return out, false, nil
					}(context, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out.Int = u
					}
				}
				if v, ok := m["Bool"]; ok {
					context := context.New(context.Location().Child(global.FieldItem("Bool")))
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
						out.Bool = u
					}
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Simple = u
			}
		}
		if v, ok := m["Complex"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Complex")))
			u, null, err := func(context global.DataContext, in interface{}) (value struct {
				String string
				Inner  struct {
					Float32 float32
				}
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
					String string
					Inner  struct {
						Float32 float32
					}
				}
				if v, ok := m["String"]; ok {
					context := context.New(context.Location().Child(global.FieldItem("String")))
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
						out.String = u
					}
				}
				if v, ok := m["Inner"]; ok {
					context := context.New(context.Location().Child(global.FieldItem("Inner")))
					u, null, err := func(context global.DataContext, in interface{}) (value struct {
						Float32 float32
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
							Float32 float32
						}
						if v, ok := m["Float32"]; ok {
							context := context.New(context.Location().Child(global.FieldItem("Float32")))
							u, null, err := func(context global.DataContext, in interface{}) (value float32, null bool, err error) {
								if in == nil {
									return value, true, nil
								}
								// nativeUnpacker
								out, null, err := pack.UnpackFloat32(context.Location(), in)
								if err != nil || null {
									return value, null, err
								}
								return out, false, nil
							}(context, v)
							if err != nil {
								return value, false, err
							}
							if !null {
								out.Float32 = u
							}
						}
						return out, false, nil
					}(context, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out.Inner = u
					}
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Complex = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Structs(out), false, nil
}
func (p packageType) UnpackSubInterface(context global.DataContext, in interface{}) (value SubInterface, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		SubInterface sub.SubInterface
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
			SubInterface sub.SubInterface
		}
		if v, ok := m["SubInterface"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("SubInterface")))
			u, null, err := func(context global.DataContext, in interface{}) (value sub.SubInterface, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// selectorUnpacker
				p := context.Package().Get("frizz.io/tests/packer/sub")
				if p != nil {
					out, null, err := p.Unpack(context, in, "SubInterface")
					if err != nil || null {
						return value, null, err
					}
					return out.(sub.SubInterface), false, nil
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
				return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "SubInterface")
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.SubInterface = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return SubInterface(out), false, nil
}
func (p packageType) UnpackSubInterfaceSlice(context global.DataContext, in interface{}) (value SubInterfaceSlice, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value []sub.SubInterface, null bool, err error) {
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
		var out = make([]sub.SubInterface, len(a))
		for i, v := range a {
			context := context.New(context.Location().Child(global.ArrayItem(i)))
			u, null, err := func(context global.DataContext, in interface{}) (value sub.SubInterface, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// selectorUnpacker
				p := context.Package().Get("frizz.io/tests/packer/sub")
				if p != nil {
					out, null, err := p.Unpack(context, in, "SubInterface")
					if err != nil || null {
						return value, null, err
					}
					return out.(sub.SubInterface), false, nil
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
				return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "SubInterface")
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
	return SubInterfaceSlice(out), false, nil
}
func (p packageType) UnpackSubMap(context global.DataContext, in interface{}) (value SubMap, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Map map[string]sub.Sub
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
			Map map[string]sub.Sub
		}
		if v, ok := m["Map"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Map")))
			u, null, err := func(context global.DataContext, in interface{}) (value map[string]sub.Sub, null bool, err error) {
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
				var out = make(map[string]sub.Sub, len(m))
				for k, v := range m {
					context := context.New(context.Location().Child(global.MapItem(k)))
					u, null, err := func(context global.DataContext, in interface{}) (value sub.Sub, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// selectorUnpacker
						p := context.Package().Get("frizz.io/tests/packer/sub")
						if p != nil {
							out, null, err := p.Unpack(context, in, "Sub")
							if err != nil || null {
								return value, null, err
							}
							return out.(sub.Sub), false, nil
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
						return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
					}(context, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out[k] = u
					}
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Map = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return SubMap(out), false, nil
}
func (p packageType) UnpackSubSlice(context global.DataContext, in interface{}) (value SubSlice, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Slice []sub.Sub
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
			Slice []sub.Sub
		}
		if v, ok := m["Slice"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Slice")))
			u, null, err := func(context global.DataContext, in interface{}) (value []sub.Sub, null bool, err error) {
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
				var out = make([]sub.Sub, len(a))
				for i, v := range a {
					context := context.New(context.Location().Child(global.ArrayItem(i)))
					u, null, err := func(context global.DataContext, in interface{}) (value sub.Sub, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// selectorUnpacker
						p := context.Package().Get("frizz.io/tests/packer/sub")
						if p != nil {
							out, null, err := p.Unpack(context, in, "Sub")
							if err != nil || null {
								return value, null, err
							}
							return out.(sub.Sub), false, nil
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
						return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
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
				out.Slice = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return SubSlice(out), false, nil
}
func (p packageType) UnpackType(context global.DataContext, in interface{}) (value Type, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// customUnpacker
	out := new(Type)
	null, err = out.Unpack(context, in)
	if err != nil || null {
		return value, null, err
	}
	return *out, false, nil
}
func (p packageType) UnpackUint(context global.DataContext, in interface{}) (value Uint, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value uint, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// nativeUnpacker
		out, null, err := pack.UnpackUint(context.Location(), in)
		if err != nil || null {
			return value, null, err
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Uint(out), false, nil
}
func (p packageType) Repack(context global.DataContext, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Ages":
		return p.RepackAges(context, in.(Ages))
	case "Alias":
		return p.RepackAlias(context, in.(Alias))
	case "AliasArray":
		return p.RepackAliasArray(context, in.(AliasArray))
	case "AliasMap":
		return p.RepackAliasMap(context, in.(AliasMap))
	case "AliasPointer":
		return p.RepackAliasPointer(context, in.(AliasPointer))
	case "AliasSlice":
		return p.RepackAliasSlice(context, in.(AliasSlice))
	case "AliasSub":
		return p.RepackAliasSub(context, in.(AliasSub))
	case "Csv":
		return p.RepackCsv(context, in.(Csv))
	case "Custom":
		return p.RepackCustom(context, in.(Custom))
	case "CustomSub":
		return p.RepackCustomSub(context, in.(CustomSub))
	case "EmbedNatives":
		return p.RepackEmbedNatives(context, in.(EmbedNatives))
	case "EmbedPointer":
		return p.RepackEmbedPointer(context, in.(EmbedPointer))
	case "EmbedQual":
		return p.RepackEmbedQual(context, in.(EmbedQual))
	case "EmbedQualPointer":
		return p.RepackEmbedQualPointer(context, in.(EmbedQualPointer))
	case "Float64":
		return p.RepackFloat64(context, in.(Float64))
	case "HasTime":
		return p.RepackHasTime(context, in.(HasTime))
	case "ImpSilent":
		return p.RepackImpSilent(context, in.(ImpSilent))
	case "Impi":
		return p.RepackImpi(context, in.(Impi))
	case "Imps":
		return p.RepackImps(context, in.(Imps))
	case "Int":
		return p.RepackInt(context, in.(Int))
	case "Interface":
		return p.RepackInterface(context, in.(Interface))
	case "InterfaceField":
		return p.RepackInterfaceField(context, in.(InterfaceField))
	case "InterfaceSlice":
		return p.RepackInterfaceSlice(context, in.(InterfaceSlice))
	case "InterfaceValidator":
		return p.RepackInterfaceValidator(context, in.(InterfaceValidator))
	case "Maps":
		return p.RepackMaps(context, in.(Maps))
	case "Natives":
		return p.RepackNatives(context, in.(Natives))
	case "Pointers":
		return p.RepackPointers(context, in.(Pointers))
	case "Private":
		return p.RepackPrivate(context, in.(Private))
	case "Qual":
		return p.RepackQual(context, in.(Qual))
	case "Silent":
		return p.RepackSilent(context, in.(Silent))
	case "Slices":
		return p.RepackSlices(context, in.(Slices))
	case "String":
		return p.RepackString(context, in.(String))
	case "Structs":
		return p.RepackStructs(context, in.(Structs))
	case "SubInterface":
		return p.RepackSubInterface(context, in.(SubInterface))
	case "SubInterfaceSlice":
		return p.RepackSubInterfaceSlice(context, in.(SubInterfaceSlice))
	case "SubMap":
		return p.RepackSubMap(context, in.(SubMap))
	case "SubSlice":
		return p.RepackSubSlice(context, in.(SubSlice))
	case "Type":
		return p.RepackType(context, in.(Type))
	case "Uint":
		return p.RepackUint(context, in.(Uint))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) RepackAges(context global.DataContext, in Ages) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(context)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
}
func (p packageType) RepackAlias(context global.DataContext, in Alias) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in Int) (value interface{}, dict bool, null bool, err error) {
		// localRepacker
		out, dict, null, err := p.RepackInt(context, in)
		if err != nil {
			return nil, false, false, err
		}
		return out, dict, null, nil
	}(context, (Int)(in))
}
func (p packageType) RepackAliasArray(context global.DataContext, in AliasArray) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in [3]string) (value interface{}, dict bool, null bool, err error) {
		// sliceRepacker
		out := make([]interface{}, len(in))
		empty := true
		for i, item := range in {
			v, _, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return pack.RepackString(in)
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
	}(context, ([3]string)(in))
}
func (p packageType) RepackAliasMap(context global.DataContext, in AliasMap) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in map[string]*Qual) (value interface{}, dict bool, null bool, err error) {
		// mapRepacker
		out := make(map[string]interface{}, len(in))
		for k, item := range in {
			v, _, _, err := func(context global.DataContext, in *Qual) (value interface{}, dict bool, null bool, err error) {
				// pointerRepacker
				if in == nil {
					return nil, false, true, nil
				}
				out, dict, null, err := func(context global.DataContext, in Qual) (value interface{}, dict bool, null bool, err error) {
					// localRepacker
					out, dict, null, err := p.RepackQual(context, in)
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, null, nil
				}(context, *in)
				if err != nil {
					return nil, false, false, err
				}
				return out, dict, false, nil
			}(context, item)
			if err != nil {
				return nil, false, false, err
			}
			out[k] = v
		}
		return out, true, len(in) == 0, nil
	}(context, (map[string]*Qual)(in))
}
func (p packageType) RepackAliasPointer(context global.DataContext, in AliasPointer) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in *Int) (value interface{}, dict bool, null bool, err error) {
		// pointerRepacker
		if in == nil {
			return nil, false, true, nil
		}
		out, dict, null, err := func(context global.DataContext, in Int) (value interface{}, dict bool, null bool, err error) {
			// localRepacker
			out, dict, null, err := p.RepackInt(context, in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(context, *in)
		if err != nil {
			return nil, false, false, err
		}
		return out, dict, false, nil
	}(context, (*Int)(in))
}
func (p packageType) RepackAliasSlice(context global.DataContext, in AliasSlice) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in []Int) (value interface{}, dict bool, null bool, err error) {
		// sliceRepacker
		out := make([]interface{}, len(in))
		empty := true
		for i, item := range in {
			v, _, null, err := func(context global.DataContext, in Int) (value interface{}, dict bool, null bool, err error) {
				// localRepacker
				out, dict, null, err := p.RepackInt(context, in)
				if err != nil {
					return nil, false, false, err
				}
				return out, dict, null, nil
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
	}(context, ([]Int)(in))
}
func (p packageType) RepackAliasSub(context global.DataContext, in AliasSub) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker
		p := context.Package().Get("frizz.io/tests/packer/sub")
		if p != nil {
			out, dict, null, err := p.Repack(context, in, "Sub")
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
		return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
	}(context, (sub.Sub)(in))
}
func (p packageType) RepackCsv(context global.DataContext, in Csv) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(context)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
}
func (p packageType) RepackCustom(context global.DataContext, in Custom) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(context)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
}
func (p packageType) RepackCustomSub(context global.DataContext, in CustomSub) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(context)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
}
func (p packageType) RepackEmbedNatives(context global.DataContext, in EmbedNatives) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Natives
		Int int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(context global.DataContext, in Natives) (value interface{}, dict bool, null bool, err error) {
			// localRepacker
			out, dict, null, err := p.RepackNatives(context, in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(context, in.Natives); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Natives"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Natives
		Int int
	})(in))
}
func (p packageType) RepackEmbedPointer(context global.DataContext, in EmbedPointer) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		*Natives
		Int int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(context global.DataContext, in *Natives) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(context global.DataContext, in Natives) (value interface{}, dict bool, null bool, err error) {
				// localRepacker
				out, dict, null, err := p.RepackNatives(context, in)
				if err != nil {
					return nil, false, false, err
				}
				return out, dict, null, nil
			}(context, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(context, in.Natives); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Natives"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		*Natives
		Int int
	})(in))
}
func (p packageType) RepackEmbedQual(context global.DataContext, in EmbedQual) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		sub.Sub
		Int int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(context global.DataContext, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
			// selectorRepacker
			p := context.Package().Get("frizz.io/tests/packer/sub")
			if p != nil {
				out, dict, null, err := p.Repack(context, in, "Sub")
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
			return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
		}(context, in.Sub); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Sub"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		sub.Sub
		Int int
	})(in))
}
func (p packageType) RepackEmbedQualPointer(context global.DataContext, in EmbedQualPointer) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		*sub.Sub
		Int int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(context global.DataContext, in *sub.Sub) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(context global.DataContext, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
				// selectorRepacker
				p := context.Package().Get("frizz.io/tests/packer/sub")
				if p != nil {
					out, dict, null, err := p.Repack(context, in, "Sub")
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
				return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
			}(context, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(context, in.Sub); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Sub"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		*sub.Sub
		Int int
	})(in))
}
func (p packageType) RepackFloat64(context global.DataContext, in Float64) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in float64) (value interface{}, dict bool, null bool, err error) {
		// nativeRepacker
		return pack.RepackNumber(in)
	}(context, (float64)(in))
}
func (p packageType) RepackHasTime(context global.DataContext, in HasTime) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Time time.Time
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(context global.DataContext, in time.Time) (value interface{}, dict bool, null bool, err error) {
			// selectorRepacker
			p := context.Package().Get("time")
			if p != nil {
				out, dict, null, err := p.Repack(context, in, "Time")
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
			return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "time", "Time")
		}(context, in.Time); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Time"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Time time.Time
	})(in))
}
func (p packageType) RepackImpSilent(context global.DataContext, in ImpSilent) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
		// nativeRepacker
		return pack.RepackNumber(in)
	}(context, (int)(in))
}
func (p packageType) RepackImpi(context global.DataContext, in Impi) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Int int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Int int
	})(in))
}
func (p packageType) RepackImps(context global.DataContext, in Imps) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		String string
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackString(in)
		}(context, in.String); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["String"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		String string
	})(in))
}
func (p packageType) RepackInt(context global.DataContext, in Int) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
		// nativeRepacker
		return pack.RepackNumber(in)
	}(context, (int)(in))
}
func (p packageType) RepackInterface(context global.DataContext, in Interface) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in interface {
		Foo() string
	}) (value interface{}, dict bool, null bool, err error) {
		// interfaceRepacker
		return pack.RepackInterface(context, false, in)
	}(context, (interface {
		Foo() string
	})(in))
}
func (p packageType) RepackInterfaceField(context global.DataContext, in InterfaceField) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Iface Interface
		Slice []Interface
		Array [3]Interface
		Map   map[string]Interface
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 5)
		empty := true
		if v, _, null, err := func(context global.DataContext, in Interface) (value interface{}, dict bool, null bool, err error) {
			// localRepacker
			out, dict, null, err := p.RepackInterface(context, in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(context, in.Iface); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Iface"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in []Interface) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.DataContext, in Interface) (value interface{}, dict bool, null bool, err error) {
					// localRepacker
					out, dict, null, err := p.RepackInterface(context, in)
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, null, nil
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
		}(context, in.Slice); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Slice"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in [3]Interface) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.DataContext, in Interface) (value interface{}, dict bool, null bool, err error) {
					// localRepacker
					out, dict, null, err := p.RepackInterface(context, in)
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, null, nil
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
		}(context, in.Array); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Array"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in map[string]Interface) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(context global.DataContext, in Interface) (value interface{}, dict bool, null bool, err error) {
					// localRepacker
					out, dict, null, err := p.RepackInterface(context, in)
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, null, nil
				}(context, item)
				if err != nil {
					return nil, false, false, err
				}
				out[k] = v
			}
			return out, true, len(in) == 0, nil
		}(context, in.Map); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Map"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Iface Interface
		Slice []Interface
		Array [3]Interface
		Map   map[string]Interface
	})(in))
}
func (p packageType) RepackInterfaceSlice(context global.DataContext, in InterfaceSlice) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in []Interface) (value interface{}, dict bool, null bool, err error) {
		// sliceRepacker
		out := make([]interface{}, len(in))
		empty := true
		for i, item := range in {
			v, _, null, err := func(context global.DataContext, in Interface) (value interface{}, dict bool, null bool, err error) {
				// localRepacker
				out, dict, null, err := p.RepackInterface(context, in)
				if err != nil {
					return nil, false, false, err
				}
				return out, dict, null, nil
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
	}(context, ([]Interface)(in))
}
func (p packageType) RepackInterfaceValidator(context global.DataContext, in InterfaceValidator) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
		// nativeRepacker
		return pack.RepackString(in)
	}(context, (string)(in))
}
func (p packageType) RepackMaps(context global.DataContext, in Maps) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Ints    map[string]int
		Strings map[string]string
		Slices  map[string][]string
		Arrays  map[string][2]int
		Maps    map[string]map[string]string
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 6)
		empty := true
		if v, _, null, err := func(context global.DataContext, in map[string]int) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return pack.RepackNumber(in)
				}(context, item)
				if err != nil {
					return nil, false, false, err
				}
				out[k] = v
			}
			return out, true, len(in) == 0, nil
		}(context, in.Ints); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Ints"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in map[string]string) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return pack.RepackString(in)
				}(context, item)
				if err != nil {
					return nil, false, false, err
				}
				out[k] = v
			}
			return out, true, len(in) == 0, nil
		}(context, in.Strings); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Strings"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in map[string][]string) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(context global.DataContext, in []string) (value interface{}, dict bool, null bool, err error) {
					// sliceRepacker
					out := make([]interface{}, len(in))
					empty := true
					for i, item := range in {
						v, _, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
							// nativeRepacker
							return pack.RepackString(in)
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
		}(context, in.Slices); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Slices"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in map[string][2]int) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(context global.DataContext, in [2]int) (value interface{}, dict bool, null bool, err error) {
					// sliceRepacker
					out := make([]interface{}, len(in))
					empty := true
					for i, item := range in {
						v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
							// nativeRepacker
							return pack.RepackNumber(in)
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
		}(context, in.Arrays); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Arrays"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in map[string]map[string]string) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(context global.DataContext, in map[string]string) (value interface{}, dict bool, null bool, err error) {
					// mapRepacker
					out := make(map[string]interface{}, len(in))
					for k, item := range in {
						v, _, _, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
							// nativeRepacker
							return pack.RepackString(in)
						}(context, item)
						if err != nil {
							return nil, false, false, err
						}
						out[k] = v
					}
					return out, true, len(in) == 0, nil
				}(context, item)
				if err != nil {
					return nil, false, false, err
				}
				out[k] = v
			}
			return out, true, len(in) == 0, nil
		}(context, in.Maps); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Maps"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Ints    map[string]int
		Strings map[string]string
		Slices  map[string][]string
		Arrays  map[string][2]int
		Maps    map[string]map[string]string
	})(in))
}
func (p packageType) RepackNatives(context global.DataContext, in Natives) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Bool      bool
		Byte      byte
		Float32   float32
		Float64   float64
		Int       int
		Int8      int8
		Int16     int16
		Int32     int32
		Int64     int64
		Uint      uint
		Uint8     uint8
		Uint16    uint16
		Uint32    uint32
		Uint64    uint64
		Rune      rune
		String    string
		PtrString *string
		PtrInt    *int
		PtrBool   *bool
		Number    json.Number
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 21)
		empty := true
		if v, _, null, err := func(context global.DataContext, in bool) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackBool(in)
		}(context, in.Bool); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Bool"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in byte) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Byte); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Byte"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in float32) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Float32); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Float32"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in float64) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Float64); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Float64"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in int8) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Int8); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int8"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in int16) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Int16); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int16"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in int32) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Int32); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int32"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in int64) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Int64); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int64"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in uint) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Uint); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Uint"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in uint8) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Uint8); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Uint8"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in uint16) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Uint16); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Uint16"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in uint32) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Uint32); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Uint32"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in uint64) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Uint64); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Uint64"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in rune) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackString(in)
		}(context, in.Rune); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Rune"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackString(in)
		}(context, in.String); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["String"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in *string) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return pack.RepackString(in)
			}(context, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(context, in.PtrString); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["PtrString"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in *int) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return pack.RepackNumber(in)
			}(context, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(context, in.PtrInt); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["PtrInt"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in *bool) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(context global.DataContext, in bool) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return pack.RepackBool(in)
			}(context, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(context, in.PtrBool); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["PtrBool"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in json.Number) (value interface{}, dict bool, null bool, err error) {
			// selectorRepacker (json.Number)
			if in == "" {
				return value, false, true, nil
			}
			return in, false, false, nil
		}(context, in.Number); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Number"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Bool      bool
		Byte      byte
		Float32   float32
		Float64   float64
		Int       int
		Int8      int8
		Int16     int16
		Int32     int32
		Int64     int64
		Uint      uint
		Uint8     uint8
		Uint16    uint16
		Uint32    uint32
		Uint64    uint64
		Rune      rune
		String    string
		PtrString *string
		PtrInt    *int
		PtrBool   *bool
		Number    json.Number
	})(in))
}
func (p packageType) RepackPointers(context global.DataContext, in Pointers) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		String      *string
		Int         *Int
		Sub         *sub.Sub
		Array       *[3]int
		Slice       *[]string
		Map         *map[string]int
		SliceString []*string
		SliceInt    []*Int
		SliceSub    []*sub.Sub
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 10)
		empty := true
		if v, _, null, err := func(context global.DataContext, in *string) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return pack.RepackString(in)
			}(context, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(context, in.String); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["String"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in *Int) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(context global.DataContext, in Int) (value interface{}, dict bool, null bool, err error) {
				// localRepacker
				out, dict, null, err := p.RepackInt(context, in)
				if err != nil {
					return nil, false, false, err
				}
				return out, dict, null, nil
			}(context, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(context, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in *sub.Sub) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(context global.DataContext, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
				// selectorRepacker
				p := context.Package().Get("frizz.io/tests/packer/sub")
				if p != nil {
					out, dict, null, err := p.Repack(context, in, "Sub")
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
				return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
			}(context, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(context, in.Sub); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Sub"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in *[3]int) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(context global.DataContext, in [3]int) (value interface{}, dict bool, null bool, err error) {
				// sliceRepacker
				out := make([]interface{}, len(in))
				empty := true
				for i, item := range in {
					v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
						// nativeRepacker
						return pack.RepackNumber(in)
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
			}(context, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(context, in.Array); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Array"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in *[]string) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(context global.DataContext, in []string) (value interface{}, dict bool, null bool, err error) {
				// sliceRepacker
				out := make([]interface{}, len(in))
				empty := true
				for i, item := range in {
					v, _, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
						// nativeRepacker
						return pack.RepackString(in)
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
			}(context, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(context, in.Slice); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Slice"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in *map[string]int) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(context global.DataContext, in map[string]int) (value interface{}, dict bool, null bool, err error) {
				// mapRepacker
				out := make(map[string]interface{}, len(in))
				for k, item := range in {
					v, _, _, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
						// nativeRepacker
						return pack.RepackNumber(in)
					}(context, item)
					if err != nil {
						return nil, false, false, err
					}
					out[k] = v
				}
				return out, true, len(in) == 0, nil
			}(context, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(context, in.Map); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Map"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in []*string) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.DataContext, in *string) (value interface{}, dict bool, null bool, err error) {
					// pointerRepacker
					if in == nil {
						return nil, false, true, nil
					}
					out, dict, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
						// nativeRepacker
						return pack.RepackString(in)
					}(context, *in)
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, false, nil
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
		}(context, in.SliceString); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["SliceString"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in []*Int) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.DataContext, in *Int) (value interface{}, dict bool, null bool, err error) {
					// pointerRepacker
					if in == nil {
						return nil, false, true, nil
					}
					out, dict, null, err := func(context global.DataContext, in Int) (value interface{}, dict bool, null bool, err error) {
						// localRepacker
						out, dict, null, err := p.RepackInt(context, in)
						if err != nil {
							return nil, false, false, err
						}
						return out, dict, null, nil
					}(context, *in)
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, false, nil
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
		}(context, in.SliceInt); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["SliceInt"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in []*sub.Sub) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.DataContext, in *sub.Sub) (value interface{}, dict bool, null bool, err error) {
					// pointerRepacker
					if in == nil {
						return nil, false, true, nil
					}
					out, dict, null, err := func(context global.DataContext, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
						// selectorRepacker
						p := context.Package().Get("frizz.io/tests/packer/sub")
						if p != nil {
							out, dict, null, err := p.Repack(context, in, "Sub")
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
						return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
					}(context, *in)
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, false, nil
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
		}(context, in.SliceSub); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["SliceSub"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		String      *string
		Int         *Int
		Sub         *sub.Sub
		Array       *[3]int
		Slice       *[]string
		Map         *map[string]int
		SliceString []*string
		SliceInt    []*Int
		SliceSub    []*sub.Sub
	})(in))
}
func (p packageType) RepackPrivate(context global.DataContext, in Private) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		i int
		s string
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.i); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["i"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackString(in)
		}(context, in.s); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["s"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		i int
		s string
	})(in))
}
func (p packageType) RepackQual(context global.DataContext, in Qual) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Sub sub.Sub
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(context global.DataContext, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
			// selectorRepacker
			p := context.Package().Get("frizz.io/tests/packer/sub")
			if p != nil {
				out, dict, null, err := p.Repack(context, in, "Sub")
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
			return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
		}(context, in.Sub); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Sub"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Sub sub.Sub
	})(in))
}
func (p packageType) RepackSilent(context global.DataContext, in Silent) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Silent silent.Silent
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(context global.DataContext, in silent.Silent) (value interface{}, dict bool, null bool, err error) {
			// selectorRepacker
			p := context.Package().Get("frizz.io/tests/packer/silent")
			if p != nil {
				out, dict, null, err := p.Repack(context, in, "Silent")
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
			return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/tests/packer/silent", "Silent")
		}(context, in.Silent); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Silent"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Silent silent.Silent
	})(in))
}
func (p packageType) RepackSlices(context global.DataContext, in Slices) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Ints      []int
		Strings   []string
		ArrayLit  [5]string
		ArrayExpr [2 * N]int
		Structs   []struct {
			Int int
		}
		Arrays [][]string
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 7)
		empty := true
		if v, _, null, err := func(context global.DataContext, in []int) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return pack.RepackNumber(in)
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
		}(context, in.Ints); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Ints"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in []string) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return pack.RepackString(in)
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
		}(context, in.Strings); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Strings"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in [5]string) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return pack.RepackString(in)
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
		}(context, in.ArrayLit); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["ArrayLit"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in [2 * N]int) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return pack.RepackNumber(in)
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
		}(context, in.ArrayExpr); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["ArrayExpr"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in []struct {
			Int int
		}) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.DataContext, in struct {
					Int int
				}) (value interface{}, dict bool, null bool, err error) {
					// structRepacker
					out := make(map[string]interface{}, 2)
					empty := true
					if v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
						// nativeRepacker
						return pack.RepackNumber(in)
					}(context, in.Int); err != nil {
						return nil, false, false, err
					} else {
						if !null {
							empty = false
							out["Int"] = v
						}
					}
					return out, false, empty, nil
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
		}(context, in.Structs); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Structs"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in [][]string) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.DataContext, in []string) (value interface{}, dict bool, null bool, err error) {
					// sliceRepacker
					out := make([]interface{}, len(in))
					empty := true
					for i, item := range in {
						v, _, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
							// nativeRepacker
							return pack.RepackString(in)
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
				if !null {
					empty = false
				}
				out[i] = v
			}
			return out, false, empty, nil
		}(context, in.Arrays); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Arrays"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Ints      []int
		Strings   []string
		ArrayLit  [5]string
		ArrayExpr [2 * N]int
		Structs   []struct {
			Int int
		}
		Arrays [][]string
	})(in))
}
func (p packageType) RepackString(context global.DataContext, in String) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
		// nativeRepacker
		return pack.RepackString(in)
	}(context, (string)(in))
}
func (p packageType) RepackStructs(context global.DataContext, in Structs) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Simple struct {
			Int  int
			Bool bool
		}
		Complex struct {
			String string
			Inner  struct {
				Float32 float32
			}
		}
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(context global.DataContext, in struct {
			Int  int
			Bool bool
		}) (value interface{}, dict bool, null bool, err error) {
			// structRepacker
			out := make(map[string]interface{}, 3)
			empty := true
			if v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return pack.RepackNumber(in)
			}(context, in.Int); err != nil {
				return nil, false, false, err
			} else {
				if !null {
					empty = false
					out["Int"] = v
				}
			}
			if v, _, null, err := func(context global.DataContext, in bool) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return pack.RepackBool(in)
			}(context, in.Bool); err != nil {
				return nil, false, false, err
			} else {
				if !null {
					empty = false
					out["Bool"] = v
				}
			}
			return out, false, empty, nil
		}(context, in.Simple); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Simple"] = v
			}
		}
		if v, _, null, err := func(context global.DataContext, in struct {
			String string
			Inner  struct {
				Float32 float32
			}
		}) (value interface{}, dict bool, null bool, err error) {
			// structRepacker
			out := make(map[string]interface{}, 3)
			empty := true
			if v, _, null, err := func(context global.DataContext, in string) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return pack.RepackString(in)
			}(context, in.String); err != nil {
				return nil, false, false, err
			} else {
				if !null {
					empty = false
					out["String"] = v
				}
			}
			if v, _, null, err := func(context global.DataContext, in struct {
				Float32 float32
			}) (value interface{}, dict bool, null bool, err error) {
				// structRepacker
				out := make(map[string]interface{}, 2)
				empty := true
				if v, _, null, err := func(context global.DataContext, in float32) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return pack.RepackNumber(in)
				}(context, in.Float32); err != nil {
					return nil, false, false, err
				} else {
					if !null {
						empty = false
						out["Float32"] = v
					}
				}
				return out, false, empty, nil
			}(context, in.Inner); err != nil {
				return nil, false, false, err
			} else {
				if !null {
					empty = false
					out["Inner"] = v
				}
			}
			return out, false, empty, nil
		}(context, in.Complex); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Complex"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Simple struct {
			Int  int
			Bool bool
		}
		Complex struct {
			String string
			Inner  struct {
				Float32 float32
			}
		}
	})(in))
}
func (p packageType) RepackSubInterface(context global.DataContext, in SubInterface) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		SubInterface sub.SubInterface
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(context global.DataContext, in sub.SubInterface) (value interface{}, dict bool, null bool, err error) {
			// selectorRepacker
			p := context.Package().Get("frizz.io/tests/packer/sub")
			if p != nil {
				out, dict, null, err := p.Repack(context, in, "SubInterface")
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
			return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "SubInterface")
		}(context, in.SubInterface); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["SubInterface"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		SubInterface sub.SubInterface
	})(in))
}
func (p packageType) RepackSubInterfaceSlice(context global.DataContext, in SubInterfaceSlice) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in []sub.SubInterface) (value interface{}, dict bool, null bool, err error) {
		// sliceRepacker
		out := make([]interface{}, len(in))
		empty := true
		for i, item := range in {
			v, _, null, err := func(context global.DataContext, in sub.SubInterface) (value interface{}, dict bool, null bool, err error) {
				// selectorRepacker
				p := context.Package().Get("frizz.io/tests/packer/sub")
				if p != nil {
					out, dict, null, err := p.Repack(context, in, "SubInterface")
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
				return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "SubInterface")
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
	}(context, ([]sub.SubInterface)(in))
}
func (p packageType) RepackSubMap(context global.DataContext, in SubMap) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Map map[string]sub.Sub
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(context global.DataContext, in map[string]sub.Sub) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(context global.DataContext, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
					// selectorRepacker
					p := context.Package().Get("frizz.io/tests/packer/sub")
					if p != nil {
						out, dict, null, err := p.Repack(context, in, "Sub")
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
					return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
				}(context, item)
				if err != nil {
					return nil, false, false, err
				}
				out[k] = v
			}
			return out, true, len(in) == 0, nil
		}(context, in.Map); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Map"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Map map[string]sub.Sub
	})(in))
}
func (p packageType) RepackSubSlice(context global.DataContext, in SubSlice) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Slice []sub.Sub
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(context global.DataContext, in []sub.Sub) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(context global.DataContext, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
					// selectorRepacker
					p := context.Package().Get("frizz.io/tests/packer/sub")
					if p != nil {
						out, dict, null, err := p.Repack(context, in, "Sub")
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
					return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/tests/packer/sub", "Sub")
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
		}(context, in.Slice); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Slice"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Slice []sub.Sub
	})(in))
}
func (p packageType) RepackType(context global.DataContext, in Type) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(context)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
}
func (p packageType) RepackUint(context global.DataContext, in Uint) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in uint) (value interface{}, dict bool, null bool, err error) {
		// nativeRepacker
		return pack.RepackNumber(in)
	}(context, (uint)(in))
}
func (p packageType) GetData(filename string) string {
	return ""
}
func (p packageType) GetType(name string) string {
	return ""
}
func (p packageType) GetImportedPackages(packages map[string]global.Package) {
	packages["frizz.io/tests/packer"] = Package
	silent.Package.GetImportedPackages(packages)
}
