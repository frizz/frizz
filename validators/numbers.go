package validators

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

// frizz
type GreaterThan json.Number

func (g GreaterThan) Validate(input interface{}) (valid bool, message string, err error) {
	return validate(input, ">", json.Number(g))
}

// frizz
type LessThan json.Number

func (l LessThan) Validate(input interface{}) (valid bool, message string, err error) {
	return validate(input, "<", json.Number(l))
}

// frizz
type GreaterThanOrEqual json.Number

func (g GreaterThanOrEqual) Validate(input interface{}) (valid bool, message string, err error) {
	return validate(input, ">=", json.Number(g))
}

// frizz
type LessThanOrEqual json.Number

func (l LessThanOrEqual) Validate(input interface{}) (valid bool, message string, err error) {
	return validate(input, "<=", json.Number(l))
}

func validate(input interface{}, operator string, comparison json.Number) (valid bool, message string, err error) {
	switch input.(type) {
	case int, int8, int16, int32, int64:
		input := MustInt64(input)
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
	case uint, uint8, uint16, uint32, uint64:
		input := MustUint64(input)
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
	case float32, float64:
		input := MustFloat64(input)
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
	}
	return true, "", nil
}
