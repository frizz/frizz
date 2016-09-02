package models

import (
	"sync"

	"context"

	"kego.io/system/node"
)

type BranchContentsInterface interface {
	Label(ctx context.Context) string
}

type NodeContentsInterface interface {
	GetNode() *node.Node
	GetName() string
}

type FileContentsInterface interface {
	NodeContentsInterface
	GetFilename() string
}

type AsyncContentsInterface interface {
	Loaded() bool
}

type RootContents struct {
	Name string
}

func (c RootContents) Label(ctx context.Context) string {
	return c.Name
}

type DataContents struct{}

func (c DataContents) Label(ctx context.Context) string {
	return "data"
}

type TypesContents struct{}

func (c TypesContents) Label(ctx context.Context) string {
	return "types"
}

type NodeContents struct {
	Node *node.Node
	Name string
}

func (c NodeContents) GetNode() *node.Node {
	return c.Node
}

func (c NodeContents) GetName() string {
	return c.Name
}

func (c NodeContents) Label(ctx context.Context) string {
	if c.Name != "" {
		return c.Name
	}
	if c.Node != nil {
		return c.Node.Label(ctx)
	}
	return ""
}

type FileContents struct {
	NodeContents
	Filename string
}

func (c FileContents) GetFilename() string {
	return c.Filename
}

type AsyncContents struct {
	FileContents
	Once sync.Once
}

func (c AsyncContents) Loaded() bool {
	return c.Node != nil
}
