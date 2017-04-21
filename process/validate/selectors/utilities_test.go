package selectors

import (
	"testing"

	"context"

	"frizz.io/process/parser"
	"frizz.io/system"
	"frizz.io/system/node"
	"frizz.io/tests"
	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
	"github.com/go-errors/errors"
)

func TestGetIndex(t *testing.T) {
	p := getUtilParser(t)
	r := p.root
	i := p.getIndex(r)
	assert.Equal(t, -1, i)
	s := p.getSiblings(r)
	assert.Equal(t, 0, s)
}

func TestGetFloat64(t *testing.T) {

	n := &node.Node{ValueNumber: 1.1}
	f := getFloat64(n)
	assert.Equal(t, 1.1, f)

	f = getFloat64(nil)
	assert.Equal(t, 0.0, f)

	f = getFloat64("1.1")
	assert.Equal(t, 1.1, f)

	f = getFloat64("a")
	assert.Equal(t, -1.0, f)

	f = getFloat64(2)
	assert.Equal(t, 2.0, f)
}

func TestGetInt32(t *testing.T) {
	i := getInt32(2.0)
	assert.Equal(t, int32(2), i)

	i = getInt32(2)
	assert.Equal(t, int32(2), i)

	i = getInt32("a")
	assert.Equal(t, int32(-1), i)
}

func TestIsNull(t *testing.T) {
	assert.True(t, isNull(exprElement{typ: system.J_NULL, value: nil}))
	assert.True(t, isNull(exprElement{typ: system.J_MAP, value: nil}))
	assert.True(t, isNull(exprElement{typ: system.J_MAP, value: &node.Node{Null: true}}))
	assert.True(t, isNull(exprElement{typ: system.J_MAP, value: &node.Node{Missing: true}}))
	assert.False(t, isNull(exprElement{typ: system.J_MAP, value: &node.Node{}}))
	assert.False(t, isNull(exprElement{typ: system.J_MAP, value: 1}))
}

func TestGetBool(t *testing.T) {
	assert.False(t, getBool(nil))
	assert.False(t, getBool(&node.Node{ValueBool: false}))
	assert.True(t, getBool(&node.Node{ValueBool: true}))
	assert.True(t, getBool(system.NewBool(true)))
	assert.False(t, getBool(system.NewBool(false)))
	assert.True(t, getBool(true))
	assert.False(t, getBool(false))
	assert.False(t, getBool("a"))
}

func TestExprElementIsTruthy(t *testing.T) {
	assert.False(t, exprElementIsTruthy(exprElement{typ: system.J_STRING, value: ""}))
	assert.True(t, exprElementIsTruthy(exprElement{typ: system.J_STRING, value: "a"}))
	assert.False(t, exprElementIsTruthy(exprElement{typ: system.J_NUMBER, value: 0.0}))
	assert.True(t, exprElementIsTruthy(exprElement{typ: system.J_NUMBER, value: 0.1}))
	assert.True(t, exprElementIsTruthy(exprElement{typ: system.J_OBJECT}))
	assert.True(t, exprElementIsTruthy(exprElement{typ: system.J_MAP}))
	assert.True(t, exprElementIsTruthy(exprElement{typ: system.J_ARRAY}))
	assert.False(t, exprElementIsTruthy(exprElement{typ: system.J_BOOL, value: false}))
	assert.True(t, exprElementIsTruthy(exprElement{typ: system.J_BOOL, value: true}))
	assert.False(t, exprElementIsTruthy(exprElement{typ: system.J_NULL}))
	assert.False(t, exprElementIsTruthy(exprElement{typ: system.J_OPERATOR}))

}

func TestGetJsonString(t *testing.T) {
	assert.Equal(t, "", getJsonString(nil))
	assert.Equal(t, "{}", getJsonString(&foo{}))
}

type foo struct{}

func (f *foo) MarshalJSON(ctx context.Context) ([]byte, error) {
	return []byte("a"), errors.New("b")
}

func getUtilParser(t *testing.T) *Parser {
	cb := tests.Context("a.b/c").Ssystem(parser.Parse)
	data := `{
		"type": "system:type",
		"id": "foo"
	}`
	n, err := node.Unmarshal(cb.Ctx(), []byte(data))
	require.NoError(t, err)
	p, err := CreateParser(cb.Ctx(), n)
	require.NoError(t, err)
	return p
}
