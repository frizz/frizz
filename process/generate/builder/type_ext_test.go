package builder_test

import (
	"testing"

	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"

	"frizz.io/tests"
	"frizz.io/tests/ext"
)

func TestStructs(t *testing.T) {
	cb := tests.New().TempGopath(true)
	defer cb.Cleanup()

	testTypeCollection(t, cb)
	testTypeAlias(t, cb)
}

func testTypeCollection(t *testing.T, cb *tests.ContextBuilder) {
	source, err := ext.InitialiseAndGenerate(t, cb, "collections", map[string]string{
		"a.json": `
		{
			"type": "system:type",
			"id": "a",
			"fields": {
				"a": {
					"type": "system:@array",
					"items": {
						"type": "json:@string"
					}
				},
				"b": {
					"type": "system:@array",
					"items": {
						"type": "system:@string"
					}
				},
				"c": {
					"type": "system:@array",
					"items": {
						"type": "system:@package"
					}
				},
				"d": {
					"type": "system:@array",
					"items": {
						"type": "system:@string",
						"interface": true
					}
				}
			}
		}`,
		"b.json": `
		{
			"type": "system:type",
			"id": "b",
			"native": "string",
			"alias": {
				"type": "json:@string"
			}
		}`,
		"c.json": `
		{
			"type": "system:type",
			"id": "c",
			"native": "string",
			"alias": {
				"type": "system:@string"
			}
		}`,
		"d.json": `
		{
			"type": "system:type",
			"id": "d",
			"native": "string",
			"alias": {
				"type": "json:@string"
			}
		}`,
	})
	require.NoError(t, err)
	assert.Contains(t, source, "A []string")
	assert.Contains(t, source, "B []*system.String")
	assert.Contains(t, source, "C []*system.Package")
	assert.Contains(t, source, "D []system.StringInterface")
}

func testTypeAlias(t *testing.T, cb *tests.ContextBuilder) {
	source, err := ext.InitialiseAndGenerate(t, cb, "alias", map[string]string{
		"a.json": `
		{
			"type": "system:type",
			"id": "a",
			"native": "string",
			"alias": {
				"type": "json:@string"
			}
		}`,
		"b.json": `
		{
			"type": "system:type",
			"id": "b",
			"native": "string",
			"alias": {
				"type": "system:@string"
			}
		}`,
		"c.json": `
		{
			"type": "system:type",
			"id": "c",
			"native": "array",
			"alias": {
				"type": "system:@array",
				"items": {
					"type": "json:@string"
				}
			}
		}`,
		"d.json": `
		{
			"type": "system:type",
			"id": "d",
			"native": "array",
			"alias": {
				"type": "system:@array",
				"items": {
					"type": "system:@string"
				}
			}
		}`,
	})
	require.NoError(t, err)
	assert.Contains(t, source, "type A string")
	assert.Contains(t, source, "type B system.String")
	assert.Contains(t, source, "type C []string")
	assert.Contains(t, source, "type D []*system.String")
}
