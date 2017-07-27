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
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Csv":
		return p.UnpackCsv(root, stack, in)
	case "EmbedNatives":
		return p.UnpackEmbedNatives(root, stack, in)
	case "AliasSlice":
		return p.UnpackAliasSlice(root, stack, in)
	case "AliasPointer":
		return p.UnpackAliasPointer(root, stack, in)
	case "Structs":
		return p.UnpackStructs(root, stack, in)
	case "EmbedQualPointer":
		return p.UnpackEmbedQualPointer(root, stack, in)
	case "Private":
		return p.UnpackPrivate(root, stack, in)
	case "AliasMap":
		return p.UnpackAliasMap(root, stack, in)
	case "String":
		return p.UnpackString(root, stack, in)
	case "Ages":
		return p.UnpackAges(root, stack, in)
	case "Type":
		return p.UnpackType(root, stack, in)
	case "EmbedQual":
		return p.UnpackEmbedQual(root, stack, in)
	case "Maps":
		return p.UnpackMaps(root, stack, in)
	case "Slices":
		return p.UnpackSlices(root, stack, in)
	case "Natives":
		return p.UnpackNatives(root, stack, in)
	case "CustomSub":
		return p.UnpackCustomSub(root, stack, in)
	case "Impi":
		return p.UnpackImpi(root, stack, in)
	case "AliasSub":
		return p.UnpackAliasSub(root, stack, in)
	case "Alias":
		return p.UnpackAlias(root, stack, in)
	case "Int":
		return p.UnpackInt(root, stack, in)
	case "Qual":
		return p.UnpackQual(root, stack, in)
	case "Custom":
		return p.UnpackCustom(root, stack, in)
	case "EmbedPointer":
		return p.UnpackEmbedPointer(root, stack, in)
	case "SubInterface":
		return p.UnpackSubInterface(root, stack, in)
	case "Pointers":
		return p.UnpackPointers(root, stack, in)
	case "InterfaceField":
		return p.UnpackInterfaceField(root, stack, in)
	case "Imps":
		return p.UnpackImps(root, stack, in)
	case "Interface":
		return p.UnpackInterface(root, stack, in)
	case "AliasArray":
		return p.UnpackAliasArray(root, stack, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) UnpackStructs(root *frizz.Root, stack frizz.Stack, in interface{}) (value Structs, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
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
			return value, false, errors.New("unpacking into struct, value should be a map")
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
			stack := stack.Append(frizz.FieldItem("Simple"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
				Int  int
				Bool bool
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
					Int  int
					Bool bool
				}
				if v, ok := m["Int"]; ok {
					stack := stack.Append(frizz.FieldItem("Int"))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// nativeUnpacker
						out, null, err := frizz.UnpackInt(stack, in)
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
						out.Int = u
					}
				}
				if v, ok := m["Bool"]; ok {
					stack := stack.Append(frizz.FieldItem("Bool"))
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
						out.Bool = u
					}
				}
				return out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Simple = u
			}
		}
		if v, ok := m["Complex"]; ok {
			stack := stack.Append(frizz.FieldItem("Complex"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
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
					return value, false, errors.New("unpacking into struct, value should be a map")
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
					stack := stack.Append(frizz.FieldItem("String"))
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
						out.String = u
					}
				}
				if v, ok := m["Inner"]; ok {
					stack := stack.Append(frizz.FieldItem("Inner"))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
						Float32 float32
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
							Float32 float32
						}
						if v, ok := m["Float32"]; ok {
							stack := stack.Append(frizz.FieldItem("Float32"))
							u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value float32, null bool, err error) {
								if in == nil {
									return value, true, nil
								}
								// nativeUnpacker
								out, null, err := frizz.UnpackFloat32(stack, in)
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
								out.Float32 = u
							}
						}
						return out, false, nil
					}(root, stack, v)
					if err != nil {
						return value, false, err
					}
					if !null {
						out.Inner = u
					}
				}
				return out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Complex = u
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
	return Structs(out), false, nil
}
func (p packer) UnpackCsv(root *frizz.Root, stack frizz.Stack, in interface{}) (value Csv, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// customUnpacker
	out := new(Csv)
	null, err = out.Unpack(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return *out, false, nil
}
func (p packer) UnpackEmbedNatives(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedNatives, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Natives
		Int int
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
			Natives
			Int int
		}
		if v, ok := m["Natives"]; ok {
			stack := stack.Append(frizz.FieldItem("Natives"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Natives, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// localUnpacker
				out, null, err := p.UnpackNatives(root, stack, in)
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
				out.Natives = u
			}
		}
		if v, ok := m["Int"]; ok {
			stack := stack.Append(frizz.FieldItem("Int"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackInt(stack, in)
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
				out.Int = u
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
	return EmbedNatives(out), false, nil
}
func (p packer) UnpackAliasSlice(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasSlice, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []Int, null bool, err error) {
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
		var out = make([]Int, len(a))
		for i, v := range a {
			stack := stack.Append(frizz.ArrayItem(i))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// localUnpacker
				out, null, err := p.UnpackInt(root, stack, in)
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
	return AliasSlice(out), false, nil
}
func (p packer) UnpackAliasPointer(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasPointer, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *Int, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// pointerUnpacker
		out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, null bool, err error) {
			if in == nil {
				return value, true, nil
			}
			// localUnpacker
			out, null, err := p.UnpackInt(root, stack, in)
			if err != nil {
				return value, false, err
			}
			if null {
				return value, true, nil
			}
			return out, false, nil
		}(root, stack, in)
		if err != nil {
			return value, false, err
		}
		if null {
			return value, true, nil
		}
		return &out, false, nil
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return AliasPointer(out), false, nil
}
func (p packer) UnpackEmbedQualPointer(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedQualPointer, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		*sub.Sub
		Int int
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
			*sub.Sub
			Int int
		}
		if v, ok := m["Sub"]; ok {
			stack := stack.Append(frizz.FieldItem("Sub"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *sub.Sub, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value sub.Sub, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// selectorUnpacker
					out, null, err := sub.Packer.UnpackSub(root, stack, in)
					if err != nil {
						return value, false, err
					}
					if null {
						return value, true, nil
					}
					return out, false, nil
				}(root, stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return &out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Sub = u
			}
		}
		if v, ok := m["Int"]; ok {
			stack := stack.Append(frizz.FieldItem("Int"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackInt(stack, in)
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
				out.Int = u
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
	return EmbedQualPointer(out), false, nil
}
func (p packer) UnpackPrivate(root *frizz.Root, stack frizz.Stack, in interface{}) (value Private, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		i int
		s string
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
			i int
			s string
		}
		if v, ok := m["i"]; ok {
			stack := stack.Append(frizz.FieldItem("i"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackInt(stack, in)
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
				out.i = u
			}
		}
		if v, ok := m["s"]; ok {
			stack := stack.Append(frizz.FieldItem("s"))
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
				out.s = u
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
	return Private(out), false, nil
}
func (p packer) UnpackAliasMap(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasMap, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]*Qual, null bool, err error) {
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
		var out = make(map[string]*Qual, len(m))
		for k, v := range m {
			stack := stack.Append(frizz.MapItem(k))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *Qual, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Qual, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// localUnpacker
					out, null, err := p.UnpackQual(root, stack, in)
					if err != nil {
						return value, false, err
					}
					if null {
						return value, true, nil
					}
					return out, false, nil
				}(root, stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return &out, false, nil
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
	return AliasMap(out), false, nil
}
func (p packer) UnpackString(root *frizz.Root, stack frizz.Stack, in interface{}) (value String, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value string, null bool, err error) {
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
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return String(out), false, nil
}
func (p packer) UnpackSlices(root *frizz.Root, stack frizz.Stack, in interface{}) (value Slices, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
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
			return value, false, errors.New("unpacking into struct, value should be a map")
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
			stack := stack.Append(frizz.FieldItem("Ints"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []int, null bool, err error) {
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
				var out = make([]int, len(a))
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// nativeUnpacker
						out, null, err := frizz.UnpackInt(stack, in)
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
				out.Ints = u
			}
		}
		if v, ok := m["Strings"]; ok {
			stack := stack.Append(frizz.FieldItem("Strings"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []string, null bool, err error) {
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
				var out = make([]string, len(a))
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
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
						out[i] = u
					}
				}
				return out[:], false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Strings = u
			}
		}
		if v, ok := m["ArrayLit"]; ok {
			stack := stack.Append(frizz.FieldItem("ArrayLit"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [5]string, null bool, err error) {
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
				var out [5]string
				if len(a) > 5 {
					return value, false, errors.Errorf("data length %d does not fit in array of length %d", len(a), 5)
				}
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
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
						out[i] = u
					}
				}
				return out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.ArrayLit = u
			}
		}
		if v, ok := m["ArrayExpr"]; ok {
			stack := stack.Append(frizz.FieldItem("ArrayExpr"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [2 * N]int, null bool, err error) {
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
				var out [2 * N]int
				if len(a) > 2*N {
					return value, false, errors.Errorf("data length %d does not fit in array of length %d", len(a), 2*N)
				}
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// nativeUnpacker
						out, null, err := frizz.UnpackInt(stack, in)
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
				return out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.ArrayExpr = u
			}
		}
		if v, ok := m["Structs"]; ok {
			stack := stack.Append(frizz.FieldItem("Structs"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []struct {
				Int int
			}, null bool, err error) {
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
				var out = make([]struct {
					Int int
				}, len(a))
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
						Int int
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
							Int int
						}
						if v, ok := m["Int"]; ok {
							stack := stack.Append(frizz.FieldItem("Int"))
							u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
								if in == nil {
									return value, true, nil
								}
								// nativeUnpacker
								out, null, err := frizz.UnpackInt(stack, in)
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
								out.Int = u
							}
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
				out.Structs = u
			}
		}
		if v, ok := m["Arrays"]; ok {
			stack := stack.Append(frizz.FieldItem("Arrays"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [][]string, null bool, err error) {
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
				var out = make([][]string, len(a))
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []string, null bool, err error) {
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
						var out = make([]string, len(a))
						for i, v := range a {
							stack := stack.Append(frizz.ArrayItem(i))
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
								out[i] = u
							}
						}
						return out[:], false, nil
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
				out.Arrays = u
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
	return Slices(out), false, nil
}
func (p packer) UnpackNatives(root *frizz.Root, stack frizz.Stack, in interface{}) (value Natives, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
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
		}
		if v, ok := m["Bool"]; ok {
			stack := stack.Append(frizz.FieldItem("Bool"))
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
				out.Bool = u
			}
		}
		if v, ok := m["Byte"]; ok {
			stack := stack.Append(frizz.FieldItem("Byte"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value byte, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackByte(stack, in)
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
				out.Byte = u
			}
		}
		if v, ok := m["Float32"]; ok {
			stack := stack.Append(frizz.FieldItem("Float32"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value float32, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackFloat32(stack, in)
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
				out.Float32 = u
			}
		}
		if v, ok := m["Float64"]; ok {
			stack := stack.Append(frizz.FieldItem("Float64"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value float64, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackFloat64(stack, in)
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
				out.Float64 = u
			}
		}
		if v, ok := m["Int"]; ok {
			stack := stack.Append(frizz.FieldItem("Int"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackInt(stack, in)
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
				out.Int = u
			}
		}
		if v, ok := m["Int8"]; ok {
			stack := stack.Append(frizz.FieldItem("Int8"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int8, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackInt8(stack, in)
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
				out.Int8 = u
			}
		}
		if v, ok := m["Int16"]; ok {
			stack := stack.Append(frizz.FieldItem("Int16"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int16, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackInt16(stack, in)
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
				out.Int16 = u
			}
		}
		if v, ok := m["Int32"]; ok {
			stack := stack.Append(frizz.FieldItem("Int32"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int32, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackInt32(stack, in)
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
				out.Int32 = u
			}
		}
		if v, ok := m["Int64"]; ok {
			stack := stack.Append(frizz.FieldItem("Int64"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int64, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackInt64(stack, in)
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
				out.Int64 = u
			}
		}
		if v, ok := m["Uint"]; ok {
			stack := stack.Append(frizz.FieldItem("Uint"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value uint, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackUint(stack, in)
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
				out.Uint = u
			}
		}
		if v, ok := m["Uint8"]; ok {
			stack := stack.Append(frizz.FieldItem("Uint8"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value uint8, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackUint8(stack, in)
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
				out.Uint8 = u
			}
		}
		if v, ok := m["Uint16"]; ok {
			stack := stack.Append(frizz.FieldItem("Uint16"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value uint16, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackUint16(stack, in)
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
				out.Uint16 = u
			}
		}
		if v, ok := m["Uint32"]; ok {
			stack := stack.Append(frizz.FieldItem("Uint32"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value uint32, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackUint32(stack, in)
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
				out.Uint32 = u
			}
		}
		if v, ok := m["Uint64"]; ok {
			stack := stack.Append(frizz.FieldItem("Uint64"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value uint64, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackUint64(stack, in)
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
				out.Uint64 = u
			}
		}
		if v, ok := m["Rune"]; ok {
			stack := stack.Append(frizz.FieldItem("Rune"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value rune, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackRune(stack, in)
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
				out.Rune = u
			}
		}
		if v, ok := m["String"]; ok {
			stack := stack.Append(frizz.FieldItem("String"))
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
				out.String = u
			}
		}
		if v, ok := m["PtrString"]; ok {
			stack := stack.Append(frizz.FieldItem("PtrString"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value string, null bool, err error) {
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
				}(root, stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return &out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.PtrString = u
			}
		}
		if v, ok := m["PtrInt"]; ok {
			stack := stack.Append(frizz.FieldItem("PtrInt"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// nativeUnpacker
					out, null, err := frizz.UnpackInt(stack, in)
					if err != nil {
						return value, false, err
					}
					if null {
						return value, true, nil
					}
					return out, false, nil
				}(root, stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return &out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.PtrInt = u
			}
		}
		if v, ok := m["PtrBool"]; ok {
			stack := stack.Append(frizz.FieldItem("PtrBool"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *bool, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value bool, null bool, err error) {
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
				}(root, stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return &out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.PtrBool = u
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
	return Natives(out), false, nil
}
func (p packer) UnpackAges(root *frizz.Root, stack frizz.Stack, in interface{}) (value Ages, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// customUnpacker
	out := new(Ages)
	null, err = out.Unpack(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return *out, false, nil
}
func (p packer) UnpackType(root *frizz.Root, stack frizz.Stack, in interface{}) (value Type, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// customUnpacker
	out := new(Type)
	null, err = out.Unpack(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return *out, false, nil
}
func (p packer) UnpackEmbedQual(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedQual, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		sub.Sub
		Int int
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
			sub.Sub
			Int int
		}
		if v, ok := m["Sub"]; ok {
			stack := stack.Append(frizz.FieldItem("Sub"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value sub.Sub, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// selectorUnpacker
				out, null, err := sub.Packer.UnpackSub(root, stack, in)
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
				out.Sub = u
			}
		}
		if v, ok := m["Int"]; ok {
			stack := stack.Append(frizz.FieldItem("Int"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackInt(stack, in)
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
				out.Int = u
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
	return EmbedQual(out), false, nil
}
func (p packer) UnpackMaps(root *frizz.Root, stack frizz.Stack, in interface{}) (value Maps, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
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
			return value, false, errors.New("unpacking into struct, value should be a map")
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
			stack := stack.Append(frizz.FieldItem("Ints"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]int, null bool, err error) {
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
				var out = make(map[string]int, len(m))
				for k, v := range m {
					stack := stack.Append(frizz.MapItem(k))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// nativeUnpacker
						out, null, err := frizz.UnpackInt(stack, in)
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
				out.Ints = u
			}
		}
		if v, ok := m["Strings"]; ok {
			stack := stack.Append(frizz.FieldItem("Strings"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]string, null bool, err error) {
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
				var out = make(map[string]string, len(m))
				for k, v := range m {
					stack := stack.Append(frizz.MapItem(k))
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
						out[k] = u
					}
				}
				return out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Strings = u
			}
		}
		if v, ok := m["Slices"]; ok {
			stack := stack.Append(frizz.FieldItem("Slices"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string][]string, null bool, err error) {
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
				var out = make(map[string][]string, len(m))
				for k, v := range m {
					stack := stack.Append(frizz.MapItem(k))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []string, null bool, err error) {
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
						var out = make([]string, len(a))
						for i, v := range a {
							stack := stack.Append(frizz.ArrayItem(i))
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
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Slices = u
			}
		}
		if v, ok := m["Arrays"]; ok {
			stack := stack.Append(frizz.FieldItem("Arrays"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string][2]int, null bool, err error) {
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
				var out = make(map[string][2]int, len(m))
				for k, v := range m {
					stack := stack.Append(frizz.MapItem(k))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [2]int, null bool, err error) {
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
						var out [2]int
						if len(a) > 2 {
							return value, false, errors.Errorf("data length %d does not fit in array of length %d", len(a), 2)
						}
						for i, v := range a {
							stack := stack.Append(frizz.ArrayItem(i))
							u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
								if in == nil {
									return value, true, nil
								}
								// nativeUnpacker
								out, null, err := frizz.UnpackInt(stack, in)
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
				out.Arrays = u
			}
		}
		if v, ok := m["Maps"]; ok {
			stack := stack.Append(frizz.FieldItem("Maps"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]map[string]string, null bool, err error) {
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
				var out = make(map[string]map[string]string, len(m))
				for k, v := range m {
					stack := stack.Append(frizz.MapItem(k))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]string, null bool, err error) {
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
						var out = make(map[string]string, len(m))
						for k, v := range m {
							stack := stack.Append(frizz.MapItem(k))
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
								out[k] = u
							}
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
				out.Maps = u
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
	return Maps(out), false, nil
}
func (p packer) UnpackInt(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// nativeUnpacker
		out, null, err := frizz.UnpackInt(stack, in)
		if err != nil {
			return value, false, err
		}
		if null {
			return value, true, nil
		}
		return out, false, nil
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return Int(out), false, nil
}
func (p packer) UnpackQual(root *frizz.Root, stack frizz.Stack, in interface{}) (value Qual, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Sub sub.Sub
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
			Sub sub.Sub
		}
		if v, ok := m["Sub"]; ok {
			stack := stack.Append(frizz.FieldItem("Sub"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value sub.Sub, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// selectorUnpacker
				out, null, err := sub.Packer.UnpackSub(root, stack, in)
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
				out.Sub = u
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
	return Qual(out), false, nil
}
func (p packer) UnpackCustomSub(root *frizz.Root, stack frizz.Stack, in interface{}) (value CustomSub, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// customUnpacker
	out := new(CustomSub)
	null, err = out.Unpack(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return *out, false, nil
}
func (p packer) UnpackImpi(root *frizz.Root, stack frizz.Stack, in interface{}) (value Impi, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		Int int
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
			Int int
		}
		if v, ok := m["Int"]; ok {
			stack := stack.Append(frizz.FieldItem("Int"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackInt(stack, in)
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
				out.Int = u
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
	return Impi(out), false, nil
}
func (p packer) UnpackAliasSub(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasSub, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value sub.Sub, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker
		out, null, err := sub.Packer.UnpackSub(root, stack, in)
		if err != nil {
			return value, false, err
		}
		if null {
			return value, true, nil
		}
		return out, false, nil
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return AliasSub(out), false, nil
}
func (p packer) UnpackAlias(root *frizz.Root, stack frizz.Stack, in interface{}) (value Alias, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// localUnpacker
		out, null, err := p.UnpackInt(root, stack, in)
		if err != nil {
			return value, false, err
		}
		if null {
			return value, true, nil
		}
		return out, false, nil
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return Alias(out), false, nil
}
func (p packer) UnpackCustom(root *frizz.Root, stack frizz.Stack, in interface{}) (value Custom, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// customUnpacker
	out := new(Custom)
	null, err = out.Unpack(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return *out, false, nil
}
func (p packer) UnpackEmbedPointer(root *frizz.Root, stack frizz.Stack, in interface{}) (value EmbedPointer, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		*Natives
		Int int
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
			*Natives
			Int int
		}
		if v, ok := m["Natives"]; ok {
			stack := stack.Append(frizz.FieldItem("Natives"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *Natives, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Natives, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// localUnpacker
					out, null, err := p.UnpackNatives(root, stack, in)
					if err != nil {
						return value, false, err
					}
					if null {
						return value, true, nil
					}
					return out, false, nil
				}(root, stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return &out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Natives = u
			}
		}
		if v, ok := m["Int"]; ok {
			stack := stack.Append(frizz.FieldItem("Int"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackInt(stack, in)
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
				out.Int = u
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
	return EmbedPointer(out), false, nil
}
func (p packer) UnpackSubInterface(root *frizz.Root, stack frizz.Stack, in interface{}) (value SubInterface, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		SubInterface sub.SubInterface
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
			SubInterface sub.SubInterface
		}
		if v, ok := m["SubInterface"]; ok {
			stack := stack.Append(frizz.FieldItem("SubInterface"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value sub.SubInterface, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// selectorUnpacker
				out, null, err := sub.Packer.UnpackSubInterface(root, stack, in)
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
				out.SubInterface = u
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
	return SubInterface(out), false, nil
}
func (p packer) UnpackPointers(root *frizz.Root, stack frizz.Stack, in interface{}) (value Pointers, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
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
			return value, false, errors.New("unpacking into struct, value should be a map")
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
			stack := stack.Append(frizz.FieldItem("String"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value string, null bool, err error) {
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
				}(root, stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return &out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.String = u
			}
		}
		if v, ok := m["Int"]; ok {
			stack := stack.Append(frizz.FieldItem("Int"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *Int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// localUnpacker
					out, null, err := p.UnpackInt(root, stack, in)
					if err != nil {
						return value, false, err
					}
					if null {
						return value, true, nil
					}
					return out, false, nil
				}(root, stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return &out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int = u
			}
		}
		if v, ok := m["Sub"]; ok {
			stack := stack.Append(frizz.FieldItem("Sub"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *sub.Sub, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value sub.Sub, null bool, err error) {
					if in == nil {
						return value, true, nil
					}
					// selectorUnpacker
					out, null, err := sub.Packer.UnpackSub(root, stack, in)
					if err != nil {
						return value, false, err
					}
					if null {
						return value, true, nil
					}
					return out, false, nil
				}(root, stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return &out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Sub = u
			}
		}
		if v, ok := m["Array"]; ok {
			stack := stack.Append(frizz.FieldItem("Array"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *[3]int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [3]int, null bool, err error) {
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
					var out [3]int
					if len(a) > 3 {
						return value, false, errors.Errorf("data length %d does not fit in array of length %d", len(a), 3)
					}
					for i, v := range a {
						stack := stack.Append(frizz.ArrayItem(i))
						u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
							if in == nil {
								return value, true, nil
							}
							// nativeUnpacker
							out, null, err := frizz.UnpackInt(stack, in)
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
					return out, false, nil
				}(root, stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return &out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Array = u
			}
		}
		if v, ok := m["Slice"]; ok {
			stack := stack.Append(frizz.FieldItem("Slice"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *[]string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []string, null bool, err error) {
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
					var out = make([]string, len(a))
					for i, v := range a {
						stack := stack.Append(frizz.ArrayItem(i))
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
				return &out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Slice = u
			}
		}
		if v, ok := m["Map"]; ok {
			stack := stack.Append(frizz.FieldItem("Map"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *map[string]int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// pointerUnpacker
				out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]int, null bool, err error) {
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
					var out = make(map[string]int, len(m))
					for k, v := range m {
						stack := stack.Append(frizz.MapItem(k))
						u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
							if in == nil {
								return value, true, nil
							}
							// nativeUnpacker
							out, null, err := frizz.UnpackInt(stack, in)
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
				}(root, stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return &out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Map = u
			}
		}
		if v, ok := m["SliceString"]; ok {
			stack := stack.Append(frizz.FieldItem("SliceString"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []*string, null bool, err error) {
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
				var out = make([]*string, len(a))
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *string, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// pointerUnpacker
						out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value string, null bool, err error) {
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
						}(root, stack, in)
						if err != nil {
							return value, false, err
						}
						if null {
							return value, true, nil
						}
						return &out, false, nil
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
				out.SliceString = u
			}
		}
		if v, ok := m["SliceInt"]; ok {
			stack := stack.Append(frizz.FieldItem("SliceInt"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []*Int, null bool, err error) {
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
				var out = make([]*Int, len(a))
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *Int, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// pointerUnpacker
						out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Int, null bool, err error) {
							if in == nil {
								return value, true, nil
							}
							// localUnpacker
							out, null, err := p.UnpackInt(root, stack, in)
							if err != nil {
								return value, false, err
							}
							if null {
								return value, true, nil
							}
							return out, false, nil
						}(root, stack, in)
						if err != nil {
							return value, false, err
						}
						if null {
							return value, true, nil
						}
						return &out, false, nil
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
				out.SliceInt = u
			}
		}
		if v, ok := m["SliceSub"]; ok {
			stack := stack.Append(frizz.FieldItem("SliceSub"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []*sub.Sub, null bool, err error) {
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
				var out = make([]*sub.Sub, len(a))
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value *sub.Sub, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// pointerUnpacker
						out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value sub.Sub, null bool, err error) {
							if in == nil {
								return value, true, nil
							}
							// selectorUnpacker
							out, null, err := sub.Packer.UnpackSub(root, stack, in)
							if err != nil {
								return value, false, err
							}
							if null {
								return value, true, nil
							}
							return out, false, nil
						}(root, stack, in)
						if err != nil {
							return value, false, err
						}
						if null {
							return value, true, nil
						}
						return &out, false, nil
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
				out.SliceSub = u
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
	return Pointers(out), false, nil
}
func (p packer) UnpackInterfaceField(root *frizz.Root, stack frizz.Stack, in interface{}) (value InterfaceField, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
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
			return value, false, errors.New("unpacking into struct, value should be a map")
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
			stack := stack.Append(frizz.FieldItem("Iface"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Interface, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// localUnpacker
				out, null, err := p.UnpackInterface(root, stack, in)
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
				out.Iface = u
			}
		}
		if v, ok := m["Slice"]; ok {
			stack := stack.Append(frizz.FieldItem("Slice"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value []Interface, null bool, err error) {
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
				var out = make([]Interface, len(a))
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Interface, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// localUnpacker
						out, null, err := p.UnpackInterface(root, stack, in)
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
				out.Slice = u
			}
		}
		if v, ok := m["Array"]; ok {
			stack := stack.Append(frizz.FieldItem("Array"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [3]Interface, null bool, err error) {
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
				var out [3]Interface
				if len(a) > 3 {
					return value, false, errors.Errorf("data length %d does not fit in array of length %d", len(a), 3)
				}
				for i, v := range a {
					stack := stack.Append(frizz.ArrayItem(i))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Interface, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// localUnpacker
						out, null, err := p.UnpackInterface(root, stack, in)
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
				return out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Array = u
			}
		}
		if v, ok := m["Map"]; ok {
			stack := stack.Append(frizz.FieldItem("Map"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value map[string]Interface, null bool, err error) {
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
				var out = make(map[string]Interface, len(m))
				for k, v := range m {
					stack := stack.Append(frizz.MapItem(k))
					u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value Interface, null bool, err error) {
						if in == nil {
							return value, true, nil
						}
						// localUnpacker
						out, null, err := p.UnpackInterface(root, stack, in)
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
				out.Map = u
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
	return InterfaceField(out), false, nil
}
func (p packer) UnpackImps(root *frizz.Root, stack frizz.Stack, in interface{}) (value Imps, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		String string
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
			String string
		}
		if v, ok := m["String"]; ok {
			stack := stack.Append(frizz.FieldItem("String"))
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
				out.String = u
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
	return Imps(out), false, nil
}
func (p packer) UnpackInterface(root *frizz.Root, stack frizz.Stack, in interface{}) (value Interface, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value interface {
		Foo() string
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// interfaceUnpacker
		out, null, err := root.UnpackInterface(stack, in)
		if err != nil {
			return value, false, err
		}
		iface, ok := out.(interface {
			Foo() string
		})
		if !ok {
			return value, false, errors.Errorf("unpacking into interface, type %T does not implement interface", out)
		}
		if null {
			return value, true, nil
		}
		return iface, false, nil
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return Interface(out), false, nil
}
func (p packer) UnpackAliasArray(root *frizz.Root, stack frizz.Stack, in interface{}) (value AliasArray, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value [3]string, null bool, err error) {
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
		var out [3]string
		if len(a) > 3 {
			return value, false, errors.Errorf("data length %d does not fit in array of length %d", len(a), 3)
		}
		for i, v := range a {
			stack := stack.Append(frizz.ArrayItem(i))
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
				out[i] = u
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
	return AliasArray(out), false, nil
}
func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "CustomSub":
		return p.RepackCustomSub(root, stack, in.(CustomSub))
	case "Impi":
		return p.RepackImpi(root, stack, in.(Impi))
	case "AliasSub":
		return p.RepackAliasSub(root, stack, in.(AliasSub))
	case "Alias":
		return p.RepackAlias(root, stack, in.(Alias))
	case "Int":
		return p.RepackInt(root, stack, in.(Int))
	case "Qual":
		return p.RepackQual(root, stack, in.(Qual))
	case "Custom":
		return p.RepackCustom(root, stack, in.(Custom))
	case "EmbedPointer":
		return p.RepackEmbedPointer(root, stack, in.(EmbedPointer))
	case "SubInterface":
		return p.RepackSubInterface(root, stack, in.(SubInterface))
	case "Pointers":
		return p.RepackPointers(root, stack, in.(Pointers))
	case "InterfaceField":
		return p.RepackInterfaceField(root, stack, in.(InterfaceField))
	case "Imps":
		return p.RepackImps(root, stack, in.(Imps))
	case "Interface":
		return p.RepackInterface(root, stack, in.(Interface))
	case "AliasArray":
		return p.RepackAliasArray(root, stack, in.(AliasArray))
	case "Csv":
		return p.RepackCsv(root, stack, in.(Csv))
	case "EmbedNatives":
		return p.RepackEmbedNatives(root, stack, in.(EmbedNatives))
	case "AliasSlice":
		return p.RepackAliasSlice(root, stack, in.(AliasSlice))
	case "AliasPointer":
		return p.RepackAliasPointer(root, stack, in.(AliasPointer))
	case "Structs":
		return p.RepackStructs(root, stack, in.(Structs))
	case "EmbedQualPointer":
		return p.RepackEmbedQualPointer(root, stack, in.(EmbedQualPointer))
	case "Private":
		return p.RepackPrivate(root, stack, in.(Private))
	case "AliasMap":
		return p.RepackAliasMap(root, stack, in.(AliasMap))
	case "String":
		return p.RepackString(root, stack, in.(String))
	case "Ages":
		return p.RepackAges(root, stack, in.(Ages))
	case "Type":
		return p.RepackType(root, stack, in.(Type))
	case "EmbedQual":
		return p.RepackEmbedQual(root, stack, in.(EmbedQual))
	case "Maps":
		return p.RepackMaps(root, stack, in.(Maps))
	case "Slices":
		return p.RepackSlices(root, stack, in.(Slices))
	case "Natives":
		return p.RepackNatives(root, stack, in.(Natives))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
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
			if in == nil {
				return nil, false, true, nil
			}
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
			return out, dict, false, nil
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
func (p packer) RepackAliasMap(root *frizz.Root, stack frizz.Stack, in AliasMap) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in map[string]*Qual) (value interface{}, dict bool, null bool, err error) {
		// mapRepacker
		out := make(map[string]interface{}, len(in))
		for k, item := range in {
			v, _, _, err := func(root *frizz.Root, stack frizz.Stack, in *Qual) (value interface{}, dict bool, null bool, err error) {
				// pointerRepacker
				if in == nil {
					return nil, false, true, nil
				}
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
				return out, dict, false, nil
			}(root, stack, item)
			if err != nil {
				return nil, false, false, err
			}
			out[k] = v
		}
		return out, true, len(in) == 0, nil
	}(root, stack, (map[string]*Qual)(in))
}
func (p packer) RepackString(root *frizz.Root, stack frizz.Stack, in String) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
		// nativeRepacker
		return frizz.RepackString(in)
	}(root, stack, (string)(in))
}
func (p packer) RepackNatives(root *frizz.Root, stack frizz.Stack, in Natives) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
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
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 20)
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
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *string) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return frizz.RepackString(in)
			}(root, stack, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(root, stack, in.PtrString); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["PtrString"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *int) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return frizz.RepackNumber(in)
			}(root, stack, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(root, stack, in.PtrInt); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["PtrInt"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in *bool) (value interface{}, dict bool, null bool, err error) {
			// pointerRepacker
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in bool) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return frizz.RepackBool(in)
			}(root, stack, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
		}(root, stack, in.PtrBool); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["PtrBool"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
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
	})(in))
}
func (p packer) RepackAges(root *frizz.Root, stack frizz.Stack, in Ages) (value interface{}, dict bool, null bool, err error) {
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
func (p packer) RepackCustomSub(root *frizz.Root, stack frizz.Stack, in CustomSub) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(root, stack)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
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
func (p packer) RepackCustom(root *frizz.Root, stack frizz.Stack, in Custom) (value interface{}, dict bool, null bool, err error) {
	// customRepacker
	out, dict, null, err := in.Repack(root, stack)
	if err != nil {
		return nil, false, false, err
	}
	return out, dict, null, nil
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
			if in == nil {
				return nil, false, true, nil
			}
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
			return out, dict, false, nil
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
func (p packer) RepackSubInterface(root *frizz.Root, stack frizz.Stack, in SubInterface) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		SubInterface sub.SubInterface
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in sub.SubInterface) (value interface{}, dict bool, null bool, err error) {
			// selectorRepacker
			out, dict, null, err := sub.Packer.RepackSubInterface(root, stack, in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}(root, stack, in.SubInterface); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["SubInterface"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		SubInterface sub.SubInterface
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
			if in == nil {
				return nil, false, true, nil
			}
			out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
				// nativeRepacker
				return frizz.RepackString(in)
			}(root, stack, *in)
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, false, nil
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
			if in == nil {
				return nil, false, true, nil
			}
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
			return out, dict, false, nil
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
			if in == nil {
				return nil, false, true, nil
			}
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
			return out, dict, false, nil
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
			if in == nil {
				return nil, false, true, nil
			}
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
			return out, dict, false, nil
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
			if in == nil {
				return nil, false, true, nil
			}
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
			return out, dict, false, nil
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
			if in == nil {
				return nil, false, true, nil
			}
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
			return out, dict, false, nil
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
					if in == nil {
						return nil, false, true, nil
					}
					out, dict, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
						// nativeRepacker
						return frizz.RepackString(in)
					}(root, stack, *in)
					if err != nil {
						return nil, false, false, err
					}
					return out, dict, false, nil
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
					if in == nil {
						return nil, false, true, nil
					}
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
					return out, dict, false, nil
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
					if in == nil {
						return nil, false, true, nil
					}
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
					return out, dict, false, nil
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
func (p packer) RepackCsv(root *frizz.Root, stack frizz.Stack, in Csv) (value interface{}, dict bool, null bool, err error) {
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
func (p packer) RepackAliasPointer(root *frizz.Root, stack frizz.Stack, in AliasPointer) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in *Int) (value interface{}, dict bool, null bool, err error) {
		// pointerRepacker
		if in == nil {
			return nil, false, true, nil
		}
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
		return out, dict, false, nil
	}(root, stack, (*Int)(in))
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

const Imports imports = 0

type imports int

func (i imports) Path() string {
	return "frizz.io/tests/packer"
}
func (i imports) Add(packers map[string]frizz.Packer, types map[string]frizz.Typer) {
	if packers != nil {
		packers["frizz.io/tests/packer"] = Packer
	}
}
