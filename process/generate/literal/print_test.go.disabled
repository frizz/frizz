package literal

import (
	"reflect"
	"testing"

	"github.com/davelondon/ktest/assert"
	"kego.io/process/generate/builder"
)

func TestGetName(t *testing.T) {

	imp := builder.Imports{"e.f/g": builder.Import{Path: "e.f/g", Name: "g", Alias: "d"}}

	s := ""
	b := true
	i := 1
	f := 1.2

	assert.Equal(t, "string", GetName(reflect.TypeOf(s), "a.b/c", imp.Add))
	assert.Equal(t, "bool", GetName(reflect.TypeOf(b), "a.b/c", imp.Add))
	assert.Equal(t, "int", GetName(reflect.TypeOf(i), "a.b/c", imp.Add))
	assert.Equal(t, "float64", GetName(reflect.TypeOf(f), "a.b/c", imp.Add))

	assert.Equal(t, "*string", GetName(reflect.TypeOf(&s), "a.b/c", imp.Add))
	assert.Equal(t, "*bool", GetName(reflect.TypeOf(&b), "a.b/c", imp.Add))
	assert.Equal(t, "*int", GetName(reflect.TypeOf(&i), "a.b/c", imp.Add))
	assert.Equal(t, "*float64", GetName(reflect.TypeOf(&f), "a.b/c", imp.Add))

	impk := builder.Imports{
		"e.f/g": builder.Import{Path: "e.f/g", Name: "g", Alias: "d"},
		"kego.io/process/generate/literal": builder.Import{Path: "kego.io/process/generate/literal", Name: "literal", Alias: "h"}}
	type MyType struct{}
	m := MyType{}

	assert.Equal(t, "MyType", GetName(reflect.TypeOf(m), "kego.io/process/generate/literal", imp.Add))
	assert.Equal(t, "h.MyType", GetName(reflect.TypeOf(m), "a.b/c", impk.Add))
	assert.Equal(t, "*MyType", GetName(reflect.TypeOf(&m), "kego.io/process/generate/literal", imp.Add))
	assert.Equal(t, "*h.MyType", GetName(reflect.TypeOf(&m), "a.b/c", impk.Add))

	asp := "github.com/davelondon/ktest/assert"
	impa := builder.Imports{
		"e.f/g": builder.Import{Path: "e.f/g", Name: "g", Alias: "d"},
		asp:     builder.Import{Path: asp, Name: "assert", Alias: "as"},
	}
	ass := assert.Assertions{} // Just using a random struct from a non system package

	assert.Equal(t, "Assertions", GetName(reflect.TypeOf(ass), asp, imp.Add))
	assert.Equal(t, "as.Assertions", GetName(reflect.TypeOf(ass), "a.b/c", impa.Add))
	assert.Equal(t, "*Assertions", GetName(reflect.TypeOf(&ass), asp, imp.Add))
	assert.Equal(t, "*as.Assertions", GetName(reflect.TypeOf(&ass), "a.b/c", impa.Add))

	as := []string{}
	ab := []bool{}
	abp := []*bool{}
	am := []MyType{}
	aass := []assert.Assertions{}
	aassp := []*assert.Assertions{}
	assert.Equal(t, "[]string", GetName(reflect.TypeOf(as), "a.b/c", imp.Add))
	assert.Equal(t, "*[]bool", GetName(reflect.TypeOf(&ab), "a.b/c", imp.Add))
	assert.Equal(t, "[]*bool", GetName(reflect.TypeOf(abp), "a.b/c", imp.Add))
	assert.Equal(t, "*[]*bool", GetName(reflect.TypeOf(&abp), "a.b/c", imp.Add))
	assert.Equal(t, "[]h.MyType", GetName(reflect.TypeOf(am), "a.b/c", impk.Add))
	assert.Equal(t, "[]as.Assertions", GetName(reflect.TypeOf(aass), "a.b/c", impa.Add))
	assert.Equal(t, "*[]as.Assertions", GetName(reflect.TypeOf(&aass), "a.b/c", impa.Add))
	assert.Equal(t, "*[]*as.Assertions", GetName(reflect.TypeOf(&aassp), "a.b/c", impa.Add))

	ms := map[string]string{}
	mb := map[string]bool{}
	mbp := map[string]*bool{}
	mm := map[string]MyType{}
	mass := map[string]assert.Assertions{}
	massp := map[string]*assert.Assertions{}
	assert.Equal(t, "map[string]string", GetName(reflect.TypeOf(ms), "a.b/c", imp.Add))
	assert.Equal(t, "*map[string]bool", GetName(reflect.TypeOf(&mb), "a.b/c", imp.Add))
	assert.Equal(t, "map[string]*bool", GetName(reflect.TypeOf(mbp), "a.b/c", imp.Add))
	assert.Equal(t, "*map[string]*bool", GetName(reflect.TypeOf(&mbp), "a.b/c", imp.Add))
	assert.Equal(t, "map[string]h.MyType", GetName(reflect.TypeOf(mm), "a.b/c", impk.Add))
	assert.Equal(t, "map[string]as.Assertions", GetName(reflect.TypeOf(mass), "a.b/c", impa.Add))
	assert.Equal(t, "*map[string]as.Assertions", GetName(reflect.TypeOf(&mass), "a.b/c", impa.Add))
	assert.Equal(t, "*map[string]*as.Assertions", GetName(reflect.TypeOf(&massp), "a.b/c", impa.Add))

	assert.Equal(t, "map[bool]string", GetName(reflect.TypeOf(map[bool]string{}), "a.b/c", imp.Add))
	assert.Equal(t, "map[*int]string", GetName(reflect.TypeOf(map[*int]string{}), "a.b/c", imp.Add))

	assert.Equal(t, "map[MyType]string", GetName(reflect.TypeOf(map[MyType]string{}), "kego.io/process/generate/literal", imp.Add))
	assert.Equal(t, "map[h.MyType]string", GetName(reflect.TypeOf(map[MyType]string{}), "a.b/c", impk.Add))
	assert.Equal(t, "map[*MyType]string", GetName(reflect.TypeOf(map[*MyType]string{}), "kego.io/process/generate/literal", imp.Add))
	assert.Equal(t, "map[*h.MyType]string", GetName(reflect.TypeOf(map[*MyType]string{}), "a.b/c", impk.Add))

	assert.Equal(t, "map[Assertions]string", GetName(reflect.TypeOf(map[assert.Assertions]string{}), asp, imp.Add))
	assert.Equal(t, "map[as.Assertions]string", GetName(reflect.TypeOf(map[assert.Assertions]string{}), "a.b/c", impa.Add))
	assert.Equal(t, "map[*Assertions]string", GetName(reflect.TypeOf(map[*assert.Assertions]string{}), asp, imp.Add))
	assert.Equal(t, "map[*as.Assertions]string", GetName(reflect.TypeOf(map[*assert.Assertions]string{}), "a.b/c", impa.Add))
}
