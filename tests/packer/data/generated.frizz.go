package data

import (
	global "frizz.io/global"
	packer "frizz.io/tests/packer"
	sub "frizz.io/tests/packer/sub"
	errors "github.com/pkg/errors"
)

const Package packageType = 0

type packageType int

func (p packageType) Path() string {
	return "frizz.io/tests/packer/data"
}
func (p packageType) Unpack(context global.Context, root global.Root, stack global.Stack, in interface{}, name string) (value interface{}, null bool, err error) {
	switch name {
	}
	return nil, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packageType) Repack(context global.Context, root global.Root, stack global.Stack, in interface{}, name string) (value interface{}, dict bool, null bool, err error) {
	switch name {
	}
	return nil, false, false, errors.Errorf("%s: type %s not found", stack, name)
}
func (p packageType) Get(name string) string {
	return ""
}
func (p packageType) Add(packages map[string]global.Package) {
	packer.Package.Add(packages)
	sub.Package.Add(packages)
}
