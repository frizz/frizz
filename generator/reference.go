package generator

import "fmt"

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
