package selectors

// notest

import (
	"fmt"
	"log"
	"os"
	"strings"

	"kego.io/system/node"
)

type logHandler struct {
	Enabled        bool
	recursionLevel int
	prefixes       map[int]string
}

//var logger = logHandler{true, 0, map[int]string{}}

var logger = logHandler{false, 0, nil}
var handler = log.New(os.Stderr, "selectors: ", 0)
var recursionMarker = "⇢ "

func (l *logHandler) formatPrefix(a ...interface{}) []interface{} {
	var arguments []interface{}
	arguments = append(
		arguments,
		strings.Repeat(recursionMarker, l.recursionLevel),
	)
	prefix, ok := l.prefixes[l.recursionLevel]
	if ok {
		arguments = append(
			arguments,
			prefix,
		)
	}
	arguments = append(
		arguments,
		a...,
	)
	return arguments
}

func (l *logHandler) Print(a ...interface{}) {
	if logger.Enabled {
		handler.Print(l.formatPrefix(a...)...)
	}
}

func (l *logHandler) Println(a ...interface{}) {
	if logger.Enabled {
		handler.Println(l.formatPrefix(a...)...)
	}
}

func (l *logHandler) IncreaseDepth() {
	if logger.Enabled {
		l.recursionLevel++
	}
}

func (l *logHandler) DecreaseDepth() {
	if logger.Enabled {
		l.recursionLevel--
	}
}

func (l *logHandler) SetPrefix(prefix ...interface{}) {
	if logger.Enabled {
		l.prefixes[l.recursionLevel] = fmt.Sprint(prefix...)
	}
}

func (l *logHandler) ClearPrefix() {
	if logger.Enabled {
		l.prefixes[l.recursionLevel] = ""
	}
}

func EnableLogger() {
	logger.prefixes = make(map[int]string)
	logger.Enabled = true
}

func getFormattedNodeMap(nodes map[*node.Node]*node.Node) []string {
	output := make([]*node.Node, 0, len(nodes))
	for _, val := range nodes {
		output = append(output, val)
	}
	return getFormattedNodeArray(output)
}

func getFormattedNodeArray(nodes []*node.Node) []string {
	var formatted []string
	for _, n := range nodes {
		if n != nil {
			formatted = append(formatted, fmt.Sprint(*n))
		} else {
			formatted = append(formatted, fmt.Sprint(nil))
		}
	}
	return formatted
}

func getFormattedTokens(tokens []*token) []string {
	var output []string
	for _, token := range tokens {
		output = append(output, fmt.Sprint(token.val))
	}
	return output
}

func getFormattedExpression(tokens []*exprElement) []string {
	var output []string
	for _, token := range tokens {
		output = append(output, fmt.Sprint(token.value))
	}
	return output
}
