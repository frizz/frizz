package literal

import (
	"reflect"
	"testing"

	"kego.io/assert"
)

func TestGetName(t *testing.T) {

	imp := map[string]string{"d": "e.f/g"}

	s := ""
	b := true
	i := 1
	f := 1.2

	assert.Equal(t, "string", GetName(reflect.TypeOf(s), "a.b/c", imp))
	assert.Equal(t, "bool", GetName(reflect.TypeOf(b), "a.b/c", imp))
	assert.Equal(t, "int", GetName(reflect.TypeOf(i), "a.b/c", imp))
	assert.Equal(t, "float64", GetName(reflect.TypeOf(f), "a.b/c", imp))

	assert.Equal(t, "*string", GetName(reflect.TypeOf(&s), "a.b/c", imp))
	assert.Equal(t, "*bool", GetName(reflect.TypeOf(&b), "a.b/c", imp))
	assert.Equal(t, "*int", GetName(reflect.TypeOf(&i), "a.b/c", imp))
	assert.Equal(t, "*float64", GetName(reflect.TypeOf(&f), "a.b/c", imp))

	impk := map[string]string{"d": "e.f/g", "h": "kego.io/literal"}
	type MyType struct{}
	m := MyType{}

	assert.Equal(t, "MyType", GetName(reflect.TypeOf(m), "kego.io/literal", imp))
	assert.Equal(t, "h.MyType", GetName(reflect.TypeOf(m), "a.b/c", impk))
	assert.Equal(t, "*MyType", GetName(reflect.TypeOf(&m), "kego.io/literal", imp))
	assert.Equal(t, "*h.MyType", GetName(reflect.TypeOf(&m), "a.b/c", impk))

	asp := "kego.io/assert"
	impa := map[string]string{"d": "e.f/g", "as": asp}
	ass := assert.Assertions{} // Just using a random struct from a non system package

	assert.Equal(t, "Assertions", GetName(reflect.TypeOf(ass), asp, imp))
	assert.Equal(t, "as.Assertions", GetName(reflect.TypeOf(ass), "a.b/c", impa))
	assert.Equal(t, "*Assertions", GetName(reflect.TypeOf(&ass), asp, imp))
	assert.Equal(t, "*as.Assertions", GetName(reflect.TypeOf(&ass), "a.b/c", impa))

	as := []string{}
	ab := []bool{}
	abp := []*bool{}
	am := []MyType{}
	aass := []assert.Assertions{}
	aassp := []*assert.Assertions{}
	assert.Equal(t, "[]string", GetName(reflect.TypeOf(as), "a.b/c", imp))
	assert.Equal(t, "*[]bool", GetName(reflect.TypeOf(&ab), "a.b/c", imp))
	assert.Equal(t, "[]*bool", GetName(reflect.TypeOf(abp), "a.b/c", imp))
	assert.Equal(t, "*[]*bool", GetName(reflect.TypeOf(&abp), "a.b/c", imp))
	assert.Equal(t, "[]h.MyType", GetName(reflect.TypeOf(am), "a.b/c", impk))
	assert.Equal(t, "[]as.Assertions", GetName(reflect.TypeOf(aass), "a.b/c", impa))
	assert.Equal(t, "*[]as.Assertions", GetName(reflect.TypeOf(&aass), "a.b/c", impa))
	assert.Equal(t, "*[]*as.Assertions", GetName(reflect.TypeOf(&aassp), "a.b/c", impa))

	ms := map[string]string{}
	mb := map[string]bool{}
	mbp := map[string]*bool{}
	mm := map[string]MyType{}
	mass := map[string]assert.Assertions{}
	massp := map[string]*assert.Assertions{}
	assert.Equal(t, "map[string]string", GetName(reflect.TypeOf(ms), "a.b/c", imp))
	assert.Equal(t, "*map[string]bool", GetName(reflect.TypeOf(&mb), "a.b/c", imp))
	assert.Equal(t, "map[string]*bool", GetName(reflect.TypeOf(mbp), "a.b/c", imp))
	assert.Equal(t, "*map[string]*bool", GetName(reflect.TypeOf(&mbp), "a.b/c", imp))
	assert.Equal(t, "map[string]h.MyType", GetName(reflect.TypeOf(mm), "a.b/c", impk))
	assert.Equal(t, "map[string]as.Assertions", GetName(reflect.TypeOf(mass), "a.b/c", impa))
	assert.Equal(t, "*map[string]as.Assertions", GetName(reflect.TypeOf(&mass), "a.b/c", impa))
	assert.Equal(t, "*map[string]*as.Assertions", GetName(reflect.TypeOf(&massp), "a.b/c", impa))

	assert.Equal(t, "map[bool]string", GetName(reflect.TypeOf(map[bool]string{}), "a.b/c", imp))
	assert.Equal(t, "map[*int]string", GetName(reflect.TypeOf(map[*int]string{}), "a.b/c", imp))

	assert.Equal(t, "map[MyType]string", GetName(reflect.TypeOf(map[MyType]string{}), "kego.io/literal", imp))
	assert.Equal(t, "map[h.MyType]string", GetName(reflect.TypeOf(map[MyType]string{}), "a.b/c", impk))
	assert.Equal(t, "map[*MyType]string", GetName(reflect.TypeOf(map[*MyType]string{}), "kego.io/literal", imp))
	assert.Equal(t, "map[*h.MyType]string", GetName(reflect.TypeOf(map[*MyType]string{}), "a.b/c", impk))

	assert.Equal(t, "map[Assertions]string", GetName(reflect.TypeOf(map[assert.Assertions]string{}), asp, imp))
	assert.Equal(t, "map[as.Assertions]string", GetName(reflect.TypeOf(map[assert.Assertions]string{}), "a.b/c", impa))
	assert.Equal(t, "map[*Assertions]string", GetName(reflect.TypeOf(map[*assert.Assertions]string{}), asp, imp))
	assert.Equal(t, "map[*as.Assertions]string", GetName(reflect.TypeOf(map[*assert.Assertions]string{}), "a.b/c", impa))
}
