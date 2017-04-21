package selectors

import (
	"testing"

	"frizz.io/process/parser"
	"frizz.io/system/node"
	"frizz.io/tests"
	"github.com/dave/ktest/assert"
	"github.com/dave/ktest/require"
)

func TestUniversalProduction(t *testing.T) {
	p := getParser(t)

	f := p.universalProduction("a")
	out, err := f(&node.Node{})
	require.NoError(t, err)
	assert.False(t, out)

	f = p.universalProduction("*")
	out, err = f(&node.Node{})
	require.NoError(t, err)
	assert.True(t, out)
}

func TestFrizzProduction(t *testing.T) {
	p := getParser(t)
	f := p.frizzProduction(nil)
	b, err := f(&node.Node{Null: true})
	assert.False(t, b)
	require.NoError(t, err)

}
func TestMatch(t *testing.T) {
	p := getParser(t)
	_, _, err := p.match([]*token{}, S_WORD)
	assert.IsError(t, err, "EGHVMCOKCS")
}

func TestGetValue(t *testing.T) {
	p := getParser(t)
	out, err := p.GetValues(":root")
	require.NoError(t, err)
	assert.Equal(t, 1, len(out))
}
func TestGetNodes(t *testing.T) {
	p := getParser(t)

	out, err := p.GetNodes(":root")
	require.NoError(t, err)
	assert.Equal(t, 1, len(out))

	// The lexer extracts a S_WORD token, which has no selector
	out, err = p.GetNodes("\"foo\"")
	assert.IsError(t, err, "QFXXVGGHSS")
	assert.HasError(t, err, "REPJTLCAMQ")

	// The lexer extracts a S_FLOAT token, which has no selector
	out, err = p.GetNodes("1.1")
	assert.IsError(t, err, "QFXXVGGHSS")
	assert.HasError(t, err, "REPJTLCAMQ")

	_, err = p.GetNodes(".a * * .b")
	assert.IsError(t, err, "QFXXVGGHSS")
	assert.HasError(t, err, "KLEORWJHSP")

}

func getParser(t *testing.T) *Parser {
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
