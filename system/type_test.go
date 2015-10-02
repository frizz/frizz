package system

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestBool(t *testing.T) {
	testBool(t, unmarshalFunc)
	testBool(t, unpackFunc)
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
			"is": ["rule"],
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
	f, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "This is the native json bool data type", f.Description)
	assert.True(t, f.Native.Exists)
	assert.Equal(t, "bool", f.Native.Value)

}

func TestType(t *testing.T) {
	testType(t, unmarshalFunc)
	testType(t, unpackFunc)
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
	f, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "This is the most basic type.", f.Description)
	assert.True(t, f.Native.Exists)
	assert.Equal(t, "object", f.Native.Value)
	assert.Equal(t, "Is this type an interface?", f.Fields["interface"].(Object).Object().Description)
	assert.Equal(t, true, f.Fields["interface"].Rule().Optional)
	r, ok := f.Fields["interface"].(*Bool_rule)
	assert.True(t, ok, "Wrong type %T\n", f.Fields["interface"])
	assert.True(t, r.Default.Exists)
	assert.Equal(t, false, r.Default.Value)

}

func TestUnregister(t *testing.T) {
	Unregister("", "")
}

func TestNativeGoType(t *testing.T) {
	n, err := nativeGoType("string")
	assert.NoError(t, err)
	assert.Equal(t, "string", n)

	n, err = nativeGoType("number")
	assert.NoError(t, err)
	assert.Equal(t, "float64", n)

	n, err = nativeGoType("bool")
	assert.NoError(t, err)
	assert.Equal(t, "bool", n)

	_, err = nativeGoType("a")
	assert.IsError(t, err, "TXQIDRBJRH")
}

func TestTypeIsNativeType(t *testing.T) {
	y := &Type{Native: NewString("bool")}
	n := y.IsNativeObject()
	assert.False(t, n)

	y = &Type{Native: NewString("array")}
	n = y.IsNativeObject()
	assert.False(t, n)

	y = &Type{Native: NewString("object")}
	n = y.IsNativeObject()
	assert.True(t, n)

	// Will also return true for any other value
	y = &Type{Native: NewString("a")}
	n = y.IsNativeObject()
	assert.True(t, n)
}

func TestTypeNativeValueGolangType(t *testing.T) {
	y := &Type{Native: NewString("bool")}
	n, err := y.NativeValueGolangType()
	assert.NoError(t, err)
	assert.Equal(t, "bool", n)
}

func TestTypeGoName(t *testing.T) {
	y := &Type{Object_base: &Object_base{Id: NewReference("a.b/c", "aa")}}
	n := y.GoName()
	assert.Equal(t, "Aa", n)
}
