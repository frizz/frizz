package generator

import (
	"fmt"
	"reflect"
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
	"kego.io/system"
)

func TestGetPointer(t *testing.T) {
	//!isNative && !isInterface => "*"
	assert.Equal(t, "", getPointer(&system.Type{Native: system.NewString("bool"), Interface: false}))
	assert.Equal(t, "", getPointer(&system.Type{Native: system.NewString("object"), Interface: true}))
	assert.Equal(t, "", getPointer(&system.Type{Native: system.NewString("bool"), Interface: true}))
	assert.Equal(t, "*", getPointer(&system.Type{Native: system.NewString("object"), Interface: false}))
	assert.Equal(t, "*", getPointer(&system.Type{Native: system.NewString("map"), Interface: false}))
	assert.Equal(t, "*", getPointer(&system.Type{Native: system.NewString("array"), Interface: false}))
}

func TestFormatTag(t *testing.T) {
	type parentStruct struct {
		*system.Object
	}
	type ruleStruct struct {
		*system.Object
		*system.Rule
	}
	parentType := &system.Type{
		Object: &system.Object{Id: system.NewReference("a.b/c", "a"), Type: system.NewReference("kego.io/system", "type")},
	}
	ruleType := &system.Type{
		Object: &system.Object{Id: system.NewReference("a.b/c", "@a"), Type: system.NewReference("kego.io/system", "type")},
	}
	json.Register("a.b/c", "a", reflect.TypeOf(&parentStruct{}), nil, 0)
	json.Register("a.b/c", "@a", reflect.TypeOf(&ruleStruct{}), nil, 0)
	system.Register("a.b/c", "a", parentType, 0)
	system.Register("a.b/c", "@a", ruleType, 0)
	defer json.Unregister("a.b/c", "a")
	defer json.Unregister("a.b/c", "@a")
	defer system.Unregister("a.b/c", "a")
	defer system.Unregister("a.b/c", "@a")

	r := &system.RuleWrapper{
		Interface: &ruleStruct{},
		Parent:    parentType,
	}
	s, err := formatTag("n", false, []byte("null"), r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`json:\"n\"`", s)

	s, err = formatTag("n", true, []byte("null"), r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`json:\"-\"`", s)

	s, err = formatTag("n", false, []byte(`"a"`), r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"a\\\",\\\"path\\\":\\\"d.e/f\\\"}}\" json:\"n\"`", s)

	parentType.Id = system.NewReference("kego.io/system", "string")
	s, err = formatTag("n", false, []byte(`"a"`), r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"value\\\":\\\"a\\\"}}\" json:\"n\"`", s)

	_, err = formatTag("n", false, []byte(`foo`), r, "d.e/f", map[string]string{})
	assert.IsError(t, err, "LKBWJTMJCF")
}

type structWithCustomMarshaler struct {
	*system.Object
	throwError bool
}

func (s *structWithCustomMarshaler) MarshalJSON() ([]byte, error) {
	if s.throwError {
		return nil, fmt.Errorf("Here's an error")
	}
	return []byte(`"foo"`), nil
}

var _ json.Marshaler = (*structWithCustomMarshaler)(nil)

func TestMarshaler(t *testing.T) {
	f := structWithCustomMarshaler{}
	s, err := f.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(`"foo"`), s)
}

// Json marshaling throws an error when fed a chan, so
// it's a convenient way to force an error.
type typeThatWillCauseJsonMarshalToError chan string

type structWithoutCustomMarshaler struct {
	A string
}

type ruleStructA struct {
	*system.Object
	*system.Rule
	B system.String
}
type ruleStructB struct {
	*system.Object
	*system.Rule
	Default system.String
}

func (r *ruleStructB) GetDefault() interface{} {
	return r.Default
}

var _ system.DefaultRule = (*ruleStructB)(nil)

type ruleStructC struct {
	*system.Object
	*system.Rule
	Default *structWithCustomMarshaler
}

func (r *ruleStructC) GetDefault() interface{} {
	return r.Default
}

var _ system.DefaultRule = (*ruleStructC)(nil)

type ruleStructD struct {
	*system.Object
	*system.Rule
	Default typeThatWillCauseJsonMarshalToError
}

func (r *ruleStructD) GetDefault() interface{} {
	return r.Default
}

var _ system.DefaultRule = (*ruleStructD)(nil)

type ruleStructE struct {
	*system.Object
	*system.Rule
	Default structWithoutCustomMarshaler
}

func (r *ruleStructE) GetDefault() interface{} {
	return r.Default
}

var _ system.DefaultRule = (*ruleStructE)(nil)

func TestGetTag(t *testing.T) {
	type parentStruct struct {
		*system.Object
	}

	parentType := &system.Type{
		Object: &system.Object{Id: system.NewReference("a.b/c", "a"), Type: system.NewReference("kego.io/system", "type")},
	}
	ruleType := &system.Type{
		Object: &system.Object{Id: system.NewReference("a.b/c", "@a"), Type: system.NewReference("kego.io/system", "type")},
	}
	json.Register("a.b/c", "@a", reflect.TypeOf(&ruleStructA{}), nil, 0)
	json.Register("a.b/c", "@b", reflect.TypeOf(&ruleStructB{}), nil, 0)
	json.Register("a.b/c", "@c", reflect.TypeOf(&ruleStructC{}), nil, 0)
	json.Register("a.b/c", "@d", reflect.TypeOf(&ruleStructD{}), nil, 0)
	json.Register("a.b/c", "@e", reflect.TypeOf(&ruleStructE{}), nil, 0)
	for _, letter := range []string{"a", "b", "c", "d", "e"} {
		json.Register("a.b/c", letter, reflect.TypeOf(&parentStruct{}), nil, 0)
		system.Register("a.b/c", letter, parentType, 0)
		system.Register("a.b/c", fmt.Sprintf("@%s", letter), ruleType, 0)
		defer json.Unregister("a.b/c", letter)
		defer json.Unregister("a.b/c", fmt.Sprintf("@%s", letter))
		defer system.Unregister("a.b/c", letter)
		defer system.Unregister("a.b/c", fmt.Sprintf("@%s", letter))
	}

	r := &system.RuleWrapper{
		Interface: &ruleStructA{},
		Parent:    parentType,
	}

	// rule has no default field
	s, err := getTag("n", false, r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`json:\"n\"`", s)

	s, err = getTag("n", true, r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`json:\"-\"`", s)

	r.Interface = &ruleStructB{Default: system.NewString("c")}
	s, err = getTag("n", false, r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"c\\\",\\\"path\\\":\\\"d.e/f\\\"}}\" json:\"n\"`", s)

	r.Interface = &ruleStructC{Default: &structWithCustomMarshaler{Object: &system.Object{Id: system.NewReference("d.e/f", "f")}}}
	s, err = getTag("n", false, r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"foo\\\",\\\"path\\\":\\\"d.e/f\\\"}}\" json:\"n\"`", s)

	r.Interface = &ruleStructC{Default: &structWithCustomMarshaler{Object: &system.Object{Id: system.NewReference("d.e/f", "f")}, throwError: true}}
	s, err = getTag("n", false, r, "d.e/f", map[string]string{})
	assert.IsError(t, err, "YIEMHYFVCD")

	r.Interface = &ruleStructD{Default: make(typeThatWillCauseJsonMarshalToError)}
	s, err = getTag("n", false, r, "d.e/f", map[string]string{})
	assert.IsError(t, err, "QQDOLAJKLU")

	r.Interface = &ruleStructE{Default: structWithoutCustomMarshaler{A: "b"}}
	s, err = getTag("n", false, r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":{\\\"A\\\":\\\"b\\\"},\\\"path\\\":\\\"d.e/f\\\"}}\" json:\"n\"`", s)

}

func TestGoTypeDescriptor(t *testing.T) {
	p := &system.StringRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/json", "@string"),
		},
		Rule: &system.Rule{},
	}
	i := Imports{}
	s, err := Type("n", p, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "string `json:\"n\"`", s)

	p = &system.StringRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@string"),
		},
		Rule: &system.Rule{
			Exclude: true,
		},
	}
	s, err = Type("n", p, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "String `json:\"-\"`", s)

	p = &system.StringRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@string"),
		},
		Rule: &system.Rule{},
	}
	s, err = Type("n", p, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "String `json:\"n\"`", s)

	s, err = Type("n", p, "kego.io/a", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "system.String `json:\"n\"`", s)

	// We're just using type here because it's a handy
	// non-native type in the system package
	pt := &system.TypeRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@type"),
		},
		Rule: &system.Rule{},
	}
	s, err = Type("n", pt, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "*Type `json:\"n\"`", s)

	// We're just using property here because it's a handy
	// non-native type in the system package
	pt = &system.TypeRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@type"),
		},
		Rule: &system.Rule{},
	}
	s, err = Type("n", pt, "kego.io/a", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "*system.Type `json:\"n\"`", s)

	type a struct{}
	type aRule struct {
		*system.Object
		*system.Rule
	}
	tyr := &system.Type{Object: &system.Object{Id: system.NewReference("b.c/d", "@a")}}
	ty := &system.Type{Object: &system.Object{Id: system.NewReference("b.c/d", "a")}}
	json.Register("b.c/d", "a", reflect.TypeOf(&a{}), nil, 0)
	json.Register("b.c/d", "@a", reflect.TypeOf(&aRule{}), nil, 0)
	system.Register("b.c/d", "a", ty, 0)
	system.Register("b.c/d", "@a", tyr, 0)
	defer json.Unregister("b.c/d", "a")
	defer json.Unregister("b.c/d", "@a")
	defer system.Unregister("b.c/d", "a")
	defer system.Unregister("b.c/d", "@a")

	pa := &aRule{
		Object: &system.Object{
			Type: system.NewReference("b.c/d", "@a"),
		},
		Rule: &system.Rule{},
	}
	ia := Imports{"b.c/d": Import{Path: "b.c/d", Name: "d", Alias: "d"}}
	s, err = Type("n", pa, "kego.io/system", ia.Add)
	assert.NoError(t, err)
	assert.Equal(t, "*d.A `json:\"n\"`", s)

	s, err = Type("n", pa, "b.c/d", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "*A `json:\"n\"`", s)

	p = &system.StringRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@string"),
		},
		Rule:    &system.Rule{},
		Default: system.NewString("a"),
	}
	s, err = Type("n", p, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "String `kego:\"{\\\"default\\\":{\\\"value\\\":\\\"a\\\"}}\" json:\"n\"`", s)

	pm := &system.MapRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@map"),
		},
		Rule: &system.Rule{},
		Items: &system.StringRule{
			Object: &system.Object{
				Type: system.NewReference("kego.io/system", "@string"),
			},
		},
	}
	s, err = Type("n", pm, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "map[string]String `json:\"n\"`", s)

	par := &system.ArrayRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@array"),
		},
		Rule: &system.Rule{},
		Items: &system.StringRule{
			Object: &system.Object{
				Type: system.NewReference("kego.io/system", "@string"),
			},
		},
	}
	s, err = Type("n", par, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "[]String `json:\"n\"`", s)

	pm = &system.MapRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@map"),
		},
		Rule: &system.Rule{},
		Items: &system.ArrayRule{
			Object: &system.Object{
				Type: system.NewReference("kego.io/system", "@array"),
			},
			Items: &system.StringRule{
				Object: &system.Object{
					Type: system.NewReference("kego.io/system", "@string"),
				},
			},
		},
	}
	s, err = Type("n", pm, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "map[string][]String `json:\"n\"`", s)

	pm = &system.MapRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@map"),
		},
		Rule: &system.Rule{},
		Items: &system.ArrayRule{
			Object: &system.Object{
				Type: system.NewReference("kego.io/system", "@array"),
			},
			Items: &system.StringRule{
				Object: &system.Object{
					Type: system.NewReference("kego.io/system", "@string"),
				},
				Default: system.NewString("a"),
			},
		},
	}
	s, err = Type("n", pm, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "map[string][]String `kego:\"{\\\"default\\\":{\\\"value\\\":\\\"a\\\"}}\" json:\"n\"`", s)

}

func TypeErrors_NeedsTypes(t *testing.T) {

	p := &system.JsonStringRule{
		Object: &system.Object{
			Type: system.NewReference("a.b/c", "notFoundType"),
		},
		Rule: &system.Rule{},
	}
	i := Imports{}
	_, err := Type("n", p, "kego.io/system", i.Add)
	// Item is an unregistered type, so errors at NewRuleHolder
	assert.IsError(t, err, "TFXFBIRXHN")
	assert.HasError(t, err, "KYCTDXKFYR")

	pm := &system.MapRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@map"),
		},
		Rule: &system.Rule{},
	}
	_, err = Type("n", pm, "kego.io/system", i.Add)
	// Collection item @map doesn't have Items field, so errors at ItemsRule
	assert.IsError(t, err, "SOGEFOPJHB")
	assert.HasError(t, err, "SUJLYBXPYS")

	type a struct{}
	type aRule struct {
		*system.Object
		*system.Rule
	}
	tyr := &system.Type{Object: &system.Object{Id: system.NewReference("b.c/d", "@a")}}
	ty := &system.Type{Object: &system.Object{Id: system.NewReference("b.c/d", "a")}}
	json.Register("b.c/d", "a", reflect.TypeOf(&a{}), nil, 0)
	json.Register("b.c/d", "@a", reflect.TypeOf(&aRule{}), nil, 0)
	system.Register("b.c/d", "a", ty, 0)
	system.Register("b.c/d", "@a", tyr, 0)
	defer json.Unregister("b.c/d", "a")
	defer json.Unregister("b.c/d", "@a")
	defer system.Unregister("b.c/d", "a")
	defer system.Unregister("b.c/d", "@a")

	pa := &aRule{
		Object: &system.Object{
			Type: system.NewReference("b.c/d", "@a"),
		},
		Rule: &system.Rule{},
	}
	_, err = Type("n", pa, "kego.io/system", i.Add)
	// This used to throw an error but since we moved to dynamic imports, it
	// should not now.
	assert.NoError(t, err)

}

func TestUnknownRule(t *testing.T) {
	testUnknownRule(t, unmarshalFunc)
	testUnknownRule(t, unpackFunc)
}
func testUnknownRule(t *testing.T, unpacker unpackerFunc) {

	unmarshalDiagram(t, unpacker)
	defer system.Unregister("kego.io/gallery", "diagram")
	defer system.Unregister("kego.io/gallery", "@diagram")

	data := `{
		"description": "This represents a gallery - it's just a list of images",
		"type": "system:type",
		"id": "gallery",
		"fields": {
			"image": {
				"type": "@diagram",
				"default": {
					"type": "diagram",
					"url": "def"
				},
				"optional": true
			},
			"foo": {
				"type": "system:@bool",
				"default": true,
				"optional": true
			},
			"ref": {
				"type": "system:@reference",
				"default": "image",
				"optional": true
			}
		}
	}`

	imp := Imports{}
	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/gallery", map[string]string{})
	assert.Error(t, err)
	assert.EqualError(t, err, "Unknown type kego.io/gallery:diagram")

	f, ok := i.(*system.Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)

	s, err := Type("n", f.Fields["image"], "kego.io/gallery", imp.Add)
	assert.NoError(t, err)
	assert.Equal(t, "*Diagram `kego:\"{\\\"default\\\":{\\\"type\\\":\\\"kego.io/gallery:diagram\\\",\\\"value\\\":{\\\"type\\\":\\\"diagram\\\",\\\"url\\\":\\\"def\\\"}}}\" json:\"n\"`", s)

	b, err := Type("n", f.Fields["foo"], "kego.io/gallery", imp.Add)
	assert.NoError(t, err)
	assert.Equal(t, "system.Bool `kego:\"{\\\"default\\\":{\\\"value\\\":true}}\" json:\"n\"`", b)

	r, err := Type("n", f.Fields["ref"], "kego.io/gallery", imp.Add)
	assert.NoError(t, err)
	assert.Equal(t, "system.Reference `kego:\"{\\\"default\\\":{\\\"type\\\":\\\"kego.io/system:reference\\\",\\\"value\\\":\\\"kego.io/gallery:image\\\"}}\" json:\"n\"`", r)

}

func unmarshalDiagram(t *testing.T, unpacker unpackerFunc) {

	diagram := `{
		"description": "This is a type of image, which just contains the url of the image",
		"type": "system:type",
		"id": "diagram",
		"fields": {
			"url": {
				"type": "system:@string"
			}
		},
		"rule": {
			"description": "Restriction rules for diagram",
			"type": "system:type",
			"is": ["system:rule"],
			"fields": {
				"default": {
					"description": "Default value",
					"type": "@diagram",
					"optional": true
				}
			}
		}
	}`

	var i interface{}
	err := unpacker([]byte(diagram), &i, "kego.io/gallery", map[string]string{})
	assert.NoError(t, err)
	d, ok := i.(*system.Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, d)

	system.Register("kego.io/gallery", d.Id.Name, d, 0)
	if d.Rule != nil {
		d.Rule.Id = system.NewReference(d.Id.Package, fmt.Sprint("@", d.Id.Name))
		system.Register("kego.io/gallery", d.Rule.Id.Name, d.Rule, 0)
	}

}
