package validators

import (
	"reflect"

	"frizz.io/common"
	"frizz.io/frizz"
	"github.com/pkg/errors"
)

// frizz
type Struct map[string][]common.Validator

func (f Struct) Validate(stack frizz.Stack, input interface{}) (valid bool, message string, err error) {
	return f.ValidateValue(stack, reflect.ValueOf(input))
}

func (f Struct) ValidateValue(stack frizz.Stack, value reflect.Value) (valid bool, message string, err error) {
	if value.Kind() != reflect.Struct {
		return false, "", errors.Errorf("%s: validator Fields can only validate structs. Found %T", stack, value.Interface())
	}
	for name, validators := range f {
		field := value.FieldByName(name)
		if field == (reflect.Value{}) {
			return false, "", errors.Errorf("%s: field %s not found in %T", stack, name, value.Interface())
		}
		inner := stack.Append(frizz.FieldItem(name))
		for _, v := range validators {
			if vv, ok := v.(common.ValueValidator); ok {
				if valid, message, err := vv.ValidateValue(inner, field); err != nil || !valid {
					return valid, message, err
				}
			} else if valid, message, err := v.Validate(inner, field.Interface()); err != nil || !valid {
				return valid, message, err
			}
		}
	}
	return true, "", nil
}

// frizz
type Length []common.Validator

func (l Length) Validate(stack frizz.Stack, input interface{}) (valid bool, message string, err error) {
	return l.ValidateValue(stack, reflect.ValueOf(input))
}

func (l Length) ValidateValue(stack frizz.Stack, value reflect.Value) (valid bool, message string, err error) {
	switch value.Kind() {
	case reflect.Map, reflect.Array, reflect.Slice, reflect.String: // nothing
	default:
		return false, "", errors.Errorf("%s: validator Length can only validate arrays, slices, maps or strings. Found %T", stack, value.Interface())
	}
	length := value.Len()
	for _, validator := range l {
		if valid, message, err := validator.Validate(stack, length); err != nil || !valid {
			return valid, message, err
		}
	}
	return true, "", nil
}

// frizz
type Keys []common.Validator

func (k Keys) Validate(stack frizz.Stack, input interface{}) (valid bool, message string, err error) {
	return k.ValidateValue(stack, reflect.ValueOf(input))
}

func (k Keys) ValidateValue(stack frizz.Stack, value reflect.Value) (valid bool, message string, err error) {
	if value.Kind() != reflect.Map {
		return false, "", errors.Errorf("%s: validator Keys can only validate maps. Found %T", stack, value.Interface())
	}
	for i := 0; i < value.Len(); i++ {
		keyValue := value.MapKeys()[i]
		inner := stack.Append(frizz.MapItem(keyValue.Interface().(string)))
		for _, validator := range k {
			if valid, message, err := validator.Validate(inner, keyValue.Interface()); err != nil || !valid {
				return valid, message, err
			}
		}
	}
	return true, "", nil
}

// frizz
type Items []common.Validator

func (i Items) Validate(stack frizz.Stack, input interface{}) (valid bool, message string, err error) {
	return i.ValidateValue(stack, reflect.ValueOf(input))
}

func (i Items) ValidateValue(stack frizz.Stack, value reflect.Value) (valid bool, message string, err error) {
	switch value.Kind() {
	case reflect.Map, reflect.Array, reflect.Slice: // nothing
	default:
		return false, "", errors.Errorf("%s: validator Items can only validate arrays, slices or maps. Found %T", stack, value.Interface())
	}
	for j := 0; j < value.Len(); j++ {
		var inner frizz.Stack
		var itemValue reflect.Value
		if value.Kind() == reflect.Map {
			key := value.MapKeys()[j]
			itemValue = value.MapIndex(key)
			inner = stack.Append(frizz.MapItem(key.Interface().(string)))
		} else {
			itemValue = value.Index(j)
			inner = stack.Append(frizz.ArrayItem(j))
		}
		for _, validator := range i {
			if valid, message, err := validator.Validate(inner, itemValue.Interface()); err != nil || !valid {
				return valid, message, err
			}
		}
	}
	return true, "", nil
}
