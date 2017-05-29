package generate

import (
	"bytes"
	"testing"
)

func TestGenerate(t *testing.T) {
	buf := &bytes.Buffer{}
	if err := Generate(buf, "frizz.io/site"); err != nil {
		t.Fatal(err.Error())
	}
}
