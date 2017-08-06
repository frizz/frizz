package validators

import (
	"fmt"

	"frizz.io/global"

	"reflect"
)

// frizz
type IsNull struct{}

func (i IsNull) Validate(context global.ValidationContext, input interface{}) (valid bool, message string, err error) {
	return i.ValidateValue(context, reflect.ValueOf(input))
}

func (i IsNull) ValidateValue(context global.ValidationContext, value reflect.Value) (valid bool, message string, err error) {
	for value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
		value = value.Elem()
	}
	if value.IsValid() {
		// input not null -> return invalid
		return false, fmt.Sprintf("%s: value %v must be null", context.Location(), value.Interface()), nil
	}
	return true, "", nil
}

// frizz
type NotNull struct{}

func (n NotNull) Validate(context global.ValidationContext, input interface{}) (valid bool, message string, err error) {
	return n.ValidateValue(context, reflect.ValueOf(input))
}

func (n NotNull) ValidateValue(context global.ValidationContext, value reflect.Value) (valid bool, message string, err error) {
	for value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
		value = value.Elem()
	}
	if !value.IsValid() {
		// input null -> return invalid
		return false, fmt.Sprintf("%s: value must not be null", context.Location()), nil
	}
	return true, "", nil
}

/*
// frizz
type Zero bool

func (z Zero) Validate(context global.ValidationContext, input interface{}) (valid bool, message string, err error) {
	var zero bool
	switch input.(type) {
	case bool:
		if !input.(bool) {
			zero = true
		}
	case string:
		if input.(string) == "" {
			zero = true
		}
	case int, int8, int16, int32, int64:
		input := MustInt64(input)
		if input == 0 {
			zero = true
		}
	case uint, uint8, uint16, uint32, uint64:
		input := MustUint64(input)
		if input == 0 {
			zero = true
		}
	case float32, float64:
		input := MustFloat64(input)
		if input == 0.0 {
			zero = true
		}
	default:
		if input == nil {
			zero = true
		}
	}
	if bool(z) && !zero {
		// Zero(true) -> must be zero -> input not zero -> return invalid
		return false, fmt.Sprintf("%s: value %v must be zero", context.Location(), input), nil
	} else if !bool(z) && zero {
		// Zero(false) -> must not be zero -> input zero -> return invalid
		return false, fmt.Sprintf("%s: value %v must not be zero", context.Location(), input), nil
	}
	return true, "", nil
}
*/
