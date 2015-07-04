package process

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"io"

	"kego.io/json"
	"kego.io/kerr"
	"kego.io/system"
)

func ScanForImports(root string, recursive bool, packagePath string) (imports map[string]string, err error) {

	imports = map[string]string{}

	scanner := func(ob interface{}) error {
		if i, ok := ob.(*system.Imports); ok {
			for name, imp := range i.Imports {
				imports[name] = imp.Value
			}
		}
		return nil
	}
	err = scanPath(root, true, recursive, scanner, packagePath, map[string]string{})
	return
}

func ScanForGlobals(root string, recursive bool, packagePath string, imports map[string]string) error {
	scanner := func(i interface{}) error {
		if _, ok := i.(*system.Type); ok {
			return nil
		}
		o, ok := i.(system.Object)
		if !ok {
			return nil
		}
		b := o.GetBase()
		if !b.Id.Exists {
			// Anything without an ID is not a global
			return nil
		}
		system.RegisterGlobal(b.Id.Package, b.Id.Name, o)
		return nil
	}
	return scanPath(root, false, recursive, scanner, packagePath, imports)
}

func ScanForTypes(root string, ignoreUnknownTypes bool, recursive bool, packagePath string, imports map[string]string) error {
	scanner := func(ob interface{}) error {
		if t, ok := ob.(*system.Type); ok {

			system.RegisterType(t.Id.Package, t.Id.Name, t)

			if t.Rule != nil {

				system.RegisterType(t.Rule.Id.Package, t.Rule.Id.Name, t.Rule)

			} else {

				// If the rule is missing, automatically create a default.
				ref := system.NewReference(packagePath, fmt.Sprintf("@%s", t.Id.Name))
				rule := &system.Type{
					Base: &system.Base{
						Description: fmt.Sprintf("Automatically created basic rule for %s", t.Id.Name),
						Type:        system.NewReference("kego.io/system", "type"),
						Id:          ref,
					},
					Is: []system.Reference{system.NewReference("kego.io/system", "rule")},
					Embed: []system.Reference{
						system.NewReference("kego.io/system", "ruleBase"),
					},
					Native:    system.NewString("object"),
					Interface: false,
				}
				system.RegisterType(ref.Package, ref.Name, rule)
			}
		}
		return nil
	}
	return scanPath(root, ignoreUnknownTypes, recursive, scanner, packagePath, imports)
}

func scanPath(root string, ignoreUnknownTypes bool, recursive bool, scan func(ob interface{}) error, packagePath string, imports map[string]string) error {

	walker := func(filePath string, file os.FileInfo, err error) error {
		if err != nil {
			return kerr.New("RSYYBBHVQK", err, "process.Scan", "walker (%s)", filePath)
		}
		if err := scanFile(filePath, ignoreUnknownTypes, scan, packagePath, imports); err != nil {
			return kerr.New("EMFAEDUFRS", err, "process.Scan", "processScannedFile (%s)", filePath)
		}
		return nil
	}

	if recursive {
		if err := filepath.Walk(root, walker); err != nil {
			return kerr.New("XHHQSAVCKK", err, "process.Scan", "filepath.Walk (scanning for types)")
		}
	} else {
		files, err := ioutil.ReadDir(root)
		if err != nil {
			return kerr.New("CDYLDBLHKT", err, "process.Scan", "ioutil.ReadDir")
		}
		for _, f := range files {
			if err := walker(filepath.Join(root, f.Name()), f, nil); err != nil {
				return kerr.New("IAPRUHFTAD", err, "process.Scan", "walker")
			}
		}
	}

	return nil
}

func scanFile(filePath string, ignoreUnknownTypes bool, scan func(ob interface{}) error, packagePath string, imports map[string]string) error {

	if !strings.HasSuffix(filePath, ".json") {
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return kerr.New("NMWROTKPLJ", err, "process.processScannedFile", "os.Open (%s)", filePath)
	}
	defer file.Close()

	if err = scanReader(file, ignoreUnknownTypes, scan, packagePath, imports); err != nil {
		return kerr.New("DHTURNTIXE", err, "process.processScannedFile", "processReader (%s)", filePath)
	}
	return nil
}

func scanReader(file io.Reader, ignoreUnknownTypes bool, scan func(ob interface{}) error, packagePath string, imports map[string]string) error {

	var i interface{}
	unknown, err := json.NewDecoder(file, packagePath, imports).Decode(&i)
	if err != nil {
		return kerr.New("DSMDNTCPOQ", err, "process.processReader", "json.NewDecoder.Decode")
	}
	if unknown && !ignoreUnknownTypes {
		return kerr.New("KWNPDUJNYP", nil, "process.processReader", "json.NewDecoder.Decode: unknown types")
	}

	return scan(i)
}
