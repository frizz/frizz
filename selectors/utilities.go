package selectors

import (
	"log"
	"strconv"

	"fmt"

	"kego.io/json"
	"kego.io/system"
	"kego.io/system/node"
)

func nodeIsMemberOfHaystack(needle *node.Node, haystack map[*node.Node]*node.Node) bool {
	_, ok := haystack[needle]
	return ok
}

func nodeIsMemberOfList(needle *node.Node, haystack []*node.Node) bool {
	for _, element := range haystack {
		if element == needle {
			return true
		}
	}
	return false
}

func (p *Parser) appendAncestorsToHaystack(n *node.Node, haystack map[*node.Node]*node.Node) {
	if !p.isRoot(n) {
		haystack[n.Parent.GetNode()] = n.Parent.GetNode()
		p.appendAncestorsToHaystack(n.Parent.GetNode(), haystack)
	}
}

func (p *Parser) nodeIsChildOfHaystackMember(needle *node.Node, haystack map[*node.Node]*node.Node) bool {
	if nodeIsMemberOfHaystack(needle, haystack) {
		return true
	}
	if p.isRoot(needle) {
		return false
	}
	return p.nodeIsChildOfHaystackMember(needle.Parent.GetNode(), haystack)
}

func (p *Parser) parents(lhs []*node.Node, rhs []*node.Node) []*node.Node {
	var results []*node.Node

	lhsHaystack := getHaystackFromNodeList(lhs)

	for _, element := range rhs {
		if !p.isRoot(element) && nodeIsMemberOfHaystack(element.Parent.GetNode(), lhsHaystack) {
			results = append(results, element)
		}
	}

	return results
}
func (p *Parser) isRoot(n *node.Node) bool {
	if n.Parent == nil {
		return true
	}
	if n == p.root {
		return true
	}
	return false
}
func (p *Parser) getKey(n *node.Node) string {
	if p.isRoot(n) {
		return ""
	}
	return n.Key
}
func (p *Parser) getIndex(n *node.Node) int {
	if p.isRoot(n) {
		return -1
	}
	return n.Index + 1
}
func (p *Parser) getSiblings(n *node.Node) int {
	if p.isRoot(n) {
		return 0
	}
	return len(n.Parent.GetNode().Array)
}

func (p *Parser) ancestors(lhs []*node.Node, rhs []*node.Node) []*node.Node {
	var results []*node.Node
	haystack := getHaystackFromNodeList(lhs)

	for _, element := range rhs {
		if p.nodeIsChildOfHaystackMember(element, haystack) {
			results = append(results, element)
		}
	}

	return results
}

func (p *Parser) siblings(lhs []*node.Node, rhs []*node.Node) []*node.Node {
	var results []*node.Node
	parents := make(map[*node.Node]*node.Node, len(lhs))

	for _, element := range lhs {
		if !p.isRoot(element) {
			parents[element.Parent.GetNode()] = element.Parent.GetNode()
		}
	}

	for _, element := range rhs {
		if !p.isRoot(element) && nodeIsMemberOfHaystack(element.Parent.GetNode(), parents) {
			results = append(results, element)
		}
	}

	return results
}

func getFloat64(in interface{}) float64 {
	if in == nil {
		return 0.0
	}
	if as_node, ok := in.(*node.Node); ok {
		return as_node.ValueNumber
	}
	if as_nativeNum, ok := in.(system.NativeNumber); ok {
		return as_nativeNum.NativeNumber()
	}
	if as_float, ok := in.(float64); ok {
		return as_float
	}
	if as_int, ok := in.(int64); ok {
		value := float64(as_int)
		return value
	}
	if as_string, ok := in.(string); ok {
		parsed_float_string, err := strconv.ParseFloat(as_string, 64)
		if err == nil {
			value := parsed_float_string
			return value
		}
		parsed_int_string, err := strconv.ParseInt(as_string, 10, 32)
		if err == nil {
			value := float64(parsed_int_string)
			return value
		}
	}
	panic(fmt.Sprintf("ERR: %T", in))
	result := float64(-1)
	log.Print("Error transforming ", in, " into Float64")
	return result
}

func getInt32(in interface{}) int32 {
	value := int32(getFloat64(in))
	if value == -1 {
		log.Print("Error transforming ", in, " into Int32")
	}
	return value
}

func isNull(in exprElement) bool {
	if in.typ == json.J_NULL {
		return true
	}
	if in.value == nil {
		return true
	}
	if n, ok := in.value.(*node.Node); ok {
		return n.Null || n.Missing
	}
	return false
}

func getBool(in interface{}) bool {
	if in == nil {
		return false
	}
	if as_node, ok := in.(*node.Node); ok {
		return as_node.ValueBool
	}
	if as_nativeBool, ok := in.(system.NativeBool); ok {
		return as_nativeBool.NativeBool()
	}
	if as_bool, ok := in.(bool); ok {
		return as_bool
	}
	return false
}

func getJsonString(in interface{}) string {
	if in == nil {
		return ""
	}
	if as_node, ok := in.(*node.Node); ok {
		return as_node.ValueString
	}
	if as_nativeStr, ok := in.(system.NativeString); ok {
		return as_nativeStr.NativeString()
	}
	if as_string, ok := in.(string); ok {
		return as_string
	}
	marshaled_result, err := json.Marshal(in)
	if err != nil {
		log.Print("Error transforming ", in, " into JSON string")
	}
	result := string(marshaled_result)
	return result
}

func exprElementIsTruthy(e exprElement) (bool, error) {
	switch e.typ {
	case json.J_STRING:
		return len(e.value.(string)) > 0, nil
	case json.J_NUMBER:
		return e.value.(float64) > 0, nil
	case json.J_OBJECT:
		return true, nil
	case json.J_MAP:
		return true, nil
	case json.J_ARRAY:
		return true, nil
	case json.J_BOOL:
		return e.value.(bool), nil
	case json.J_NULL:
		return false, nil
	default:
		return false, nil
	}
}

func exprElementsMatch(lhs exprElement, rhs exprElement) bool {
	return lhs.typ == rhs.typ
}

func getHaystackFromNodeList(nodes []*node.Node) map[*node.Node]*node.Node {
	hashmap := make(map[*node.Node]*node.Node, len(nodes))
	for _, node := range nodes {
		hashmap[node] = node
	}
	return hashmap
}
