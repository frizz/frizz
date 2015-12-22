package types

import (
	"kego.io/system"
	_ "kego.io/system/types"
)

func init() {
	ptr0 := &system.Reference{Package: "kego.io/demo/common/units", Name: "@rectangle"}
	ptr1 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr2 := &system.Object{Description: "Automatically created basic rule for rectangle", Id: ptr0, Rules: []system.RuleInterface(nil), Type: ptr1}
	ptr3 := &system.Reference{Package: "kego.io/system", Name: "rule"}
	ptr4 := system.String("object")
	ptr5 := &system.Type{Object: ptr2, Embed: []*system.Reference{ptr3}, Fields: map[string]system.RuleInterface(nil), Native: &ptr4, Rule: (*system.Type)(nil)}
	ptr6 := &system.Reference{Package: "kego.io/demo/common/units", Name: "rectangle"}
	ptr7 := &system.Reference{Package: "kego.io/system", Name: "type"}
	ptr8 := &system.Object{Id: ptr6, Rules: []system.RuleInterface(nil), Type: ptr7}
	ptr9 := &system.Reference{Package: "kego.io/system", Name: "@int"}
	ptr10 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr9}
	ptr11 := &system.Rule{}
	ptr12 := &system.IntRule{Object: ptr10, Rule: ptr11, Default: (*system.Int)(nil), Maximum: (*system.Int)(nil), Minimum: (*system.Int)(nil), MultipleOf: (*system.Int)(nil)}
	ptr13 := &system.Reference{Package: "kego.io/system", Name: "@int"}
	ptr14 := &system.Object{Id: (*system.Reference)(nil), Rules: []system.RuleInterface(nil), Type: ptr13}
	ptr15 := &system.Rule{}
	ptr16 := &system.IntRule{Object: ptr14, Rule: ptr15, Default: (*system.Int)(nil), Maximum: (*system.Int)(nil), Minimum: (*system.Int)(nil), MultipleOf: (*system.Int)(nil)}
	ptr17 := system.String("object")
	ptr18 := &system.Type{Object: ptr8, Embed: []*system.Reference(nil), Fields: map[string]system.RuleInterface{"height": ptr12, "width": ptr16}, Native: &ptr17, Rule: (*system.Type)(nil)}
	system.Register("kego.io/demo/common/units", "@rectangle", ptr5, 0xe9ce1340e46b290e)
	system.Register("kego.io/demo/common/units", "rectangle", ptr18, 0xe9ce1340e46b290e)
}
