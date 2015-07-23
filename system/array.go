package system

import (
	"reflect"

	"fmt"

	"kego.io/kerr"
)

func (r *Array_rule) Enforce(data interface{}, path string, aliases map[string]string) (bool, string, error) {

	if !r.MaxItems.Exists && !r.MinItems.Exists && !r.UniqueItems {
		// We should return early here in order to prevent needless reflection
		return true, "", nil
	}

	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Slice {
		return false, "", kerr.New("NFWPLTOJLP", nil, "Array_rule.Enforce", "val.Kind %s should be slice.", val.Kind())
	}

	// This is the maximum number of items allowed in the array
	// MaxItems Int
	if r.MaxItems.Exists {
		if val.Len() > r.MaxItems.Value {
			return false, fmt.Sprintf("MaxItems: length %d should not be greater than %d", val.Len(), r.MaxItems.Value), nil
		}
	}

	// This is the minimum number of items allowed in the array
	// MinItems Int
	if r.MinItems.Exists {
		if val.Len() < r.MinItems.Value {
			return false, fmt.Sprintf("MinItems: length %d should not be less than %d", val.Len(), r.MinItems.Value), nil
		}
	}

	// If this is true, each item must be unique
	// UniqueItems Bool `kego:"{\"default\":{\"value\":false}}"`
	if r.UniqueItems {
		for i := 0; i < val.Len(); i++ {
			for j := 0; j < val.Len(); j++ {
				if i == j {
					continue
				}
				if reflect.DeepEqual(val.Index(i).Interface(), val.Index(j).Interface()) {
					return false, fmt.Sprintf("UniqueItems: array contains duplicate item %v", val.Index(i).Interface()), nil
				}
			}
		}
	}

	return true, "", nil
}