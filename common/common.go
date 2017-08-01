package common

import (
	"reflect"

	"frizz.io/frizz"
)

// frizz
type Validator interface {
	Validate(stack frizz.Stack, input interface{}) (valid bool, message string, err error)
}

type ValueValidator interface {
	ValidateValue(stack frizz.Stack, input reflect.Value) (valid bool, message string, err error)
}
