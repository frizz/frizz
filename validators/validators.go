package validators

import "frizz.io/common"

// frizz
type Keys struct {
	Validators []common.Validator
}

func (k Keys)Validate(input interface{}) (valid bool, message string, err error) {
	return true, "", nil
}

// frizz
type Items struct {
	Validators []common.Validator
}

func (i Items)Validate(input interface{}) (valid bool, message string, err error) {
	return true, "", nil
}

// frizz
type Regex struct {
	Regex string
}

func (r Regex)Validate(input interface{}) (valid bool, message string, err error) {
	return true, "", nil
}