package system

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"kego.io/json"
)

func TestGoTypeDescriptor(t *testing.T) {

	var i interface{}
	json.Unmarshal([]byte(`{
			"type": "system:property",
			"item": {
				"type": "json:@string"
			}
		}`), &i, "kego.io/system", map[string]string{})
	p, ok := i.(*Property)
	assert.True(t, ok, "Wrong type")
	s, err := p.GoTypeDescriptor("kego.io/system", map[string]string{})
	assert.NoError(t, err)
	assert.Equal(t, "string", s)
}
