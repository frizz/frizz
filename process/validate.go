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

func ValidateCommand(set settings) error {
	if set.verbose {
		fmt.Print("Validating... ")
	}
	if err := Validate(set); err != nil {
		if u, ok := err.(kerr.UniqueError); ok {
			if _, ok := u.Source().(TypesChangedError); ok {
				if set.verbose {
					fmt.Println("Types have changed - rebuilding...")
				}
				if err := RunAllCommands(set); err != nil {
					return err
				}
				return nil
			}
		}
		return err
	}
	if set.verbose {
		fmt.Println("OK.")
	}
	return nil
}

func Validate(set settings) error {

	walker := func(filePath string, file os.FileInfo, err error) error {
		if err != nil {
			return kerr.New("GFBBIERGIY", err, "process.Validate", "walker (%s)", filePath)
		}
		if err := validateFile(filePath, set); err != nil {
			return kerr.New("QJGXAUKPTI", err, "process.Validate", "validateFile (%s)", filePath)
		}
		return nil
	}

	if set.recursive {
		if err := filepath.Walk(set.dir, walker); err != nil {
			return kerr.New("GCKFJQJUXK", err, "process.Validate", "filepath.Walk")
		}
	} else {
		files, err := ioutil.ReadDir(set.dir)
		if err != nil {
			return kerr.New("RJXRHBYVUW", err, "process.Validate", "ioutil.ReadDir")
		}
		for _, f := range files {
			if err := walker(filepath.Join(set.dir, f.Name()), f, nil); err != nil {
				return kerr.New("UJBOWKFUMS", err, "process.Validate", "walker")
			}
		}
	}

	return nil
}

func validateFile(filePath string, set settings) error {

	bytes, hash, err := openFile(filePath, set)
	if err != nil {
		return kerr.New("XXYPVKLNBQ", err, "process.validateFile", "openFile")
	}
	if bytes == nil {
		return nil
	}

	if err = validateBytes(bytes, hash, set); err != nil {
		return kerr.New("GFVGDBDTNQ", err, "process.validateFile", "validateReader (%s)", filePath)
	}
	return nil
}

func validateBytes(bytes []byte, hash uint64, set settings) error {
	var i interface{}
	err := json.Unmarshal(bytes, &i, set.path, set.aliases)
	if up, ok := err.(json.UnknownPackageError); ok {
		return kerr.New("QPOGRNXWMH", err, "process.validateReader", "json.NewDecoder: unknown package %s", up.UnknownPackage)
	} else if ut, ok := err.(json.UnknownTypeError); ok {
		return kerr.New("PJABFRVFLF", nil, "process.validateReader", "json.NewDecoder: unknown type %s", ut.UnknownType)
	} else if err != nil {
		return kerr.New("QIVNOQKCQF", err, "process.validateReader", "json.NewDecoder.Decode")
	}
	if err := validateUnknown(i, hash, set.path, set.aliases); err != nil {
		return kerr.New("RVKNMWKQHD", err, "process.validateReader", "validateUnknown")
	}
	return nil
}

func validateUnknown(data interface{}, hash uint64, path string, aliases map[string]string) error {

	ob, ok := data.(system.Object)
	if !ok {
		return kerr.New("RSSFIFHCOF", nil, "process.validateUnknown", "data %T does not implement system.Object", data)
	}

	if ty, ok := ob.(*system.Type); ok {
		_, h, ok := system.GetType(ty.Id.Package, ty.Id.Name)
		if !ok {
			return TypesChangedError{fmt.Sprintf("New type %s found - run kego again to rebuild", ty.Id.Value())}
		}
		if hash != h {
			return TypesChangedError{fmt.Sprintf("Type %s has changed - run kego again to rebuild", ty.Id.Value())}
		}
	}

	t, ok := ob.GetBase().Type.GetType()
	if !ok {
		return kerr.New("BNQTCEDVGV", nil, "process.validateUnknown", "Type.GetType %s not found", ob.GetBase().Type.Value())
	}

	partialRuleHolder := system.NewMinimalRuleHolder(t, path, aliases)

	rules := t.Rules
	if system.RulesApplyToObjects(data) {
		rules = append(rules, data.(system.Object).GetBase().Rules...)
	}

	return validateObject(partialRuleHolder, rules, data, path, aliases)

}

type TypesChangedError struct {
	Message string
}

func (v TypesChangedError) Error() string {
	return v.Message
}

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return v.Message
}

func validateObject(rule *system.RuleHolder, rules []system.Rule, data interface{}, path string, aliases map[string]string) error {

	// Validate the actual object
	if v, ok := data.(system.Validator); ok {
		ok, message, err := v.Validate(path, aliases)
		if err != nil {
			return kerr.New("RUGJLUAFAN", err, "process.validateObject", "v.Validate")
		}
		if !ok {
			return ValidationError{message}
		}
	}

	if rule.Rule != nil {
		e, ok := rule.Rule.(system.Enforcer)
		if ok {
			ok, message, err := e.Enforce(data, path, aliases)
			if err != nil {
				return kerr.New("EBEMISLGDX", err, "process.validateObject", "e.Enforce (main)")
			}
			if !ok {
				return ValidationError{message}
			}
		}
	}

	if rules != nil && len(rules) > 0 {

		j := &selectors.Element{Data: data, Value: reflect.ValueOf(data), Rule: rule}
		p, err := selectors.CreateParser(j, path, aliases)
		if err != nil {
			return kerr.New("AIWLGYGGAY", err, "process.validateObject", "selectors.CreateParser")
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
				ok, message, err := e.Enforce(match.Data, path, aliases)
				if err != nil {
					return kerr.New("MGHHDYTXVV", err, "process.validateObject", "e.Enforce")
				}
				if !ok {
					return ValidationError{message}
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
			return kerr.New("YFNERJIKWF", err, "process.validateObject", "rule.ItemsRule (array)")
		}
		rules := rule.Rule.(system.Object).GetBase().Rules
		return validateArrayChildren(items, rules, data, path, aliases)
	case "map":
		items, err := rule.ItemsRule()
		if err != nil {
			return kerr.New("PRPQQJKIKF", err, "process.validateObject", "rule.ItemsRule (map)")
		}
		rules := rule.Rule.(system.Object).GetBase().Rules
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
		return kerr.New("FJBEEDBLOK", nil, "process.validateObjectChildren", "value.Kind %s (%T) must be Struct", value.Kind(), data)
	}

	rules := []system.Rule{}
	if system.RulesApplyToObjects(data) {
		rules = data.(system.Object).GetBase().Rules
	}

	for name, field := range itemsRule.ParentType.Fields {
		child, _, _, found, _, err := system.GetObjectField(value, system.GoName(name))
		if err != nil {
			return kerr.New("XTUKWWRDHH", err, "process.validateObjectChildren", "system.GetField (%s)", name)
		}
		if !field.GetRuleBase().Optional && !found {
			return kerr.New("ETODESNSET", nil, "process.validateObjectChildren", "Field %s is missing and not optional", name)
		}
		childRule, err := system.NewRuleHolder(field, path, aliases)
		if err != nil {
			return kerr.New("IQOXVXBLRO", err, "process.validateObjectChildren", "system.NewRuleHolder (%s)", name)
		}
		ob, ok := field.(system.Object)
		if !ok {
			return kerr.New("XRTVWVUAMP", nil, "process.validateObjectChildren", "field does not implement system.Object")
		}

		allRules := append(rules, ob.GetBase().Rules...)

		// if we have additional rules on the main field rule, we should add them to allRules
		if len(childRule.Rule.(system.Object).GetBase().Rules) > 0 {
			allRules = append(allRules, childRule.Rule.(system.Object).GetBase().Rules...)
		}

		if err = validateObject(childRule, allRules, child, path, aliases); err != nil {
			return kerr.New("YJYSAOQWSJ", err, "process.validateObjectChildren", "validateObject (%s)", name)
		}
	}
	return nil
}

// TODO: Generate some code to remove the reflection from this
func validateArrayChildren(itemsRule *system.RuleHolder, rules []system.Rule, data interface{}, path string, aliases map[string]string) error {

	value := reflect.Indirect(reflect.ValueOf(data))
	if value.Kind() != reflect.Slice {
		return kerr.New("HQOKLQABIA", nil, "process.validateArrayChildren", "value.Kind %s must be Slice", value.Kind())
	}

	// if we have additional rules on the main items rule, we should add them to rules
	if len(itemsRule.Rule.(system.Object).GetBase().Rules) > 0 {
		rules = append(rules, itemsRule.Rule.(system.Object).GetBase().Rules...)
	}

	for i := 0; i < value.Len(); i++ {
		child := value.Index(i).Interface()
		if err := validateObject(itemsRule, rules, child, path, aliases); err != nil {
			return kerr.New("DKVEPIWTPI", err, "process.validateArrayChildren", "validateObject")
		}
	}
	return nil
}

// TODO: Generate some code to remove the reflection from this
func validateMapChildren(itemsRule *system.RuleHolder, rules []system.Rule, data interface{}, path string, aliases map[string]string) error {

	value := reflect.Indirect(reflect.ValueOf(data))
	if value.Kind() != reflect.Map {
		return kerr.New("GMEQUWCFBQ", nil, "process.validateMapChildren", "value.Kind %s must be Map", value.Kind())
	}
	if value.Type().Key() != reflect.TypeOf("") {
		return kerr.New("BPEHPQWUSG", nil, "process.validateMapChildren", "Map key %s must be string", value.Type().Key())
	}

	// if we have additional rules on the main items rule, we should add them to rules
	if len(itemsRule.Rule.(system.Object).GetBase().Rules) > 0 {
		rules = append(rules, itemsRule.Rule.(system.Object).GetBase().Rules...)
	}

	for _, key := range value.MapKeys() {
		child := value.MapIndex(key).Interface()
		if err := validateObject(itemsRule, rules, child, path, aliases); err != nil {
			return kerr.New("YLONAMFUAG", err, "process.validateMapChildren", "validateObject key %s %#v", key, child)
		}
	}
	return nil
}
