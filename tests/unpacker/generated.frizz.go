package unpacker

import (
	system "frizz.io/system"
	sub "frizz.io/tests/unpacker/sub"
	errors "github.com/pkg/errors"
)

var Unpackers = struct {
	Int      func(interface{}) (Int, error)
	String   func(interface{}) (String, error)
	Qual     func(interface{}) (Qual, error)
	Pointers func(interface{}) (Pointers, error)
	Maps     func(interface{}) (Maps, error)
	Slices   func(interface{}) (Slices, error)
	Structs  func(interface{}) (Structs, error)
	Natives  func(interface{}) (Natives, error)
}{
	Int:      unpacker_Int,
	Maps:     unpacker_Maps,
	Natives:  unpacker_Natives,
	Pointers: unpacker_Pointers,
	Qual:     unpacker_Qual,
	Slices:   unpacker_Slices,
	String:   unpacker_String,
	Structs:  unpacker_Structs,
}

func unpacker_Int(in interface{}) (value Int, err error) {
	out, err := system.Convert_int(in)
	if err != nil {
		return value, err
	}
	return Int(out), nil
}
func unpacker_String(in interface{}) (value String, err error) {
	out, err := system.Convert_string(in)
	if err != nil {
		return value, err
	}
	return String(out), nil
}
func unpacker_Qual(in interface{}) (value Qual, err error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("error unpacking into struct, value should be a map")
	}
	var out Qual
	if v, ok := m["sub"]; ok {
		u, err := func(in interface{}) (value sub.Sub, err error) {
			out, err := sub.Unpackers.Sub(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Sub = u
	}
	return out, nil
}
func unpacker_Pointers(in interface{}) (value Pointers, err error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("error unpacking into struct, value should be a map")
	}
	var out Pointers
	if v, ok := m["string"]; ok {
		u, err := func(in interface{}) (value *string, err error) {
			out, err := func(in interface{}) (value string, err error) {
				out, err := system.Convert_string(in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.String = u
	}
	if v, ok := m["int"]; ok {
		u, err := func(in interface{}) (value *Int, err error) {
			out, err := func(in interface{}) (value Int, err error) {
				out, err := unpacker_Int(in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	if v, ok := m["sub"]; ok {
		u, err := func(in interface{}) (value *sub.Sub, err error) {
			out, err := func(in interface{}) (value sub.Sub, err error) {
				out, err := sub.Unpackers.Sub(in)
				if err != nil {
					return value, err
				}
				return out, nil
			}(in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Sub = u
	}
	if v, ok := m["array"]; ok {
		u, err := func(in interface{}) (value *[3]int, err error) {
			out, err := func(in interface{}) (value [3]int, err error) {
				a, ok := in.([]interface{})
				if !ok {
					return value, errors.New("error unpacking into slice, value should be an array")
				}
				var out [3]int
				if len(a) > 3 {
					return value, errors.Errorf("data length %d does not fit in array of length %d", len(a), 3)
				}
				for i, v := range a {
					u, err := func(in interface{}) (value int, err error) {
						out, err := system.Convert_int(in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(v)
					if err != nil {
						return value, err
					}
					out[i] = u
				}
				return out, nil
			}(in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Array = u
	}
	if v, ok := m["slice"]; ok {
		u, err := func(in interface{}) (value *[]string, err error) {
			out, err := func(in interface{}) (value []string, err error) {
				a, ok := in.([]interface{})
				if !ok {
					return value, errors.New("error unpacking into slice, value should be an array")
				}
				var out = make([]string, len(a))
				for i, v := range a {
					u, err := func(in interface{}) (value string, err error) {
						out, err := system.Convert_string(in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(v)
					if err != nil {
						return value, err
					}
					out[i] = u
				}
				return out[:], nil
			}(in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Slice = u
	}
	if v, ok := m["map"]; ok {
		u, err := func(in interface{}) (value *map[string]int, err error) {
			out, err := func(in interface{}) (value map[string]int, err error) {
				m, ok := in.(map[string]interface{})
				if !ok {
					return value, errors.New("error unpacking into map, value should be a map")
				}
				var out = make(map[string]int, len(m))
				for k, v := range m {
					u, err := func(in interface{}) (value int, err error) {
						out, err := system.Convert_int(in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(v)
					if err != nil {
						return value, err
					}
					out[k] = u
				}
				return out, nil
			}(in)
			if err != nil {
				return value, err
			}
			return &out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Map = u
	}
	if v, ok := m["slice-string"]; ok {
		u, err := func(in interface{}) (value []*string, err error) {
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("error unpacking into slice, value should be an array")
			}
			var out = make([]*string, len(a))
			for i, v := range a {
				u, err := func(in interface{}) (value *string, err error) {
					out, err := func(in interface{}) (value string, err error) {
						out, err := system.Convert_string(in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(in)
					if err != nil {
						return value, err
					}
					return &out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(v)
		if err != nil {
			return value, err
		}
		out.SliceString = u
	}
	if v, ok := m["slice-int"]; ok {
		u, err := func(in interface{}) (value []*Int, err error) {
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("error unpacking into slice, value should be an array")
			}
			var out = make([]*Int, len(a))
			for i, v := range a {
				u, err := func(in interface{}) (value *Int, err error) {
					out, err := func(in interface{}) (value Int, err error) {
						out, err := unpacker_Int(in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(in)
					if err != nil {
						return value, err
					}
					return &out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(v)
		if err != nil {
			return value, err
		}
		out.SliceInt = u
	}
	if v, ok := m["slice-sub"]; ok {
		u, err := func(in interface{}) (value []*sub.Sub, err error) {
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("error unpacking into slice, value should be an array")
			}
			var out = make([]*sub.Sub, len(a))
			for i, v := range a {
				u, err := func(in interface{}) (value *sub.Sub, err error) {
					out, err := func(in interface{}) (value sub.Sub, err error) {
						out, err := sub.Unpackers.Sub(in)
						if err != nil {
							return value, err
						}
						return out, nil
					}(in)
					if err != nil {
						return value, err
					}
					return &out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(v)
		if err != nil {
			return value, err
		}
		out.SliceSub = u
	}
	return out, nil
}
func unpacker_Maps(in interface{}) (value Maps, err error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("error unpacking into struct, value should be a map")
	}
	var out Maps
	if v, ok := m["ints"]; ok {
		u, err := func(in interface{}) (value map[string]int, err error) {
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("error unpacking into map, value should be a map")
			}
			var out = make(map[string]int, len(m))
			for k, v := range m {
				u, err := func(in interface{}) (value int, err error) {
					out, err := system.Convert_int(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Ints = u
	}
	if v, ok := m["strings"]; ok {
		u, err := func(in interface{}) (value map[string]string, err error) {
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("error unpacking into map, value should be a map")
			}
			var out = make(map[string]string, len(m))
			for k, v := range m {
				u, err := func(in interface{}) (value string, err error) {
					out, err := system.Convert_string(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Strings = u
	}
	if v, ok := m["slices"]; ok {
		u, err := func(in interface{}) (value map[string][]string, err error) {
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("error unpacking into map, value should be a map")
			}
			var out = make(map[string][]string, len(m))
			for k, v := range m {
				u, err := func(in interface{}) (value []string, err error) {
					a, ok := in.([]interface{})
					if !ok {
						return value, errors.New("error unpacking into slice, value should be an array")
					}
					var out = make([]string, len(a))
					for i, v := range a {
						u, err := func(in interface{}) (value string, err error) {
							out, err := system.Convert_string(in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(v)
						if err != nil {
							return value, err
						}
						out[i] = u
					}
					return out[:], nil
				}(v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Slices = u
	}
	if v, ok := m["arrays"]; ok {
		u, err := func(in interface{}) (value map[string][2]int, err error) {
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("error unpacking into map, value should be a map")
			}
			var out = make(map[string][2]int, len(m))
			for k, v := range m {
				u, err := func(in interface{}) (value [2]int, err error) {
					a, ok := in.([]interface{})
					if !ok {
						return value, errors.New("error unpacking into slice, value should be an array")
					}
					var out [2]int
					if len(a) > 2 {
						return value, errors.Errorf("data length %d does not fit in array of length %d", len(a), 2)
					}
					for i, v := range a {
						u, err := func(in interface{}) (value int, err error) {
							out, err := system.Convert_int(in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(v)
						if err != nil {
							return value, err
						}
						out[i] = u
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Arrays = u
	}
	if v, ok := m["maps"]; ok {
		u, err := func(in interface{}) (value map[string]map[string]string, err error) {
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("error unpacking into map, value should be a map")
			}
			var out = make(map[string]map[string]string, len(m))
			for k, v := range m {
				u, err := func(in interface{}) (value map[string]string, err error) {
					m, ok := in.(map[string]interface{})
					if !ok {
						return value, errors.New("error unpacking into map, value should be a map")
					}
					var out = make(map[string]string, len(m))
					for k, v := range m {
						u, err := func(in interface{}) (value string, err error) {
							out, err := system.Convert_string(in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(v)
						if err != nil {
							return value, err
						}
						out[k] = u
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out[k] = u
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Maps = u
	}
	return out, nil
}
func unpacker_Slices(in interface{}) (value Slices, err error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("error unpacking into struct, value should be a map")
	}
	var out Slices
	if v, ok := m["ints"]; ok {
		u, err := func(in interface{}) (value []int, err error) {
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("error unpacking into slice, value should be an array")
			}
			var out = make([]int, len(a))
			for i, v := range a {
				u, err := func(in interface{}) (value int, err error) {
					out, err := system.Convert_int(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Ints = u
	}
	if v, ok := m["strings"]; ok {
		u, err := func(in interface{}) (value []string, err error) {
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("error unpacking into slice, value should be an array")
			}
			var out = make([]string, len(a))
			for i, v := range a {
				u, err := func(in interface{}) (value string, err error) {
					out, err := system.Convert_string(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Strings = u
	}
	if v, ok := m["array-lit"]; ok {
		u, err := func(in interface{}) (value [5]string, err error) {
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("error unpacking into slice, value should be an array")
			}
			var out [5]string
			if len(a) > 5 {
				return value, errors.Errorf("data length %d does not fit in array of length %d", len(a), 5)
			}
			for i, v := range a {
				u, err := func(in interface{}) (value string, err error) {
					out, err := system.Convert_string(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.ArrayLit = u
	}
	if v, ok := m["array-expr"]; ok {
		u, err := func(in interface{}) (value [2 * N]int, err error) {
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("error unpacking into slice, value should be an array")
			}
			var out [2 * N]int
			if len(a) > 2*N {
				return value, errors.Errorf("data length %d does not fit in array of length %d", len(a), 2*N)
			}
			for i, v := range a {
				u, err := func(in interface{}) (value int, err error) {
					out, err := system.Convert_int(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.ArrayExpr = u
	}
	if v, ok := m["structs"]; ok {
		u, err := func(in interface{}) (value []struct {
			Int int
		}, err error) {
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("error unpacking into slice, value should be an array")
			}
			var out = make([]struct {
				Int int
			}, len(a))
			for i, v := range a {
				u, err := func(in interface{}) (value struct {
					Int int
				}, err error) {
					m, ok := in.(map[string]interface{})
					if !ok {
						return value, errors.New("error unpacking into struct, value should be a map")
					}
					var out struct {
						Int int
					}
					if v, ok := m["int"]; ok {
						u, err := func(in interface{}) (value int, err error) {
							out, err := system.Convert_int(in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(v)
						if err != nil {
							return value, err
						}
						out.Int = u
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Structs = u
	}
	if v, ok := m["arrays"]; ok {
		u, err := func(in interface{}) (value [][]string, err error) {
			a, ok := in.([]interface{})
			if !ok {
				return value, errors.New("error unpacking into slice, value should be an array")
			}
			var out = make([][]string, len(a))
			for i, v := range a {
				u, err := func(in interface{}) (value []string, err error) {
					a, ok := in.([]interface{})
					if !ok {
						return value, errors.New("error unpacking into slice, value should be an array")
					}
					var out = make([]string, len(a))
					for i, v := range a {
						u, err := func(in interface{}) (value string, err error) {
							out, err := system.Convert_string(in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(v)
						if err != nil {
							return value, err
						}
						out[i] = u
					}
					return out[:], nil
				}(v)
				if err != nil {
					return value, err
				}
				out[i] = u
			}
			return out[:], nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Arrays = u
	}
	return out, nil
}
func unpacker_Structs(in interface{}) (value Structs, err error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("error unpacking into struct, value should be a map")
	}
	var out Structs
	if v, ok := m["simple"]; ok {
		u, err := func(in interface{}) (value struct {
			Int  int
			Bool bool
		}, err error) {
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("error unpacking into struct, value should be a map")
			}
			var out struct {
				Int  int
				Bool bool
			}
			if v, ok := m["int"]; ok {
				u, err := func(in interface{}) (value int, err error) {
					out, err := system.Convert_int(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out.Int = u
			}
			if v, ok := m["bool"]; ok {
				u, err := func(in interface{}) (value bool, err error) {
					out, err := system.Convert_bool(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out.Bool = u
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Simple = u
	}
	if v, ok := m["complex"]; ok {
		u, err := func(in interface{}) (value struct {
			String string
			Inner  struct {
				Float32 float32
			}
		}, err error) {
			m, ok := in.(map[string]interface{})
			if !ok {
				return value, errors.New("error unpacking into struct, value should be a map")
			}
			var out struct {
				String string
				Inner  struct {
					Float32 float32
				}
			}
			if v, ok := m["string"]; ok {
				u, err := func(in interface{}) (value string, err error) {
					out, err := system.Convert_string(in)
					if err != nil {
						return value, err
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out.String = u
			}
			if v, ok := m["inner"]; ok {
				u, err := func(in interface{}) (value struct {
					Float32 float32
				}, err error) {
					m, ok := in.(map[string]interface{})
					if !ok {
						return value, errors.New("error unpacking into struct, value should be a map")
					}
					var out struct {
						Float32 float32
					}
					if v, ok := m["float-32"]; ok {
						u, err := func(in interface{}) (value float32, err error) {
							out, err := system.Convert_float32(in)
							if err != nil {
								return value, err
							}
							return out, nil
						}(v)
						if err != nil {
							return value, err
						}
						out.Float32 = u
					}
					return out, nil
				}(v)
				if err != nil {
					return value, err
				}
				out.Inner = u
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Complex = u
	}
	return out, nil
}
func unpacker_Natives(in interface{}) (value Natives, err error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return value, errors.New("error unpacking into struct, value should be a map")
	}
	var out Natives
	if v, ok := m["bool"]; ok {
		u, err := func(in interface{}) (value bool, err error) {
			out, err := system.Convert_bool(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Bool = u
	}
	if v, ok := m["byte"]; ok {
		u, err := func(in interface{}) (value byte, err error) {
			out, err := system.Convert_byte(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Byte = u
	}
	if v, ok := m["float-32"]; ok {
		u, err := func(in interface{}) (value float32, err error) {
			out, err := system.Convert_float32(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Float32 = u
	}
	if v, ok := m["float-64"]; ok {
		u, err := func(in interface{}) (value float64, err error) {
			out, err := system.Convert_float64(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Float64 = u
	}
	if v, ok := m["int"]; ok {
		u, err := func(in interface{}) (value int, err error) {
			out, err := system.Convert_int(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Int = u
	}
	if v, ok := m["int-8"]; ok {
		u, err := func(in interface{}) (value int8, err error) {
			out, err := system.Convert_int8(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Int8 = u
	}
	if v, ok := m["int-16"]; ok {
		u, err := func(in interface{}) (value int16, err error) {
			out, err := system.Convert_int16(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Int16 = u
	}
	if v, ok := m["int-32"]; ok {
		u, err := func(in interface{}) (value int32, err error) {
			out, err := system.Convert_int32(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Int32 = u
	}
	if v, ok := m["int-64"]; ok {
		u, err := func(in interface{}) (value int64, err error) {
			out, err := system.Convert_int64(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Int64 = u
	}
	if v, ok := m["uint"]; ok {
		u, err := func(in interface{}) (value uint, err error) {
			out, err := system.Convert_uint(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Uint = u
	}
	if v, ok := m["uint-8"]; ok {
		u, err := func(in interface{}) (value uint8, err error) {
			out, err := system.Convert_uint8(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Uint8 = u
	}
	if v, ok := m["uint-16"]; ok {
		u, err := func(in interface{}) (value uint16, err error) {
			out, err := system.Convert_uint16(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Uint16 = u
	}
	if v, ok := m["uint-32"]; ok {
		u, err := func(in interface{}) (value uint32, err error) {
			out, err := system.Convert_uint32(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Uint32 = u
	}
	if v, ok := m["uint-64"]; ok {
		u, err := func(in interface{}) (value uint64, err error) {
			out, err := system.Convert_uint64(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Uint64 = u
	}
	if v, ok := m["rune"]; ok {
		u, err := func(in interface{}) (value rune, err error) {
			out, err := system.Convert_rune(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.Rune = u
	}
	if v, ok := m["string"]; ok {
		u, err := func(in interface{}) (value string, err error) {
			out, err := system.Convert_string(in)
			if err != nil {
				return value, err
			}
			return out, nil
		}(v)
		if err != nil {
			return value, err
		}
		out.String = u
	}
	return out, nil
}
