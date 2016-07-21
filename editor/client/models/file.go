package models

import "kego.io/system/node"

type FileModel struct {
	Package  bool
	Type     bool
	Filename string
	Node     *node.Node
	LoadHash uint64
	SaveHash uint64
	Hash     uint64
}

func (f *FileModel) Changed() bool {
	prevHash := f.LoadHash
	if f.SaveHash > 0 {
		prevHash = f.SaveHash
	}
	return f.Hash != prevHash
}
