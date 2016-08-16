package models

import (
	"sync"

	"context"

	"kego.io/system/node"
)

type BranchContentsInterface interface {
	Label(ctx context.Context) string
}

type AsyncInterface interface {
	Loaded() bool
}

type NodeContentsInterface interface {
	GetNode() *node.Node
	GetName() string
}

type SourceContentsInterface interface {
	NodeContentsInterface
	GetFilename() string
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
	return c.Node.Label(ctx)
}

type SourceContents struct {
	NodeContents
	Once     sync.Once
	Filename string
}

func (c SourceContents) Label(ctx context.Context) string {
	return c.Name
}

func (c SourceContents) Loaded() bool {
	return c.Node != nil
}

func (c SourceContents) GetFilename() string {
	return c.Filename
}
