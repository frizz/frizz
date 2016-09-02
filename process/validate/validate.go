package validate // import "kego.io/process/validate"

// ke: {"package": {"complete": true}}

import (
	"context"

	"github.com/davelondon/kerr"
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
	errors, err = ValidateNode(ctx, n)
	if err != nil {
		return nil, kerr.Wrap("RVKNMWKQHD", err)
	}
	return errors, nil
}

func BuildRulesNode(ctx context.Context, n *node.Node, cache map[*node.Node][]system.RuleInterface) error {

	if n.Value == nil || n.Null || n.Missing {
		return nil
	}

	// Start with the rules from the type
	rules := n.Type.Rules

	if system.RulesApplyToObjects(n.Value) {
		ob := n.Value.(system.ObjectInterface).GetObject(nil)
		rules = append(rules, ob.Rules...)
	}

	if err := buildRulesObject(ctx, n, cache, rules); err != nil {
		return kerr.Wrap("KCXXUSBBEH", err)
	}
	return nil
}

func buildRulesObject(ctx context.Context, n *node.Node, cache map[*node.Node][]system.RuleInterface, rules []system.RuleInterface) error {

	if n.Value == nil || n.Null || n.Missing {
		return nil
	}

	if n.Rule != nil && n.Rule.Interface != nil {
		cache[n] = append(cache[n], n.Rule.Interface)
	}

	if rules != nil && len(rules) > 0 {
		p, err := selectors.CreateParser(ctx, n)
		if err != nil {
			return kerr.Wrap("AIWLGYGGAY", err)
		}
		for _, rule := range rules {
			base := rule.GetRule(nil)
			selector := ":root"
			if base.Selector != "" {
				selector = base.Selector
			}
			matches, err := p.GetNodes(selector)
			if err != nil {
				return kerr.Wrap("UKOCCFJWAB", err)
			}
			for _, match := range matches {
				cache[match] = append(cache[match], rule)
			}
		}
	}

	// Validate the children
	switch n.Type.NativeJsonType() {
	case json.J_OBJECT:
		if err := buildRulesObjectChildren(ctx, n, cache); err != nil {
			return kerr.Wrap("WLGMAVXNQN", err)
		}
	case json.J_ARRAY, json.J_MAP:
		if err := buildRulesCollectionChildren(ctx, n, cache); err != nil {
			return kerr.Wrap("WIDHJGYYMT", err)
		}
	}

	return nil
}

func buildRulesObjectChildren(ctx context.Context, n *node.Node, cache map[*node.Node][]system.RuleInterface) error {

	rules := []system.RuleInterface{}

	if system.RulesApplyToObjects(n.Value) {
		ob := n.Value.(system.ObjectInterface).GetObject(nil)
		rules = ob.Rules
	}

	for name, fieldRule := range n.Type.Fields {
		child := n.Map[name]
		fieldRules := rules
		fieldOb := fieldRule.(system.ObjectInterface).GetObject(nil)
		fieldRules = append(fieldRules, fieldOb.Rules...)
		if err := buildRulesObject(ctx, child, cache, fieldRules); err != nil {
			return kerr.Wrap("QXUGWLDDGN", err)
		}
	}

	return nil
}

func buildRulesCollectionChildren(ctx context.Context, n *node.Node, cache map[*node.Node][]system.RuleInterface) error {

	rules := n.Rule.Interface.(system.ObjectInterface).GetObject(nil).Rules

	// If we have additional rules on the main items rule, we should add them to rules
	items, err := n.Rule.ItemsRule()
	if err != nil {
		return kerr.Wrap("YFNERJIKWF", err)
	}
	itemsRuleOb := items.Interface.(system.ObjectInterface).GetObject(nil)
	if len(itemsRuleOb.Rules) > 0 {
		rules = append(rules, itemsRuleOb.Rules...)
	}

	for _, child := range n.Array {
		err := buildRulesObject(ctx, child, cache, rules)
		if err != nil {
			return kerr.Wrap("FWKXQRMPQJ", err)
		}
	}
	for _, child := range n.Map {
		err := buildRulesObject(ctx, child, cache, rules)
		if err != nil {
			return kerr.Wrap("MKYHIDMORV", err)
		}
	}
	return nil
}

func ValidateNode(ctx context.Context, n *node.Node) (errors []ValidationError, err error) {

	// First validate all the nodes
	for _, current := range n.Flatten(true) {
		// Validate the actual object
		if v, ok := current.Value.(system.Validator); ok {
			failed, messages, err := v.Validate(ctx)
			if err != nil {
				return nil, kerr.Wrap("JXKTDTIIYG", err)
			}
			if failed {
				for _, message := range messages {
					errors = append(errors, ValidationError{Struct: kerr.New("KULDIJUYFB", message), Source: current})
				}
			}
		}
	}

	// Then build a list of all nodes that have rules
	cache := map[*node.Node][]system.RuleInterface{}
	if err := BuildRulesNode(ctx, n, cache); err != nil {
		return nil, kerr.Wrap("YPUHTXPGRA", err)
	}

	// Then enforce the rules
	for current, rules := range cache {
		for _, rule := range rules {
			e, ok := rule.(system.Enforcer)
			if !ok {
				continue
			}
			failed, messages, err := e.Enforce(ctx, current.Value)
			if err != nil {
				return nil, kerr.Wrap("EBEMISLGDX", err)
			}
			if failed {
				for _, message := range messages {
					errors = append(errors, ValidationError{Struct: kerr.New("HLKQWDCMRN", message), Source: current})
				}
			}
		}
	}
	return errors, nil
}

type ValidationError struct {
	kerr.Struct
	Source *node.Node
}

type ValidationCommandError struct {
	kerr.Struct
}
