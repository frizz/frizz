package unpacker

import (
	system "frizz.io/system"
	errors "github.com/pkg/errors"
)

func unpackStructs(in interface{}) (value Structs, err error) {
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
func unpackNatives(in interface{}) (value Natives, err error) {
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
