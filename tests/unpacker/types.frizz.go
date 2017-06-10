package unpacker

import (
	frizz "frizz.io/frizz"
	sub "frizz.io/tests/unpacker/sub"
	errors "github.com/pkg/errors"
)

var Unpackers = struct {
	Ages             func(*frizz.Root, frizz.Stack, interface{}) (Ages, error)
	Csv              func(*frizz.Root, frizz.Stack, interface{}) (Csv, error)
	Type             func(*frizz.Root, frizz.Stack, interface{}) (Type, error)
	Custom           func(*frizz.Root, frizz.Stack, interface{}) (Custom, error)
	EmbedNatives     func(*frizz.Root, frizz.Stack, interface{}) (EmbedNatives, error)
	EmbedPointer     func(*frizz.Root, frizz.Stack, interface{}) (EmbedPointer, error)
	EmbedQualPointer func(*frizz.Root, frizz.Stack, interface{}) (EmbedQualPointer, error)
	EmbedQual        func(*frizz.Root, frizz.Stack, interface{}) (EmbedQual, error)
	InterfaceField   func(*frizz.Root, frizz.Stack, interface{}) (InterfaceField, error)
	Impi             func(*frizz.Root, frizz.Stack, interface{}) (Impi, error)
	Imps             func(*frizz.Root, frizz.Stack, interface{}) (Imps, error)
	Interface        func(*frizz.Root, frizz.Stack, interface{}) (Interface, error)
	Private          func(*frizz.Root, frizz.Stack, interface{}) (Private, error)
	AliasSub         func(*frizz.Root, frizz.Stack, interface{}) (AliasSub, error)
	AliasSlice       func(*frizz.Root, frizz.Stack, interface{}) (AliasSlice, error)
	AliasArray       func(*frizz.Root, frizz.Stack, interface{}) (AliasArray, error)
	AliasMap         func(*frizz.Root, frizz.Stack, interface{}) (AliasMap, error)
	AliasPointer     func(*frizz.Root, frizz.Stack, interface{}) (AliasPointer, error)
	Alias            func(*frizz.Root, frizz.Stack, interface{}) (Alias, error)
	Int              func(*frizz.Root, frizz.Stack, interface{}) (Int, error)
	String           func(*frizz.Root, frizz.Stack, interface{}) (String, error)
	Qual             func(*frizz.Root, frizz.Stack, interface{}) (Qual, error)
	Pointers         func(*frizz.Root, frizz.Stack, interface{}) (Pointers, error)
	Maps             func(*frizz.Root, frizz.Stack, interface{}) (Maps, error)
	Slices           func(*frizz.Root, frizz.Stack, interface{}) (Slices, error)
	Structs          func(*frizz.Root, frizz.Stack, interface{}) (Structs, error)
	Natives          func(*frizz.Root, frizz.Stack, interface{}) (Natives, error)
}{
	Ages:             unpack_Ages,
	Alias:            unpack_Alias,
	AliasArray:       unpack_AliasArray,
	AliasMap:         unpack_AliasMap,
	AliasPointer:     unpack_AliasPointer,
	AliasSlice:       unpack_AliasSlice,
	AliasSub:         unpack_AliasSub,
	Csv:              unpack_Csv,
	Custom:           unpack_Custom,
	EmbedNatives:     unpack_EmbedNatives,
	EmbedPointer:     unpack_EmbedPointer,
	EmbedQual:        unpack_EmbedQual,
	EmbedQualPointer: unpack_EmbedQualPointer,
	Impi:             unpack_Impi,
	Imps:             unpack_Imps,
	Int:              unpack_Int,
	Interface:        unpack_Interface,
	InterfaceField:   unpack_InterfaceField,
	Maps:             unpack_Maps,
	Natives:          unpack_Natives,
	Pointers:         unpack_Pointers,
	Private:          unpack_Private,
	Qual:             unpack_Qual,
	Slices:           unpack_Slices,
	String:           unpack_String,
	Structs:          unpack_Structs,
	Type:             unpack_Type,
}

func unpack_Ages(root *frizz.Root, stack frizz.Stack, in interface{}) (value Ages, err error) {
	out := new(Ages)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func unpack_Csv(root *frizz.Root, stack frizz.Stack, in interface{}) (value Csv, err error) {
	out := new(Csv)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func unpack_Type(root *frizz.Root, stack frizz.Stack, in interface{}) (value Type, err error) {
	out := new(Type)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func unpack_Custom(root *frizz.Root, stack frizz.Stack, in interface{}) (value Custom, err error) {
	out := new(Custom)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func unpack_EmbedNatives(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedNatives, err error) {
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
			out, err := unpack_Natives(root, stack, in)
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
func unpack_EmbedPointer(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedPointer, err error) {
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
				out, err := unpack_Natives(root, stack, in)
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
func unpack_EmbedQualPointer(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedQualPointer, err error) {
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
				out, err := sub.Unpackers.Sub(root, stack, in)
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
func unpack_EmbedQual(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedQual, err error) {
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
			out, err := sub.Unpackers.Sub(root, stack, in)
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
func unpack_InterfaceField(root *frizz.Root, stack frizz.Stack, in interface{}) (value InterfaceField, err error) {
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
			out, err := unpack_Interface(root, stack, in)
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
					out, err := unpack_Interface(root, stack, in)
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
					out, err := unpack_Interface(root, stack, in)
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
					out, err := unpack_Interface(root, stack, in)
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
func unpack_Impi(root *frizz.Root, stack frizz.Stack, in interface{}) (value Impi, err error) {
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
func unpack_Imps(root *frizz.Root, stack frizz.Stack, in interface{}) (value Imps, err error) {
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
func unpack_Interface(root *frizz.Root, stack frizz.Stack, in interface{}) (value Interface, err error) {
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
func unpack_Private(root *frizz.Root, stack frizz.Stack, in interface{}) (value Private, err error) {
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
func unpack_AliasSub(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasSub, err error) {
	// selectorUnpacker
	out, err := sub.Unpackers.Sub(root, stack, in)
	if err != nil {
		return value, err
	}
	return AliasSub(out), nil
}
func unpack_AliasSlice(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasSlice, err error) {
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
			out, err := unpack_Int(root, stack, in)
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
func unpack_AliasArray(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasArray, err error) {
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
func unpack_AliasMap(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasMap, err error) {
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
				out, err := unpack_Qual(root, stack, in)
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
func unpack_AliasPointer(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasPointer, err error) {
	// pointerUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, err error) {
		// localUnpacker
		out, err := unpack_Int(root, stack, in)
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
func unpack_Alias(root *frizz.Root, stack frizz.Stack, in interface{}) (value Alias, err error) {
	// localUnpacker
	out, err := unpack_Int(root, stack, in)
	if err != nil {
		return value, err
	}
	return Alias(out), nil
}
func unpack_Int(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, err error) {
	// nativeUnpacker
	out, err := frizz.UnpackInt(stack, in)
	if err != nil {
		return value, err
	}
	return Int(out), nil
}
func unpack_String(root *frizz.Root, stack frizz.Stack, in interface{}) (value String, err error) {
	// nativeUnpacker
	out, err := frizz.UnpackString(stack, in)
	if err != nil {
		return value, err
	}
	return String(out), nil
}
func unpack_Qual(root *frizz.Root, stack frizz.Stack, in interface{}) (value Qual, err error) {
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
			out, err := sub.Unpackers.Sub(root, stack, in)
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
func unpack_Pointers(root *frizz.Root, stack frizz.Stack, in interface{}) (value Pointers, err error) {
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
				out, err := unpack_Int(root, stack, in)
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
				out, err := sub.Unpackers.Sub(root, stack, in)
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
						out, err := unpack_Int(root, stack, in)
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
						out, err := sub.Unpackers.Sub(root, stack, in)
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
func unpack_Maps(root *frizz.Root, stack frizz.Stack, in interface{}) (value Maps, err error) {
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
func unpack_Slices(root *frizz.Root, stack frizz.Stack, in interface{}) (value Slices, err error) {
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
func unpack_Structs(root *frizz.Root, stack frizz.Stack, in interface{}) (value Structs, err error) {
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
func unpack_Natives(root *frizz.Root, stack frizz.Stack, in interface{}) (value Natives, err error) {
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
func init() {
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Ages", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Ages(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Csv", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Csv(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Type", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Type(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Custom", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Custom(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "EmbedNatives", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_EmbedNatives(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "EmbedPointer", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_EmbedPointer(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "EmbedQualPointer", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_EmbedQualPointer(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "EmbedQual", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_EmbedQual(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "InterfaceField", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_InterfaceField(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Impi", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Impi(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Imps", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Imps(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Interface", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Interface(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Private", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Private(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "AliasSub", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_AliasSub(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "AliasSlice", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_AliasSlice(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "AliasArray", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_AliasArray(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "AliasMap", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_AliasMap(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "AliasPointer", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_AliasPointer(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Alias", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Alias(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Int", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Int(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "String", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_String(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Qual", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Qual(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Pointers", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Pointers(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Maps", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Maps(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Slices", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Slices(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Structs", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Structs(root, stack, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Natives", func(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
		return unpack_Natives(root, stack, in)
	})
}
