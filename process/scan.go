package process

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/surge/cityhash"
	"kego.io/json"
	"kego.io/kerr"
	"kego.io/system"
)

func ScanForPackage(set settings) error {

	found := false
	var pkg *system.Package
	var hash uint64
	scanner := func(ob interface{}, source []byte, h uint64) error {
		if p, ok := ob.(*system.Package); ok {
			if found {
				return kerr.New("NCUXDSPSDA", nil, "Found two package objects. Each package should only have one.")
			}
			pkg = p
			hash = h
			found = true
		}
		return nil
	}
	if err := scanPath(true, true, scanner, set); err != nil {
		return kerr.New("KOFQKHLSIT", err, "scanPath")
	}
	if pkg == nil {
		pkg = &system.Package{
			Object_base: &system.Object_base{
				Type: system.NewReference("kego.io/system", "package"),
			},
		}
	}
	system.RegisterPackage(set.path, pkg, hash)
	return nil
}

func ScanForGlobals(set settings) error {
	scanner := func(ob interface{}, source []byte, hash uint64) error {
		if _, ok := ob.(*system.Type); ok {
			return nil
		}
		o, ok := ob.(system.Object)
		if !ok {
			return nil
		}
		b := o.Object()
		if !b.Id.Exists {
			// Anything without an ID is not a global
			return nil
		}
		if err := system.Register(b.Id.Package, b.Id.Name, o, hash); err != nil {
			return kerr.New("IYSUOFVSGR", err, "system.Register")
		}
		return nil
	}
	return scanPath(false, false, scanner, set)
}

func HasSourceFiles(set settings) (bool, error) {
	hasFiles := false
	scanner := func(ob interface{}, source []byte, hash uint64) error {
		hasFiles = true
		return nil
	}
	if err := scanPath(true, false, scanner, set); err != nil {
		return false, kerr.New("VAMBAWHTCS", err, "scanPath (ScanForKegoFiles)")
	}
	return hasFiles, nil
}

func ScanForSource(set settings) error {
	scanner := func(ob interface{}, source []byte, hash uint64) error {
		o, ok := ob.(system.Object)
		if !ok {
			return nil
		}
		b := o.Object()
		if !b.Id.Exists {
			// Anything without an ID is not a global
			return nil
		}
		if err := system.RegisterSource(b.Id, b.Type, source, hash); err != nil {
			return kerr.New("DFDDLHNYUW", err, "system.RegisterSource")
		}
		return nil
	}
	return scanPath(false, false, scanner, set)
}

func ScanForTypes(ignoreUnknownTypes bool, set settings) error {
	scanner := func(ob interface{}, source []byte, hash uint64) error {
		if t, ok := ob.(*system.Type); ok {

			if err := system.Register(t.Id.Package, t.Id.Name, t, hash); err != nil {
				return kerr.New("XIANHMCTKB", err, "system.Register (type)")
			}

			if t.Base != nil {

				if !t.Interface {
					return kerr.New("OOUDHHYVOE", nil, "Only interface types can have a base type")
				}

				if t.Base.Id.Exists {
					return kerr.New("OINLQNXBJK", nil, "Base types should not have id specified")
				}

				t.Base.Id = system.NewReference(t.Id.Package, fmt.Sprint("$", t.Id.Name))

				if err := system.Register(t.Base.Id.Package, t.Base.Id.Name, t.Base, hash); err != nil {
					return kerr.New("DRKPPLVLRO", err, "system.Register (base)")
				}
			}

			if t.Rule != nil {

				if t.Rule.Id.Exists {
					return kerr.New("LOTEAIWAAW", nil, "Rule types should not have id specified")
				}

				t.Rule.Id = system.NewReference(t.Id.Package, fmt.Sprint("@", t.Id.Name))

				if err := system.Register(t.Rule.Id.Package, t.Rule.Id.Name, t.Rule, hash); err != nil {
					return kerr.New("GNOJJEYXHX", err, "system.Register (rule)")
				}

			} else {

				// If the rule is missing, automatically create a default.
				ref := system.NewReference(t.Id.Package, fmt.Sprint("@", t.Id.Name))
				rule := &system.Type{
					Object_base: &system.Object_base{
						Description: fmt.Sprintf("Automatically created basic rule for %s", t.Id.Name),
						Type:        system.NewReference("kego.io/system", "type"),
						Id:          ref,
					},
					Is:        []system.Reference{system.NewReference("kego.io/system", "rule")},
					Native:    system.NewString("object"),
					Interface: false,
				}
				if err := system.Register(ref.Package, ref.Name, rule, hash); err != nil {
					return kerr.New("YWGPPDVMCV", err, "system.Register (automatic rule)")
				}
			}
		}
		return nil
	}
	return scanPath(ignoreUnknownTypes, false, scanner, set)
}

func scanPath(ignoreUnknownTypes bool, ignoreUnknownPackages bool, scan func(ob interface{}, source []byte, hash uint64) error, set settings) error {

	walker := func(filePath string, file os.FileInfo, err error) error {
		if err != nil {
			return kerr.New("RSYYBBHVQK", err, "walker (%s)", filePath)
		}
		if err := scanFile(filePath, ignoreUnknownTypes, ignoreUnknownPackages, scan, set); err != nil {
			return kerr.New("EMFAEDUFRS", err, "processScannedFile (%s)", filePath)
		}
		return nil
	}

	if set.recursive {
		if err := filepath.Walk(set.dir, walker); err != nil {
			return kerr.New("XHHQSAVCKK", err, "filepath.Walk (scanning for types)")
		}
	} else {
		files, err := ioutil.ReadDir(set.dir)
		if err != nil {
			return kerr.New("CDYLDBLHKT", err, "ioutil.ReadDir")
		}
		for _, f := range files {
			if err := walker(filepath.Join(set.dir, f.Name()), f, nil); err != nil {
				return kerr.New("IAPRUHFTAD", err, "walker")
			}
		}
	}

	return nil
}

func scanFile(filePath string, ignoreUnknownTypes bool, ignoreUnknownPackages bool, scan func(ob interface{}, source []byte, hash uint64) error, set settings) error {

	bytes, hash, err := openFile(filePath, set)
	if err != nil {
		return kerr.New("JHSOCKOTHE", err, "openFile")
	}
	if bytes == nil {
		return nil
	}

	if err = scanBytes(bytes, hash, ignoreUnknownTypes, ignoreUnknownPackages, scan, set); err != nil {
		return kerr.New("DHTURNTIXE", err, "processReader (%s)", filePath)
	}
	return nil
}

func scanBytes(file []byte, hash uint64, ignoreUnknownTypes bool, ignoreUnknownPackages bool, scan func(ob interface{}, source []byte, hash uint64) error, set settings) error {

	var i interface{}
	err := json.Unmarshal(file, &i, set.path, set.aliases)

	if ut, ok := err.(json.UnknownTypeError); ok {
		if !ignoreUnknownTypes {
			return kerr.New("FKCPTUWJWW", nil, "json.NewDecoder.Decode: unknown type %s", ut.UnknownType)
		}
	} else if up, ok := err.(json.UnknownPackageError); ok {
		if !ignoreUnknownPackages {
			return kerr.New("KWNPDUJNYP", nil, "json.NewDecoder.Decode: unknown package %s", up.UnknownPackage)
		}
	} else if err != nil {
		return kerr.New("DSMDNTCPOQ", err, "json.NewDecoder.Decode")
	}

	return scan(i, file, hash)
}

// openFile opens a file, optionally converts from yml to json, and returns a byte slice and a hash of the contents.
func openFile(filePath string, set settings) ([]byte, uint64, error) {

	if !strings.HasSuffix(filePath, ".json") && !strings.HasSuffix(filePath, ".yaml") && !strings.HasSuffix(filePath, ".yml") {
		return nil, 0, nil
	}

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, 0, kerr.New("NMWROTKPLJ", err, "ioutil.ReadFile (%s)", filePath)
	}

	if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
		j, err := yaml.YAMLToJSON(bytes)
		if err != nil {
			return nil, 0, kerr.New("FAFJCYESRH", err, "yaml.YAMLToJSON")
		}
		bytes = j
	}

	relative, err := filepath.Rel(set.dir, filePath)
	if err != nil {
		return nil, 0, kerr.New("MDNIWARJEG", err, "filepath.Rel")
	}

	hash, err := getHash(relative, set.path, set.aliases, bytes)
	if err != nil {
		return nil, 0, kerr.New("GKUPQSADWQ", err, "getHash")
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

	bytes, err := json.MarshalPlain(holder)
	if err != nil {
		return 0, kerr.New("TGAEJVECIF", err, "json.Marshal")
	}

	hash := cityhash.CityHash64(bytes, uint32(len(bytes)))

	return hash, nil
}
