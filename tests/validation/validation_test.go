package validation_test

import (
	"fmt"
	"testing"

	"frizz.io/frizz"
	"frizz.io/tests/validation"
)

func TestValidate(t *testing.T) {
	p, err := frizz.Package("frizz.io/tests/validation", validation.Packers...)
	if err != nil {
		t.Fatal(err.Error())
	}
	for k, v := range p {
		fmt.Printf("%#v %#v\n", k, v)
	}

}
