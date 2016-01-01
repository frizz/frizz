package validate_test

import (
	"testing"

	"kego.io/process/validate"
)

func TestValidate(t *testing.T) {
	validate.Validate_NeedsTypes(t)
}
