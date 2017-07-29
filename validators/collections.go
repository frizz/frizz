package validators

import (
	"reflect"

	"frizz.io/common"
	"github.com/pkg/errors"
)

// frizz
type Struct map[string][]common.Validator

func (f Struct) Validate(input interface{}) (valid bool, message string, err error) {
	return f.ValidateValue(reflect.ValueOf(input))
}

func (f Struct) ValidateValue(value reflect.Value) (valid bool, message string, err error) {
	if value.Kind() != reflect.Struct {
		return false, "", errors.Errorf("validator Fields can only validate structs. Got %T", value.Interface())
	}
	for name, validators := range f {
		field := value.FieldByName(name)
		if field == (reflect.Value{}) {
			return false, "", errors.Errorf("field %s not found in %T", name, value.Interface())
		}
		for _, v := range validators {
			valid, message, err = v.Validate(field.Interface())
			if err != nil {
				return false, "", err
			}
			if !valid {
				return false, message, nil
			}
		}
	}
	return true, "", nil
}

// frizz
type Keys []common.Validator

func (k Keys) Validate(input interface{}) (valid bool, message string, err error) {
	return k.ValidateValue(reflect.ValueOf(input))
}

func (k Keys) ValidateValue(value reflect.Value) (valid bool, message string, err error) {
	if value.Kind() != reflect.Map {
		return false, "", errors.Errorf("validator Keys can only validate maps. Got %T", value.Interface())
	}
	for i := 0; i < value.Len(); i++ {
		keyValue := value.MapKeys()[i]
		for _, validator := range k {
			valid, message, err = validator.Validate(keyValue.Interface())
			if err != nil {
				return false, "", err
			}
			if !valid {
				return false, message, nil
			}
		}
	}
	return true, "", nil
}

// frizz
type Items []common.Validator

func (i Items) Validate(input interface{}) (valid bool, message string, err error) {
	return i.ValidateValue(reflect.ValueOf(input))
}

func (i Items) ValidateValue(value reflect.Value) (valid bool, message string, err error) {
	if value.Kind() != reflect.Array && value.Kind() != reflect.Map && value.Kind() != reflect.Slice {
		return false, "", errors.Errorf("validator Items can only validate array, slice or map. Found %T", value.Interface())
	}
	for j := 0; j < value.Len(); j++ {
		var itemValue reflect.Value
		if value.Kind() == reflect.Map {
			itemValue = value.MapIndex(value.MapKeys()[j])
		} else {
			itemValue = value.Index(j)
		}
		for _, validator := range i {
			valid, message, err = validator.Validate(itemValue.Interface())
			if err != nil {
				return false, "", err
			}
			if !valid {
				return false, message, nil
			}
		}
	}
	return true, "", nil
}
