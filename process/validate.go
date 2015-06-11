package process

import (
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"kego.io/json"
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
	if err := json.NewDecoder(file, packagePath, imports).Decode(&i); err != nil {
		return kerr.New("QIVNOQKCQF", err, "process.validateReader", "json.NewDecoder.Decode")
	}
	if err := validateUnknown(i, packagePath, imports); err != nil {
		return kerr.New("RVKNMWKQHD", err, "process.validateReader", "validateUnknown")
	}
	return nil
}

func validateUnknown(data interface{}, path string, imports map[string]string) error {

	tp, ok := data.(system.Typer)
	if !ok {
		return kerr.New("RSSFIFHCOF", nil, "process.validateUnknown", "Input %T is not *system.Object", data)
	}
	t, ok := tp.GetType()
	if !ok {
		return kerr.New("BNQTCEDVGV", nil, "process.validateUnknown", "Type not found")
	}

	// We construct a partial RuleHolder because we don't actually have a Rule object for the
	// main object in a file

	partialRuleHolder := &system.RuleHolder{ParentType: t, Path: path, Imports: imports}

	return validateObject(partialRuleHolder, data, path, imports)

}

func validateObject(rule *system.RuleHolder, data interface{}, path string, imports map[string]string) error {

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

	// TODO: Get selectors working
	/*
		for selector, rule := range rules {
			if selector == "doesnt apply to the current object" {
				continue
			}
			if e, ok := rule.(system.Enforcer); ok {
				ok, message, err := e.Enforce(data, path, imports)
				if err != nil {
					kerr.New("MGHHDYTXVV", err, "process.validateObject", "e.Enforce")
				}
				if !ok {
					return kerr.New("FRXEXSTARP", nil, "process.validateObject", "Broken rule. %s. %#v", message, data)
				}
			}
		}
	*/

	// Validate the children
	switch rule.ParentType.Native.Value {
	case "object":
		return validateObjectChildren(rule, data, path, imports)
	case "array":
		items, err := rule.ItemsRule()
		if err != nil {
			return kerr.New("YFNERJIKWF", err, "process.validateObject", "rule.ItemsRule (array)")
		}
		return validateArrayChildren(items, data, path, imports)
	case "map":
		items, err := rule.ItemsRule()
		if err != nil {
			return kerr.New("PRPQQJKIKF", err, "process.validateObject", "rule.ItemsRule (map)")
		}
		return validateMapChildren(items, data, path, imports)
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

	for name, property := range itemsRule.ParentType.Properties {
		child, _, _, found, _, err := system.GetObjectField(value, system.IdToGoName(name))
		if err != nil {
			return kerr.New("XTUKWWRDHH", err, "process.validateObjectChildren", "system.GetField (%s)", name)
		}
		if !property.Optional && !found {
			return kerr.New("ETODESNSET", nil, "process.validateObjectChildren", "Field %s is missing and not optional", name)
		}
		childRule, err := system.NewRuleHolder(property.Item, path, imports)
		if err != nil {
			return kerr.New("IQOXVXBLRO", err, "process.validateObjectChildren", "system.NewRuleHolder (%s)", name)
		}
		if err = validateObject(childRule, child, path, imports); err != nil {
			return kerr.New("YJYSAOQWSJ", err, "process.validateObjectChildren", "validateObject (%s)", name)
		}
	}
	return nil
}

// TODO: Generate some code to remove the reflection from this
func validateArrayChildren(itemsRule *system.RuleHolder, data interface{}, path string, imports map[string]string) error {

	value := reflect.Indirect(reflect.ValueOf(data))
	if value.Kind() != reflect.Slice {
		return kerr.New("HQOKLQABIA", nil, "process.validateArrayChildren", "value.Kind %s must be Slice", value.Kind())
	}

	for i := 0; i < value.Len(); i++ {
		child := value.Index(i).Interface()
		if err := validateObject(itemsRule, child, path, imports); err != nil {
			return kerr.New("DKVEPIWTPI", err, "process.validateArrayChildren", "validateObject")
		}
	}
	return nil
}

// TODO: Generate some code to remove the reflection from this
func validateMapChildren(itemsRule *system.RuleHolder, data interface{}, path string, imports map[string]string) error {

	value := reflect.Indirect(reflect.ValueOf(data))
	if value.Kind() != reflect.Map {
		return kerr.New("GMEQUWCFBQ", nil, "process.validateMapChildren", "value.Kind %s must be Map", value.Kind())
	}
	if value.Type().Key() != reflect.TypeOf("") {
		return kerr.New("BPEHPQWUSG", nil, "process.validateMapChildren", "Map key %s must be string", value.Type().Key())
	}

	for _, key := range value.MapKeys() {
		child := value.MapIndex(key).Interface()
		if err := validateObject(itemsRule, child, path, imports); err != nil {
			return kerr.New("YLONAMFUAG", err, "process.validateMapChildren", "validateObject")
		}
	}
	return nil
}
