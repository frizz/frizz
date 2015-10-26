package system_test

import (
	"testing"

	"kego.io/json"
	"kego.io/kerr/assert"
	"kego.io/system"
	"kego.io/system/node"
)

func TestBool(t *testing.T) {
	testBool(t, unmarshalFunc)
	testBool(t, unpackFunc)
	testBool(t, repackFunc)
}
func testBool(t *testing.T, unpacker unpackerFunc) {

	data := `{
		"description": "This is the native json bool data type",
		"type": "type",
		"id": "bool",
		"native": "bool",
		"rule": {
			"description": "Restriction rules for bools",
			"type": "type",
			"embed": ["rule"],
			"fields": {
				"default": {
					"description": "Default value of this is missing or null",
					"type": "@bool",
					"optional": true
				}
			}
		}
	}`

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*system.Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "This is the native json bool data type", f.Description)
	assert.True(t, f.Native.Exists)
	assert.Equal(t, "bool", f.Native.Value)

}

func TestType(t *testing.T) {
	testType(t, unmarshalFunc)
	testType(t, unpackFunc)
	testType(t, repackFunc)
}
func testType(t *testing.T, unpacker unpackerFunc) {

	data := `{
		"description": "This is the most basic type.",
		"type": "type",
		"id": "type",
		"fields": {
			"is": {
				"description": "Array of interface types that this type should support",
				"type": "@array",
				"items": {
					"type": "@reference"
				},
				"optional": true
			},
			"native": {
				"description": "This is the native json type that represents this type. If omitted, default is object.",
				"type": "@string",
				"enum": ["string", "number", "bool", "array", "object", "map"],
				"default": "object",
				"optional": true
			},
			"interface": {
				"description": "Is this type an interface?",
				"type": "@bool",
				"default": false,
				"optional": true
			},
			"fields": {
				"description": "Each field is listed with it's type",
				"type": "@map",
				"items": {
					"type": "@rule"
				},
				"optional": true
			},
			"rule": {
				"description": "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.",
				"type": "@type",
				"optional": true
			}
		},
		"rule": {
			"description": "Restriction rules for types",
			"type": "type",
			"is": ["rule"]
		}
	}`

	var i interface{}
	err := unpacker([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*system.Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "This is the most basic type.", f.Description)
	assert.True(t, f.Native.Exists)
	assert.Equal(t, "object", f.Native.Value)
	assert.Equal(t, "Is this type an interface?", f.Fields["interface"].(system.ObjectInterface).GetObject().Description)
	assert.Equal(t, true, f.Fields["interface"].GetRule().Optional)
	r, ok := f.Fields["interface"].(*system.BoolRule)
	assert.True(t, ok, "Wrong type %T\n", f.Fields["interface"])
	assert.True(t, r.Default.Exists)
	assert.Equal(t, false, r.Default.Value)

}

type unpackerFunc func([]byte, *interface{}, string, map[string]string) error

func unmarshalFunc(data []byte, i *interface{}, path string, aliases map[string]string) error {
	return json.Unmarshal(data, i, path, aliases)
}
func unpackFunc(data []byte, i *interface{}, path string, aliases map[string]string) error {
	var j interface{}
	if err := json.UnmarshalPlain(data, &j); err != nil {
		return err
	}
	return json.Unpack(json.NewJsonUnpacker(j), i, path, aliases)
}

func repackFunc(data []byte, i *interface{}, path string, aliases map[string]string) error {
	var n node.Node
	if err := json.UnmarshalPlainContext(data, &n, path, aliases); err != nil {
		return err
	}
	return json.Unpack(&n, i, path, aliases)
}
