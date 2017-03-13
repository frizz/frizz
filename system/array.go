package system

import (
	"reflect"

	"fmt"

	"context"

	"github.com/dave/kerr"
)

func (r *ArrayRule) Enforce(ctx context.Context, data interface{}) (fail bool, messages []string, err error) {

	if r.MaxItems == nil && r.MinItems == nil && !r.UniqueItems {
		// We should return early here in order to prevent needless reflection
		return
	}

	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Slice {
		return true, nil, kerr.New("OWTAUVVFBL", "val.Kind %s should be slice.", val.Kind())
	}

	// This is the maximum number of items allowed in the array
	// MaxItems Int
	if r.MaxItems != nil {
		if val.Len() > r.MaxItems.Value() {
			fail = true
			messages = append(messages, fmt.Sprintf("MaxItems: length %d should not be greater than %d", val.Len(), r.MaxItems.Value()))
		}
	}

	// This is the minimum number of items allowed in the array
	// MinItems Int
	if r.MinItems != nil {
		if val.Len() < r.MinItems.Value() {
			fail = true
			messages = append(messages, fmt.Sprintf("MinItems: length %d should not be less than %d", val.Len(), r.MinItems.Value()))
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
					fail = true
					messages = append(messages, fmt.Sprintf("UniqueItems: array contains duplicate item %v", val.Index(i).Interface()))
				}
			}
		}
	}
	return
}

var _ Enforcer = (*ArrayRule)(nil)

func (a *ArrayRule) GetItemsRule() RuleInterface {
	return a.Items
}

var _ CollectionRule = (*ArrayRule)(nil)
