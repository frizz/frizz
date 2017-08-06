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

func (g GreaterThan) Validate(context global.ValidationContext, input interface{}) (valid bool, message string, err error) {
	return g.ValidateValue(context, reflect.ValueOf(input))
}

func (g GreaterThan) ValidateValue(context global.ValidationContext, value reflect.Value) (valid bool, message string, err error) {
	return compare(context, value, ">", json.Number(g))
}

// frizz
type LessThan json.Number

func (l LessThan) Validate(context global.ValidationContext, input interface{}) (valid bool, message string, err error) {
	return l.ValidateValue(context, reflect.ValueOf(input))
}

func (l LessThan) ValidateValue(context global.ValidationContext, value reflect.Value) (valid bool, message string, err error) {
	return compare(context, value, "<", json.Number(l))
}

// frizz
type GreaterThanOrEqual json.Number

func (g GreaterThanOrEqual) Validate(context global.ValidationContext, input interface{}) (valid bool, message string, err error) {
	return g.ValidateValue(context, reflect.ValueOf(input))
}

func (g GreaterThanOrEqual) ValidateValue(context global.ValidationContext, value reflect.Value) (valid bool, message string, err error) {
	return compare(context, value, ">=", json.Number(g))
}

// frizz
type LessThanOrEqual json.Number

func (l LessThanOrEqual) Validate(context global.ValidationContext, input interface{}) (valid bool, message string, err error) {
	return l.ValidateValue(context, reflect.ValueOf(input))
}

func (l LessThanOrEqual) ValidateValue(context global.ValidationContext, value reflect.Value) (valid bool, message string, err error) {
	return compare(context, value, "<=", json.Number(l))
}

// frizz
type Equal system.Raw

func (e Equal) Validate(context global.ValidationContext, input interface{}) (valid bool, message string, err error) {
	return e.ValidateValue(context, reflect.ValueOf(input))
}

func (e Equal) ValidateValue(context global.ValidationContext, value reflect.Value) (valid bool, message string, err error) {
	if !value.IsValid() {
		// nil value -> return valid
		return true, "", nil
	}
	comparison := e.Value
	switch comparison := comparison.(type) {
	case json.Number:
		return compare(context, value, "==", comparison)
	case string:
		if value.Type().Kind() != reflect.String {
			return false, "", errors.Errorf("%s: validators.Equal can only accept numeric or string types. Found %T", context.Location(), value.Interface())
		}
		input := value.Convert(stringtype).Interface().(string)
		if input != comparison {
			return false, fmt.Sprintf("%s: value %#v must be equal to %#v", context.Location(), input, comparison), nil
		}
		return true, "", nil
	}
	return false, "", errors.Errorf("%s: unsupported type %T for validatros.Equal value", context.Location(), comparison)
}

func compare(context global.ValidationContext, value reflect.Value, operator string, comparison json.Number) (valid bool, message string, err error) {
	if !value.IsValid() {
		// nil value -> return valid
		return true, "", nil
	}
	switch value.Type().Kind() {
	case reflect.Interface, reflect.Ptr:
		// interface or ptr: recurse with elem
		return compare(context, value.Elem(), operator, comparison)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		input := value.Convert(int64type).Interface().(int64)
		i, err := strconv.ParseInt(string(comparison), 10, 64)
		if err != nil {
			return false, "", errors.Wrapf(err, "%s: type %T can only be compared with int64, not %s", context.Location(), value.Interface(), comparison)
		}
		if operator == ">" {
			if input <= i {
				return false, fmt.Sprintf("%s: value %v must be greater than %v", context.Location(), input, i), nil
			}
		} else if operator == "<" {
			if input >= i {
				return false, fmt.Sprintf("%s: value %v must be less than %v", context.Location(), input, i), nil
			}
		} else if operator == ">=" {
			if input < i {
				return false, fmt.Sprintf("%s: value %v must be greater than or equal to %v", context.Location(), input, i), nil
			}
		} else if operator == "<=" {
			if input > i {
				return false, fmt.Sprintf("%s: value %v must be less than or equal to %v", context.Location(), input, i), nil
			}
		} else if operator == "==" {
			if input != i {
				return false, fmt.Sprintf("%s: value %v must be equal to %v", context.Location(), input, i), nil
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		input := value.Convert(uint64type).Interface().(uint64)
		i, err := strconv.ParseUint(string(comparison), 10, 64)
		if err != nil {
			return false, "", errors.Wrapf(err, "%s: type %T can only be compared with uint64, not %s", context.Location(), value.Interface(), comparison)
		}
		if operator == ">" {
			if input <= i {
				return false, fmt.Sprintf("%s: value %v must be greater than %v", context.Location(), input, i), nil
			}
		} else if operator == "<" {
			if input >= i {
				return false, fmt.Sprintf("%s: value %v must be less than %v", context.Location(), input, i), nil
			}
		} else if operator == ">=" {
			if input < i {
				return false, fmt.Sprintf("%s: value %v must be greater than or equal to %v", context.Location(), input, i), nil
			}
		} else if operator == "<=" {
			if input > i {
				return false, fmt.Sprintf("%s: value %v must be less than or equal to %v", context.Location(), input, i), nil
			}
		} else if operator == "==" {
			if input != i {
				return false, fmt.Sprintf("%s: value %v must be equal to %v", context.Location(), input, i), nil
			}
		}
	case reflect.Float32, reflect.Float64:
		input := value.Convert(float64type).Interface().(float64)
		f, err := strconv.ParseFloat(string(comparison), 64)
		if err != nil {
			return false, "", errors.Wrapf(err, "%s: type %T can only be compared with float64, not %s", context.Location(), value.Interface(), comparison)
		}
		if operator == ">" {
			if input <= f {
				return false, fmt.Sprintf("%s: value %v must be greater than %v", context.Location(), input, f), nil
			}
		} else if operator == "<" {
			if input >= f {
				return false, fmt.Sprintf("%s: value %v must be less than %v", context.Location(), input, f), nil
			}
		} else if operator == ">=" {
			if input < f {
				return false, fmt.Sprintf("%s: value %v must be greater than or equal to %v", context.Location(), input, f), nil
			}
		} else if operator == "<=" {
			if input > f {
				return false, fmt.Sprintf("%s: value %v must be less than or equal to %v", context.Location(), input, f), nil
			}
		} else if operator == "==" {
			if input != f {
				return false, fmt.Sprintf("%s: value %v must be equal to %v", context.Location(), input, f), nil
			}
		}
	default:
		return false, "", errors.Errorf("%s: unknown type for numeric comparison %T", context.Location(), value.Interface())
	}
	return true, "", nil
}
