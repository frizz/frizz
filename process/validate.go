package process

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"fmt"

	"kego.io/json"
	"kego.io/kerr"
	"kego.io/selectors"
	"kego.io/system"
)

func ValidateCommand(set Settings) (typesChanged bool, err error) {
	if set.Verbose {
		fmt.Print("Validating... ")
	}
	if err := Validate(set); err != nil {
		if _, ok := kerr.Source(err).(TypesChangedError); ok {
			if set.Verbose {
				fmt.Println("Types have changed - rebuilding...")
			}
			if err := RunAllCommands(set); err != nil {
				return true, err
			}
			return true, nil
		}
		return false, err
	}
	if set.Verbose {
		fmt.Println("OK.")
	}
	return false, nil
}

func Validate(set Settings) error {

	walker := func(filePath string, file os.FileInfo, err error) error {
		if err != nil {
			return kerr.New("GFBBIERGIY", err, "walker (%s)", filePath)
		}
		if err := validateFile(filePath, set); err != nil {
			return kerr.New("QJGXAUKPTI", err, "validateFile (%s)", filePath)
		}
		return nil
	}

	if set.Recursive {
		if err := filepath.Walk(set.Dir, walker); err != nil {
			return kerr.New("GCKFJQJUXK", err, "filepath.Walk")
		}
	} else {
		files, err := ioutil.ReadDir(set.Dir)
		if err != nil {
			return kerr.New("RJXRHBYVUW", err, "ioutil.ReadDir")
		}
		for _, f := range files {
			if err := walker(filepath.Join(set.Dir, f.Name()), f, nil); err != nil {
				return kerr.New("UJBOWKFUMS", err, "walker")
			}
		}
	}

	return nil
}

func validateFile(filePath string, set Settings) error {

	bytes, hash, err := openFile(filePath, set)
	if err != nil {
		return kerr.New("XXYPVKLNBQ", err, "openFile")
	}
	if bytes == nil {
		return nil
	}

	if err = validateBytes(bytes, hash, set); err != nil {
		return kerr.New("GFVGDBDTNQ", err, "validateReader (%s)", filePath)
	}
	return nil
}

func validateBytes(bytes []byte, hash uint64, set Settings) error {
	var i interface{}
	err := json.Unmarshal(bytes, &i, set.Path, set.Aliases)
	if up, ok := err.(json.UnknownPackageError); ok {
		return kerr.New("QPOGRNXWMH", err, "json.NewDecoder: unknown package %s", up.UnknownPackage)
	} else if ut, ok := err.(json.UnknownTypeError); ok {
		return kerr.New("PJABFRVFLF", nil, "json.NewDecoder: unknown type %s", ut.UnknownType)
	} else if err != nil {
		return kerr.New("QIVNOQKCQF", err, "json.NewDecoder.Decode")
	}
	if err := validateUnknown(i, hash, set.Path, set.Aliases); err != nil {
		return kerr.New("RVKNMWKQHD", err, "validateUnknown")
	}
	return nil
}

func validateUnknown(data interface{}, hash uint64, path string, aliases map[string]string) error {

	ob, ok := data.(system.Object)
	if !ok {
		return kerr.New("RSSFIFHCOF", nil, "data %T does not implement system.Object", data)
	}

	if ty, ok := ob.(*system.Type); ok {
		h, ok := system.GetGlobal(ty.Id.Package, ty.Id.Name)
		if !ok {
			return TypesChangedError{kerr.New("XCBNOPEEUY", nil, "New type %s found - run kego again to rebuild", ty.Id.Value())}
		}
		if hash != h.Hash {
			return TypesChangedError{kerr.New("KRKBNITUON", nil, "Type %s has changed - run kego again to rebuild", ty.Id.Value())}
		}
	}

	t, ok := ob.Object().Type.GetType()
	if !ok {
		return kerr.New("BNQTCEDVGV", nil, "Type.GetType %s not found", ob.Object().Type.Value())
	}

	partialRuleHolder := system.NewMinimalRuleHolder(t)

	rules := t.Rules
	if system.RulesApplyToObjects(data) {
		rules = append(rules, data.(system.Object).Object().Rules...)
	}

	return validateObject(partialRuleHolder, rules, data, path, aliases)

}

type TypesChangedError struct {
	kerr.Struct
}

type ValidationError struct {
	kerr.Struct
}

func validateObject(rule *system.RuleHolder, rules []system.Rule, data interface{}, path string, aliases map[string]string) error {

	// Validate the actual object
	if v, ok := data.(system.Validator); ok {
		ok, message, err := v.Validate(path, aliases)
		if err != nil {
			return kerr.New("RUGJLUAFAN", err, "v.Validate")
		}
		if !ok {
			return ValidationError{kerr.New("KULDIJUYFB", nil, message)}
		}
	}

	if rule.Rule != nil {
		e, ok := rule.Rule.(system.Enforcer)
		if ok {
			ok, message, err := e.Enforce(data, path, aliases)
			if err != nil {
				return kerr.New("EBEMISLGDX", err, "e.Enforce (main)")
			}
			if !ok {
				return ValidationError{kerr.New("HLKQWDCMRN", nil, message)}
			}
		}
	}

	if rules != nil && len(rules) > 0 {

		j := &selectors.Element{Data: data, Value: reflect.ValueOf(data), Rule: rule}
		p, err := selectors.CreateParser(j, path, aliases)
		if err != nil {
			return kerr.New("AIWLGYGGAY", err, "selectors.CreateParser")
		}

		for _, rule := range rules {

			base := rule.Rule()

			selector := ":root"
			if base.Selector != "" {
				selector = base.Selector
			}

			e, ok := rule.(system.Enforcer)
			if !ok {
				return kerr.New("ABVWHMMXGG", nil, "rule %T does not implement system.Enforcer", rule)
			}
			matches, err := p.GetElements(selector)
			if err != nil {
				return kerr.New("UKOCCFJWAB", err, "p.GetJsonElements (%s)", selector)
			}
			for _, match := range matches {
				ok, message, err := e.Enforce(match.Data, path, aliases)
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
	switch rule.ParentType.Native.Value {
	case "object":
		return validateObjectChildren(rule, data, path, aliases)
	case "array":
		items, err := rule.ItemsRule()
		if err != nil {
			return kerr.New("YFNERJIKWF", err, "rule.ItemsRule (array)")
		}
		rules := rule.Rule.(system.Object).Object().Rules
		return validateArrayChildren(items, rules, data, path, aliases)
	case "map":
		items, err := rule.ItemsRule()
		if err != nil {
			return kerr.New("PRPQQJKIKF", err, "rule.ItemsRule (map)")
		}
		rules := rule.Rule.(system.Object).Object().Rules
		return validateMapChildren(items, rules, data, path, aliases)
	}

	return nil
}

// TODO: Generate some code to remove the reflection from this
func validateObjectChildren(itemsRule *system.RuleHolder, data interface{}, path string, aliases map[string]string) error {

	if data == nil {
		return nil
	}

	value := reflect.Indirect(reflect.ValueOf(data))
	if value.Kind() != reflect.Struct {
		return kerr.New("FJBEEDBLOK", nil, "value.Kind %s (%T) must be Struct", value.Kind(), data)
	}

	rules := []system.Rule{}
	if system.RulesApplyToObjects(data) {
		rules = data.(system.Object).Object().Rules
	}

	for name, field := range itemsRule.ParentType.Fields {
		child, _, _, found, _, err := system.GetObjectField(value, system.GoName(name))
		if err != nil {
			return kerr.New("XTUKWWRDHH", err, "system.GetField (%s)", name)
		}
		if !field.Rule().Optional && !found {
			return kerr.New("ETODESNSET", nil, "Field %s is missing and not optional", name)
		}
		childRule, err := system.NewRuleHolder(field)
		if err != nil {
			return kerr.New("IQOXVXBLRO", err, "system.NewRuleHolder (%s)", name)
		}
		ob, ok := field.(system.Object)
		if !ok {
			return kerr.New("XRTVWVUAMP", nil, "field does not implement system.Object")
		}

		allRules := append(rules, ob.Object().Rules...)

		// if we have additional rules on the main field rule, we should add them to allRules
		if len(childRule.Rule.(system.Object).Object().Rules) > 0 {
			allRules = append(allRules, childRule.Rule.(system.Object).Object().Rules...)
		}

		if err = validateObject(childRule, allRules, child, path, aliases); err != nil {
			return kerr.New("YJYSAOQWSJ", err, "validateObject (%s)", name)
		}
	}
	return nil
}

// TODO: Generate some code to remove the reflection from this
func validateArrayChildren(itemsRule *system.RuleHolder, rules []system.Rule, data interface{}, path string, aliases map[string]string) error {

	value := reflect.Indirect(reflect.ValueOf(data))
	if value.Kind() != reflect.Slice {
		return kerr.New("HQOKLQABIA", nil, "value.Kind %s must be Slice", value.Kind())
	}

	// if we have additional rules on the main items rule, we should add them to rules
	if len(itemsRule.Rule.(system.Object).Object().Rules) > 0 {
		rules = append(rules, itemsRule.Rule.(system.Object).Object().Rules...)
	}

	for i := 0; i < value.Len(); i++ {
		child := value.Index(i).Interface()
		if err := validateObject(itemsRule, rules, child, path, aliases); err != nil {
			return kerr.New("DKVEPIWTPI", err, "validateObject")
		}
	}
	return nil
}

// TODO: Generate some code to remove the reflection from this
func validateMapChildren(itemsRule *system.RuleHolder, rules []system.Rule, data interface{}, path string, aliases map[string]string) error {

	value := reflect.Indirect(reflect.ValueOf(data))
	if value.Kind() != reflect.Map {
		return kerr.New("GMEQUWCFBQ", nil, "value.Kind %s must be Map", value.Kind())
	}
	if value.Type().Key() != reflect.TypeOf("") {
		return kerr.New("BPEHPQWUSG", nil, "Map key %s must be string", value.Type().Key())
	}

	// if we have additional rules on the main items rule, we should add them to rules
	if len(itemsRule.Rule.(system.Object).Object().Rules) > 0 {
		rules = append(rules, itemsRule.Rule.(system.Object).Object().Rules...)
	}

	for _, key := range value.MapKeys() {
		child := value.MapIndex(key).Interface()
		if err := validateObject(itemsRule, rules, child, path, aliases); err != nil {
			return kerr.New("YLONAMFUAG", err, "validateObject key %s %#v", key, child)
		}
	}
	return nil
}
