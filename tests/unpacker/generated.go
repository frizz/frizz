package unpacker

import (
	system "frizz.io/system"
	errors "github.com/pkg/errors"
)

func unpackNatives(in interface{}) (interface{}, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, errors.Errorf("error unpacking into %s, value should be a map", "Natives")
	}
	out := new(Natives)
	if v, ok := m["bool"]; ok {
		c, err := system.Convert_bool(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Bool = c
	}
	if v, ok := m["byte"]; ok {
		c, err := system.Convert_byte(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Byte = c
	}
	if v, ok := m["float-32"]; ok {
		c, err := system.Convert_float32(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Float32 = c
	}
	if v, ok := m["float-64"]; ok {
		c, err := system.Convert_float64(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Float64 = c
	}
	if v, ok := m["int"]; ok {
		c, err := system.Convert_int(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Int = c
	}
	if v, ok := m["int-8"]; ok {
		c, err := system.Convert_int8(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Int8 = c
	}
	if v, ok := m["int-16"]; ok {
		c, err := system.Convert_int16(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Int16 = c
	}
	if v, ok := m["int-32"]; ok {
		c, err := system.Convert_int32(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Int32 = c
	}
	if v, ok := m["int-64"]; ok {
		c, err := system.Convert_int64(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Int64 = c
	}
	if v, ok := m["uint"]; ok {
		c, err := system.Convert_uint(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Uint = c
	}
	if v, ok := m["uint-8"]; ok {
		c, err := system.Convert_uint8(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Uint8 = c
	}
	if v, ok := m["uint-16"]; ok {
		c, err := system.Convert_uint16(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Uint16 = c
	}
	if v, ok := m["uint-32"]; ok {
		c, err := system.Convert_uint32(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Uint32 = c
	}
	if v, ok := m["uint-64"]; ok {
		c, err := system.Convert_uint64(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Uint64 = c
	}
	if v, ok := m["rune"]; ok {
		c, err := system.Convert_rune(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.Rune = c
	}
	if v, ok := m["string"]; ok {
		c, err := system.Convert_string(v)
		if err != nil {
			return nil, errors.WithMessage(err, "error unpacking")
		}
		out.String = c
	}
	return out, nil
}
