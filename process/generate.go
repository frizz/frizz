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
	"kego.io/uerr"
)

func functions() template.FuncMap {
	return template.FuncMap{
		"quote":     strconv.Quote,
		"import":    importStatement,
		"ternary":   ternary,
		"map":       mapHelper,
		"reference": system.GoReference,
		"name":      getPackageNameFromPath,
	}
}
func templates() *template.Template {
	tpl := template.New("").Funcs(functions())
	for name, unpacker := range _bindata {
		asset, err := unpacker()
		if err != nil {
			panic(err)
		}
		tpl = template.Must(tpl.New(name).Parse(string(asset.bytes)))
	}
	return tpl
}

type templateData struct {
	Types   map[string]*system.Type
	Path    string
	Imports map[string]string
}

func Generate(dir string, packagePath string, imports map[string]string, test bool) error {

	types := system.GetAllTypesInPackage(packagePath)
	if len(types) == 0 {
		return uerr.New("HQLAEMCHBM", nil, "process.Generate", "No types found")
	}

	data := templateData{Types: types, Path: packagePath, Imports: imports}

	var rendered bytes.Buffer
	if err := templates().ExecuteTemplate(&rendered, "main.tmpl", data); err != nil {
		return uerr.New("SGHJCEHQMF", err, "process.Generate", "tpl.ExecuteTemplate")
	}

	formatted, err := format.Source(rendered.Bytes())
	if err != nil {
		return uerr.New("XTKWMEDWKI", err, "process.Generate", "format.Source:\n%s\n", rendered.Bytes())
	}

	if test {
		fmt.Println(string(formatted))
		return nil
	}

	if err := save(dir, formatted); err != nil {
		return uerr.New("EKXNBDCASD", err, "process.Generate", "save")
	}

	return nil
}

func save(dir string, contents []byte) error {

	name := "generated.go"
	file := filepath.Join(dir, name)

	backup(file, filepath.Join(dir, fmt.Sprintf("%s.backup", name)))

	output, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return uerr.New("NWLWHSGJWP", err, "process.save", "os.OpenFile (could not open output file)")
	}
	defer output.Close()

	output.Write(contents)

	return nil
}

func backup(file string, backup string) {

	if _, err := os.Stat(backup); err == nil {
		os.Remove(backup)
	}

	if _, err := os.Stat(file); err == nil {
		os.Rename(file, backup)
	}
}
func getPackageNameFromPath(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
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
		return nil, uerr.New("AHGBMCNALB", nil, "process.mapHelper", "Must be an even number of values. Got %v", len(values))
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, uerr.New("WLHGIPIEUI", nil, "process.mapHelper", "All keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
