package stores

import (
	"context"

	"github.com/davelondon/kerr"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/flux"
	"kego.io/system"
	"kego.io/system/node"
)

type FileStore struct {
	*flux.Store
	ctx context.Context
	app *App

	files map[*node.Node]*models.FileModel
}

type fileNotif string

func (b fileNotif) IsNotif() {}

const (
	FileChangedStateChange fileNotif = "FileChange"
)

func NewFileStore(ctx context.Context) *FileStore {
	s := &FileStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
		files: map[*node.Node]*models.FileModel{},
	}
	s.Init(s)
	return s
}

func (s *FileStore) All() map[*node.Node]*models.FileModel {
	return s.files
}

func (s *FileStore) fileByName(filename string) *models.FileModel {
	for _, file := range s.files {
		if filename == file.Filename {
			return file
		}
	}
	return nil
}

func (s *FileStore) Changed() bool {
	for _, fi := range s.files {
		if fi.Changed() {
			return true
		}
	}
	return false
}

func (s *FileStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.Add:
		payload.Wait(s.app.Nodes)
		if action.Parent != nil {
			// only root nodes are files
			break
		}
		switch action.Direction() {
		case actions.New, actions.Redo:
			n := action.Node
			n.RecomputeHash(s.ctx, true)
			fm := &models.FileModel{
				Package:  false,
				Type:     *action.Type.Id == *system.NewReference("kego.io/system", "type"),
				Filename: action.BranchFile,
				Node:     n,
				Hash:     n.Hash(),
			}
			s.files[n] = fm
		case actions.Undo:
			delete(s.files, action.Node)
		}
		payload.Notify(nil, FileChangedStateChange)
	case *actions.SaveFileSuccess:
		for _, file := range action.Response.Files {
			f := s.fileByName(file.File)
			if f == nil {
				s.app.Fail <- kerr.New("FPXNHFJNOR", "File %s not found", file.File)
			}
			f.SaveHash = file.Hash
		}
		payload.Notify(nil, FileChangedStateChange)
	case *actions.InitialState:
		payload.Wait(s.app.Nodes)
		n := s.app.Package.Node()
		fm := &models.FileModel{
			Package:  true,
			Type:     false,
			Filename: s.app.Package.Filename(),
			Node:     n,
			Hash:     n.Hash(),
			LoadHash: n.Hash(),
		}
		s.files[n] = fm

		for _, ti := range s.app.Types.All() {
			n := ti.Node
			fm := &models.FileModel{
				Package:  false,
				Type:     true,
				Filename: ti.File,
				Node:     n,
				Hash:     n.Hash(),
				LoadHash: n.Hash(),
			}
			s.files[n] = fm
		}
	case *actions.LoadFileSuccess:
		sci, ok := action.Branch.Contents.(models.FileContentsInterface)
		if !ok {
			break
		}
		n := sci.GetNode()
		fm := &models.FileModel{
			Package:  false,
			Type:     false,
			Filename: sci.GetFilename(),
			Node:     n,
			Hash:     n.Hash(),
			LoadHash: n.Hash(),
		}
		s.files[n] = fm
	}

	if m, ok := payload.Action.(actions.Mutator); ok {
		payload.Wait(s.app.Nodes)
		n := m.CommonAncestor().Root()
		f, ok := s.files[n]
		if ok {
			prevChanged := f.Changed()
			f.Hash = n.Hash()
			newChanged := f.Changed()
			if prevChanged != newChanged {
				// state changed!
				payload.Notify(n, FileChangedStateChange)
			}
		}
	}
	return true
}
