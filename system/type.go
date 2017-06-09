package system

import (
	"frizz.io/common"
)

// frizz
type Type struct {
	Fields map[string]Field
}

// frizz
type Field struct {
	Validators []common.Validator
}
