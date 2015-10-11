package validate_test

import (
	"testing"

	"kego.io/process/validate"
	_ "kego.io/system/types"
)

func TestValidate(t *testing.T) {
	validate.Validate_NeedsTypes(t)
}
