package validators

import "frizz.io/system/common"

type Keys struct {
	Validators []common.Validator
}

type Items struct {
	Validators []common.Validator
}

type Regex struct {
	Regex string
}
