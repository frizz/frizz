//go:generate ke kego.io/process/validate/tests
package tests // import "kego.io/process/validate/tests"

// ke: {"package": {"notest": true}}

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

func (r *CRule) Enforce(ctx context.Context, data interface{}) (fail bool, messages []string, err error) {
	if r.Fail != nil && r.Fail.Value() {
		fail = true
		messages = append(messages, "Fail")
	}
	return
}
