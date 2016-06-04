package editors

import (
	"bytes"
	"math/rand"

	"golang.org/x/net/context"
	"kego.io/editor/client/clientctx"
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

	editors.Set("object", new(ObjectEditor))
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
