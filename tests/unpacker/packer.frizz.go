package unpacker

import (
	frizz "frizz.io/frizz"
	sub "frizz.io/tests/unpacker/sub"
	errors "github.com/pkg/errors"
)

const Packer packer = 0

type packer int

func (p packer) Path() string {
	return "frizz.io/tests/unpacker"
}
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (interface{}, error) {
	switch name {
	case "CustomSub":
		return p.UnpackCustomSub(root, stack, in)
	case "Ages":
		return p.UnpackAges(root, stack, in)
	case "Csv":
		return p.UnpackCsv(root, stack, in)
	case "Type":
		return p.UnpackType(root, stack, in)
	case "Custom":
		return p.UnpackCustom(root, stack, in)
	case "EmbedNatives":
		return p.UnpackEmbedNatives(root, stack, in)
	case "EmbedPointer":
		return p.UnpackEmbedPointer(root, stack, in)
	case "EmbedQualPointer":
		return p.UnpackEmbedQualPointer(root, stack, in)
	case "EmbedQual":
		return p.UnpackEmbedQual(root, stack, in)
	case "InterfaceField":
		return p.UnpackInterfaceField(root, stack, in)
	case "Impi":
		return p.UnpackImpi(root, stack, in)
	case "Imps":
		return p.UnpackImps(root, stack, in)
	case "Interface":
		return p.UnpackInterface(root, stack, in)
	case "Private":
		return p.UnpackPrivate(root, stack, in)
	case "AliasSub":
		return p.UnpackAliasSub(root, stack, in)
	case "AliasSlice":
		return p.UnpackAliasSlice(root, stack, in)
	case "AliasArray":
		return p.UnpackAliasArray(root, stack, in)
	case "AliasMap":
		return p.UnpackAliasMap(root, stack, in)
	case "AliasPointer":
		return p.UnpackAliasPointer(root, stack, in)
	case "Alias":
		return p.UnpackAlias(root, stack, in)
	case "Int":
		return p.UnpackInt(root, stack, in)
	case "String":
		return p.UnpackString(root, stack, in)
	case "Qual":
		return p.UnpackQual(root, stack, in)
	case "Pointers":
		return p.UnpackPointers(root, stack, in)
	case "Maps":
		return p.UnpackMaps(root, stack, in)
	case "Slices":
		return p.UnpackSlices(root, stack, in)
	case "Structs":
		return p.UnpackStructs(root, stack, in)
	case "Natives":
		return p.UnpackNatives(root, stack, in)
	}
	return nil, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) UnpackCustomSub(root *frizz.Root, stack frizz.Stack, in interface{}) (value CustomSub, err error) {
	out := new(CustomSub)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func (p packer) UnpackAges(root *frizz.Root, stack frizz.Stack, in interface{}) (value Ages, err error) {
	out := new(Ages)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func (p packer) UnpackCsv(root *frizz.Root, stack frizz.Stack, in interface{}) (value Csv, err error) {
	out := new(Csv)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func (p packer) UnpackType(root *frizz.Root, stack frizz.Stack, in interface{}) (value Type, err error) {
	out := new(Type)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func (p packer) UnpackCustom(root *frizz.Root, stack frizz.Stack, in interface{}) (value Custom, err error) {
	out := new(Custom)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func (p packer) UnpackEmbedNatives(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedNatives, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out EmbedNatives
	if v, ok := m["Natives"]; ok {
		stack := stack.Append(frizz.FieldItem("Natives"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Natives, err error) {
			// localUnpacker
			out, err := p.UnpackNatives(root, stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Natives = u
	}
	if v, ok := m["Int"]; ok {
		stack := stack.Append(frizz.FieldItem("Int"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	return out, nil
}
func (p packer) UnpackEmbedPointer(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedPointer, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out EmbedPointer
	if v, ok := m["Natives"]; ok {
		stack := stack.Append(frizz.FieldItem("Natives"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *Natives, err error) {
			// pointerUnpacker
			out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Natives, err error) {
				// localUnpacker
				out, err := p.UnpackNatives(root, stack, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(root, stack, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Natives = u
	}
	if v, ok := m["Int"]; ok {
		stack := stack.Append(frizz.FieldItem("Int"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	return out, nil
}
func (p packer) UnpackEmbedQualPointer(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedQualPointer, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out EmbedQualPointer
	if v, ok := m["Sub"]; ok {
		stack := stack.Append(frizz.FieldItem("Sub"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *sub.Sub, err error) {
			// pointerUnpacker
			out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value sub.Sub, err error) {
				// selectorUnpacker
				out, err := sub.Packer.UnpackSub(root, stack, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(root, stack, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Sub = u
	}
	if v, ok := m["Int"]; ok {
		stack := stack.Append(frizz.FieldItem("Int"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	return out, nil
}
func (p packer) UnpackEmbedQual(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedQual, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out EmbedQual
	if v, ok := m["Sub"]; ok {
		stack := stack.Append(frizz.FieldItem("Sub"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value sub.Sub, err error) {
			// selectorUnpacker
			out, err := sub.Packer.UnpackSub(root, stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Sub = u
	}
	if v, ok := m["Int"]; ok {
		stack := stack.Append(frizz.FieldItem("Int"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	return out, nil
}
func (p packer) UnpackInterfaceField(root *frizz.Root, stack frizz.Stack, in interface{}) (value InterfaceField, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out InterfaceField
	if v, ok := m["Iface"]; ok {
		stack := stack.Append(frizz.FieldItem("Iface"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Interface, err error) {
			// localUnpacker
			out, err := p.UnpackInterface(root, stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Iface = u
	}
	if v, ok := m["Slice"]; ok {
		stack := stack.Append(frizz.FieldItem("Slice"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []Interface, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]Interface, len(a))
			for i, v := range a {
				stack := stack.Append(frizz.ArrayItem(i))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Interface, err error) {
					// localUnpacker
					out, err := p.UnpackInterface(root, stack, in)
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
		out.Slice = u
	}
	if v, ok := m["Array"]; ok {
		stack := stack.Append(frizz.FieldItem("Array"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [3]Interface, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out [3]Interface
			if len(a) > 3 {
				return value, errors.Errorf("data length %d does not fit in array of length %d", len(a), 3)
			}
			for i, v := range a {
				stack := stack.Append(frizz.ArrayItem(i))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Interface, err error) {
					// localUnpacker
					out, err := p.UnpackInterface(root, stack, in)
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
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Array = u
	}
	if v, ok := m["Map"]; ok {
		stack := stack.Append(frizz.FieldItem("Map"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]Interface, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]Interface, len(m))
			for k, v := range m {
				stack := stack.Append(frizz.MapItem(k))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Interface, err error) {
					// localUnpacker
					out, err := p.UnpackInterface(root, stack, in)
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
		out.Map = u
	}
	return out, nil
}
func (p packer) UnpackImpi(root *frizz.Root, stack frizz.Stack, in interface{}) (value Impi, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Impi
	if v, ok := m["Int"]; ok {
		stack := stack.Append(frizz.FieldItem("Int"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	return out, nil
}
func (p packer) UnpackImps(root *frizz.Root, stack frizz.Stack, in interface{}) (value Imps, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Imps
	if v, ok := m["String"]; ok {
		stack := stack.Append(frizz.FieldItem("String"))
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
		out.String = u
	}
	return out, nil
}
func (p packer) UnpackInterface(root *frizz.Root, stack frizz.Stack, in interface{}) (value Interface, err error) {
	// interfaceUnpacker
	out, err := root.UnpackInterface(stack, in)
	if err != nil {
		return value, err
	}
	iface, ok := out.(Interface)
	if !ok {
		return value, errors.Errorf("unpacking into interface, type %T does not implement interface", out)
	}
	return iface, nil
}
func (p packer) UnpackPrivate(root *frizz.Root, stack frizz.Stack, in interface{}) (value Private, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Private
	if v, ok := m["i"]; ok {
		stack := stack.Append(frizz.FieldItem("i"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.i = u
	}
	if v, ok := m["s"]; ok {
		stack := stack.Append(frizz.FieldItem("s"))
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
		out.s = u
	}
	return out, nil
}
func (p packer) UnpackAliasSub(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasSub, err error) {
	// selectorUnpacker
	out, err := sub.Packer.UnpackSub(root, stack, in)
	if err != nil {
		return value, err
	}
	return AliasSub(out), nil
}
func (p packer) UnpackAliasSlice(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasSlice, err error) {
	// sliceUnpacker
	a, ok := in.([]interface{})
	if !ok {
		return value, errors.New("unpacking into slice, value should be an array")
	}
	var out = make(AliasSlice, len(a))
	for i, v := range a {
		stack := stack.Append(frizz.ArrayItem(i))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, err error) {
			// localUnpacker
			out, err := p.UnpackInt(root, stack, in)
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
}
func (p packer) UnpackAliasArray(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasArray, err error) {
	// sliceUnpacker
	a, ok := in.([]interface{})
	if !ok {
		return value, errors.New("unpacking into slice, value should be an array")
	}
	var out AliasArray
	if len(a) > 3 {
		return value, errors.Errorf("data length %d does not fit in array of length %d", len(a), 3)
	}
	for i, v := range a {
		stack := stack.Append(frizz.ArrayItem(i))
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
		out[i] = u
	}
	return out, nil
}
func (p packer) UnpackAliasMap(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasMap, err error) {
	// mapUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into map, value should be a map")
	}
	var out = make(AliasMap, len(m))
	for k, v := range m {
		stack := stack.Append(frizz.MapItem(k))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *Qual, err error) {
			// pointerUnpacker
			out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Qual, err error) {
				// localUnpacker
				out, err := p.UnpackQual(root, stack, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(root, stack, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out[k] = u
	}
	return out, nil
}
func (p packer) UnpackAliasPointer(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasPointer, err error) {
	// pointerUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, err error) {
		// localUnpacker
		out, err := p.UnpackInt(root, stack, in)
		if err != nil {
			return value, err
		}
		return out, nil
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return AliasPointer(&out), nil
}
func (p packer) UnpackAlias(root *frizz.Root, stack frizz.Stack, in interface{}) (value Alias, err error) {
	// localUnpacker
	out, err := p.UnpackInt(root, stack, in)
	if err != nil {
		return value, err
	}
	return Alias(out), nil
}
func (p packer) UnpackInt(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, err error) {
	// nativeUnpacker
	out, err := frizz.UnpackInt(stack, in)
	if err != nil {
		return value, err
	}
	return Int(out), nil
}
func (p packer) UnpackString(root *frizz.Root, stack frizz.Stack, in interface{}) (value String, err error) {
	// nativeUnpacker
	out, err := frizz.UnpackString(stack, in)
	if err != nil {
		return value, err
	}
	return String(out), nil
}
func (p packer) UnpackQual(root *frizz.Root, stack frizz.Stack, in interface{}) (value Qual, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Qual
	if v, ok := m["Sub"]; ok {
		stack := stack.Append(frizz.FieldItem("Sub"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value sub.Sub, err error) {
			// selectorUnpacker
			out, err := sub.Packer.UnpackSub(root, stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Sub = u
	}
	return out, nil
}
func (p packer) UnpackPointers(root *frizz.Root, stack frizz.Stack, in interface{}) (value Pointers, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Pointers
	if v, ok := m["String"]; ok {
		stack := stack.Append(frizz.FieldItem("String"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *string, err error) {
			// pointerUnpacker
			out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value string, err error) {
				// nativeUnpacker
				out, err := frizz.UnpackString(stack, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(root, stack, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.String = u
	}
	if v, ok := m["Int"]; ok {
		stack := stack.Append(frizz.FieldItem("Int"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *Int, err error) {
			// pointerUnpacker
			out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, err error) {
				// localUnpacker
				out, err := p.UnpackInt(root, stack, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(root, stack, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	if v, ok := m["Sub"]; ok {
		stack := stack.Append(frizz.FieldItem("Sub"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *sub.Sub, err error) {
			// pointerUnpacker
			out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value sub.Sub, err error) {
				// selectorUnpacker
				out, err := sub.Packer.UnpackSub(root, stack, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(root, stack, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Sub = u
	}
	if v, ok := m["Array"]; ok {
		stack := stack.Append(frizz.FieldItem("Array"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *[3]int, err error) {
			// pointerUnpacker
			out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [3]int, err error) {
				// sliceUnpacker
				a, ok := in.([]interface{})
				if !ok {
					return value, errors.New("unpacking into slice, value should be an array")
				}
				var out [3]int
				if len(a) > 3 {
					return value, errors.Errorf("data length %d does not fit in array of length %d", len(a), 3)
				}
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
					u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
						// nativeUnpacker
						out, err := frizz.UnpackInt(stack, in)
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
				return out, nil
			}(root, stack, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Array = u
	}
	if v, ok := m["Slice"]; ok {
		stack := stack.Append(frizz.FieldItem("Slice"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *[]string, err error) {
			// pointerUnpacker
			out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []string, err error) {
				// sliceUnpacker
				a, ok := in.([]interface{})
				if !ok {
					return value, errors.New("unpacking into slice, value should be an array")
				}
				var out = make([]string, len(a))
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
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
					out[i] = u
				}
				return out[:], nil
			}(root, stack, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Slice = u
	}
	if v, ok := m["Map"]; ok {
		stack := stack.Append(frizz.FieldItem("Map"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *map[string]int, err error) {
			// pointerUnpacker
			out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]int, err error) {
				// mapUnpacker
				m, ok := in.(map[string]interface{})
				if !ok {
					return value, errors.New("unpacking into map, value should be a map")
				}
				var out = make(map[string]int, len(m))
				for k, v := range m {
					stack := stack.Append(frizz.MapItem(k))
					u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
						// nativeUnpacker
						out, err := frizz.UnpackInt(stack, in)
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
			}(root, stack, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Map = u
	}
	if v, ok := m["SliceString"]; ok {
		stack := stack.Append(frizz.FieldItem("SliceString"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []*string, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]*string, len(a))
			for i, v := range a {
				stack := stack.Append(frizz.ArrayItem(i))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *string, err error) {
					// pointerUnpacker
					out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value string, err error) {
						// nativeUnpacker
						out, err := frizz.UnpackString(stack, in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(root, stack, in)
					if err != nil {
						return value, err
					}
					return &out, nil
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
		out.SliceString = u
	}
	if v, ok := m["SliceInt"]; ok {
		stack := stack.Append(frizz.FieldItem("SliceInt"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []*Int, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]*Int, len(a))
			for i, v := range a {
				stack := stack.Append(frizz.ArrayItem(i))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *Int, err error) {
					// pointerUnpacker
					out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, err error) {
						// localUnpacker
						out, err := p.UnpackInt(root, stack, in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(root, stack, in)
					if err != nil {
						return value, err
					}
					return &out, nil
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
		out.SliceInt = u
	}
	if v, ok := m["SliceSub"]; ok {
		stack := stack.Append(frizz.FieldItem("SliceSub"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []*sub.Sub, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]*sub.Sub, len(a))
			for i, v := range a {
				stack := stack.Append(frizz.ArrayItem(i))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *sub.Sub, err error) {
					// pointerUnpacker
					out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value sub.Sub, err error) {
						// selectorUnpacker
						out, err := sub.Packer.UnpackSub(root, stack, in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(root, stack, in)
					if err != nil {
						return value, err
					}
					return &out, nil
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
		out.SliceSub = u
	}
	return out, nil
}
func (p packer) UnpackMaps(root *frizz.Root, stack frizz.Stack, in interface{}) (value Maps, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Maps
	if v, ok := m["Ints"]; ok {
		stack := stack.Append(frizz.FieldItem("Ints"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]int, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]int, len(m))
			for k, v := range m {
				stack := stack.Append(frizz.MapItem(k))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackInt(stack, in)
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
		out.Ints = u
	}
	if v, ok := m["Strings"]; ok {
		stack := stack.Append(frizz.FieldItem("Strings"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]string, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]string, len(m))
			for k, v := range m {
				stack := stack.Append(frizz.MapItem(k))
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
				out[k] = u
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Strings = u
	}
	if v, ok := m["Slices"]; ok {
		stack := stack.Append(frizz.FieldItem("Slices"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string][]string, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string][]string, len(m))
			for k, v := range m {
				stack := stack.Append(frizz.MapItem(k))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []string, err error) {
					// sliceUnpacker
					a, ok := in.([]interface{})
					if !ok {
						return value, errors.New("unpacking into slice, value should be an array")
					}
					var out = make([]string, len(a))
					for i, v := range a {
						stack := stack.Append(frizz.ArrayItem(i))
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
						out[i] = u
					}
					return out[:], nil
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
		out.Slices = u
	}
	if v, ok := m["Arrays"]; ok {
		stack := stack.Append(frizz.FieldItem("Arrays"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string][2]int, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string][2]int, len(m))
			for k, v := range m {
				stack := stack.Append(frizz.MapItem(k))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [2]int, err error) {
					// sliceUnpacker
					a, ok := in.([]interface{})
					if !ok {
						return value, errors.New("unpacking into slice, value should be an array")
					}
					var out [2]int
					if len(a) > 2 {
						return value, errors.Errorf("data length %d does not fit in array of length %d", len(a), 2)
					}
					for i, v := range a {
						stack := stack.Append(frizz.ArrayItem(i))
						u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
							// nativeUnpacker
							out, err := frizz.UnpackInt(stack, in)
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
		out.Arrays = u
	}
	if v, ok := m["Maps"]; ok {
		stack := stack.Append(frizz.FieldItem("Maps"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]map[string]string, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]map[string]string, len(m))
			for k, v := range m {
				stack := stack.Append(frizz.MapItem(k))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]string, err error) {
					// mapUnpacker
					m, ok := in.(map[string]interface{})
					if !ok {
						return value, errors.New("unpacking into map, value should be a map")
					}
					var out = make(map[string]string, len(m))
					for k, v := range m {
						stack := stack.Append(frizz.MapItem(k))
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
						out[k] = u
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
		out.Maps = u
	}
	return out, nil
}
func (p packer) UnpackSlices(root *frizz.Root, stack frizz.Stack, in interface{}) (value Slices, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Slices
	if v, ok := m["Ints"]; ok {
		stack := stack.Append(frizz.FieldItem("Ints"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []int, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]int, len(a))
			for i, v := range a {
				stack := stack.Append(frizz.ArrayItem(i))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackInt(stack, in)
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
		out.Ints = u
	}
	if v, ok := m["Strings"]; ok {
		stack := stack.Append(frizz.FieldItem("Strings"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []string, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]string, len(a))
			for i, v := range a {
				stack := stack.Append(frizz.ArrayItem(i))
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
				out[i] = u
			}
			return out[:], nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Strings = u
	}
	if v, ok := m["ArrayLit"]; ok {
		stack := stack.Append(frizz.FieldItem("ArrayLit"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [5]string, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out [5]string
			if len(a) > 5 {
				return value, errors.Errorf("data length %d does not fit in array of length %d", len(a), 5)
			}
			for i, v := range a {
				stack := stack.Append(frizz.ArrayItem(i))
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
				out[i] = u
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.ArrayLit = u
	}
	if v, ok := m["ArrayExpr"]; ok {
		stack := stack.Append(frizz.FieldItem("ArrayExpr"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [2 * N]int, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out [2 * N]int
			if len(a) > 2*N {
				return value, errors.Errorf("data length %d does not fit in array of length %d", len(a), 2*N)
			}
			for i, v := range a {
				stack := stack.Append(frizz.ArrayItem(i))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackInt(stack, in)
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
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.ArrayExpr = u
	}
	if v, ok := m["Structs"]; ok {
		stack := stack.Append(frizz.FieldItem("Structs"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []struct {
			Int int
		}, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]struct {
				Int int
			}, len(a))
			for i, v := range a {
				stack := stack.Append(frizz.ArrayItem(i))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
					Int int
				}, err error) {
					// structUnpacker
					m, ok := in.(map[string]interface{})
					if !ok {
						return value, errors.New("unpacking into struct, value should be a map")
					}
					var out struct {
						Int int
					}
					if v, ok := m["Int"]; ok {
						stack := stack.Append(frizz.FieldItem("Int"))
						u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
							// nativeUnpacker
							out, err := frizz.UnpackInt(stack, in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(root, stack, v)
						if err != nil {
							return value, err
						}
						out.Int = u
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
		out.Structs = u
	}
	if v, ok := m["Arrays"]; ok {
		stack := stack.Append(frizz.FieldItem("Arrays"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [][]string, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([][]string, len(a))
			for i, v := range a {
				stack := stack.Append(frizz.ArrayItem(i))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []string, err error) {
					// sliceUnpacker
					a, ok := in.([]interface{})
					if !ok {
						return value, errors.New("unpacking into slice, value should be an array")
					}
					var out = make([]string, len(a))
					for i, v := range a {
						stack := stack.Append(frizz.ArrayItem(i))
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
						out[i] = u
					}
					return out[:], nil
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
		out.Arrays = u
	}
	return out, nil
}
func (p packer) UnpackStructs(root *frizz.Root, stack frizz.Stack, in interface{}) (value Structs, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Structs
	if v, ok := m["Simple"]; ok {
		stack := stack.Append(frizz.FieldItem("Simple"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
			Int  int
			Bool bool
		}, err error) {
			// structUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into struct, value should be a map")
			}
			var out struct {
				Int  int
				Bool bool
			}
			if v, ok := m["Int"]; ok {
				stack := stack.Append(frizz.FieldItem("Int"))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackInt(stack, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(root, stack, v)
				if err != nil {
					return value, err
				}
				out.Int = u
			}
			if v, ok := m["Bool"]; ok {
				stack := stack.Append(frizz.FieldItem("Bool"))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value bool, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackBool(stack, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(root, stack, v)
				if err != nil {
					return value, err
				}
				out.Bool = u
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Simple = u
	}
	if v, ok := m["Complex"]; ok {
		stack := stack.Append(frizz.FieldItem("Complex"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
			String string
			Inner  struct {
				Float32 float32
			}
		}, err error) {
			// structUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into struct, value should be a map")
			}
			var out struct {
				String string
				Inner  struct {
					Float32 float32
				}
			}
			if v, ok := m["String"]; ok {
				stack := stack.Append(frizz.FieldItem("String"))
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
				out.String = u
			}
			if v, ok := m["Inner"]; ok {
				stack := stack.Append(frizz.FieldItem("Inner"))
				u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
					Float32 float32
				}, err error) {
					// structUnpacker
					m, ok := in.(map[string]interface{})
					if !ok {
						return value, errors.New("unpacking into struct, value should be a map")
					}
					var out struct {
						Float32 float32
					}
					if v, ok := m["Float32"]; ok {
						stack := stack.Append(frizz.FieldItem("Float32"))
						u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value float32, err error) {
							// nativeUnpacker
							out, err := frizz.UnpackFloat32(stack, in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(root, stack, v)
						if err != nil {
							return value, err
						}
						out.Float32 = u
					}
					return out, nil
				}(root, stack, v)
				if err != nil {
					return value, err
				}
				out.Inner = u
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Complex = u
	}
	return out, nil
}
func (p packer) UnpackNatives(root *frizz.Root, stack frizz.Stack, in interface{}) (value Natives, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Natives
	if v, ok := m["Bool"]; ok {
		stack := stack.Append(frizz.FieldItem("Bool"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value bool, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackBool(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Bool = u
	}
	if v, ok := m["Byte"]; ok {
		stack := stack.Append(frizz.FieldItem("Byte"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value byte, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackByte(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Byte = u
	}
	if v, ok := m["Float32"]; ok {
		stack := stack.Append(frizz.FieldItem("Float32"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value float32, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackFloat32(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Float32 = u
	}
	if v, ok := m["Float64"]; ok {
		stack := stack.Append(frizz.FieldItem("Float64"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value float64, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackFloat64(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Float64 = u
	}
	if v, ok := m["Int"]; ok {
		stack := stack.Append(frizz.FieldItem("Int"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	if v, ok := m["Int8"]; ok {
		stack := stack.Append(frizz.FieldItem("Int8"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int8, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt8(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Int8 = u
	}
	if v, ok := m["Int16"]; ok {
		stack := stack.Append(frizz.FieldItem("Int16"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int16, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt16(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Int16 = u
	}
	if v, ok := m["Int32"]; ok {
		stack := stack.Append(frizz.FieldItem("Int32"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int32, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt32(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Int32 = u
	}
	if v, ok := m["Int64"]; ok {
		stack := stack.Append(frizz.FieldItem("Int64"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int64, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt64(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Int64 = u
	}
	if v, ok := m["Uint"]; ok {
		stack := stack.Append(frizz.FieldItem("Uint"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value uint, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackUint(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Uint = u
	}
	if v, ok := m["Uint8"]; ok {
		stack := stack.Append(frizz.FieldItem("Uint8"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value uint8, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackUint8(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Uint8 = u
	}
	if v, ok := m["Uint16"]; ok {
		stack := stack.Append(frizz.FieldItem("Uint16"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value uint16, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackUint16(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Uint16 = u
	}
	if v, ok := m["Uint32"]; ok {
		stack := stack.Append(frizz.FieldItem("Uint32"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value uint32, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackUint32(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Uint32 = u
	}
	if v, ok := m["Uint64"]; ok {
		stack := stack.Append(frizz.FieldItem("Uint64"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value uint64, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackUint64(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Uint64 = u
	}
	if v, ok := m["Rune"]; ok {
		stack := stack.Append(frizz.FieldItem("Rune"))
		u, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value rune, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackRune(stack, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(root, stack, v)
		if err != nil {
			return value, err
		}
		out.Rune = u
	}
	if v, ok := m["String"]; ok {
		stack := stack.Append(frizz.FieldItem("String"))
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
		out.String = u
	}
	return out, nil
}
