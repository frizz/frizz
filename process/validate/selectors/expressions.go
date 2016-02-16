package selectors

import (
	"strings"

	"kego.io/json"
	"kego.io/system/node"
)

type exprElement struct {
	value interface{}
	typ   json.Type
}

var precedenceMap = map[string]int{
	"*":  1,
	"/":  1,
	"%":  1,
	"+":  2,
	"-":  2,
	"<=": 3,
	">=": 3,
	"<":  3,
	">":  3,
	"$=": 3,
	"^=": 3,
	"*=": 3,
	"=":  3,
	"!=": 3,
	"&&": 4,
	"||": 5,
}

var comparatorMap = map[string]func(lhs exprElement, rhs exprElement) exprElement{
	"*": func(lhs exprElement, rhs exprElement) exprElement {
		value := getFloat64(lhs.value) * getFloat64(rhs.value)
		return exprElement{value, json.J_NUMBER}
	},
	"/": func(lhs exprElement, rhs exprElement) exprElement {
		value := getFloat64(lhs.value) / getFloat64(rhs.value)
		return exprElement{value, json.J_NUMBER}
	},
	"%": func(lhs exprElement, rhs exprElement) exprElement {
		value := float64(getInt32(lhs.value) % getInt32(rhs.value))
		return exprElement{value, json.J_NUMBER}
	},
	"+": func(lhs exprElement, rhs exprElement) exprElement {
		value := getFloat64(lhs.value) + getFloat64(rhs.value)
		return exprElement{value, json.J_NUMBER}
	},
	"-": func(lhs exprElement, rhs exprElement) exprElement {
		value := getFloat64(lhs.value) - getFloat64(rhs.value)
		return exprElement{value, json.J_NUMBER}
	},
	"<=": func(lhs exprElement, rhs exprElement) exprElement {
		if getFloat64(lhs.value) <= getFloat64(rhs.value) {
			return exprElement{true, json.J_BOOL}
		}
		return exprElement{false, json.J_BOOL}
	},
	"<": func(lhs exprElement, rhs exprElement) exprElement {
		if getFloat64(lhs.value) < getFloat64(rhs.value) {
			return exprElement{true, json.J_BOOL}
		}
		return exprElement{false, json.J_BOOL}
	},
	">=": func(lhs exprElement, rhs exprElement) exprElement {
		if getFloat64(lhs.value) >= getFloat64(rhs.value) {
			return exprElement{true, json.J_BOOL}
		}
		return exprElement{false, json.J_BOOL}
	},
	">": func(lhs exprElement, rhs exprElement) exprElement {
		if getFloat64(lhs.value) > getFloat64(rhs.value) {
			return exprElement{true, json.J_BOOL}
		}
		return exprElement{false, json.J_BOOL}
	},
	"$=": func(lhs exprElement, rhs exprElement) exprElement {
		lhs_str := getJsonString(lhs.value)
		rhs_str := getJsonString(rhs.value)
		if strings.HasSuffix(lhs_str, rhs_str) {
			return exprElement{true, json.J_BOOL}
		}
		return exprElement{false, json.J_BOOL}
	},
	"^=": func(lhs exprElement, rhs exprElement) exprElement {
		lhs_str := getJsonString(lhs.value)
		rhs_str := getJsonString(rhs.value)
		if strings.Index(lhs_str, rhs_str) == 0 {
			return exprElement{true, json.J_BOOL}
		}
		return exprElement{false, json.J_BOOL}
	},
	"*=": func(lhs exprElement, rhs exprElement) exprElement {
		lhs_str := getJsonString(lhs.value)
		rhs_str := getJsonString(rhs.value)
		if strings.Index(lhs_str, rhs_str) != -1 {
			return exprElement{true, json.J_BOOL}
		}
		return exprElement{false, json.J_BOOL}
	},
	"=": func(lhs exprElement, rhs exprElement) exprElement {
		if isNull(lhs) || isNull(rhs) {
			result := isNull(lhs) && isNull(rhs)
			return exprElement{result, json.J_BOOL}
		}
		lhs_string := getJsonString(lhs.value)
		rhs_string := getJsonString(rhs.value)
		if lhs_string == rhs_string {
			return exprElement{true, json.J_BOOL}
		}
		return exprElement{false, json.J_BOOL}
	},
	"!=": func(lhs exprElement, rhs exprElement) exprElement {
		if isNull(lhs) || isNull(rhs) {
			// ke: {"block": {"notest": true}}
			result := !isNull(lhs) || !isNull(rhs)
			return exprElement{result, json.J_BOOL}
		}
		lhs_string := getJsonString(lhs.value)
		rhs_string := getJsonString(rhs.value)
		if lhs_string != rhs_string {
			return exprElement{true, json.J_BOOL}
		}
		return exprElement{false, json.J_BOOL}
	},
	"&&": func(lhs exprElement, rhs exprElement) exprElement {
		if getBool(lhs.value) && getBool(rhs.value) {
			return exprElement{true, json.J_BOOL}
		}
		return exprElement{false, json.J_BOOL}
	},
	"||": func(lhs exprElement, rhs exprElement) exprElement {
		if getBool(lhs.value) || getBool(rhs.value) {
			return exprElement{true, json.J_BOOL}
		}
		return exprElement{false, json.J_BOOL}
	},
}

func (p *Parser) evaluateExpressionWithPrecedence(elements []*exprElement, precedenceLevel int) []*exprElement {
	var newExpression []*exprElement
	var i int
	var element *exprElement
	for i = 0; i < len(elements); i++ {
		element = elements[i]
		if element.typ == json.J_OPERATOR && precedenceMap[element.value.(string)] == precedenceLevel {
			lhs := newExpression[len(newExpression)-1]
			rhs := elements[i+1]
			newExpression = newExpression[0 : len(newExpression)-1]
			if exprElementsMatch(*lhs, *rhs) {
				result := comparatorMap[element.value.(string)](*lhs, *rhs)
				newExpression = append(newExpression, &result)
			} else {
				// Should we match type1 != type2?
				//if element.value.(string) == "!=" {
				//	newExpression = append(newExpression, &exprElement{true, json.J_BOOL})
				//} else {
				logger.Print("Cannot compare ", lhs.value, " and ", rhs.value, "; types differ: ", rhs.typ, " != ", lhs.typ)
				newExpression = append(newExpression, &exprElement{false, json.J_BOOL})
				//}
			}
			// Skip ahead an additional step more than normal
			// since we just pulled off one more element than normal.
			i += 1
		} else {
			newExpression = append(newExpression, element)
		}
	}
	return newExpression
}

func (p *Parser) evaluateExpression(elements []*exprElement) exprElement {
	for i := 1; i <= 5; i++ {
		elements = p.evaluateExpressionWithPrecedence(elements, i)
	}
	if len(elements) > 1 {
		// ke: {"block": {"notest": true}}
		panic("More than one expression result")
	}
	return *elements[0]
}

func (p *Parser) parseExpression(tokens []*token, current *node.Node) exprElement {
	var finalTokens []*exprElement

	logger.Print("Parsing expression ", getFormattedTokens(tokens))

	// If this expression is wrapped with parentheses, let's remove them.
	if tokens[0].typ == S_PAREN && tokens[0].val == "(" {
		tokens = tokens[1 : len(tokens)-1]
	}

	for i := 0; i < len(tokens); i++ {
		thisToken := *tokens[i]
		if thisToken.typ == S_PAREN && thisToken.val == "(" {
			var leftCount = 0
			var rightCount = 0
			var j = 0
			var subExprToken *token
			for j, subExprToken = range tokens[i:] {
				if subExprToken.typ == S_PAREN && subExprToken.val == "(" {
					leftCount++
				} else if subExprToken.typ == S_PAREN && subExprToken.val == ")" {
					rightCount++
				}
				if leftCount == rightCount {
					break
				}
			}
			subexprResult := p.parseExpression(tokens[i:i+j+1], current)
			finalTokens = append(finalTokens, &subexprResult)
			// Let's move the cursor to the end of the subexpression above.
			i = i + j
		} else if thisToken.typ == S_PVAR {
			// This is for the value of the node being processed.
			finalTokens = append(finalTokens, &exprElement{current.Value, current.JsonType})
		} else if thisToken.typ == S_STRING || thisToken.typ == S_BOOL || thisToken.typ == S_NIL || thisToken.typ == S_NUMBER {
			// Let's copy these kinds of tokens down as expression elements.
			var jType json.Type
			switch thisToken.typ {
			case S_BOOL:
				jType = json.J_BOOL
			case S_NIL:
				jType = json.J_NULL
			case S_NUMBER:
				jType = json.J_NUMBER
			default:
				jType = json.J_STRING
			}
			finalTokens = append(finalTokens, &exprElement{thisToken.val, jType})
		} else if thisToken.typ == S_BINOP {
			finalTokens = append(finalTokens, &exprElement{thisToken.val, json.J_OPERATOR})
		} else {
			logger.Print("Unexpected token ", thisToken, " ignored.")
		}
	}

	return p.evaluateExpression(finalTokens)
}
