package types

import (
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Base{Description: "Automatically created basic rule for globalRequest", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@globalRequest", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Base: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Base{Description: "Automatically created basic rule for globalResponse", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@globalResponse", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Type{Base: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr4 := &system.Base{Description: "Automatically created basic rule for message", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@message", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr5 := &system.Type{Base: ptr4, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr6 := &system.Base{Description: "Automatically created basic rule for messageBase", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "@messageBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr7 := &system.Type{Base: ptr6, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr8 := &system.Base{Description: "This is sent by the client to request a global.", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "globalRequest", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr9 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr10 := &system.RuleBase{}
	ptr11 := &system.String_rule{Base: ptr9, RuleBase: ptr10, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr12 := &system.Type{Base: ptr8, Embed: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "messageBase", Exists: true}}, Fields: map[string]system.Rule{"name": ptr11}, Is: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr13 := &system.Base{Description: "This is returned when the client requests a global.", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "globalResponse", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr14 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr15 := &system.RuleBase{}
	ptr16 := &system.String_rule{Base: ptr14, RuleBase: ptr15, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr17 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr18 := &system.RuleBase{}
	ptr19 := &system.Bool_rule{Base: ptr17, RuleBase: ptr18, Default: system.Bool{}}
	ptr20 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr21 := &system.RuleBase{}
	ptr22 := &system.String_rule{Base: ptr20, RuleBase: ptr21, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr23 := &system.Type{Base: ptr13, Embed: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "messageBase", Exists: true}}, Fields: map[string]system.Rule{"data": ptr16, "found": ptr19, "name": ptr22}, Is: []system.Reference{system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr24 := &system.Base{Description: "This is the message interface", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "message", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr25 := &system.Type{Base: ptr24, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr26 := &system.Base{Description: "This is a base class embedded in each message", Id: system.Reference{Package: "kego.io/editor/shared/messages", Name: "messageBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr27 := &system.Base{Description: "A globally unique string identifying this message", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr28 := &system.RuleBase{}
	ptr29 := &system.String_rule{Base: ptr27, RuleBase: ptr28, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr30 := &system.Base{Description: "If this is a response, this is the message guid that we are responding to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr31 := &system.RuleBase{}
	ptr32 := &system.String_rule{Base: ptr30, RuleBase: ptr31, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr33 := &system.Type{Base: ptr26, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"guid": ptr29, "request": ptr32}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/editor/shared/messages", "@globalRequest", ptr1, 0xb72faee9d9111cb1)
	system.Register("kego.io/editor/shared/messages", "@globalResponse", ptr3, 0x12e106b47dc01e79)
	system.Register("kego.io/editor/shared/messages", "@message", ptr5, 0x6cbf132e7d248f63)
	system.Register("kego.io/editor/shared/messages", "@messageBase", ptr7, 0x5a67e563d231a890)
	system.Register("kego.io/editor/shared/messages", "globalRequest", ptr12, 0xb72faee9d9111cb1)
	system.Register("kego.io/editor/shared/messages", "globalResponse", ptr23, 0x12e106b47dc01e79)
	system.Register("kego.io/editor/shared/messages", "message", ptr25, 0x6cbf132e7d248f63)
	system.Register("kego.io/editor/shared/messages", "messageBase", ptr33, 0x5a67e563d231a890)
}
