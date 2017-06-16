package unpacker

import (
	"fmt"
	"go/ast"
	"go/parser"
	"strconv"

	"github.com/pkg/errors"

	"strings"

	"bytes"
	"go/printer"

	"frizz.io/frizz"
	"frizz.io/tests/unpacker/sub"
)

// frizz
type CustomSub sub.Sub

func (c *CustomSub) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}) error {
	s, err := sub.Unpackers.Sub(root, stack, in)
	if err != nil {
		return err
	}
	s.String += "-b"
	*c = CustomSub(s)
	return nil
}

func (c *CustomSub) Repack(root *frizz.Root, stack frizz.Stack) (interface{}, error) {
	return nil, nil
}

// frizz
type Ages map[string]int

func (a *Ages) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}) error {
	str, ok := in.(string)
	if !ok {
		return errors.Errorf("%s: ages type must be string", stack)
	}
	names := strings.Split(str, ",")
	ages := Ages{}
	for _, name := range names {
		parts := strings.Split(name, ":")
		if len(parts) != 2 {
			return errors.Errorf("%s: converting %s to ages", stack, str)
		}
		i, err := strconv.ParseInt(parts[1], 10, 0)
		if err != nil {
			return errors.Wrapf(err, "%s: converting %s to int", stack, parts[1])
		}
		ages[parts[0]] = int(i)
	}
	*a = ages
	return nil
}

func (a *Ages) Repack(root *frizz.Root, stack frizz.Stack) (interface{}, error) {
	var parts []string
	for k, v := range *a {
		parts = append(parts, fmt.Sprintf("%s:%s", k, v))
	}
	return strings.Join(parts, ","), nil
}

// frizz
type Csv []int

func (c *Csv) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}) error {
	str, ok := in.(string)
	if !ok {
		return errors.Errorf("%s: csv type must be string", stack)
	}
	parts := strings.Split(str, ",")
	csv := Csv{}
	for _, part := range parts {
		i, err := strconv.ParseInt(part, 10, 0)
		if err != nil {
			return errors.Wrapf(err, "%s: converting %s to int", stack, part)
		}
		csv = append(csv, int(i))
	}
	*c = csv
	return nil
}

func (c *Csv) Repack(root *frizz.Root, stack frizz.Stack) (interface{}, error) {
	var out string
	for i, v := range *c {
		if i != 0 {
			out += ","
		}
		out += fmt.Sprint(v)
	}
	return out, nil
}

// frizz
type Type struct {
	Path string
	Name string
}

func (t *Type) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}) error {
	str, ok := in.(string)
	if !ok {
		return errors.Errorf("%s: custom type must be string", stack)
	}
	expr, err := parser.ParseExpr(str)
	if err != nil {
		return errors.Wrapf(err, "%s: parsing expr", stack)
	}
	switch expr := expr.(type) {
	case *ast.SelectorExpr:
		x, ok := expr.X.(*ast.Ident)
		if !ok {
			return errors.Errorf("%s: expr.X must be *ast.Ident", stack)
		}
		path, ok := root.Imports[x.Name]
		if !ok {
			return errors.Errorf("%s: alias %s not found in imports", stack, x.Name)
		}
		*t = Type{Path: path, Name: expr.Sel.Name}
		return nil
	case *ast.Ident:
		*t = Type{Path: root.Path, Name: expr.Name}
		return nil
	default:
		return errors.Errorf("%s: expr must be *ast.SelectorExpr or *ast.Ident", stack)
	}
}

func (t *Type) Repack(root *frizz.Root, stack frizz.Stack) (interface{}, error) {
	if t.Path == root.Path {
		return t.Name, nil
	}
	for alias, path := range root.Imports {
		if path == t.Path {
			return fmt.Sprintf("%s.%s", alias, t.Name), nil
		}
	}
	return nil, errors.Errorf("%s: can't find %s in imports", stack, t.Path)
}

// frizz
type Custom struct {
	ast.Expr
}

func (c *Custom) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}) error {
	str, ok := in.(string)
	if !ok {
		return errors.Errorf("%s: custom type must be string", stack)
	}
	expr, err := parser.ParseExpr(str)
	if err != nil {
		return errors.Wrapf(err, "%s: parsing expr", stack)
	}
	*c = Custom{Expr: expr}
	return nil
}

func (c *Custom) Repack(root *frizz.Root, stack frizz.Stack) (interface{}, error) {
	buf := &bytes.Buffer{}
	err := printer.Fprint(buf, nil, c.Expr)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return buf.String(), nil
}

// frizz
type EmbedNatives struct {
	Natives
	Int int
}

// frizz
type EmbedPointer struct {
	*Natives
	Int int
}

// frizz
type EmbedQualPointer struct {
	*sub.Sub
	Int int
}

// frizz
type EmbedQual struct {
	sub.Sub
	Int int
}

// frizz
type InterfaceField struct {
	Iface Interface
	Slice []Interface
	Array [3]Interface
	Map   map[string]Interface
}

// frizz
type Impi struct {
	Int int
}

func (i Impi) Foo() string {
	// notest
	return fmt.Sprint(i.Int)
}

// frizz
type Imps struct {
	String string
}

func (i Imps) Foo() string {
	// notest
	return i.String
}

// frizz
type Interface interface {
	Foo() string
}

// frizz
type Private struct {
	i int
	s string
}

// frizz
type AliasSub sub.Sub

// frizz
type AliasSlice []Int

// frizz
type AliasArray [3]string

// frizz
type AliasMap map[string]*Qual

// frizz
type AliasPointer *Int

// frizz
type Alias Int

// frizz
type Int int

// frizz
type String string

// frizz
type Qual struct {
	Sub sub.Sub
}

// frizz
type Pointers struct {
	String      *string
	Int         *Int
	Sub         *sub.Sub
	Array       *[3]int
	Slice       *[]string
	Map         *map[string]int
	SliceString []*string
	SliceInt    []*Int
	SliceSub    []*sub.Sub
}

// frizz
type Maps struct {
	Ints    map[string]int
	Strings map[string]string
	Slices  map[string][]string
	Arrays  map[string][2]int
	Maps    map[string]map[string]string
}

const N = 2

// frizz
type Slices struct {
	Ints      []int
	Strings   []string
	ArrayLit  [5]string
	ArrayExpr [2 * N]int
	Structs   []struct {
		Int int
	}
	Arrays [][]string
}

// frizz
type Structs struct {
	Simple struct {
		Int  int
		Bool bool
	}
	Complex struct {
		String string
		Inner  struct {
			Float32 float32
		}
	}
}

// frizz
type Natives struct {
	Bool    bool
	Byte    byte
	Float32 float32
	Float64 float64
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Uint    uint
	Uint8   uint8
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Rune    rune
	String  string
}
