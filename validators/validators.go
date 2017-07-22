package validators

import (
	"regexp"

	"fmt"

	"frizz.io/common"
	"github.com/pkg/errors"
)

// frizz
type Keys struct {
	Validators []common.Validator
}

func (k Keys) Validate(input interface{}) (valid bool, message string, err error) {
	return true, "", nil
}

// frizz
type Items struct {
	Validators []common.Validator
}

func (i Items) Validate(input interface{}) (valid bool, message string, err error) {
	return true, "", nil
}

// frizz
type Regex struct {
	Regex  string
	Invert bool
}

func (r Regex) Validate(input interface{}) (valid bool, message string, err error) {
	var s string
	switch input := input.(type) {
	case string:
		s = input
	case fmt.Stringer:
		s = input.String()
	default:
		return false, "", errors.Errorf("input to Regex validator should be string or fmt.Stringer, got %T.", input)
	}
	matched, err := regexp.Match(r.Regex, []byte(s))
	if err != nil {
		return false, "", errors.WithStack(err)
	}
	if !r.Invert && !matched {
		return false, fmt.Sprintf("input %#v did not match regex %#v", input, r.Regex), nil
	}
	if r.Invert && matched {
		return false, fmt.Sprintf("input %#v matched regex %#v", input, r.Regex), nil
	}
	return true, "", nil
}
