package system

import (
	"fmt"
	"reflect"

	"kego.io/kerr"
)

func (r *Map_rule) Enforce(data interface{}, path string, aliases map[string]string) (bool, string, error) {

	if !r.MaxItems.Exists && !r.MinItems.Exists {
		// We should return early here in order to prevent needless reflection
		return true, "", nil
	}

	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Map {
		return false, "", kerr.New("NFWPLTOJLP", nil, "Array_rule.Enforce", "val.Kind %s should be slice.", val.Kind())
	}

	// This is the maximum number of items allowed in the map
	// MaxItems Int
	if r.MaxItems.Exists {
		if val.Len() > r.MaxItems.Value {
			return false, fmt.Sprintf("MaxItems: length %d should not be greater than %d", val.Len(), r.MaxItems.Value), nil
		}
	}

	// This is the minimum number of items allowed in the map
	// MinItems Int
	if r.MinItems.Exists {
		if val.Len() < r.MinItems.Value {
			return false, fmt.Sprintf("MinItems: length %d should not be less than %d", val.Len(), r.MinItems.Value), nil
		}
	}

	return true, "", nil
}
