package validators

import (
	"encoding/json"
	"fmt"
	"strconv"

	"reflect"

	"frizz.io/system"
	"github.com/pkg/errors"
)

// frizz
type GreaterThan json.Number

func (g GreaterThan) Validate(input interface{}) (valid bool, message string, err error) {
	return g.ValidateValue(reflect.ValueOf(input))
}

func (g GreaterThan) ValidateValue(value reflect.Value) (valid bool, message string, err error) {
	return compare(value, ">", json.Number(g))
}

// frizz
type LessThan json.Number

func (l LessThan) Validate(input interface{}) (valid bool, message string, err error) {
	return l.ValidateValue(reflect.ValueOf(input))
}

func (l LessThan) ValidateValue(value reflect.Value) (valid bool, message string, err error) {
	return compare(value, "<", json.Number(l))
}

// frizz
type GreaterThanOrEqual json.Number

func (g GreaterThanOrEqual) Validate(input interface{}) (valid bool, message string, err error) {
	return g.ValidateValue(reflect.ValueOf(input))
}

func (g GreaterThanOrEqual) ValidateValue(value reflect.Value) (valid bool, message string, err error) {
	return compare(value, ">=", json.Number(g))
}

// frizz
type LessThanOrEqual json.Number

func (l LessThanOrEqual) Validate(input interface{}) (valid bool, message string, err error) {
	return l.ValidateValue(reflect.ValueOf(input))
}

func (l LessThanOrEqual) ValidateValue(value reflect.Value) (valid bool, message string, err error) {
	return compare(value, "<=", json.Number(l))
}

// frizz
type Equal system.Raw

func (e Equal) Validate(input interface{}) (valid bool, message string, err error) {
	return e.ValidateValue(reflect.ValueOf(input))
}

func (e Equal) ValidateValue(value reflect.Value) (valid bool, message string, err error) {
	comparison := e.Value
	switch comparison := comparison.(type) {
	case json.Number:
		return compare(value, "==", comparison)
	case string:
		if value.Type().Kind() != reflect.String {
			return false, "", errors.Errorf("validators.Equal can only accept numeric or string types. Found %T", value.Interface())
		}
		input := value.Convert(stringtype).Interface().(string)
		if input != comparison {
			return false, fmt.Sprintf("value %#v must be equal to %#v", input, comparison), nil
		}
		return true, "", nil
	}
	return false, "", errors.Errorf("unsupported type %T for validatros.Equal value", comparison)
}

func compare(value reflect.Value, operator string, comparison json.Number) (valid bool, message string, err error) {
	switch value.Type().Kind() {
	case reflect.Interface, reflect.Ptr:
		// interface or ptr: recurse with elem
		return compare(value.Elem(), operator, comparison)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		input := value.Convert(int64type).Interface().(int64)
		i, err := strconv.ParseInt(string(comparison), 10, 64)
		if err != nil {
			return false, "", errors.Wrapf(err, "type %T can only be compared with int64, not %s", value.Interface(), comparison)
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
		} else if operator == "==" {
			if input != i {
				return false, fmt.Sprintf("value %v must be equal to %v", input, i), nil
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		input := value.Convert(uint64type).Interface().(uint64)
		i, err := strconv.ParseUint(string(comparison), 10, 64)
		if err != nil {
			return false, "", errors.Wrapf(err, "type %T can only be compared with uint64, not %s", value.Interface(), comparison)
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
		} else if operator == "==" {
			if input != i {
				return false, fmt.Sprintf("value %v must be equal to %v", input, i), nil
			}
		}
	case reflect.Float32, reflect.Float64:
		input := value.Convert(float64type).Interface().(float64)
		f, err := strconv.ParseFloat(string(comparison), 64)
		if err != nil {
			return false, "", errors.Wrapf(err, "type %T can only be compared with float64, not %s", value.Interface(), comparison)
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
		} else if operator == "==" {
			if input != f {
				return false, fmt.Sprintf("value %v must be equal to %v", input, f), nil
			}
		}
	default:
		return false, "", errors.Errorf("unknown type for numeric comparison %T", value.Interface())
	}
	return true, "", nil
}
