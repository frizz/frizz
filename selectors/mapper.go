package jsonselect

import (
	"kego.io/kerr"
	"kego.io/system"
)

type nativeType string

const (
	J_STRING  nativeType = "string"
	J_NUMBER  nativeType = "number"
	J_MAP     nativeType = "map"
	J_OBJ     nativeType = "object"
	J_ARRAY   nativeType = "array"
	J_BOOLEAN nativeType = "boolean"
	J_NULL    nativeType = "null"

	// Not actually a type, obviously
	J_OPER nativeType = "oper"
)

type jsonNode struct {
	value      interface{}
	native     nativeType
	ktype      system.Reference
	json       *Json
	parent     *jsonNode
	parent_key string
	idx        int
	siblings   int
}

func (p *Parser) getFlooredDocumentMap(node *jsonNode) ([]*jsonNode, error) {
	var newMap []*jsonNode
	newMap, err := p.getNodes(node.json, newMap, nil, "", -1, -1)
	if err != nil {
		return nil, kerr.New("XDXAAIMPLH", err, "jsonselect.getFlooredDocumentMap", "getNodes")
	}

	if logger.Enabled {
		logger.Print("Floored document map for ", node, " reduced node count to ", len(newMap))
	}

	return newMap, nil
}

func (p *Parser) getNodes(jdoc *Json, nodes []*jsonNode, parent *jsonNode, parent_key string, idx int, siblings int) ([]*jsonNode, error) {
	node := jsonNode{}
	node.parent = parent
	node.json = jdoc
	if len(parent_key) > 0 {
		node.parent_key = parent_key
	}
	if idx > -1 {
		node.idx = idx
	}
	if siblings > -1 {
		node.siblings = siblings
	}

	node.ktype = jdoc.Rule.ParentType.Type

	switch jdoc.Rule.ParentType.Native.Value {
	case "bool":
		value, exists := jdoc.Data.(system.NativeBool).NativeBool()
		if !exists {
			node.value = nil
			node.native = J_NULL
			break
		}
		node.value = value
		node.native = J_BOOLEAN
		break
	case "number":
		value, exists := jdoc.Data.(system.NativeNumber).NativeNumber()
		if !exists {
			node.value = nil
			node.native = J_NULL
			break
		}
		node.value = value
		node.native = J_NUMBER
		break
	case "string":
		value, exists := jdoc.Data.(system.NativeString).NativeString()
		if !exists {
			node.value = nil
			node.native = J_NULL
			break
		}
		node.value = value
		node.native = J_STRING
		break
	case "array":
		node.value = jdoc.Data
		node.native = J_ARRAY
		length := jdoc.Value.Len()
		itemsRule, err := jdoc.Rule.ItemsRule()
		if err != nil {
			return nil, kerr.New("PXGPNCVEFH", err, "jsonselect.getNodes", "jdoc.Rule.ItemsRule (array)")
		}
		for i := 0; i < length; i++ {
			object, _, value, found, _, err := system.GetArrayMember(jdoc.Value, i)
			if err != nil {
				return nil, kerr.New("UNSRRWKJTM", err, "jsonselect.getNodes", "system.GetArrayMember")
			}
			if !found {
				continue
			}
			child := &Json{Data: object, Rule: itemsRule, Value: value}
			// TODO: Why is the index i+1 here? Take time to understand this!
			nodes, err = p.getNodes(child, nodes, &node, "", i+1, length)
			if err != nil {
				return nil, kerr.New("LFQQBAJXJU", err, "jsonselect.getNodes", "getNodes (array)")
			}
		}
		break
	case "map":
		node.value = jdoc.Data
		node.native = J_MAP
		itemsRule, err := jdoc.Rule.ItemsRule()
		if err != nil {
			return nil, kerr.New("SYGEFDHBTO", err, "jsonselect.getNodes", "jdoc.Rule.ItemsRule (map)")
		}
		for _, k := range jdoc.Value.MapKeys() {
			key, ok := k.Interface().(string)
			if !ok {
				return nil, kerr.New("QYPJSPHNRN", nil, "jsonselect.getNodes", "Map nodes must be strings, not %T", k.Interface())
			}
			object, _, value, found, _, err := system.GetMapMember(jdoc.Value, key)
			if err != nil {
				return nil, kerr.New("JMUJMBBLWU", err, "jsonselect.getNodes", "system.GetMapMember")
			}
			if !found {
				continue
			}
			child := &Json{Data: object, Rule: itemsRule, Value: value}
			nodes, err = p.getNodes(child, nodes, &node, key, -1, -1)
			if err != nil {
				return nil, kerr.New("HEMMRWSOEU", err, "jsonselect.getNodes", "getNodes (map)")
			}
		}
		break
	case "object":
		node.value = jdoc.Data
		node.native = J_OBJ
		for key, property := range jdoc.Rule.ParentType.Properties {
			object, _, value, found, _, err := system.GetObjectField(jdoc.Value, system.IdToGoName(key))
			if err != nil {
				return nil, kerr.New("JMUJMBBLWU", err, "jsonselect.getNodes", "system.GetMapMember")
			}
			if !found {
				continue
			}
			itemRule, err := system.NewRuleHolder(property.Item, p.path, p.imports)
			if err != nil {
				return nil, kerr.New("WWQNGYIUKN", err, "jsonselect.getNodes", "system.NewRuleHolder")
			}
			child := &Json{Data: object, Rule: itemRule, Value: value}
			nodes, err = p.getNodes(child, nodes, &node, key, -1, -1)
			if err != nil {
				return nil, kerr.New("NQWTHARJES", err, "jsonselect.getNodes", "p.getNodes")
			}
		}
		break
	}
	nodes = append(nodes, &node)
	return nodes, nil
}

func (p *Parser) mapDocument() error {
	var nodes []*jsonNode
	n, err := p.getNodes(p.Data, nodes, nil, "", -1, -1)
	if err != nil {
		return kerr.New("DBFKOECICC", err, "jsonselect.mapDocument", "p.getNodes")
	}
	p.nodes = n
	return nil
}
