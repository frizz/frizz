package builder_test

import (
	"testing"

	"kego.io/process/generate/builder"
)

func TestGoTypeDescriptorErrors(t *testing.T) {
	builder.TypeErrors_NeedsTypes(t)
}
