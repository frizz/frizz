package process

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"kego.io/json"
	"kego.io/jsonselect"
	"kego.io/kerr"
	"kego.io/system"
)

func Validate(root string, packagePath string, imports map[string]string) error {

	walker := func(filePath string, file os.FileInfo, err error) error {
		if err != nil {
			return kerr.New("GFBBIERGIY", err, "process.Validate", "walker (%s)", filePath)
		}
		if err := validateFile(filePath, packagePath, imports); err != nil {
			return kerr.New("QJGXAUKPTI", err, "process.Validate", "validateFile (%s)", filePath)
		}
		return nil
	}

	if err := filepath.Walk(root, walker); err != nil {
		return kerr.New("GCKFJQJUXK", err, "process.Validate", "filepath.Walk")
	}
	return nil
}

func validateFile(filePath string, packagePath string, imports map[string]string) error {

	if !strings.HasSuffix(filePath, ".json") {
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return kerr.New("TIODJJGDCL", err, "process.validateFile", "os.Open (%s)", filePath)
	}
	defer file.Close()

	if err = validateReader(file, packagePath, imports); err != nil {
		return kerr.New("GFVGDBDTNQ", err, "process.validateFile", "validateReader (%s)", filePath)
	}
	return nil
}

func validateReader(file io.Reader, packagePath string, imports map[string]string) error {
	var i interface{}
	unknown, err := json.NewDecoder(file, packagePath, imports).Decode(&i)
	if err != nil {
		return kerr.New("QIVNOQKCQF", err, "process.validateReader", "json.NewDecoder.Decode")
	}
	if unknown {
		return kerr.New("PJABFRVFLF", nil, "process.validateReader", "json.NewDecoder: unknown types")
	}
	if err := validateUnknown(i, packagePath, imports); err != nil {
		return kerr.New("RVKNMWKQHD", err, "process.validateReader", "validateUnknown")
	}
	return nil
}

func validateUnknown(data interface{}, path string, imports map[string]string) error {

	tp, ok := data.(system.Typer)
	if !ok {
		return kerr.New("RSSFIFHCOF", nil, "process.validateUnknown", "data %T does not implement system.Typer", data)
	}
	t, ok := tp.GetType()
	if !ok {
		return kerr.New("BNQTCEDVGV", nil, "process.validateUnknown", "tp.GetType not found")
	}

	partialRuleHolder := system.NewMinimalRuleHolder(t, path, imports)

	rules := []system.Rule{}
	if data.(system.Ruler).RulesApply() == system.RULES_APPLY_TO_OBJECTS {
		rules = data.(system.Ruler).GetRules()
	}

	return validateObject(partialRuleHolder, rules, data, path, imports)

}

func validateObject(rule *system.RuleHolder, rules []system.Rule, data interface{}, path string, imports map[string]string) error {

	// Validate the actual object
	if v, ok := data.(system.Validator); ok {
		ok, message, err := v.Validate(path, imports)
		if err != nil {
			return kerr.New("RUGJLUAFAN", err, "process.validateObject", "v.Validate")
		}
		if !ok {
			return kerr.New("DCIARXKRXN", nil, "process.validateObject", "Invalid data. %s. %#v", message, data)
		}
	}

	if rule.Rule != nil {
		e, ok := rule.Rule.(system.Enforcer)
		if ok {
			ok, message, err := e.Enforce(data, path, imports)
			if err != nil {
				return kerr.New("EBEMISLGDX", err, "process.validateObject", "e.Enforce (main)")
			}
			if !ok {
				return kerr.New("KKOFBHILXM", nil, "process.validateObject", "Broken rule. %s. %#v", message, data)
			}
		}
	}

	if rules != nil && len(rules) > 0 {

		j := &jsonselect.Element{Data: data, Value: reflect.ValueOf(data), Rule: rule}
		p, err := jsonselect.CreateParser(j, path, imports)
		if err != nil {
			return kerr.New("AIWLGYGGAY", err, "process.validateObject", "jsonselect.CreateParser")
		}

		for _, rule := range rules {

			base := rule.GetRuleBase()

			selector := ":root"
			if base.Selector != "" {
				selector = base.Selector
			}

			e, ok := rule.(system.Enforcer)
			if !ok {
				return kerr.New("ABVWHMMXGG", nil, "process.validateObject", "rule %T does not implement system.Enforcer", rule)
			}
			matches, err := p.GetElements(selector)
			if err != nil {
				return kerr.New("UKOCCFJWAB", err, "process.validateObject", "p.GetJsonElements (%s)", selector)
			}
			for _, match := range matches {
				ok, message, err := e.Enforce(match.Data, path, imports)
				if err != nil {
					return kerr.New("MGHHDYTXVV", err, "process.validateObject", "e.Enforce")
				}
				if !ok {
					return kerr.New("FRXEXSTARP", nil, "process.validateObject", "Broken rule. %s. %#v", message, match.Data)
				}

			}
		}
	}

	// Validate the children
	switch rule.ParentType.Native.Value {
	case "object":
		return validateObjectChildren(rule, data, path, imports)
	case "array":
		items, err := rule.ItemsRule()
		if err != nil {
			return kerr.New("YFNERJIKWF", err, "process.validateObject", "rule.ItemsRule (array)")
		}
		rules := rule.Rule.(system.Ruler).GetRules()
		return validateArrayChildren(items, rules, data, path, imports)
	case "map":
		items, err := rule.ItemsRule()
		if err != nil {
			return kerr.New("PRPQQJKIKF", err, "process.validateObject", "rule.ItemsRule (map)")
		}
		rules := rule.Rule.(system.Ruler).GetRules()
		return validateMapChildren(items, rules, data, path, imports)
	}

	return nil
}

// TODO: Generate some code to remove the reflection from this
func validateObjectChildren(itemsRule *system.RuleHolder, data interface{}, path string, imports map[string]string) error {

	if data == nil {
		return nil
	}

	value := reflect.Indirect(reflect.ValueOf(data))
	if value.Kind() != reflect.Struct {
		return kerr.New("FJBEEDBLOK", nil, "process.validateObjectChildren", "value.Kind %s (%T) must be Struct", value.Kind(), data)
	}

	rules := []system.Rule{}
	if data.(system.Ruler).RulesApply() == system.RULES_APPLY_TO_OBJECTS {
		rules = data.(system.Ruler).GetRules()
	}

	for name, field := range itemsRule.ParentType.Fields {
		child, _, _, found, _, err := system.GetObjectField(value, system.IdToGoName(name))
		if err != nil {
			return kerr.New("XTUKWWRDHH", err, "process.validateObjectChildren", "system.GetField (%s)", name)
		}
		if !field.GetRuleBase().Optional && !found {
			return kerr.New("ETODESNSET", nil, "process.validateObjectChildren", "Field %s is missing and not optional", name)
		}
		childRule, err := system.NewRuleHolder(field, path, imports)
		if err != nil {
			return kerr.New("IQOXVXBLRO", err, "process.validateObjectChildren", "system.NewRuleHolder (%s)", name)
		}
		ob, ok := field.(system.Object)
		if !ok {
			return kerr.New("XRTVWVUAMP", nil, "process.validateObjectChildren", "field does not implement system.Object")
		}
		allRules := append(rules, ob.GetBase().Rules...)
		if err = validateObject(childRule, allRules, child, path, imports); err != nil {
			return kerr.New("YJYSAOQWSJ", err, "process.validateObjectChildren", "validateObject (%s)", name)
		}
	}
	return nil
}

// TODO: Generate some code to remove the reflection from this
func validateArrayChildren(itemsRule *system.RuleHolder, rules []system.Rule, data interface{}, path string, imports map[string]string) error {

	value := reflect.Indirect(reflect.ValueOf(data))
	if value.Kind() != reflect.Slice {
		return kerr.New("HQOKLQABIA", nil, "process.validateArrayChildren", "value.Kind %s must be Slice", value.Kind())
	}

	for i := 0; i < value.Len(); i++ {
		child := value.Index(i).Interface()
		if err := validateObject(itemsRule, rules, child, path, imports); err != nil {
			return kerr.New("DKVEPIWTPI", err, "process.validateArrayChildren", "validateObject")
		}
	}
	return nil
}

// TODO: Generate some code to remove the reflection from this
func validateMapChildren(itemsRule *system.RuleHolder, rules []system.Rule, data interface{}, path string, imports map[string]string) error {

	value := reflect.Indirect(reflect.ValueOf(data))
	if value.Kind() != reflect.Map {
		return kerr.New("GMEQUWCFBQ", nil, "process.validateMapChildren", "value.Kind %s must be Map", value.Kind())
	}
	if value.Type().Key() != reflect.TypeOf("") {
		return kerr.New("BPEHPQWUSG", nil, "process.validateMapChildren", "Map key %s must be string", value.Type().Key())
	}

	for _, key := range value.MapKeys() {
		child := value.MapIndex(key).Interface()
		if err := validateObject(itemsRule, rules, child, path, imports); err != nil {
			return kerr.New("YLONAMFUAG", err, "process.validateMapChildren", "validateObject key %s %#v", key, child)
		}
	}
	return nil
}
