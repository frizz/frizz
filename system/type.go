package system

import (
	"reflect"

	"frizz.io/common"
	"frizz.io/frizz"
	"github.com/pkg/errors"
)

// frizz
type Type struct {
	Validators []common.Validator
}

func (t Type) Validate(stack frizz.Stack, value interface{}) (valid bool, message string, err error) {
	return t.ValidateValue(stack, reflect.ValueOf(value))
}

func (t Type) ValidateValue(stack frizz.Stack, value reflect.Value) (valid bool, message string, err error) {
	switch value.Kind() {
	case reflect.Invalid, reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return false, "", errors.Errorf("%s: can't validate kind %s", stack, value.Kind())
	case reflect.Ptr, reflect.Interface:
		valid, message, err = t.Validate(stack, value.Elem())
		if err != nil {
			return false, "", err
		}
		if !valid {
			return false, message, nil
		}
	default:
		for _, v := range t.Validators {
			if vv, ok := v.(common.ValueValidator); ok {
				valid, message, err = vv.ValidateValue(stack, value)
				if err != nil {
					return false, "", err
				}
				if !valid {
					return false, message, nil
				}
			} else {
				valid, message, err = v.Validate(stack, value.Interface())
				if err != nil {
					return false, "", err
				}
				if !valid {
					return false, message, nil
				}
			}
		}
	}
	return true, "", nil
}
