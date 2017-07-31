package validation

import (
	frizz "frizz.io/frizz"
	system "frizz.io/system"
	validators "frizz.io/validators"
)

const Imports imports = 0

type imports int

func (i imports) Path() string {
	return "frizz.io/tests/validation"
}
func (i imports) Add(packers map[string]frizz.Packer, types map[string]frizz.Typer) {
	system.Imports.Add(packers, types)
	validators.Imports.Add(packers, types)
}
