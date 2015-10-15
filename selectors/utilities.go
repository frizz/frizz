package selectors

import (
	"log"
	"strconv"

	"kego.io/json"
)

func nodeIsMemberOfHaystack(needle *node, haystack map[*Element]*node) bool {
	_, ok := haystack[needle.element]
	return ok
}

func nodeIsMemberOfList(needle *node, haystack []*node) bool {
	for _, element := range haystack {
		if element.element == needle.element {
			return true
		}
	}
	return false
}

func appendAncestorsToHaystack(n *node, haystack map[*Element]*node) {
	if n.Parent != nil {
		haystack[n.Parent.element] = n.Parent
		appendAncestorsToHaystack(n.Parent, haystack)
	}
}

func nodeIsChildOfHaystackMember(needle *node, haystack map[*Element]*node) bool {
	if nodeIsMemberOfHaystack(needle, haystack) {
		return true
	}
	if needle.Parent == nil {
		return false
	}
	return nodeIsChildOfHaystackMember(needle.Parent, haystack)
}

func parents(lhs []*node, rhs []*node) []*node {
	var results []*node

	lhsHaystack := getHaystackFromNodeList(lhs)

	for _, element := range rhs {
		if element.Parent != nil && nodeIsMemberOfHaystack(element.Parent, lhsHaystack) {
			results = append(results, element)
		}
	}

	return results
}

func ancestors(lhs []*node, rhs []*node) []*node {
	var results []*node
	haystack := getHaystackFromNodeList(lhs)

	for _, element := range rhs {
		if nodeIsChildOfHaystackMember(element, haystack) {
			results = append(results, element)
		}
	}

	return results
}

func siblings(lhs []*node, rhs []*node) []*node {
	var results []*node
	parents := make(map[*Element]*node, len(lhs))

	for _, element := range lhs {
		if element.Parent != nil {
			parents[element.Parent.element] = element.Parent
		}
	}

	for _, element := range rhs {
		if element.Parent != nil && nodeIsMemberOfHaystack(element.Parent, parents) {
			results = append(results, element)
		}
	}

	return results
}

func getFloat64(in interface{}) float64 {
	as_float, ok := in.(float64)
	if ok {
		return as_float
	}
	as_int, ok := in.(int64)
	if ok {
		value := float64(as_int)
		return value
	}
	as_string, ok := in.(string)
	if ok {
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
	panic("ERR")
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

func getJsonString(in interface{}) string {
	as_string, ok := in.(string)
	if ok {
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

func getHaystackFromNodeList(nodes []*node) map[*Element]*node {
	hashmap := make(map[*Element]*node, len(nodes))
	for _, node := range nodes {
		hashmap[node.element] = node
	}
	return hashmap
}
