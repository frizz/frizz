package process

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"kego.io/json"
	"kego.io/system"
)

func Scan(dir string, packageName string, packagePath string, imports map[string]string, types map[string]*system.Type) error {

	walker := func(path string, file os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".json") {

			file, err := ioutil.ReadFile(path)
			if err != nil {
				return fmt.Errorf("Error reading file %v in %v:\n%v\n", path, packageName, err)
			}

			var i interface{}
			if err := json.Unmarshal(file, &i, packagePath, imports); err != nil {
				return fmt.Errorf("Error when processing %v:\n%v\n", path, err)
			}

			t, ok := i.(*system.Type)
			if ok {
				fullname := fmt.Sprintf("%s:%s", packagePath, t.Id)
				types[fullname] = t
				system.RegisterType(fullname, t)
				if t.Rule != nil {
					rulename := fmt.Sprintf("%s:%s", packagePath, t.Rule.Id)
					types[rulename] = t.Rule
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
						Interface: system.NewBool(false),
					}
					types[rulename] = rule
					system.RegisterType(rulename, rule)
				}
			}

			return nil

		}
		return nil
	}

	if err := filepath.Walk(dir, walker); err != nil {
		return fmt.Errorf("Error while scanning for types:\n%v\n", err)
	}
	return nil
}
