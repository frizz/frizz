package system

import (
	"testing"

	"frizz.io/frizz"
	"frizz.io/validators"

	"reflect"

	"frizz.io/common"
)

func TestType(t *testing.T) {
	data, err := frizz.Package("frizz.io/system", Packer, validators.Packer)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) != 1 {
		t.Fatal("Expected 1 item to be decoded")
	}
	typ, ok := data["type"]
	if !ok {
		t.Fatal("Expected item to be named 'type'")
	}
	expected := Type{
		Fields: map[string]Field{
			"Fields": {
				Validators: []common.Validator{
					validators.Keys{
						Validators: []common.Validator{
							validators.Regex{
								Regex: "^_?[a-z][-a-z]*$",
							},
						},
					},
					validators.Items{
						Validators: []common.Validator{},
					},
				},
			},
		},
	}
	if !reflect.DeepEqual(typ, expected) {
		t.Fatalf("unmarshaled type not what is expected. got: %#v, expected: %#v", typ, expected)
	}
}
