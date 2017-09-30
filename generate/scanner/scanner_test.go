package scanner

import (
	"testing"

	"github.com/dave/patsy/builder"
	"github.com/dave/patsy/vos"
)

func TestScan_ignores_generated_file(t *testing.T) {
	env := vos.Mock()
	b, err := builder.New(env, "a.b")
	if err != nil {
		t.Error(err)
	}
	packagePath, _, err := b.Package("c", map[string]string{
		"c.go":               "package c",
		"generated.frizz.go": "sadsda",
	})
	if err != nil {
		t.Error(err)
	}
	s := New(packagePath, env)
	err = s.Scan()
	if err != nil {
		t.Error(err)
	}
}

func TestScan_ignores_errors(t *testing.T) {
	env := vos.Mock()
	b, err := builder.New(env, "a.b")
	if err != nil {
		t.Error(err)
	}
	packagePath, _, err := b.Package("c", map[string]string{
		"c.go": "package c\n\nvar i int = d.Foo()",
	})
	if err != nil {
		t.Error(err)
	}
	s := New(packagePath, env)
	err = s.Scan()
	if err != nil {
		t.Error(err)
	}
}
