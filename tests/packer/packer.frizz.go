package packer

import (
	frizz "frizz.io/frizz"
	sub "frizz.io/tests/packer/sub"
	errors "github.com/pkg/errors"
)

const Packer packer = 0

type packer int

func (p packer) Path() string {
	return "frizz.io/tests/packer"
}
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, err error) {
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
	// customUnpacker
	out := new(CustomSub)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func (p packer) UnpackAges(root *frizz.Root, stack frizz.Stack, in interface{}) (value Ages, err error) {
	// customUnpacker
	out := new(Ages)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func (p packer) UnpackCsv(root *frizz.Root, stack frizz.Stack, in interface{}) (value Csv, err error) {
	// customUnpacker
	out := new(Csv)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func (p packer) UnpackType(root *frizz.Root, stack frizz.Stack, in interface{}) (value Type, err error) {
	// customUnpacker
	out := new(Type)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func (p packer) UnpackCustom(root *frizz.Root, stack frizz.Stack, in interface{}) (value Custom, err error) {
	// customUnpacker
	out := new(Custom)
	if err := out.Unpack(root, stack, in); err != nil {
		return value, err
	}
	return *out, nil
}
func (p packer) UnpackEmbedNatives(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedNatives, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Natives
		Int int
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			Natives
			Int int
		}
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return EmbedNatives(out), nil
}
func (p packer) UnpackEmbedPointer(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedPointer, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		*Natives
		Int int
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			*Natives
			Int int
		}
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return EmbedPointer(out), nil
}
func (p packer) UnpackEmbedQualPointer(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedQualPointer, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		*sub.Sub
		Int int
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			*sub.Sub
			Int int
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return EmbedQualPointer(out), nil
}
func (p packer) UnpackEmbedQual(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedQual, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		sub.Sub
		Int int
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			sub.Sub
			Int int
		}
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return EmbedQual(out), nil
}
func (p packer) UnpackInterfaceField(root *frizz.Root, stack frizz.Stack, in interface{}) (value InterfaceField, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Iface Interface
		Slice []Interface
		Array [3]Interface
		Map   map[string]Interface
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			Iface Interface
			Slice []Interface
			Array [3]Interface
			Map   map[string]Interface
		}
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return InterfaceField(out), nil
}
func (p packer) UnpackImpi(root *frizz.Root, stack frizz.Stack, in interface{}) (value Impi, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Impi(out), nil
}
func (p packer) UnpackImps(root *frizz.Root, stack frizz.Stack, in interface{}) (value Imps, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		String string
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			String string
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Imps(out), nil
}
func (p packer) UnpackInterface(root *frizz.Root, stack frizz.Stack, in interface{}) (value Interface, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value interface {
		Foo() string
	}, err error) {
		// interfaceUnpacker
		out, err := root.UnpackInterface(stack, in)
		if err != nil {
			return value, err
		}
		iface, ok := out.(interface {
			Foo() string
		})
		if !ok {
			return value, errors.Errorf("unpacking into interface, type %T does not implement interface", out)
		}
		return iface, nil
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Interface(out), nil
}
func (p packer) UnpackPrivate(root *frizz.Root, stack frizz.Stack, in interface{}) (value Private, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		i int
		s string
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			i int
			s string
		}
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Private(out), nil
}
func (p packer) UnpackAliasSub(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasSub, err error) {
	// aliasUnpacker
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
	return AliasSub(out), nil
}
func (p packer) UnpackAliasSlice(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasSlice, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []Int, err error) {
		// sliceUnpacker
		a, ok := in.([]interface{})
		if !ok {
			return value, errors.New("unpacking into slice, value should be an array")
		}
		var out = make([]Int, len(a))
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return AliasSlice(out), nil
}
func (p packer) UnpackAliasArray(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasArray, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [3]string, err error) {
		// sliceUnpacker
		a, ok := in.([]interface{})
		if !ok {
			return value, errors.New("unpacking into slice, value should be an array")
		}
		var out [3]string
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return AliasArray(out), nil
}
func (p packer) UnpackAliasMap(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasMap, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]*Qual, err error) {
		// mapUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into map, value should be a map")
		}
		var out = make(map[string]*Qual, len(m))
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return AliasMap(out), nil
}
func (p packer) UnpackAliasPointer(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasPointer, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *Int, err error) {
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return AliasPointer(out), nil
}
func (p packer) UnpackAlias(root *frizz.Root, stack frizz.Stack, in interface{}) (value Alias, err error) {
	// aliasUnpacker
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
	return Alias(out), nil
}
func (p packer) UnpackInt(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, err error) {
		// nativeUnpacker
		out, err := frizz.UnpackInt(stack, in)
		if err != nil {
			return value, err
		}
		return out, nil
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Int(out), nil
}
func (p packer) UnpackString(root *frizz.Root, stack frizz.Stack, in interface{}) (value String, err error) {
	// aliasUnpacker
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
	return String(out), nil
}
func (p packer) UnpackQual(root *frizz.Root, stack frizz.Stack, in interface{}) (value Qual, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Sub sub.Sub
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			Sub sub.Sub
		}
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Qual(out), nil
}
func (p packer) UnpackPointers(root *frizz.Root, stack frizz.Stack, in interface{}) (value Pointers, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		String      *string
		Int         *Int
		Sub         *sub.Sub
		Array       *[3]int
		Slice       *[]string
		Map         *map[string]int
		SliceString []*string
		SliceInt    []*Int
		SliceSub    []*sub.Sub
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Pointers(out), nil
}
func (p packer) UnpackMaps(root *frizz.Root, stack frizz.Stack, in interface{}) (value Maps, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Ints    map[string]int
		Strings map[string]string
		Slices  map[string][]string
		Arrays  map[string][2]int
		Maps    map[string]map[string]string
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			Ints    map[string]int
			Strings map[string]string
			Slices  map[string][]string
			Arrays  map[string][2]int
			Maps    map[string]map[string]string
		}
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Maps(out), nil
}
func (p packer) UnpackSlices(root *frizz.Root, stack frizz.Stack, in interface{}) (value Slices, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Ints      []int
		Strings   []string
		ArrayLit  [5]string
		ArrayExpr [2 * N]int
		Structs   []struct {
			Int int
		}
		Arrays [][]string
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Slices(out), nil
}
func (p packer) UnpackStructs(root *frizz.Root, stack frizz.Stack, in interface{}) (value Structs, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
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
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Structs(out), nil
}
func (p packer) UnpackNatives(root *frizz.Root, stack frizz.Stack, in interface{}) (value Natives, err error) {
	// aliasUnpacker
	out, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Bool    bool
		Byte    byte
		Float32 float32
		Float64 float64
		Int     int
		Int8    int8
		Int16   int16
		Int32   int32
		Int64   int64
		Uint    uint
		Uint8   uint8
		Uint16  uint16
		Uint32  uint32
		Uint64  uint64
		Rune    rune
		String  string
	}, err error) {
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, errors.New("unpacking into struct, value should be a map")
		}
		var out struct {
			Bool    bool
			Byte    byte
			Float32 float32
			Float64 float64
			Int     int
			Int8    int8
			Int16   int16
			Int32   int32
			Int64   int64
			Uint    uint
			Uint8   uint8
			Uint16  uint16
			Uint32  uint32
			Uint64  uint64
			Rune    rune
			String  string
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
	}(root, stack, in)
	if err != nil {
		return value, err
	}
	return Natives(out), nil
}
func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "CustomSub":
		return p.RepackCustomSub(root, stack, in.(CustomSub))
	case "Ages":
		return p.RepackAges(root, stack, in.(Ages))
	case "Csv":
		return p.RepackCsv(root, stack, in.(Csv))
	case "Type":
		return p.RepackType(root, stack, in.(Type))
	case "Custom":
		return p.RepackCustom(root, stack, in.(Custom))
	case "EmbedNatives":
		return p.RepackEmbedNatives(root, stack, in.(EmbedNatives))
	case "EmbedPointer":
		return p.RepackEmbedPointer(root, stack, in.(EmbedPointer))
	case "EmbedQualPointer":
		return p.RepackEmbedQualPointer(root, stack, in.(EmbedQualPointer))
	case "EmbedQual":
		return p.RepackEmbedQual(root, stack, in.(EmbedQual))
	case "InterfaceField":
		return p.RepackInterfaceField(root, stack, in.(InterfaceField))
	case "Impi":
		return p.RepackImpi(root, stack, in.(Impi))
	case "Imps":
		return p.RepackImps(root, stack, in.(Imps))
	case "Interface":
		return p.RepackInterface(root, stack, in.(Interface))
	case "Private":
		return p.RepackPrivate(root, stack, in.(Private))
	case "AliasSub":
		return p.RepackAliasSub(root, stack, in.(AliasSub))
	case "AliasSlice":
		return p.RepackAliasSlice(root, stack, in.(AliasSlice))
	case "AliasArray":
		return p.RepackAliasArray(root, stack, in.(AliasArray))
	case "AliasMap":
		return p.RepackAliasMap(root, stack, in.(AliasMap))
	case "AliasPointer":
		return p.RepackAliasPointer(root, stack, in.(AliasPointer))
	case "Alias":
		return p.RepackAlias(root, stack, in.(Alias))
	case "Int":
		return p.RepackInt(root, stack, in.(Int))
	case "String":
		return p.RepackString(root, stack, in.(String))
	case "Qual":
		return p.RepackQual(root, stack, in.(Qual))
	case "Pointers":
		return p.RepackPointers(root, stack, in.(Pointers))
	case "Maps":
		return p.RepackMaps(root, stack, in.(Maps))
	case "Slices":
		return p.RepackSlices(root, stack, in.(Slices))
	case "Structs":
		return p.RepackStructs(root, stack, in.(Structs))
	case "Natives":
		return p.RepackNatives(root, stack, in.(Natives))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) RepackCustomSub(root *frizz.Root, stack frizz.Stack, in CustomSub) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(root, stack)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
}
func (p packer) RepackAges(root *frizz.Root, stack frizz.Stack, in Ages) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(root, stack)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
}
func (p packer) RepackCsv(root *frizz.Root, stack frizz.Stack, in Csv) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(root, stack)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
}
func (p packer) RepackType(root *frizz.Root, stack frizz.Stack, in Type) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(root, stack)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
}
func (p packer) RepackCustom(root *frizz.Root, stack frizz.Stack, in Custom) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(root, stack)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
}
func (p packer) RepackEmbedNatives(root *frizz.Root, stack frizz.Stack, in EmbedNatives) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		Natives
		Int int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in Natives) (value interface{}, dict bool, null bool, err error) {
			// localRepacker
			out, dict, null, err := p.RepackNatives(root, stack, in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.Natives); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Natives"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		Natives
		Int int
	})(in))
}
func (p packer) RepackEmbedPointer(root *frizz.Root, stack frizz.Stack, in EmbedPointer) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		*Natives
		Int int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *Natives) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in Natives) (value interface{}, dict bool, null bool, err error) {
				// localRepacker
				out, dict, null, err := p.RepackNatives(root, stack, in)
				if err != nil {
					return nil, false, false, err
				}
				return out, dict, null, nil
			}(root, stack, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.Natives); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Natives"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		*Natives
		Int int
	})(in))
}
func (p packer) RepackEmbedQualPointer(root *frizz.Root, stack frizz.Stack, in EmbedQualPointer) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		*sub.Sub
		Int int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *sub.Sub) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
				// selectorRepacker
				out, dict, null, err := sub.Packer.RepackSub(root, stack, in)
				if err != nil {
					return nil, false, false, err
				}
				return out, dict, null, nil
			}(root, stack, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.Sub); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Sub"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		*sub.Sub
		Int int
	})(in))
}
func (p packer) RepackEmbedQual(root *frizz.Root, stack frizz.Stack, in EmbedQual) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		sub.Sub
		Int int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
			// selectorRepacker
			out, dict, null, err := sub.Packer.RepackSub(root, stack, in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.Sub); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Sub"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		sub.Sub
		Int int
	})(in))
}
func (p packer) RepackInterfaceField(root *frizz.Root, stack frizz.Stack, in InterfaceField) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		Iface Interface
		Slice []Interface
		Array [3]Interface
		Map   map[string]Interface
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 5)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in Interface) (value interface{}, dict bool, null bool, err error) {
			// localRepacker
			out, dict, null, err := p.RepackInterface(root, stack, in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.Iface); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Iface"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in []Interface) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in Interface) (value interface{}, dict bool, null bool, err error) {
					// localRepacker
					out, dict, null, err := p.RepackInterface(root, stack, in)
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
		}(root, stack, in.Slice); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Slice"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in [3]Interface) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in Interface) (value interface{}, dict bool, null bool, err error) {
					// localRepacker
					out, dict, null, err := p.RepackInterface(root, stack, in)
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
		}(root, stack, in.Array); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Array"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in map[string]Interface) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(root *frizz.Root, stack frizz.Stack, in Interface) (value interface{}, dict bool, null bool, err error) {
					// localRepacker
					out, dict, null, err := p.RepackInterface(root, stack, in)
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
		}(root, stack, in.Map); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Map"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		Iface Interface
		Slice []Interface
		Array [3]Interface
		Map   map[string]Interface
	})(in))
}
func (p packer) RepackImpi(root *frizz.Root, stack frizz.Stack, in Impi) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		Int int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		Int int
	})(in))
}
func (p packer) RepackImps(root *frizz.Root, stack frizz.Stack, in Imps) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		String string
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackString(in)
		}(root, stack, in.String); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["String"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		String string
	})(in))
}
func (p packer) RepackInterface(root *frizz.Root, stack frizz.Stack, in Interface) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in interface {
		Foo() string
	}) (value interface{}, dict bool, null bool, err error) {
		// interfaceRepacker
		return root.RepackInterface(stack, false, in)
	}(root, stack, (interface {
		Foo() string
	})(in))
}
func (p packer) RepackPrivate(root *frizz.Root, stack frizz.Stack, in Private) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		i int
		s string
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.i); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["i"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackString(in)
		}(root, stack, in.s); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["s"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		i int
		s string
	})(in))
}
func (p packer) RepackAliasSub(root *frizz.Root, stack frizz.Stack, in AliasSub) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker
		out, dict, null, err := sub.Packer.RepackSub(root, stack, in)
		if err != nil {
			return nil, false, false, err
		}
		return out, dict, null, nil
	}(root, stack, (sub.Sub)(in))
}
func (p packer) RepackAliasSlice(root *frizz.Root, stack frizz.Stack, in AliasSlice) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in []Int) (value interface{}, dict bool, null bool, err error) {
		// sliceRepacker
		out := make([]interface{}, len(in))
		empty := true
		for i, item := range in {
			v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in Int) (value interface{}, dict bool, null bool, err error) {
				// localRepacker
				out, dict, null, err := p.RepackInt(root, stack, in)
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
	}(root, stack, ([]Int)(in))
}
func (p packer) RepackAliasArray(root *frizz.Root, stack frizz.Stack, in AliasArray) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in [3]string) (value interface{}, dict bool, null bool, err error) {
		// sliceRepacker
		out := make([]interface{}, len(in))
		empty := true
		for i, item := range in {
			v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return frizz.RepackString(in)
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
	}(root, stack, ([3]string)(in))
}
func (p packer) RepackAliasMap(root *frizz.Root, stack frizz.Stack, in AliasMap) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in map[string]*Qual) (value interface{}, dict bool, null bool, err error) {
		// mapRepacker
		out := make(map[string]interface{}, len(in))
		for k, item := range in {
			v, _, _, err := func(root *frizz.Root, stack frizz.Stack, in *Qual) (value interface{}, dict bool, null bool, err error) {
				// pointerRepacker
				out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in Qual) (value interface{}, dict bool, null bool, err error) {
					// localRepacker
					out, dict, null, err := p.RepackQual(root, stack, in)
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, null, nil
				}(root, stack, *in)
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
	}(root, stack, (map[string]*Qual)(in))
}
func (p packer) RepackAliasPointer(root *frizz.Root, stack frizz.Stack, in AliasPointer) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in *Int) (value interface{}, dict bool, null bool, err error) {
		// pointerRepacker
		out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in Int) (value interface{}, dict bool, null bool, err error) {
			// localRepacker
			out, dict, null, err := p.RepackInt(root, stack, in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, *in)
		if err != nil {
			return nil, false, false, err
		}
		return out, dict, null, nil
	}(root, stack, (*Int)(in))
}
func (p packer) RepackAlias(root *frizz.Root, stack frizz.Stack, in Alias) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in Int) (value interface{}, dict bool, null bool, err error) {
		// localRepacker
		out, dict, null, err := p.RepackInt(root, stack, in)
		if err != nil {
			return nil, false, false, err
		}
		return out, dict, null, nil
	}(root, stack, (Int)(in))
}
func (p packer) RepackInt(root *frizz.Root, stack frizz.Stack, in Int) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
		// nativeRepacker
		return frizz.RepackNumber(in)
	}(root, stack, (int)(in))
}
func (p packer) RepackString(root *frizz.Root, stack frizz.Stack, in String) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
		// nativeRepacker
		return frizz.RepackString(in)
	}(root, stack, (string)(in))
}
func (p packer) RepackQual(root *frizz.Root, stack frizz.Stack, in Qual) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		Sub sub.Sub
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
			// selectorRepacker
			out, dict, null, err := sub.Packer.RepackSub(root, stack, in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.Sub); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Sub"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		Sub sub.Sub
	})(in))
}
func (p packer) RepackPointers(root *frizz.Root, stack frizz.Stack, in Pointers) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
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
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *string) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return frizz.RepackString(in)
			}(root, stack, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.String); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["String"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *Int) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in Int) (value interface{}, dict bool, null bool, err error) {
				// localRepacker
				out, dict, null, err := p.RepackInt(root, stack, in)
				if err != nil {
					return nil, false, false, err
				}
				return out, dict, null, nil
			}(root, stack, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *sub.Sub) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
				// selectorRepacker
				out, dict, null, err := sub.Packer.RepackSub(root, stack, in)
				if err != nil {
					return nil, false, false, err
				}
				return out, dict, null, nil
			}(root, stack, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.Sub); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Sub"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *[3]int) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in [3]int) (value interface{}, dict bool, null bool, err error) {
				// sliceRepacker
				out := make([]interface{}, len(in))
				empty := true
				for i, item := range in {
					v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
						// nativeRepacker
						return frizz.RepackNumber(in)
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
			}(root, stack, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.Array); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Array"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *[]string) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in []string) (value interface{}, dict bool, null bool, err error) {
				// sliceRepacker
				out := make([]interface{}, len(in))
				empty := true
				for i, item := range in {
					v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
						// nativeRepacker
						return frizz.RepackString(in)
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
			}(root, stack, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.Slice); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Slice"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *map[string]int) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in map[string]int) (value interface{}, dict bool, null bool, err error) {
				// mapRepacker
				out := make(map[string]interface{}, len(in))
				for k, item := range in {
					v, _, _, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
						// nativeRepacker
						return frizz.RepackNumber(in)
					}(root, stack, item)
					if err != nil {
						return nil, false, false, err
					}
					out[k] = v
				}
				return out, true, len(in) == 0, nil
			}(root, stack, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.Map); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Map"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in []*string) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *string) (value interface{}, dict bool, null bool, err error) {
					// pointerRepacker
					out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
						// nativeRepacker
						return frizz.RepackString(in)
					}(root, stack, *in)
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
		}(root, stack, in.SliceString); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["SliceString"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in []*Int) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *Int) (value interface{}, dict bool, null bool, err error) {
					// pointerRepacker
					out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in Int) (value interface{}, dict bool, null bool, err error) {
						// localRepacker
						out, dict, null, err := p.RepackInt(root, stack, in)
						if err != nil {
							return nil, false, false, err
						}
						return out, dict, null, nil
					}(root, stack, *in)
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
		}(root, stack, in.SliceInt); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["SliceInt"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in []*sub.Sub) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *sub.Sub) (value interface{}, dict bool, null bool, err error) {
					// pointerRepacker
					out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in sub.Sub) (value interface{}, dict bool, null bool, err error) {
						// selectorRepacker
						out, dict, null, err := sub.Packer.RepackSub(root, stack, in)
						if err != nil {
							return nil, false, false, err
						}
						return out, dict, null, nil
					}(root, stack, *in)
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
		}(root, stack, in.SliceSub); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["SliceSub"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
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
func (p packer) RepackMaps(root *frizz.Root, stack frizz.Stack, in Maps) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		Ints    map[string]int
		Strings map[string]string
		Slices  map[string][]string
		Arrays  map[string][2]int
		Maps    map[string]map[string]string
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 6)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in map[string]int) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return frizz.RepackNumber(in)
				}(root, stack, item)
				if err != nil {
					return nil, false, false, err
				}
				out[k] = v
			}
			return out, true, len(in) == 0, nil
		}(root, stack, in.Ints); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Ints"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in map[string]string) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return frizz.RepackString(in)
				}(root, stack, item)
				if err != nil {
					return nil, false, false, err
				}
				out[k] = v
			}
			return out, true, len(in) == 0, nil
		}(root, stack, in.Strings); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Strings"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in map[string][]string) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(root *frizz.Root, stack frizz.Stack, in []string) (value interface{}, dict bool, null bool, err error) {
					// sliceRepacker
					out := make([]interface{}, len(in))
					empty := true
					for i, item := range in {
						v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
							// nativeRepacker
							return frizz.RepackString(in)
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
		}(root, stack, in.Slices); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Slices"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in map[string][2]int) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(root *frizz.Root, stack frizz.Stack, in [2]int) (value interface{}, dict bool, null bool, err error) {
					// sliceRepacker
					out := make([]interface{}, len(in))
					empty := true
					for i, item := range in {
						v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
							// nativeRepacker
							return frizz.RepackNumber(in)
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
		}(root, stack, in.Arrays); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Arrays"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in map[string]map[string]string) (value interface{}, dict bool, null bool, err error) {
			// mapRepacker
			out := make(map[string]interface{}, len(in))
			for k, item := range in {
				v, _, _, err := func(root *frizz.Root, stack frizz.Stack, in map[string]string) (value interface{}, dict bool, null bool, err error) {
					// mapRepacker
					out := make(map[string]interface{}, len(in))
					for k, item := range in {
						v, _, _, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
							// nativeRepacker
							return frizz.RepackString(in)
						}(root, stack, item)
						if err != nil {
							return nil, false, false, err
						}
						out[k] = v
					}
					return out, true, len(in) == 0, nil
				}(root, stack, item)
				if err != nil {
					return nil, false, false, err
				}
				out[k] = v
			}
			return out, true, len(in) == 0, nil
		}(root, stack, in.Maps); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Maps"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		Ints    map[string]int
		Strings map[string]string
		Slices  map[string][]string
		Arrays  map[string][2]int
		Maps    map[string]map[string]string
	})(in))
}
func (p packer) RepackSlices(root *frizz.Root, stack frizz.Stack, in Slices) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
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
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in []int) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return frizz.RepackNumber(in)
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
		}(root, stack, in.Ints); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Ints"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in []string) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return frizz.RepackString(in)
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
		}(root, stack, in.Strings); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Strings"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in [5]string) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return frizz.RepackString(in)
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
		}(root, stack, in.ArrayLit); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["ArrayLit"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in [2 * N]int) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return frizz.RepackNumber(in)
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
		}(root, stack, in.ArrayExpr); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["ArrayExpr"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in []struct {
			Int int
		}) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in struct {
					Int int
				}) (value interface{}, dict bool, null bool, err error) {
					// structRepacker
					out := make(map[string]interface{}, 2)
					empty := true
					if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
						// nativeRepacker
						return frizz.RepackNumber(in)
					}(root, stack, in.Int); err != nil {
						return nil, false, false, err
					} else {
						if !null {
							empty = false
							out["Int"] = v
						}
					}
					return out, false, empty, nil
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
		}(root, stack, in.Structs); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Structs"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in [][]string) (value interface{}, dict bool, null bool, err error) {
			// sliceRepacker
			out := make([]interface{}, len(in))
			empty := true
			for i, item := range in {
				v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in []string) (value interface{}, dict bool, null bool, err error) {
					// sliceRepacker
					out := make([]interface{}, len(in))
					empty := true
					for i, item := range in {
						v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
							// nativeRepacker
							return frizz.RepackString(in)
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
				if !null {
					empty = false
				}
				out[i] = v
			}
			return out, false, empty, nil
		}(root, stack, in.Arrays); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Arrays"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
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
func (p packer) RepackStructs(root *frizz.Root, stack frizz.Stack, in Structs) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
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
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in struct {
			Int  int
			Bool bool
		}) (value interface{}, dict bool, null bool, err error) {
			// structRepacker
			out := make(map[string]interface{}, 3)
			empty := true
			if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return frizz.RepackNumber(in)
			}(root, stack, in.Int); err != nil {
				return nil, false, false, err
			} else {
				if !null {
					empty = false
					out["Int"] = v
				}
			}
			if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in bool) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return frizz.RepackBool(in)
			}(root, stack, in.Bool); err != nil {
				return nil, false, false, err
			} else {
				if !null {
					empty = false
					out["Bool"] = v
				}
			}
			return out, false, empty, nil
		}(root, stack, in.Simple); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Simple"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in struct {
			String string
			Inner  struct {
				Float32 float32
			}
		}) (value interface{}, dict bool, null bool, err error) {
			// structRepacker
			out := make(map[string]interface{}, 3)
			empty := true
			if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return frizz.RepackString(in)
			}(root, stack, in.String); err != nil {
				return nil, false, false, err
			} else {
				if !null {
					empty = false
					out["String"] = v
				}
			}
			if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in struct {
				Float32 float32
			}) (value interface{}, dict bool, null bool, err error) {
				// structRepacker
				out := make(map[string]interface{}, 2)
				empty := true
				if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in float32) (value interface{}, dict bool, null bool, err error) {
					// nativeRepacker
					return frizz.RepackNumber(in)
				}(root, stack, in.Float32); err != nil {
					return nil, false, false, err
				} else {
					if !null {
						empty = false
						out["Float32"] = v
					}
				}
				return out, false, empty, nil
			}(root, stack, in.Inner); err != nil {
				return nil, false, false, err
			} else {
				if !null {
					empty = false
					out["Inner"] = v
				}
			}
			return out, false, empty, nil
		}(root, stack, in.Complex); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Complex"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
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
func (p packer) RepackNatives(root *frizz.Root, stack frizz.Stack, in Natives) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		Bool    bool
		Byte    byte
		Float32 float32
		Float64 float64
		Int     int
		Int8    int8
		Int16   int16
		Int32   int32
		Int64   int64
		Uint    uint
		Uint8   uint8
		Uint16  uint16
		Uint32  uint32
		Uint64  uint64
		Rune    rune
		String  string
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 17)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in bool) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackBool(in)
		}(root, stack, in.Bool); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Bool"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in byte) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Byte); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Byte"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in float32) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Float32); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Float32"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in float64) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Float64); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Float64"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int8) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Int8); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int8"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int16) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Int16); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int16"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int32) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Int32); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int32"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int64) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Int64); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int64"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in uint) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Uint); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Uint"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in uint8) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Uint8); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Uint8"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in uint16) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Uint16); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Uint16"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in uint32) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Uint32); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Uint32"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in uint64) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Uint64); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Uint64"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in rune) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackString(in)
		}(root, stack, in.Rune); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Rune"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackString(in)
		}(root, stack, in.String); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["String"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		Bool    bool
		Byte    byte
		Float32 float32
		Float64 float64
		Int     int
		Int8    int8
		Int16   int16
		Int32   int32
		Int64   int64
		Uint    uint
		Uint8   uint8
		Uint16  uint16
		Uint32  uint32
		Uint64  uint64
		Rune    rune
		String  string
	})(in))
}
