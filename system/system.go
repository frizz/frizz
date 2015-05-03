//go:generate kego
package system // import "kego.io/system"

import (
	"fmt"
	"strings"
	"sync"
)

func IdToGoReference(id string, packagePath string, localImports map[string]string, localPackagePath string) (string, error) {
	typeName := IdToGoName(id)
	return GoReference(typeName, packagePath, localImports, localPackagePath)
}
func GoReference(typeName string, packagePath string, localImports map[string]string, localPackagePath string) (string, error) {
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
	return "", fmt.Errorf("Error in system.IdToGoReference: package path %v not found in local context imports.\n", packagePath)
}

func IdToGoName(id string) string {
	if strings.HasPrefix(id, "@") {
		// Rules are prefixed with @ in the json ID, and suffixed with _rule in the
		// golang type name.
		return fmt.Sprintf("%v%v_rule", strings.ToUpper(id[1:2]), id[2:])
	} else {
		return fmt.Sprintf("%v%v", strings.ToUpper(id[0:1]), id[1:])
	}
}

var types struct {
	sync.RWMutex
	m map[string]*Type
}

func RegisterType(name string, typ *Type) {
	types.Lock()
	if types.m == nil {
		types.m = make(map[string]*Type)
	}
	types.m[name] = typ
	types.Unlock()
}
func UnregisterType(name string) {
	types.Lock()
	if types.m == nil {
		return
	}
	delete(types.m, name)
	types.Unlock()
}
func GetType(name string) (*Type, bool) {
	types.RLock()
	t, ok := types.m[name]
	types.RUnlock()
	if !ok {
		return nil, false
	}
	return t, true
}
