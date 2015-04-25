//go:generate kego
package system // import "kego.io/system"

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"kego.io/json"
)

type Defaulter interface {
	HasDefault() bool
	GetDefault() ([]byte, error)
}

// Basic interface implemented by everything that is composed of an Object
type Basic interface {
	Base() *Object
}

// SetContext satisfies the json.Contexter interface, which allows the json unmarshal
// function to store the unmarshal context in every object.
func (o *Object) SetContext(context *json.Context) {
	o.Context = context
}

func IdToGoReference(id string, packagePath string, localImports map[string]string, localPackagePath string) (string, error) {
	typeName := IdToGoName(id)
	if packagePath == localPackagePath {
		return typeName, nil
	} else {

		// TODO: move this to a util package
		searchMapForValue := func(in map[string]string, value string) (key string, found bool) {
			for k, v := range in {
				if v == value {
					return k, true
				}
			}
			return "", false
		}

		localPackageName, ok := searchMapForValue(localImports, packagePath)
		if !ok {
			return "", fmt.Errorf("Error in system.IdToGoReference: package path %v not found in local context imports.\n", packagePath)
		}
		return fmt.Sprintf("%v.%v", localPackageName, typeName), nil
	}
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

type Context_rule struct {
	*Object
}

func (r *Context_rule) Base() *Object {
	return r.Object
}

func init() {
	json.RegisterType("kego.io/json:context", reflect.TypeOf(&json.Context{}))
	json.RegisterType("kego.io/json:@context", reflect.TypeOf(&Context_rule{}))
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
