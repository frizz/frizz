package selectors

import (
	"kego.io/json"
	"kego.io/kerr"
	"kego.io/system"
)

type node struct {
	Value    interface{}
	JsonType json.Type
	Type     *system.Type
	element  *Element
	Parent   *node
	Key      string
	Index    int
	Siblings int
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

func (p *Parser) getNodes(element *Element, nodes []*node, parent *node, key string, index int, siblings int) ([]*node, error) {
	n := node{}
	n.Parent = parent
	n.element = element
	if len(key) > 0 {
		n.Key = key
	}
	if index > -1 {
		n.Index = index
	}
	if siblings > -1 {
		n.Siblings = siblings
	}

	oi, ok := element.Data.(system.ObjectInterface)
	if ok {
		o := oi.GetObject()
		n.Type, ok = o.Type.GetType()
		if !ok {
			return nil, kerr.New("HGENTDRWHL", nil, "Type.GetType not found")
		}
	} else {
		n.Type = element.Rule.ParentType
	}

	switch n.Type.Native.Value {
	case "bool":
		value, exists := element.Data.(system.NativeBool).NativeBool()
		if !exists {
			n.Value = nil
			n.JsonType = json.J_NULL
			break
		}
		n.Value = value
		n.JsonType = json.J_BOOL
		break
	case "number":
		value, exists := element.Data.(system.NativeNumber).NativeNumber()
		if !exists {
			n.Value = nil
			n.JsonType = json.J_NULL
			break
		}
		n.Value = value
		n.JsonType = json.J_NUMBER
		break
	case "string":
		value, exists := element.Data.(system.NativeString).NativeString()
		if !exists {
			n.Value = nil
			n.JsonType = json.J_NULL
			break
		}
		n.Value = value
		n.JsonType = json.J_STRING
		break
	case "array":
		n.Value = element.Data
		n.JsonType = json.J_ARRAY
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
		n.Value = element.Data
		n.JsonType = json.J_MAP
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
		n.Value = element.Data
		n.JsonType = json.J_OBJECT
		for key, field := range n.Type.Fields {
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
