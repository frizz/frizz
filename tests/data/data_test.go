package data

import (
	"fmt"
	"testing"

	"frizz.io/frizz"
)

func TestAddImports(t *testing.T) {
	obs, err := frizz.Package(Imports)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v", obs)

}
