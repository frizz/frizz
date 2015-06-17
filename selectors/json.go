package selectors

import (
	"reflect"

	"kego.io/system"
)

type Element struct {
	Data  interface{}
	Value reflect.Value
	Rule  *system.RuleHolder
}
