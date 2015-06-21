package process

import (
	"bytes"
	"fmt"
	"go/format"
	"strconv"
	"text/template"

	"strings"

	"kego.io/kerr"
	"kego.io/literal"
	"kego.io/system"
)

func Generate(file fileType, path string, imports map[string]string) (source []byte, err error) {

	switch file {
	case F_CMD_TYPES, F_CMD_MAIN, F_CMD_VALIDATE:
		var templateName string
		switch file {
		case F_CMD_MAIN:
			templateName = "cmd_main.tmpl"
		case F_CMD_TYPES:
			templateName = "cmd_types.tmpl"
		case F_CMD_VALIDATE:
			templateName = "cmd_validate.tmpl"
		}
		cmdData := cmdDataStruct{Path: path, Imports: imports}
		source, err = executeTemplateAndFormat(cmdData, templateName)
		if err != nil {
			err = kerr.New("NXIWSECRLL", err, "process.Generate", "executeTemplateAndFormat (cmd)")
		}
	case F_MAIN, F_TYPES:
		types := system.GetAllTypesInPackage(path)
		if len(types) == 0 {
			err = kerr.New("HQLAEMCHBM", nil, "process.Generate", "No types found")
			return
		}
		switch file {
		case F_MAIN:
			mainData := mainDataStruct{Types: types, Path: path, Imports: imports}
			source, err = executeTemplateAndFormat(mainData, "main.tmpl")
			if err != nil {
				err = kerr.New("XTIEALKSXN", err, "process.Generate", "executeTemplateAndFormat (main)")
			}
		case F_TYPES:

			typesData := typesDataStruct{}

			typesPath := fmt.Sprintf("%s/types", path)

			typesData.NonTypesPath = path
			typesData.Path = typesPath

			// Clone the imports map because we have to add the base package
			typesData.Imports = map[string]string{}
			for k, v := range imports {
				typesData.Imports[k] = v
			}

			// Add the base package, but not if it's system or json
			if path != "kego.io/system" {
				name := getPackageNameFromPath(path)
				typesData.Imports[name] = path
			}

			typesData.Types = map[system.Reference]string{}
			pointersOrder := []string{}
			pointersMap := map[string]string{}
			for r, t := range types {
				typesData.Types[r] = literal.Build(t, pointersMap, &pointersOrder, typesPath, typesData.Imports)
			}
			typesData.Pointers = orderPointers(pointersOrder, pointersMap)

			//typesData := getTypesDataFromMainData(mainData)
			source, err = executeTemplateAndFormat(typesData, "types.tmpl")
			if err != nil {
				err = kerr.New("UURNHCUYAI", err, "process.Generate", "executeTemplateAndFormat (types)")
			}
		}
	case F_GLOBALS:
		globals := system.GetAllGlobalsInPackage(path)
		if len(globals) == 0 {
			return
		}

		globalData := globalDataStruct{Path: path, Imports: imports}
		globalData.Globals = map[system.Reference]string{}
		pointersOrder := []string{}
		pointersMap := map[string]string{}
		for r, g := range globals {
			globalData.Globals[r] = literal.Build(g, pointersMap, &pointersOrder, path, imports)
		}

		globalData.Pointers = orderPointers(pointersOrder, pointersMap)

		source, err = executeTemplateAndFormat(globalData, "global.tmpl")
		if err != nil {
			err = kerr.New("GPDPVPHXCY", err, "process.Generate", "executeTemplateAndFormat (global)")
		}
	}
	return
}

func orderPointers(pointersOrder []string, pointersMap map[string]string) []pointer {
	pointers := []pointer{}
	for _, name := range pointersOrder {
		pointers = append(pointers, pointer{name, pointersMap[name]})
	}
	return pointers
}

type pointer struct {
	Name   string
	Source string
}

func executeTemplateAndFormat(data interface{}, templateName string) ([]byte, error) {

	var rendered bytes.Buffer
	if err := templates().ExecuteTemplate(&rendered, templateName, data); err != nil {
		return nil, kerr.New("SGHJCEHQMF", err, "process.executeTemplateAndFormat", "templates.ExecuteTemplate")
	}

	formatted, err := format.Source(rendered.Bytes())
	if err != nil {
		return nil, kerr.New("XTKWMEDWKI", err, "process.executeTemplateAndFormat", "format.Source:\n%s\n", rendered.Bytes())
	}

	return formatted, nil

}

type mainDataStruct struct {
	Types   map[system.Reference]*system.Type
	Path    string
	Imports map[string]string
}
type globalDataStruct struct {
	Pointers []pointer
	Globals  map[system.Reference]string
	Path     string
	Imports  map[string]string
}
type typesDataStruct struct {
	Pointers     []pointer
	Types        map[system.Reference]string
	Path         string
	Imports      map[string]string
	NonTypesPath string
}
type cmdDataStruct struct {
	Path    string
	Imports map[string]string
}

func functions() template.FuncMap {
	return template.FuncMap{
		"quote":        strconv.Quote,
		"import":       importStatement,
		"types":        typesImportStatement,
		"ternary":      ternary,
		"map":          mapHelper,
		"reference":    system.GoReference,
		"reference_id": system.IdToGoReference,
		"name":         getPackageNameFromPath,
		"description":  description,
		"goname":       system.IdToGoName,
		"gotype":       system.GoTypeDescriptor,
	}
}

func description(i interface{}) (string, error) {
	o, ok := i.(system.Object)
	if !ok {
		return "", kerr.New("HQIOAQBDAX", nil, "process.description", "input does not implement system.Object")
	}
	b := o.GetBase()
	if b.Description == "" {
		return "", nil
	}
	return fmt.Sprintf("// %s", b.Description), nil
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
// TODO: Lots of packages have a different name to the path...
// TODO: Work out what to do here.
func getPackageNameFromPath(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

// importStatement generates the source code for an import statement. If the
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

// typesImportStatement generates the source code for an import statement
// for the types sub-package.
func typesImportStatement(path string, currentPackage string) string {
	newPath := fmt.Sprintf("%s/types", path)
	if currentPackage == newPath {
		return ""
	}
	return fmt.Sprintf("_ %s", strconv.Quote(newPath))
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
		return nil, kerr.New("AHGBMCNALB", nil, "process.mapHelper", "Must be an even number of values. Got %v", len(values))
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, kerr.New("WLHGIPIEUI", nil, "process.mapHelper", "All keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
