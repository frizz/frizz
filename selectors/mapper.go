package selectors

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

type node struct {
	value      interface{}
	native     nativeType
	ktyperef   system.Reference
	ktype      *system.Type
	element    *Element
	parent     *node
	parent_key string
	idx        int
	siblings   int
}

func (p *Parser) mapDocument() error {
	var nodes []*node
	n, err := p.getNodes(p.Data, nodes, nil, "", -1, -1)
	if err != nil {
		return kerr.New("DBFKOECICC", err, "p.getNodes")
	}
	p.nodes = n
	return nil
}

func (p *Parser) getFlooredDocumentMap(n *node) ([]*node, error) {
	var newMap []*node
	newMap, err := p.getNodes(n.element, newMap, nil, "", -1, -1)
	if err != nil {
		return nil, kerr.New("XDXAAIMPLH", err, "getNodes")
	}

	if logger.Enabled {
		logger.Print("Floored document map for ", n, " reduced node count to ", len(newMap))
	}

	return newMap, nil
}

func (p *Parser) getNodes(element *Element, nodes []*node, parent *node, parent_key string, idx int, siblings int) ([]*node, error) {
	n := node{}
	n.parent = parent
	n.element = element
	if len(parent_key) > 0 {
		n.parent_key = parent_key
	}
	if idx > -1 {
		n.idx = idx
	}
	if siblings > -1 {
		n.siblings = siblings
	}

	ob, ok := element.Data.(system.Object)
	if ok {
		base := ob.GetBase()
		n.ktyperef = base.Type
		n.ktype, ok = base.Type.GetType()
		if !ok {
			return nil, kerr.New("HGENTDRWHL", nil, "Type.GetType not found")
		}
	} else {
		n.ktyperef = element.Rule.ParentType.Id
		n.ktype = element.Rule.ParentType
	}

	switch n.ktype.Native.Value {
	case "bool":
		value, exists := element.Data.(system.NativeBool).NativeBool()
		if !exists {
			n.value = nil
			n.native = J_NULL
			break
		}
		n.value = value
		n.native = J_BOOLEAN
		break
	case "number":
		value, exists := element.Data.(system.NativeNumber).NativeNumber()
		if !exists {
			n.value = nil
			n.native = J_NULL
			break
		}
		n.value = value
		n.native = J_NUMBER
		break
	case "string":
		value, exists := element.Data.(system.NativeString).NativeString()
		if !exists {
			n.value = nil
			n.native = J_NULL
			break
		}
		n.value = value
		n.native = J_STRING
		break
	case "array":
		n.value = element.Data
		n.native = J_ARRAY
		length := element.Value.Len()
		itemsRule, err := element.Rule.ItemsRule()
		if err != nil {
			return nil, kerr.New("PXGPNCVEFH", err, "jdoc.Rule.ItemsRule (array)")
		}
		for i := 0; i < length; i++ {
			object, _, value, found, _, err := system.GetArrayMember(element.Value, i)
			if err != nil {
				return nil, kerr.New("UNSRRWKJTM", err, "system.GetArrayMember")
			}
			if !found {
				continue
			}
			child := &Element{Data: object, Rule: itemsRule, Value: value}
			// TODO: Why is the index i+1 here? Take time to understand this!
			nodes, err = p.getNodes(child, nodes, &n, "", i+1, length)
			if err != nil {
				return nil, kerr.New("LFQQBAJXJU", err, "getNodes (array)")
			}
		}
		break
	case "map":
		n.value = element.Data
		n.native = J_MAP
		itemsRule, err := element.Rule.ItemsRule()
		if err != nil {
			return nil, kerr.New("SYGEFDHBTO", err, "jdoc.Rule.ItemsRule (map)")
		}
		for _, k := range element.Value.MapKeys() {
			key, ok := k.Interface().(string)
			if !ok {
				return nil, kerr.New("QYPJSPHNRN", nil, "Map nodes must be strings, not %T", k.Interface())
			}
			object, _, value, found, _, err := system.GetMapMember(element.Value, key)
			if err != nil {
				return nil, kerr.New("JMUJMBBLWU", err, "system.GetMapMember")
			}
			if !found {
				continue
			}
			child := &Element{Data: object, Rule: itemsRule, Value: value}
			nodes, err = p.getNodes(child, nodes, &n, key, -1, -1)
			if err != nil {
				return nil, kerr.New("HEMMRWSOEU", err, "getNodes (map)")
			}
		}
		break
	case "object":
		n.value = element.Data
		n.native = J_OBJ
		for key, field := range n.ktype.Fields {
			object, _, value, found, _, err := system.GetObjectField(element.Value, system.GoName(key))
			if err != nil {
				return nil, kerr.New("JMUJMBBLWU", err, "system.GetMapMember")
			}
			if !found {
				continue
			}
			itemRule, err := system.NewRuleHolder(field)
			if err != nil {
				return nil, kerr.New("WWQNGYIUKN", err, "system.NewRuleHolder")
			}
			child := &Element{Data: object, Rule: itemRule, Value: value}
			nodes, err = p.getNodes(child, nodes, &n, key, -1, -1)
			if err != nil {
				return nil, kerr.New("NQWTHARJES", err, "p.getNodes")
			}
		}
		break
	}
	nodes = append(nodes, &n)
	return nodes, nil
}
