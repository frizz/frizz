package generate

import (
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/tests"
)

func TestValidate(t *testing.T) {

	cb := tests.Context("a.b/c").Alias("d.e/f", "f")

	b, err := ValidateCommand(cb.Ctx())
	assert.NoError(t, err)
	assert.Equal(t, `package main

import (
	_ "a.b/c"
	_ "d.e/f"
	"kego.io/process/validate/command"
	_ "kego.io/system"
)

func main() {
	command.Main("a.b/c")
}
`, string(b))

	cb = tests.Context("a.b/c").Alias("d.e/\"", "f")

	b, err = ValidateCommand(cb.Ctx())
	assert.IsError(t, err, "IIIRBBXASR")

}
