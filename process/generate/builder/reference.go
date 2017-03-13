package builder

import (
	"github.com/dave/jennifer/jen"
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
	return jen.Qual(reference.Package, system.GoName(reference.Name))
}

func InterfaceReference(reference *system.Reference) *jen.Statement {
	return jen.Qual(reference.Package, system.GoInterfaceName(reference.Name))
}
