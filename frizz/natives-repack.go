package frizz

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func RepackNative(stack Stack, name string, in interface{}) (interface{}, error) {
	// notest
	switch name {
	case "bool":
		return RepackBool(in), nil
	case "byte", "float32", "float64", "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32", "int64", "uint64":
		return RepackNumber(in), nil
	case "rune", "string":
		return RepackString(in), nil
	}
	return nil, errors.Errorf("%s: unknown type %s", stack, name)
}

func RepackBool(in interface{}) interface{} {
	return in.(bool)
}

func RepackNumber(in interface{}) interface{} {
	return json.Number(fmt.Sprintf("%v", in))
}

func RepackString(in interface{}) interface{} {
	return fmt.Sprintf("%v", in)
}
