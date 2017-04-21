package editors

import (
	"context"

	"frizz.io/editor/client/clientctx"
)

func Register(ctx context.Context) {
	// Don't do this. Implement the Editable interface instead. We can't do this
	// for system types so we use this method instead.
	editors := clientctx.FromContext(ctx)

	editors.Set("string", new(StringEditor))
	editors.Set("frizz.io/json:string", new(StringEditor))
	editors.Set("frizz.io/system:string", new(StringEditor))

	editors.Set("number", new(NumberEditor))
	editors.Set("frizz.io/json:number", new(NumberEditor))
	editors.Set("frizz.io/system:number", new(NumberEditor))
	editors.Set("frizz.io/system:int", new(NumberEditor))

	editors.Set("bool", new(BoolEditor))
	editors.Set("frizz.io/json:bool", new(BoolEditor))
	editors.Set("frizz.io/system:bool", new(BoolEditor))

	editors.Set("frizz.io/system:object", new(ObjectEditor))
}
