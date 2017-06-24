package frizz

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

func RepackNative(stack Stack, name string, in interface{}) (value interface{}, dict bool, null bool, err error) {
	// notest
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
