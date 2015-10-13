package types

import (
	"kego.io/demo/common/units"
	_ "kego.io/demo/common/units/types"
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Object{Description: "Automatically created basic rule for icon", Id: system.Reference{Package: "kego.io/demo/common/images", Name: "@icon", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Object: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Object{Description: "Restriction rules for images", Id: system.Reference{Package: "kego.io/demo/common/images", Name: "@image", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr4 := &system.Rule{Optional: true}
	ptr5 := &system.BoolRule{Object: ptr3, Rule: ptr4, Default: system.Bool{}}
	ptr6 := &system.Type{Object: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface{"secure": ptr5}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr7 := &system.Object{Description: "Automatically created basic rule for photo", Id: system.Reference{Package: "kego.io/demo/common/images", Name: "@photo", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr8 := &system.Type{Object: ptr7, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr9 := &system.Object{Description: "This is a type of image, which just contains the url of the image", Id: system.Reference{Package: "kego.io/demo/common/images", Name: "icon", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr10 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr11 := &system.Rule{}
	ptr12 := &system.StringRule{Object: ptr10, Rule: ptr11, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr13 := &system.Type{Object: ptr9, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"url": ptr12}, Is: []system.Reference{system.Reference{Package: "kego.io/demo/common/images", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Object{Description: "This interface type represents all images - icon and photo", Id: system.Reference{Package: "kego.io/demo/common/images", Name: "image", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Type{Object: ptr14, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: ptr6}
	ptr16 := &system.Object{Description: "This represents an image, and contains path, server and protocol separately", Id: system.Reference{Package: "kego.io/demo/common/images", Name: "photo", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr17 := &system.Object{Description: "The path for the url - e.g. /foo/bar.jpg", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr18 := &system.Rule{}
	ptr19 := &system.StringRule{Object: ptr17, Rule: ptr18, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{Value: "^/.*$", Exists: true}}
	ptr20 := &system.Object{Description: "The protocol for the url - e.g. http or https", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr21 := &system.Rule{Optional: true}
	ptr22 := &system.StringRule{Object: ptr20, Rule: ptr21, Default: system.String{Value: "http", Exists: true}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr23 := &system.Object{Description: "The server for the url - e.g. www.google.com", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr24 := &system.Rule{}
	ptr25 := &system.StringRule{Object: ptr23, Rule: ptr24, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr26 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/demo/common/units", Name: "@rectangle", Exists: true}}
	ptr27 := &system.Rule{}
	ptr28 := &units.RectangleRule{Object: ptr26, Rule: ptr27}
	ptr29 := &system.Type{Object: ptr16, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"path": ptr19, "protocol": ptr22, "server": ptr25, "size": ptr28}, Is: []system.Reference{system.Reference{Package: "kego.io/demo/common/images", Name: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/common/images", "@icon", ptr1, 0x7c5035ca01145c14)
	system.Register("kego.io/demo/common/images", "@image", ptr6, 0xd22690479466a7f)
	system.Register("kego.io/demo/common/images", "@photo", ptr8, 0xf2b64533e434a543)
	system.Register("kego.io/demo/common/images", "icon", ptr13, 0x7c5035ca01145c14)
	system.Register("kego.io/demo/common/images", "image", ptr15, 0xd22690479466a7f)
	system.Register("kego.io/demo/common/images", "photo", ptr29, 0xf2b64533e434a543)
}
