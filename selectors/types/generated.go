package types

import (
	"kego.io/system"

	_ "kego.io/system/types"
)

func init() {

	ptr8728703680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623232 := &system.Object{Context: ptr8728703680, Description: "Automatically created basic rule for typed", Id: "@typed", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728650592 := &system.Type{Object: ptr8728623232, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728342624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728782704 := &system.Object{Context: ptr8728342624, Description: "", Id: "expr", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728698016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728783712 := &system.Object{Context: ptr8728698016, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728697984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728783824 := &system.Object{Context: ptr8728697984, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553024 := &system.String_rule{Object: ptr8728783824, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488624 := &system.Property{Object: ptr8728783712, Defaulter: false, Item: ptr8728553024, Optional: false}

	ptr8728698528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728783936 := &system.Object{Context: ptr8728698528, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728698496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728784272 := &system.Object{Context: ptr8728698496, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728264096 := &system.Bool_rule{Object: ptr8728784272, Default: system.Bool{Value: false, Exists: false}}

	ptr8728488672 := &system.Property{Object: ptr8728783936, Defaulter: false, Item: ptr8728264096, Optional: false}

	ptr8728342464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728784384 := &system.Object{Context: ptr8728342464, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728342432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728784496 := &system.Object{Context: ptr8728342432, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728264544 := &system.Bool_rule{Object: ptr8728784496, Default: system.Bool{Value: false, Exists: false}}

	ptr8728489152 := &system.Property{Object: ptr8728784384, Defaulter: false, Item: ptr8728264544, Optional: false}

	ptr8728696096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728782816 := &system.Object{Context: ptr8728696096, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728695936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728782928 := &system.Object{Context: ptr8728695936, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728440016 := &system.Number_rule{Object: ptr8728782928, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728488384 := &system.Property{Object: ptr8728782816, Defaulter: false, Item: ptr8728440016, Optional: false}

	ptr8728696672 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728783040 := &system.Object{Context: ptr8728696672, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728696512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728783152 := &system.Object{Context: ptr8728696512, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728440096 := &system.Number_rule{Object: ptr8728783152, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728488432 := &system.Property{Object: ptr8728783040, Defaulter: false, Item: ptr8728440096, Optional: false}

	ptr8728697120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728783264 := &system.Object{Context: ptr8728697120, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728697088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728783376 := &system.Object{Context: ptr8728697088, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728552544 := &system.String_rule{Object: ptr8728783376, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488480 := &system.Property{Object: ptr8728783264, Defaulter: false, Item: ptr8728552544, Optional: false}

	ptr8728697568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728783488 := &system.Object{Context: ptr8728697568, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728697536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728783600 := &system.Object{Context: ptr8728697536, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728552704 := &system.String_rule{Object: ptr8728783600, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488576 := &system.Property{Object: ptr8728783488, Defaulter: false, Item: ptr8728552704, Optional: false}

	ptr8728574144 := &system.Type{Object: ptr8728782704, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"null": ptr8728488624, "true": ptr8728488672, "false": ptr8728489152, "int": ptr8728488384, "float": ptr8728488432, "string": ptr8728488480, "string2": ptr8728488576}, Rule: (*system.Type)(nil)}

	ptr8728691520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728622336 := &system.Object{Context: ptr8728691520, Description: "", Id: "basic", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728343296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728622672 := &system.Object{Context: ptr8728343296, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728343264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623008 := &system.Object{Context: ptr8728343264, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728343232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623344 := &system.Object{Context: ptr8728343232, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728552864 := &system.String_rule{Object: ptr8728623344, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728422272 := &system.Map_rule{Object: ptr8728623008, Items: ptr8728552864, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728487328 := &system.Property{Object: ptr8728622672, Defaulter: false, Item: ptr8728422272, Optional: false}

	ptr8728343744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623456 := &system.Object{Context: ptr8728343744, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728343712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623680 := &system.Object{Context: ptr8728343712, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553184 := &system.String_rule{Object: ptr8728623680, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488144 := &system.Property{Object: ptr8728623456, Defaulter: false, Item: ptr8728553184, Optional: false}

	ptr8728344832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623792 := &system.Object{Context: ptr8728344832, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728344736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728624240 := &system.Object{Context: ptr8728344736, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728344704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728624352 := &system.Object{Context: ptr8728344704, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728344672 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728624464 := &system.Object{Context: ptr8728344672, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553504 := &system.String_rule{Object: ptr8728624464, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728422592 := &system.Map_rule{Object: ptr8728624352, Items: ptr8728553504, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728422464 := &system.Array_rule{Object: ptr8728624240, Items: ptr8728422592, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728488192 := &system.Property{Object: ptr8728623792, Defaulter: false, Item: ptr8728422464, Optional: false}

	ptr8728345536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728624576 := &system.Object{Context: ptr8728345536, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728345440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728624688 := &system.Object{Context: ptr8728345440, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728345408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728624800 := &system.Object{Context: ptr8728345408, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553664 := &system.String_rule{Object: ptr8728624800, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728422656 := &system.Array_rule{Object: ptr8728624688, Items: ptr8728553664, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728488816 := &system.Property{Object: ptr8728624576, Defaulter: false, Item: ptr8728422656, Optional: false}

	ptr8728346400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728624912 := &system.Object{Context: ptr8728346400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728346208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728625024 := &system.Object{Context: ptr8728346208, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728346176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728281200 := &system.Object{Context: ptr8728346176, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553824 := &system.String_rule{Object: ptr8728281200, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728422784 := &system.Array_rule{Object: ptr8728625024, Items: ptr8728553824, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728489008 := &system.Property{Object: ptr8728624912, Defaulter: false, Item: ptr8728422784, Optional: false}

	ptr8728691392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728780800 := &system.Object{Context: ptr8728691392, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728691232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728781136 := &system.Object{Context: ptr8728691232, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728438416 := &system.Number_rule{Object: ptr8728781136, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728489104 := &system.Property{Object: ptr8728780800, Defaulter: false, Item: ptr8728438416, Optional: false}

	ptr8728575584 := &system.Type{Object: ptr8728622336, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"name": ptr8728487328, "favoriteColor": ptr8728488144, "languagesSpoken": ptr8728488192, "seatingPreference": ptr8728488816, "drinkPreference": ptr8728489008, "weight": ptr8728489104}, Rule: (*system.Type)(nil)}

	ptr8728342848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728784048 := &system.Object{Context: ptr8728342848, Description: "Automatically created basic rule for expr", Id: "@expr", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728572992 := &system.Type{Object: ptr8728784048, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728694880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728782144 := &system.Object{Context: ptr8728694880, Description: "", Id: "collision", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728694752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728782256 := &system.Object{Context: ptr8728694752, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728694720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728782368 := &system.Object{Context: ptr8728694720, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728694688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728782480 := &system.Object{Context: ptr8728694688, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728551424 := &system.String_rule{Object: ptr8728782480, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728423168 := &system.Map_rule{Object: ptr8728782368, Items: ptr8728551424, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728488000 := &system.Property{Object: ptr8728782256, Defaulter: false, Item: ptr8728423168, Optional: false}

	ptr8728573712 := &system.Type{Object: ptr8728782144, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"number": ptr8728488000}, Rule: (*system.Type)(nil)}

	ptr8728693888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728782032 := &system.Object{Context: ptr8728693888, Description: "Automatically created basic rule for c", Id: "@c", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728573424 := &system.Type{Object: ptr8728782032, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728699360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728786512 := &system.Object{Context: ptr8728699360, Description: "", Id: "sibling", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728698432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728787184 := &system.Object{Context: ptr8728698432, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728698400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728787296 := &system.Object{Context: ptr8728698400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728698240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728787408 := &system.Object{Context: ptr8728698240, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728441296 := &system.Number_rule{Object: ptr8728787408, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728423616 := &system.Map_rule{Object: ptr8728787296, Items: ptr8728441296, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728490400 := &system.Property{Object: ptr8728787184, Defaulter: false, Item: ptr8728423616, Optional: false}

	ptr8728699232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728787520 := &system.Object{Context: ptr8728699232, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728699200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728787632 := &system.Object{Context: ptr8728699200, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728699040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728787744 := &system.Object{Context: ptr8728699040, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728441376 := &system.Number_rule{Object: ptr8728787744, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728423744 := &system.Map_rule{Object: ptr8728787632, Items: ptr8728441376, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728490496 := &system.Property{Object: ptr8728787520, Defaulter: false, Item: ptr8728423744, Optional: false}

	ptr8728696256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728786624 := &system.Object{Context: ptr8728696256, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728696064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728786736 := &system.Object{Context: ptr8728696064, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728441136 := &system.Number_rule{Object: ptr8728786736, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728490208 := &system.Property{Object: ptr8728786624, Defaulter: false, Item: ptr8728441136, Optional: false}

	ptr8728696992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728786848 := &system.Object{Context: ptr8728696992, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728696800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728786960 := &system.Object{Context: ptr8728696800, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728441216 := &system.Number_rule{Object: ptr8728786960, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728490256 := &system.Property{Object: ptr8728786848, Defaulter: false, Item: ptr8728441216, Optional: false}

	ptr8728697408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728787072 := &system.Object{Context: ptr8728697408, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728490304 := &system.Property{Object: ptr8728787072, Defaulter: false, Item: map[string]interface{}{"type": "@c", "_path": "kego.io/jsonselect", "_imports": map[string]string{}}, Optional: false}

	ptr8728650016 := &system.Type{Object: ptr8728786512, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"d": ptr8728490400, "e": ptr8728490496, "a": ptr8728490208, "b": ptr8728490256, "c": ptr8728490304}, Rule: (*system.Type)(nil)}

	ptr8728694976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728786400 := &system.Object{Context: ptr8728694976, Description: "Automatically created basic rule for polykids", Id: "@polykids", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728649728 := &system.Type{Object: ptr8728786400, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728691136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728785504 := &system.Object{Context: ptr8728691136, Description: "Automatically created basic rule for kid", Id: "@kid", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728574720 := &system.Type{Object: ptr8728785504, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728703264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728787968 := &system.Object{Context: ptr8728703264, Description: "", Id: "typed", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728701856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728788864 := &system.Object{Context: ptr8728701856, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728490736 := &system.Property{Object: ptr8728788864, Defaulter: false, Item: map[string]interface{}{"type": "@image", "_path": "kego.io/jsonselect", "_imports": map[string]string{}}, Optional: false}

	ptr8728702560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728622448 := &system.Object{Context: ptr8728702560, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728702464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728622560 := &system.Object{Context: ptr8728702464, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728702432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728622784 := &system.Object{Context: ptr8728702432, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728554624 := &system.String_rule{Object: ptr8728622784, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728424000 := &system.Array_rule{Object: ptr8728622560, Items: ptr8728554624, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728490832 := &system.Property{Object: ptr8728622448, Defaulter: false, Item: ptr8728424000, Optional: false}

	ptr8728703136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728622896 := &system.Object{Context: ptr8728703136, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728702976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623120 := &system.Object{Context: ptr8728702976, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728439216 := &system.Number_rule{Object: ptr8728623120, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728490928 := &system.Property{Object: ptr8728622896, Defaulter: false, Item: ptr8728439216, Optional: false}

	ptr8728700640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728788080 := &system.Object{Context: ptr8728700640, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728700608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728788192 := &system.Object{Context: ptr8728700608, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728700576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728788304 := &system.Object{Context: ptr8728700576, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728554304 := &system.String_rule{Object: ptr8728788304, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728423680 := &system.Map_rule{Object: ptr8728788192, Items: ptr8728554304, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728489680 := &system.Property{Object: ptr8728788080, Defaulter: false, Item: ptr8728423680, Optional: false}

	ptr8728701088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728788416 := &system.Object{Context: ptr8728701088, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728701056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728788528 := &system.Object{Context: ptr8728701056, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728554464 := &system.String_rule{Object: ptr8728788528, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728490448 := &system.Property{Object: ptr8728788416, Defaulter: false, Item: ptr8728554464, Optional: false}

	ptr8728701568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728788640 := &system.Object{Context: ptr8728701568, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728701536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728788752 := &system.Object{Context: ptr8728701536, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728423872 := &system.Map_rule{Object: ptr8728788752, Items: map[string]interface{}{"type": "@kid", "_path": "kego.io/jsonselect", "_imports": map[string]string{}}, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728490592 := &system.Property{Object: ptr8728788640, Defaulter: false, Item: ptr8728423872, Optional: false}

	ptr8728650448 := &system.Type{Object: ptr8728787968, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"avatar": ptr8728490736, "drinkPreference": ptr8728490832, "weight": ptr8728490928, "name": ptr8728489680, "favoriteColor": ptr8728490448, "kids": ptr8728490592}, Rule: (*system.Type)(nil)}

	ptr8728343136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728784160 := &system.Object{Context: ptr8728343136, Description: "", Id: "image", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728575440 := &system.Type{Object: ptr8728784160, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728693472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728781024 := &system.Object{Context: ptr8728693472, Description: "", Id: "c", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728692000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728781248 := &system.Object{Context: ptr8728692000, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728691840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728781360 := &system.Object{Context: ptr8728691840, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728438976 := &system.Number_rule{Object: ptr8728781360, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728487616 := &system.Property{Object: ptr8728781248, Defaulter: false, Item: ptr8728438976, Optional: false}

	ptr8728692576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728781472 := &system.Object{Context: ptr8728692576, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728692416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728781584 := &system.Object{Context: ptr8728692416, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728439056 := &system.Number_rule{Object: ptr8728781584, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728487664 := &system.Property{Object: ptr8728781472, Defaulter: false, Item: ptr8728439056, Optional: false}

	ptr8728693344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728781696 := &system.Object{Context: ptr8728693344, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728693312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728781808 := &system.Object{Context: ptr8728693312, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728693152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728781920 := &system.Object{Context: ptr8728693152, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728439136 := &system.Number_rule{Object: ptr8728781920, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728422848 := &system.Map_rule{Object: ptr8728781808, Items: ptr8728439136, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728487712 := &system.Property{Object: ptr8728781696, Defaulter: false, Item: ptr8728422848, Optional: false}

	ptr8728573280 := &system.Type{Object: ptr8728781024, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8728487616, "b": ptr8728487664, "c": ptr8728487712}, Rule: (*system.Type)(nil)}

	ptr8728692992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728785952 := &system.Object{Context: ptr8728692992, Description: "Automatically created basic rule for photo", Id: "@photo", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728575152 := &system.Type{Object: ptr8728785952, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728699776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728787856 := &system.Object{Context: ptr8728699776, Description: "Automatically created basic rule for sibling", Id: "@sibling", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728650160 := &system.Type{Object: ptr8728787856, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728346112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728784720 := &system.Object{Context: ptr8728346112, Description: "", Id: "kid", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728344608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728784832 := &system.Object{Context: ptr8728344608, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728344576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728784944 := &system.Object{Context: ptr8728344576, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553344 := &system.String_rule{Object: ptr8728784944, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488912 := &system.Property{Object: ptr8728784832, Defaulter: false, Item: ptr8728553344, Optional: false}

	ptr8728345280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728785056 := &system.Object{Context: ptr8728345280, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728345248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728785168 := &system.Object{Context: ptr8728345248, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553984 := &system.String_rule{Object: ptr8728785168, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488960 := &system.Property{Object: ptr8728785056, Defaulter: false, Item: ptr8728553984, Optional: false}

	ptr8728345920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728785280 := &system.Object{Context: ptr8728345920, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728345888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728785392 := &system.Object{Context: ptr8728345888, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728864256 := &system.Bool_rule{Object: ptr8728785392, Default: system.Bool{Value: false, Exists: false}}

	ptr8728489056 := &system.Property{Object: ptr8728785280, Defaulter: false, Item: ptr8728864256, Optional: true}

	ptr8728574576 := &system.Type{Object: ptr8728784720, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"language": ptr8728488912, "level": ptr8728488960, "preferred": ptr8728489056}, Rule: (*system.Type)(nil)}

	ptr8728692480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728785616 := &system.Object{Context: ptr8728692480, Description: "", Id: "photo", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728692288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728785728 := &system.Object{Context: ptr8728692288, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728692256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728785840 := &system.Object{Context: ptr8728692256, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728554144 := &system.String_rule{Object: ptr8728785840, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728489488 := &system.Property{Object: ptr8728785728, Defaulter: false, Item: ptr8728554144, Optional: false}

	ptr8728575008 := &system.Type{Object: ptr8728785616, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/jsonselect:image", Package: "kego.io/jsonselect", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"url": ptr8728489488}, Rule: (*system.Type)(nil)}

	ptr8728343808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728784608 := &system.Object{Context: ptr8728343808, Description: "Automatically created basic rule for image", Id: "@image", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728574288 := &system.Type{Object: ptr8728784608, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728694368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728786064 := &system.Object{Context: ptr8728694368, Description: "", Id: "polykids", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728694208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728786176 := &system.Object{Context: ptr8728694208, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728694080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728786288 := &system.Object{Context: ptr8728694080, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728427456 := &system.Array_rule{Object: ptr8728786288, Items: map[string]interface{}{"_imports": map[string]string{}, "type": "@kid", "_path": "kego.io/jsonselect"}, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728489872 := &system.Property{Object: ptr8728786176, Defaulter: false, Item: ptr8728427456, Optional: false}

	ptr8728656064 := &system.Type{Object: ptr8728786064, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8728489872}, Rule: (*system.Type)(nil)}

	ptr8728695296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728782592 := &system.Object{Context: ptr8728695296, Description: "Automatically created basic rule for collision", Id: "@collision", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728573856 := &system.Type{Object: ptr8728782592, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728690944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728780912 := &system.Object{Context: ptr8728690944, Description: "Automatically created basic rule for basic", Id: "@basic", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728572848 := &system.Type{Object: ptr8728780912, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/jsonselect:@basic", ptr8728572848)

	system.RegisterType("kego.io/jsonselect:@c", ptr8728573424)

	system.RegisterType("kego.io/jsonselect:@collision", ptr8728573856)

	system.RegisterType("kego.io/jsonselect:@expr", ptr8728572992)

	system.RegisterType("kego.io/jsonselect:@image", ptr8728574288)

	system.RegisterType("kego.io/jsonselect:@kid", ptr8728574720)

	system.RegisterType("kego.io/jsonselect:@photo", ptr8728575152)

	system.RegisterType("kego.io/jsonselect:@polykids", ptr8728649728)

	system.RegisterType("kego.io/jsonselect:@sibling", ptr8728650160)

	system.RegisterType("kego.io/jsonselect:@typed", ptr8728650592)

	system.RegisterType("kego.io/jsonselect:basic", ptr8728575584)

	system.RegisterType("kego.io/jsonselect:c", ptr8728573280)

	system.RegisterType("kego.io/jsonselect:collision", ptr8728573712)

	system.RegisterType("kego.io/jsonselect:expr", ptr8728574144)

	system.RegisterType("kego.io/jsonselect:image", ptr8728575440)

	system.RegisterType("kego.io/jsonselect:kid", ptr8728574576)

	system.RegisterType("kego.io/jsonselect:photo", ptr8728575008)

	system.RegisterType("kego.io/jsonselect:polykids", ptr8728656064)

	system.RegisterType("kego.io/jsonselect:sibling", ptr8728650016)

	system.RegisterType("kego.io/jsonselect:typed", ptr8728650448)

}
