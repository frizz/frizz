package types

import (
	"kego.io/system"

	_ "kego.io/system/types"
)

func init() {

	ptr8728497440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728840608 := &system.Object{Context: ptr8728497440, Description: "", Id: "collision", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728497312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728840720 := &system.Object{Context: ptr8728497312, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728497280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728840832 := &system.Object{Context: ptr8728497280, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728497248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728840944 := &system.Object{Context: ptr8728497248, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728585792 := &system.String_rule{Object: ptr8728840944, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728554368 := &system.Map_rule{Object: ptr8728840832, Items: ptr8728585792, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728489008 := &system.Property{Object: ptr8728840720, Defaulter: false, Item: ptr8728554368, Optional: false}

	ptr8728707792 := &system.Type{Object: ptr8728840608, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"number": ptr8728489008}, Rule: (*system.Type)(nil)}

	ptr8728346272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728839488 := &system.Object{Context: ptr8728346272, Description: "Automatically created basic rule for basic", Id: "@basic", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728707072 := &system.Type{Object: ptr8728839488, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728498688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728842848 := &system.Object{Context: ptr8728498688, Description: "Automatically created basic rule for expr", Id: "@expr", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728707216 := &system.Type{Object: ptr8728842848, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728854944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728842288 := &system.Object{Context: ptr8728854944, Description: "Automatically created basic rule for sibling", Id: "@sibling", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728707648 := &system.Type{Object: ptr8728842288, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728500864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728843744 := &system.Object{Context: ptr8728500864, Description: "Automatically created basic rule for kid", Id: "@kid", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728708368 := &system.Type{Object: ptr8728843744, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728496448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728840496 := &system.Object{Context: ptr8728496448, Description: "Automatically created basic rule for c", Id: "@c", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728707504 := &system.Type{Object: ptr8728840496, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728500448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728842960 := &system.Object{Context: ptr8728500448, Description: "", Id: "kid", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728499424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728843072 := &system.Object{Context: ptr8728499424, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728499392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728843184 := &system.Object{Context: ptr8728499392, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728586752 := &system.String_rule{Object: ptr8728843184, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488672 := &system.Property{Object: ptr8728843072, Defaulter: false, Item: ptr8728586752, Optional: false}

	ptr8728499872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728843296 := &system.Object{Context: ptr8728499872, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728499840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728843408 := &system.Object{Context: ptr8728499840, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728586912 := &system.String_rule{Object: ptr8728843408, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488912 := &system.Property{Object: ptr8728843296, Defaulter: false, Item: ptr8728586912, Optional: false}

	ptr8728500320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728843520 := &system.Object{Context: ptr8728500320, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728500288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728843632 := &system.Object{Context: ptr8728500288, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728262912 := &system.Bool_rule{Object: ptr8728843632, Default: system.Bool{Value: false, Exists: false}}

	ptr8728489056 := &system.Property{Object: ptr8728843520, Defaulter: false, Item: ptr8728262912, Optional: true}

	ptr8728710528 := &system.Type{Object: ptr8728842960, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"language": ptr8728488672, "level": ptr8728488912, "preferred": ptr8728489056}, Rule: (*system.Type)(nil)}

	ptr8728502176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728844192 := &system.Object{Context: ptr8728502176, Description: "Automatically created basic rule for polykids", Id: "@polykids", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728708800 := &system.Type{Object: ptr8728844192, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728855200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728844304 := &system.Object{Context: ptr8728855200, Description: "", Id: "sibling", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728343520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728844416 := &system.Object{Context: ptr8728343520, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728343296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728844528 := &system.Object{Context: ptr8728343296, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728718624 := &system.Number_rule{Object: ptr8728844528, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728489968 := &system.Property{Object: ptr8728844416, Defaulter: false, Item: ptr8728718624, Optional: false}

	ptr8728344672 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728844640 := &system.Object{Context: ptr8728344672, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728344320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728844752 := &system.Object{Context: ptr8728344320, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728717024 := &system.Number_rule{Object: ptr8728844752, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728490016 := &system.Property{Object: ptr8728844640, Defaulter: false, Item: ptr8728717024, Optional: false}

	ptr8728345024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728844864 := &system.Object{Context: ptr8728345024, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728490112 := &system.Property{Object: ptr8728844864, Defaulter: false, Item: map[string]interface{}{"type": "@c", "_path": "kego.io/jsonselect", "_imports": map[string]string{}}, Optional: false}

	ptr8728346112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728844976 := &system.Object{Context: ptr8728346112, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728346080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728845088 := &system.Object{Context: ptr8728346080, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728345792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728845200 := &system.Object{Context: ptr8728345792, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728717104 := &system.Number_rule{Object: ptr8728845200, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728554560 := &system.Map_rule{Object: ptr8728845088, Items: ptr8728717104, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728490256 := &system.Property{Object: ptr8728844976, Defaulter: false, Item: ptr8728554560, Optional: false}

	ptr8728855072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728845312 := &system.Object{Context: ptr8728855072, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728855040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728845424 := &system.Object{Context: ptr8728855040, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728854880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728845536 := &system.Object{Context: ptr8728854880, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728717184 := &system.Number_rule{Object: ptr8728845536, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728554688 := &system.Map_rule{Object: ptr8728845424, Items: ptr8728717184, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728490352 := &system.Property{Object: ptr8728845312, Defaulter: false, Item: ptr8728554688, Optional: false}

	ptr8728709088 := &system.Type{Object: ptr8728844304, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8728489968, "b": ptr8728490016, "c": ptr8728490112, "d": ptr8728490256, "e": ptr8728490352}, Rule: (*system.Type)(nil)}

	ptr8728345760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728656224 := &system.Object{Context: ptr8728345760, Description: "", Id: "basic", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728345824 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728657344 := &system.Object{Context: ptr8728345824, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728344448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728657456 := &system.Object{Context: ptr8728344448, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728586272 := &system.String_rule{Object: ptr8728657456, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728489632 := &system.Property{Object: ptr8728657344, Defaulter: false, Item: ptr8728586272, Optional: false}

	ptr8728342848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728657568 := &system.Object{Context: ptr8728342848, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728342752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728281200 := &system.Object{Context: ptr8728342752, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728342720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728838144 := &system.Object{Context: ptr8728342720, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728342688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728838256 := &system.Object{Context: ptr8728342688, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728584192 := &system.String_rule{Object: ptr8728838256, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728554112 := &system.Map_rule{Object: ptr8728838144, Items: ptr8728584192, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728553984 := &system.Array_rule{Object: ptr8728281200, Items: ptr8728554112, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728489728 := &system.Property{Object: ptr8728657568, Defaulter: false, Item: ptr8728553984, Optional: false}

	ptr8728343776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728838368 := &system.Object{Context: ptr8728343776, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728343648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728838480 := &system.Object{Context: ptr8728343648, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728343616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728838592 := &system.Object{Context: ptr8728343616, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728585312 := &system.String_rule{Object: ptr8728838592, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728553344 := &system.Array_rule{Object: ptr8728838480, Items: ptr8728585312, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728487376 := &system.Property{Object: ptr8728838368, Defaulter: false, Item: ptr8728553344, Optional: false}

	ptr8728344736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728838704 := &system.Object{Context: ptr8728344736, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728344640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728838816 := &system.Object{Context: ptr8728344640, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728344608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728838928 := &system.Object{Context: ptr8728344608, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728585472 := &system.String_rule{Object: ptr8728838928, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728553600 := &system.Array_rule{Object: ptr8728838816, Items: ptr8728585472, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728487472 := &system.Property{Object: ptr8728838704, Defaulter: false, Item: ptr8728553600, Optional: false}

	ptr8728345632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728839040 := &system.Object{Context: ptr8728345632, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728345472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728839376 := &system.Object{Context: ptr8728345472, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728715344 := &system.Number_rule{Object: ptr8728839376, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728487616 := &system.Property{Object: ptr8728839040, Defaulter: false, Item: ptr8728715344, Optional: false}

	ptr8728344032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728656560 := &system.Object{Context: ptr8728344032, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728344000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728656896 := &system.Object{Context: ptr8728344000, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728343968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728657232 := &system.Object{Context: ptr8728343968, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728585952 := &system.String_rule{Object: ptr8728657232, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728553856 := &system.Map_rule{Object: ptr8728656896, Items: ptr8728585952, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728488864 := &system.Property{Object: ptr8728656560, Defaulter: false, Item: ptr8728553856, Optional: false}

	ptr8728710672 := &system.Type{Object: ptr8728656224, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"favoriteColor": ptr8728489632, "languagesSpoken": ptr8728489728, "seatingPreference": ptr8728487376, "drinkPreference": ptr8728487472, "weight": ptr8728487616, "name": ptr8728488864}, Rule: (*system.Type)(nil)}

	ptr8728497856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728841056 := &system.Object{Context: ptr8728497856, Description: "Automatically created basic rule for collision", Id: "@collision", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728707936 := &system.Type{Object: ptr8728841056, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728498112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728841168 := &system.Object{Context: ptr8728498112, Description: "", Id: "expr", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728495040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728839152 := &system.Object{Context: ptr8728495040, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728495008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728839264 := &system.Object{Context: ptr8728495008, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728585632 := &system.String_rule{Object: ptr8728839264, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728487328 := &system.Property{Object: ptr8728839152, Defaulter: false, Item: ptr8728585632, Optional: false}

	ptr8728495584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728841728 := &system.Object{Context: ptr8728495584, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728495552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728841840 := &system.Object{Context: ptr8728495552, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728586112 := &system.String_rule{Object: ptr8728841840, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728487424 := &system.Property{Object: ptr8728841728, Defaulter: false, Item: ptr8728586112, Optional: false}

	ptr8728496224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728841952 := &system.Object{Context: ptr8728496224, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728496160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728842064 := &system.Object{Context: ptr8728496160, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728586432 := &system.String_rule{Object: ptr8728842064, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728487520 := &system.Property{Object: ptr8728841952, Defaulter: false, Item: ptr8728586432, Optional: false}

	ptr8728497024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728842176 := &system.Object{Context: ptr8728497024, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728496992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728842512 := &system.Object{Context: ptr8728496992, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728259792 := &system.Bool_rule{Object: ptr8728842512, Default: system.Bool{Value: false, Exists: false}}

	ptr8728487568 := &system.Property{Object: ptr8728842176, Defaulter: false, Item: ptr8728259792, Optional: false}

	ptr8728497728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728842624 := &system.Object{Context: ptr8728497728, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728497696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728842736 := &system.Object{Context: ptr8728497696, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728261552 := &system.Bool_rule{Object: ptr8728842736, Default: system.Bool{Value: false, Exists: false}}

	ptr8728487904 := &system.Property{Object: ptr8728842624, Defaulter: false, Item: ptr8728261552, Optional: false}

	ptr8728498656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728841280 := &system.Object{Context: ptr8728498656, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728498496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728841392 := &system.Object{Context: ptr8728498496, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728716864 := &system.Number_rule{Object: ptr8728841392, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728489296 := &system.Property{Object: ptr8728841280, Defaulter: false, Item: ptr8728716864, Optional: false}

	ptr8728494464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728841504 := &system.Object{Context: ptr8728494464, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728494240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728841616 := &system.Object{Context: ptr8728494240, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728716944 := &system.Number_rule{Object: ptr8728841616, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728489344 := &system.Property{Object: ptr8728841504, Defaulter: false, Item: ptr8728716944, Optional: false}

	ptr8728708224 := &system.Type{Object: ptr8728841168, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"string": ptr8728487328, "string2": ptr8728487424, "null": ptr8728487520, "true": ptr8728487568, "false": ptr8728487904, "int": ptr8728489296, "float": ptr8728489344}, Rule: (*system.Type)(nil)}

	ptr8728496032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728839600 := &system.Object{Context: ptr8728496032, Description: "", Id: "c", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728494560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728839712 := &system.Object{Context: ptr8728494560, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728494400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728839824 := &system.Object{Context: ptr8728494400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728715824 := &system.Number_rule{Object: ptr8728839824, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728488528 := &system.Property{Object: ptr8728839712, Defaulter: false, Item: ptr8728715824, Optional: false}

	ptr8728495136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728839936 := &system.Object{Context: ptr8728495136, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728494976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728840048 := &system.Object{Context: ptr8728494976, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728715904 := &system.Number_rule{Object: ptr8728840048, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728488576 := &system.Property{Object: ptr8728839936, Defaulter: false, Item: ptr8728715904, Optional: false}

	ptr8728495904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728840160 := &system.Object{Context: ptr8728495904, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728495872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728840272 := &system.Object{Context: ptr8728495872, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728495712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728840384 := &system.Object{Context: ptr8728495712, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728715984 := &system.Number_rule{Object: ptr8728840384, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728557888 := &system.Map_rule{Object: ptr8728840272, Items: ptr8728715984, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728488624 := &system.Property{Object: ptr8728840160, Defaulter: false, Item: ptr8728557888, Optional: false}

	ptr8728707360 := &system.Type{Object: ptr8728839600, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8728488528, "b": ptr8728488576, "c": ptr8728488624}, Rule: (*system.Type)(nil)}

	ptr8728501760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728843856 := &system.Object{Context: ptr8728501760, Description: "", Id: "polykids", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728501632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728843968 := &system.Object{Context: ptr8728501632, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728501536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728844080 := &system.Object{Context: ptr8728501536, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728554048 := &system.Array_rule{Object: ptr8728844080, Items: map[string]interface{}{"_path": "kego.io/jsonselect", "_imports": map[string]string{}, "type": "@kid"}, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728489536 := &system.Property{Object: ptr8728843968, Defaulter: false, Item: ptr8728554048, Optional: false}

	ptr8728708656 := &system.Type{Object: ptr8728843856, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8728489536}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/jsonselect:@basic", ptr8728707072)

	system.RegisterType("kego.io/jsonselect:@c", ptr8728707504)

	system.RegisterType("kego.io/jsonselect:@collision", ptr8728707936)

	system.RegisterType("kego.io/jsonselect:@expr", ptr8728707216)

	system.RegisterType("kego.io/jsonselect:@kid", ptr8728708368)

	system.RegisterType("kego.io/jsonselect:@polykids", ptr8728708800)

	system.RegisterType("kego.io/jsonselect:@sibling", ptr8728707648)

	system.RegisterType("kego.io/jsonselect:basic", ptr8728710672)

	system.RegisterType("kego.io/jsonselect:c", ptr8728707360)

	system.RegisterType("kego.io/jsonselect:collision", ptr8728707792)

	system.RegisterType("kego.io/jsonselect:expr", ptr8728708224)

	system.RegisterType("kego.io/jsonselect:kid", ptr8728710528)

	system.RegisterType("kego.io/jsonselect:polykids", ptr8728708656)

	system.RegisterType("kego.io/jsonselect:sibling", ptr8728709088)

}
