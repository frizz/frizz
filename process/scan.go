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

func Scan(root string, ignoreUnknownTypes bool, recursive bool, packagePath string, imports map[string]string) error {

	walker := func(filePath string, file os.FileInfo, err error) error {
		if err != nil {
			return kerr.New("RSYYBBHVQK", err, "process.Scan", "walker (%s)", filePath)
		}
		if err := processScannedFile(filePath, ignoreUnknownTypes, packagePath, imports); err != nil {
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

func processScannedFile(filePath string, ignoreUnknownTypes bool, packagePath string, imports map[string]string) error {

	if !strings.HasSuffix(filePath, ".json") {
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return kerr.New("NMWROTKPLJ", err, "process.processScannedFile", "os.Open (%s)", filePath)
	}
	defer file.Close()

	if err = processReader(file, ignoreUnknownTypes, packagePath, imports); err != nil {
		return kerr.New("DHTURNTIXE", err, "process.processScannedFile", "processReader (%s)", filePath)
	}
	return nil
}

func processReader(file io.Reader, ignoreUnknownTypes bool, packagePath string, imports map[string]string) error {

	var i interface{}
	unknown, err := json.NewDecoder(file, packagePath, imports).Decode(&i)
	if err != nil {
		return kerr.New("DSMDNTCPOQ", err, "process.processReader", "json.NewDecoder.Decode")
	}
	if unknown && !ignoreUnknownTypes {
		return kerr.New("KWNPDUJNYP", nil, "process.processReader", "json.NewDecoder.Decode: unknown types")
	}

	processScannedObject(i, packagePath, imports)
	return nil
}

func processScannedObject(i interface{}, packagePath string, imports map[string]string) {

	t, ok := i.(*system.Type)
	if ok {

		fullname := fmt.Sprintf("%s:%s", packagePath, t.Id)
		system.RegisterType(fullname, t)

		if t.Rule != nil {

			rulename := fmt.Sprintf("%s:%s", packagePath, t.Rule.Id)
			system.RegisterType(rulename, t.Rule)

		} else {

			// If the rule is missing, automatically create a default.
			id := fmt.Sprintf("@%s", t.Id)
			rulename := fmt.Sprintf("%s:%s", packagePath, id)
			rule := &system.Type{
				Base: &system.Base{
					Description: fmt.Sprintf("Automatically created basic rule for %s", t.Id),
					Type:        system.NewReference("kego.io/system", "type"),
					Id:          id,
					Context:     t.Base.Context.Clone(),
				},
				Is: []system.Reference{system.NewReference("kego.io/system", "rule")},
				Embed: []system.Reference{
					system.NewReference("kego.io/system", "ruleBase"),
				},
				Native:    system.NewString("object"),
				Interface: false,
			}
			system.RegisterType(rulename, rule)
		}
	}
}
