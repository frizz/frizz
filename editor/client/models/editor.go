package models

import (
	"context"

	"reflect"

	"github.com/dave/kerr"
	"kego.io/context/jsonctx"
	"kego.io/editor/client/clientctx"
	"kego.io/editor/client/editable"
	"kego.io/system"
	"kego.io/system/node"
)

type EditorModel struct {
	Node    *node.Node
	Deleted bool
	Data    map[interface{}]interface{}
}

func NewEditor(n *node.Node) *EditorModel {
	return &EditorModel{
		Node: n,
		Data: map[interface{}]interface{}{},
	}
}

func GetEmbedEditable(ctx context.Context, node *node.Node, embed *system.Reference) (editable.Editable, error) {

	if node == nil || node.Null || node.Missing {
		return nil, nil
	}

	if *node.Type.Id == *embed {
		return GetEditable(ctx, node), nil
	}

	jcache := jsonctx.FromContext(ctx)
	nf, df, ok := jcache.GetNewFunc(embed.Package, embed.Name)
	if !ok {
		return nil, kerr.New("DGWDERFPVV", "Can't find %s in jsonctx", embed.String())
	}
	t := reflect.TypeOf(nf())
	if df != nil {
		t = t.Elem()
	}

	v := node.Val
	for v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	var field reflect.Value
	for i := 0; i < v.Type().NumField(); i++ {
		f := v.Type().Field(i)
		if f.Anonymous && f.Type == t {
			field = v.Field(i)
			break
		}
	}
	if field == (reflect.Value{}) {
		return nil, kerr.New("UDBOWYUBER", "Can't find %s field in struct", t)
	}

	// This is the recommended method of presenting an custom editor.
	if ed, ok := field.Interface().(editable.Editable); ok {
		return ed, nil
	}

	editors := clientctx.FromContext(ctx)

	// Don't do this. Implement the Editable interface instead. We can't do this
	// for system types so we use this method instead.
	if e, ok := editors.Get(embed.String()); ok {
		return e, nil
	}

	return nil, nil

}

func GetEditable(ctx context.Context, node *node.Node) editable.Editable {

	if node == nil || node.Null || node.Missing {
		return nil
	}

	// This is the recommended method of presenting an custom editor.
	if ed, ok := node.Value.(editable.Editable); ok {
		return ed
	}

	editors := clientctx.FromContext(ctx)

	// Don't do this. Implement the Editable interface instead. We can't do
	// this for system types so we use this method instead.
	if e, ok := editors.Get(node.Type.Id.String()); ok {
		return e
	}

	if node.JsonType != "" {
		if e, ok := editors.Get(string(node.JsonType)); ok {
			return e
		}
	}

	return nil
}
