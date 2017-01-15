package builder

import (
	"fmt"

	"github.com/davelondon/jennifer/jen"
	"kego.io/system"
)

func Reference(reference *system.Reference) *jen.Statement {
	if reference.Package == "kego.io/json" {
		switch reference.Name {
		case "string":
			return jen.String()
		case "number":
			return jen.Float64()
		case "bool":
			return jen.Bool()
		}
	}
	return jen.Id(fmt.Sprintf("%s.%s", reference.Package, system.GoName(reference.Name)))
}

func InterfaceReference(reference *system.Reference) *jen.Statement {
	return jen.Id(fmt.Sprintf("%s.%s", reference.Package, system.GoInterfaceName(reference.Name)))
}

/*
func Reference(path string, name string, localPath string, getAlias func(string) string) string {
	if path == localPath {
		return name
	}
	if path == "kego.io/json" {
		if name == "String" {
			// Built-in native json string type
			return "string"
		}
		if name == "Number" {
			// Built-in native json number type
			return "float64"
		}
		if name == "Bool" {
			// Built-in native json bool type
			return "bool"
		}
	}
	return fmt.Sprintf("%s.%s", getAlias(path), name)
}
*/
