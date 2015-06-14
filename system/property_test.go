package system

import (
	"testing"

	"reflect"

	"fmt"

	"kego.io/assert"
	"kego.io/json"
)

func TestPropertyGetPointer(t *testing.T) {
	//!isNative && !isInterface => "*"
	assert.Equal(t, "", getPointer(&Type{Native: NewString("bool"), Interface: false}))
	assert.Equal(t, "", getPointer(&Type{Native: NewString("object"), Interface: true}))
	assert.Equal(t, "", getPointer(&Type{Native: NewString("bool"), Interface: true}))
	assert.Equal(t, "*", getPointer(&Type{Native: NewString("object"), Interface: false}))
	assert.Equal(t, "*", getPointer(&Type{Native: NewString("map"), Interface: false}))
	assert.Equal(t, "*", getPointer(&Type{Native: NewString("array"), Interface: false}))
}

func TestPropertyFormatTag(t *testing.T) {
	type parentStruct struct {
		*Base
	}
	type ruleStruct struct {
		*Base
		*RuleBase
	}
	parentType := &Type{
		Base: &Base{Context: &Context{Package: "a.b/c"}, Id: "a", Type: NewReference("kego.io/system", "type")},
	}
	ruleType := &Type{
		Base: &Base{Context: &Context{Package: "a.b/c"}, Id: "@a", Type: NewReference("kego.io/system", "type")},
	}
	json.RegisterType("a.b/c:a", reflect.TypeOf(&parentStruct{}))
	json.RegisterType("a.b/c:@a", reflect.TypeOf(&ruleStruct{}))
	RegisterType("a.b/c:a", parentType)
	RegisterType("a.b/c:@a", ruleType)
	defer json.UnregisterType("a.b/c:a")
	defer json.UnregisterType("a.b/c:@a")
	defer UnregisterType("a.b/c:a")
	defer UnregisterType("a.b/c:@a")

	r := &RuleHolder{
		Rule:       &ruleStruct{},
		RuleType:   ruleType,
		ParentType: parentType,
		Path:       "d.e/f",
		Imports:    map[string]string{},
	}
	s, err := formatTag([]byte("null"), r)
	assert.NoError(t, err)
	assert.Equal(t, "", s)

	s, err = formatTag([]byte(`"a"`), r)
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"a\\\",\\\"path\\\":\\\"d.e/f\\\"}}\"`", s)

	parentType.Context.Package = "kego.io/system"
	parentType.Id = "string"
	s, err = formatTag([]byte(`"a"`), r)
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"value\\\":\\\"a\\\"}}\"`", s)

	_, err = formatTag([]byte(`foo`), r)
	assert.IsError(t, err, "LKBWJTMJCF")
}

type structWithCustomMarshaler struct {
	*Base
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

func TestPropertyGetTag(t *testing.T) {
	type parentStruct struct {
		*Base
	}
	type structWithoutCustomMarshaler struct {
		A string
	}
	type ruleStruct struct {
		*Base
		*RuleBase
		B String
		C *structWithCustomMarshaler
		D typeThatWillCauseJsonMarshalToError
		E structWithoutCustomMarshaler
	}
	parentType := &Type{
		Base: &Base{Context: &Context{Package: "a.b/c"}, Id: "a", Type: NewReference("kego.io/system", "type")},
	}
	ruleType := &Type{
		Base: &Base{Context: &Context{Package: "a.b/c"}, Id: "@a", Type: NewReference("kego.io/system", "type")},
	}
	json.RegisterType("a.b/c:a", reflect.TypeOf(&parentStruct{}))
	json.RegisterType("a.b/c:@a", reflect.TypeOf(&ruleStruct{}))
	RegisterType("a.b/c:a", parentType)
	RegisterType("a.b/c:@a", ruleType)
	defer json.UnregisterType("a.b/c:a")
	defer json.UnregisterType("a.b/c:@a")
	defer UnregisterType("a.b/c:a")
	defer UnregisterType("a.b/c:@a")

	r := &RuleHolder{
		Rule:       &ruleStruct{},
		RuleType:   ruleType,
		ParentType: parentType,
		Path:       "d.e/f",
		Imports:    map[string]string{},
	}

	// ruleType has no defaulter property
	s, err := getTag(r)
	assert.NoError(t, err)
	assert.Equal(t, "", s)

	ruleType.Properties = map[string]*Property{
		"b": &Property{
			Defaulter: true,
		},
	}

	// rule has no default field
	s, err = getTag(r)
	assert.NoError(t, err)
	assert.Equal(t, "", s)

	/*
		r.Rule = map[string]interface{}{"b": "c"}
		s, err = getTag(r)
		assert.NoError(t, err)
		assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"c\\\",\\\"path\\\":\\\"d.e/f\\\"}}\"`", s)

		r.Rule = map[string]interface{}{"d": "e"}
		s, err = getTag(r)
		assert.NoError(t, err)
		assert.Equal(t, "", s)

		r.Rule = map[string]interface{}{"b": make(typeThatWillCauseJsonMarshalToError)}
		s, err = getTag(r)
		assert.IsError(t, err, "FYMGUTAOCR")
	*/

	r.Rule = &ruleStruct{B: NewString("c")}

	s, err = getTag(r)
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"c\\\",\\\"path\\\":\\\"d.e/f\\\"}}\"`", s)

	ruleType.Properties = map[string]*Property{
		"c": &Property{
			Defaulter: true,
		},
	}
	r.Rule = &ruleStruct{C: &structWithCustomMarshaler{Base: &Base{Id: "f"}}}
	s, err = getTag(r)
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"foo\\\",\\\"path\\\":\\\"d.e/f\\\"}}\"`", s)

	r.Rule = &ruleStruct{C: &structWithCustomMarshaler{Base: &Base{Id: "f"}, throwError: true}}
	s, err = getTag(r)
	assert.IsError(t, err, "YIEMHYFVCD")

	ruleType.Properties = map[string]*Property{
		"d": &Property{
			Defaulter: true,
		},
	}
	r.Rule = &ruleStruct{D: make(typeThatWillCauseJsonMarshalToError)}
	s, err = getTag(r)
	assert.IsError(t, err, "QQDOLAJKLU")

	ruleType.Properties = map[string]*Property{
		"e": &Property{
			Defaulter: true,
		},
	}
	r.Rule = &ruleStruct{E: structWithoutCustomMarshaler{A: "b"}}
	s, err = getTag(r)
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":{\\\"A\\\":\\\"b\\\"},\\\"path\\\":\\\"d.e/f\\\"}}\"`", s)

}

func TestGoName(t *testing.T) {
	p := &Property{}
	s := p.GoName("foo")
	assert.Equal(t, "Foo", s)
	s = p.GoName("@foo")
	assert.Equal(t, "Foo_rule", s)
}

func TestGoTypeDescriptor(t *testing.T) {
	p := &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		Item: &JsonString_rule{
			Base: &Base{
				Type: NewReference("kego.io/json", "@string"),
			},
		},
	}
	s, err := p.GoTypeDescriptor("kego.io/system", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "string", s)

	p = &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		Item: &String_rule{
			Base: &Base{
				Type: NewReference("kego.io/system", "@string"),
			},
		},
	}
	s, err = p.GoTypeDescriptor("kego.io/system", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "String", s)

	s, err = p.GoTypeDescriptor("kego.io/a", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "system.String", s)

	p = &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		// We're just using property here because it's a handy
		// non-native type in the system package
		Item: &Property_rule{
			Base: &Base{
				Type: NewReference("kego.io/system", "@property"),
			},
		},
	}
	s, err = p.GoTypeDescriptor("kego.io/system", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "*Property", s)

	p = &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		// We're just using property here because it's a handy
		// non-native type in the system package
		Item: &Property_rule{
			Base: &Base{
				Type: NewReference("kego.io/system", "@property"),
			},
		},
	}
	s, err = p.GoTypeDescriptor("kego.io/a", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "*system.Property", s)

	type a struct{}
	type a_rule struct {
		*Base
		*RuleBase
	}
	tyr := &Type{Base: &Base{Id: "@a"}}
	ty := &Type{Base: &Base{Id: "a", Context: &Context{Package: "b.c/d"}}}
	json.RegisterType("b.c/d:a", reflect.TypeOf(&a{}))
	json.RegisterType("b.c/d:@a", reflect.TypeOf(&a_rule{}))
	RegisterType("b.c/d:a", ty)
	RegisterType("b.c/d:@a", tyr)
	defer json.UnregisterType("b.c/d:a")
	defer json.UnregisterType("b.c/d:@a")
	defer UnregisterType("b.c/d:a")
	defer UnregisterType("b.c/d:@a")

	p = &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		Item: &a_rule{
			Base: &Base{
				Type: NewReference("b.c/d", "@a"),
			},
		},
	}
	s, err = p.GoTypeDescriptor("kego.io/system", map[string]string{
		"d": "b.c/d",
	})
	assert.NoError(t, err)
	assert.Equal(t, "*d.A", s)

	s, err = p.GoTypeDescriptor("b.c/d", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "*A", s)

	p = &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		Item: &String_rule{
			Base: &Base{
				Type: NewReference("kego.io/system", "@string"),
			},
			Default: NewString("a"),
		},
	}
	s, err = p.GoTypeDescriptor("kego.io/system", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "String `kego:\"{\\\"default\\\":{\\\"value\\\":\\\"a\\\"}}\"`", s)

	p = &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		Item: &Map_rule{
			Base: &Base{
				Type: NewReference("kego.io/system", "@map"),
			},
			Items: &String_rule{
				Base: &Base{
					Type: NewReference("kego.io/system", "@string"),
				},
			},
		},
	}
	s, err = p.GoTypeDescriptor("kego.io/system", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "map[string]String", s)

	p = &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		Item: &Array_rule{
			Base: &Base{
				Type: NewReference("kego.io/system", "@array"),
			},
			Items: &String_rule{
				Base: &Base{
					Type: NewReference("kego.io/system", "@string"),
				},
			},
		},
	}
	s, err = p.GoTypeDescriptor("kego.io/system", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "[]String", s)

	p = &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		Item: &Map_rule{
			Base: &Base{
				Type: NewReference("kego.io/system", "@map"),
			},
			Items: &Array_rule{
				Base: &Base{
					Type: NewReference("kego.io/system", "@array"),
				},
				Items: &String_rule{
					Base: &Base{
						Type: NewReference("kego.io/system", "@string"),
					},
				},
			},
		},
	}
	s, err = p.GoTypeDescriptor("kego.io/system", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "map[string][]String", s)

	p = &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		Item: &Map_rule{
			Base: &Base{
				Type: NewReference("kego.io/system", "@map"),
			},
			Items: &Array_rule{
				Base: &Base{
					Type: NewReference("kego.io/system", "@array"),
				},
				Items: &String_rule{
					Base: &Base{
						Type: NewReference("kego.io/system", "@string"),
					},
					Default: NewString("a"),
				},
			},
		},
	}
	s, err = p.GoTypeDescriptor("kego.io/system", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "map[string][]String `kego:\"{\\\"default\\\":{\\\"value\\\":\\\"a\\\"}}\"`", s)

}

func GoTypeDescriptorErrors_NeedsTypes(t *testing.T) {

	p := &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		Item: &JsonString_rule{
			Base: &Base{
				Type: NewReference("a.b/c", "notFoundType"),
			},
		},
	}
	_, err := p.GoTypeDescriptor("kego.io/system", map[string]string{})
	// Item is an unregistered type, so errors at NewRuleHolder
	assert.IsError(t, err, "QKCXSRPMOQ")

	p = &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		Item: &Map_rule{
			Base: &Base{
				Type: NewReference("kego.io/system", "@map"),
			},
		},
	}
	_, err = p.GoTypeDescriptor("kego.io/system", map[string]string{})
	// Collection item @map doesn't have Items field, so errors at collectionPrefixInnerRule
	assert.IsError(t, err, "SNATGPVLAS")

	type a struct{}
	type a_rule struct {
		*Base
		*RuleBase
	}
	tyr := &Type{Base: &Base{Id: "@a"}}
	ty := &Type{Base: &Base{Id: "a", Context: &Context{Package: "b.c/d"}}}
	json.RegisterType("b.c/d:a", reflect.TypeOf(&a{}))
	json.RegisterType("b.c/d:@a", reflect.TypeOf(&a_rule{}))
	RegisterType("b.c/d:a", ty)
	RegisterType("b.c/d:@a", tyr)
	defer json.UnregisterType("b.c/d:a")
	defer json.UnregisterType("b.c/d:@a")
	defer UnregisterType("b.c/d:a")
	defer UnregisterType("b.c/d:@a")

	p = &Property{
		Base: &Base{
			Type: NewReference("kego.io/system", "property"),
		},
		Item: &a_rule{
			Base: &Base{
				Type: NewReference("b.c/d", "@a"),
			},
		},
	}
	_, err = p.GoTypeDescriptor("kego.io/system", map[string]string{})
	// Item a_rule is a valid type but package b.c/d is not in the final
	// imports specified to GoTypeDescriptor, so we get an error at
	// inner.parentType.GoTypeReference
	assert.IsError(t, err, "CGNYBDFUNP")

	/*
		p = &Property{
			Base: &Base{
				Type: NewReference("kego.io/system", "property"),
			},
			Item: map[string]interface{}{
				"type":     "kego.io/system:@string",
				"default":  make(typeThatWillCauseJsonMarshalToError),
				"_path":    "kego.io/system",
				"_imports": map[string]string{},
			},
		}
		_, err = p.GoTypeDescriptor("kego.io/system", map[string]string{})
		// Item default is a chan, so errors at getTag
		assert.IsError(t, err, "LKIXCKRCYG")
	*/
}
