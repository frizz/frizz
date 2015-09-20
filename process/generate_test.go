package process

import (
	"testing"

	"regexp"

	"kego.io/kerr/assert"
	"kego.io/system"
)

func TestGenerateCommand_errors(t *testing.T) {

	_, err := GenerateCommand(C_TYPES, settings{path: "\""})
	// Quote in the path will generate malformed source
	assert.IsError(t, err, "CRBYOUOHPG")

	ty := &system.Type{
		Object_base: &system.Object_base{Id: system.NewReference("b.c/d", "a corrupt"), Type: system.NewReference("kego.io/system", "type")},
	}
	system.Register("b.c/d", "a", ty, 0)
	defer system.Unregister("b.c/d", "a")

	_, err = GenerateSource(S_STRUCTS, settings{path: "b.c/d"})
	// Corrupt type ID causes error from source formatter
	assert.IsError(t, err, "CRBYOUOHPG")

	_, err = GenerateSource(S_TYPES, settings{path: "b.c/d", aliases: map[string]string{"\"": "\""}})
	// Quote in the alias path will generate malformed source
	assert.IsError(t, err, "CRBYOUOHPG")
}
func getImports(t *testing.T, source string) string {
	r := regexp.MustCompile(`(?m)import \(([^)]*)\)`)
	matches := r.FindStringSubmatch(source)
	assert.Equal(t, 2, len(matches))
	return matches[1]
}
func TestGenerateSource(t *testing.T) {

	ty := &system.Type{
		Object_base: &system.Object_base{Id: system.NewReference("b.c/d", "a"), Type: system.NewReference("kego.io/system", "type")},
	}
	system.Register("b.c/d", "a", ty, 0)
	defer system.Unregister("b.c/d", "a")

	source, err := GenerateSource(S_STRUCTS, settings{path: "b.c/d", aliases: map[string]string{"f.g/h": "e"}})
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package d\n")
	imp := getImports(t, string(source))
	assert.Contains(t, imp, "\t\"kego.io/json\"\n")
	assert.Contains(t, imp, "\t\"kego.io/system\"\n")
	assert.Contains(t, imp, "\t\"reflect\"\n")
	assert.NotContains(t, imp, "\"f.g/h\"")
	assert.Contains(t, string(source), "\ntype A struct {\n\t*system.Object_base\n}\n")
	assert.Contains(t, string(source), "json.Register(\"b.c/d\", \"a\", reflect.TypeOf(&A{}), 0x0)\n")

	source, err = GenerateSource(S_TYPES, settings{path: "b.c/d", aliases: map[string]string{"f.g/h": "e"}})
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package types")
	imp = getImports(t, string(source))
	assert.Contains(t, imp, "\t_ \"f.g/h/types\"\n")
	assert.Contains(t, imp, "\t\"kego.io/system\"\n")
	assert.Contains(t, imp, "\t_ \"kego.io/system/types\"\n")
	assert.NotContains(t, imp, "\"b.c/d\"")
	assert.NotContains(t, imp, "\"f.g/h\"")
	assert.Contains(t, string(source), "system.Register(\"b.c/d\", \"a\", ptr1, 0x0)")

	source, err = GenerateCommand(C_TYPES, settings{path: "b.c/d", aliases: map[string]string{"f.g/h": "e"}})
	assert.NoError(t, err)
	assert.Contains(t, string(source), "package main\n")
	imp = getImports(t, string(source))
	assert.Contains(t, imp, "_ \"b.c/d\"\n")
	assert.Contains(t, imp, "_ \"f.g/h\"\n")
	assert.Contains(t, imp, "_ \"f.g/h/types\"\n")
	assert.Contains(t, imp, "\"fmt\"\n")
	assert.Contains(t, imp, "\t\"kego.io/process\"\n")
	assert.Contains(t, imp, "_ \"kego.io/system\"\n")
	assert.Contains(t, imp, "_ \"kego.io/system/types\"\n")
	assert.Contains(t, imp, "\"os\"\n")
	assert.Contains(t, string(source), "process.Generate(process.S_TYPES, set)")

}
