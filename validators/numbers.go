package validators

import (
	"encoding/json"
	"fmt"
	"strconv"

	"reflect"

	"github.com/pkg/errors"
)

// frizz
type GreaterThan json.Number

func (g GreaterThan) Validate(input interface{}) (valid bool, message string, err error) {
	return g.ValidateValue(reflect.ValueOf(input))
}

func (g GreaterThan) ValidateValue(value reflect.Value) (valid bool, message string, err error) {
	return validate(value, ">", json.Number(g))
}

// frizz
type LessThan json.Number

func (l LessThan) Validate(input interface{}) (valid bool, message string, err error) {
	return l.ValidateValue(reflect.ValueOf(input))
}

func (l LessThan) ValidateValue(value reflect.Value) (valid bool, message string, err error) {
	return validate(value, "<", json.Number(l))
}

// frizz
type GreaterThanOrEqual json.Number

func (g GreaterThanOrEqual) Validate(input interface{}) (valid bool, message string, err error) {
	return g.ValidateValue(reflect.ValueOf(input))
}

func (g GreaterThanOrEqual) ValidateValue(value reflect.Value) (valid bool, message string, err error) {
	return validate(value, ">=", json.Number(g))
}

// frizz
type LessThanOrEqual json.Number

func (l LessThanOrEqual) Validate(input interface{}) (valid bool, message string, err error) {
	return l.ValidateValue(reflect.ValueOf(input))
}

func (l LessThanOrEqual) ValidateValue(value reflect.Value) (valid bool, message string, err error) {
	return validate(value, "<=", json.Number(l))
}

func validate(value reflect.Value, operator string, comparison json.Number) (valid bool, message string, err error) {
	switch value.Type().Kind() {
	case reflect.Interface, reflect.Ptr:
		// interface or ptr: recurse with elem
		return validate(value.Elem(), operator, comparison)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		input := MustInt64Value(value)
		if input == 0 {
			return true, "", nil
		}
		i, err := strconv.ParseInt(string(comparison), 10, 64)
		if err != nil {
			return false, "", errors.Wrapf(err, "unpacking int64")
		}
		if operator == ">" {
			if input <= i {
				return false, fmt.Sprintf("value %v must be greater than %v", input, i), nil
			}
		} else if operator == "<" {
			if input >= i {
				return false, fmt.Sprintf("value %v must be less than %v", input, i), nil
			}
		} else if operator == ">=" {
			if input < i {
				return false, fmt.Sprintf("value %v must be greater than or equal to %v", input, i), nil
			}
		} else if operator == "<=" {
			if input > i {
				return false, fmt.Sprintf("value %v must be less than or equal to %v", input, i), nil
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		input := MustUint64Value(value)
		if input == 0 {
			return true, "", nil
		}
		i, err := strconv.ParseUint(string(comparison), 10, 64)
		if err != nil {
			return false, "", errors.Wrapf(err, "unpacking uint64")
		}
		if operator == ">" {
			if input <= i {
				return false, fmt.Sprintf("value %v must be greater than %v", input, i), nil
			}
		} else if operator == "<" {
			if input >= i {
				return false, fmt.Sprintf("value %v must be less than %v", input, i), nil
			}
		} else if operator == ">=" {
			if input < i {
				return false, fmt.Sprintf("value %v must be greater than or equal to %v", input, i), nil
			}
		} else if operator == "<=" {
			if input > i {
				return false, fmt.Sprintf("value %v must be less than or equal to %v", input, i), nil
			}
		}
	case reflect.Float32, reflect.Float64:
		input := MustFloat64Value(value)
		if input == 0.0 {
			return true, "", nil
		}
		f, err := strconv.ParseFloat(string(comparison), 64)
		if err != nil {
			return false, "", errors.Wrapf(err, "unpacking float64")
		}
		if operator == ">" {
			if input <= f {
				return false, fmt.Sprintf("value %v must be greater than %v", input, f), nil
			}
		} else if operator == "<" {
			if input >= f {
				return false, fmt.Sprintf("value %v must be less than %v", input, f), nil
			}
		} else if operator == ">=" {
			if input < f {
				return false, fmt.Sprintf("value %v must be greater than or equal to %v", input, f), nil
			}
		} else if operator == "<=" {
			if input > f {
				return false, fmt.Sprintf("value %v must be less than or equal to %v", input, f), nil
			}
		}
	default:
		return false, "", errors.Errorf("unknown type for numeric comparison %T", value.Interface())
	}
	return true, "", nil
}
