package validators

import (
	"fmt"

	"frizz.io/frizz"
)

// frizz
type NotZero struct{}

func (n NotZero) Validate(stack frizz.Stack, input interface{}) (valid bool, message string, err error) {
	var fail bool
	switch input.(type) {
	case bool:
		if !input.(bool) {
			fail = true
		}
	case string:
		if input.(string) == "" {
			fail = true
		}
	case int, int8, int16, int32, int64:
		input := MustInt64(input)
		if input == 0 {
			fail = true
		}
	case uint, uint8, uint16, uint32, uint64:
		input := MustUint64(input)
		if input == 0 {
			fail = true
		}
	case float32, float64:
		input := MustFloat64(input)
		if input == 0.0 {
			fail = true
		}
	default:
		if input == nil {
			fail = true
		}
	}
	if fail {
		return false, fmt.Sprintf("%s: must not be zero value", stack), nil
	}
	return true, "", nil
}
