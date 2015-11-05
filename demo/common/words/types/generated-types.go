package types

import (
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@localizer"}
	ptr1 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr2 := &system.Object{Description: "Automatically created basic rule for localizer", Id: ptr0, Rules: []system.RuleInterface(nil), Type: ptr1}
	ptr3 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr4 := system.String("object")
	ptr5 := &system.Type{Object: ptr2, Embed: []*system.Reference{ptr3}, Fields: map[string]system.RuleInterface(nil), Is: []*system.Reference(nil), Native: &ptr4, Rule: (*system.Type)(nil)}
	ptr6 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@simple"}
	ptr7 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr8 := &system.Object{Description: "Automatically created basic rule for simple", Id: ptr6, Rules: []system.RuleInterface(nil), Type: ptr7}
	ptr9 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr10 := system.String("object")
	ptr11 := &system.Type{Object: ptr8, Embed: []*system.Reference{ptr9}, Fields: map[string]system.RuleInterface(nil), Is: []*system.Reference(nil), Native: &ptr10, Rule: (*system.Type)(nil)}
	ptr12 := &system.Reference{Package: "kego.io/demo/common/words", Name: "@translation"}
	ptr13 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr14 := &system.Object{Description: "Automatically created basic rule for translation", Id: ptr12, Rules: []system.RuleInterface(nil), Type: ptr13}
	ptr15 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr16 := system.String("object")
	ptr17 := &system.Type{Object: ptr14, Embed: []*system.Reference{ptr15}, Fields: map[string]system.RuleInterface(nil), Is: []*system.Reference(nil), Native: &ptr16, Rule: (*system.Type)(nil)}
	ptr18 := &system.Reference{Package: "kego.io/demo/common/words", Name: "localizer"}
	ptr19 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr20 := &system.Object{Id: ptr18, Rules: []system.RuleInterface(nil), Type: ptr19}
	ptr21 := system.String("object")
	ptr22 := &system.Type{Object: ptr20, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Interface: true, Is: []*system.Reference(nil), Native: &ptr21, Rule: (*system.Type)(nil)}
	ptr23 := &system.Reference{Package: "kego.io/demo/common/words", Name: "simple"}
	ptr24 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr25 := &system.Object{Id: ptr23, Rules: []system.RuleInterface(nil), Type: ptr24}
	ptr26 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr27 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr26}
	ptr28 := &system.Rule{}
	ptr29 := &system.StringRule{Object: ptr27, Rule: ptr28, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr30 := &system.Reference{Package: "kego.io/demo/common/words", Name: "localizer"}
	ptr31 := system.String("object")
	ptr32 := &system.Type{Object: ptr25, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"string": ptr29}, Is: []*system.Reference{ptr30}, Native: &ptr31, Rule: (*system.Type)(nil)}
	ptr33 := &system.Reference{Package: "kego.io/demo/common/words", Name: "translation"}
	ptr34 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr35 := &system.Object{Description: "This represents a translated string", Id: ptr33, Rules: []system.RuleInterface(nil), Type: ptr34}
	ptr36 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr37 := &system.Object{Description: "The original English string", Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr36}
	ptr38 := &system.Rule{Optional: true}
	ptr39 := system.String("http")
	ptr40 := &system.StringRule{Object: ptr37, Rule: ptr38, Default: &ptr39, Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr41 := &system.Reference{Package: "kego.io/system", Name: "@map"}
	ptr42 := &system.Object{Description: "The translated strings", Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr41}
	ptr43 := &system.Rule{}
	ptr44 := &system.Reference{Package: "kego.io/system", Name: "@string"}
	ptr45 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr44}
	ptr46 := &system.Rule{}
	ptr47 := &system.StringRule{Object: ptr45, Rule: ptr46, Default: (*system.String)(nil), Enum: []string(nil), Equal: (*system.String)(nil), Format: (*system.String)(nil), MaxLength: (*system.Int)(nil), MinLength: (*system.Int)(nil), Pattern: (*system.String)(nil)}
	ptr48 := &system.MapRule{Object: ptr42, Rule: ptr43, Items: ptr47, MaxItems: (*system.Int)(nil), MinItems: (*system.Int)(nil)}
	ptr49 := &system.Reference{Package: "kego.io/demo/common/words", Name: "localizer"}
	ptr50 := system.String("object")
	ptr51 := &system.Type{Object: ptr35, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"english": ptr40, "translations": ptr48}, Is: []*system.Reference{ptr49}, Native: &ptr50, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/common/words", "@localizer", ptr5, 0x72d207b91bb0393)
	system.Register("kego.io/demo/common/words", "@simple", ptr11, 0xb367436bb1ddac59)
	system.Register("kego.io/demo/common/words", "@translation", ptr17, 0x20d3acc377a84c88)
	system.Register("kego.io/demo/common/words", "localizer", ptr22, 0x72d207b91bb0393)
	system.Register("kego.io/demo/common/words", "simple", ptr32, 0xb367436bb1ddac59)
	system.Register("kego.io/demo/common/words", "translation", ptr51, 0x20d3acc377a84c88)
}
