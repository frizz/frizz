package helper

import (
	"strconv"

	"kego.io/json"
	"kego.io/kerr"
	"kego.io/system"
	_ "kego.io/system/types"
)

type Node struct {
	ValueString string
	ValueNumber float64
	ValueBool   bool
	Value       interface{}
	Null        bool
	Missing     bool
	Array       []*Node
	Map         map[string]*Node
	Fields      map[string]*Node
	Rule        *system.RuleHolder
	Type        *system.Type
}

func Unpack(input []byte, path string, aliases map[string]string) (*Node, error) {
	var v interface{}
	if err := json.UnmarshalPlain(input, &v); err != nil {
		return nil, kerr.New("GLCYSUHDGD", err, "UnmarshalPlain")
	}
	n, err := get(v, true, nil, path, aliases)
	if err != nil {
		return nil, kerr.New("FUYLKYTQYD", err, "get (read)")
	}
	return n, nil
}

func get(value interface{}, exists bool, rule *system.RuleHolder, path string, aliases map[string]string) (*Node, error) {

	objectType, err := extractObjectType(value, rule, path, aliases)
	if err != nil {
		return nil, kerr.New("RBDBRRUVMM", err, "extractObjectType")
	}

	n := &Node{}
	n.Rule = rule
	n.Type = objectType
	n.Null = value == nil
	n.Missing = !exists

	if n.Null && objectType.Native.Value != "object" {
		// For null input, we should skip processing the value or children, unless
		// the native type is object. For object types we should still process the
		// child fields because information from the parent type is added.
		return n, nil
	}

	if n.Missing {
		// If the field doesn't exist in the input json, we should add the field info, but
		// we shouldn't try to add info for children. This would result in infinite recursion
		// where one of the fields is the same type as the parent - e.g. Type.Rule.
		return n, nil
	}

	switch objectType.Native.Value {
	case "string":
		s, ok := value.(string)
		if !ok {
			return nil, kerr.New("RKKSUYTCIA", nil, "Type %T should be a string", value)
		}
		n.ValueString = s
	case "number":
		f, ok := value.(float64)
		if !ok {
			return nil, kerr.New("RNSWFUUTHB", nil, "Type %T should be a float64", value)
		}
		n.ValueNumber = f
	case "bool":
		b, ok := value.(bool)
		if !ok {
			return nil, kerr.New("QGKJRAQUQI", nil, "Type %T should be a bool", value)
		}
		n.ValueBool = b
	case "array":
		a, ok := value.([]interface{})
		if !ok {
			return nil, kerr.New("CTJQUOKRTK", nil, "Type %T should be a []interface{}", value)
		}
		c, ok := n.Rule.Rule.(system.CollectionRule)
		if !ok {
			return nil, kerr.New("IUTONSPQOL", nil, "Rule %t must implement *system.CollectionRule for array types", n.Rule.Rule)
		}
		childRule, err := system.NewRuleHolder(c.GetItemsRule())
		if err != nil {
			return nil, kerr.New("KPIBIOCTGF", err, "NewRuleHolder (array)")
		}
		for i, child := range a {
			childNode, err := get(child, true, childRule, path, aliases)
			if err != nil {
				return nil, kerr.New("VWWYPDIJKP", err, "get (array #%d)", i)
			}
			n.Array = append(n.Array, childNode)
		}
	case "map":
		m, ok := value.(map[string]interface{})
		if !ok {
			return nil, kerr.New("IPWEPTWVYY", nil, "Type %T should be a map[string]interface{}", value)
		}
		c, ok := n.Rule.Rule.(system.CollectionRule)
		if !ok {
			return nil, kerr.New("RTQUNQEKUY", nil, "Rule %t must implement *system.CollectionRule for map types", n.Rule.Rule)
		}
		childRule, err := system.NewRuleHolder(c.GetItemsRule())
		if err != nil {
			return nil, kerr.New("SBFTRGJNAO", err, "NewRuleHolder (map)")
		}
		n.Map = map[string]*Node{}
		for name, child := range m {
			childNode, err := get(child, true, childRule, path, aliases)
			if err != nil {
				return nil, kerr.New("HTOPDOKPRE", err, "get (map '%s')", name)
			}
			n.Map[name] = childNode
		}
	case "object":
		o := map[string]interface{}{}
		if value != nil {
			var ok bool
			o, ok = value.(map[string]interface{})
			if !ok {
				return nil, kerr.New("CVCRNWMDYF", nil, "Type %T should be a map[string]interface{}", value)
			}
		}
		n.Fields = map[string]*Node{}

		fields := map[string]system.Rule{}
		if err := extractFields(fields, n.Type); err != nil {
			return nil, kerr.New("LPWTOSATQE", err, "extractFields (%s)", n.Type.Id.Value())
		}

		for name, r := range fields {
			rule, err := system.NewRuleHolder(r)
			if err != nil {
				return nil, kerr.New("YWFSOLOBXH", err, "NewRuleHolder (field '%s')", name)
			}
			child, ok := o[name]
			childNode, err := get(child, ok, rule, path, aliases)
			if err != nil {
				return nil, kerr.New("LJUGPMWNPD", err, "get (field '%s')", name)
			}
			n.Fields[name] = childNode
		}
		for name, _ := range o {
			_, ok := fields[name]
			if !ok {
				return nil, kerr.New("SRANLETJRS", nil, "Extra field %s", name)
			}
		}
	}
	return n, nil
}

func extractObjectType(value interface{}, rule *system.RuleHolder, path string, aliases map[string]string) (*system.Type, error) {

	if rule != nil && !rule.ParentType.Interface {
		// If we have a rule, and it's not an interface, then we just return the
		// parent type of the rule.
		return rule.ParentType, nil
	}

	ob, ok := value.(map[string]interface{})
	if !ok {
		return nil, kerr.New("DLSQRFLINL", nil, "Input %T should be an object if rule is nil or an interface type", value)
	}
	typeField, ok := ob["type"]
	if !ok {
		return nil, kerr.New("HBJVDKAKBJ", nil, "Input must have type field if rule is nil or an interface type")
	}
	typeString, ok := typeField.(string)
	if !ok {
		return nil, kerr.New("RQGIJPCKFD", nil, "Type field %T must be a string", typeField)
	}
	var r system.Reference
	if err := r.UnmarshalJSON([]byte(strconv.Quote(typeString)), path, aliases); err != nil {
		return nil, kerr.New("YXHGIBXCOC", err, "UnmarshalInterface (type)")
	}
	t, ok := r.GetType()
	if !ok {
		return nil, kerr.New("IJFMJJWVCA", nil, "Could not find type %s", r.Value())
	}
	return t, nil

}

func extractFields(fields map[string]system.Rule, t *system.Type) error {

	getType := func(r system.Reference) (*system.Type, error) {
		g, ok := system.GetGlobal(r.Package, r.Name)
		if !ok {
			return nil, kerr.New("GJPKXQBKYH", nil, "Can't find global %s", r.Value())
		}
		t, ok := g.Object.(*system.Type)
		if !ok {
			return nil, kerr.New("BKYQWKFTIA", nil, "Global %T should *system.Type", g)
		}
		return t, nil
	}
	if !t.Basic && !t.Interface {
		// All types apart from Basic types embed system:$object (system.Object_base)
		ob, err := getType(system.NewReference("kego.io/system", "$object"))
		if err != nil {
			return kerr.New("YRFWOTIGFT", err, "getType (system:$object)")
		}
		if err := extractFields(fields, ob); err != nil {
			return kerr.New("DTQEFALIMM", err, "extractFields (system:$object)")
		}
	}
	for _, ifaceRef := range t.Is {
		// All interfaces that have a base type are embedded
		iface, err := getType(ifaceRef)
		if err != nil {
			return kerr.New("HQORJWWNJH", err, "getType (iface %s)", ifaceRef.Value())
		}
		if iface.Base == nil {
			continue
		}
		if err := extractFields(fields, iface.Base); err != nil {
			return kerr.New("WWILLDXRUI", err, "extractFields (iface %s)", ifaceRef.Value())
		}
	}
	for _, embedRef := range t.Embed {
		embed, err := getType(embedRef)
		if err != nil {
			return kerr.New("SLIRILCARQ", err, "getType (embed %s)", embedRef.Value())
		}
		if err := extractFields(fields, embed); err != nil {
			return kerr.New("JWAPCVIYBJ", err, "extractFields (embed %s)", embedRef.Value())
		}
	}
	for name, rule := range t.Fields {
		if _, ok := fields[name]; ok {
			return kerr.New("BARXPFXQNB", nil, "Duplicate field %s", name)
		}
		fields[name] = rule
	}
	return nil
}
