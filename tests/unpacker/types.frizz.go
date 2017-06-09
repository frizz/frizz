package unpacker

import (
	frizz "frizz.io/frizz"
	sub "frizz.io/tests/unpacker/sub"
	errors "github.com/pkg/errors"
)

var Unpackers = struct {
	InterfaceField func(*frizz.Root, frizz.Stack, interface{}) (InterfaceField, error)
	Impi           func(*frizz.Root, frizz.Stack, interface{}) (Impi, error)
	Imps           func(*frizz.Root, frizz.Stack, interface{}) (Imps, error)
	Interface      func(*frizz.Root, frizz.Stack, interface{}) (Interface, error)
	Private        func(*frizz.Root, frizz.Stack, interface{}) (Private, error)
	AliasSub       func(*frizz.Root, frizz.Stack, interface{}) (AliasSub, error)
	AliasSlice     func(*frizz.Root, frizz.Stack, interface{}) (AliasSlice, error)
	AliasArray     func(*frizz.Root, frizz.Stack, interface{}) (AliasArray, error)
	AliasMap       func(*frizz.Root, frizz.Stack, interface{}) (AliasMap, error)
	AliasPointer   func(*frizz.Root, frizz.Stack, interface{}) (AliasPointer, error)
	Alias          func(*frizz.Root, frizz.Stack, interface{}) (Alias, error)
	Int            func(*frizz.Root, frizz.Stack, interface{}) (Int, error)
	String         func(*frizz.Root, frizz.Stack, interface{}) (String, error)
	Qual           func(*frizz.Root, frizz.Stack, interface{}) (Qual, error)
	Pointers       func(*frizz.Root, frizz.Stack, interface{}) (Pointers, error)
	Maps           func(*frizz.Root, frizz.Stack, interface{}) (Maps, error)
	Slices         func(*frizz.Root, frizz.Stack, interface{}) (Slices, error)
	Structs        func(*frizz.Root, frizz.Stack, interface{}) (Structs, error)
	Natives        func(*frizz.Root, frizz.Stack, interface{}) (Natives, error)
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

func unpacker_InterfaceField(r *frizz.Root, s frizz.Stack, in interface{}) (value InterfaceField, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out InterfaceField
	if v, ok := m["Iface"]; ok {
		s := s.Append(frizz.FieldItem("Iface"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value Interface, err error) {
			// localUnpacker
			out, err := unpacker_Interface(r, s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Iface = u
	}
	if v, ok := m["Slice"]; ok {
		s := s.Append(frizz.FieldItem("Slice"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value []Interface, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]Interface, len(a))
			for i, v := range a {
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value Interface, err error) {
					// localUnpacker
					out, err := unpacker_Interface(r, s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Slice = u
	}
	if v, ok := m["Array"]; ok {
		s := s.Append(frizz.FieldItem("Array"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value [3]Interface, err error) {
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
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value Interface, err error) {
					// localUnpacker
					out, err := unpacker_Interface(r, s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Array = u
	}
	if v, ok := m["Map"]; ok {
		s := s.Append(frizz.FieldItem("Map"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value map[string]Interface, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]Interface, len(m))
			for k, v := range m {
				s := s.Append(frizz.MapItem(k))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value Interface, err error) {
					// localUnpacker
					out, err := unpacker_Interface(r, s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Map = u
	}
	return out, nil
}
func unpacker_Impi(r *frizz.Root, s frizz.Stack, in interface{}) (value Impi, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Impi
	if v, ok := m["Int"]; ok {
		s := s.Append(frizz.FieldItem("Int"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	return out, nil
}
func unpacker_Imps(r *frizz.Root, s frizz.Stack, in interface{}) (value Imps, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Imps
	if v, ok := m["String"]; ok {
		s := s.Append(frizz.FieldItem("String"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackString(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.String = u
	}
	return out, nil
}
func unpacker_Interface(r *frizz.Root, s frizz.Stack, in interface{}) (value Interface, err error) {
	// interfaceUnpacker
	out, err := r.UnpackInterface(s, in)
	if err != nil {
		return value, err
	}
	iface, ok := out.(Interface)
	if !ok {
		return value, errors.Errorf("unpacking into interface, type %T does not implement interface", out)
	}
	return iface, nil
}
func unpacker_Private(r *frizz.Root, s frizz.Stack, in interface{}) (value Private, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Private
	if v, ok := m["i"]; ok {
		s := s.Append(frizz.FieldItem("i"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.i = u
	}
	if v, ok := m["s"]; ok {
		s := s.Append(frizz.FieldItem("s"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackString(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.s = u
	}
	return out, nil
}
func unpacker_AliasSub(r *frizz.Root, s frizz.Stack, in interface{}) (value AliasSub, err error) {
	// selectorUnpacker
	out, err := sub.Unpackers.Sub(r, s, in)
	if err != nil {
		return value, err
	}
	return AliasSub(out), nil
}
func unpacker_AliasSlice(r *frizz.Root, s frizz.Stack, in interface{}) (value AliasSlice, err error) {
	// sliceUnpacker
	a, ok := in.([]interface{})
	if !ok {
		return value, errors.New("unpacking into slice, value should be an array")
	}
	var out = make(AliasSlice, len(a))
	for i, v := range a {
		s := s.Append(frizz.ArrayItem(i))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value Int, err error) {
			// localUnpacker
			out, err := unpacker_Int(r, s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out[i] = u
	}
	return out[:], nil
}
func unpacker_AliasArray(r *frizz.Root, s frizz.Stack, in interface{}) (value AliasArray, err error) {
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
		s := s.Append(frizz.ArrayItem(i))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackString(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out[i] = u
	}
	return out, nil
}
func unpacker_AliasMap(r *frizz.Root, s frizz.Stack, in interface{}) (value AliasMap, err error) {
	// mapUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into map, value should be a map")
	}
	var out = make(AliasMap, len(m))
	for k, v := range m {
		s := s.Append(frizz.MapItem(k))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value *Qual, err error) {
			// pointerUnpacker
			out, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value Qual, err error) {
				// localUnpacker
				out, err := unpacker_Qual(r, s, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(r, s, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out[k] = u
	}
	return out, nil
}
func unpacker_AliasPointer(r *frizz.Root, s frizz.Stack, in interface{}) (value AliasPointer, err error) {
	// pointerUnpacker
	out, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value Int, err error) {
		// localUnpacker
		out, err := unpacker_Int(r, s, in)
		if err != nil {
			return value, err
		}
		return out, nil
	}(r, s, in)
	if err != nil {
		return value, err
	}
	return AliasPointer(&out), nil
}
func unpacker_Alias(r *frizz.Root, s frizz.Stack, in interface{}) (value Alias, err error) {
	// localUnpacker
	out, err := unpacker_Int(r, s, in)
	if err != nil {
		return value, err
	}
	return Alias(out), nil
}
func unpacker_Int(r *frizz.Root, s frizz.Stack, in interface{}) (value Int, err error) {
	// nativeUnpacker
	out, err := frizz.UnpackInt(s, in)
	if err != nil {
		return value, err
	}
	return Int(out), nil
}
func unpacker_String(r *frizz.Root, s frizz.Stack, in interface{}) (value String, err error) {
	// nativeUnpacker
	out, err := frizz.UnpackString(s, in)
	if err != nil {
		return value, err
	}
	return String(out), nil
}
func unpacker_Qual(r *frizz.Root, s frizz.Stack, in interface{}) (value Qual, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Qual
	if v, ok := m["Sub"]; ok {
		s := s.Append(frizz.FieldItem("Sub"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value sub.Sub, err error) {
			// selectorUnpacker
			out, err := sub.Unpackers.Sub(r, s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Sub = u
	}
	return out, nil
}
func unpacker_Pointers(r *frizz.Root, s frizz.Stack, in interface{}) (value Pointers, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Pointers
	if v, ok := m["String"]; ok {
		s := s.Append(frizz.FieldItem("String"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value *string, err error) {
			// pointerUnpacker
			out, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
				// nativeUnpacker
				out, err := frizz.UnpackString(s, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(r, s, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.String = u
	}
	if v, ok := m["Int"]; ok {
		s := s.Append(frizz.FieldItem("Int"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value *Int, err error) {
			// pointerUnpacker
			out, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value Int, err error) {
				// localUnpacker
				out, err := unpacker_Int(r, s, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(r, s, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	if v, ok := m["Sub"]; ok {
		s := s.Append(frizz.FieldItem("Sub"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value *sub.Sub, err error) {
			// pointerUnpacker
			out, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value sub.Sub, err error) {
				// selectorUnpacker
				out, err := sub.Unpackers.Sub(r, s, in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(r, s, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Sub = u
	}
	if v, ok := m["Array"]; ok {
		s := s.Append(frizz.FieldItem("Array"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value *[3]int, err error) {
			// pointerUnpacker
			out, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value [3]int, err error) {
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
					s := s.Append(frizz.ArrayItem(i))
					u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int, err error) {
						// nativeUnpacker
						out, err := frizz.UnpackInt(s, in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(r, s, v)
					if err != nil {
						return value, err
					}
					out[i] = u
				}
				return out, nil
			}(r, s, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Array = u
	}
	if v, ok := m["Slice"]; ok {
		s := s.Append(frizz.FieldItem("Slice"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value *[]string, err error) {
			// pointerUnpacker
			out, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value []string, err error) {
				// sliceUnpacker
				a, ok := in.([]interface{})
				if !ok {
					return value, errors.New("unpacking into slice, value should be an array")
				}
				var out = make([]string, len(a))
				for i, v := range a {
					s := s.Append(frizz.ArrayItem(i))
					u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
						// nativeUnpacker
						out, err := frizz.UnpackString(s, in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(r, s, v)
					if err != nil {
						return value, err
					}
					out[i] = u
				}
				return out[:], nil
			}(r, s, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Slice = u
	}
	if v, ok := m["Map"]; ok {
		s := s.Append(frizz.FieldItem("Map"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value *map[string]int, err error) {
			// pointerUnpacker
			out, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value map[string]int, err error) {
				// mapUnpacker
				m, ok := in.(map[string]interface{})
				if !ok {
					return value, errors.New("unpacking into map, value should be a map")
				}
				var out = make(map[string]int, len(m))
				for k, v := range m {
					s := s.Append(frizz.MapItem(k))
					u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int, err error) {
						// nativeUnpacker
						out, err := frizz.UnpackInt(s, in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(r, s, v)
					if err != nil {
						return value, err
					}
					out[k] = u
				}
				return out, nil
			}(r, s, in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Map = u
	}
	if v, ok := m["SliceString"]; ok {
		s := s.Append(frizz.FieldItem("SliceString"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value []*string, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]*string, len(a))
			for i, v := range a {
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value *string, err error) {
					// pointerUnpacker
					out, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
						// nativeUnpacker
						out, err := frizz.UnpackString(s, in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(r, s, in)
					if err != nil {
						return value, err
					}
					return &out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.SliceString = u
	}
	if v, ok := m["SliceInt"]; ok {
		s := s.Append(frizz.FieldItem("SliceInt"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value []*Int, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]*Int, len(a))
			for i, v := range a {
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value *Int, err error) {
					// pointerUnpacker
					out, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value Int, err error) {
						// localUnpacker
						out, err := unpacker_Int(r, s, in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(r, s, in)
					if err != nil {
						return value, err
					}
					return &out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.SliceInt = u
	}
	if v, ok := m["SliceSub"]; ok {
		s := s.Append(frizz.FieldItem("SliceSub"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value []*sub.Sub, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]*sub.Sub, len(a))
			for i, v := range a {
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value *sub.Sub, err error) {
					// pointerUnpacker
					out, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value sub.Sub, err error) {
						// selectorUnpacker
						out, err := sub.Unpackers.Sub(r, s, in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(r, s, in)
					if err != nil {
						return value, err
					}
					return &out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.SliceSub = u
	}
	return out, nil
}
func unpacker_Maps(r *frizz.Root, s frizz.Stack, in interface{}) (value Maps, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Maps
	if v, ok := m["Ints"]; ok {
		s := s.Append(frizz.FieldItem("Ints"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value map[string]int, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]int, len(m))
			for k, v := range m {
				s := s.Append(frizz.MapItem(k))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackInt(s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Ints = u
	}
	if v, ok := m["Strings"]; ok {
		s := s.Append(frizz.FieldItem("Strings"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value map[string]string, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]string, len(m))
			for k, v := range m {
				s := s.Append(frizz.MapItem(k))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackString(s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Strings = u
	}
	if v, ok := m["Slices"]; ok {
		s := s.Append(frizz.FieldItem("Slices"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value map[string][]string, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string][]string, len(m))
			for k, v := range m {
				s := s.Append(frizz.MapItem(k))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value []string, err error) {
					// sliceUnpacker
					a, ok := in.([]interface{})
					if !ok {
						return value, errors.New("unpacking into slice, value should be an array")
					}
					var out = make([]string, len(a))
					for i, v := range a {
						s := s.Append(frizz.ArrayItem(i))
						u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
							// nativeUnpacker
							out, err := frizz.UnpackString(s, in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(r, s, v)
						if err != nil {
							return value, err
						}
						out[i] = u
					}
					return out[:], nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Slices = u
	}
	if v, ok := m["Arrays"]; ok {
		s := s.Append(frizz.FieldItem("Arrays"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value map[string][2]int, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string][2]int, len(m))
			for k, v := range m {
				s := s.Append(frizz.MapItem(k))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value [2]int, err error) {
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
						s := s.Append(frizz.ArrayItem(i))
						u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int, err error) {
							// nativeUnpacker
							out, err := frizz.UnpackInt(s, in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(r, s, v)
						if err != nil {
							return value, err
						}
						out[i] = u
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Arrays = u
	}
	if v, ok := m["Maps"]; ok {
		s := s.Append(frizz.FieldItem("Maps"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value map[string]map[string]string, err error) {
			// mapUnpacker
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("unpacking into map, value should be a map")
			}
			var out = make(map[string]map[string]string, len(m))
			for k, v := range m {
				s := s.Append(frizz.MapItem(k))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value map[string]string, err error) {
					// mapUnpacker
					m, ok := in.(map[string]interface{})
					if !ok {
						return value, errors.New("unpacking into map, value should be a map")
					}
					var out = make(map[string]string, len(m))
					for k, v := range m {
						s := s.Append(frizz.MapItem(k))
						u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
							// nativeUnpacker
							out, err := frizz.UnpackString(s, in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(r, s, v)
						if err != nil {
							return value, err
						}
						out[k] = u
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Maps = u
	}
	return out, nil
}
func unpacker_Slices(r *frizz.Root, s frizz.Stack, in interface{}) (value Slices, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Slices
	if v, ok := m["Ints"]; ok {
		s := s.Append(frizz.FieldItem("Ints"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value []int, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]int, len(a))
			for i, v := range a {
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackInt(s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Ints = u
	}
	if v, ok := m["Strings"]; ok {
		s := s.Append(frizz.FieldItem("Strings"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value []string, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([]string, len(a))
			for i, v := range a {
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackString(s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Strings = u
	}
	if v, ok := m["ArrayLit"]; ok {
		s := s.Append(frizz.FieldItem("ArrayLit"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value [5]string, err error) {
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
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackString(s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.ArrayLit = u
	}
	if v, ok := m["ArrayExpr"]; ok {
		s := s.Append(frizz.FieldItem("ArrayExpr"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value [2 * N]int, err error) {
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
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackInt(s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.ArrayExpr = u
	}
	if v, ok := m["Structs"]; ok {
		s := s.Append(frizz.FieldItem("Structs"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value []struct {
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
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value struct {
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
						s := s.Append(frizz.FieldItem("Int"))
						u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int, err error) {
							// nativeUnpacker
							out, err := frizz.UnpackInt(s, in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(r, s, v)
						if err != nil {
							return value, err
						}
						out.Int = u
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Structs = u
	}
	if v, ok := m["Arrays"]; ok {
		s := s.Append(frizz.FieldItem("Arrays"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value [][]string, err error) {
			// sliceUnpacker
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("unpacking into slice, value should be an array")
			}
			var out = make([][]string, len(a))
			for i, v := range a {
				s := s.Append(frizz.ArrayItem(i))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value []string, err error) {
					// sliceUnpacker
					a, ok := in.([]interface{})
					if !ok {
						return value, errors.New("unpacking into slice, value should be an array")
					}
					var out = make([]string, len(a))
					for i, v := range a {
						s := s.Append(frizz.ArrayItem(i))
						u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
							// nativeUnpacker
							out, err := frizz.UnpackString(s, in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(r, s, v)
						if err != nil {
							return value, err
						}
						out[i] = u
					}
					return out[:], nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Arrays = u
	}
	return out, nil
}
func unpacker_Structs(r *frizz.Root, s frizz.Stack, in interface{}) (value Structs, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Structs
	if v, ok := m["Simple"]; ok {
		s := s.Append(frizz.FieldItem("Simple"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value struct {
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
				s := s.Append(frizz.FieldItem("Int"))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackInt(s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out.Int = u
			}
			if v, ok := m["Bool"]; ok {
				s := s.Append(frizz.FieldItem("Bool"))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value bool, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackBool(s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out.Bool = u
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Simple = u
	}
	if v, ok := m["Complex"]; ok {
		s := s.Append(frizz.FieldItem("Complex"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value struct {
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
				s := s.Append(frizz.FieldItem("String"))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
					// nativeUnpacker
					out, err := frizz.UnpackString(s, in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out.String = u
			}
			if v, ok := m["Inner"]; ok {
				s := s.Append(frizz.FieldItem("Inner"))
				u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value struct {
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
						s := s.Append(frizz.FieldItem("Float32"))
						u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value float32, err error) {
							// nativeUnpacker
							out, err := frizz.UnpackFloat32(s, in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(r, s, v)
						if err != nil {
							return value, err
						}
						out.Float32 = u
					}
					return out, nil
				}(r, s, v)
				if err != nil {
					return value, err
				}
				out.Inner = u
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Complex = u
	}
	return out, nil
}
func unpacker_Natives(r *frizz.Root, s frizz.Stack, in interface{}) (value Natives, err error) {
	// structUnpacker
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("unpacking into struct, value should be a map")
	}
	var out Natives
	if v, ok := m["Bool"]; ok {
		s := s.Append(frizz.FieldItem("Bool"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value bool, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackBool(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Bool = u
	}
	if v, ok := m["Byte"]; ok {
		s := s.Append(frizz.FieldItem("Byte"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value byte, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackByte(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Byte = u
	}
	if v, ok := m["Float32"]; ok {
		s := s.Append(frizz.FieldItem("Float32"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value float32, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackFloat32(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Float32 = u
	}
	if v, ok := m["Float64"]; ok {
		s := s.Append(frizz.FieldItem("Float64"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value float64, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackFloat64(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Float64 = u
	}
	if v, ok := m["Int"]; ok {
		s := s.Append(frizz.FieldItem("Int"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	if v, ok := m["Int8"]; ok {
		s := s.Append(frizz.FieldItem("Int8"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int8, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt8(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Int8 = u
	}
	if v, ok := m["Int16"]; ok {
		s := s.Append(frizz.FieldItem("Int16"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int16, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt16(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Int16 = u
	}
	if v, ok := m["Int32"]; ok {
		s := s.Append(frizz.FieldItem("Int32"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int32, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt32(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Int32 = u
	}
	if v, ok := m["Int64"]; ok {
		s := s.Append(frizz.FieldItem("Int64"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value int64, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackInt64(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Int64 = u
	}
	if v, ok := m["Uint"]; ok {
		s := s.Append(frizz.FieldItem("Uint"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value uint, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackUint(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Uint = u
	}
	if v, ok := m["Uint8"]; ok {
		s := s.Append(frizz.FieldItem("Uint8"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value uint8, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackUint8(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Uint8 = u
	}
	if v, ok := m["Uint16"]; ok {
		s := s.Append(frizz.FieldItem("Uint16"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value uint16, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackUint16(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Uint16 = u
	}
	if v, ok := m["Uint32"]; ok {
		s := s.Append(frizz.FieldItem("Uint32"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value uint32, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackUint32(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Uint32 = u
	}
	if v, ok := m["Uint64"]; ok {
		s := s.Append(frizz.FieldItem("Uint64"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value uint64, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackUint64(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Uint64 = u
	}
	if v, ok := m["Rune"]; ok {
		s := s.Append(frizz.FieldItem("Rune"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value rune, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackRune(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.Rune = u
	}
	if v, ok := m["String"]; ok {
		s := s.Append(frizz.FieldItem("String"))
		u, err := func(r *frizz.Root, s frizz.Stack, in interface{}) (value string, err error) {
			// nativeUnpacker
			out, err := frizz.UnpackString(s, in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(r, s, v)
		if err != nil {
			return value, err
		}
		out.String = u
	}
	return out, nil
}
func init() {
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "InterfaceField", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_InterfaceField(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Impi", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Impi(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Imps", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Imps(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Interface", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Interface(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Private", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Private(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "AliasSub", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_AliasSub(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "AliasSlice", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_AliasSlice(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "AliasArray", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_AliasArray(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "AliasMap", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_AliasMap(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "AliasPointer", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_AliasPointer(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Alias", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Alias(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Int", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Int(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "String", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_String(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Qual", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Qual(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Pointers", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Pointers(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Maps", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Maps(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Slices", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Slices(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Structs", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Structs(r, s, in)
	})
	frizz.DefaultRegistry.Set("frizz.io/tests/unpacker", "Natives", func(r *frizz.Root, s frizz.Stack, in interface{}) (interface{}, error) {
		return unpacker_Natives(r, s, in)
	})
}
