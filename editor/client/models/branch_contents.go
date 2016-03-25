package models

import (
	"fmt"

	"kego.io/system/node"
)

type BranchContentsInterface interface {
	Label() string
	Loaded() bool
	Async() bool
}

type BranchContents struct{}

func (c BranchContents) Label() string {
	return ""
}
func (c BranchContents) Loaded() bool {
	return true
}
func (c BranchContents) Async() bool {
	return false
}

type BranchContentsAsyncInterface interface {
	LoadAction(signal chan struct{}) interface{}
}

type RootContents struct {
	BranchContents
}

type DataContents struct {
	BranchContents
}

func (c DataContents) Label() string {
	return "data"
}

type TypesContents struct {
	BranchContents
}

func (c TypesContents) Label() string {
	return "types"
}

type NodeContentsInterface interface {
	GetNode() *node.Node
}

type NodeContents struct {
	BranchContents
	Node *node.Node
	Name string
}

func (c NodeContents) GetNode() *node.Node {
	return c.Node
}

func (c NodeContents) Label() string {
	if c.Name != "" {
		return c.Name
	}
	if c.Node.Index > -1 {
		return fmt.Sprintf("[%d]", c.Node.Index)
	}
	return c.Node.Key
}

type SourceContents struct {
	BranchContents
	NodeContents
	Name     string
	Filename string
	loaded   bool
}

func (c SourceContents) Label() string {
	return c.Name
}

func (c SourceContents) Loaded() bool {
	return c.loaded
}

func (c SourceContents) Async() bool {
	return true
}
