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
	ptr8 := &system.Object_base{Description: "Automatically created basic rule for globalRequest", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@globalRequest", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr9 := &system.Type{Object_base: ptr8, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr10 := &system.Object_base{Description: "Automatically created basic rule for globalResponse", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@globalResponse", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr11 := &system.Type{Object_base: ptr10, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr12 := &system.Object_base{Description: "Automatically created basic rule for message", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@message", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr13 := &system.Type{Object_base: ptr12, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Object_base{Description: "This is sent by the client to request a global.", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "globalRequest", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr16 := &system.Rule_base{}
	ptr17 := &system.String_rule{Object_base: ptr15, Rule_base: ptr16, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr18 := &system.Type{Object_base: ptr14, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"name": ptr17}, Is: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr19 := &system.Object_base{Description: "This is returned when the client requests a global.", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "globalResponse", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr20 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr21 := &system.Rule_base{}
	ptr22 := &system.String_rule{Object_base: ptr20, Rule_base: ptr21, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr23 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr24 := &system.Rule_base{}
	ptr25 := &system.Bool_rule{Object_base: ptr23, Rule_base: ptr24, Default: system.Bool{}}
	ptr26 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr27 := &system.Rule_base{}
	ptr28 := &system.String_rule{Object_base: ptr26, Rule_base: ptr27, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr29 := &system.Type{Object_base: ptr19, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"data": ptr22, "found": ptr25, "name": ptr28}, Is: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr30 := &system.Object_base{Description: "This is the message interface", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr31 := &system.Type{Object_base: ptr30, Base: ptr7, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/editor/shared/messages", "$message", ptr7, 0xaabb15d6475cfec1)
	system.Register("kego.io/editor/shared/messages", "@globalRequest", ptr9, 0x136af35684ddacf7)
	system.Register("kego.io/editor/shared/messages", "@globalResponse", ptr11, 0x6b2590de8a9ff1d6)
	system.Register("kego.io/editor/shared/messages", "@message", ptr13, 0xaabb15d6475cfec1)
	system.Register("kego.io/editor/shared/messages", "globalRequest", ptr18, 0x136af35684ddacf7)
	system.Register("kego.io/editor/shared/messages", "globalResponse", ptr29, 0x6b2590de8a9ff1d6)
	system.Register("kego.io/editor/shared/messages", "message", ptr31, 0xaabb15d6475cfec1)
}
