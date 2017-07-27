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
	}, null bool, err error) {
		if in == nil {
			return value, true, nil
		}
		// structUnpacker
		m, ok := in.(map[string]interface{})
		if !ok {
			return value, false, errors.New("unpacking into struct, value should be a map")
		}
		if len(m) == 0 {
			return value, true, nil
		}
		var out struct {
			String string
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
	}) (value interface{}, dict bool, null bool, err error) {
		// structRepacker
		out := make(map[string]interface{}, 2)
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
		return out, false, empty, nil
	}(root, stack, (struct {
		String string
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
		return "ewoJIl9pbXBvcnQiOiB7CgkJInN5c3RlbSI6ICJmcml6ei5pby9zeXN0ZW0iLAoJCSJ2YWxpZGF0b3JzIjogImZyaXp6LmlvL3ZhbGlkYXRvcnMiCgl9LAoJIl90eXBlIjogInN5c3RlbS5UeXBlIiwKCSJGaWVsZHMiOiB7CgkJIlN0cmluZyI6IHsKCQkJIlZhbGlkYXRvcnMiOiBbCgkJCQl7CgkJCQkJIl90eXBlIjogInZhbGlkYXRvcnMuUmVnZXgiLAoJCQkJCSJSZWdleCI6ICJeZm9vLiokIgoJCQkJfQoJCQldCgkJfQoJfQp9"
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
	validators.Imports.Add(packers, types)
	system.Imports.Add(packers, types)
}
