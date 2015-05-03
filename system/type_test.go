package system

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"kego.io/json"
)

func TestBool(t *testing.T) {
	data := `{
		"description": "This is the native json bool data type",
		"type": "type",
		"id": "bool",
		"native": "bool",
		"rule": {
			"description": "Restriction rules for bools",
			"type": "type",
			"id": "@bool",
			"is": ["rule"],
			"properties": {
				"default": {
					"description": "Default value of this is missing or null",
					"type": "property",
					"optional": true,
					"item": {
						"type": "@bool"
					}
				}
			}
		}
	}`

	var i interface{}
	err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "This is the native json bool data type", f.Description)
	assert.True(t, f.Native.Exists)
	assert.Equal(t, "bool", f.Native.Value)
	assert.Equal(t, "@bool", f.Rule.Id)

}

func TestType(t *testing.T) {
	data := `{
		"description": "This is the most basic type.",
		"type": "type",
		"id": "type",
		"properties": {
			"extends": {
				"description": "Type which this should extend",
				"type": "property",
				"optional": true,
				"item": {
					"type": "@reference",
					"default": "kego.io/system:object"
				}
			},
			"is": {
				"description": "Array of interface types that this type should support",
				"type": "property",
				"optional": true,
				"item": {
					"type": "@array",
					"items": {
						"type": "@reference"
					}
				}
			},
			"native": {
				"description": "This is the native json type that represents this type. If omitted, default is object.",
				"type": "property",
				"optional": true,
				"item": {
					"type": "@string",
					"enum": ["string", "number", "bool", "array", "object", "map"],
					"default": "object"
				}
			},
			"interface": {
				"description": "Is this type an interface?",
				"type": "property",
				"optional": true,
				"item": {
					"type": "@bool",
					"default": false
				}
			},
			"properties": {
				"description": "Each field is listed with it's type",
				"type": "property",
				"optional": true,
				"item": {
					"type": "@map",
					"items": {
						"type": "@property"
					}
				}
			},
			"rule": {
				"description": "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.",
				"optional": true,
				"type": "property",
				"item": {
					"type": "@type"
				}
			}
		},
		"rule": {
			"description": "Restriction rules for types",
			"type": "type",
			"id": "@type",
			"is": ["rule"]
		}
	}`

	var i interface{}
	err := json.Unmarshal([]byte(data), &i, "kego.io/system", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)
	assert.Equal(t, "This is the most basic type.", f.Description)
	assert.True(t, f.Native.Exists)
	assert.Equal(t, "object", f.Native.Value)
	assert.Equal(t, "@type", f.Rule.Id)
	assert.Equal(t, "Is this type an interface?", f.Properties["interface"].Description)
	assert.Equal(t, true, f.Properties["interface"].Optional.Value)
	r, ok := f.Properties["interface"].Item.(*Bool_rule)
	assert.True(t, ok, "Wrong type %T\n", f.Properties["interface"].Item)
	assert.True(t, r.Default.Exists)
	assert.Equal(t, false, r.Default.Value)

}

func unmarshalDiagram(t *testing.T) {
	diagram := `{
		"description": "This is a type of image, which just contains the url of the image",
		"type": "system:type",
		"id": "diagram",
		"properties": {
			"url": {
				"type": "system:property",
				"item": {
					"type": "system:@string"
				}
			}
		},
		"rule": {
			"description": "Restriction rules for diagram",
			"type": "system:type",
			"id": "@diagram",
			"is": ["system:rule"],
			"properties": {
				"default": {
					"description": "Default value",
					"type": "system:property",
					"optional": true,
					"defaulter": true,
					"item": {
						"type": "@diagram"
					}
				}
			}
		}
	}`

	var i interface{}
	err := json.Unmarshal([]byte(diagram), &i, "kego.io/gallery", map[string]string{})
	assert.NoError(t, err)
	d, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, d)

	fullname := fmt.Sprintf("%s:%s", "kego.io/gallery", d.Id)
	RegisterType(fullname, d)
	if d.Rule != nil {
		rulename := fmt.Sprintf("%s:%s", "kego.io/gallery", d.Rule.Id)
		RegisterType(rulename, d.Rule)
	}

}

func TestUnknownRule(t *testing.T) {

	unmarshalDiagram(t)

	data := `{
		"description": "This represents a gallery - it's just a list of images",
		"type": "system:type",
		"id": "gallery",
		"properties": {
			"image": {
				"type": "system:property",
				"optional": true,
				"item": {
					"type": "@diagram",
					"default": {
						"type": "diagram",
						"url": "def"
					}
				}
			},
			"foo": {
				"type": "system:property",
				"optional": true,
				"item": {
					"type": "system:@bool",
					"default": true
				}
			},
			"ref": {
				"type": "system:property",
				"optional": true,
				"item": {
					"type": "system:@reference",
					"default": "image"
				}
			}
		}
	}`

	var i interface{}
	err := json.Unmarshal([]byte(data), &i, "kego.io/gallery", map[string]string{})
	assert.NoError(t, err)
	f, ok := i.(*Type)
	assert.True(t, ok, "Type %T not correct", i)
	assert.NotNil(t, f)

	s, err := f.Properties["image"].GoTypeDescriptor("kego.io/gallery", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "*Diagram `kego:\"{\\\"default\\\":{\\\"type\\\":\\\"kego.io/gallery:diagram\\\",\\\"value\\\":{\\\"type\\\":\\\"diagram\\\",\\\"url\\\":\\\"def\\\"},\\\"path\\\":\\\"kego.io/gallery\\\"}}\"`", s)

	b, err := f.Properties["foo"].GoTypeDescriptor("kego.io/gallery", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "system.Bool `kego:\"{\\\"default\\\":{\\\"type\\\":\\\"kego.io/system:bool\\\",\\\"value\\\":true,\\\"path\\\":\\\"kego.io/gallery\\\"}}\"`", b)
	//

	r, err := f.Properties["ref"].GoTypeDescriptor("kego.io/gallery", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "system.Reference `kego:\"{\\\"default\\\":{\\\"type\\\":\\\"kego.io/system:reference\\\",\\\"value\\\":\\\"kego.io/gallery:image\\\",\\\"path\\\":\\\"kego.io/gallery\\\"}}\"`", r)
}
