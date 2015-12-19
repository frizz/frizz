package validate // import "kego.io/process/validate"

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/net/context"
	"kego.io/context/cmdctx"
	"kego.io/json"
	"kego.io/ke"
	"kego.io/kerr"
	"kego.io/process/scan"
	"kego.io/selectors"
	"kego.io/system"
	"kego.io/system/node"
)

func Validate(ctx context.Context) error {

	cmd := cmdctx.FromContext(ctx)

	walker := func(filePath string, file os.FileInfo, err error) error {
		if err != nil {
			return kerr.New("GFBBIERGIY", err, "walker (%s)", filePath)
		}
		if err := validateFile(ctx, filePath); err != nil {
			return kerr.New("QJGXAUKPTI", err, "validateFile (%s)", filePath)
		}
		return nil
	}

	if cmd.Recursive {
		if err := filepath.Walk(cmd.Dir, walker); err != nil {
			return kerr.New("GCKFJQJUXK", err, "filepath.Walk")
		}
	} else {
		files, err := ioutil.ReadDir(cmd.Dir)
		if err != nil {
			return kerr.New("RJXRHBYVUW", err, "ioutil.ReadDir")
		}
		for _, f := range files {
			if err := walker(filepath.Join(cmd.Dir, f.Name()), f, nil); err != nil {
				return kerr.New("UJBOWKFUMS", err, "walker")
			}
		}
	}

	return nil
}

func validateFile(ctx context.Context, filePath string) error {

	bytes, hash, err := scan.OpenFile(ctx, filePath)
	if err != nil {
		return kerr.New("XXYPVKLNBQ", err, "openFile")
	}
	if bytes == nil {
		return nil
	}

	if err = validateBytes(ctx, bytes, hash); err != nil {
		return kerr.New("GFVGDBDTNQ", err, "validateReader (%s)", filePath)
	}
	return nil
}

func validateBytes(ctx context.Context, bytes []byte, hash uint64) error {
	n := &node.Node{}
	err := ke.UnmarshalNode(ctx, bytes, n)
	if up, ok := err.(json.UnknownPackageError); ok {
		return kerr.New("QPOGRNXWMH", err, "unknown package %s", up.UnknownPackage)
	} else if ut, ok := err.(json.UnknownTypeError); ok {
		return kerr.New("PJABFRVFLF", nil, "unknown type %s", ut.UnknownType)
	} else if err != nil {
		return kerr.New("QIVNOQKCQF", err, "UnmarshalNode")
	}

	if err := validateNode(ctx, n, hash); err != nil {
		return kerr.New("RVKNMWKQHD", err, "validateUnknown")
	}
	return nil
}

func validateNode(ctx context.Context, node *node.Node, hash uint64) error {

	if node.Value == nil || node.Null || node.Missing {
		return nil
	}

	if *node.Type.Id == *system.NewReference("kego.io/system", "type") {
		t := node.Value.(*system.Type)
		h, ok := system.GetGlobal(t.Id.Package, t.Id.Name)
		if !ok {
			return TypesChangedError{kerr.New("XCBNOPEEUY", nil, "New type %s found - run kego again to rebuild", node.Type.Id.Value())}
		}
		if hash != h.Hash {
			return TypesChangedError{kerr.New("KRKBNITUON", nil, "Type %s has changed - run kego again to rebuild", node.Type.Id.Value())}
		}
	}

	// Start with the rules from the type
	rules := node.Type.Rules

	// Add the rules from the object, but only if they apply
	if system.RulesApplyToObjects(node.Value) {
		rules = append(rules, node.Value.(system.ObjectInterface).GetObject(nil).Rules...)
	}

	return validateObject(ctx, node, rules)

}

type TypesChangedError struct {
	kerr.Struct
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
				ok, message, err := e.Enforce(ctx, match.Value)
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
		if len(child.Rule.Interface.(system.ObjectInterface).GetObject(nil).Rules) > 0 {
			allRules = append(allRules, child.Rule.Interface.(system.ObjectInterface).GetObject(nil).Rules...)
		}

		if err := validateObject(ctx, child, allRules); err != nil {
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
		if err := validateObject(ctx, child, rules); err != nil {
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
		if err := validateObject(ctx, child, rules); err != nil {
			return kerr.New("YLONAMFUAG", err, "validateObject map key %s", n)
		}
	}
	return nil
}
