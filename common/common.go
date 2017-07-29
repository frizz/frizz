package common

import "reflect"

// frizz
type Validator interface {
	Validate(input interface{}) (valid bool, message string, err error)
}

type ValueValidator interface {
	ValidateValue(input reflect.Value) (valid bool, message string, err error)
}
