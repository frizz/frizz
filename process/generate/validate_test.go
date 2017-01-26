package generate

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/tests"
)

func TestValidate(t *testing.T) {

	cb := tests.Context("a.b/c").Alias("f", "d.e/f")

	b, err := ValidateCommand(cb.Ctx())
	require.NoError(t, err)
	assert.Equal(t, `package main

import (
	_ "a.b/c"
	_ "d.e/f"
	command "kego.io/process/validate/command"
	_ "kego.io/system"
)

func main() {
	command.ValidateMain("a.b/c")
}
`, string(b))

	cb = tests.Context("a.b/c").Alias("f", "d.e/\"")

	b, err = ValidateCommand(cb.Ctx())
	assert.IsError(t, err, "IIIRBBXASR")

}
