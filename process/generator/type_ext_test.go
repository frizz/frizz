package generator_test

import (
	"testing"

	"kego.io/process/generator"
	_ "kego.io/system/types"
)

func TestGoTypeDescriptorErrors(t *testing.T) {
	generator.TypeErrors_NeedsTypes(t)
}
