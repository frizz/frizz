package generator_test

import (
	"testing"

	"kego.io/generator"
)

func TestGoTypeDescriptorErrors(t *testing.T) {
	generator.TypeErrors_NeedsTypes(t)
}
