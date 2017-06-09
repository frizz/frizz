package unpacker

import (
	"fmt"

	"frizz.io/tests/unpacker/sub"
)

/*
// frizz
type Custom struct {
	ast.Expr
}

func (Custom) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}) (interface{}, error) {
	str, ok := in.(string)
	if !ok {
		return nil, errors.Errorf("%s: custom type must be string", stack)
	}
	expr, err := parser.ParseExpr(str)
	if err != nil {
		return nil, errors.Wrapf(err, "%s: parsing type string", stack)
	}
	return Custom(expr), nil
}
*/

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
