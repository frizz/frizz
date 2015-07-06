package process

import (
	"testing"

	"regexp"

	"kego.io/kerr/assert"
	"kego.io/system"
)

func TestGenerate_errors(t *testing.T) {

	_, err := Generate(F_CMD_TYPES, "\"", map[string]string{})
	// Quote in the path will generate malformed source
	assert.IsError(t, err, "CRBYOUOHPG")

	ty := &system.Type{
		Base: &system.Base{Id: system.NewReference("b.c/d", "a corrupt"), Type: system.NewReference("kego.io/system", "type")},
	}
	system.RegisterType("b.c/d", "a", ty)
	defer system.UnregisterType("b.c/d", "a")

	_, err = Generate(F_MAIN, "b.c/d", map[string]string{})
	// Corrupt type ID causes error from source formatter
	assert.IsError(t, err, "CRBYOUOHPG")

	_, err = Generate(F_TYPES, "b.c/d", map[string]string{"\"": "\""})
	// Quote in the import path will generate malformed source
	assert.IsError(t, err, "CRBYOUOHPG")
}
func getImports(t *testing.T, source string) string {
	r := regexp.MustCompile(`(?m)import \(([^)]*)\)`)
	matches := r.FindStringSubmatch(source)
	assert.Equal(t, 2, len(matches))
	return matches[1]
}
func TestGenerate(t *testing.T) {

	ty := &system.Type{
		Base: &system.Base{Id: system.NewReference("b.c/d", "a"), Type: system.NewReference("kego.io/system", "type")},
	}
	system.RegisterType("b.c/d", "a", ty)
	defer system.UnregisterType("b.c/d", "a")

	source, err := Generate(F_MAIN, "b.c/d", map[string]string{"f.g/h": "e"})
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package d\n")
	imp := getImports(t, string(source))
	assert.Contains(t, imp, "\tjson \"kego.io/json\"\n")
	assert.Contains(t, imp, "\tsystem \"kego.io/system\"\n")
	assert.Contains(t, imp, "\treflect \"reflect\"\n")
	assert.NotContains(t, imp, "\"f.g/h\"")
	assert.Contains(t, string(source), "\ntype A struct {\n\t*system.Base\n}\n")
	assert.Contains(t, string(source), "json.RegisterType(\"b.c/d\", \"a\", reflect.TypeOf(&A{}))\n")

	source, err = Generate(F_TYPES, "b.c/d", map[string]string{"f.g/h": "e"})
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package types")
	imp = getImports(t, string(source))
	assert.Contains(t, imp, "\t_ \"f.g/h/types\"\n")
	assert.Contains(t, imp, "\tsystem \"kego.io/system\"\n")
	assert.Contains(t, imp, "\t_ \"kego.io/system/types\"\n")
	assert.NotContains(t, imp, "\"b.c/d\"")
	assert.NotContains(t, imp, "\"f.g/h\"")
	assert.Contains(t, string(source), "system.RegisterType(\"b.c/d\", \"a\"")

	source, err = Generate(F_CMD_TYPES, "b.c/d", map[string]string{"f.g/h": "e"})
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package main\n")
	imp = getImports(t, string(source))
	assert.Contains(t, imp, "_ \"b.c/d\"\n")
	assert.Contains(t, imp, "_ \"f.g/h\"\n")
	assert.Contains(t, imp, "_ \"f.g/h/types\"\n")
	assert.Contains(t, imp, "fmt \"fmt\"\n")
	assert.Contains(t, imp, "\tprocess \"kego.io/process\"\n")
	assert.Contains(t, imp, "_ \"kego.io/system\"\n")
	assert.Contains(t, imp, "_ \"kego.io/system/types\"\n")
	assert.Contains(t, imp, "os \"os\"\n")
	assert.Contains(t, string(source), "process.GenerateFiles(process.F_TYPES, dir, test, recursive, verbose, path, imports)")

}
