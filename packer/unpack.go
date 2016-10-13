package packer

import (
	"context"
	"strings"

	"kego.io/context/envctx"

	"fmt"

	"github.com/davelondon/kerr"
)

type unpackStruct struct {
	unknownType    string // have we encountered an unknown type?
	unknownPackage string // have we encountered an unknown package?
}

// Unpacker unpacks the data from in into the object. If iface is true, we're
// unpacking into an interface. In this situation, map types are always
// specified in typed form: {"type": "mytype", "value": {"foo": "bar"}}.
type Unpacker interface {
	Unpack(ctx context.Context, in Packed, iface bool) error
}

type UnknownPackageError struct {
	kerr.Struct
	UnknownPackage string
}

type UnknownTypeError struct {
	kerr.Struct
	UnknownType string
}

type InitializableType interface {
	InitializeType(path string, name string) error
}

// if we tried to unmarshal an incompatible type, we return this from the InitializeType
// funciton. We include the package path and name of the unmarshaled object.
type InitializableTypeError struct {
	UnmarshalledPath string
	UnmarshalledName string
	IntoPath         string
	IntoName         string
}

func (i InitializableTypeError) Error() string {
	return fmt.Sprintf("Tried to unmarshal %s:%s into %s:%s", i.UnmarshalledPath, i.UnmarshalledName, i.IntoPath, i.IntoName)
}

func GetReferencePartsFromTypeString(ctx context.Context, typeString string) (path string, name string, err error) {

	env := envctx.FromContext(ctx)

	if strings.Contains(typeString, "/") {
		// If the type name contains a slash, I'm assuming it's a fully
		// qualified type name of the form "kego.io/system:type".
		// TODO: Improve this with a regex?
		parts := strings.Split(typeString, ":")

		// We hard-code system and json to prevent them having to always be
		// specified in the aliases
		if parts[0] == "kego.io/system" {
			return "kego.io/system", parts[1], nil
		} else if parts[0] == "kego.io/json" {
			return "kego.io/json", parts[1], nil
		}

		_, found := findKey(env.Aliases, parts[0])
		if !found && parts[0] != env.Path {
			return "", "", UnknownPackageError{
				Struct:         kerr.New("KJSOXDESFD", "Unknown package %s", parts[0]),
				UnknownPackage: parts[0],
			}
		}
		return parts[0], parts[1], nil
	} else if strings.Contains(typeString, ":") {
		// If the type name contains a colon, I'm assuming it's an abreviated
		// qualified type name of the form "system:type". We should look the
		// package name up in the aliases map.
		// TODO: Improve this with a regex?
		parts := strings.Split(typeString, ":")

		// We hard-code system and json to prevent them having to always be
		// specified in the aliases
		if parts[0] == "system" {
			return "kego.io/system", parts[1], nil
		} else if parts[0] == "json" {
			return "kego.io/json", parts[1], nil
		}

		packagePath, ok := env.Aliases[parts[0]]
		if !ok {
			return "", "", UnknownPackageError{
				Struct:         kerr.New("DKKFLKDKYI", "Unknown package %s", parts[0]),
				UnknownPackage: parts[0],
			}
		}
		return packagePath, parts[1], nil
	} else {
		return env.Path, typeString, nil
	}
}
func findKey(m map[string]string, value string) (string, bool) {
	for k, v := range m {
		if value == v {
			return k, true
		}
	}
	return "", false
}
