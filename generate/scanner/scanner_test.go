package scanner

import (
	"testing"

	"fmt"
	"go/ast"

	"reflect"

	"frizz.io/generate/jast"
	"github.com/dave/jennifer/jen"
	"github.com/dave/patsy/builder"
	"github.com/dave/patsy/vos"
)

func TestSilentImport(t *testing.T) {
	env := vos.Mock()
	b, err := builder.New(env, "a.b")
	if err != nil {
		t.Fatal(err)
	}
	_, _, err = b.Package("silent", map[string]string{
		"silent.go": `
package silent

// frizz
type Silent interface {
	Silent()
}
`,
	})
	if err != nil {
		t.Fatal(err)
	}
	packagePathUse, _, err := b.Package("use", map[string]string{
		"use.go": `
package use

import (
	"time"

	"a.b/silent"
)

// frizz
type Uses silent.Silent

// frizz
type Time time.Time

`,
	})
	if err != nil {
		t.Fatal(err)
	}
	s := New(packagePathUse, env)
	err = s.Scan()
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(s.Imports, []string{"a.b/silent"}) {
		t.Fatal("Unexpected:", s.Imports)
	}

}

func TestJast(t *testing.T) {
	env := vos.Mock()
	b, err := builder.New(env, "a.b")
	if err != nil {
		t.Fatal(err)
	}
	packagePath, _, err := b.Package("c", map[string]string{
		"c.go": `
package c

import ggg "frizz.io/global"

// frizz
type Foo struct {
	i ggg.Packable
}
`,
	})
	if err != nil {
		t.Fatal(err)
	}
	s := New(packagePath, env)
	err = s.Scan()
	if err != nil {
		t.Fatal(err)
	}

	// find the SelectorExpr corresponding to "ggg.Packable"
	se := s.Types["Foo"].Spec.(*ast.StructType).Fields.List[0].Type.(*ast.SelectorExpr)

	j := jast.New(nil, s.Info.Uses)
	jf := jen.NewFile("a")
	jf.Func().Id("bar").Params().Block(
		jen.Var().Id("g").Add(j.Expr(se)),
	)

	expected := `package a

import global "frizz.io/global"

func bar() {
	var g global.Packable
}

`
	output := fmt.Sprintf("%#v\n", jf)

	if expected != output {
		t.Fatal("Unexpected", output)
	}

}

func TestScanner_Scan(t *testing.T) {
	env := vos.Mock()
	b, err := builder.New(env, "a.b")
	if err != nil {
		t.Fatal(err)
	}
	packagePath, _, err := b.Package("c", map[string]string{
		"c.go": `package c

import "frizz.io/global"

// frizz
type Foo struct {
	i int
}
func (f *Foo) Unpack(context global.DataContext, in interface{}) (null bool, err error) {}
func (f *Foo) Repack(context global.DataContext) (value interface{}, dict bool, null bool, err error) {}
`,
		"d.go": `package c`,
		"c_test.go": `
package c_test

import "testing"

func TestFoo(t *testing.T) {}`,
	})
	if err != nil {
		t.Fatal(err)
	}
	s := New(packagePath, env)
	err = s.Scan()
	if err != nil {
		t.Fatal(err)
	}
}

func TestScan_ignores_generated_file(t *testing.T) {
	env := vos.Mock()
	b, err := builder.New(env, "a.b")
	if err != nil {
		t.Fatal(err)
	}
	packagePath, _, err := b.Package("c", map[string]string{
		"c.go":               "package c",
		"generated.frizz.go": "sadsda",
	})
	if err != nil {
		t.Fatal(err)
	}
	s := New(packagePath, env)
	err = s.Scan()
	if err != nil {
		t.Fatal(err)
	}
}

func TestScan_ignores_errors(t *testing.T) {
	env := vos.Mock()
	b, err := builder.New(env, "a.b")
	if err != nil {
		t.Fatal(err)
	}
	packagePath, _, err := b.Package("c", map[string]string{
		"c.go": "package c\n\nvar i int = d.Foo()",
	})
	if err != nil {
		t.Fatal(err)
	}
	s := New(packagePath, env)
	err = s.Scan()
	if err != nil {
		t.Fatal(err)
	}
}
