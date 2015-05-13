package process

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"kego.io/system"
	"kego.io/uerr"
)

func TestGenerate_errors(t *testing.T) {

	_, err := Generate("a.b/c", map[string]string{})
	// No types
	uerr.Assert(t, err, "HQLAEMCHBM")

	ty := &system.Type{
		Object: &system.Object{Id: "a corrupt", Type: system.NewReference("kego.io/system", "type")},
	}
	system.RegisterType("b.c/d:a", ty)
	defer system.UnregisterType("b.c/d:a")

	_, err = Generate("b.c/d", map[string]string{})
	// Corrupt type ID causes error from source formatter
	uerr.Assert(t, err, "XTKWMEDWKI")
}
func TestGenerate(t *testing.T) {

	ty := &system.Type{
		Object: &system.Object{Id: "a", Type: system.NewReference("kego.io/system", "type")},
	}
	system.RegisterType("b.c/d:a", ty)
	defer system.UnregisterType("b.c/d:a")

	b, err := Generate("b.c/d", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, `package d

import (
	"reflect"

	"kego.io/json"

	"kego.io/system"
)

//***********************************************************
//*** a ***
//***********************************************************

type A struct {
}

func jsonTypes() map[string]reflect.Type {
	return map[string]reflect.Type{

		"b.c/d:a": reflect.TypeOf(&A{}),
	}
}

func systemTypes() map[string]*system.Type {
	return map[string]*system.Type{

		"b.c/d:a": &system.Type{Object: &system.Object{Context: (*system.Context)(nil), Description: "", Id: "a", Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "", Exists: false}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)},
	}
}

func init() {
	json.RegisterPackage("b.c/d", jsonTypes)
	system.RegisterPackage("b.c/d", systemTypes)
}
`, string(b))

}

func TestGetPackageNameFromPath(t *testing.T) {
	n := getPackageNameFromPath("a.b/c")
	assert.Equal(t, "c", n)

	// Double check we're not splitting on backslash
	n = getPackageNameFromPath(`a.b/c\d`)
	assert.Equal(t, `c\d`, n)
}

func TestImportStatement(t *testing.T) {
	i := importStatement("a", "b.c/a", "b.c/a")
	assert.Equal(t, "", i)

	i = importStatement("a", "b.c/a", "d.e/f")
	assert.Equal(t, `"b.c/a"`, i)

	i = importStatement("a", "b.c/d", "e.f/g")
	assert.Equal(t, `a "b.c/d"`, i)
}

func TestTernary(t *testing.T) {
	a := ternary(true, "b", "c")
	assert.Equal(t, "b", a)

	a = ternary(false, "b", "c")
	assert.Equal(t, "c", a)
}

func TestMapHelper(t *testing.T) {
	m, err := mapHelper("a", "b", "c", 2)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(m))
	assert.Equal(t, "b", m["a"])
	assert.Equal(t, 2, m["c"])

	_, err = mapHelper("a", "b", "c")
	uerr.Assert(t, err, "AHGBMCNALB")

	_, err = mapHelper(1, "a")
	uerr.Assert(t, err, "WLHGIPIEUI")
}
