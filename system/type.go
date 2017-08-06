package system

import (
	"reflect"

	"frizz.io/common"
	"frizz.io/global"
	"github.com/pkg/errors"
)

// frizz
type Type struct {
	Validators []common.Validator
}

func (t Type) Validate(context global.ValidationContext, value interface{}) (valid bool, message string, err error) {
	return t.ValidateValue(context, reflect.ValueOf(value))
}

func (t Type) ValidateValue(context global.ValidationContext, value reflect.Value) (valid bool, message string, err error) {
	switch value.Kind() {
	case reflect.Invalid, reflect.Chan, reflect.Func, reflect.UnsafePointer:
		return false, "", errors.Errorf("%s: can't validate kind %s", context.Location(), value.Kind())
	case reflect.Ptr, reflect.Interface:
		if valid, message, err := t.ValidateValue(context, value.Elem()); err != nil || !valid {
			return valid, message, err
		}
	default:
		for _, v := range t.Validators {
			if vv, ok := v.(common.ValueValidator); ok {
				if valid, message, err := vv.ValidateValue(context, value); err != nil || !valid {
					return valid, message, err
				}
			} else {
				if valid, message, err := v.Validate(context, value.Interface()); err != nil || !valid {
					return valid, message, err
				}
			}
		}
	}
	return true, "", nil
}
