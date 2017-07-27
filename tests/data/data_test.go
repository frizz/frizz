package data

import (
	"testing"
	"frizz.io/frizz"
	"fmt"
)

func TestAddImports(t *testing.T) {
	packers := map[string]frizz.Packer{}
	AddImports(packers, nil)
	var pkrs []frizz.Packer
	for _, p := range packers {
		pkrs = append(pkrs, p)
	}

	obs, err := frizz.Package("frizz.io/tests/data", pkrs...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v", obs)

}

