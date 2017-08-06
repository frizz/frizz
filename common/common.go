package common

import (
	"reflect"

	"frizz.io/global"
)

// frizz
type Validator interface {
	Validate(context global.ValidationContext, input interface{}) (valid bool, message string, err error)
}

type ValueValidator interface {
	ValidateValue(context global.ValidationContext, input reflect.Value) (valid bool, message string, err error)
}
