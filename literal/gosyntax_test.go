package fmt

import (
	"testing"

	"kego.io/assert"
)

func TestGoSyntax(t *testing.T) {

	type a struct {
		as string
	}
	type b struct {
		ba *a
	}
	foo := b{&a{"c"}}

	s := GoSyntax("", map[string]string{}, foo)
	assert.Equal(t, `fmt.b{ba:&fmt.a{as:"c"}}`, s)

	s = GoSyntax("kego.io/fmt", map[string]string{"a": "a.b/c"}, foo)
	assert.Equal(t, `b{ba:&a{as:"c"}}`, s)

	s = GoSyntax("kego.io/system", map[string]string{"f": "kego.io/fmt"}, foo)
	assert.Equal(t, `f.b{ba:&f.a{as:"c"}}`, s)

	s = GoSyntax("kego.io/system", map[string]string{"fmt": "foo.com/fmt", "f1": "kego.io/fmt"}, foo)
	assert.Equal(t, `f1.b{ba:&f1.a{as:"c"}}`, s)
}

func TestGoSyntaxTypes(t *testing.T) {

	type a struct {
		as string
	}
	type b struct {
		bs  string
		bm  map[string]string
		bmt map[string]a
		bmp map[string]*a
		ba  []string
		bat []a
		bap []*a
	}
	foo := b{
		bs:  "a",
		bm:  map[string]string{"b": "c"},
		bmt: map[string]a{"d": a{"e"}},
		bmp: map[string]*a{"f": &a{"g"}},
		ba:  []string{"h", "i"},
		bat: []a{a{"j"}, a{"k"}},
		bap: []*a{&a{"l"}, &a{"m"}},
	}

	s := GoSyntax("", map[string]string{}, foo)
	assert.Equal(t, "fmt.b{bs:\"a\", bm:map[string]string{\"b\":\"c\"}, bmt:map[string]fmt.a{\"d\":fmt.a{as:\"e\"}}, bmp:map[string]*fmt.a{\"f\":&fmt.a{as:\"g\"}}, ba:[]string{\"h\", \"i\"}, bat:[]fmt.a{fmt.a{as:\"j\"}, fmt.a{as:\"k\"}}, bap:[]*fmt.a{&fmt.a{as:\"l\"}, &fmt.a{as:\"m\"}}}", s)

	s = GoSyntax("kego.io/fmt", map[string]string{"a": "a.b/c"}, foo)
	assert.Equal(t, "b{bs:\"a\", bm:map[string]string{\"b\":\"c\"}, bmt:map[string]a{\"d\":a{as:\"e\"}}, bmp:map[string]*a{\"f\":&a{as:\"g\"}}, ba:[]string{\"h\", \"i\"}, bat:[]a{a{as:\"j\"}, a{as:\"k\"}}, bap:[]*a{&a{as:\"l\"}, &a{as:\"m\"}}}", s)

	s = GoSyntax("kego.io/system", map[string]string{"f": "kego.io/fmt"}, foo)
	assert.Equal(t, "f.b{bs:\"a\", bm:map[string]string{\"b\":\"c\"}, bmt:map[string]f.a{\"d\":f.a{as:\"e\"}}, bmp:map[string]*f.a{\"f\":&f.a{as:\"g\"}}, ba:[]string{\"h\", \"i\"}, bat:[]f.a{f.a{as:\"j\"}, f.a{as:\"k\"}}, bap:[]*f.a{&f.a{as:\"l\"}, &f.a{as:\"m\"}}}", s)

	s = GoSyntax("kego.io/system", map[string]string{"fmt": "foo.com/fmt", "f1": "kego.io/fmt"}, foo)
	assert.Equal(t, "f1.b{bs:\"a\", bm:map[string]string{\"b\":\"c\"}, bmt:map[string]f1.a{\"d\":f1.a{as:\"e\"}}, bmp:map[string]*f1.a{\"f\":&f1.a{as:\"g\"}}, ba:[]string{\"h\", \"i\"}, bat:[]f1.a{f1.a{as:\"j\"}, f1.a{as:\"k\"}}, bap:[]*f1.a{&f1.a{as:\"l\"}, &f1.a{as:\"m\"}}}", s)

}

func TestGoSyntaxTypesNil(t *testing.T) {

	type a struct {
		as string
	}
	type b struct {
		a   *a
		bm  map[string]string
		bmt map[string]a
		bmp map[string]*a
		ba  []string
		bat []a
		bap []*a
	}
	foo := b{
		a:   nil,
		bm:  nil,
		bmt: nil,
		bmp: nil,
		ba:  nil,
		bat: nil,
		bap: nil,
	}

	s := GoSyntax("", map[string]string{}, foo)
	assert.Equal(t, "fmt.b{a:(*fmt.a)(nil), bm:map[string]string(nil), bmt:map[string]fmt.a(nil), bmp:map[string]*fmt.a(nil), ba:[]string(nil), bat:[]fmt.a(nil), bap:[]*fmt.a(nil)}", s)

	s = GoSyntax("kego.io/fmt", map[string]string{"a": "a.b/c"}, foo)
	assert.Equal(t, "b{a:(*a)(nil), bm:map[string]string(nil), bmt:map[string]a(nil), bmp:map[string]*a(nil), ba:[]string(nil), bat:[]a(nil), bap:[]*a(nil)}", s)

	s = GoSyntax("kego.io/system", map[string]string{"f": "kego.io/fmt"}, foo)
	assert.Equal(t, "f.b{a:(*f.a)(nil), bm:map[string]string(nil), bmt:map[string]f.a(nil), bmp:map[string]*f.a(nil), ba:[]string(nil), bat:[]f.a(nil), bap:[]*f.a(nil)}", s)

	s = GoSyntax("kego.io/system", map[string]string{"fmt": "foo.com/fmt", "f1": "kego.io/fmt"}, foo)
	assert.Equal(t, "f1.b{a:(*f1.a)(nil), bm:map[string]string(nil), bmt:map[string]f1.a(nil), bmp:map[string]*f1.a(nil), ba:[]string(nil), bat:[]f1.a(nil), bap:[]*f1.a(nil)}", s)

}
