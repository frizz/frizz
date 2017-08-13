package data

import (
	global "frizz.io/global"
	pack "frizz.io/pack"
	packer "frizz.io/tests/packer"
	sub "frizz.io/tests/packer/sub"
	errors "github.com/pkg/errors"
)

const Package packageType = 0

type packageType int

func (p packageType) Path() string {
	return "frizz.io/tests/data"
}
func (p packageType) Unpack(context global.DataContext, in interface{}, name string) (value interface{}, null bool, err error) {
	return nil, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) Repack(context global.DataContext, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	return nil, false, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) GetData(filename string) string {
	switch filename {
	case "natives.frizz.json":
		return "ewoJIl9pbXBvcnQiOiB7CgkJInBhY2tlciI6ICJmcml6ei5pby90ZXN0cy9wYWNrZXIiCgl9LAoJIl90eXBlIjogInBhY2tlci5OYXRpdmVzIiwKCSJJbnQiOiAyCn0="
	case "sub-interface.frizz.json":
		return "ewoJIl9pbXBvcnQiOiB7CgkJInBhY2tlciI6ICJmcml6ei5pby90ZXN0cy9wYWNrZXIiLAoJCSJzdWIiOiAiZnJpenouaW8vdGVzdHMvcGFja2VyL3N1YiIKCX0sCgkiX3R5cGUiOiAicGFja2VyLlN1YkludGVyZmFjZSIsCgkiU3ViSW50ZXJmYWNlIjogewoJCSJfdHlwZSI6ICJzdWIuU3ViIiwKCQkiU3RyaW5nIjogImZvbyIKCX0KfQ=="
	}
	return ""
}
func (p packageType) GetType(name string) string {
	return ""
}
func (p packageType) GetImportedPackages(packages map[string]global.Package) {
	packages["frizz.io/tests/data"] = Package
	packer.Package.GetImportedPackages(packages)
	sub.Package.GetImportedPackages(packages)
}
func (p packageType) Loader(loader global.Loader) dataType {
	return dataType{loader}
}

var Data = Package.Loader(pack.DefaultLoader)

type dataType struct {
	loader global.Loader
}

func (d dataType) Natives() packer.Natives {
	return d.loader.Load(Package, "natives.frizz.json").(packer.Natives)
}
func (d dataType) SubInterface() packer.SubInterface {
	return d.loader.Load(Package, "sub-interface.frizz.json").(packer.SubInterface)
}
