package process

import (
	"testing"

	"kego.io/assert"
	"kego.io/system"
)

func Test_typesImportStatement(t *testing.T) {

	i := typesImportStatement("a.b/c", "d.e/f")
	assert.Equal(t, "_ \"a.b/c/types\"", i)

	i = typesImportStatement("a.b/c", "a.b/c/types")
	assert.Equal(t, "", i)
}

func Test_importStatement(t *testing.T) {

	i := importStatement("a", "b.c/d", "e.f/g")
	assert.Equal(t, "a \"b.c/d\"", i)

	i = importStatement("a", "b.c/a", "e.f/g")
	assert.Equal(t, "\"b.c/a\"", i)

	i = importStatement("a", "b.c/a", "b.c/a")
	assert.Equal(t, "", i)
}

func TestGenerate_errors(t *testing.T) {

	_, err := Generate(F_MAIN, "a.b/c", map[string]string{})
	// No types
	assert.IsError(t, err, "HQLAEMCHBM")

	_, err = Generate(F_CMD_TYPES, "\"", map[string]string{})
	// Quote in the path will generate malformed source
	assert.IsError(t, err, "NXIWSECRLL")
	assert.HasError(t, err, "XTKWMEDWKI")

	ty := &system.Type{
		Object: &system.Object{Id: "a corrupt", Type: system.NewReference("kego.io/system", "type")},
	}
	system.RegisterType("b.c/d:a", ty)
	defer system.UnregisterType("b.c/d:a")

	_, err = Generate(F_MAIN, "b.c/d", map[string]string{})
	// Corrupt type ID causes error from source formatter
	assert.IsError(t, err, "XTIEALKSXN")

	_, err = Generate(F_TYPES, "b.c/d", map[string]string{"\"": "\""})
	// Quote in the import path will generate malformed source
	assert.IsError(t, err, "UURNHCUYAI")
	assert.HasError(t, err, "XTKWMEDWKI")
}
func TestGenerate(t *testing.T) {

	ty := &system.Type{
		Object: &system.Object{Id: "a", Type: system.NewReference("kego.io/system", "type")},
	}
	system.RegisterType("b.c/d:a", ty)
	defer system.UnregisterType("b.c/d:a")

	source, err := Generate(F_MAIN, "b.c/d", map[string]string{"e": "f.g/h"})
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package d\n")
	assert.Contains(t, string(source), "\t\"reflect\"\n")
	assert.Contains(t, string(source), "\t\"kego.io/json\"\n")
	assert.Contains(t, string(source), "\t\"kego.io/system\"\n")
	assert.Contains(t, string(source), "e \"f.g/h\"\n")
	assert.Contains(t, string(source), "\ntype A struct {\n}\n")
	assert.Contains(t, string(source), "json.RegisterType(\"b.c/d:a\", reflect.TypeOf(&A{}))\n")

	source, err = Generate(F_TYPES, "b.c/d", map[string]string{"e": "f.g/h"})
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package types")
	assert.Contains(t, string(source), "\t\"kego.io/system\"\n")
	assert.Contains(t, string(source), "_ \"kego.io/system/types\"\n")
	assert.Contains(t, string(source), "\t\"b.c/d\"\n")
	assert.Contains(t, string(source), "e \"f.g/h\"\n")
	assert.Contains(t, string(source), "system.RegisterType(\"b.c/d:a\"")

	source, err = Generate(F_CMD_TYPES, "b.c/d", map[string]string{"e": "f.g/h"})
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package main\n")
	assert.Contains(t, string(source), "\t\"kego.io/process\"\n")
	assert.Contains(t, string(source), "_ \"kego.io/system\"\n")
	assert.Contains(t, string(source), "_ \"kego.io/system/types\"\n")
	assert.Contains(t, string(source), "_ \"b.c/d\"\n")
	assert.Contains(t, string(source), "_ \"f.g/h\"\n")
	assert.Contains(t, string(source), "_ \"f.g/h/types\"\n")
	assert.Contains(t, string(source), "process.GenerateFiles(process.F_TYPES)")

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
	assert.IsError(t, err, "AHGBMCNALB")

	_, err = mapHelper(1, "a")
	assert.IsError(t, err, "WLHGIPIEUI")
}
