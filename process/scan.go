package process

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"io"

	"bytes"

	"github.com/ghodss/yaml"
	"kego.io/json"
	"kego.io/kerr"
	"kego.io/system"
)

func ScanForAliases(set settings) (map[string]string, error) {

	aliases := map[string]string{}

	scanner := func(ob interface{}) error {
		if i, ok := ob.(*system.Imports); ok {
			for path, alias := range i.Imports {
				aliases[path] = alias
			}
		}
		return nil
	}
	if err := scanPath(true, true, scanner, set); err != nil {
		return nil, err
	}
	return aliases, nil
}

func ScanForGlobals(set settings) error {
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
	return scanPath(false, false, scanner, set)
}

func ScanForTypes(ignoreUnknownTypes bool, set settings) error {
	scanner := func(ob interface{}) error {
		if t, ok := ob.(*system.Type); ok {

			system.RegisterType(t.Id.Package, t.Id.Name, t)

			if t.Rule != nil {

				system.RegisterType(t.Rule.Id.Package, t.Rule.Id.Name, t.Rule)

			} else {

				// If the rule is missing, automatically create a default.
				ref := system.NewReference(set.path, fmt.Sprintf("@%s", t.Id.Name))
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
	return scanPath(ignoreUnknownTypes, false, scanner, set)
}

func scanPath(ignoreUnknownTypes bool, ignoreUnknownPackages bool, scan func(ob interface{}) error, set settings) error {

	walker := func(filePath string, file os.FileInfo, err error) error {
		if err != nil {
			return kerr.New("RSYYBBHVQK", err, "process.Scan", "walker (%s)", filePath)
		}
		if err := scanFile(filePath, ignoreUnknownTypes, ignoreUnknownPackages, scan, set.path, set.aliases); err != nil {
			return kerr.New("EMFAEDUFRS", err, "process.Scan", "processScannedFile (%s)", filePath)
		}
		return nil
	}

	if set.recursive {
		if err := filepath.Walk(set.dir, walker); err != nil {
			return kerr.New("XHHQSAVCKK", err, "process.Scan", "filepath.Walk (scanning for types)")
		}
	} else {
		files, err := ioutil.ReadDir(set.dir)
		if err != nil {
			return kerr.New("CDYLDBLHKT", err, "process.Scan", "ioutil.ReadDir")
		}
		for _, f := range files {
			if err := walker(filepath.Join(set.dir, f.Name()), f, nil); err != nil {
				return kerr.New("IAPRUHFTAD", err, "process.Scan", "walker")
			}
		}
	}

	return nil
}

func scanFile(filePath string, ignoreUnknownTypes bool, ignoreUnknownPackages bool, scan func(ob interface{}) error, packagePath string, aliases map[string]string) error {

	if !strings.HasSuffix(filePath, ".json") && !strings.HasSuffix(filePath, ".yaml") && !strings.HasSuffix(filePath, ".yml") {
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return kerr.New("NMWROTKPLJ", err, "process.processScannedFile", "os.Open (%s)", filePath)
	}
	defer file.Close()

	var reader io.Reader
	if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
		y, err := ioutil.ReadAll(file)
		if err != nil {
			return kerr.New("AXNOMOAWDF", err, "process.scanFile", "ioutil.ReadAll (yml)")
		}
		j, err := yaml.YAMLToJSON(y)
		if err != nil {
			return kerr.New("FAFJCYESRH", err, "process.scanFile", "yaml.YAMLToJSON")
		}
		reader = bytes.NewReader(j)
	} else {
		reader = file
	}

	if err = scanReader(reader, ignoreUnknownTypes, ignoreUnknownPackages, scan, packagePath, aliases); err != nil {
		return kerr.New("DHTURNTIXE", err, "process.processScannedFile", "processReader (%s)", filePath)
	}
	return nil
}

func scanReader(file io.Reader, ignoreUnknownTypes bool, ignoreUnknownPackages bool, scan func(ob interface{}) error, packagePath string, aliases map[string]string) error {

	var i interface{}
	err := json.NewDecoder(file, packagePath, aliases).Decode(&i)

	if ut, ok := err.(json.UnknownTypeError); ok {
		if !ignoreUnknownTypes {
			return kerr.New("FKCPTUWJWW", nil, "process.processReader", "json.NewDecoder.Decode: unknown type %s", ut.UnknownType)
		}
	} else if up, ok := err.(json.UnknownPackageError); ok {
		if !ignoreUnknownPackages {
			return kerr.New("KWNPDUJNYP", nil, "process.processReader", "json.NewDecoder.Decode: unknown package %s", up.UnknownPackage)
		}
	} else if err != nil {
		return kerr.New("DSMDNTCPOQ", err, "process.processReader", "json.NewDecoder.Decode")
	}

	return scan(i)
}
