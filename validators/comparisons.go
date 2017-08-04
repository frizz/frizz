package validators

import (
	"encoding/json"
	"fmt"
	"strconv"

	"reflect"

	"frizz.io/global"
	"frizz.io/system"
	"github.com/pkg/errors"
)

// frizz
type GreaterThan json.Number

func (g GreaterThan) Validate(stack global.Stack, input interface{}) (valid bool, message string, err error) {
	return g.ValidateValue(stack, reflect.ValueOf(input))
}

func (g GreaterThan) ValidateValue(stack global.Stack, value reflect.Value) (valid bool, message string, err error) {
	return compare(stack, value, ">", json.Number(g))
}

// frizz
type LessThan json.Number

func (l LessThan) Validate(stack global.Stack, input interface{}) (valid bool, message string, err error) {
	return l.ValidateValue(stack, reflect.ValueOf(input))
}

func (l LessThan) ValidateValue(stack global.Stack, value reflect.Value) (valid bool, message string, err error) {
	return compare(stack, value, "<", json.Number(l))
}

// frizz
type GreaterThanOrEqual json.Number

func (g GreaterThanOrEqual) Validate(stack global.Stack, input interface{}) (valid bool, message string, err error) {
	return g.ValidateValue(stack, reflect.ValueOf(input))
}

func (g GreaterThanOrEqual) ValidateValue(stack global.Stack, value reflect.Value) (valid bool, message string, err error) {
	return compare(stack, value, ">=", json.Number(g))
}

// frizz
type LessThanOrEqual json.Number

func (l LessThanOrEqual) Validate(stack global.Stack, input interface{}) (valid bool, message string, err error) {
	return l.ValidateValue(stack, reflect.ValueOf(input))
}

func (l LessThanOrEqual) ValidateValue(stack global.Stack, value reflect.Value) (valid bool, message string, err error) {
	return compare(stack, value, "<=", json.Number(l))
}

// frizz
type Equal system.Raw

func (e Equal) Validate(stack global.Stack, input interface{}) (valid bool, message string, err error) {
	return e.ValidateValue(stack, reflect.ValueOf(input))
}

func (e Equal) ValidateValue(stack global.Stack, value reflect.Value) (valid bool, message string, err error) {
	if !value.IsValid() {
		// nil value -> return valid
		return true, "", nil
	}
	comparison := e.Value
	switch comparison := comparison.(type) {
	case json.Number:
		return compare(stack, value, "==", comparison)
	case string:
		if value.Type().Kind() != reflect.String {
			return false, "", errors.Errorf("%s: validators.Equal can only accept numeric or string types. Found %T", stack, value.Interface())
		}
		input := value.Convert(stringtype).Interface().(string)
		if input != comparison {
			return false, fmt.Sprintf("%s: value %#v must be equal to %#v", stack, input, comparison), nil
		}
		return true, "", nil
	}
	return false, "", errors.Errorf("%s: unsupported type %T for validatros.Equal value", stack, comparison)
}

func compare(stack global.Stack, value reflect.Value, operator string, comparison json.Number) (valid bool, message string, err error) {
	if !value.IsValid() {
		// nil value -> return valid
		return true, "", nil
	}
	switch value.Type().Kind() {
	case reflect.Interface, reflect.Ptr:
		// interface or ptr: recurse with elem
		return compare(stack, value.Elem(), operator, comparison)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		input := value.Convert(int64type).Interface().(int64)
		i, err := strconv.ParseInt(string(comparison), 10, 64)
		if err != nil {
			return false, "", errors.Wrapf(err, "%s: type %T can only be compared with int64, not %s", stack, value.Interface(), comparison)
		}
		if operator == ">" {
			if input <= i {
				return false, fmt.Sprintf("%s: value %v must be greater than %v", stack, input, i), nil
			}
		} else if operator == "<" {
			if input >= i {
				return false, fmt.Sprintf("%s: value %v must be less than %v", stack, input, i), nil
			}
		} else if operator == ">=" {
			if input < i {
				return false, fmt.Sprintf("%s: value %v must be greater than or equal to %v", stack, input, i), nil
			}
		} else if operator == "<=" {
			if input > i {
				return false, fmt.Sprintf("%s: value %v must be less than or equal to %v", stack, input, i), nil
			}
		} else if operator == "==" {
			if input != i {
				return false, fmt.Sprintf("%s: value %v must be equal to %v", stack, input, i), nil
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		input := value.Convert(uint64type).Interface().(uint64)
		i, err := strconv.ParseUint(string(comparison), 10, 64)
		if err != nil {
			return false, "", errors.Wrapf(err, "%s: type %T can only be compared with uint64, not %s", stack, value.Interface(), comparison)
		}
		if operator == ">" {
			if input <= i {
				return false, fmt.Sprintf("%s: value %v must be greater than %v", stack, input, i), nil
			}
		} else if operator == "<" {
			if input >= i {
				return false, fmt.Sprintf("%s: value %v must be less than %v", stack, input, i), nil
			}
		} else if operator == ">=" {
			if input < i {
				return false, fmt.Sprintf("%s: value %v must be greater than or equal to %v", stack, input, i), nil
			}
		} else if operator == "<=" {
			if input > i {
				return false, fmt.Sprintf("%s: value %v must be less than or equal to %v", stack, input, i), nil
			}
		} else if operator == "==" {
			if input != i {
				return false, fmt.Sprintf("%s: value %v must be equal to %v", stack, input, i), nil
			}
		}
	case reflect.Float32, reflect.Float64:
		input := value.Convert(float64type).Interface().(float64)
		f, err := strconv.ParseFloat(string(comparison), 64)
		if err != nil {
			return false, "", errors.Wrapf(err, "%s: type %T can only be compared with float64, not %s", stack, value.Interface(), comparison)
		}
		if operator == ">" {
			if input <= f {
				return false, fmt.Sprintf("%s: value %v must be greater than %v", stack, input, f), nil
			}
		} else if operator == "<" {
			if input >= f {
				return false, fmt.Sprintf("%s: value %v must be less than %v", stack, input, f), nil
			}
		} else if operator == ">=" {
			if input < f {
				return false, fmt.Sprintf("%s: value %v must be greater than or equal to %v", stack, input, f), nil
			}
		} else if operator == "<=" {
			if input > f {
				return false, fmt.Sprintf("%s: value %v must be less than or equal to %v", stack, input, f), nil
			}
		} else if operator == "==" {
			if input != f {
				return false, fmt.Sprintf("%s: value %v must be equal to %v", stack, input, f), nil
			}
		}
	default:
		return false, "", errors.Errorf("%s: unknown type for numeric comparison %T", stack, value.Interface())
	}
	return true, "", nil
}
