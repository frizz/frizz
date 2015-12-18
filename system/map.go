package system

import (
	"fmt"
	"reflect"

	"golang.org/x/net/context"

	"kego.io/kerr"
)

func (r *MapRule) Enforce(ctx context.Context, data interface{}) (bool, string, error) {

	if r.MaxItems == nil && r.MinItems == nil {
		// We should return early here in order to prevent needless reflection
		return true, "", nil
	}

	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Map {
		return false, "", kerr.New("NFWPLTOJLP", nil, "val.Kind %s should be slice.", val.Kind())
	}

	// This is the maximum number of items allowed in the map
	// MaxItems Int
	if r.MaxItems != nil {
		if val.Len() > r.MaxItems.Value() {
			return false, fmt.Sprintf("MaxItems: length %d should not be greater than %d", val.Len(), r.MaxItems.Value()), nil
		}
	}

	// This is the minimum number of items allowed in the map
	// MinItems Int
	if r.MinItems != nil {
		if val.Len() < r.MinItems.Value() {
			return false, fmt.Sprintf("MinItems: length %d should not be less than %d", val.Len(), r.MinItems.Value()), nil
		}
	}

	return true, "", nil
}

var _ Enforcer = (*MapRule)(nil)

func (m *MapRule) GetItemsRule() RuleInterface {
	return m.Items
}

var _ CollectionRule = (*MapRule)(nil)
