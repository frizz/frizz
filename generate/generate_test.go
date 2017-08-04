package generate

import (
	"bytes"
	"testing"

	"github.com/dave/patsy"
	"github.com/dave/patsy/vos"
)

func TestGenerate(t *testing.T) {
	buf := &bytes.Buffer{}
	env := vos.Os()
	path := "frizz.io/tests/packer"
	dir, err := patsy.Dir(env, path)
	if err != nil {
		t.Fatal(err.Error())
	}
	if _, err := Generate(buf, env, path, dir); err != nil {
		t.Fatal(err.Error())
	}
}
