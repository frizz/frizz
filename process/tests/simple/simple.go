//go:generate ke kego.io/process/tests/simple
package simple // import "kego.io/process/tests/simple"

import (
	"golang.org/x/net/context"
	"kego.io/system"
)

var _ system.StringInterface = (*A)(nil)

func (a A) GetString(ctx context.Context) *system.String {
	return a.B
}

type C interface{}

var _ system.Enforcer = (*CRule)(nil)

func (r *CRule) Enforce(ctx context.Context, data interface{}) (success bool, message string, err error) {
	if r.Fail != nil && r.Fail.Value() {
		return false, "Fail", nil
	}
	return true, "", nil
}
