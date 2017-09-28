package usessilent

import (
	json "encoding/json"
	global "frizz.io/global"
	silent "frizz.io/tests/packer/silent"
	errors "github.com/pkg/errors"
	time "time"
)

const Package packageType = 0

type packageType int

func (p packageType) Path() string {
	return "frizz.io/tests/packer/usessilent"
}
func (p packageType) Unpack(context global.DataContext, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Time":
		return p.UnpackTime(context, in)
	case "Uses":
		return p.UnpackUses(context, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) UnpackTime(context global.DataContext, in interface{}) (value Time, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value time.Time, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker
		p := context.Package().Get("time")
		if p != nil {
			out, null, err := p.Unpack(context, in, "Time")
			if err != nil || null {
				return value, null, err
			}
			return out.(time.Time), false, nil
		}
		var vi interface{} = &value
		if vu, ok := vi.(json.Unmarshaler); ok {
			b, err := json.Marshal(in)
			if err != nil {
				return value, false, err
			}
			if err := vu.UnmarshalJSON(b); err != nil {
				return value, false, err
			}
			return value, false, nil
		}
		return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "time", "Time")
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Time(out), false, nil
}
func (p packageType) UnpackUses(context global.DataContext, in interface{}) (value Uses, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(context global.DataContext, in interface{}) (value silent.Silent, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// selectorUnpacker
		p := context.Package().Get("frizz.io/tests/packer/silent")
		if p != nil {
			out, null, err := p.Unpack(context, in, "Silent")
			if err != nil || null {
				return value, null, err
			}
			return out.(silent.Silent), false, nil
		}
		var vi interface{} = &value
		if vu, ok := vi.(json.Unmarshaler); ok {
			b, err := json.Marshal(in)
			if err != nil {
				return value, false, err
			}
			if err := vu.UnmarshalJSON(b); err != nil {
				return value, false, err
			}
			return value, false, nil
		}
		return value, false, errors.Errorf("%s: can't unpack %s.%s", context.Location(), "frizz.io/tests/packer/silent", "Silent")
	}(context, in)
	if err != nil || null {
		return value, null, err
	}
	return Uses(out), false, nil
}
func (p packageType) Repack(context global.DataContext, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Time":
		return p.RepackTime(context, in.(Time))
	case "Uses":
		return p.RepackUses(context, in.(Uses))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", context.Location(), name)
}
func (p packageType) RepackTime(context global.DataContext, in Time) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in time.Time) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker
		p := context.Package().Get("time")
		if p != nil {
			out, dict, null, err := p.Repack(context, in, "Time")
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}
		var vi interface{} = &in
		if vu, ok := vi.(json.Marshaler); ok {
			b, err := vu.MarshalJSON()
			if err != nil {
				return nil, false, false, err
			}
			if err := json.Unmarshal(b, &value); err != nil {
				return nil, false, false, err
			}
			return value, b[0] == '{', len(b) == 4 && string(b) == "null", nil
		}
		return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "time", "Time")
	}(context, (time.Time)(in))
}
func (p packageType) RepackUses(context global.DataContext, in Uses) (value interface{}, dict bool, null bool, err error) {
	return func(context global.DataContext, in silent.Silent) (value interface{}, dict bool, null bool, err error) {
		// selectorRepacker
		p := context.Package().Get("frizz.io/tests/packer/silent")
		if p != nil {
			out, dict, null, err := p.Repack(context, in, "Silent")
			if err != nil {
				return nil, false, false, err
			}
			return out, dict, null, nil
		}
		var vi interface{} = &in
		if vu, ok := vi.(json.Marshaler); ok {
			b, err := vu.MarshalJSON()
			if err != nil {
				return nil, false, false, err
			}
			if err := json.Unmarshal(b, &value); err != nil {
				return nil, false, false, err
			}
			return value, b[0] == '{', len(b) == 4 && string(b) == "null", nil
		}
		return value, false, false, errors.Errorf("%s: can't repack %s.%s", context.Location(), "frizz.io/tests/packer/silent", "Silent")
	}(context, (silent.Silent)(in))
}
func (p packageType) GetData(filename string) string {
	return ""
}
func (p packageType) GetType(name string) string {
	return ""
}
func (p packageType) GetImportedPackages(packages map[string]global.Package) {
	packages["frizz.io/tests/packer/usessilent"] = Package
	silent.Package.GetImportedPackages(packages)
}
