package process // import "kego.io/process"

import (
	"fmt"

	"kego.io/kerr"
	"kego.io/process/validate"
)

func FormatError(err error) string {
	source := kerr.Source(err)
	if m, ok := source.(validate.ValidationError); ok {
		return fmt.Sprint("Error: ", m.Description)
	}
	if t, ok := err.(CommandError); ok {
		return t.Description
	}
	return err.Error()
}
