package validators

import (
	"fmt"
	"reflect"
)

func MustInt64(input interface{}) int64 {
	switch input := input.(type) {
	case int:
		return int64(input)
	case int8:
		return int64(input)
	case int16:
		return int64(input)
	case int32:
		return int64(input)
	case int64:
		return input
	default:
		panic(fmt.Sprintf("MustInt64 must be given a int, int8, int16, int32 or int64. Found %T", input))
	}
}

func MustUint64(input interface{}) uint64 {
	switch input := input.(type) {
	case uint:
		return uint64(input)
	case uint8:
		return uint64(input)
	case uint16:
		return uint64(input)
	case uint32:
		return uint64(input)
	case uint64:
		return input
	default:
		panic(fmt.Sprintf("MustUint64 must be given a uint, uint8, uint16, uint32 or uint64. Found %T", input))
	}
}

func MustFloat64(input interface{}) float64 {
	switch input := input.(type) {
	case float32:
		return float64(input)
	case float64:
		return input
	default:
		panic(fmt.Sprintf("MustFloat64 must be given a float32 or float64. Found %T", input))
	}
}

var stringtype = reflect.TypeOf("")
var int64type = reflect.TypeOf(int64(0))
var uint64type = reflect.TypeOf(uint64(0))
var float64type = reflect.TypeOf(float64(0))
