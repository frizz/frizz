package unpacker

import (
	context "context"
	system "frizz.io/system"
	sub "frizz.io/tests/unpacker/sub"
	errors "github.com/pkg/errors"
)

var Unpackers = struct {
	InterfaceField func(context.Context, interface{}) (InterfaceField, error)
	Impi           func(context.Context, interface{}) (Impi, error)
	Imps           func(context.Context, interface{}) (Imps, error)
	Interface      func(context.Context, interface{}) (Interface, error)
	Private        func(context.Context, interface{}) (Private, error)
	AliasSub       func(context.Context, interface{}) (AliasSub, error)
	AliasSlice     func(context.Context, interface{}) (AliasSlice, error)
	AliasArray     func(context.Context, interface{}) (AliasArray, error)
	AliasMap       func(context.Context, interface{}) (AliasMap, error)
	AliasPointer   func(context.Context, interface{}) (AliasPointer, error)
	Alias          func(context.Context, interface{}) (Alias, error)
	Int            func(context.Context, interface{}) (Int, error)
	String         func(context.Context, interface{}) (String, error)
	Qual           func(context.Context, interface{}) (Qual, error)
	Pointers       func(context.Context, interface{}) (Pointers, error)
	Maps           func(context.Context, interface{}) (Maps, error)
	Slices         func(context.Context, interface{}) (Slices, error)
	Structs        func(context.Context, interface{}) (Structs, error)
	Natives        func(context.Context, interface{}) (Natives, error)
}{
	Alias:          unpacker_Alias,
	AliasArray:     unpacker_AliasArray,
	AliasMap:       unpacker_AliasMap,
	AliasPointer:   unpacker_AliasPointer,
	AliasSlice:     unpacker_AliasSlice,
	AliasSub:       unpacker_AliasSub,
	Impi:           unpacker_Impi,
	Imps:           unpacker_Imps,
	Int:            unpacker_Int,
	Interface:      unpacker_Interface,
	InterfaceField: unpacker_InterfaceField,
	Maps:           unpacker_Maps,
	Natives:        unpacker_Natives,
	Pointers:       unpacker_Pointers,
	Private:        unpacker_Private,
	Qual:           unpacker_Qual,
	Slices:         unpacker_Slices,
	String:         unpacker_String,
	Structs:        unpacker_Structs,
}

func unpacker_InterfaceField(ctx context.Context, in interface{}) (value InterfaceField, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out InterfaceField
	if v, ok := m["Iface"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value Interface, err error) {
			// localUnpacker
			out, err := unpacker_Interface(ctx, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Iface = u
	}
	if v, ok := m["Slice"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value []Interface, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]Interface, len(a))
			for i, v := range a {
				u, err := func(ctx context.Context, in interface{}) (value Interface, err error) {
					// localUnpacker
					out, err := unpacker_Interface(ctx, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Slice = u
	}
	if v, ok := m["Array"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value [3]Interface, err error) {
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
				u, err := func(ctx context.Context, in interface{}) (value Interface, err error) {
					// localUnpacker
					out, err := unpacker_Interface(ctx, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Array = u
	}
	if v, ok := m["Map"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value map[string]Interface, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]Interface, len(m))
			for k, v := range m {
				u, err := func(ctx context.Context, in interface{}) (value Interface, err error) {
					// localUnpacker
					out, err := unpacker_Interface(ctx, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Map = u
	}
	return out, nil
}
func unpacker_Impi(ctx context.Context, in interface{}) (value Impi, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Impi
	if v, ok := m["Int"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := system.Convert_int(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	return out, nil
}
func unpacker_Imps(ctx context.Context, in interface{}) (value Imps, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Imps
	if v, ok := m["String"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := system.Convert_string(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.String = u
	}
	return out, nil
}
func unpacker_Interface(ctx context.Context, in interface{}) (value Interface, err error) {
	// interfaceUnpacker
	out, err := system.UnpackInterface(ctx, in)
	if err != nil {
		return value, err
	}
	iface, ok := out.(Interface)
	if !ok {
		return value, errors.Errorf("unpacking into interface, type %T does not implement interface", out)
	}
	return iface, nil
}
func unpacker_Private(ctx context.Context, in interface{}) (value Private, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Private
	if v, ok := m["i"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := system.Convert_int(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.i = u
	}
	if v, ok := m["s"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := system.Convert_string(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.s = u
	}
	return out, nil
}
func unpacker_AliasSub(ctx context.Context, in interface{}) (value AliasSub, err error) {
	// selectorUnpacker
	out, err := sub.Unpackers.Sub(ctx, in)
	if err != nil {
		return value, err
	}
	return AliasSub(out), nil
}
func unpacker_AliasSlice(ctx context.Context, in interface{}) (value AliasSlice, err error) {
	// sliceUnpacker
	a, ok := in.([]interface{})
	if !ok {
		return value, errors.New("unpacking into slice, value should be an array")
	}
	var out = make(AliasSlice, len(a))
	for i, v := range a {
		u, err := func(ctx context.Context, in interface{}) (value Int, err error) {
			// localUnpacker
			out, err := unpacker_Int(ctx, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out[i] = u
	}
	return out[:], nil
}
func unpacker_AliasArray(ctx context.Context, in interface{}) (value AliasArray, err error) {
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
		u, err := func(ctx context.Context, in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := system.Convert_string(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out[i] = u
	}
	return out, nil
}
func unpacker_AliasMap(ctx context.Context, in interface{}) (value AliasMap, err error) {
	// mapUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into map, value should be a map")
	}
	var out = make(AliasMap, len(m))
	for k, v := range m {
		u, err := func(ctx context.Context, in interface{}) (value *Qual, err error) {
			// pointerUnpacker
			out, err := func(ctx context.Context, in interface{}) (value Qual, err error) {
				// localUnpacker
				out, err := unpacker_Qual(ctx, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(ctx, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out[k] = u
	}
	return out, nil
}
func unpacker_AliasPointer(ctx context.Context, in interface{}) (value AliasPointer, err error) {
	// pointerUnpacker
	out, err := func(ctx context.Context, in interface{}) (value Int, err error) {
		// localUnpacker
		out, err := unpacker_Int(ctx, in)
		if err != nil {
			return value, err
		}
		return out, nil
	}(ctx, in)
	if err != nil {
		return value, err
	}
	return AliasPointer(&out), nil
}
func unpacker_Alias(ctx context.Context, in interface{}) (value Alias, err error) {
	// localUnpacker
	out, err := unpacker_Int(ctx, in)
	if err != nil {
		return value, err
	}
	return Alias(out), nil
}
func unpacker_Int(ctx context.Context, in interface{}) (value Int, err error) {
	// nativeUnpacker
	out, err := system.Convert_int(in)
	if err != nil {
		return value, err
	}
	return Int(out), nil
}
func unpacker_String(ctx context.Context, in interface{}) (value String, err error) {
	// nativeUnpacker
	out, err := system.Convert_string(in)
	if err != nil {
		return value, err
	}
	return String(out), nil
}
func unpacker_Qual(ctx context.Context, in interface{}) (value Qual, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Qual
	if v, ok := m["Sub"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value sub.Sub, err error) {
			// selectorUnpacker
			out, err := sub.Unpackers.Sub(ctx, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Sub = u
	}
	return out, nil
}
func unpacker_Pointers(ctx context.Context, in interface{}) (value Pointers, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Pointers
	if v, ok := m["String"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value *string, err error) {
			// pointerUnpacker
			out, err := func(ctx context.Context, in interface{}) (value string, err error) {
				// nativeUnpacker
				out, err := system.Convert_string(in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(ctx, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.String = u
	}
	if v, ok := m["Int"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value *Int, err error) {
			// pointerUnpacker
			out, err := func(ctx context.Context, in interface{}) (value Int, err error) {
				// localUnpacker
				out, err := unpacker_Int(ctx, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(ctx, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	if v, ok := m["Sub"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value *sub.Sub, err error) {
			// pointerUnpacker
			out, err := func(ctx context.Context, in interface{}) (value sub.Sub, err error) {
				// selectorUnpacker
				out, err := sub.Unpackers.Sub(ctx, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(ctx, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Sub = u
	}
	if v, ok := m["Array"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value *[3]int, err error) {
			// pointerUnpacker
			out, err := func(ctx context.Context, in interface{}) (value [3]int, err error) {
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
					u, err := func(ctx context.Context, in interface{}) (value int, err error) {
						// nativeUnpacker
						out, err := system.Convert_int(in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(ctx, v)
					if err != nil {
						return value, err
					}
					out[i] = u
				}
				return out, nil
			}(ctx, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Array = u
	}
	if v, ok := m["Slice"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value *[]string, err error) {
			// pointerUnpacker
			out, err := func(ctx context.Context, in interface{}) (value []string, err error) {
				// sliceUnpacker
				a, ok := in.([]interface{})
				if !ok {
					return value, errors.New("unpacking into slice, value should be an array")
				}
				var out = make([]string, len(a))
				for i, v := range a {
					u, err := func(ctx context.Context, in interface{}) (value string, err error) {
						// nativeUnpacker
						out, err := system.Convert_string(in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(ctx, v)
					if err != nil {
						return value, err
					}
					out[i] = u
				}
				return out[:], nil
			}(ctx, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Slice = u
	}
	if v, ok := m["Map"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value *map[string]int, err error) {
			// pointerUnpacker
			out, err := func(ctx context.Context, in interface{}) (value map[string]int, err error) {
				// mapUnpacker
				m, ok := in.(map[string]interface{})
				if !ok {
					return value, errors.New("unpacking into map, value should be a map")
				}
				var out = make(map[string]int, len(m))
				for k, v := range m {
					u, err := func(ctx context.Context, in interface{}) (value int, err error) {
						// nativeUnpacker
						out, err := system.Convert_int(in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(ctx, v)
					if err != nil {
						return value, err
					}
					out[k] = u
				}
				return out, nil
			}(ctx, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Map = u
	}
	if v, ok := m["SliceString"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value []*string, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]*string, len(a))
			for i, v := range a {
				u, err := func(ctx context.Context, in interface{}) (value *string, err error) {
					// pointerUnpacker
					out, err := func(ctx context.Context, in interface{}) (value string, err error) {
						// nativeUnpacker
						out, err := system.Convert_string(in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(ctx, in)
					if err != nil {
						return value, err
					}
					return &out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.SliceString = u
	}
	if v, ok := m["SliceInt"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value []*Int, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]*Int, len(a))
			for i, v := range a {
				u, err := func(ctx context.Context, in interface{}) (value *Int, err error) {
					// pointerUnpacker
					out, err := func(ctx context.Context, in interface{}) (value Int, err error) {
						// localUnpacker
						out, err := unpacker_Int(ctx, in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(ctx, in)
					if err != nil {
						return value, err
					}
					return &out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.SliceInt = u
	}
	if v, ok := m["SliceSub"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value []*sub.Sub, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]*sub.Sub, len(a))
			for i, v := range a {
				u, err := func(ctx context.Context, in interface{}) (value *sub.Sub, err error) {
					// pointerUnpacker
					out, err := func(ctx context.Context, in interface{}) (value sub.Sub, err error) {
						// selectorUnpacker
						out, err := sub.Unpackers.Sub(ctx, in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(ctx, in)
					if err != nil {
						return value, err
					}
					return &out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.SliceSub = u
	}
	return out, nil
}
func unpacker_Maps(ctx context.Context, in interface{}) (value Maps, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Maps
	if v, ok := m["Ints"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value map[string]int, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]int, len(m))
			for k, v := range m {
				u, err := func(ctx context.Context, in interface{}) (value int, err error) {
					// nativeUnpacker
					out, err := system.Convert_int(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Ints = u
	}
	if v, ok := m["Strings"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value map[string]string, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]string, len(m))
			for k, v := range m {
				u, err := func(ctx context.Context, in interface{}) (value string, err error) {
					// nativeUnpacker
					out, err := system.Convert_string(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Strings = u
	}
	if v, ok := m["Slices"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value map[string][]string, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string][]string, len(m))
			for k, v := range m {
				u, err := func(ctx context.Context, in interface{}) (value []string, err error) {
					// sliceUnpacker
					a, ok := in.([]interface{})
					if !ok {
						return value, errors.New("unpacking into slice, value should be an array")
					}
					var out = make([]string, len(a))
					for i, v := range a {
						u, err := func(ctx context.Context, in interface{}) (value string, err error) {
							// nativeUnpacker
							out, err := system.Convert_string(in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(ctx, v)
						if err != nil {
							return value, err
						}
						out[i] = u
					}
					return out[:], nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Slices = u
	}
	if v, ok := m["Arrays"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value map[string][2]int, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string][2]int, len(m))
			for k, v := range m {
				u, err := func(ctx context.Context, in interface{}) (value [2]int, err error) {
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
						u, err := func(ctx context.Context, in interface{}) (value int, err error) {
							// nativeUnpacker
							out, err := system.Convert_int(in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(ctx, v)
						if err != nil {
							return value, err
						}
						out[i] = u
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Arrays = u
	}
	if v, ok := m["Maps"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value map[string]map[string]string, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]map[string]string, len(m))
			for k, v := range m {
				u, err := func(ctx context.Context, in interface{}) (value map[string]string, err error) {
					// mapUnpacker
					m, ok := in.(map[string]interface{})
					if !ok {
						return value, errors.New("unpacking into map, value should be a map")
					}
					var out = make(map[string]string, len(m))
					for k, v := range m {
						u, err := func(ctx context.Context, in interface{}) (value string, err error) {
							// nativeUnpacker
							out, err := system.Convert_string(in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(ctx, v)
						if err != nil {
							return value, err
						}
						out[k] = u
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Maps = u
	}
	return out, nil
}
func unpacker_Slices(ctx context.Context, in interface{}) (value Slices, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Slices
	if v, ok := m["Ints"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value []int, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]int, len(a))
			for i, v := range a {
				u, err := func(ctx context.Context, in interface{}) (value int, err error) {
					// nativeUnpacker
					out, err := system.Convert_int(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Ints = u
	}
	if v, ok := m["Strings"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value []string, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]string, len(a))
			for i, v := range a {
				u, err := func(ctx context.Context, in interface{}) (value string, err error) {
					// nativeUnpacker
					out, err := system.Convert_string(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Strings = u
	}
	if v, ok := m["ArrayLit"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value [5]string, err error) {
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
				u, err := func(ctx context.Context, in interface{}) (value string, err error) {
					// nativeUnpacker
					out, err := system.Convert_string(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.ArrayLit = u
	}
	if v, ok := m["ArrayExpr"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value [2 * N]int, err error) {
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
				u, err := func(ctx context.Context, in interface{}) (value int, err error) {
					// nativeUnpacker
					out, err := system.Convert_int(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.ArrayExpr = u
	}
	if v, ok := m["Structs"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value []struct {
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
				u, err := func(ctx context.Context, in interface{}) (value struct {
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
						u, err := func(ctx context.Context, in interface{}) (value int, err error) {
							// nativeUnpacker
							out, err := system.Convert_int(in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(ctx, v)
						if err != nil {
							return value, err
						}
						out.Int = u
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Structs = u
	}
	if v, ok := m["Arrays"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value [][]string, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([][]string, len(a))
			for i, v := range a {
				u, err := func(ctx context.Context, in interface{}) (value []string, err error) {
					// sliceUnpacker
					a, ok := in.([]interface{})
					if !ok {
						return value, errors.New("unpacking into slice, value should be an array")
					}
					var out = make([]string, len(a))
					for i, v := range a {
						u, err := func(ctx context.Context, in interface{}) (value string, err error) {
							// nativeUnpacker
							out, err := system.Convert_string(in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(ctx, v)
						if err != nil {
							return value, err
						}
						out[i] = u
					}
					return out[:], nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Arrays = u
	}
	return out, nil
}
func unpacker_Structs(ctx context.Context, in interface{}) (value Structs, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Structs
	if v, ok := m["Simple"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value struct {
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
				u, err := func(ctx context.Context, in interface{}) (value int, err error) {
					// nativeUnpacker
					out, err := system.Convert_int(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out.Int = u
			}
			if v, ok := m["Bool"]; ok {
				u, err := func(ctx context.Context, in interface{}) (value bool, err error) {
					// nativeUnpacker
					out, err := system.Convert_bool(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out.Bool = u
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Simple = u
	}
	if v, ok := m["Complex"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value struct {
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
				u, err := func(ctx context.Context, in interface{}) (value string, err error) {
					// nativeUnpacker
					out, err := system.Convert_string(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out.String = u
			}
			if v, ok := m["Inner"]; ok {
				u, err := func(ctx context.Context, in interface{}) (value struct {
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
						u, err := func(ctx context.Context, in interface{}) (value float32, err error) {
							// nativeUnpacker
							out, err := system.Convert_float32(in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(ctx, v)
						if err != nil {
							return value, err
						}
						out.Float32 = u
					}
					return out, nil
				}(ctx, v)
				if err != nil {
					return value, err
				}
				out.Inner = u
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Complex = u
	}
	return out, nil
}
func unpacker_Natives(ctx context.Context, in interface{}) (value Natives, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Natives
	if v, ok := m["Bool"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value bool, err error) {
			// nativeUnpacker
			out, err := system.Convert_bool(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Bool = u
	}
	if v, ok := m["Byte"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value byte, err error) {
			// nativeUnpacker
			out, err := system.Convert_byte(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Byte = u
	}
	if v, ok := m["Float32"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value float32, err error) {
			// nativeUnpacker
			out, err := system.Convert_float32(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Float32 = u
	}
	if v, ok := m["Float64"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value float64, err error) {
			// nativeUnpacker
			out, err := system.Convert_float64(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Float64 = u
	}
	if v, ok := m["Int"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := system.Convert_int(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	if v, ok := m["Int8"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value int8, err error) {
			// nativeUnpacker
			out, err := system.Convert_int8(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Int8 = u
	}
	if v, ok := m["Int16"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value int16, err error) {
			// nativeUnpacker
			out, err := system.Convert_int16(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Int16 = u
	}
	if v, ok := m["Int32"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value int32, err error) {
			// nativeUnpacker
			out, err := system.Convert_int32(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Int32 = u
	}
	if v, ok := m["Int64"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value int64, err error) {
			// nativeUnpacker
			out, err := system.Convert_int64(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Int64 = u
	}
	if v, ok := m["Uint"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value uint, err error) {
			// nativeUnpacker
			out, err := system.Convert_uint(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Uint = u
	}
	if v, ok := m["Uint8"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value uint8, err error) {
			// nativeUnpacker
			out, err := system.Convert_uint8(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Uint8 = u
	}
	if v, ok := m["Uint16"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value uint16, err error) {
			// nativeUnpacker
			out, err := system.Convert_uint16(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Uint16 = u
	}
	if v, ok := m["Uint32"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value uint32, err error) {
			// nativeUnpacker
			out, err := system.Convert_uint32(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Uint32 = u
	}
	if v, ok := m["Uint64"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value uint64, err error) {
			// nativeUnpacker
			out, err := system.Convert_uint64(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Uint64 = u
	}
	if v, ok := m["Rune"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value rune, err error) {
			// nativeUnpacker
			out, err := system.Convert_rune(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.Rune = u
	}
	if v, ok := m["String"]; ok {
		u, err := func(ctx context.Context, in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := system.Convert_string(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(ctx, v)
		if err != nil {
			return value, err
		}
		out.String = u
	}
	return out, nil
}
func init() {
	system.Registry.Set(system.RegistryTypeKey{
		Name: "InterfaceField",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_InterfaceField(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "Impi",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_Impi(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "Imps",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_Imps(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "Interface",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_Interface(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "Private",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_Private(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "AliasSub",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_AliasSub(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "AliasSlice",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_AliasSlice(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "AliasArray",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_AliasArray(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "AliasMap",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_AliasMap(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "AliasPointer",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_AliasPointer(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "Alias",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_Alias(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "Int",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_Int(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "String",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_String(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "Qual",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_Qual(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "Pointers",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_Pointers(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "Maps",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_Maps(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "Slices",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_Slices(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "Structs",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_Structs(ctx, in)
	}})
	system.Registry.Set(system.RegistryTypeKey{
		Name: "Natives",
		Path: "frizz.io/tests/unpacker",
	}, system.RegistryType{Unpacker: func(ctx context.Context, in interface{}) (interface{}, error) {
		return unpacker_Natives(ctx, in)
	}})
}
