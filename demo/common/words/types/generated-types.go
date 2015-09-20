package types

import (
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Object_base{Description: "Automatically created basic rule for localizer", Id: system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Object_base: ptr0, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Object_base{Description: "Automatically created basic rule for simple", Id: system.Reference{Package: "kego.io/demo/common/words", Name: "@simple", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Type{Object_base: ptr2, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr4 := &system.Object_base{Description: "Automatically created basic rule for translation", Id: system.Reference{Package: "kego.io/demo/common/words", Name: "@translation", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr5 := &system.Type{Object_base: ptr4, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr6 := &system.Object_base{Id: system.Reference{Package: "kego.io/demo/common/words", Name: "localizer", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr7 := &system.Type{Object_base: ptr6, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr8 := &system.Object_base{Id: system.Reference{Package: "kego.io/demo/common/words", Name: "simple", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr9 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr10 := &system.Rule_base{}
	ptr11 := &system.String_rule{Object_base: ptr9, Rule_base: ptr10, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr12 := &system.Type{Object_base: ptr8, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"string": ptr11}, Is: []system.Reference{system.Reference{Package: "kego.io/demo/common/words", Name: "localizer", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr13 := &system.Object_base{Description: "This represents a translated string", Id: system.Reference{Package: "kego.io/demo/common/words", Name: "translation", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr14 := &system.Object_base{Description: "The original English string", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr15 := &system.Rule_base{Optional: true}
	ptr16 := &system.String_rule{Object_base: ptr14, Rule_base: ptr15, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr17 := &system.Object_base{Description: "The translated strings", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr18 := &system.Rule_base{}
	ptr19 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr20 := &system.Rule_base{}
	ptr21 := &system.String_rule{Object_base: ptr19, Rule_base: ptr20, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr22 := &system.Map_rule{Object_base: ptr17, Rule_base: ptr18, Items: ptr21, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr23 := &system.Type{Object_base: ptr13, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"english": ptr16, "translations": ptr22}, Is: []system.Reference{system.Reference{Package: "kego.io/demo/common/words", Name: "localizer", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/common/words", "@localizer", ptr1, 0x72d207b91bb0393)
	system.Register("kego.io/demo/common/words", "@simple", ptr3, 0xb367436bb1ddac59)
	system.Register("kego.io/demo/common/words", "@translation", ptr5, 0x20d3acc377a84c88)
	system.Register("kego.io/demo/common/words", "localizer", ptr7, 0x72d207b91bb0393)
	system.Register("kego.io/demo/common/words", "simple", ptr12, 0xb367436bb1ddac59)
	system.Register("kego.io/demo/common/words", "translation", ptr23, 0x20d3acc377a84c88)
}
