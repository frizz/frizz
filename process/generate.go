package process

import (
	"bytes"
	"fmt"
	"go/format"
	"strconv"
	"text/template"

	"strings"

	"kego.io/system"
	"kego.io/uerr"
)

func Generate(path string, imports map[string]string) ([]byte, error) {

	types := system.GetAllTypesInPackage(path)
	if len(types) == 0 {
		return nil, uerr.New("HQLAEMCHBM", nil, "process.Generate", "No types found")
	}

	data := templateData{Types: types, Path: path, Imports: imports}

	var rendered bytes.Buffer
	if err := templates().ExecuteTemplate(&rendered, "main.tmpl", data); err != nil {
		return nil, uerr.New("SGHJCEHQMF", err, "process.Generate", "tpl.ExecuteTemplate")
	}

	formatted, err := format.Source(rendered.Bytes())
	if err != nil {
		return nil, uerr.New("XTKWMEDWKI", err, "process.Generate", "format.Source:\n%s\n", rendered.Bytes())
	}

	return formatted, nil

}

type templateData struct {
	Types   map[string]*system.Type
	Path    string
	Imports map[string]string
}

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

// getPackageNameFromPath returns the name of the package, given the package
// path. Note this is golang packages not file system paths, so we always use
// forward slash instead of os.PathSeparator
func getPackageNameFromPath(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

// importStatement generates the scoure code for an import statement. If the
// package alias is different to the name from the package path, we specify the
// alias. If we're trying to import the local package, we output nothing.
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

// ternary is a helper for template logic
func ternary(condition bool, valueIfTrue string, valueIfFalse string) string {
	if condition {
		return valueIfTrue
	} else {
		return valueIfFalse
	}
}

// mapHelper allows us to create a map inside a template
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
