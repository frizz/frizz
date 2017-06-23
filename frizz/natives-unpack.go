package frizz

import (
	"encoding/json"

	"strconv"

	"github.com/pkg/errors"
)

// unsupported: "complex64", "complex128", "error", "uintptr"

const (
	errorMessage      = "%s: unpacking into %s, value %#v should be %s"
	errorMessageFinal = "%s: converting %v to %s"
)

func UnpackNative(stack Stack, name string, in interface{}) (interface{}, error) {
	// notest
	switch name {
	case "bool":
		return UnpackBool(stack, in)
	case "byte":
		return UnpackByte(stack, in)
	case "float32":
		return UnpackFloat32(stack, in)
	case "float64":
		return UnpackFloat64(stack, in)
	case "int":
		return UnpackInt(stack, in)
	case "int8":
		return UnpackInt8(stack, in)
	case "int16":
		return UnpackInt16(stack, in)
	case "int32":
		return UnpackInt32(stack, in)
	case "uint":
		return UnpackUint(stack, in)
	case "uint8":
		return UnpackUint8(stack, in)
	case "uint16":
		return UnpackUint16(stack, in)
	case "uint32":
		return UnpackUint32(stack, in)
	case "int64":
		return UnpackInt64(stack, in)
	case "uint64":
		return UnpackUint64(stack, in)
	case "rune":
		return UnpackRune(stack, in)
	case "string":
		return UnpackString(stack, in)
	}
	return nil, errors.Errorf("%s: unknown type %s", stack, name)
}

func UnpackBool(stack Stack, in interface{}) (bool, error) {
	b, ok := in.(bool)
	if !ok {
		return false, errors.Errorf(errorMessage, stack, "bool", in, "bool")
	}
	return b, nil
}

func UnpackByte(stack Stack, in interface{}) (byte, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "byte", in, "json.Number")
	}
	i, err := strconv.ParseUint(string(n), 10, 8)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "byte")
	}
	return byte(i), nil
}

func UnpackFloat32(stack Stack, in interface{}) (float32, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "float32", in, "json.Number")
	}
	f, err := strconv.ParseFloat(string(n), 32)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "float32")
	}
	return float32(f), nil
}

func UnpackFloat64(stack Stack, in interface{}) (float64, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "float64", in, "json.Number")
	}
	f, err := strconv.ParseFloat(string(n), 64)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "float64")
	}
	return f, nil
}

func UnpackInt(stack Stack, in interface{}) (int, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "int", in, "json.Number")
	}
	i, err := strconv.ParseInt(string(n), 10, 0)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "int")
	}
	return int(i), nil
}

func UnpackInt8(stack Stack, in interface{}) (int8, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "int8", in, "json.Number")
	}
	i, err := strconv.ParseInt(string(n), 10, 8)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "int8")
	}
	return int8(i), nil
}

func UnpackInt16(stack Stack, in interface{}) (int16, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "int16", in, "json.Number")
	}
	i, err := strconv.ParseInt(string(n), 10, 16)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "int16")
	}
	return int16(i), nil
}

func UnpackInt32(stack Stack, in interface{}) (int32, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "int32", in, "json.Number")
	}
	i, err := strconv.ParseInt(string(n), 10, 32)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "int32")
	}
	return int32(i), nil
}

func UnpackInt64(stack Stack, in interface{}) (int64, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "int64", in, "json.Number")
	}
	i, err := strconv.ParseInt(string(n), 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "int64")
	}
	return i, nil
}

func UnpackUint(stack Stack, in interface{}) (uint, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "uint", in, "json.Number")
	}
	i, err := strconv.ParseUint(string(n), 10, 0)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "uint")
	}
	return uint(i), nil
}

func UnpackUint8(stack Stack, in interface{}) (uint8, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "uint8", in, "json.Number")
	}
	i, err := strconv.ParseUint(string(n), 10, 8)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "uint8")
	}
	return uint8(i), nil
}

func UnpackUint16(stack Stack, in interface{}) (uint16, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "uint16", in, "json.Number")
	}
	i, err := strconv.ParseUint(string(n), 10, 16)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "uint16")
	}
	return uint16(i), nil
}

func UnpackUint32(stack Stack, in interface{}) (uint32, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "uint32", in, "json.Number")
	}
	i, err := strconv.ParseUint(string(n), 10, 32)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "uint32")
	}
	return uint32(i), nil
}

func UnpackUint64(stack Stack, in interface{}) (uint64, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "uint64", in, "json.Number")
	}
	i, err := strconv.ParseUint(string(n), 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, stack, n, "uint64")
	}
	return i, nil
}

func UnpackString(stack Stack, in interface{}) (string, error) {
	s, ok := in.(string)
	if !ok {
		return "", errors.Errorf(errorMessage, stack, "string", in, "string")
	}
	return s, nil
}

func UnpackRune(stack Stack, in interface{}) (rune, error) {
	s, ok := in.(string)
	if !ok {
		return 0, errors.Errorf(errorMessage, stack, "rune", in, "string")
	}
	var out rune
	for i, r := range s {
		if i > 0 {
			return rune(0), errors.Errorf("%s: unpacking into rune: string should have a single rune", stack)
		}
		out = r
	}
	return out, nil
}
