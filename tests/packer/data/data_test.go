package data

import (
	"testing"

	"frizz.io/frizz"
	"frizz.io/tests/packer"
	"frizz.io/tests/packer/sub"
)

// TestImports checks that a data-only package (without type definitions) works properly.
func TestImports(t *testing.T) {
	obs, err := frizz.Package(Imports)
	if err != nil {
		t.Fatal(err)
	}
	if len(obs) != 1 {
		t.Fail()
	}
	d, ok := obs["sub-interface"]
	if !ok {
		t.Fail()
	}
	o, ok := d.(packer.SubInterface)
	if !ok {
		t.Fail()
	}
	s, ok := o.SubInterface.(sub.Sub)
	if !ok {
		t.Fail()
	}
	if s.String != "foo" {
		t.Fail()
	}
}
