package node // import "kego.io/system/node"

import (
	"kego.io/json"
	"kego.io/kerr"
	"kego.io/system"
)

type Node struct {
	Parent      *Node
	Key         string            // in an object or a map, this is the key
	Index       int               // in an array, this is the index
	Origin      *system.Reference // in an object, this is the type that the field originated from - e.g. perhaps an embedded type
	ValueString string
	ValueNumber float64
	ValueBool   bool
	Value       interface{} // unmarshalled value
	Null        bool        // null is true if the json is null or the field is missing
	Missing     bool        // missing is only true if the field is missing
	Array       []*Node
	Map         map[string]*Node
	Rule        *system.RuleWrapper
	Type        *system.Type
	JsonType    json.Type
}

// Unpack unpacks a node from an unpackable
func (n *Node) Unpack(in json.Packed, path string, aliases map[string]string) error {
	if err := n.extract(nil, "", -1, &system.Reference{}, 0, in, true, nil, path, aliases); err != nil {
		return kerr.New("FUYLKYTQYD", err, "extract")
	}
	return nil
}

var _ json.ContextUnpacker = (*Node)(nil)

func (n *Node) extract(parent *Node, key string, index int, origin *system.Reference, siblings int, in json.Packed, exists bool, rule *system.RuleWrapper, path string, aliases map[string]string) error {

	objectType, err := extractType(in, rule, path, aliases)
	if err != nil {
		return kerr.New("RBDBRRUVMM", err, "extractType")
	}

	n.Parent = parent
	n.Key = key
	n.Index = index
	n.Rule = rule
	if n.Rule == nil {
		n.Rule = system.WrapEmptyRule(objectType)
	}
	n.Type = objectType
	n.Missing = !exists
	n.Origin = origin
	if in == nil {
		n.JsonType = json.J_NULL
		n.Null = true
	} else {
		n.JsonType = in.Type()
		n.Null = n.JsonType == json.J_NULL
	}

	// for objects and maps, UpType() from the unpackable is always J_MAP,
	// so we correct it for object types here.
	if n.JsonType == json.J_MAP && n.Type.NativeJsonType() == json.J_OBJECT {
		n.JsonType = json.J_OBJECT
	}

	// validate json type
	if n.JsonType != json.J_NULL && n.JsonType != n.Type.NativeJsonType() {
		return kerr.New("VEPLUIJXSN", nil, "json type is %s but object type is %s", n.JsonType, n.Type.NativeJsonType())
	}

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

	if n.Missing {
		// If the field doesn't exist in the input json, we should add the field info, but
		// we shouldn't try to add info for children. This would result in infinite recursion
		// where one of the fields is the same type as the parent - e.g. Type.Rule.
		return nil
	}

	if n.Null && objectType.Native.Value() != "object" {
		// For null input, we should skip processing the value or children, unless
		// the native type is object. For object types we should still process the
		// child fields because information from the parent type is added.
		return nil
	}

	switch objectType.Native.Value() {
	case "string":
		if in.Type() != json.J_STRING {
			return kerr.New("RKKSUYTCIA", nil, "Type %s should be a string", in.Type())
		}
		n.ValueString = in.String()
	case "number":
		if in.Type() != json.J_NUMBER {
			return kerr.New("RNSWFUUTHB", nil, "Type %s should be a float64", in.Type())
		}
		n.ValueNumber = in.Number()
	case "bool":
		if in.Type() != json.J_BOOL {
			return kerr.New("QGKJRAQUQI", nil, "Type %s should be a bool", in.Type())
		}
		n.ValueBool = in.Bool()
	case "array":
		if in.Type() != json.J_ARRAY {
			return kerr.New("CTJQUOKRTK", nil, "Type %s should be a []interface{}", in.Type())
		}
		c, ok := n.Rule.Interface.(system.CollectionRule)
		if !ok {
			return kerr.New("IUTONSPQOL", nil, "Rule %t must implement *CollectionRule for array types", n.Rule.Interface)
		}
		childRule, err := system.WrapRule(c.GetItemsRule())
		if err != nil {
			return kerr.New("KPIBIOCTGF", err, "NewRuleHolder (array)")
		}
		children := in.Array()
		for i, child := range children {
			childNode := &Node{}
			if err := childNode.extract(n, "", i, &system.Reference{}, len(children), child, true, childRule, path, aliases); err != nil {
				return kerr.New("VWWYPDIJKP", err, "get (array #%d)", i)
			}
			n.Array = append(n.Array, childNode)
		}
	case "map":
		if in.Type() != json.J_MAP {
			return kerr.New("IPWEPTWVYY", nil, "Type %s should be a map[string]interface{}", in.Type())
		}
		c, ok := n.Rule.Interface.(system.CollectionRule)
		if !ok {
			return kerr.New("RTQUNQEKUY", nil, "Rule %t must implement *CollectionRule for map types", n.Rule.Interface)
		}
		childRule, err := system.WrapRule(c.GetItemsRule())
		if err != nil {
			return kerr.New("SBFTRGJNAO", err, "NewRuleHolder (map)")
		}
		n.Map = map[string]*Node{}
		children := in.Map()
		for name, child := range children {
			childNode := &Node{}
			if err := childNode.extract(n, name, -1, &system.Reference{}, 0, child, true, childRule, path, aliases); err != nil {
				return kerr.New("HTOPDOKPRE", err, "get (map '%s')", name)
			}
			n.Map[name] = childNode
		}
	case "object":
		m := map[string]json.Packed{}
		if in != nil && in.Type() != json.J_NULL {
			if in.Type() != json.J_MAP {
				return kerr.New("CVCRNWMDYF", nil, "Type %s should be a map[string]interface{}", in.Type())
			}
			m = in.Map()
		}
		n.Map = map[string]*Node{}

		fields := map[string]*system.Field{}
		if err := extractFields(fields, n.Type); err != nil {
			return kerr.New("LPWTOSATQE", err, "extractFields (%s)", n.Type.Id.Value())
		}

		for name, f := range fields {
			rule, err := system.WrapRule(f.Rule)
			if err != nil {
				return kerr.New("YWFSOLOBXH", err, "NewRuleHolder (field '%s')", name)
			}
			child, ok := m[name]
			childNode := &Node{}
			if err := childNode.extract(n, name, -1, f.Origin, 0, child, ok, rule, path, aliases); err != nil {
				return kerr.New("LJUGPMWNPD", err, "get (field '%s')", name)
			}
			n.Map[name] = childNode
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

func extractType(in json.Packed, rule *system.RuleWrapper, path string, aliases map[string]string) (*system.Type, error) {

	if rule != nil && rule.Parent.Interface && rule.Struct.Interface {
		return nil, kerr.New("TDXTPGVFAK", nil, "Can't have interface type and rule at the same time")
	}

	if rule != nil && !rule.Parent.Interface && !rule.Struct.Interface {
		// If we have a rule, and it's not an interface, then we just return the
		// parent type of the rule.
		return rule.Parent, nil
	}

	// if the rule is nil (e.g. unpacking into an unknown type) or the type is an interface, we
	// ensure the input is a map
	if rule == nil || rule.Parent.Interface {
		if in == nil {
			return nil, nil
		}
		switch in.Type() {
		case json.J_MAP:
			break
		default:
			return nil, kerr.New("DLSQRFLINL", nil, "Input %s should be J_MAP if rule is nil or an interface type", in.Type())
		}

	}

	// if the rule is an interface rule, we ensure the input is a map or a native value
	if rule != nil && rule.Struct.Interface {
		if in == nil {
			return nil, nil
		}
		switch in.Type() {
		case json.J_MAP:
			break
		case json.J_STRING, json.J_NUMBER, json.J_BOOL:
			// if the input value is a native value, we will be unpacking into the parent
			// type of the rule
			return rule.Parent, nil
		default:
			return nil, kerr.New("SNYLGBJYTM", nil, "Input %s should be J_MAP, J_STRING, J_NUMBER or J_BOOL if rule is interface rule", in.Type())
		}
	}

	ob := in.Map()
	typeField, ok := ob["type"]
	if !ok {
		return nil, kerr.New("HBJVDKAKBJ", nil, "Input must have type field if rule is nil or an interface type")
	}
	var r system.Reference
	if err := r.Unpack(typeField, path, aliases); err != nil {
		return nil, kerr.New("YXHGIBXCOC", err, "Unpack (type)")
	}
	t, ok := r.GetType()
	if !ok {
		return nil, kerr.New("IJFMJJWVCA", nil, "Could not find type %s", r.Value())
	}

	return t, nil
}

func extractFields(fields map[string]*system.Field, t *system.Type) error {

	getType := func(r *system.Reference) (*system.Type, error) {
		g, ok := system.GetGlobal(r.Package, r.Name)
		if !ok {
			return nil, kerr.New("GJPKXQBKYH", nil, "Can't find global %s", r.Value())
		}
		t, ok := g.Object.(*system.Type)
		if !ok {
			return nil, kerr.New("BKYQWKFTIA", nil, "Global %T should *Type", g)
		}
		return t, nil
	}
	if !t.Basic && !t.Interface {
		// All types apart from Basic types embed system:object
		ob, err := getType(system.NewReference("kego.io/system", "object"))
		if err != nil {
			return kerr.New("YRFWOTIGFT", err, "getType (system:object)")
		}
		if err := extractFields(fields, ob); err != nil {
			return kerr.New("DTQEFALIMM", err, "extractFields (system:object)")
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
		fields[name] = &system.Field{Name: name, Rule: rule, Origin: t.Id}
	}
	return nil
}
