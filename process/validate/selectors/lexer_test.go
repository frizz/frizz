package selectors

import (
	"testing"

	"kego.io/kerr/assert"
)

func TestLexUnterminated(t *testing.T) {
	_, err := lex("(()", selectorScanner)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Unterminated expression")
}

func TestLexFloat(t *testing.T) {
	tokens, err := lex("1.1", selectorScanner)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, S_FLOAT, tokens[0].typ)
}

func TestLexWord(t *testing.T) {
	tokens, err := lex("\"a\"", selectorScanner)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, S_WORD, tokens[0].typ)
}
