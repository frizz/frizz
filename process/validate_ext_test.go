package process_test

import (
	"testing"

	"kego.io/process"
	_ "kego.io/system/types"
)

func TestValidate(t *testing.T) {
	process.Validate_NeedsTypes(t)
}
