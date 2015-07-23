package process

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
	"kego.io/cityhash"
	"kego.io/json"
	"kego.io/kerr"
	"kego.io/system"
)

func ScanForAliases(set settings) (map[string]string, error) {

	aliases := map[string]string{}

	scanner := func(ob interface{}, hash uint64) error {
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
	scanner := func(i interface{}, hash uint64) error {
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
	scanner := func(ob interface{}, hash uint64) error {
		if t, ok := ob.(*system.Type); ok {

			system.RegisterType(t.Id.Package, t.Id.Name, t, hash)

			if t.Rule != nil {

				system.RegisterType(t.Rule.Id.Package, t.Rule.Id.Name, t.Rule, hash)

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
				system.RegisterType(ref.Package, ref.Name, rule, hash)
			}
		}
		return nil
	}
	return scanPath(ignoreUnknownTypes, false, scanner, set)
}

func scanPath(ignoreUnknownTypes bool, ignoreUnknownPackages bool, scan func(ob interface{}, hash uint64) error, set settings) error {

	walker := func(filePath string, file os.FileInfo, err error) error {
		if err != nil {
			return kerr.New("RSYYBBHVQK", err, "process.Scan", "walker (%s)", filePath)
		}
		if err := scanFile(filePath, ignoreUnknownTypes, ignoreUnknownPackages, scan, set); err != nil {
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

func scanFile(filePath string, ignoreUnknownTypes bool, ignoreUnknownPackages bool, scan func(ob interface{}, hash uint64) error, set settings) error {

	bytes, hash, err := openFile(filePath, set)
	if err != nil {
		return kerr.New("JHSOCKOTHE", err, "process.scanFile", "openFile")
	}
	if bytes == nil {
		return nil
	}

	if err = scanBytes(bytes, hash, ignoreUnknownTypes, ignoreUnknownPackages, scan, set); err != nil {
		return kerr.New("DHTURNTIXE", err, "process.scanFile", "processReader (%s)", filePath)
	}
	return nil
}

func scanBytes(file []byte, hash uint64, ignoreUnknownTypes bool, ignoreUnknownPackages bool, scan func(ob interface{}, hash uint64) error, set settings) error {

	var i interface{}
	err := json.Unmarshal(file, &i, set.path, set.aliases)

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

	return scan(i, hash)
}

// openFile opens a file, optionally converts from yml to json, and returns a byte slice and a hash of the contents.
func openFile(filePath string, set settings) ([]byte, uint64, error) {

	if !strings.HasSuffix(filePath, ".json") && !strings.HasSuffix(filePath, ".yaml") && !strings.HasSuffix(filePath, ".yml") {
		return nil, 0, nil
	}

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, 0, kerr.New("NMWROTKPLJ", err, "process.openFile", "ioutil.ReadFile (%s)", filePath)
	}

	if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
		j, err := yaml.YAMLToJSON(bytes)
		if err != nil {
			return nil, 0, kerr.New("FAFJCYESRH", err, "process.openFile", "yaml.YAMLToJSON")
		}
		bytes = j
	}

	relative, err := filepath.Rel(set.dir, filePath)
	if err != nil {
		return nil, 0, kerr.New("MDNIWARJEG", err, "process.openFile", "filepath.Rel")
	}

	hash, err := getHash(relative, set.path, set.aliases, bytes)
	if err != nil {
		return nil, 0, kerr.New("GKUPQSADWQ", err, "process.openFile", "getHash")
	}

	return bytes, hash, nil

}

// gets the hash of a file, including data about the file path, package path and import aliases
func getHash(relativeFilePath string, packagePath string, aliases map[string]string, content []byte) (uint64, error) {

	if aliases == nil {
		// to stop null / {} confusion in json
		aliases = map[string]string{}
	}

	holder := struct {
		File    string
		Path    string
		Aliases map[string]string
		Content []byte
	}{relativeFilePath, packagePath, aliases, content}

	bytes, err := json.Marshal(holder)
	if err != nil {
		return 0, kerr.New("TGAEJVECIF", err, "process.getHash", "json.Marshal")
	}

	hash := cityhash.CityHash64(bytes, uint32(len(bytes)))

	return hash, nil
}