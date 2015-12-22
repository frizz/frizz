package types

import (
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Reference{Package: "kego.io/editor/shared/messages", Name: "@message"}
	ptr1 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr2 := &system.Object{Description: "Automatically created basic rule for message", Id: ptr0, Rules: []system.RuleInterface(nil), Type: ptr1}
	ptr3 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr4 := system.String("object")
	ptr5 := &system.Type{Object: ptr2, Embed: []*system.Reference{ptr3}, Fields: map[string]system.RuleInterface(nil), Native: &ptr4, Rule: (*system.Type)(nil)}
	ptr6 := &system.Reference{Package: "kego.io/editor/shared/messages", Name: "@sourceRequest"}
	ptr7 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr8 := &system.Object{Description: "Automatically created basic rule for sourceRequest", Id: ptr6, Rules: []system.RuleInterface(nil), Type: ptr7}
	ptr9 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr10 := system.String("object")
	ptr11 := &system.Type{Object: ptr8, Embed: []*system.Reference{ptr9}, Fields: map[string]system.RuleInterface(nil), Native: &ptr10, Rule: (*system.Type)(nil)}
	ptr12 := &system.Reference{Package: "kego.io/editor/shared/messages", Name: "@sourceResponse"}
	ptr13 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr14 := &system.Object{Description: "Automatically created basic rule for sourceResponse", Id: ptr12, Rules: []system.RuleInterface(nil), Type: ptr13}
	ptr15 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr16 := system.String("object")
	ptr17 := &system.Type{Object: ptr14, Embed: []*system.Reference{ptr15}, Fields: map[string]system.RuleInterface(nil), Native: &ptr16, Rule: (*system.Type)(nil)}
	ptr18 := &system.Reference{Package: "kego.io/editor/shared/messages", Name: "message"}
	ptr19 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr20 := &system.Object{Description: "This is a base class embedded in each message", Id: ptr18, Rules: []system.RuleInterface(nil), Type: ptr19}
	ptr21 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr22 := &system.Object{Description: "A globally unique string identifying this message", Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr21}
	ptr23 := &system.Rule{}
	ptr24 := &system.StringRule{Object: ptr22, Rule: ptr23, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr25 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr26 := &system.Object{Description: "If this is a response, this is the message guid that we are responding to", Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr25}
	ptr27 := &system.Rule{}
	ptr28 := &system.StringRule{Object: ptr26, Rule: ptr27, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr29 := system.String("object")
	ptr30 := &system.Type{Object: ptr20, Basic: true, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"guid": ptr24, "request": ptr28}, Native: &ptr29, Rule: (*system.Type)(nil)}
	ptr31 := &system.Reference{Package: "kego.io/editor/shared/messages", Name: "sourceRequest"}
	ptr32 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr33 := &system.Object{Description: "This is sent by the client to request a source file.", Id: ptr31, Rules: []system.RuleInterface(nil), Type: ptr32}
	ptr34 := &system.Reference{Package: "kego.io/editor/shared/messages", Name: "message"}
	ptr35 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr36 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr35}
	ptr37 := &system.Rule{}
	ptr38 := &system.StringRule{Object: ptr36, Rule: ptr37, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr39 := system.String("object")
	ptr40 := &system.Type{Object: ptr33, Embed: []*system.Reference{ptr34}, Fields: map[string]system.RuleInterface{"name": ptr38}, Native: &ptr39, Rule: (*system.Type)(nil)}
	ptr41 := &system.Reference{Package: "kego.io/editor/shared/messages", Name: "sourceResponse"}
	ptr42 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr43 := &system.Object{Description: "This is returned when the client requests a source.", Id: ptr41, Rules: []system.RuleInterface(nil), Type: ptr42}
	ptr44 := &system.Reference{Package: "kego.io/editor/shared/messages", Name: "message"}
	ptr45 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr46 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr45}
	ptr47 := &system.Rule{}
	ptr48 := &system.StringRule{Object: ptr46, Rule: ptr47, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr49 := &system.Reference{Package: "kego.io/system", Name: "@bool"}
	ptr50 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr49}
	ptr51 := &system.Rule{}
	ptr52 := &system.BoolRule{Object: ptr50, Rule: ptr51, Default: (*system.Bool)(nil)}
	ptr53 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr54 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr53}
	ptr55 := &system.Rule{}
	ptr56 := &system.StringRule{Object: ptr54, Rule: ptr55, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr57 := system.String("object")
	ptr58 := &system.Type{Object: ptr43, Embed: []*system.Reference{ptr44}, Fields: map[string]system.RuleInterface{"data": ptr48, "found": ptr52, "name": ptr56}, Native: &ptr57, Rule: (*system.Type)(nil)}
	system.Register("kego.io/editor/shared/messages", "@message", ptr5, 0x426981ef1fd02cc3)
	system.Register("kego.io/editor/shared/messages", "@sourceRequest", ptr11, 0x5a863e2afe29ca1b)
	system.Register("kego.io/editor/shared/messages", "@sourceResponse", ptr17, 0xc4b44e683f330a44)
	system.Register("kego.io/editor/shared/messages", "message", ptr30, 0x426981ef1fd02cc3)
	system.Register("kego.io/editor/shared/messages", "sourceRequest", ptr40, 0x5a863e2afe29ca1b)
	system.Register("kego.io/editor/shared/messages", "sourceResponse", ptr58, 0xc4b44e683f330a44)
}
