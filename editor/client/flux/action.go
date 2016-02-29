package flux

import "reflect"

type ActionInterface interface {
}

type Action struct {
	waiting map[reflect.Type]bool
}
