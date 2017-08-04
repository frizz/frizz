package common

import (
	"reflect"

	"frizz.io/global"
)

// frizz
type Validator interface {
	Validate(stack global.Stack, input interface{}) (valid bool, message string, err error)
}

type ValueValidator interface {
	ValidateValue(stack global.Stack, input reflect.Value) (valid bool, message string, err error)
}
