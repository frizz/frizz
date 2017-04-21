package generate

import (
	"testing"

	"frizz.io/tests"
	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
)

func TestValidate(t *testing.T) {

	cb := tests.Context("a.b/c").Alias("f", "d.e/f")

	b, err := ValidateCommand(cb.Ctx())
	require.NoError(t, err)
	assert.Equal(t, `package main

import (
	_ "a.b/c"
	_ "d.e/f"
	command "frizz.io/process/validate/command"
	_ "frizz.io/system"
)

func main() {
	command.ValidateMain("a.b/c")
}
`, string(b))

	cb = tests.Context("a.b/c").Alias("f", "d.e/\"")

	b, err = ValidateCommand(cb.Ctx())
	assert.IsError(t, err, "IIIRBBXASR")

}
