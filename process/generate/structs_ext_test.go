package generate_test

import (
	"testing"

	"regexp"

	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
	"kego.io/process/generate"
	"kego.io/tests"
)

/*
func TestStructs(t *testing.T) {
	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	testAlias(t, cb)
	testGenerateSource(t, cb)

}

func testGenerateSource(t *testing.T, cb *tests.ContextBuilder) {
	source, err := ext.InitialiseAndGenerate(t, cb, "a", map[string]string{
		"amap.json": `
		{
			"type": "system:type",
			"id": "amap",
			"native": "map",
			"custom-kind": "map"
		}`,
		"aarr.json": `
		{
			"type": "system:type",
			"id": "aarr",
			"native": "array",
			"custom-kind": "array"
		}`,
		"a.json": `
		{
			"description": "d",
			"type": "system:type",
			"id": "a"
		}`,
		"b.json": `
		{
			"description": "bd",
			"type": "system:type",
			"id": "b",
			"fields": {
				"a": {
					"description": "fd",
					"type": "system:@string"
				}
			}
		}`,
		"an.json": `
		{
			"type": "system:type",
			"id": "an",
			"native": "bool",
			"alias": {
				"type": "json:@bool"
			}
		}`,
		"ai.json": `
		{
			"type": "system:type",
			"id": "ai",
			"interface": true
		}`,
	})
	require.NoError(t, err)
	assert.Contains(t, source, "package a\n")
	imp := getImports(t, source)
	assert.Contains(t, imp, "\t\"context\"\n")
	assert.Contains(t, imp, "\t\"kego.io/context/jsonctx\"\n")
	assert.Contains(t, imp, "\t\"kego.io/system\"\n")
	assert.Contains(t, imp, "\t\"reflect\"\n")
	assert.NotContains(t, imp, "\"f.g/h\"")
	assert.NotContains(t, imp, "Amap") // Native collections shouldn't go into the generated file
	assert.NotContains(t, imp, "Aarr") // Native collections shouldn't go into the generated file
	assert.Contains(t, source, "\n// d\ntype A struct {\n\t*system.Object\n}\n")
	assert.Contains(t, source, "pkg.InitType(\"a\", reflect.TypeOf((*A)(nil)), reflect.TypeOf((*ARule)(nil)), reflect.TypeOf((*AInterface)(nil)).Elem())\n")

	// Type "B" is a struct with a description on the field
	assert.Contains(t, source, "// fd\n")

	// Type "Ai" is an interface
	assert.Contains(t, source, "pkg.InitType(\"ai\", reflect.TypeOf((*Ai)(nil)).Elem(), reflect.TypeOf((*AiRule)(nil)), nil)\n")

	// Type "An" should not be a struct because it's a bool native. However, it should be registered.
	assert.Contains(t, source, "type An bool")
	assert.Contains(t, source, "pkg.InitType(\"an\", reflect.TypeOf((*An)(nil)), reflect.TypeOf((*AnRule)(nil)), reflect.TypeOf((*AnInterface)(nil)).Elem())\n")
}

func testAlias(t *testing.T, cb *tests.ContextBuilder) {
	source, err := ext.InitialiseAndGenerate(t, cb, "c", map[string]string{
		"type-alias-map-json-string.json": `
			{
				"description": "Alias of json string map",
				"type": "system:type",
				"id": "type-alias-map-json-string",
				"alias": {
					"type": "system:@map",
					"items": {
						"type": "json:@string"
					}
				}
			}
		`,
		"type-alias-array-json-string.json": `
			{
				"description": "Alias of json string array",
				"type": "system:type",
				"id": "type-alias-array-json-string",
				"alias": {
					"type": "system:@array",
					"items": {
						"type": "json:@string"
					}
				}
			}
		`,
		"type-alias-map-system-string.json": `
			{
				"description": "Alias of system string map",
				"type": "system:type",
				"id": "type-alias-map-system-string",
				"alias": {
					"type": "system:@map",
					"items": {
						"type": "system:@string"
					}
				}
			}
		`,
		"type-alias-array-system-string.json": `
			{
				"description": "Alias of system string array",
				"type": "system:type",
				"id": "type-alias-array-system-string",
				"alias": {
					"type": "system:@array",
					"items": {
						"type": "system:@string"
					}
				}
			}
		`,
		"type-alias-system-string.json": `
			{
				"description": "Alias of system string",
				"type": "system:type",
				"id": "type-alias-system-string",
				"alias": {
					"type": "system:@string"
				}
			}
		`,
		"type-alias-json-string.json": `
			{
				"description": "Alias of json string",
				"type": "system:type",
				"id": "type-alias-json-string",
				"alias": {
					"type": "json:@string"
				}
			}
		`,
		"type-alias-custom-value.json": `
			{
				"description": "Alias of custom value type",
				"type": "system:type",
				"id": "type-alias-custom-value",
				"custom": true,
				"custom-kind": "value"
			}
		`,
		"type-alias-custom-collection.json": `
			{
				"description": "Alias of custom collection type",
				"type": "system:type",
				"id": "type-alias-custom-collection",
				"custom": true,
				"custom-kind": "map"
			}
		`,
		"type-alias-custom-struct.json": `
			{
				"description": "Alias of custom struct type",
				"type": "system:type",
				"id": "type-alias-custom-struct",
				"custom": true,
				"custom-kind": "struct"
			}
		`,
		"struct-containing-alias-type-fields.json": `
			{
				"type": "system:type",
				"id": "struct-containing-alias-type-fields",
				"fields": {
					"field-alias-json-string": {
						"type": "@type-alias-json-string"
					},
					"field-alias-system-string": {
						"type": "@type-alias-system-string"
					},
					"field-alias-map-json-string": {
						"type": "@type-alias-map-json-string"
					},
					"field-alias-array-json-string": {
						"type": "@type-alias-array-json-string"
					},
					"field-alias-map-system-string": {
						"type": "@type-alias-map-system-string"
					},
					"field-alias-array-system-string": {
						"type": "@type-alias-array-system-string"
					},
					"field-alias-custom-value": {
						"type": "@type-alias-custom-value"
					},
					"field-alias-custom-collection": {
						"type": "@type-alias-custom-collection"
					},
					"field-alias-custom-struct": {
						"type": "@type-alias-custom-struct"
					}
				}
			}
		`,
	})
	require.NoError(t, err)

	assert.Contains(t, source, "type TypeAliasMapJsonString map[string]string\n")
	assert.Regexp(t, `FieldAliasMapJsonString\s+TypeAliasMapJsonString`, source)

	assert.Contains(t, source, "type TypeAliasArrayJsonString []string\n")
	assert.Regexp(t, `FieldAliasArrayJsonString\s+TypeAliasArrayJsonString`, source)

	assert.Contains(t, source, "type TypeAliasMapSystemString map[string]*system.String\n")
	assert.Regexp(t, `FieldAliasMapSystemString\s+TypeAliasMapSystemString`, source)

	assert.Contains(t, source, "type TypeAliasArraySystemString []*system.String\n")
	assert.Regexp(t, `FieldAliasArraySystemString\s+TypeAliasArraySystemString`, source)

	assert.Contains(t, source, "type TypeAliasSystemString system.String\n")
	assert.Regexp(t, `FieldAliasSystemString\s+\*TypeAliasSystemString`, source)

	assert.Contains(t, source, "type TypeAliasJsonString string\n")
	assert.Regexp(t, `FieldAliasJsonString\s+\*TypeAliasJsonString`, source)

	assert.NotRegexp(t, `type TypeAliasCustomValue\s`, source)
	assert.Regexp(t, `FieldAliasCustomValue\s+\*TypeAliasCustomValue`, source)

	assert.NotRegexp(t, `type TypeAliasCustomCollection\s`, source)
	assert.Regexp(t, `FieldAliasCustomCollection\s+TypeAliasCustomCollection`, source)

	assert.NotRegexp(t, `type TypeAliasCustomStruct\s`, source)
	assert.Regexp(t, `FieldAliasCustomStruct\s+\*TypeAliasCustomStruct`, source)

}
*/

func TestGenerateSourceErr1(t *testing.T) {
	cb := tests.Context("a.b/c").Sempty()
	_, err := generate.Structs(cb.Ctx(), cb.Env())
	assert.IsError(t, err, "DQVQWTKRSK")
}

func TestGenerateSourceNoTypes(t *testing.T) {
	cb := tests.Context("a.b/c").Spkg("a.b/c")
	_, err := generate.Structs(cb.Ctx(), cb.Env())
	require.NoError(t, err)
}

func getImports(t *testing.T, source string) string {
	r := regexp.MustCompile(`(?m)import \(([^)]*)\)`)
	matches := r.FindStringSubmatch(source)
	assert.Equal(t, 2, len(matches))
	return matches[1]
}
