package literal

import (
	"testing"

	"kego.io/kerr/assert"
	"kego.io/process/generate/builder"
)

func TestBuild(t *testing.T) {

	type a struct {
		As string
	}
	type b struct {
		Ba *a
	}
	foo := &b{&a{"c"}}

	p := []Pointer{}
	i := builder.Imports{}
	n := Build(foo, &p, "", i.Add)
	assert.Equal(t, "ptr1", n.Name)
	assert.Equal(t, `&literal.a{As:"c"}`, p[0].Source)
	assert.Equal(t, `&literal.b{Ba:ptr0}`, p[1].Source)

	p = []Pointer{}
	i = builder.Imports{"a.b/c": builder.Import{Path: "a.b/c", Name: "c", Alias: "a"}}
	n = Build(foo, &p, "kego.io/process/generate/literal", i.Add)
	assert.Equal(t, "ptr1", n.Name)
	assert.Equal(t, `&a{As:"c"}`, p[0].Source)
	assert.Equal(t, `&b{Ba:ptr0}`, p[1].Source)

	p = []Pointer{}
	i = builder.Imports{
		"foo.com/literal":                  builder.Import{Path: "foo.com/literal", Name: "literal", Alias: "literal"},
		"kego.io/process/generate/literal": builder.Import{Path: "kego.io/process/generate/literal", Name: "literal", Alias: "f"}}
	n = Build(foo, &p, "kego.io/system", i.Add)
	assert.Equal(t, "ptr1", n.Name)
	assert.Equal(t, `&f.a{As:"c"}`, p[0].Source)
	assert.Equal(t, `&f.b{Ba:ptr0}`, p[1].Source)

}

func TestBuildPlain(t *testing.T) {

	type in int
	type st string
	type bo bool
	type fl float64
	type a struct {
		Ain in
		Ast st
		Abo bo
		Afl fl
		Bin *in
		Bst *st
		Bbo *bo
		Bfl *fl
	}
	vin := in(2)
	vst := st("c")
	vbo := bo(true)
	vfl := fl(2.0)
	foo := &a{vin, vst, vbo, vfl, &vin, &vst, &vbo, &vfl}

	p := []Pointer{}
	i := builder.Imports{}
	n := Build(foo, &p, "", i.Add)

	//for i, po := range p {
	//	fmt.Println(i, po.Source)
	//}

	assert.Equal(t, "ptr4", n.Name)
	assert.Equal(t, `literal.in(2)`, p[0].Source)
	assert.Equal(t, `literal.st("c")`, p[1].Source)
	assert.Equal(t, `literal.bo(true)`, p[2].Source)
	assert.Equal(t, `literal.fl(2)`, p[3].Source)
	assert.Equal(t, `&literal.a{Ain:literal.in(2), Ast:literal.st("c"), Abo:literal.bo(true), Afl:literal.fl(2), Bin:&ptr0, Bst:&ptr1, Bbo:&ptr2, Bfl:&ptr3}`, p[4].Source)

}

func TestOrder(t *testing.T) {
	foo := &map[string]string{
		"e": "f",
		"g": "h",
		"i": "j",
		"a": "b",
		"c": "d",
	}
	p := []Pointer{}
	i := builder.Imports{}
	n := Build(foo, &p, "", i.Add)
	assert.Equal(t, "ptr0", n.Name)
	assert.Equal(t, `&map[string]string{"a":"b", "c":"d", "e":"f", "g":"h", "i":"j"}`, p[0].Source)
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

	p := []Pointer{}
	i := builder.Imports{}
	n := Build(foo, &p, "", i.Add)
	assert.Equal(t, "ptr3", n.Name)
	assert.Equal(t, `&literal.b{Bs:"a", Bm:map[string]string{"b":"c"}, Bmt:map[string]literal.a{"d":literal.a{As:"e"}}, Bmp:map[string]*literal.a{"f":ptr0}, Ba:[]string{"h", "i"}, Bat:[]literal.a{literal.a{As:"j"}, literal.a{As:"k"}}, Bap:[]*literal.a{ptr1, ptr2}}`, p[3].Source)
	assert.Equal(t, `&literal.a{As:"g"}`, p[0].Source)
	assert.Equal(t, `&literal.a{As:"l"}`, p[1].Source)
	assert.Equal(t, `&literal.a{As:"m"}`, p[2].Source)

	p = []Pointer{}
	i = builder.Imports{"a.b/c": builder.Import{Path: "a.b/c", Alias: "a", Name: "c"}}
	n = Build(foo, &p, "kego.io/process/generate/literal", i.Add)
	assert.Equal(t, "ptr3", n.Name)
	assert.Equal(t, `&b{Bs:"a", Bm:map[string]string{"b":"c"}, Bmt:map[string]a{"d":a{As:"e"}}, Bmp:map[string]*a{"f":ptr0}, Ba:[]string{"h", "i"}, Bat:[]a{a{As:"j"}, a{As:"k"}}, Bap:[]*a{ptr1, ptr2}}`, p[3].Source)
	assert.Equal(t, `&a{As:"g"}`, p[0].Source)
	assert.Equal(t, `&a{As:"l"}`, p[1].Source)
	assert.Equal(t, `&a{As:"m"}`, p[2].Source)

	p = []Pointer{}
	i = builder.Imports{
		"foo.com/literal":                  builder.Import{Path: "foo.com/literal", Name: "literal", Alias: "literal"},
		"kego.io/process/generate/literal": builder.Import{Path: "kego.io/process/generate/literal", Name: "literal", Alias: "f"}}
	n = Build(foo, &p, "kego.io/system", i.Add)
	assert.Equal(t, "ptr3", n.Name)
	assert.Equal(t, `&f.b{Bs:"a", Bm:map[string]string{"b":"c"}, Bmt:map[string]f.a{"d":f.a{As:"e"}}, Bmp:map[string]*f.a{"f":ptr0}, Ba:[]string{"h", "i"}, Bat:[]f.a{f.a{As:"j"}, f.a{As:"k"}}, Bap:[]*f.a{ptr1, ptr2}}`, p[3].Source)
	assert.Equal(t, `&f.a{As:"g"}`, p[0].Source)
	assert.Equal(t, `&f.a{As:"l"}`, p[1].Source)
	assert.Equal(t, `&f.a{As:"m"}`, p[2].Source)

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

	p := []Pointer{}
	i := builder.Imports{}
	n := Build(foo, &p, "", i.Add)
	assert.Equal(t, "ptr0", n.Name)
	assert.Equal(t, `&literal.b{A:(*literal.a)(nil), Bm:map[string]string(nil), Bmt:map[string]literal.a(nil), Bmp:map[string]*literal.a(nil), Ba:[]string(nil), Bat:[]literal.a(nil), Bap:[]*literal.a(nil)}`, p[0].Source)

	p = []Pointer{}
	i = builder.Imports{"a.b/c": builder.Import{Path: "a.b/c", Alias: "a", Name: "c"}}
	n = Build(foo, &p, "kego.io/process/generate/literal", i.Add)
	assert.Equal(t, "ptr0", n.Name)
	assert.Equal(t, `&b{A:(*a)(nil), Bm:map[string]string(nil), Bmt:map[string]a(nil), Bmp:map[string]*a(nil), Ba:[]string(nil), Bat:[]a(nil), Bap:[]*a(nil)}`, p[0].Source)

	p = []Pointer{}
	i = builder.Imports{
		"foo.com/literal":                  builder.Import{Path: "foo.com/literal", Name: "literal", Alias: "literal"},
		"kego.io/process/generate/literal": builder.Import{Path: "kego.io/process/generate/literal", Name: "literal", Alias: "f"}}
	n = Build(foo, &p, "kego.io/system", i.Add)
	assert.Equal(t, "ptr0", n.Name)
	assert.Equal(t, `&f.b{A:(*f.a)(nil), Bm:map[string]string(nil), Bmt:map[string]f.a(nil), Bmp:map[string]*f.a(nil), Ba:[]string(nil), Bat:[]f.a(nil), Bap:[]*f.a(nil)}`, p[0].Source)
}
