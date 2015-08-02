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
		*system.Base
	}
	type ruleStruct struct {
		*system.Base
		*system.RuleBase
	}
	parentType := &system.Type{
		Base: &system.Base{Id: system.NewReference("a.b/c", "a"), Type: system.NewReference("kego.io/system", "type")},
	}
	ruleType := &system.Type{
		Base: &system.Base{Id: system.NewReference("a.b/c", "@a"), Type: system.NewReference("kego.io/system", "type")},
	}
	json.RegisterType("a.b/c", "a", reflect.TypeOf(&parentStruct{}), 0)
	json.RegisterType("a.b/c", "@a", reflect.TypeOf(&ruleStruct{}), 0)
	system.RegisterType("a.b/c", "a", parentType, 0)
	system.RegisterType("a.b/c", "@a", ruleType, 0)
	defer json.UnregisterType("a.b/c", "a")
	defer json.UnregisterType("a.b/c", "@a")
	defer system.UnregisterType("a.b/c", "a")
	defer system.UnregisterType("a.b/c", "@a")

	r := &system.RuleHolder{
		Rule:       &ruleStruct{},
		RuleType:   ruleType,
		ParentType: parentType,
	}
	s, err := formatTag([]byte("null"), r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", s)

	s, err = formatTag([]byte(`"a"`), r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"a\\\",\\\"path\\\":\\\"d.e/f\\\"}}\"`", s)

	parentType.Id = system.NewReference("kego.io/system", "string")
	s, err = formatTag([]byte(`"a"`), r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"value\\\":\\\"a\\\"}}\"`", s)

	_, err = formatTag([]byte(`foo`), r, "d.e/f", map[string]string{})
	assert.IsError(t, err, "LKBWJTMJCF")
}

type structWithCustomMarshaler struct {
	*system.Base
	throwError bool
}

func (s *structWithCustomMarshaler) MarshalJSON() ([]byte, error) {
	if s.throwError {
		return nil, fmt.Errorf("Here's an error")
	}
	return []byte(`"foo"`), nil
}

func TestMarshaler(t *testing.T) {
	f := structWithCustomMarshaler{}
	s, err := f.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, []byte(`"foo"`), s)
}

// Json marshaling throws an error when fed a chan, so
// it's a convenient way to force an error.
type typeThatWillCauseJsonMarshalToError chan string

func TestGetTag(t *testing.T) {
	type parentStruct struct {
		*system.Base
	}
	type structWithoutCustomMarshaler struct {
		A string
	}
	type ruleStructA struct {
		*system.Base
		*system.RuleBase
		B system.String
	}
	type ruleStructB struct {
		*system.Base
		*system.RuleBase
		Default system.String
	}
	type ruleStructC struct {
		*system.Base
		*system.RuleBase
		Default *structWithCustomMarshaler
	}
	type ruleStructD struct {
		*system.Base
		*system.RuleBase
		Default typeThatWillCauseJsonMarshalToError
	}
	type ruleStructE struct {
		*system.Base
		*system.RuleBase
		Default structWithoutCustomMarshaler
	}
	parentType := &system.Type{
		Base: &system.Base{Id: system.NewReference("a.b/c", "a"), Type: system.NewReference("kego.io/system", "type")},
	}
	ruleType := &system.Type{
		Base: &system.Base{Id: system.NewReference("a.b/c", "@a"), Type: system.NewReference("kego.io/system", "type")},
	}
	json.RegisterType("a.b/c", "@a", reflect.TypeOf(&ruleStructA{}), 0)
	json.RegisterType("a.b/c", "@b", reflect.TypeOf(&ruleStructB{}), 0)
	json.RegisterType("a.b/c", "@c", reflect.TypeOf(&ruleStructC{}), 0)
	json.RegisterType("a.b/c", "@d", reflect.TypeOf(&ruleStructD{}), 0)
	json.RegisterType("a.b/c", "@e", reflect.TypeOf(&ruleStructE{}), 0)
	for _, letter := range []string{"a", "b", "c", "d", "e"} {
		json.RegisterType("a.b/c", letter, reflect.TypeOf(&parentStruct{}), 0)
		system.RegisterType("a.b/c", letter, parentType, 0)
		system.RegisterType("a.b/c", fmt.Sprintf("@%s", letter), ruleType, 0)
		defer json.UnregisterType("a.b/c", letter)
		defer json.UnregisterType("a.b/c", fmt.Sprintf("@%s", letter))
		defer system.UnregisterType("a.b/c", letter)
		defer system.UnregisterType("a.b/c", fmt.Sprintf("@%s", letter))
	}

	r := &system.RuleHolder{
		Rule:       &ruleStructA{},
		RuleType:   ruleType,
		ParentType: parentType,
	}

	// rule has no default field
	s, err := getTag(r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "", s)

	r.Rule = &ruleStructB{Default: system.NewString("c")}
	s, err = getTag(r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"c\\\",\\\"path\\\":\\\"d.e/f\\\"}}\"`", s)

	r.Rule = &ruleStructC{Default: &structWithCustomMarshaler{Base: &system.Base{Id: system.NewReference("d.e/f", "f")}}}
	s, err = getTag(r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"foo\\\",\\\"path\\\":\\\"d.e/f\\\"}}\"`", s)

	r.Rule = &ruleStructC{Default: &structWithCustomMarshaler{Base: &system.Base{Id: system.NewReference("d.e/f", "f")}, throwError: true}}
	s, err = getTag(r, "d.e/f", map[string]string{})
	assert.IsError(t, err, "YIEMHYFVCD")

	r.Rule = &ruleStructD{Default: make(typeThatWillCauseJsonMarshalToError)}
	s, err = getTag(r, "d.e/f", map[string]string{})
	assert.IsError(t, err, "QQDOLAJKLU")

	r.Rule = &ruleStructE{Default: structWithoutCustomMarshaler{A: "b"}}
	s, err = getTag(r, "d.e/f", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":{\\\"A\\\":\\\"b\\\"},\\\"path\\\":\\\"d.e/f\\\"}}\"`", s)

}

func TestGoTypeDescriptor(t *testing.T) {
	p := &system.String_rule{
		Base: &system.Base{
			Type: system.NewReference("kego.io/json", "@string"),
		},
	}
	i := Imports{}
	s, err := Type(p, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "string", s)

	p = &system.String_rule{
		Base: &system.Base{
			Type: system.NewReference("kego.io/system", "@string"),
		},
	}
	s, err = Type(p, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "String", s)

	s, err = Type(p, "kego.io/a", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "system.String", s)

	// We're just using type here because it's a handy
	// non-native type in the system package
	pt := &system.Type_rule{
		Base: &system.Base{
			Type: system.NewReference("kego.io/system", "@type"),
		},
	}
	s, err = Type(pt, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "*Type", s)

	// We're just using property here because it's a handy
	// non-native type in the system package
	pt = &system.Type_rule{
		Base: &system.Base{
			Type: system.NewReference("kego.io/system", "@type"),
		},
	}
	s, err = Type(pt, "kego.io/a", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "*system.Type", s)

	type a struct{}
	type a_rule struct {
		*system.Base
		*system.RuleBase
	}
	tyr := &system.Type{Base: &system.Base{Id: system.NewReference("b.c/d", "@a")}}
	ty := &system.Type{Base: &system.Base{Id: system.NewReference("b.c/d", "a")}}
	json.RegisterType("b.c/d", "a", reflect.TypeOf(&a{}), 0)
	json.RegisterType("b.c/d", "@a", reflect.TypeOf(&a_rule{}), 0)
	system.RegisterType("b.c/d", "a", ty, 0)
	system.RegisterType("b.c/d", "@a", tyr, 0)
	defer json.UnregisterType("b.c/d", "a")
	defer json.UnregisterType("b.c/d", "@a")
	defer system.UnregisterType("b.c/d", "a")
	defer system.UnregisterType("b.c/d", "@a")

	pa := &a_rule{
		Base: &system.Base{
			Type: system.NewReference("b.c/d", "@a"),
		},
	}
	ia := Imports{"b.c/d": Import{Path: "b.c/d", Name: "d", Alias: "d"}}
	s, err = Type(pa, "kego.io/system", ia.Add)
	assert.NoError(t, err)
	assert.Equal(t, "*d.A", s)

	s, err = Type(pa, "b.c/d", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "*A", s)

	p = &system.String_rule{
		Base: &system.Base{
			Type: system.NewReference("kego.io/system", "@string"),
		},
		Default: system.NewString("a"),
	}
	s, err = Type(p, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "String `kego:\"{\\\"default\\\":{\\\"value\\\":\\\"a\\\"}}\"`", s)

	pm := &system.Map_rule{
		Base: &system.Base{
			Type: system.NewReference("kego.io/system", "@map"),
		},
		Items: &system.String_rule{
			Base: &system.Base{
				Type: system.NewReference("kego.io/system", "@string"),
			},
		},
	}
	s, err = Type(pm, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "map[string]String", s)

	par := &system.Array_rule{
		Base: &system.Base{
			Type: system.NewReference("kego.io/system", "@array"),
		},
		Items: &system.String_rule{
			Base: &system.Base{
				Type: system.NewReference("kego.io/system", "@string"),
			},
		},
	}
	s, err = Type(par, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "[]String", s)

	pm = &system.Map_rule{
		Base: &system.Base{
			Type: system.NewReference("kego.io/system", "@map"),
		},
		Items: &system.Array_rule{
			Base: &system.Base{
				Type: system.NewReference("kego.io/system", "@array"),
			},
			Items: &system.String_rule{
				Base: &system.Base{
					Type: system.NewReference("kego.io/system", "@string"),
				},
			},
		},
	}
	s, err = Type(pm, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "map[string][]String", s)

	pm = &system.Map_rule{
		Base: &system.Base{
			Type: system.NewReference("kego.io/system", "@map"),
		},
		Items: &system.Array_rule{
			Base: &system.Base{
				Type: system.NewReference("kego.io/system", "@array"),
			},
			Items: &system.String_rule{
				Base: &system.Base{
					Type: system.NewReference("kego.io/system", "@string"),
				},
				Default: system.NewString("a"),
			},
		},
	}
	s, err = Type(pm, "kego.io/system", i.Add)
	assert.NoError(t, err)
	assert.Equal(t, "map[string][]String `kego:\"{\\\"default\\\":{\\\"value\\\":\\\"a\\\"}}\"`", s)

}

func TypeErrors_NeedsTypes(t *testing.T) {

	p := &system.JsonString_rule{
		Base: &system.Base{
			Type: system.NewReference("a.b/c", "notFoundType"),
		},
	}
	i := Imports{}
	_, err := Type(p, "kego.io/system", i.Add)
	// Item is an unregistered type, so errors at NewRuleHolder
	assert.IsError(t, err, "TFXFBIRXHN")
	assert.HasError(t, err, "PFGWISOHRR")

	pm := &system.Map_rule{
		Base: &system.Base{
			Type: system.NewReference("kego.io/system", "@map"),
		},
	}
	_, err = Type(pm, "kego.io/system", i.Add)
	// Collection item @map doesn't have Items field, so errors at collectionPrefixInnerRule
	assert.IsError(t, err, "SOGEFOPJHB")
	assert.HasError(t, err, "VYTHGJTSNJ")

	type a struct{}
	type a_rule struct {
		*system.Base
		*system.RuleBase
	}
	tyr := &system.Type{Base: &system.Base{Id: system.NewReference("b.c/d", "@a")}}
	ty := &system.Type{Base: &system.Base{Id: system.NewReference("b.c/d", "a")}}
	json.RegisterType("b.c/d", "a", reflect.TypeOf(&a{}), 0)
	json.RegisterType("b.c/d", "@a", reflect.TypeOf(&a_rule{}), 0)
	system.RegisterType("b.c/d", "a", ty, 0)
	system.RegisterType("b.c/d", "@a", tyr, 0)
	defer json.UnregisterType("b.c/d", "a")
	defer json.UnregisterType("b.c/d", "@a")
	defer system.UnregisterType("b.c/d", "a")
	defer system.UnregisterType("b.c/d", "@a")

	pa := &a_rule{
		Base: &system.Base{
			Type: system.NewReference("b.c/d", "@a"),
		},
	}
	_, err = Type(pa, "kego.io/system", i.Add)
	// This used to throw an error but since we moved to dynamic imports, it
	// should not now.
	assert.NoError(t, err)

}

func unmarshalDiagram(t *testing.T) {
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
			"id": "@diagram",
			"is": ["system:rule"],
			"embed": ["system:ruleBase"],
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
	err := json.Unmarshal([]byte(diagram), &i, "kego.io/gallery", map[string]string{})
	assert.NoError(t, err)
	d, ok := i.(*system.Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, d)

	system.RegisterType("kego.io/gallery", d.Id.Name, d, 0)
	if d.Rule != nil {
		system.RegisterType("kego.io/gallery", d.Rule.Id.Name, d.Rule, 0)
	}

}

func TestUnknownRule(t *testing.T) {

	unmarshalDiagram(t)
	defer system.UnregisterType("kego.io/gallery", "diagram")
	defer system.UnregisterType("kego.io/gallery", "@diagram")

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
	err := json.Unmarshal([]byte(data), &i, "kego.io/gallery", map[string]string{})
	assert.EqualError(t, err, "Unknown type kego.io/gallery:diagram")
	f, ok := i.(*system.Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)

	s, err := Type(f.Fields["image"], "kego.io/gallery", imp.Add)
	assert.NoError(t, err)
	assert.Equal(t, "*Diagram `kego:\"{\\\"default\\\":{\\\"type\\\":\\\"kego.io/gallery:diagram\\\",\\\"value\\\":{\\\"type\\\":\\\"diagram\\\",\\\"url\\\":\\\"def\\\"}}}\"`", s)

	b, err := Type(f.Fields["foo"], "kego.io/gallery", imp.Add)
	assert.NoError(t, err)
	assert.Equal(t, "system.Bool `kego:\"{\\\"default\\\":{\\\"value\\\":true}}\"`", b)

	r, err := Type(f.Fields["ref"], "kego.io/gallery", imp.Add)
	assert.NoError(t, err)
	assert.Equal(t, "system.Reference `kego:\"{\\\"default\\\":{\\\"type\\\":\\\"kego.io/system:reference\\\",\\\"value\\\":\\\"kego.io/gallery:image\\\"}}\"`", r)

}
