package validate // import "kego.io/process/validate"

// ke: {"package": {"complete": true}}

import (
	"github.com/davelondon/kerr"
	"golang.org/x/net/context"
	"kego.io/context/envctx"
	"kego.io/json"
	"kego.io/process/scanner"
	"kego.io/process/validate/selectors"
	"kego.io/system"
	"kego.io/system/node"
)

func ValidatePackage(ctx context.Context) (errors []ValidationError, err error) {

	env := envctx.FromContext(ctx)

	files := scanner.ScanDirToFiles(ctx, env.Dir, env.Recursive)
	bytes := scanner.ScanFilesToBytes(ctx, files)
	for c := range bytes {
		if c.Err != nil {
			return nil, kerr.Wrap("IHSVWAUAYW", c.Err)
		}
		ve, err := validateBytes(ctx, c.Bytes)
		if err != nil {
			return nil, kerr.Wrap("KWLWXKWHLF", err)
		}
		if len(ve) > 0 {
			errors = append(errors, ve...)
		}
	}

	return
}

func validateBytes(ctx context.Context, bytes []byte) (errors []ValidationError, err error) {
	n, err := node.Unmarshal(ctx, bytes)
	if err != nil {
		return nil, kerr.Wrap("QIVNOQKCQF", err)
	}
	errors, err = ValidateNode(ctx, n, true)
	if err != nil {
		return nil, kerr.Wrap("RVKNMWKQHD", err)
	}
	return errors, nil
}

func ValidateNode(ctx context.Context, node *node.Node, children bool) (errors []ValidationError, err error) {

	if node.Value == nil || node.Null || node.Missing {
		return
	}

	// Start with the rules from the type
	rules := node.Type.Rules

	// Add the rules from the object, but only if they apply
	if system.RulesApplyToObjects(node.Value) {
		rules = append(rules, node.Value.(system.ObjectInterface).GetObject(nil).Rules...)
	}

	errors, err = validateObject(ctx, node, rules, children)
	if err != nil {
		return nil, kerr.Wrap("KQBLWNRGTL", err)
	}

	return errors, nil
}

type ValidationCommandError struct {
	kerr.Struct
}

type ValidationError struct {
	kerr.Struct
	Source *node.Node
}

func validateObject(ctx context.Context, node *node.Node, rules []system.RuleInterface, children bool) (errors []ValidationError, err error) {

	if node.Value == nil || node.Null || node.Missing {
		return nil, nil
	}

	// Validate the actual object
	if v, ok := node.Value.(system.Validator); ok {
		failed, messages, err := v.Validate(ctx)
		if err != nil {
			return nil, kerr.Wrap("RUGJLUAFAN", err)
		}
		if failed {
			for _, message := range messages {
				errors = append(errors, ValidationError{Struct: kerr.New("KULDIJUYFB", message), Source: node})
			}
			// if the object isn't valid we don't want to continue
			return errors, nil
		}
	}

	if node.Rule.Interface != nil {
		if e, ok := node.Rule.Interface.(system.Enforcer); ok {
			failed, messages, err := e.Enforce(ctx, node.Value)
			if err != nil {
				return nil, kerr.Wrap("EBEMISLGDX", err)
			}
			if failed {
				for _, message := range messages {
					errors = append(errors, ValidationError{Struct: kerr.New("HLKQWDCMRN", message), Source: node})
				}
			}
		}
	}

	if rules != nil && len(rules) > 0 {

		p, err := selectors.CreateParser(ctx, node)
		if err != nil {
			return nil, kerr.Wrap("AIWLGYGGAY", err)
		}

		for _, rule := range rules {

			base := rule.GetRule(nil)

			selector := ":root"
			if base.Selector != "" {
				selector = base.Selector
			}

			e, ok := rule.(system.Enforcer)
			if !ok {
				return nil, kerr.New("ABVWHMMXGG", "rule %T does not implement system.Enforcer", rule)
			}
			matches, err := p.GetNodes(selector)
			if err != nil {
				return nil, kerr.Wrap("UKOCCFJWAB", err)
			}
			for _, match := range matches {
				failed, messages, err := e.Enforce(ctx, match.Value)
				if err != nil {
					return nil, kerr.Wrap("MGHHDYTXVV", err)
				}
				if failed {
					for _, message := range messages {
						errors = append(errors, ValidationError{Struct: kerr.New("HAOXUVTFEX", message), Source: node})
					}
				}
			}
		}
	}

	if !children {
		return
	}

	// Validate the children
	switch node.Type.NativeJsonType() {
	case json.J_OBJECT:
		ve, err := validateObjectChildren(ctx, node)
		if err != nil {
			return nil, kerr.Wrap("WLGMAVXNQN", err)
		}
		if len(ve) > 0 {
			errors = append(errors, ve...)
		}
	case json.J_ARRAY:
		items, err := node.Rule.ItemsRule()
		if err != nil {
			return nil, kerr.Wrap("YFNERJIKWF", err)
		}
		rules := node.Rule.Interface.(system.ObjectInterface).GetObject(nil).Rules
		ve, err := validateArrayChildren(ctx, node, items, rules)
		if err != nil {
			return nil, kerr.Wrap("SFBYGQOIPN", err)
		}
		if len(ve) > 0 {
			errors = append(errors, ve...)
		}
	case json.J_MAP:
		items, err := node.Rule.ItemsRule()
		if err != nil {
			return nil, kerr.Wrap("PRPQQJKIKF", err)
		}
		rules := node.Rule.Interface.(system.ObjectInterface).GetObject(nil).Rules
		ve, err := validateMapChildren(ctx, node, items, rules)
		if err != nil {
			return nil, kerr.Wrap("RFQVHTNHGQ", err)
		}
		if len(ve) > 0 {
			errors = append(errors, ve...)
		}
	}

	return
}

func validateObjectChildren(ctx context.Context, node *node.Node) (errors []ValidationError, err error) {

	if node.Value == nil || node.Null || node.Missing {
		return
	}

	rules := []system.RuleInterface{}
	if system.RulesApplyToObjects(node.Value) {
		rules = node.Value.(system.ObjectInterface).GetObject(nil).Rules
	}

	for name, field := range node.Type.Fields {
		child, ok := node.Map[name]
		if !field.GetRule(nil).Optional && !ok {
			return nil, kerr.New("ETODESNSET", "Field %s is missing and not optional", name)
		}
		ob := field.(system.ObjectInterface).GetObject(nil)
		rules = append(rules, ob.Rules...)

		ve, err := validateObject(ctx, child, rules, true)
		if err != nil {
			return nil, kerr.Wrap("YJYSAOQWSJ", err)
		}
		if len(ve) > 0 {
			errors = append(errors, ve...)
		}
	}
	return
}

func validateArrayChildren(ctx context.Context, node *node.Node, itemsRule *system.RuleWrapper, rules []system.RuleInterface) (errors []ValidationError, err error) {

	// if we have additional rules on the main items rule, we should add them to rules
	if len(itemsRule.Interface.(system.ObjectInterface).GetObject(nil).Rules) > 0 {
		rules = append(rules, itemsRule.Interface.(system.ObjectInterface).GetObject(nil).Rules...)
	}

	for _, child := range node.Array {
		ve, err := validateObject(ctx, child, rules, true)
		if err != nil {
			return nil, kerr.Wrap("DKVEPIWTPI", err)
		}
		if len(ve) > 0 {
			errors = append(errors, ve...)
		}
	}
	return
}

func validateMapChildren(ctx context.Context, node *node.Node, itemsRule *system.RuleWrapper, rules []system.RuleInterface) (errors []ValidationError, err error) {

	// if we have additional rules on the main items rule, we should add them to rules
	if len(itemsRule.Interface.(system.ObjectInterface).GetObject(nil).Rules) > 0 {
		rules = append(rules, itemsRule.Interface.(system.ObjectInterface).GetObject(nil).Rules...)
	}

	for _, child := range node.Map {
		ve, err := validateObject(ctx, child, rules, true)
		if err != nil {
			return nil, kerr.Wrap("YLONAMFUAG", err)
		}
		if len(ve) > 0 {
			errors = append(errors, ve...)
		}
	}
	return
}
