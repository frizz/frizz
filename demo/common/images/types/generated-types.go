package types

import (
	"kego.io/demo/common/units"
	_ "kego.io/demo/common/units/types"
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Base{Description: "Automatically created basic rule for icon", Id: system.Reference{Package: "kego.io/demo/common/images", Name: "@icon", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Base: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Base{Description: "Restriction rules for images", Id: system.Reference{Package: "kego.io/demo/common/images", Name: "@image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr4 := &system.RuleBase{Optional: true}
	ptr5 := &system.Bool_rule{Base: ptr3, RuleBase: ptr4, Default: system.Bool{}}
	ptr6 := &system.Type{Base: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"secure": ptr5}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr7 := &system.Base{Description: "Automatically created basic rule for photo", Id: system.Reference{Package: "kego.io/demo/common/images", Name: "@photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr8 := &system.Type{Base: ptr7, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr9 := &system.Base{Description: "This is a type of image, which just contains the url of the image", Id: system.Reference{Package: "kego.io/demo/common/images", Name: "icon", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr10 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr11 := &system.RuleBase{}
	ptr12 := &system.String_rule{Base: ptr10, RuleBase: ptr11, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr13 := &system.Type{Base: ptr9, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"url": ptr12}, Is: []system.Reference{system.Reference{Package: "kego.io/demo/common/images", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Base{Description: "This interface type represents all images - icon and photo", Id: system.Reference{Package: "kego.io/demo/common/images", Name: "image", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Type{Base: ptr14, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: ptr6}
	ptr16 := &system.Base{Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "kego.io/demo/common/images", Name: "photo", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr17 := &system.Base{Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr18 := &system.RuleBase{}
	ptr19 := &system.String_rule{Base: ptr17, RuleBase: ptr18, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{Value: "^/.*$", Exists: true}}
	ptr20 := &system.Base{Description: "The protocol for the url - e.g. http or https", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr21 := &system.RuleBase{Optional: true}
	ptr22 := &system.String_rule{Base: ptr20, RuleBase: ptr21, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr23 := &system.Base{Description: "The server for the url - e.g. www.google.com", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr24 := &system.RuleBase{}
	ptr25 := &system.String_rule{Base: ptr23, RuleBase: ptr24, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr26 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/demo/common/units", Name: "@rectangle", Exists: true}}
	ptr27 := &system.RuleBase{}
	ptr28 := &units.Rectangle_rule{Base: ptr26, RuleBase: ptr27}
	ptr29 := &system.Type{Base: ptr16, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"path": ptr19, "protocol": ptr22, "server": ptr25, "size": ptr28}, Is: []system.Reference{system.Reference{Package: "kego.io/demo/common/images", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/common/images", "@icon", ptr1, 0x7c5035ca01145c14)
	system.Register("kego.io/demo/common/images", "@image", ptr6, 0x707fd069bb76ad90)
	system.Register("kego.io/demo/common/images", "@photo", ptr8, 0xf2b64533e434a543)
	system.Register("kego.io/demo/common/images", "icon", ptr13, 0x7c5035ca01145c14)
	system.Register("kego.io/demo/common/images", "image", ptr15, 0x707fd069bb76ad90)
	system.Register("kego.io/demo/common/images", "photo", ptr29, 0xf2b64533e434a543)
}
