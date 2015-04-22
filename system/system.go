package system

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/kego/json"
)

type Reference struct {
	Value   string
	Package string
	Type    string
	Exists  bool
}

func (out *Reference) UnmarshalJSON(in []byte, context *json.Context) error {
	var s *string
	if err := json.Unmarshal(in, &s, context); err != nil {
		return err
	}
	if s == nil {
		out.Exists = false
	} else {
		out.Exists = true
		if strings.Contains(*s, "/") {
			// If the type name contains a slash, I'm assuming it's a fully qualified type name of
			// the form "github.com/kego/system:type".
			// TODO: Improve this with a regex?
			parts := strings.Split(*s, ":")
			out.Package = parts[0]
			out.Type = parts[1]
		} else if strings.Contains(*s, ":") {
			// If the type name contains a colon, I'm assuming it's an abreviated qualified type name of
			// the form "system:type". We should look the package name up in the imports map.
			// TODO: Improve this with a regex?
			parts := strings.Split(*s, ":")
			packagePath, ok := context.Imports[parts[0]]
			if !ok {
				return fmt.Errorf("Error in system.FullTypeName: package name %v not found in imports.\n", parts[0])
			}
			out.Package = packagePath
			out.Type = parts[1]
		} else {
			out.Package = context.PackagePath
			out.Type = *s
		}
		out.Value = fmt.Sprintf("%s:%s", out.Package, out.Type)
	}
	return nil
}

type Number struct {
	Value  float64
	Exists bool
}

func (out *Number) UnmarshalJSON(in []byte, context *json.Context) error {
	var f *float64
	if err := json.Unmarshal(in, &f, context); err != nil {
		return err
	}
	if f == nil {
		out.Exists = false
	} else {
		out.Exists = true
		out.Value = *f
	}
	return nil
}

type String struct {
	Value  string
	Exists bool
}

func (out *String) UnmarshalJSON(in []byte, context *json.Context) error {
	var s *string
	if err := json.Unmarshal(in, &s, context); err != nil {
		return err
	}
	if s == nil {
		out.Exists = false
	} else {
		out.Exists = true
		out.Value = *s
	}
	return nil
}

type Bool struct {
	Value  bool
	Exists bool
}

func (out *Bool) UnmarshalJSON(in []byte, context *json.Context) error {
	var b *bool
	if err := json.Unmarshal(in, &b, context); err != nil {
		return err
	}
	if b == nil {
		out.Exists = false
	} else {
		out.Exists = true
		out.Value = *b
	}
	return nil
}

type Bool_rule struct {
	Default Bool
}

// This is the most basic type.
type Object struct {
	// Description for the developer
	Description String
	// All global objects should have an id.
	Id String
	// Type of the object.
	Type Reference
}

type Property struct {
	Object
	Optional Bool `kego:"{\"default\": false}"`
	Item     Rule
}

type Rule interface{}

// This is the most basic type.
type Type struct {
	Object
	// Type which this should extend
	Extends Reference `kego:"{\"default\": \"github.com/kego/system:object\"}"`
	// Is this type an interface?
	Interface Bool `kego:"{\"default\": false}"`
	// Array of interface types that this type should support
	Is []Reference
	// This is the native json type that represents this type. If omitted, default is object.
	Native     String `kego:"{\"default\": \"object\"}"`
	Properties map[string]*Property
	Rule       *Type
}

func init() {
	json.RegisterType("github.com/kego/system:object", reflect.TypeOf(&Object{}))
	json.RegisterType("github.com/kego/system:type", reflect.TypeOf(&Type{}))
	json.RegisterType("github.com/kego/system:@bool", reflect.TypeOf(&Bool_rule{}))
}
