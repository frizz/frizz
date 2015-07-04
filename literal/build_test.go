package literal

import (
	"fmt"
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/generator"
)

func TestBuild(t *testing.T) {

	type a struct {
		As string
	}
	type b struct {
		Ba *a
	}
	foo := &b{&a{"c"}}

	o := []string{}
	p := map[string]string{}
	i := generator.NewImports_test()
	n := Build(foo, p, &o, "", i.Add_test)
	inner := o[0]
	outer := o[1]
	assert.Equal(t, n, outer)
	assert.Equal(t, `&literal.a{As:"c"}`, p[inner])
	assert.Equal(t, fmt.Sprintf(`&literal.b{Ba:%s}`, inner), p[outer])

	o = []string{}
	p = map[string]string{}
	i = generator.NewImports_test(map[string]string{"a.b/c": "a"})
	n = Build(foo, p, &o, "kego.io/literal", i.Add_test)
	inner = o[0]
	outer = o[1]
	assert.Equal(t, n, outer)
	assert.Equal(t, `&a{As:"c"}`, p[inner])
	assert.Equal(t, fmt.Sprintf(`&b{Ba:%s}`, inner), p[outer])

	o = []string{}
	p = map[string]string{}
	i = generator.NewImports_test(map[string]string{"foo.com/literal": "literal", "kego.io/literal": "f"})
	n = Build(foo, p, &o, "kego.io/system", i.Add_test)
	inner = o[0]
	outer = o[1]
	assert.Equal(t, n, outer)
	assert.Equal(t, `&f.a{As:"c"}`, p[inner])
	assert.Equal(t, fmt.Sprintf(`&f.b{Ba:%s}`, inner), p[outer])

}

func TestBuildTypes(t *testing.T) {

	type a struct {
		As string
	}
	type b struct {
		Bs  string
		Bm  map[string]string
		Bmt map[string]a
		Bmp map[string]*a
		Ba  []string
		Bat []a
		Bap []*a
	}
	foo := &b{
		Bs:  "a",
		Bm:  map[string]string{"b": "c"},
		Bmt: map[string]a{"d": a{"e"}},
		Bmp: map[string]*a{"f": &a{"g"}},
		Ba:  []string{"h", "i"},
		Bat: []a{a{"j"}, a{"k"}},
		Bap: []*a{&a{"l"}, &a{"m"}},
	}

	o := []string{}
	p := map[string]string{}
	i := generator.NewImports_test()
	n := Build(foo, p, &o, "", i.Add_test)
	inner1 := o[0]
	inner2 := o[1]
	inner3 := o[2]
	outer := o[3]
	assert.Equal(t, n, outer)
	assert.Equal(t, fmt.Sprintf(`&literal.b{Bs:"a", Bm:map[string]string{"b":"c"}, Bmt:map[string]literal.a{"d":literal.a{As:"e"}}, Bmp:map[string]*literal.a{"f":%s}, Ba:[]string{"h", "i"}, Bat:[]literal.a{literal.a{As:"j"}, literal.a{As:"k"}}, Bap:[]*literal.a{%s, %s}}`, inner1, inner2, inner3), p[outer])
	assert.Equal(t, `&literal.a{As:"g"}`, p[inner1])
	assert.Equal(t, `&literal.a{As:"l"}`, p[inner2])
	assert.Equal(t, `&literal.a{As:"m"}`, p[inner3])

	o = []string{}
	p = map[string]string{}
	i = generator.NewImports_test(map[string]string{"a.b/c": "a"})
	n = Build(foo, p, &o, "kego.io/literal", i.Add_test)
	inner1 = o[0]
	inner2 = o[1]
	inner3 = o[2]
	outer = o[3]
	assert.Equal(t, n, outer)
	assert.Equal(t, fmt.Sprintf(`&b{Bs:"a", Bm:map[string]string{"b":"c"}, Bmt:map[string]a{"d":a{As:"e"}}, Bmp:map[string]*a{"f":%s}, Ba:[]string{"h", "i"}, Bat:[]a{a{As:"j"}, a{As:"k"}}, Bap:[]*a{%s, %s}}`, inner1, inner2, inner3), p[outer])
	assert.Equal(t, `&a{As:"g"}`, p[inner1])
	assert.Equal(t, `&a{As:"l"}`, p[inner2])
	assert.Equal(t, `&a{As:"m"}`, p[inner3])

	o = []string{}
	p = map[string]string{}
	i = generator.NewImports_test(map[string]string{"foo.com/literal": "literal", "kego.io/literal": "f"})
	n = Build(foo, p, &o, "kego.io/system", i.Add_test)
	inner1 = o[0]
	inner2 = o[1]
	inner3 = o[2]
	outer = o[3]
	assert.Equal(t, n, outer)
	assert.Equal(t, fmt.Sprintf(`&f.b{Bs:"a", Bm:map[string]string{"b":"c"}, Bmt:map[string]f.a{"d":f.a{As:"e"}}, Bmp:map[string]*f.a{"f":%s}, Ba:[]string{"h", "i"}, Bat:[]f.a{f.a{As:"j"}, f.a{As:"k"}}, Bap:[]*f.a{%s, %s}}`, inner1, inner2, inner3), p[outer])
	assert.Equal(t, `&f.a{As:"g"}`, p[inner1])
	assert.Equal(t, `&f.a{As:"l"}`, p[inner2])
	assert.Equal(t, `&f.a{As:"m"}`, p[inner3])
}

func TestBuildTypesNil(t *testing.T) {

	type a struct {
		As string
	}
	type b struct {
		A   *a
		Bm  map[string]string
		Bmt map[string]a
		Bmp map[string]*a
		Ba  []string
		Bat []a
		Bap []*a
	}
	foo := &b{
		A:   nil,
		Bm:  nil,
		Bmt: nil,
		Bmp: nil,
		Ba:  nil,
		Bat: nil,
		Bap: nil,
	}

	o := []string{}
	p := map[string]string{}
	i := generator.NewImports_test(map[string]string{})
	n := Build(foo, p, &o, "", i.Add_test)
	assert.Equal(t, n, o[0])
	assert.Equal(t, `&literal.b{A:(*literal.a)(nil), Bm:map[string]string(nil), Bmt:map[string]literal.a(nil), Bmp:map[string]*literal.a(nil), Ba:[]string(nil), Bat:[]literal.a(nil), Bap:[]*literal.a(nil)}`, p[n])

	o = []string{}
	p = map[string]string{}
	i = generator.NewImports_test(map[string]string{"a.b/c": "a"})
	n = Build(foo, p, &o, "kego.io/literal", i.Add_test)
	assert.Equal(t, n, o[0])
	assert.Equal(t, `&b{A:(*a)(nil), Bm:map[string]string(nil), Bmt:map[string]a(nil), Bmp:map[string]*a(nil), Ba:[]string(nil), Bat:[]a(nil), Bap:[]*a(nil)}`, p[n])

	o = []string{}
	p = map[string]string{}
	i = generator.NewImports_test(map[string]string{"foo.com/literal": "literal", "kego.io/literal": "f"})
	n = Build(foo, p, &o, "kego.io/system", i.Add_test)
	assert.Equal(t, n, o[0])
	assert.Equal(t, `&f.b{A:(*f.a)(nil), Bm:map[string]string(nil), Bmt:map[string]f.a(nil), Bmp:map[string]*f.a(nil), Ba:[]string(nil), Bat:[]f.a(nil), Bap:[]*f.a(nil)}`, p[n])
}
