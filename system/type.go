package system

import "frizz.io/system/common"

type Type struct {
	Fields map[string]Field
}

type Field struct {
	Validators []common.Validator
}
