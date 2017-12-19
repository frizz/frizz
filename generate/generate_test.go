package generate

import (
	"bytes"
	"testing"

	"github.com/dave/patsy/vos"
)

/*
func TestGenerate_selector(t *testing.T) {
	env := vos.Mock()

	b, err := builder.New(env, "ns")
	if err != nil {
		t.Fatal(err)
	}
	defer b.Cleanup()

	path, _, err := b.Package("a", map[string]string{
		"a.go": `
package a

import (
	"ns/b"
	"ns/c"
	"time"
)

// frizz
type A struct {
	B b.B
	C c.C
	T time.Time
}
`,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = b.Package("b", map[string]string{
		"b.go": `
package b

// frizz
type B string
`,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = b.Package("c", map[string]string{
		"c.go": `
package c

type C string
`,
	})
	if err != nil {
		t.Fatal(err)
	}

	scan, err := Scan(env, path)
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range scan.Packages {
		fmt.Println("=========================", p.Path)
		buf := &bytes.Buffer{}
		if _, err := p.Generate(buf); err != nil {
			t.Fatal(err)
		}
		fmt.Println(buf.String())
	}
}
*/

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
