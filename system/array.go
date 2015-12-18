package system

import (
	"reflect"

	"fmt"

	"golang.org/x/net/context"
	"kego.io/kerr"
)

func (r *ArrayRule) Enforce(ctx context.Context, data interface{}) (bool, string, error) {

	if r.MaxItems == nil && r.MinItems == nil && !r.UniqueItems {
		// We should return early here in order to prevent needless reflection
		return true, "", nil
	}

	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Slice {
		return false, "", kerr.New("NFWPLTOJLP", nil, "val.Kind %s should be slice.", val.Kind())
	}

	// This is the maximum number of items allowed in the array
	// MaxItems Int
	if r.MaxItems != nil {
		if val.Len() > r.MaxItems.Value() {
			return false, fmt.Sprintf("MaxItems: length %d should not be greater than %d", val.Len(), r.MaxItems.Value()), nil
		}
	}

	// This is the minimum number of items allowed in the array
	// MinItems Int
	if r.MinItems != nil {
		if val.Len() < r.MinItems.Value() {
			return false, fmt.Sprintf("MinItems: length %d should not be less than %d", val.Len(), r.MinItems.Value()), nil
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

var _ Enforcer = (*ArrayRule)(nil)

func (a *ArrayRule) GetItemsRule() RuleInterface {
	return a.Items
}

var _ CollectionRule = (*ArrayRule)(nil)
