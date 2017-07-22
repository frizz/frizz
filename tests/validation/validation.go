package validation

import (
	"frizz.io/frizz"
	"frizz.io/system"
	"frizz.io/validators"
)

var Packers = []frizz.Packer{Packer, system.Packer, validators.Packer}

// frizz
type Simple struct {
	String string
}
