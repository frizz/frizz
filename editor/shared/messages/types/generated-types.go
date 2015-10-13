package types

import (
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Object{Description: "Automatically created basic rule for message", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@message", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Object: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Object{Description: "Automatically created basic rule for sourceRequest", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@sourceRequest", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Type{Object: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr4 := &system.Object{Description: "Automatically created basic rule for sourceResponse", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@sourceResponse", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr5 := &system.Type{Object: ptr4, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr6 := &system.Object{Description: "This is a base class embedded in each message", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr7 := &system.Object{Description: "A globally unique string identifying this message", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr8 := &system.Rule{}
	ptr9 := &system.StringRule{Object: ptr7, Rule: ptr8, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr10 := &system.Object{Description: "If this is a response, this is the message guid that we are responding to", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr11 := &system.Rule{}
	ptr12 := &system.StringRule{Object: ptr10, Rule: ptr11, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr13 := &system.Type{Object: ptr6, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"guid": ptr9, "request": ptr12}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Object{Description: "This is sent by the client to request a source file.", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "sourceRequest", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr16 := &system.Rule{}
	ptr17 := &system.StringRule{Object: ptr15, Rule: ptr16, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr18 := &system.Type{Object: ptr14, Embed: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}}, Fields: map[string]system.RuleInterface{"name": ptr17}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr19 := &system.Object{Description: "This is returned when the client requests a source.", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "sourceResponse", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr20 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr21 := &system.Rule{}
	ptr22 := &system.StringRule{Object: ptr20, Rule: ptr21, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr23 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr24 := &system.Rule{}
	ptr25 := &system.BoolRule{Object: ptr23, Rule: ptr24, Default: system.Bool{}}
	ptr26 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr27 := &system.Rule{}
	ptr28 := &system.StringRule{Object: ptr26, Rule: ptr27, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr29 := &system.Type{Object: ptr19, Embed: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}}, Fields: map[string]system.RuleInterface{"data": ptr22, "found": ptr25, "name": ptr28}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/editor/shared/messages", "@message", ptr1, 0x426981ef1fd02cc3)
	system.Register("kego.io/editor/shared/messages", "@sourceRequest", ptr3, 0x5a863e2afe29ca1b)
	system.Register("kego.io/editor/shared/messages", "@sourceResponse", ptr5, 0xc4b44e683f330a44)
	system.Register("kego.io/editor/shared/messages", "message", ptr13, 0x426981ef1fd02cc3)
	system.Register("kego.io/editor/shared/messages", "sourceRequest", ptr18, 0x5a863e2afe29ca1b)
	system.Register("kego.io/editor/shared/messages", "sourceResponse", ptr29, 0xc4b44e683f330a44)
}
