//go:generate kego
package system // import "kego.io/system"
import (
	"fmt"
	"strings"

	"kego.io/uerr"
)

func IdToGoName(id string) string {
	if strings.HasPrefix(id, "@") {
		// Rules are prefixed with @ in the json ID, and suffixed with _rule in the
		// golang type name.
		return fmt.Sprintf("%v%v_rule", strings.ToUpper(id[1:2]), id[2:])
	} else {
		return fmt.Sprintf("%v%v", strings.ToUpper(id[0:1]), id[1:])
	}
}
func IdToGoReference(packagePath string, id string, localPackagePath string, localImports map[string]string) (string, error) {
	typeName := IdToGoName(id)
	return GoReference(packagePath, typeName, localPackagePath, localImports)
}
func GoReference(packagePath string, typeName string, localPackagePath string, localImports map[string]string) (string, error) {
	if packagePath == localPackagePath {
		return typeName, nil
	}
	if packagePath == "kego.io/json" {
		if typeName == "String" {
			// Built-in native json string type
			return "string", nil
		}
		if typeName == "Number" {
			// Built-in native json number type
			return "float64", nil
		}
		if typeName == "Bool" {
			// Built-in native json bool type
			return "bool", nil
		}
		return fmt.Sprintf("json.%v", typeName), nil
	}
	if packagePath == "kego.io/system" {
		return fmt.Sprintf("system.%v", typeName), nil
	}
	for alias, path := range localImports {
		if packagePath == path {
			return fmt.Sprintf("%v.%v", alias, typeName), nil
		}
	}
	return "", uerr.New("HXHJIHQDBE", nil, "GoReference", "package path %v not found in local context imports", packagePath)
}
