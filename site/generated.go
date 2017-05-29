package site

import (
	system "frizz.io/system"
	errors "github.com/pkg/errors"
)

func unpackPage(in interface{}) (interface{}, error) {
	m, ok := in.(map[string]interface{})
	if !ok {
		return nil, errors.Errorf("error unpacking into %s, value should be a map", "Page")
	}
	out := new(Page)
	if v, ok := m["title"]; ok {
		c, err := system.Convert_string(v)
		if err != nil {
			return nil, errors.Wrap(err, "error unpacking")
		}
		out.Title = c
	}
	if v, ok := m["body"]; ok {
		c, err := system.Convert_string(v)
		if err != nil {
			return nil, errors.Wrap(err, "error unpacking")
		}
		out.Body = c
	}
	if v, ok := m["height"]; ok {
		c, err := system.Convert_uint8(v)
		if err != nil {
			return nil, errors.Wrap(err, "error unpacking")
		}
		out.Height = c
	}
	return out, nil
}
