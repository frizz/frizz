//go:generate frizz frizz.io/process/validate/tests
package tests // import "frizz.io/process/validate/tests"

// notest

import (
	"context"

	"frizz.io/system"
)

var _ system.StringInterface = (*A)(nil)

func (a *A) GetString(ctx context.Context) *system.String {
	if a == nil {
		return nil
	}
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
