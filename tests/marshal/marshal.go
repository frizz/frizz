package marshal

import (
	"context"
	"testing"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/ke"
)

type TestValue func(t *testing.T, v interface{})

type MarshalledString string

type UnmarshalError func(t *testing.T, err error)

func Run(t *testing.T, ctx context.Context, json string, options ...interface{}) {
	var v interface{}
	var testValue TestValue
	var unmarshalError UnmarshalError
	var marshalledString MarshalledString
	for _, op := range options {
		if v, ok := op.(TestValue); ok {
			testValue = v
			continue
		}
		if v, ok := op.(MarshalledString); ok {
			marshalledString = v
			continue
		}
		if v, ok := op.(UnmarshalError); ok {
			unmarshalError = v
			continue
		}
		panic("unknown option")
	}
	err := ke.Unmarshal(ctx, []byte(json), &v)
	if unmarshalError == nil {
		require.NoError(t, err)
	} else {
		unmarshalError(t, err)
		return
	}
	if testValue != nil {
		testValue(t, v)
	}
	b, err := ke.Marshal(ctx, v)
	require.NoError(t, err)
	if marshalledString != "" {
		assert.JSONEq(t, string(marshalledString), string(b))
	} else {
		assert.JSONEq(t, json, string(b))
	}
}
