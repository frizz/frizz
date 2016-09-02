package models

import (
	"context"

	"kego.io/system"
	"kego.io/system/node"
)

type AddPopupModel struct {
	ctx     context.Context
	Visible bool
	Parent  *node.Node
	Node    *node.Node
	Types   []*system.Type
	Success bool
}

func (m *AddPopupModel) HasName() bool {
	return m.IsFile() || m.IsMap()
}

func (m *AddPopupModel) IsMap() bool {
	return m.Parent != nil && m.Parent.Type.IsNativeMap()
}

func (m *AddPopupModel) IsArray() bool {
	return m.Parent != nil && m.Parent.Type.IsNativeArray()
}

func (m *AddPopupModel) IsField() bool {
	return m.Parent != nil && m.Parent.Type.IsNativeObject()
}

func (m *AddPopupModel) IsFile() bool {
	return m.Parent == nil
}
