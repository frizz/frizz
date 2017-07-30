package validation

import (
	frizz "frizz.io/frizz"
	system "frizz.io/system"
	validators "frizz.io/validators"
	errors "github.com/pkg/errors"
)

const Packer packer = 0

type packer int

func (p packer) Path() string {
	return "frizz.io/tests/validation"
}
func (p packer) Unpack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	case "Simple":
		return p.UnpackSimple(root, stack, in)
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) UnpackSimple(root *frizz.Root, stack frizz.Stack, in interface{}) (value Simple, null bool, err error) {
	if in == nil {
		return value, true, nil
	}
	// aliasUnpacker
	out, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value struct {
		String string
		Int    int
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.Errorf("%s: unpacking into struct, value should be a map", stack)
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out struct {
			String string
			Int    int
		}
		if v, ok := m["String"]; ok {
			stack := stack.Append(frizz.FieldItem("String"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value string, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackString(stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.String = u
			}
		}
		if v, ok := m["Int"]; ok {
			stack := stack.Append(frizz.FieldItem("Int"))
			u, null, err := func(root *frizz.Root, stack frizz.Stack, in interface{}) (value int, null bool, err error) {
				if in == nil {
					return value, true, nil
				}
				// nativeUnpacker
				out, null, err := frizz.UnpackInt(stack, in)
				if err != nil {
					return value, false, err
				}
				if null {
					return value, true, nil
				}
				return out, false, nil
			}(root, stack, v)
			if err != nil {
				return value, false, err
			}
			if !null {
				out.Int = u
			}
		}
		return out, false, nil
	}(root, stack, in)
	if err != nil {
		return value, false, err
	}
	if null {
		return value, true, nil
	}
	return Simple(out), false, nil
}
func (p packer) Repack(root *frizz.Root, stack frizz.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	case "Simple":
		return p.RepackSimple(root, stack, in.(Simple))
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packer) RepackSimple(root *frizz.Root, stack frizz.Stack, in Simple) (value interface{}, dict bool, null bool, err error) {
	return func(root *frizz.Root, stack frizz.Stack, in struct {
		String string
		Int    int
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 3)
		empty := true
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in string) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackString(in)
		}(root, stack, in.String); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["String"] = v
			}
		}
		if v, _, null, err := func(root *frizz.Root, stack frizz.Stack, in int) (value interface{}, dict bool, null bool, err error) {
			// nativeRepacker
			return frizz.RepackNumber(in)
		}(root, stack, in.Int); err != nil {
			return nil, false, false, err
		} else {
			if !null {
				empty = false
				out["Int"] = v
			}
		}
		return out, false, empty, nil
	}(root, stack, (struct {
		String string
		Int    int
	})(in))
}

const Types types = 0

type types int

func (t types) Path() string {
	return "frizz.io/tests/validation"
}
func (t types) Get(name string) string {
	switch name {
	case "Simple":
		return "ewoJIl9pbXBvcnQiOiB7CgkJInN5c3RlbSI6ICJmcml6ei5pby9zeXN0ZW0iLAoJCSJ2YWxpZGF0b3JzIjogImZyaXp6LmlvL3ZhbGlkYXRvcnMiCgl9LAoJIl90eXBlIjogInN5c3RlbS5UeXBlIiwKCSJWYWxpZGF0b3JzIjogWwoJCXsKCQkJIl90eXBlIjogInZhbGlkYXRvcnMuU3RydWN0IiwKCQkJIl92YWx1ZSI6IHsKCQkJCSJTdHJpbmciOiBbCgkJCQkJewoJCQkJCQkiX3R5cGUiOiAidmFsaWRhdG9ycy5SZWdleCIsCgkJCQkJCSJSZWdleCI6ICJeZm9vLiokIgoJCQkJCX0KCQkJCV0sCgkJCQkiSW50IjogWwoJCQkJCXsKCQkJCQkJIl90eXBlIjogInZhbGlkYXRvcnMuR3JlYXRlclRoYW4iLAoJCQkJCQkiX3ZhbHVlIjogMgoJCQkJCX0KCQkJCV0KCQkJfQoJCX0KCV0KfQ=="
	}
	return ""
}

const Imports imports = 0

type imports int

func (i imports) Path() string {
	return "frizz.io/tests/validation"
}
func (i imports) Add(packers map[string]frizz.Packer, types map[string]frizz.Typer) {
	if packers != nil {
		packers["frizz.io/tests/validation"] = Packer
	}
	if types != nil {
		types["frizz.io/tests/validation"] = Types
	}
	system.Imports.Add(packers, types)
	validators.Imports.Add(packers, types)
}
