package system

import (
	"reflect"

	"frizz.io/common"
	"github.com/pkg/errors"
)

// frizz
type Type struct {
	Validators []common.Validator
	//Fields map[string][]common.Validator // Validators for struct fields
	//Keys   []common.Validator            // Validators for map keys
	//Items  []common.Validator            // Validators for map / array items
	//Value  []common.Validator            // Validators for native values
}

func (t Type) Validate(value interface{}) (valid bool, message string, err error) {
	return t.validate(reflect.ValueOf(value))
}

func (t Type) validate(value reflect.Value) (valid bool, message string, err error) {

	switch value.Kind() {
	case reflect.Invalid, reflect.Chan, reflect.Func, reflect.UnsafePointer:

		return false, "", errors.Errorf("can't validate kind %s", value.Kind())

	case reflect.Ptr, reflect.Interface:

		valid, message, err = t.Validate(value.Elem())
		if err != nil {
			return false, "", err
		}
		if !valid {
			return false, message, nil
		}

	default:

		for _, v := range t.Validators {
			if vv, ok := v.(common.ValueValidator); ok {
				valid, message, err = vv.ValidateValue(value)
				if err != nil {
					return false, "", err
				}
				if !valid {
					return false, message, nil
				}
			} else {
				valid, message, err = v.Validate(value.Interface())
				if err != nil {
					return false, "", err
				}
				if !valid {
					return false, message, nil
				}
			}
		}

		/*
			case reflect.Array, reflect.Map, reflect.Slice:

				// validate map keys
				if len(t.Keys) > 0 && value.Kind() == reflect.Map {
					for i := 0; i < value.Len(); i++ {
						keyValue := value.MapKeys()[i]
						for _, validator := range t.Items {
							valid, message, err = validator.Validate(keyValue.Interface())
							if err != nil {
								return false, "", err
							}
							if !valid {
								return false, message, nil
							}
						}
					}
				}

				// validate collection items
				if len(t.Items) > 0 {
					for i := 0; i < value.Len(); i++ {
						var itemValue reflect.Value
						if value.Kind() == reflect.Map {
							itemValue = value.MapIndex(value.MapKeys()[i])
						} else {
							itemValue = value.Index(i)
						}
						for _, validator := range t.Items {
							valid, message, err = validator.Validate(itemValue.Interface())
							if err != nil {
								return false, "", err
							}
							if !valid {
								return false, message, nil
							}
						}
					}
				}

			case reflect.Struct:

				if len(t.Fields) > 0 {
					for i := 0; i < value.Type().NumField(); i++ {
						f := value.Type().Field(i)
						validators, ok := t.Fields[f.Name]
						if !ok || len(validators) == 0 {
							continue
						}
						fieldValue := value.Field(i)
						for _, validator := range validators {
							valid, message, err = validator.Validate(fieldValue.Interface())
							if err != nil {
								return false, "", err
							}
							if !valid {
								return false, message, nil
							}
						}
					}
				}

			default:

				// native value
				if len(t.Value) > 0 {
					for _, validator := range t.Value {
						valid, message, err = validator.Validate(value.Interface())
						if err != nil {
							return false, "", err
						}
						if !valid {
							return false, message, nil
						}
					}
				}
		*/
	}
	return true, "", nil
}
