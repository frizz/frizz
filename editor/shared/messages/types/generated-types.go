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
	ptr14 := &system.Object_base{Description: "Automatically created basic rule for sourceRequest", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@sourceRequest", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Type{Object_base: ptr14, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr16 := &system.Object_base{Description: "Automatically created basic rule for sourceResponse", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@sourceResponse", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr17 := &system.Type{Object_base: ptr16, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr18 := &system.Object_base{Description: "This is sent by the client to request a global.", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "globalRequest", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr19 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr20 := &system.Rule_base{}
	ptr21 := &system.String_rule{Object_base: ptr19, Rule_base: ptr20, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr22 := &system.Type{Object_base: ptr18, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"name": ptr21}, Is: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr23 := &system.Object_base{Description: "This is returned when the client requests a global.", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "globalResponse", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr24 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr25 := &system.Rule_base{}
	ptr26 := &system.String_rule{Object_base: ptr24, Rule_base: ptr25, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr27 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr28 := &system.Rule_base{}
	ptr29 := &system.Bool_rule{Object_base: ptr27, Rule_base: ptr28, Default: system.Bool{}}
	ptr30 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr31 := &system.Rule_base{}
	ptr32 := &system.String_rule{Object_base: ptr30, Rule_base: ptr31, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr33 := &system.Type{Object_base: ptr23, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"data": ptr26, "found": ptr29, "name": ptr32}, Is: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr34 := &system.Object_base{Description: "This is the message interface", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr35 := &system.Type{Object_base: ptr34, Base: ptr7, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr36 := &system.Object_base{Description: "This is sent by the client to request a source file.", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "sourceRequest", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr37 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr38 := &system.Rule_base{}
	ptr39 := &system.String_rule{Object_base: ptr37, Rule_base: ptr38, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr40 := &system.Type{Object_base: ptr36, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"name": ptr39}, Is: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr41 := &system.Object_base{Description: "This is returned when the client requests a source.", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "sourceResponse", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr42 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr43 := &system.Rule_base{}
	ptr44 := &system.String_rule{Object_base: ptr42, Rule_base: ptr43, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr45 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr46 := &system.Rule_base{}
	ptr47 := &system.Bool_rule{Object_base: ptr45, Rule_base: ptr46, Default: system.Bool{}}
	ptr48 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr49 := &system.Rule_base{}
	ptr50 := &system.String_rule{Object_base: ptr48, Rule_base: ptr49, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr51 := &system.Type{Object_base: ptr41, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"data": ptr44, "found": ptr47, "name": ptr50}, Is: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/editor/shared/messages", "$message", ptr7, 0xaabb15d6475cfec1)
	system.Register("kego.io/editor/shared/messages", "@globalRequest", ptr9, 0x136af35684ddacf7)
	system.Register("kego.io/editor/shared/messages", "@globalResponse", ptr11, 0x6b2590de8a9ff1d6)
	system.Register("kego.io/editor/shared/messages", "@message", ptr13, 0xaabb15d6475cfec1)
	system.Register("kego.io/editor/shared/messages", "@sourceRequest", ptr15, 0xe7f2816c1a88470b)
	system.Register("kego.io/editor/shared/messages", "@sourceResponse", ptr17, 0x4421b9f7cf76b1b0)
	system.Register("kego.io/editor/shared/messages", "globalRequest", ptr22, 0x136af35684ddacf7)
	system.Register("kego.io/editor/shared/messages", "globalResponse", ptr33, 0x6b2590de8a9ff1d6)
	system.Register("kego.io/editor/shared/messages", "message", ptr35, 0xaabb15d6475cfec1)
	system.Register("kego.io/editor/shared/messages", "sourceRequest", ptr40, 0xe7f2816c1a88470b)
	system.Register("kego.io/editor/shared/messages", "sourceResponse", ptr51, 0x4421b9f7cf76b1b0)
}
