package validate // import "kego.io/process/validate"

import (
	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/context/envctx"
	"kego.io/json"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/process/scanutils"
	"kego.io/selectors"
	"kego.io/system"
	"kego.io/system/node"
)

func ValidatePackage(ctx context.Context) error {

	env := envctx.FromContext(ctx)
	cmd := cmdctx.FromContext(ctx)

	files := scanutils.ScanDirToFiles(ctx, cmd.Dir, env.Recursive)
	bytes := scanutils.ScanFilesToBytes(ctx, files)
	for c := range bytes {
		if c.Err != nil {
			return kerr.New("IHSVWAUAYW", c.Err, "ScanFilesToBytes")
		}
		if err := validateBytes(ctx, c.Bytes); err != nil {
			return kerr.New("KWLWXKWHLF", err, "validateBytes (%s)", c.File)
		}
	}

	return nil
}

func validateBytes(ctx context.Context, bytes []byte) error {
	n := node.NewNode()
	err := ke.UnmarshalUntyped(ctx, bytes, n)
	if up, ok := err.(json.UnknownPackageError); ok {
		return kerr.New("QPOGRNXWMH", err, "unknown package %s", up.UnknownPackage)
	} else if ut, ok := err.(json.UnknownTypeError); ok {
		return kerr.New("PJABFRVFLF", err, "unknown type %s", ut.UnknownType)
	} else if err != nil {
		return kerr.New("QIVNOQKCQF", err, "UnmarshalNode")
	}

	if err := validateNode(ctx, n); err != nil {
		return kerr.New("RVKNMWKQHD", err, "validateUnknown")
	}
	return nil
}

func validateNode(ctx context.Context, node *node.Node) error {

	if node.Value == nil || node.Null || node.Missing {
		return nil
	}

	// Start with the rules from the type
	rules := node.Type.Rules

	// Add the rules from the object, but only if they apply
	if system.RulesApplyToObjects(node.Value) {
		rules = append(rules, node.Value.(system.ObjectInterface).GetObject(nil).Rules...)
	}

	return validateObject(ctx, node, rules)

}

type ValidationError struct {
	kerr.Struct
}

func validateObject(ctx context.Context, node *node.Node, rules []system.RuleInterface) error {

	if node.Value == nil || node.Null || node.Missing {
		return nil
	}

	// Validate the actual object
	if v, ok := node.Value.(system.Validator); ok {
		ok, message, err := v.Validate(ctx)
		if err != nil {
			return kerr.New("RUGJLUAFAN", err, "v.Validate")
		}
		if !ok {
			return ValidationError{kerr.New("KULDIJUYFB", nil, message)}
		}
	}

	if node.Rule.Interface != nil {
		e, ok := node.Rule.Interface.(system.Enforcer)
		if ok {
			ok, message, err := e.Enforce(ctx, node.Value)
			if err != nil {
				return kerr.New("EBEMISLGDX", err, "e.Enforce (main)")
			}
			if !ok {
				return ValidationError{kerr.New("HLKQWDCMRN", nil, message)}
			}
		}
	}

	if rules != nil && len(rules) > 0 {

		p, err := selectors.CreateParser(ctx, node)
		if err != nil {
			return kerr.New("AIWLGYGGAY", err, "selectors.CreateParser")
		}

		for _, rule := range rules {

			base := rule.GetRule(nil)

			selector := ":root"
			if base.Selector != "" {
				selector = base.Selector
			}

			e, ok := rule.(system.Enforcer)
			if !ok {
				return kerr.New("ABVWHMMXGG", nil, "rule %T does not implement system.Enforcer", rule)
			}
			matches, err := p.GetNodes(selector)
			if err != nil {
				return kerr.New("UKOCCFJWAB", err, "p.GetJsonElements (%s)", selector)
			}
			for _, match := range matches {
				ok, message, err := e.Enforce(ctx, match.GetNode().Value)
				if err != nil {
					return kerr.New("MGHHDYTXVV", err, "e.Enforce")
				}
				if !ok {
					return ValidationError{kerr.New("HAOXUVTFEX", nil, message)}
				}
			}
		}
	}

	// Validate the children
	switch node.Type.NativeJsonType() {
	case json.J_OBJECT:
		return validateObjectChildren(ctx, node)
	case json.J_ARRAY:
		items, err := node.Rule.ItemsRule()
		if err != nil {
			return kerr.New("YFNERJIKWF", err, "rule.ItemsRule (array)")
		}
		rules := node.Rule.Interface.(system.ObjectInterface).GetObject(nil).Rules
		return validateArrayChildren(ctx, node, items, rules)
	case json.J_MAP:
		items, err := node.Rule.ItemsRule()
		if err != nil {
			return kerr.New("PRPQQJKIKF", err, "rule.ItemsRule (map)")
		}
		rules := node.Rule.Interface.(system.ObjectInterface).GetObject(nil).Rules
		return validateMapChildren(ctx, node, items, rules)
	}

	return nil
}

func validateObjectChildren(ctx context.Context, node *node.Node) error {

	if node.Value == nil || node.Null || node.Missing {
		return nil
	}

	rules := []system.RuleInterface{}
	if system.RulesApplyToObjects(node.Value) {
		rules = node.Value.(system.ObjectInterface).GetObject(nil).Rules
	}

	for name, field := range node.Type.Fields {
		child, ok := node.Map[name]
		if !field.GetRule(nil).Optional && !ok {
			return kerr.New("ETODESNSET", nil, "Field %s is missing and not optional", name)
		}
		ob, ok := field.(system.ObjectInterface)
		if !ok {
			return kerr.New("XRTVWVUAMP", nil, "field does not implement system.ObjectInterface")
		}

		allRules := append(rules, ob.GetObject(nil).Rules...)

		// if we have additional rules on the main field rule, we should add them to allRules
		if len(child.GetNode().Rule.Interface.(system.ObjectInterface).GetObject(nil).Rules) > 0 {
			allRules = append(allRules, child.GetNode().Rule.Interface.(system.ObjectInterface).GetObject(nil).Rules...)
		}

		if err := validateObject(ctx, child.GetNode(), allRules); err != nil {
			return kerr.New("YJYSAOQWSJ", err, "validateObject (%s)", name)
		}
	}
	return nil
}

func validateArrayChildren(ctx context.Context, node *node.Node, itemsRule *system.RuleWrapper, rules []system.RuleInterface) error {

	// if we have additional rules on the main items rule, we should add them to rules
	if len(itemsRule.Interface.(system.ObjectInterface).GetObject(nil).Rules) > 0 {
		rules = append(rules, itemsRule.Interface.(system.ObjectInterface).GetObject(nil).Rules...)
	}

	for i, child := range node.Array {
		if err := validateObject(ctx, child.GetNode(), rules); err != nil {
			return kerr.New("DKVEPIWTPI", err, "validateObject array index %s", i)
		}
	}
	return nil
}

func validateMapChildren(ctx context.Context, node *node.Node, itemsRule *system.RuleWrapper, rules []system.RuleInterface) error {

	// if we have additional rules on the main items rule, we should add them to rules
	if len(itemsRule.Interface.(system.ObjectInterface).GetObject(nil).Rules) > 0 {
		rules = append(rules, itemsRule.Interface.(system.ObjectInterface).GetObject(nil).Rules...)
	}

	for n, child := range node.Map {
		if err := validateObject(ctx, child.GetNode(), rules); err != nil {
			return kerr.New("YLONAMFUAG", err, "validateObject map key %s", n)
		}
	}
	return nil
}
