package pack

import (
	"encoding/json"

	"strconv"

	"fmt"

	"frizz.io/global"
	"github.com/pkg/errors"
)

func RepackNative(stack global.Location, name string, in interface{}) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "bool":
		return RepackBool(in)
	case "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64":
		return RepackNumber(in)
	case "rune", "string":
		return RepackString(in)
	}
	return nil, false, false, errors.Errorf("%s: unknown type %s", stack, name)
}

func RepackBool(in interface{}) (value interface{}, dict bool, null bool, err error) {
	out := in.(bool)
	if out == false {
		return out, false, true, nil
	}
	return out, false, false, nil
}

func RepackNumber(in interface{}) (value interface{}, dict bool, null bool, err error) {
	out := fmt.Sprintf("%v", in)
	if out == "0" || out == "0.0" {
		return json.Number(out), false, true, nil
	}
	return json.Number(out), false, false, nil
}

func RepackString(in interface{}) (value interface{}, dict bool, null bool, err error) {
	var out string
	if r, ok := in.(rune); ok {
		if r == 0 {
			return out, false, true, nil
		}
		out = fmt.Sprintf("%c", r)
	} else {
		out = fmt.Sprintf("%s", in)
		if out == "" {
			return out, false, true, nil
		}
	}
	return out, false, false, nil
}

// unsupported: "complex64", "complex128", "error", "uintptr"

const (
	errorMessage      = "%s: unpacking into %s, value %#v should be %s"
	errorMessageFinal = "%s: converting %v to %s"
)

func UnpackNative(location global.Location, name string, in interface{}) (value interface{}, null bool, err error) {
	// notest
	switch name {
	case "bool":
		return UnpackBool(location, in)
	case "byte":
		return UnpackByte(location, in)
	case "float32":
		return UnpackFloat32(location, in)
	case "float64":
		return UnpackFloat64(location, in)
	case "int":
		return UnpackInt(location, in)
	case "int8":
		return UnpackInt8(location, in)
	case "int16":
		return UnpackInt16(location, in)
	case "int32":
		return UnpackInt32(location, in)
	case "uint":
		return UnpackUint(location, in)
	case "uint8":
		return UnpackUint8(location, in)
	case "uint16":
		return UnpackUint16(location, in)
	case "uint32":
		return UnpackUint32(location, in)
	case "int64":
		return UnpackInt64(location, in)
	case "uint64":
		return UnpackUint64(location, in)
	case "rune":
		return UnpackRune(location, in)
	case "string":
		return UnpackString(location, in)
	}
	return nil, false, errors.Errorf("%s: unknown type %s", location, name)
}

func UnpackBool(location global.Location, in interface{}) (bool, bool, error) {
	if in == nil {
		return false, false, nil
	}
	b, ok := in.(bool)
	if !ok {
		return false, false, errors.Errorf(errorMessage, location, "bool", in, "bool")
	}
	return b, false, nil
}

func UnpackByte(location global.Location, in interface{}) (byte, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "byte", in, "json.Number")
	}
	i, err := strconv.ParseUint(string(n), 10, 8)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "byte")
	}
	return byte(i), false, nil
}

func UnpackFloat32(location global.Location, in interface{}) (float32, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "float32", in, "json.Number")
	}
	f, err := strconv.ParseFloat(string(n), 32)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "float32")
	}
	return float32(f), false, nil
}

func UnpackFloat64(location global.Location, in interface{}) (float64, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "float64", in, "json.Number")
	}
	f, err := strconv.ParseFloat(string(n), 64)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "float64")
	}
	return f, false, nil
}

func UnpackInt(location global.Location, in interface{}) (int, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "int", in, "json.Number")
	}
	i, err := strconv.ParseInt(string(n), 10, 0)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "int")
	}
	return int(i), false, nil
}

func UnpackInt8(location global.Location, in interface{}) (int8, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "int8", in, "json.Number")
	}
	i, err := strconv.ParseInt(string(n), 10, 8)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "int8")
	}
	return int8(i), false, nil
}

func UnpackInt16(location global.Location, in interface{}) (int16, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "int16", in, "json.Number")
	}
	i, err := strconv.ParseInt(string(n), 10, 16)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "int16")
	}
	return int16(i), false, nil
}

func UnpackInt32(location global.Location, in interface{}) (int32, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "int32", in, "json.Number")
	}
	i, err := strconv.ParseInt(string(n), 10, 32)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "int32")
	}
	return int32(i), false, nil
}

func UnpackInt64(location global.Location, in interface{}) (int64, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "int64", in, "json.Number")
	}
	i, err := strconv.ParseInt(string(n), 10, 64)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "int64")
	}
	return i, false, nil
}

func UnpackUint(location global.Location, in interface{}) (uint, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "uint", in, "json.Number")
	}
	i, err := strconv.ParseUint(string(n), 10, 0)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "uint")
	}
	return uint(i), false, nil
}

func UnpackUint8(location global.Location, in interface{}) (uint8, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "uint8", in, "json.Number")
	}
	i, err := strconv.ParseUint(string(n), 10, 8)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "uint8")
	}
	return uint8(i), false, nil
}

func UnpackUint16(location global.Location, in interface{}) (uint16, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "uint16", in, "json.Number")
	}
	i, err := strconv.ParseUint(string(n), 10, 16)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "uint16")
	}
	return uint16(i), false, nil
}

func UnpackUint32(location global.Location, in interface{}) (uint32, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "uint32", in, "json.Number")
	}
	i, err := strconv.ParseUint(string(n), 10, 32)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "uint32")
	}
	return uint32(i), false, nil
}

func UnpackUint64(location global.Location, in interface{}) (uint64, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	n, ok := in.(json.Number)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "uint64", in, "json.Number")
	}
	i, err := strconv.ParseUint(string(n), 10, 64)
	if err != nil {
		return 0, false, errors.Wrapf(err, errorMessageFinal, location, n, "uint64")
	}
	return i, false, nil
}

func UnpackString(location global.Location, in interface{}) (string, bool, error) {
	if in == nil {
		return "", false, nil
	}
	s, ok := in.(string)
	if !ok {
		return "", false, errors.Errorf(errorMessage, location, "string", in, "string")
	}
	return s, false, nil
}

func UnpackRune(location global.Location, in interface{}) (rune, bool, error) {
	if in == nil {
		return 0, false, nil
	}
	s, ok := in.(string)
	if !ok {
		return 0, false, errors.Errorf(errorMessage, location, "rune", in, "string")
	}
	var out rune
	for i, r := range s {
		if i > 0 {
			return rune(0), false, errors.Errorf("%s: unpacking into rune: string should have a single rune", location)
		}
		out = r
	}
	return out, false, nil
}
