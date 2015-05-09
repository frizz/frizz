package process

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"kego.io/json"
	"kego.io/system"
	"kego.io/uerr"
)

func Scan(root string, packagePath string, imports map[string]string) error {

	walker := func(filePath string, file os.FileInfo, err error) error {
		if err := processScannedFile(filePath, packagePath, imports); err != nil {
			return uerr.New("EMFAEDUFRS", err, "process.Scan", "processScannedFile (%s)", filePath)
		}
		return nil
	}

	if err := filepath.Walk(root, walker); err != nil {
		return uerr.New("XHHQSAVCKK", err, "process.Scan", "filepath.Walk (scanning for types)")
	}
	return nil
}

func processScannedFile(filePath string, packagePath string, imports map[string]string) error {

	if !strings.HasSuffix(filePath, ".json") {
		return nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return uerr.New("NMWROTKPLJ", err, "process.processScannedFile", "os.Open (%s)", filePath)
	}
	defer file.Close()

	var i interface{}
	if err = json.NewDecoder(file, packagePath, imports).Decode(&i); err != nil {
		return uerr.New("DSMDNTCPOQ", err, "process.processScannedFile", "json.NewDecoder.Decode (%s)", filePath)
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
				Object: &system.Object{
					Description: fmt.Sprintf("Automatically created basic rule for %s", t.Id),
					Type:        system.NewReference("kego.io/system", "type"),
					Id:          id,
				},
				Is:        []system.Reference{system.NewReference("kego.io/system", "rule")},
				Extends:   system.NewReference("kego.io/system", "object"),
				Native:    system.NewString("object"),
				Interface: false,
			}
			system.RegisterType(rulename, rule)
		}
	}
}
