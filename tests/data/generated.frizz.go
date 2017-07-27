package data

import (
	frizz "frizz.io/frizz"
	packer "frizz.io/tests/packer"
	sub "frizz.io/tests/packer/sub"
)

func AddImports(packers map[string]frizz.Packer, types map[string]frizz.Typer) {
	packer.AddImports(packers, types)
	sub.AddImports(packers, types)
}
