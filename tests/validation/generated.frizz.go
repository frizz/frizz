package validation

import (
	global "frizz.io/global"
	pack "frizz.io/pack"
	system "frizz.io/system"
	validators "frizz.io/validators"
	errors "github.com/pkg/errors"
)

const Package packageType = 0

type packageType int

func (p packageType) Path() string {
	return "frizz.io/tests/validation"
}
func (p packageType) Unpack(context global.DataContext, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "ValidateTest":
		return p.UnpackValidateTest(context, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) UnpackValidateTest(context global.DataContext, in interface{}) (value ValidateTest, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value struct {
		Int int
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into struct, value should be a map, found: %#v", context.Location(), in)
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out struct {
			Int int
		}
		if v, ok := m["Int"]; ok {
			context := context.New(context.Location().Child(global.FieldItem("Int")))
			u, null, err := func(context global.DataContext, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := pack.UnpackInt(context.Location(), in)
				if err != nil || null {
					return value, null, err
				}
				return out, false, nil
			}(context, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int = u
			}
		}
		return out, false, nil
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return ValidateTest(out), false, nil
}
func (p packageType) Repack(context global.DataContext, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "ValidateTest":
		return p.RepackValidateTest(context, in.(ValidateTest))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) RepackValidateTest(context global.DataContext, in ValidateTest) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in struct {
		Int int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
		empty := true
		if v, _, null, err := func(context global.DataContext, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return pack.RepackNumber(in)
		}(context, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(context, (struct {
		Int int
	})(in))
}
func (p packageType) GetData(filename string) string {
	switch filename {
	case "validate-test.frizz.json":
		return "ewoJIl9pbXBvcnQiOiB7CgkJInN5cyI6ICJmcml6ei5pby9zeXN0ZW0iLAoJCSJ2YWwiOiAiZnJpenouaW8vdmFsaWRhdG9ycyIKCX0sCgkiX3R5cGUiOiAic3lzLlR5cGUiLAoJIlZhbGlkYXRvcnMiOiBbCgkJewoJCQkiX3R5cGUiOiAidmFsLlN0cnVjdCIsCgkJCSJfdmFsdWUiOiB7CgkJCQkiSW50IjogWwoJCQkJCXsKCQkJCQkJIl90eXBlIjogInZhbC5HcmVhdGVyVGhhbiIsCgkJCQkJCSJfdmFsdWUiOiAyCgkJCQkJfQoJCQkJXQoJCQl9CgkJfQoJXQp9"
	}
	return ""
}
func (p packageType) GetType(name string) string {
	switch name {
	case "ValidateTest":
		return "validate-test.frizz.json"
	}
	return ""
}
func (p packageType) GetImportedPackages(packages map[string]global.Package) {
	packages["frizz.io/tests/validation"] = Package
	system.Package.GetImportedPackages(packages)
	validators.Package.GetImportedPackages(packages)
}
func (p packageType) Loader(loader global.Loader) dataType {
	return dataType{loader}
}

var Data = Package.Loader(pack.DefaultLoader)

type dataType struct {
	loader global.Loader
}

func (d dataType) ValidateTest() system.Type {
	return d.loader.Load(Package, "validate-test.frizz.json").(system.Type)
}
