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
		mainData := mainDataStruct{Types: types, Path: path, Imports: imports}
		switch file {
		case F_MAIN:
			source, err = executeTemplateAndFormat(mainData, "main.tmpl")
			if err != nil {
				err = kerr.New("XTIEALKSXN", err, "process.Generate", "executeTemplateAndFormat (main)")
			}
		case F_TYPES:
			typesData := getTypesDataFromMainData(mainData)
			source, err = executeTemplateAndFormat(typesData, "types.tmpl")
			if err != nil {
				err = kerr.New("UURNHCUYAI", err, "process.Generate", "executeTemplateAndFormat (types)")
			}
		}
	case F_DATA:
		globals := system.GetAllGlobalsInPackage(path)
		if len(globals) == 0 {
			return
		}

		globalData := globalDataStruct{Path: path, Imports: imports}
		globalData.Globals = map[string]string{}
		order := []string{}
		pointersMap := map[string]string{}
		for n, g := range globals {
			var r *system.Reference
			r, err = system.NewReferenceFromString(n, path, imports)
			if err != nil {
				err = kerr.New("MHDPFJFMKC", err, "process.Generate", "system.NewReferenceFromString")
				return
			}
			globalData.Globals[r.Id] = literal.Build(g, pointersMap, &order, path, imports)
		}
		pointers := []string{}
		for _, n := range order {
			pointers = append(pointers, fmt.Sprintf("var %s = %s", n, pointersMap[n]))
		}
		globalData.Pointers = pointers

		source, err = executeTemplateAndFormat(globalData, "global.tmpl")
		if err != nil {
			err = kerr.New("GPDPVPHXCY", err, "process.Generate", "executeTemplateAndFormat (global)")
		}
	}
	return
}

func getTypesDataFromMainData(data mainDataStruct) (new typesDataStruct) {

	new.NonTypesPath = data.Path
	new.Path = fmt.Sprintf("%s/types", data.Path)

	// Clone the imports map because we have to add the base package
	new.Imports = map[string]string{}
	for k, v := range data.Imports {
		new.Imports[k] = v
	}

	// Add the base package, but not if it's system or json
	if data.Path != "kego.io/system" && data.Path != "kego.io/json" {
		name := getPackageNameFromPath(data.Path)
		new.Imports[name] = data.Path
	}

	new.Types = map[string]string{}
	order := []string{}
	pointersMap := map[string]string{}
	for n, t := range data.Types {
		new.Types[n] = literal.Build(t, pointersMap, &order, new.Path, new.Imports)
	}

	pointers := []string{}
	for _, n := range order {
		pointers = append(pointers, fmt.Sprintf("%s := %s", n, pointersMap[n]))
	}

	new.Pointers = pointers

	return
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
	Types   map[string]*system.Type
	Path    string
	Imports map[string]string
}
type globalDataStruct struct {
	Pointers []string
	Globals  map[string]string
	Path     string
	Imports  map[string]string
}
type typesDataStruct struct {
	Pointers     []string
	Types        map[string]string
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
