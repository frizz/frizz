package system

import (
	"fmt"
	"reflect"

	"context"

	"github.com/davelondon/kerr"
)

func (r *MapRule) Enforce(ctx context.Context, data interface{}) (fail bool, messages []string, err error) {

	if r.MaxItems == nil && r.MinItems == nil {
		// We should return early here in order to prevent needless reflection
		return
	}

	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Map {
		return true, nil, kerr.New("NFWPLTOJLP", "val.Kind %s should be slice.", val.Kind())
	}

	// This is the maximum number of items allowed in the map
	// MaxItems Int
	if r.MaxItems != nil {
		if val.Len() > r.MaxItems.Value() {
			fail = true
			messages = append(messages, fmt.Sprintf("MaxItems: length %d should not be greater than %d", val.Len(), r.MaxItems.Value()))
		}
	}

	// This is the minimum number of items allowed in the map
	// MinItems Int
	if r.MinItems != nil {
		if val.Len() < r.MinItems.Value() {
			fail = true
			messages = append(messages, fmt.Sprintf("MinItems: length %d should not be less than %d", val.Len(), r.MinItems.Value()))
		}
	}

	return
}

var _ Enforcer = (*MapRule)(nil)

func (m *MapRule) GetItemsRule() RuleInterface {
	return m.Items
}

var _ CollectionRule = (*MapRule)(nil)
