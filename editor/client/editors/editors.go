package editors

import (
	"bytes"
	"math/rand"

	"time"

	"github.com/davelondon/vecty"
	"github.com/davelondon/vecty/elem"
	"github.com/davelondon/vecty/prop"
	"golang.org/x/net/context"
	"kego.io/editor/client/actions"
	"kego.io/editor/client/clientctx"
	"kego.io/editor/client/models"
	"kego.io/editor/client/stores"
	"kego.io/system"
	"kego.io/system/node"
)

func Register(ctx context.Context) {
	// Don't do this. Implement the Editable interface instead. We can't do this for system types
	// so we use this method instead.
	editors := clientctx.FromContext(ctx)

	editors.Set("string", new(StringEditor))
	editors.Set("kego.io/json:string", new(StringEditor))
	editors.Set("kego.io/system:string", new(StringEditor))

	editors.Set("number", new(NumberEditor))
	editors.Set("kego.io/json:number", new(NumberEditor))
	editors.Set("kego.io/system:number", new(NumberEditor))
	editors.Set("kego.io/system:int", new(NumberEditor))

	editors.Set("bool", new(BoolEditor))
	editors.Set("kego.io/json:bool", new(BoolEditor))
	editors.Set("kego.io/system:bool", new(BoolEditor))
}

func change(app *stores.App, model *models.EditorModel, valueFunc func() interface{}) {
	v := valueFunc()
	go func() {
		<-time.After(time.Millisecond * 50)
		if v == valueFunc() {
			app.Dispatch(&actions.EditorValueChange50ms{
				Editor: model,
				Value:  v,
			})
		}
	}()
	go func() {
		<-time.After(time.Millisecond * 500)
		if v == valueFunc() {
			app.Dispatch(&actions.EditorValueChange500ms{
				Editor: model,
				Value:  v,
			})
		}
	}()
}

func helpBlock(ctx context.Context, n *node.Node) vecty.Markup {
	if n.Rule == nil {
		return vecty.List{}
	}
	description := n.Rule.Interface.(system.ObjectInterface).GetObject(ctx).Description
	if description == "" {
		return vecty.List{}
	}
	return elem.Paragraph(
		prop.Class("help-block"),
		vecty.Text(description),
	)
}

func errorBlock(ctx context.Context, m *models.NodeModel) vecty.Markup {
	if !m.Invalid {
		return vecty.List{}
	}

	errors := vecty.List{}
	for _, e := range m.Errors {
		errors = append(errors, elem.ListItem(vecty.Text(e.Description)))
	}
	return elem.Paragraph(
		prop.Class("help-block text-danger"),
		elem.UnorderedList(errors),
	)
}

func randomId() string {
	randInt := func(min int, max int) int {
		return min + rand.Intn(max-min)
	}
	var result bytes.Buffer
	var temp string
	for i := 0; i < 20; {
		if string(randInt(65, 90)) != temp {
			temp = string(randInt(65, 90))
			result.WriteString(temp)
			i++
		}
	}
	return result.String()
}
