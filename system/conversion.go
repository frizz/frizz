package system

import (
	"encoding/json"

	"strconv"

	"github.com/pkg/errors"
)

// unsupported: "complex64", "complex128", "error", "uintptr"

const (
	errorMessage      = "error unpacking into %s, value %#v should be %s"
	errorMessageFinal = "error converting %v to %s"
)

func Convert(name string, in interface{}) (interface{}, error) {
	switch name {
	case "bool":
		return Convert_bool(in)
	case "byte":
		return Convert_byte(in)
	case "float32":
		return Convert_float32(in)
	case "float64":
		return Convert_float64(in)
	case "int":
		return Convert_int(in)
	case "int8":
		return Convert_int8(in)
	case "int16":
		return Convert_int16(in)
	case "int32":
		return Convert_int32(in)
	case "uint":
		return Convert_uint(in)
	case "uint8":
		return Convert_uint8(in)
	case "uint16":
		return Convert_uint16(in)
	case "uint32":
		return Convert_uint32(in)
	case "int64":
		return Convert_int64(in)
	case "uint64":
		return Convert_uint64(in)
	case "rune":
		return Convert_rune(in)
	case "string":
		return Convert_string(in)
	}
	return nil, errors.Errorf("unknown type %s", name)
}

func Convert_bool(in interface{}) (bool, error) {
	b, ok := in.(bool)
	if !ok {
		return false, errors.Errorf(errorMessage, "bool", in, "bool") // VPGKI
	}
	return b, nil
}

func Convert_byte(in interface{}) (byte, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "byte", in, "json.Number") // LKYXV
	}
	i, err := strconv.ParseUint(string(n), 10, 8)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "byte") // JPAPE
	}
	return byte(i), nil
}

func Convert_float32(in interface{}) (float32, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "float32", in, "json.Number") // RBEXS
	}
	f, err := strconv.ParseFloat(string(n), 32)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "float32") // SAHPU
	}
	return float32(f), nil
}

func Convert_float64(in interface{}) (float64, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "float64", in, "json.Number") // YRENT
	}
	f, err := strconv.ParseFloat(string(n), 64)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "float64") // INKLX
	}
	return f, nil
}

func Convert_int(in interface{}) (int, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "int", in, "json.Number") // DMSHN
	}
	i, err := strconv.ParseInt(string(n), 10, 0)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "int") // UGYQH
	}
	return int(i), nil
}

func Convert_int8(in interface{}) (int8, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "int8", in, "json.Number") // PQDKT
	}
	i, err := strconv.ParseInt(string(n), 10, 8)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "int8") // JWHEX
	}
	return int8(i), nil
}

func Convert_int16(in interface{}) (int16, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "int16", in, "json.Number") // YBJBD
	}
	i, err := strconv.ParseInt(string(n), 10, 16)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "int16") // VJGCV
	}
	return int16(i), nil
}

func Convert_int32(in interface{}) (int32, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "int32", in, "json.Number") // DCWYT
	}
	i, err := strconv.ParseInt(string(n), 10, 32)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "int32") // BRQTC
	}
	return int32(i), nil
}

func Convert_int64(in interface{}) (int64, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "int64", in, "json.Number") // DOFWU
	}
	i, err := strconv.ParseInt(string(n), 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "int64") // EPAHS
	}
	return i, nil
}

func Convert_uint(in interface{}) (uint, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "uint", in, "json.Number") // CIJOH
	}
	i, err := strconv.ParseUint(string(n), 10, 0)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "uint") // FMPHN
	}
	return uint(i), nil
}

func Convert_uint8(in interface{}) (uint8, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "uint8", in, "json.Number") // DBUNO
	}
	i, err := strconv.ParseUint(string(n), 10, 8)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "uint8") // HBACM
	}
	return uint8(i), nil
}

func Convert_uint16(in interface{}) (uint16, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "uint16", in, "json.Number") // SKVIO
	}
	i, err := strconv.ParseUint(string(n), 10, 16)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "uint16") // CXCUQ
	}
	return uint16(i), nil
}

func Convert_uint32(in interface{}) (uint32, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "uint32", in, "json.Number") // UJMCO
	}
	i, err := strconv.ParseUint(string(n), 10, 32)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "uint32") // ILJLU
	}
	return uint32(i), nil
}

func Convert_uint64(in interface{}) (uint64, error) {
	n, ok := in.(json.Number)
	if !ok {
		return 0, errors.Errorf(errorMessage, "uint64", in, "json.Number") // KCTKC
	}
	i, err := strconv.ParseUint(string(n), 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, errorMessageFinal, n, "uint64") // QSFRS
	}
	return i, nil
}

func Convert_string(in interface{}) (string, error) {
	s, ok := in.(string)
	if !ok {
		return "", errors.Errorf(errorMessage, "string", in, "string") // YTFIW
	}
	return s, nil
}

func Convert_rune(in interface{}) (rune, error) {
	s, ok := in.(string)
	if !ok {
		return 0, errors.Errorf(errorMessage, "rune", in, "string") // MPHSY
	}
	var out rune
	for i, r := range s {
		if i > 0 {
			return rune(0), errors.New("error unpacking into rune: string should have a single rune") // HGHCR
		}
		out = r
	}
	return out, nil
}
