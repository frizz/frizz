package validators

import (
	"regexp"

	"fmt"

	"frizz.io/global"
	"github.com/pkg/errors"
)

// frizz
type Regex struct {
	Regex  string
	Invert bool
}

func (r Regex) Validate(stack global.Stack, input interface{}) (valid bool, message string, err error) {
	var s string
	switch input := input.(type) {
	case string:
		s = input
	case fmt.Stringer:
		s = input.String()
	default:
		return false, "", errors.Errorf("%s: validators.Regex can only validate string or fmt.Stringer, found %T", stack, input)
	}
	matched, err := regexp.Match(r.Regex, []byte(s))
	if err != nil {
		return false, "", errors.WithStack(err)
	}
	if !r.Invert && !matched {
		return false, fmt.Sprintf("%s: value %#v did not match regex %#v", stack, input, r.Regex), nil
	}
	if r.Invert && matched {
		return false, fmt.Sprintf("%s: value %#v matched regex %#v", stack, input, r.Regex), nil
	}
	return true, "", nil
}
