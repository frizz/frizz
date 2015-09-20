package client

import (
	"fmt"
	"reflect"

	"honnef.co/go/js/dom"

	"kego.io/editor/client/tree"
	"kego.io/kerr"
	"kego.io/system"
)

type element struct {
	name   string
	index  int
	data   interface{}
	value  reflect.Value
	rule   *system.RuleHolder
	exists bool
	ktype  *system.Type
	label  *dom.HTMLDivElement
}

func (e *element) Initialise(div dom.Element) {
	label := dom.GetWindow().Document().CreateElement("div").(*dom.HTMLDivElement)
	if e.index > -1 {
		label.SetTextContent(fmt.Sprint("[", e.index, "]"))
	} else {
		label.SetTextContent(e.name)
	}
	e.label = label
	div.AppendChild(label)
}

var _ = tree.Item(&element{})

// addElement adds a ke object element to the tree, and recursively
// adds children
func addElement(el *element, parent *tree.Node) error {
	n := tree.NewNode(el)
	parent.AppendNodes(n)
	n.Open()

	if el.data == nil {
		return nil
	}
	return addChildren(el, n)
}

func addChildren(el *element, n *tree.Node) error {

	ob, ok := el.data.(system.Object)
	if ok {
		base := ob.Object()
		el.ktype, ok = base.Type.GetType()
		if !ok {
			return kerr.New("EVTTBNAQQO", nil, "Type.GetType not found")
		}
	} else {
		el.ktype = el.rule.ParentType
	}

	switch el.ktype.Native.Value {
	case "bool":
		if el.ktype.Id.Value() == "kego.io/json:bool" {
			el.exists = true
			break
		}
		_, exists := el.data.(system.NativeBool).NativeBool()
		if !exists {
			el.exists = false
			break
		}
		el.exists = true
		break
	case "number":
		if el.ktype.Id.Value() == "kego.io/json:number" {
			el.exists = true
			break
		}
		_, exists := el.data.(system.NativeNumber).NativeNumber()
		if !exists {
			el.exists = false
			break
		}
		el.exists = true
		break
	case "string":
		if el.ktype.Id.Value() == "kego.io/json:string" {
			el.exists = true
			break
		}
		_, exists := el.data.(system.NativeString).NativeString()
		if !exists {
			el.exists = false
			break
		}
		el.exists = true
		break
	case "array":
		el.exists = true
		length := el.value.Len()
		itemsRule, err := el.rule.ItemsRule()
		if err != nil {
			return kerr.New("TURFPFGSIB", err, "jdoc.Rule.ItemsRule (array)")
		}
		for i := 0; i < length; i++ {
			object, _, value, found, _, err := system.GetArrayMember(el.value, i)
			if err != nil {
				return kerr.New("FEQEQEUTLA", err, "system.GetArrayMember")
			}
			if !found {
				return kerr.New("BUHOQJODKR", nil, "Array member at index %d not found", i)
			}
			child := &element{data: object, rule: itemsRule, value: value, index: i}
			if err := addElement(child, n); err != nil {
				return kerr.New("TKHPKVXKFH", err, "addNodes (array)")
			}
		}
		break
	case "map":
		el.exists = true
		itemsRule, err := el.rule.ItemsRule()
		if err != nil {
			return kerr.New("HQKDLCRHRU", err, "jdoc.Rule.ItemsRule (map)")
		}
		for _, k := range el.value.MapKeys() {
			key, ok := k.Interface().(string)
			if !ok {
				return kerr.New("HWTNSTWMSM", nil, "Map nodes must be strings, not %T", k.Interface())
			}
			object, _, value, found, _, err := system.GetMapMember(el.value, key)
			if err != nil {
				return kerr.New("HKLREGADQC", err, "system.GetMapMember")
			}
			if !found {
				return kerr.New("IQXHYKRBGW", nil, "Map member for key %s not found", key)
			}
			child := &element{data: object, rule: itemsRule, value: value, index: -1, name: key}
			if err := addElement(child, n); err != nil {
				return kerr.New("XWNWOLIPQV", err, "addNodes (map)")
			}
		}
		break
	case "object":
		el.exists = true
		for name, rule := range el.ktype.Fields {
			object, _, value, found, _, err := system.GetObjectField(el.value, system.GoName(name))
			if err != nil {
				return kerr.New("WQXDRFTQHW", err, "system.GetObjectField")
			}
			itemRule, err := system.NewRuleHolder(rule)
			if err != nil {
				return kerr.New("XBBAUGBNWY", err, "system.NewRuleHolder")
			}
			child := &element{data: object, rule: itemRule, value: value, index: -1, name: name, exists: found}
			if err := addElement(child, n); err != nil {
				return kerr.New("RXOYBWPLQG", err, "addNodes (object)")
			}
		}
		break
	}
	return nil
}
