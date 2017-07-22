package system

import (
	"frizz.io/common"
)

// frizz
type Type struct {
	For    string
	Fields map[string]Field
}

// frizz
type Field struct {
	Validators []common.Validator
}
