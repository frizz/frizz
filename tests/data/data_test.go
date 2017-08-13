package data

import (
	"testing"

	"frizz.io/frizz"
	"frizz.io/tests/packer"
	"frizz.io/tests/packer/sub"
)

// TestImports checks that a data-only package (without type definitions) works properly.
func TestImports(t *testing.T) {
	obs, err := frizz.Package(Package)
	if err != nil {
		t.Fatal(err)
	}
	d, ok := obs["sub-interface"]
	if !ok {
		t.Fatal("Can't find sub-interface in package")
	}
	o, ok := d.(packer.SubInterface)
	if !ok {
		t.Fatalf("Wrong type %T for sub-interface", d)
	}
	s, ok := o.SubInterface.(sub.Sub)
	if !ok {
		t.Fatalf("Wrong type %T for o.SubInterface", o.SubInterface)
	}
	if s.String != "foo" {
		t.Fatalf("Expected Foo, got %s", s.String)
	}
}

func TestData(t *testing.T) {
	n := Data.Natives()
	if n.Int != 2 {
		t.Fail()
	}
}
