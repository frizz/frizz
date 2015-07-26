package process

import (
	"kego.io/kerr"
	"kego.io/process/generator"
	"kego.io/system"
)

// TODO: this is WIP - currently stuck getting comments to line up.
// TODO: See http://stackoverflow.com/questions/31628613 for my question.
func GenerateStructs(set settings) error {

	if err := ScanForTypes(true, set); err != nil {
		return kerr.New("TUIWNNGUSH", err, "ScanForTypes")
	}

	types := system.GetAllTypesInPackage(set.path)
	for _, t := range types {
		if err := generator.UpdateStruct(set.path, set.dir, t.Type); err != nil {
			return kerr.New("IGHWQIUVVM", err, "generator.UpdateStruct")
		}
	}

	return nil
}
