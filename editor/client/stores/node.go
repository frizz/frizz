package stores

import (
	"strconv"

	"time"

	"github.com/davelondon/kerr"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/models"
	"kego.io/flux"
	"kego.io/json"
	"kego.io/process/validate"
	"kego.io/system/node"
)

type NodeStore struct {
	*flux.Store
	ctx context.Context
	app *App

	selected *node.Node
}

func (s *NodeStore) Selected() *node.Node {
	return s.selected
}

type nodeNotif string

func (b nodeNotif) IsNotif() {}

const (
	NodeInitialised            nodeNotif = "NodeInitialised"
	NodeValueChanged           nodeNotif = "NodeValueChanged"
	NodeDescendantValueChanged nodeNotif = "NodeDescendantValueChanged"
	NodeFocus                  nodeNotif = "NodeFocus"
)

func NewNodeStore(ctx context.Context) *NodeStore {
	s := &NodeStore{
		Store: &flux.Store{},
		ctx:   ctx,
		app:   FromContext(ctx),
	}
	s.Init(s)
	return s
}

func (s *NodeStore) Handle(payload *flux.Payload) bool {
	switch action := payload.Action.(type) {
	case *actions.ArrayOrder:
		if !action.Model.Node.Type.IsNativeArray() {
			s.app.Fail <- kerr.New("EPBQVIICFM", "Must be array")
			break
		}
		if err := action.Model.Node.ReorderArrayChild(action.OldIndex, action.NewIndex); err != nil {
			s.app.Fail <- kerr.Wrap("QNHBUSLMXF", err)
		}

	case *actions.BranchSelecting:
		if ni, ok := action.Branch.Contents.(models.NodeContentsInterface); ok {
			s.selected = ni.GetNode()
		} else {
			s.selected = nil
		}
	case *actions.BranchSelected:
		if ni, ok := action.Branch.Contents.(models.NodeContentsInterface); ok {
			s.selected = ni.GetNode()
		} else {
			s.selected = nil
		}
	case *actions.DeleteNode:
		parent := action.Node.Parent
		if parent.Type.IsNativeMap() {
			if err := parent.DeleteMapChild(action.Node.Key); err != nil {
				s.app.Fail <- kerr.Wrap("AKHHXJBDSL", err)
			}
		} else if parent.Type.IsNativeArray() {
			if err := parent.DeleteArrayChild(action.Node.Index); err != nil {
				s.app.Fail <- kerr.Wrap("PWYVESSUXV", err)
			}
		} else if parent.Type.IsNativeObject() {
			if err := parent.DeleteObjectChild(s.ctx, action.Node.Key); err != nil {
				s.app.Fail <- kerr.Wrap("AYBJTREOOE", err)
			}
		}
	case *actions.InitializeNode:
		if action.Parent.Type.IsNativeArray() {
			action.Node.AddToArray(s.ctx, action.Parent, len(action.Parent.Array), false)
		} else if action.Parent.Type.IsNativeMap() {
			action.Node.AddToMap(s.ctx, action.Parent, action.Key, false)
		}
		if err := action.Node.SetValueZero(s.ctx, false, action.Type); err != nil {
			s.app.Fail <- kerr.Wrap("NLSRNQGLLW", err)
		}
	case *actions.EditorFocus:
		s.app.Notify(action.Editor.Node, NodeFocus)
	case *actions.EditorValueChange:
		payload.Wait(s.app.Editors)
		go func() {
			<-time.After(time.Millisecond * 50)
			if action.Editor.TemporaryValue == action.Value {
				n := action.Editor.Node
				switch n.JsonType {
				case json.J_STRING:
					val := action.Value.(string)
					if n.ValueString == val {
						return
					}
					n.SetValueString(s.ctx, val)
				case json.J_BOOL:
					val := action.Value.(bool)
					if n.ValueBool == val {
						return
					}
					n.SetValueBool(s.ctx, val)
				case json.J_NUMBER:
					val, err := strconv.ParseFloat(action.Value.(string), 64)
					if err != nil {
						return
					}
					if n.ValueNumber == val {
						return
					}
					n.SetValueNumber(s.ctx, val)
				}

				if err := validate.ValidateNode(s.ctx, n); err != nil {
					if ve, ok := kerr.Source(err).(validate.ValidationError); ok {
						action.Editor.Invalid = true
						action.Editor.Error = ve.Description
					} else {
						s.app.Fail <- kerr.Wrap("BPGDPLCXKK", err)
					}
				} else {
					action.Editor.Invalid = false
					action.Editor.Error = ""
				}

				s.app.Notify(action.Editor, EditorValueChanged)
				s.app.Notify(n, NodeValueChanged)

				c := n.Parent
				for c != nil {
					s.app.Notify(c, NodeDescendantValueChanged)
					c = c.Parent
				}

			}
		}()
	}
	return true
}
