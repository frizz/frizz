package system

import (
	"testing"

	"reflect"

	"fmt"

	"github.com/stretchr/testify/assert"
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
		*Object
	}
	type ruleStruct struct {
		*Object
	}
	parentType := &Type{
		Object: &Object{Context: &Context{Package: "a.b/c"}, Id: "a", Type: NewReference("kego.io/system", "type")},
	}
	ruleType := &Type{
		Object: &Object{Context: &Context{Package: "a.b/c"}, Id: "@a", Type: NewReference("kego.io/system", "type")},
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
		rule:       &ruleStruct{},
		ruleType:   ruleType,
		parentType: parentType,
		path:       "d.e/f",
		imports:    map[string]string{},
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
	assert.Error(t, err)
}

type structWithCustomMarshaler struct {
	*Object
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

type typeThatWillCauseJsonMarshalToError chan string

func TestPropertyGetTag(t *testing.T) {
	type parentStruct struct {
		*Object
	}
	type structWithoutCustomMarshaler struct {
		A string
	}
	type ruleStruct struct {
		*Object
		B String
		C *structWithCustomMarshaler
		D typeThatWillCauseJsonMarshalToError
		E structWithoutCustomMarshaler
	}
	parentType := &Type{
		Object: &Object{Context: &Context{Package: "a.b/c"}, Id: "a", Type: NewReference("kego.io/system", "type")},
	}
	ruleType := &Type{
		Object: &Object{Context: &Context{Package: "a.b/c"}, Id: "@a", Type: NewReference("kego.io/system", "type")},
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
		rule:       &ruleStruct{},
		ruleType:   ruleType,
		parentType: parentType,
		path:       "d.e/f",
		imports:    map[string]string{},
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

	r.rule = &ruleStruct{B: NewString("c")}

	s, err = getTag(r)
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"c\\\",\\\"path\\\":\\\"d.e/f\\\"}}\"`", s)

	ruleType.Properties = map[string]*Property{
		"c": &Property{
			Defaulter: true,
		},
	}
	r.rule = &ruleStruct{C: &structWithCustomMarshaler{Object: &Object{Id: "f"}}}
	s, err = getTag(r)
	assert.NoError(t, err)
	assert.Equal(t, "`kego:\"{\\\"default\\\":{\\\"type\\\":\\\"a.b/c:a\\\",\\\"value\\\":\\\"foo\\\",\\\"path\\\":\\\"d.e/f\\\"}}\"`", s)

	r.rule = &ruleStruct{C: &structWithCustomMarshaler{Object: &Object{Id: "f"}, throwError: true}}
	s, err = getTag(r)
	assert.Error(t, err)

	ruleType.Properties = map[string]*Property{
		"d": &Property{
			Defaulter: true,
		},
	}
	r.rule = &ruleStruct{D: make(typeThatWillCauseJsonMarshalToError)}
	s, err = getTag(r)
	assert.Error(t, err)

	ruleType.Properties = map[string]*Property{
		"e": &Property{
			Defaulter: true,
		},
	}
	r.rule = &ruleStruct{E: structWithoutCustomMarshaler{A: "b"}}
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
	var i interface{}
	json.Unmarshal([]byte(`{
			"type": "system:property",
			"item": {
				"type": "json:@string"
			}
		}`), &i, "kego.io/system", map[string]string{})
	p, ok := i.(*Property)
	assert.True(t, ok, "Wrong type")
	s, err := p.GoTypeDescriptor("kego.io/system", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "string", s)
}
