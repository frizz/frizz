package generate

import (
	"bytes"
	"testing"

	"github.com/dave/patsy/vos"
)

func TestGenerate(t *testing.T) {
	env := vos.Os()
	path := "frizz.io/tests/packer"
	scan, err := Scan(env, path)
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range scan.Packages {
		buf := &bytes.Buffer{}
		if _, err := p.Generate(buf); err != nil {
			t.Fatal(err.Error())
		}
	}
}
