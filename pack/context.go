package pack

import (
	"bytes"

	"frizz.io/global"
	"frizz.io/utils"
)

func NewContext(path string, all map[string]global.Package) *context {
	return &context{
		path: path,
		all:  all,
	}
}

type context struct {
	path string
	all  map[string]global.Package
}

func (c *context) Path() string {
	return c.path
}

func (c *context) Register(packages ...global.Package) {
	for _, p := range packages {
		c.all[p.Path()] = p
	}
}

func (c *context) Package(path string) global.Package {
	p, ok := c.all[path]
	if !ok {
		return nil
	}
	return p
}

func (c *context) Unmarshal(in []byte) (interface{}, error) {
	data, err := utils.DecodeReader(bytes.NewBuffer(in))
	if err != nil {
		return nil, err
	}
	v, null, err := Unpack(c, data)
	if err != nil || null {
		return nil, err
	}
	if null {
		return nil, nil
	}
	return v, nil
}
