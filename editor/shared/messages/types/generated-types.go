package types

import (
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Object_base{Description: "This is a base class embedded in each message", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "$message", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Object_base{Description: "A globally unique string identifying this message", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr2 := &system.Rule_base{}
	ptr3 := &system.String_rule{Object_base: ptr1, Rule_base: ptr2, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr4 := &system.Object_base{Description: "If this is a response, this is the message guid that we are responding to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr5 := &system.Rule_base{}
	ptr6 := &system.String_rule{Object_base: ptr4, Rule_base: ptr5, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr7 := &system.Type{Object_base: ptr0, Base: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"guid": ptr3, "request": ptr6}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr8 := &system.Object_base{Description: "Automatically created basic rule for message", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@message", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr9 := &system.Type{Object_base: ptr8, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr10 := &system.Object_base{Description: "Automatically created basic rule for sourceRequest", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@sourceRequest", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr11 := &system.Type{Object_base: ptr10, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr12 := &system.Object_base{Description: "Automatically created basic rule for sourceResponse", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@sourceResponse", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr13 := &system.Type{Object_base: ptr12, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Object_base{Description: "This is the message interface", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Type{Object_base: ptr14, Base: ptr7, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr16 := &system.Object_base{Description: "This is sent by the client to request a source file.", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "sourceRequest", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr17 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr18 := &system.Rule_base{}
	ptr19 := &system.String_rule{Object_base: ptr17, Rule_base: ptr18, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr20 := &system.Type{Object_base: ptr16, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"name": ptr19}, Is: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr21 := &system.Object_base{Description: "This is returned when the client requests a source.", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "sourceResponse", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr22 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr23 := &system.Rule_base{}
	ptr24 := &system.String_rule{Object_base: ptr22, Rule_base: ptr23, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr25 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr26 := &system.Rule_base{}
	ptr27 := &system.Bool_rule{Object_base: ptr25, Rule_base: ptr26, Default: system.Bool{}}
	ptr28 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr29 := &system.Rule_base{}
	ptr30 := &system.String_rule{Object_base: ptr28, Rule_base: ptr29, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr31 := &system.Type{Object_base: ptr21, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"data": ptr24, "found": ptr27, "name": ptr30}, Is: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/editor/shared/messages", "$message", ptr7, 0xaabb15d6475cfec1)
	system.Register("kego.io/editor/shared/messages", "@message", ptr9, 0xaabb15d6475cfec1)
	system.Register("kego.io/editor/shared/messages", "@sourceRequest", ptr11, 0xe7f2816c1a88470b)
	system.Register("kego.io/editor/shared/messages", "@sourceResponse", ptr13, 0x4421b9f7cf76b1b0)
	system.Register("kego.io/editor/shared/messages", "message", ptr15, 0xaabb15d6475cfec1)
	system.Register("kego.io/editor/shared/messages", "sourceRequest", ptr20, 0xe7f2816c1a88470b)
	system.Register("kego.io/editor/shared/messages", "sourceResponse", ptr31, 0x4421b9f7cf76b1b0)
}
