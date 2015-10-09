package system

import (
	"kego.io/json"
	"kego.io/kerr"
)

type Node struct {
	Parent      *Node
	Key         string    // in an object or a map, this is the key
	Index       int       // in an array, this is the index
	Origin      Reference // in an object, this is the type that the field originated from
	ValueString string
	ValueNumber float64
	ValueBool   bool
	Value       interface{}
	Null        bool
	Missing     bool
	Array       []*Node
	Map         map[string]*Node
	Fields      map[string]*Node
	Rule        *RuleHolder
	Type        *Type
	JsonType    int
}

func (n *Node) Unpack(in json.Unpackable, path string, aliases map[string]string) error {
	if err := n.extract(nil, "", -1, Reference{}, in, true, nil, path, aliases); err != nil {
		return kerr.New("FUYLKYTQYD", err, "get (read)")
	}
	return nil
}

var _ json.ContextUnpacker = (*Node)(nil)

func (n *Node) extract(parent *Node, key string, index int, origin Reference, in json.Unpackable, exists bool, rule *RuleHolder, path string, aliases map[string]string) error {

	objectType, err := extractType(in, rule, path, aliases)
	if err != nil {
		return kerr.New("RBDBRRUVMM", err, "extractObjectType")
	}

	n.Parent = parent
	n.Key = key
	n.Index = index
	n.Rule = rule
	n.Type = objectType
	n.Missing = !exists
	n.Origin = origin
	if in != nil {
		n.JsonType = in.UpType()
	} else {
		n.JsonType = json.J_NULL
	}
	n.Null = n.JsonType == json.J_NULL

	if rule == nil {
		if err := json.Unpack(in, &n.Value, path, aliases); err != nil {
			return kerr.New("CQMWGPLYIJ", err, "Unpack")
		}
	} else {
		t, err := rule.GetReflectType()
		if err != nil {
			return kerr.New("DQJDYPIANO", err, "GetReflectType")
		}
		if err := json.UnpackFragment(in, &n.Value, t, path, aliases); err != nil {
			return kerr.New("PEVKGFFHLL", err, "UnpackFragment")
		}
	}

	if n.Null && objectType.Native.Value != "object" {
		// For null input, we should skip processing the value or children, unless
		// the native type is object. For object types we should still process the
		// child fields because information from the parent type is added.
		return nil
	}

	if n.Missing {
		// If the field doesn't exist in the input json, we should add the field info, but
		// we shouldn't try to add info for children. This would result in infinite recursion
		// where one of the fields is the same type as the parent - e.g. Type.Rule.
		return nil
	}

	switch objectType.Native.Value {
	case "string":
		if in.UpType() != json.J_STRING {
			return kerr.New("RKKSUYTCIA", nil, "Type %s should be a string", in.UpType())
		}
		n.ValueString = in.UpString()
	case "number":
		if in.UpType() != json.J_NUMBER {
			return kerr.New("RNSWFUUTHB", nil, "Type %s should be a float64", in.UpType())
		}
		n.ValueNumber = in.UpNumber()
	case "bool":
		if in.UpType() != json.J_BOOL {
			return kerr.New("QGKJRAQUQI", nil, "Type %s should be a bool", in.UpType())
		}
		n.ValueBool = in.UpBool()
	case "array":
		if in.UpType() != json.J_ARRAY {
			return kerr.New("CTJQUOKRTK", nil, "Type %s should be a []interface{}", in.UpType())
		}
		c, ok := n.Rule.Rule.(CollectionRule)
		if !ok {
			return kerr.New("IUTONSPQOL", nil, "Rule %t must implement *CollectionRule for array types", n.Rule.Rule)
		}
		childRule, err := NewRuleHolder(c.GetItemsRule())
		if err != nil {
			return kerr.New("KPIBIOCTGF", err, "NewRuleHolder (array)")
		}
		for i, child := range in.UpArray() {
			childNode := &Node{}
			if err := childNode.extract(n, "", i, Reference{}, child, true, childRule, path, aliases); err != nil {
				return kerr.New("VWWYPDIJKP", err, "get (array #%d)", i)
			}
			n.Array = append(n.Array, childNode)
		}
	case "map":
		if in.UpType() != json.J_MAP {
			return kerr.New("IPWEPTWVYY", nil, "Type %s should be a map[string]interface{}", in.UpType())
		}
		c, ok := n.Rule.Rule.(CollectionRule)
		if !ok {
			return kerr.New("RTQUNQEKUY", nil, "Rule %t must implement *CollectionRule for map types", n.Rule.Rule)
		}
		childRule, err := NewRuleHolder(c.GetItemsRule())
		if err != nil {
			return kerr.New("SBFTRGJNAO", err, "NewRuleHolder (map)")
		}
		n.Map = map[string]*Node{}
		for name, child := range in.UpMap() {
			childNode := &Node{}
			if err := childNode.extract(n, name, -1, Reference{}, child, true, childRule, path, aliases); err != nil {
				return kerr.New("HTOPDOKPRE", err, "get (map '%s')", name)
			}
			n.Map[name] = childNode
		}
	case "object":
		m := map[string]json.Unpackable{}
		if in != nil && in.UpType() != json.J_NULL {
			if in.UpType() != json.J_MAP {
				return kerr.New("CVCRNWMDYF", nil, "Type %s should be a map[string]interface{}", in.UpType())
			}
			m = in.UpMap()
		}
		n.Fields = map[string]*Node{}

		fields := map[string]*Field{}
		if err := extractFields(fields, n.Type); err != nil {
			return kerr.New("LPWTOSATQE", err, "extractFields (%s)", n.Type.Id.Value())
		}

		for name, f := range fields {
			rule, err := NewRuleHolder(f.Rule)
			if err != nil {
				return kerr.New("YWFSOLOBXH", err, "NewRuleHolder (field '%s')", name)
			}
			child, ok := m[name]
			childNode := &Node{}
			if err := childNode.extract(n, name, -1, f.Origin, child, ok, rule, path, aliases); err != nil {
				return kerr.New("LJUGPMWNPD", err, "get (field '%s')", name)
			}
			n.Fields[name] = childNode
		}
		for name, _ := range m {
			_, ok := fields[name]
			if !ok {
				return kerr.New("SRANLETJRS", nil, "Extra field %s", name)
			}
		}
	}
	return nil
}

type Field struct {
	Rule   Rule
	Origin Reference
}

func extractType(in json.Unpackable, rule *RuleHolder, path string, aliases map[string]string) (*Type, error) {

	if rule != nil && !rule.ParentType.Interface {
		// If we have a rule, and it's not an interface, then we just return the
		// parent type of the rule.
		return rule.ParentType, nil
	}

	if in.UpType() != json.J_MAP {
		return nil, kerr.New("DLSQRFLINL", nil, "Input %s should be J_MAP if rule is nil or an interface type", in.UpType())
	}
	ob := in.UpMap()
	typeField, ok := ob["type"]
	if !ok {
		return nil, kerr.New("HBJVDKAKBJ", nil, "Input must have type field if rule is nil or an interface type")
	}
	var r Reference
	if err := r.Unpack(typeField, path, aliases); err != nil {
		return nil, kerr.New("YXHGIBXCOC", err, "Unpack (type)")
	}
	t, ok := r.GetType()
	if !ok {
		return nil, kerr.New("IJFMJJWVCA", nil, "Could not find type %s", r.Value())
	}
	return t, nil

}

func extractFields(fields map[string]*Field, t *Type) error {

	getType := func(r Reference) (*Type, error) {
		g, ok := GetGlobal(r.Package, r.Name)
		if !ok {
			return nil, kerr.New("GJPKXQBKYH", nil, "Can't find global %s", r.Value())
		}
		t, ok := g.Object.(*Type)
		if !ok {
			return nil, kerr.New("BKYQWKFTIA", nil, "Global %T should *Type", g)
		}
		return t, nil
	}
	if !t.Basic && !t.Interface {
		// All types apart from Basic types embed system:$object (Object_base)
		ob, err := getType(NewReference("kego.io/system", "$object"))
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
		fields[name] = &Field{rule, t.Id}
	}
	return nil
}
