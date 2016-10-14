package builder

import (
	"reflect"
	"testing"

	"kego.io/process/parser"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/system"
	"kego.io/tests"
)

/*
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
	ctx := tests.Context("d.e/f").
		JtypePathRule("a.b/c", "a", reflect.TypeOf(&parentStruct{}), reflect.TypeOf(&ruleStruct{})).
		StypePath("a.b/c", "a", parentType).
		StypePath("a.b/c", "@a", ruleType).Ctx()

	r := &system.RuleWrapper{
		Interface: &ruleStruct{},
		Parent:    parentType,
	}
	s, err := formatTag(ctx, "n", []byte("null"), r)
	require.NoError(t, err)
	assert.Equal(t, "`json:\"n\"`", s)

	s, err = formatTag(ctx, "", nil, nil)
	require.NoError(t, err)
	assert.Equal(t, "", s)

	s, err = formatTag(ctx, "`", []byte("null"), r)
	require.NoError(t, err)
	assert.Equal(t, "\"json:\\\"`\\\"\"", s)

	s, err = formatTag(ctx, "n", []byte(`"a"`), r)
	require.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"a\\\",\\\"path\\\":\\\"d.e/f\\\"}}\" json:\"n\"`", s)

	parentType.Id = system.NewReference("kego.io/system", "string")
	s, err = formatTag(ctx, "n", []byte(`"a"`), r)
	require.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"value\\\":\\\"a\\\"}}\" json:\"n\"`", s)

	_, err = formatTag(ctx, "n", []byte(`foo`), r)
	assert.IsError(t, err, "LKBWJTMJCF")
}


type structWithCustomMarshaler struct {
	*system.Object
	throwError bool
}

func (s *structWithCustomMarshaler) MarshalJSON(ctx context.Context) ([]byte, error) {
	if s.throwError {
		return nil, fmt.Errorf("Here's an error")
	}
	return []byte(`"foo"`), nil
}

var _ packer.Marshaler = (*structWithCustomMarshaler)(nil)

func TestMarshaler(t *testing.T) {
	f := structWithCustomMarshaler{}
	s, err := f.MarshalJSON(envctx.Empty)
	require.NoError(t, err)
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
	B *system.String
}
type ruleStructB struct {
	*system.Object
	*system.Rule
	Default *system.String
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

type ruleStructF struct {
	*system.Object
	*system.Rule
}

func (r *ruleStructF) GetDefault() interface{} {
	return nil
}

var _ system.DefaultRule = (*ruleStructF)(nil)

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

	ps := reflect.TypeOf(&parentStruct{})
	ctx := tests.Context("d.e/f").
		AllPath("a.b/c", "a", ps, reflect.TypeOf(&ruleStructA{}), parentType, ruleType).
		AllPath("a.b/c", "b", ps, reflect.TypeOf(&ruleStructB{}), parentType, ruleType).
		AllPath("a.b/c", "c", ps, reflect.TypeOf(&ruleStructC{}), parentType, ruleType).
		AllPath("a.b/c", "d", ps, reflect.TypeOf(&ruleStructD{}), parentType, ruleType).
		AllPath("a.b/c", "e", ps, reflect.TypeOf(&ruleStructE{}), parentType, ruleType).
		Ctx()

	r := &system.RuleWrapper{
		Interface: &ruleStructA{},
		Parent:    parentType,
	}

	// rule has no default field
	s, err := getTag(ctx, "n", r)
	require.NoError(t, err)
	assert.Equal(t, "`json:\"n\"`", s)

	r.Interface = &ruleStructB{Default: system.NewString("c")}
	s, err = getTag(ctx, "n", r)
	require.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"c\\\",\\\"path\\\":\\\"d.e/f\\\"}}\" json:\"n\"`", s)

	r.Interface = &ruleStructC{Default: &structWithCustomMarshaler{Object: &system.Object{Id: system.NewReference("d.e/f", "f")}}}
	s, err = getTag(ctx, "n", r)
	require.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"foo\\\",\\\"path\\\":\\\"d.e/f\\\"}}\" json:\"n\"`", s)

	r.Interface = &ruleStructC{Default: &structWithCustomMarshaler{Object: &system.Object{Id: system.NewReference("d.e/f", "f")}, throwError: true}}
	s, err = getTag(ctx, "n", r)
	assert.IsError(t, err, "YIEMHYFVCD")

	r.Interface = &ruleStructD{Default: make(typeThatWillCauseJsonMarshalToError)}
	s, err = getTag(ctx, "n", r)
	assert.IsError(t, err, "QQDOLAJKLU")

	r.Interface = &ruleStructE{Default: structWithoutCustomMarshaler{A: "b"}}
	s, err = getTag(ctx, "n", r)
	require.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":{\\\"A\\\":\\\"b\\\"},\\\"path\\\":\\\"d.e/f\\\"}}\" json:\"n\"`", s)

	r.Interface = &ruleStructF{}
	s, err = getTag(ctx, "n", r)
	require.NoError(t, err)
	assert.Equal(t, "`json:\"n\"`", s)

}
*/

func TestGoTypeDescriptor(t *testing.T) {

	cb := tests.Context("kego.io/system").Jauto().Sauto(parser.Parse)

	p := &system.StringRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/json", "@string"),
		},
		Rule: &system.Rule{},
	}
	i := Imports{}
	s, err := FieldTypeDefinition(cb.Ctx(), "n", p, "kego.io/system", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "string `json:\"n\"`", s)

	p = &system.StringRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/json", "@string"),
		},
		Rule: &system.Rule{
			Interface: true,
		},
	}
	s, err = FieldTypeDefinition(cb.Ctx(), "n", p, "kego.io/system", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "json.StringInterface `json:\"n\"`", s)

	p = &system.StringRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@string"),
		},
		Rule: &system.Rule{},
	}
	s, err = FieldTypeDefinition(cb.Ctx(), "n", p, "kego.io/system", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "*String `json:\"n\"`", s)

	cb.Path("kego.io/a")

	s, err = FieldTypeDefinition(cb.Ctx(), "n", p, "kego.io/a", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "*system.String `json:\"n\"`", s)

	cb.Path("kego.io/system")

	// We're just using type here because it's a handy
	// non-native type in the system package
	pt := &system.TypeRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@type"),
		},
		Rule: &system.Rule{},
	}
	s, err = FieldTypeDefinition(cb.Ctx(), "n", pt, "kego.io/system", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "*Type `json:\"n\"`", s)

	cb.Path("kego.io/a")

	// We're just using property here because it's a handy
	// non-native type in the system package
	pt = &system.TypeRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@type"),
		},
		Rule: &system.Rule{},
	}
	s, err = FieldTypeDefinition(cb.Ctx(), "n", pt, "kego.io/a", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "*system.Type `json:\"n\"`", s)

	cb.Path("kego.io/system")

	type a struct{}
	type aRule struct {
		*system.Object
		*system.Rule
	}
	tyr := &system.Type{Object: &system.Object{Id: system.NewReference("b.c/d", "@a")}, Native: system.NewString("object")}
	ty := &system.Type{Object: &system.Object{Id: system.NewReference("b.c/d", "a")}, Native: system.NewString("object")}

	cb.AllPath("b.c/d", "a", reflect.TypeOf(&a{}), reflect.TypeOf(&aRule{}), ty, tyr)

	pa := &aRule{
		Object: &system.Object{
			Type: system.NewReference("b.c/d", "@a"),
		},
		Rule: &system.Rule{},
	}
	ia := Imports{"b.c/d": Import{Path: "b.c/d", Name: "d", Alias: "d"}}
	s, err = FieldTypeDefinition(cb.Ctx(), "n", pa, "kego.io/system", ia.Add)
	require.NoError(t, err)
	assert.Equal(t, "*d.A `json:\"n\"`", s)

	p = &system.StringRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@string"),
		},
		Rule:    &system.Rule{},
		Default: system.NewString("a"),
	}
	s, err = FieldTypeDefinition(cb.Ctx(), "n", p, "kego.io/system", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "*String `json:\"n\"`", s)

	pm := &system.MapRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@map"),
		},
		Rule: &system.Rule{},
		Items: &system.StringRule{
			Object: &system.Object{
				Type: system.NewReference("kego.io/system", "@string"),
			},
			Rule: &system.Rule{},
		},
	}
	s, err = FieldTypeDefinition(cb.Ctx(), "n", pm, "kego.io/system", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "map[string]*String `json:\"n\"`", s)

	par := &system.ArrayRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@array"),
		},
		Rule: &system.Rule{},
		Items: &system.StringRule{
			Object: &system.Object{
				Type: system.NewReference("kego.io/system", "@string"),
			},
			Rule: &system.Rule{},
		},
	}
	s, err = FieldTypeDefinition(cb.Ctx(), "n", par, "kego.io/system", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "[]*String `json:\"n\"`", s)

	pm = &system.MapRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@map"),
		},
		Rule: &system.Rule{},
		Items: &system.ArrayRule{
			Object: &system.Object{
				Type: system.NewReference("kego.io/system", "@array"),
			},
			Rule: &system.Rule{},
			Items: &system.StringRule{
				Object: &system.Object{
					Type: system.NewReference("kego.io/system", "@string"),
				},
				Rule: &system.Rule{},
			},
		},
	}
	s, err = FieldTypeDefinition(cb.Ctx(), "n", pm, "kego.io/system", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "map[string][]*String `json:\"n\"`", s)

	pm = &system.MapRule{
		Object: &system.Object{
			Type: system.NewReference("kego.io/system", "@map"),
		},
		Rule: &system.Rule{},
		Items: &system.ArrayRule{
			Object: &system.Object{
				Type: system.NewReference("kego.io/system", "@array"),
			},
			Rule: &system.Rule{},
			Items: &system.StringRule{
				Object: &system.Object{
					Type: system.NewReference("kego.io/system", "@string"),
				},
				Rule:    &system.Rule{},
				Default: system.NewString("a"),
			},
		},
	}
	s, err = FieldTypeDefinition(cb.Ctx(), "n", pm, "kego.io/system", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "map[string][]*String `json:\"n\"`", s)

	cb.Path("b.c/d")

	s, err = FieldTypeDefinition(cb.Ctx(), "n", pa, "b.c/d", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "*A `json:\"n\"`", s)

	dr := &system.DummyRule{
		Object: &system.Object{
			Type: system.NewReference("b.c/d", "@a"),
		},
		Rule: &system.Rule{},
	}
	s, err = FieldTypeDefinition(cb.Ctx(), "n", dr, "b.c/d", i.Add)
	require.NoError(t, err)
	assert.Equal(t, "*A `json:\"n\"`", s)

}

func TypeErrors_NeedsTypes(t *testing.T) {

	cb := tests.Context("kego.io/system").Jauto().Sauto(parser.Parse)

	i := Imports{}

	type a struct{}
	type aRule struct {
		*system.Object
		*system.Rule
	}
	tyr := &system.Type{Object: &system.Object{Id: system.NewReference("b.c/d", "@a")}, Native: system.NewString("object")}
	ty := &system.Type{Object: &system.Object{Id: system.NewReference("b.c/d", "a")}, Native: system.NewString("object")}

	cb.AllPath("b.c/d", "a", reflect.TypeOf(&a{}), reflect.TypeOf(&aRule{}), ty, tyr)

	pa := &aRule{
		Object: &system.Object{
			Type: system.NewReference("b.c/d", "@a"),
		},
		Rule: &system.Rule{},
	}
	_, err := FieldTypeDefinition(cb.Ctx(), "n", pa, "kego.io/system", i.Add)
	// This used to throw an error but since we moved to dynamic imports, it
	// should not now.
	require.NoError(t, err)

}

/*
func TestUnknownRule(t *testing.T) {
	testUnknownRule(t, unpacker.Unmarshal)
	testUnknownRule(t, unpacker.Unpack)
}
func testUnknownRule(t *testing.T, unpacker unpacker.Interface) {

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
			"embed": ["system:rule"],
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
	err := unpacker(tests.PathCtx("kego.io/gallery"), []byte(diagram), &i)
	require.NoError(t, err)
	d, ok := i.(*system.Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, d)

	system.Register("kego.io/gallery", d.Id.Name, d, 0)
	d.Rule.Id = system.NewReference(d.Id.Package, fmt.Sprint("@", d.Id.Name))
	system.Register("kego.io/gallery", d.Rule.Id.Name, d.Rule, 0)

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
	err := unpacker(tests.PathCtx("kego.io/gallery"), []byte(data), &i)
	assert.Error(t, err)
	assert.EqualError(t, err, "Unknown type kego.io/gallery:diagram")

	f, ok := i.(*system.Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)

	s, err := Type(tests.PathCtx("kego.io/gallery"), "n", f.Fields["image"], "kego.io/gallery", imp.Add)
	require.NoError(t, err)
	assert.Equal(t, "*Diagram `kego:\"{\\\"default\\\":{\\\"type\\\":\\\"kego.io/gallery:diagram\\\",\\\"value\\\":{\\\"type\\\":\\\"diagram\\\",\\\"url\\\":\\\"def\\\"}}}\" json:\"n\"`", s)

	b, err := Type(tests.PathCtx("kego.io/gallery"), "n", f.Fields["foo"], "kego.io/gallery", imp.Add)
	require.NoError(t, err)
	assert.Equal(t, "*system.Bool `kego:\"{\\\"default\\\":{\\\"value\\\":true}}\" json:\"n\"`", b)

	r, err := Type(tests.PathCtx("kego.io/gallery"), "n", f.Fields["ref"], "kego.io/gallery", imp.Add)
	require.NoError(t, err)
	assert.Equal(t, "*system.Reference `kego:\"{\\\"default\\\":{\\\"type\\\":\\\"kego.io/system:reference\\\",\\\"value\\\":\\\"kego.io/gallery:image\\\"}}\" json:\"n\"`", r)

}
*/
