package generate_test

import (
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/process"
	"kego.io/process/generate"
	"kego.io/tests"
)

func TestStructs(t *testing.T) {
	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	testAlias(t, cb)
	testNative(t, cb)

}

func testNative(t *testing.T, cb *tests.ContextBuilder) {
	source := initialise(t, cb, "a", map[string]string{
		"type-native-string.json": `
			{
				"description": "Native string",
				"type": "system:type",
				"id": "type-native-string",
				"native": "string"
			}
		`,
		"type-native-number.json": `
			{
				"description": "Native number",
				"type": "system:type",
				"id": "type-native-number",
				"native": "number"
			}
		`,
		"type-native-bool.json": `
			{
				"description": "Native bool",
				"type": "system:type",
				"id": "type-native-bool",
				"native": "bool"
			}
		`,
		"struct-containing-native-type-fields.json": `
			{
				"type": "system:type",
				"id": "struct-containing-native-type-fields",
				"fields": {
					"field-native-string": {
						"type": "@type-native-string"
					},
					"field-native-number": {
						"type": "@type-native-number"
					},
					"field-native-bool": {
						"type": "@type-native-bool"
					}
				}
			}
		`,
	})
	assert.Contains(t, source, "type TypeNativeString string\n")
	assert.Contains(t, source, "type TypeNativeNumber float64\n")
	assert.Contains(t, source, "type TypeNativeBool bool\n")
	assert.Regexp(t, `FieldNativeString\s+\*TypeNativeString`, source)
	assert.Regexp(t, `FieldNativeNumber\s+\*TypeNativeNumber`, source)
	assert.Regexp(t, `FieldNativeBool\s+\*TypeNativeBool`, source)
}

func testAlias(t *testing.T, cb *tests.ContextBuilder) {
	source := initialise(t, cb, "a", map[string]string{
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
					}
				}
			}
		`,
	})
	assert.Contains(t, source, "type TypeAliasMapJsonString map[string]string\n")
	assert.Contains(t, source, "type TypeAliasArrayJsonString []string\n")
	assert.Contains(t, source, "type TypeAliasSystemString *system.String\n")
	assert.Contains(t, source, "type TypeAliasJsonString string\n")
	assert.Regexp(t, `FieldAliasJsonString\s+\*TypeAliasJsonString`, source)
	assert.Regexp(t, `FieldAliasSystemString\s+\*TypeAliasSystemString`, source)
	assert.Regexp(t, `FieldAliasMapJsonString\s+\*TypeAliasMapJsonString`, source)
	assert.Regexp(t, `FieldAliasArrayJsonString\s+\*TypeAliasArrayJsonString`, source)
}

func initialise(t *testing.T, cb *tests.ContextBuilder, name string, files map[string]string) string {
	path, _ := cb.TempPackage(name, files)
	ctx, _, err := process.Initialise(cb.Ctx(), &process.Options{
		Path: path,
	})
	require.NoError(t, err)
	cb.SetCtx(ctx)
	source, err := generate.Structs(ctx, cb.Env())
	require.NoError(t, err)
	return string(source)
}
