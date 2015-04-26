package process

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strconv"
	"text/template"

	"strings"

	"kego.io/system"
)

var (
	tpl *template.Template
)

func init() {
	funcMap := template.FuncMap{
		"quote":     strconv.Quote,
		"import":    importStatement,
		"ternary":   ternary,
		"map":       mapHelper,
		"reference": makeReference,
	}
	tpl = template.New("foo.tmpl").Funcs(funcMap)
	for templateName, templateAssetUnpackerFunction := range _bindata {
		templateAsset, err := templateAssetUnpackerFunction()
		if err != nil {
			panic(err)
		}
		tpl = template.Must(tpl.New(templateName).Parse(string(templateAsset.bytes)))
	}
}

func Generate(dir string, packageName string, packagePath string, imports map[string]string, types map[string]*system.Type, testMode bool) error {

	if len(types) == 0 {
		return fmt.Errorf("Error generating types: No types found.")
	}

	data := struct {
		PackageName string
		PackagePath string
		Types       map[string]*system.Type
		Imports     map[string]string
	}{PackageName: packageName, PackagePath: packagePath, Types: types, Imports: imports}

	var rendered bytes.Buffer
	if err := tpl.ExecuteTemplate(&rendered, "main.tmpl", data); err != nil {
		return fmt.Errorf("Error executing templates:\n%v\n", err)
	}

	formatted, err := format.Source(rendered.Bytes())
	if err != nil {
		return fmt.Errorf("Error formatting generated source:\n%v\n%s\n", err, rendered.Bytes())
	}

	if testMode {
		fmt.Printf(string(formatted))
		return nil
	}

	filename := "generated.go"
	typesPath := filepath.Join(dir, filename)
	backupPath := filepath.Join(dir, fmt.Sprintf("%v.backup", filename))

	if _, err := os.Stat(backupPath); err == nil {
		os.Remove(backupPath)
	}

	if _, err := os.Stat(typesPath); err == nil {
		os.Rename(typesPath, backupPath)
	}

	output, err := os.OpenFile(typesPath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("Could not open output file: %s", err)
	}
	defer output.Close()

	output.Write(formatted)

	return nil
}

func importStatement(name string, path string, currentPackage string) string {
	if currentPackage == path {
		return ""
	}
	parts := strings.Split(path, "/")
	if parts[len(parts)-1] == name {
		return strconv.Quote(path)
	} else {
		return fmt.Sprintf("%s %s", name, strconv.Quote(path))
	}

}
func ternary(condition bool, valueIfTrue string, valueIfFalse string) string {
	if condition {
		return valueIfTrue
	} else {
		return valueIfFalse
	}
}
func mapHelper(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func makeReference(name string, path string, imports map[string]string, localPath string) (string, error) {
	if path == localPath {
		return name, nil
	}
	for alias, importPath := range imports {
		if path == importPath {
			return fmt.Sprintf("%s.%s", alias, name), nil
		}
	}
	return "", fmt.Errorf("Error in process.makeReference: Can't find %v in imports.\n", path)
}
