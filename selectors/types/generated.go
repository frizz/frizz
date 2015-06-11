package types

import (
	"kego.io/system"

	_ "kego.io/system/types"
)

func init() {

	ptr8728774624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728622560 := &system.Object{Context: ptr8728774624, Description: "", Id: "typed", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728756000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728622672 := &system.Object{Context: ptr8728756000, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728755968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728622784 := &system.Object{Context: ptr8728755968, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728755936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728622896 := &system.Object{Context: ptr8728755936, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728554144 := &system.String_rule{Object: ptr8728622896, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728423872 := &system.Map_rule{Object: ptr8728622784, Items: ptr8728554144, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728490400 := &system.Property{Object: ptr8728622672, Defaulter: false, Item: ptr8728423872, Optional: false}

	ptr8728772896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623008 := &system.Object{Context: ptr8728772896, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728772864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623120 := &system.Object{Context: ptr8728772864, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728554304 := &system.String_rule{Object: ptr8728623120, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728490496 := &system.Property{Object: ptr8728623008, Defaulter: false, Item: ptr8728554304, Optional: false}

	ptr8728773376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623232 := &system.Object{Context: ptr8728773376, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728773344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623344 := &system.Object{Context: ptr8728773344, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728424000 := &system.Map_rule{Object: ptr8728623344, Items: map[string]interface{}{"type": "@kid", "_path": "kego.io/jsonselect", "_imports": map[string]string{}}, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728490544 := &system.Property{Object: ptr8728623232, Defaulter: false, Item: ptr8728424000, Optional: false}

	ptr8728772992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623456 := &system.Object{Context: ptr8728772992, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728772800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623568 := &system.Object{Context: ptr8728772800, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728772768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623680 := &system.Object{Context: ptr8728772768, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728554464 := &system.String_rule{Object: ptr8728623680, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728424128 := &system.Array_rule{Object: ptr8728623568, Items: ptr8728554464, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728490688 := &system.Property{Object: ptr8728623456, Defaulter: false, Item: ptr8728424128, Optional: false}

	ptr8728773920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623792 := &system.Object{Context: ptr8728773920, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728773824 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728623904 := &system.Object{Context: ptr8728773824, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728773792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728624016 := &system.Object{Context: ptr8728773792, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728554624 := &system.String_rule{Object: ptr8728624016, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728422400 := &system.Array_rule{Object: ptr8728623904, Items: ptr8728554624, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728487904 := &system.Property{Object: ptr8728623792, Defaulter: false, Item: ptr8728422400, Optional: false}

	ptr8728774496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728624128 := &system.Object{Context: ptr8728774496, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728774336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728624240 := &system.Object{Context: ptr8728774336, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728438336 := &system.Number_rule{Object: ptr8728624240, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728489104 := &system.Property{Object: ptr8728624128, Defaulter: false, Item: ptr8728438336, Optional: false}

	ptr8728655920 := &system.Type{Object: ptr8728622560, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"name": ptr8728490400, "favoriteColor": ptr8728490496, "kids": ptr8728490544, "seatingPreference": ptr8728490688, "drinkPreference": ptr8728487904, "weight": ptr8728489104}, Rule: (*system.Type)(nil)}

	ptr8728752640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728718960 := &system.Object{Context: ptr8728752640, Description: "Automatically created basic rule for collision", Id: "@collision", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728573856 := &system.Type{Object: ptr8728718960, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728345696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728721648 := &system.Object{Context: ptr8728345696, Description: "Automatically created basic rule for kid", Id: "@kid", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728574288 := &system.Type{Object: ptr8728721648, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728754784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728622448 := &system.Object{Context: ptr8728754784, Description: "Automatically created basic rule for sibling", Id: "@sibling", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728575584 := &system.Type{Object: ptr8728622448, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728749408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728722096 := &system.Object{Context: ptr8728749408, Description: "Automatically created basic rule for polykids", Id: "@polykids", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728574720 := &system.Type{Object: ptr8728722096, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728752224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728718400 := &system.Object{Context: ptr8728752224, Description: "", Id: "collision", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728752096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728718624 := &system.Object{Context: ptr8728752096, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728752064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728718736 := &system.Object{Context: ptr8728752064, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728752032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728718848 := &system.Object{Context: ptr8728752032, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728551424 := &system.String_rule{Object: ptr8728718848, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728423168 := &system.Map_rule{Object: ptr8728718736, Items: ptr8728551424, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728488000 := &system.Property{Object: ptr8728718624, Defaulter: false, Item: ptr8728423168, Optional: false}

	ptr8728573712 := &system.Type{Object: ptr8728718400, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"number": ptr8728488000}, Rule: (*system.Type)(nil)}

	ptr8728750816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728715488 := &system.Object{Context: ptr8728750816, Description: "", Id: "c", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728749344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728715712 := &system.Object{Context: ptr8728749344, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728749184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728715824 := &system.Object{Context: ptr8728749184, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728438976 := &system.Number_rule{Object: ptr8728715824, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728487616 := &system.Property{Object: ptr8728715712, Defaulter: false, Item: ptr8728438976, Optional: false}

	ptr8728749920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728716048 := &system.Object{Context: ptr8728749920, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728749760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728716160 := &system.Object{Context: ptr8728749760, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728439056 := &system.Number_rule{Object: ptr8728716160, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728487664 := &system.Property{Object: ptr8728716048, Defaulter: false, Item: ptr8728439056, Optional: false}

	ptr8728750688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728716496 := &system.Object{Context: ptr8728750688, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728750656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728716832 := &system.Object{Context: ptr8728750656, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728750496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728717056 := &system.Object{Context: ptr8728750496, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728439136 := &system.Number_rule{Object: ptr8728717056, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728422848 := &system.Map_rule{Object: ptr8728716832, Items: ptr8728439136, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728487712 := &system.Property{Object: ptr8728716496, Defaulter: false, Item: ptr8728422848, Optional: false}

	ptr8728573280 := &system.Type{Object: ptr8728715488, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8728487616, "b": ptr8728487664, "c": ptr8728487712}, Rule: (*system.Type)(nil)}

	ptr8728751232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728718288 := &system.Object{Context: ptr8728751232, Description: "Automatically created basic rule for c", Id: "@c", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728573424 := &system.Type{Object: ptr8728718288, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728748864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728715264 := &system.Object{Context: ptr8728748864, Description: "", Id: "basic", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728343296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728715600 := &system.Object{Context: ptr8728343296, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728343264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728715936 := &system.Object{Context: ptr8728343264, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728343232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728716272 := &system.Object{Context: ptr8728343232, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728552864 := &system.String_rule{Object: ptr8728716272, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728422272 := &system.Map_rule{Object: ptr8728715936, Items: ptr8728552864, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728487328 := &system.Property{Object: ptr8728715600, Defaulter: false, Item: ptr8728422272, Optional: false}

	ptr8728343744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728716384 := &system.Object{Context: ptr8728343744, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728343712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728716608 := &system.Object{Context: ptr8728343712, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553184 := &system.String_rule{Object: ptr8728716608, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488144 := &system.Property{Object: ptr8728716384, Defaulter: false, Item: ptr8728553184, Optional: false}

	ptr8728344832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728716720 := &system.Object{Context: ptr8728344832, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728344736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728717168 := &system.Object{Context: ptr8728344736, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728344704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728717280 := &system.Object{Context: ptr8728344704, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728344672 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728717392 := &system.Object{Context: ptr8728344672, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553504 := &system.String_rule{Object: ptr8728717392, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728422592 := &system.Map_rule{Object: ptr8728717280, Items: ptr8728553504, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728422464 := &system.Array_rule{Object: ptr8728717168, Items: ptr8728422592, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728488192 := &system.Property{Object: ptr8728716720, Defaulter: false, Item: ptr8728422464, Optional: false}

	ptr8728345536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728717504 := &system.Object{Context: ptr8728345536, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728345440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728717616 := &system.Object{Context: ptr8728345440, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728345408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728717728 := &system.Object{Context: ptr8728345408, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553664 := &system.String_rule{Object: ptr8728717728, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728422656 := &system.Array_rule{Object: ptr8728717616, Items: ptr8728553664, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728488768 := &system.Property{Object: ptr8728717504, Defaulter: false, Item: ptr8728422656, Optional: false}

	ptr8728346400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728717840 := &system.Object{Context: ptr8728346400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728346208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728717952 := &system.Object{Context: ptr8728346208, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728346176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728718064 := &system.Object{Context: ptr8728346176, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553824 := &system.String_rule{Object: ptr8728718064, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728422784 := &system.Array_rule{Object: ptr8728717952, Items: ptr8728553824, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728488912 := &system.Property{Object: ptr8728717840, Defaulter: false, Item: ptr8728422784, Optional: false}

	ptr8728748736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728718176 := &system.Object{Context: ptr8728748736, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728748576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728718512 := &system.Object{Context: ptr8728748576, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728438416 := &system.Number_rule{Object: ptr8728718512, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728489056 := &system.Property{Object: ptr8728718176, Defaulter: false, Item: ptr8728438416, Optional: false}

	ptr8728575440 := &system.Type{Object: ptr8728715264, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"name": ptr8728487328, "favoriteColor": ptr8728488144, "languagesSpoken": ptr8728488192, "seatingPreference": ptr8728488768, "drinkPreference": ptr8728488912, "weight": ptr8728489056}, Rule: (*system.Type)(nil)}

	ptr8728345088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728720528 := &system.Object{Context: ptr8728345088, Description: "", Id: "kid", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728343648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728720976 := &system.Object{Context: ptr8728343648, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728343616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728721088 := &system.Object{Context: ptr8728343616, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553344 := &system.String_rule{Object: ptr8728721088, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488720 := &system.Property{Object: ptr8728720976, Defaulter: false, Item: ptr8728553344, Optional: false}

	ptr8728344256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728721200 := &system.Object{Context: ptr8728344256, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728344224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728721312 := &system.Object{Context: ptr8728344224, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553984 := &system.String_rule{Object: ptr8728721312, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488816 := &system.Property{Object: ptr8728721200, Defaulter: false, Item: ptr8728553984, Optional: false}

	ptr8728344928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728721424 := &system.Object{Context: ptr8728344928, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728344896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728721536 := &system.Object{Context: ptr8728344896, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728863728 := &system.Bool_rule{Object: ptr8728721536, Default: system.Bool{Value: false, Exists: false}}

	ptr8728488864 := &system.Property{Object: ptr8728721424, Defaulter: false, Item: ptr8728863728, Optional: true}

	ptr8728575296 := &system.Type{Object: ptr8728720528, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"language": ptr8728488720, "level": ptr8728488816, "preferred": ptr8728488864}, Rule: (*system.Type)(nil)}

	ptr8728775040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728624352 := &system.Object{Context: ptr8728775040, Description: "Automatically created basic rule for typed", Id: "@typed", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728649728 := &system.Type{Object: ptr8728624352, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728754240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728722208 := &system.Object{Context: ptr8728754240, Description: "", Id: "sibling", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728750400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728722320 := &system.Object{Context: ptr8728750400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728750208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728722432 := &system.Object{Context: ptr8728750208, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728440336 := &system.Number_rule{Object: ptr8728722432, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728489824 := &system.Property{Object: ptr8728722320, Defaulter: false, Item: ptr8728440336, Optional: false}

	ptr8728751392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728722544 := &system.Object{Context: ptr8728751392, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728751008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728722656 := &system.Object{Context: ptr8728751008, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728440416 := &system.Number_rule{Object: ptr8728722656, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728489872 := &system.Property{Object: ptr8728722544, Defaulter: false, Item: ptr8728440416, Optional: false}

	ptr8728751744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728722768 := &system.Object{Context: ptr8728751744, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728489920 := &system.Property{Object: ptr8728722768, Defaulter: false, Item: map[string]interface{}{"type": "@c", "_path": "kego.io/jsonselect", "_imports": map[string]string{}}, Optional: false}

	ptr8728753088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728722880 := &system.Object{Context: ptr8728753088, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728753024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728722992 := &system.Object{Context: ptr8728753024, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728752832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728723104 := &system.Object{Context: ptr8728752832, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728440496 := &system.Number_rule{Object: ptr8728723104, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728423488 := &system.Map_rule{Object: ptr8728722992, Items: ptr8728440496, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728490016 := &system.Property{Object: ptr8728722880, Defaulter: false, Item: ptr8728423488, Optional: false}

	ptr8728754080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728723216 := &system.Object{Context: ptr8728754080, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728754048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728723328 := &system.Object{Context: ptr8728754048, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728753792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728622336 := &system.Object{Context: ptr8728753792, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728440576 := &system.Number_rule{Object: ptr8728622336, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728423616 := &system.Map_rule{Object: ptr8728723328, Items: ptr8728440576, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728490112 := &system.Property{Object: ptr8728723216, Defaulter: false, Item: ptr8728423616, Optional: false}

	ptr8728575008 := &system.Type{Object: ptr8728722208, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8728489824, "b": ptr8728489872, "c": ptr8728489920, "d": ptr8728490016, "e": ptr8728490112}, Rule: (*system.Type)(nil)}

	ptr8728748288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728715376 := &system.Object{Context: ptr8728748288, Description: "Automatically created basic rule for basic", Id: "@basic", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728572848 := &system.Type{Object: ptr8728715376, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728748768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728721760 := &system.Object{Context: ptr8728748768, Description: "", Id: "polykids", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728748512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728721872 := &system.Object{Context: ptr8728748512, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728748416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728721984 := &system.Object{Context: ptr8728748416, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728423232 := &system.Array_rule{Object: ptr8728721984, Items: map[string]interface{}{"type": "@kid", "_path": "kego.io/jsonselect", "_imports": map[string]string{}}, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728489248 := &system.Property{Object: ptr8728721872, Defaulter: false, Item: ptr8728423232, Optional: false}

	ptr8728574576 := &system.Type{Object: ptr8728721760, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8728489248}, Rule: (*system.Type)(nil)}

	ptr8728342624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728719072 := &system.Object{Context: ptr8728342624, Description: "", Id: "expr", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728342464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728720752 := &system.Object{Context: ptr8728342464, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728342432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728720864 := &system.Object{Context: ptr8728342432, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728264512 := &system.Bool_rule{Object: ptr8728720864, Default: system.Bool{Value: false, Exists: false}}

	ptr8728489152 := &system.Property{Object: ptr8728720752, Defaulter: false, Item: ptr8728264512, Optional: false}

	ptr8728753440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728719184 := &system.Object{Context: ptr8728753440, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728753280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728719296 := &system.Object{Context: ptr8728753280, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728440016 := &system.Number_rule{Object: ptr8728719296, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728488384 := &system.Property{Object: ptr8728719184, Defaulter: false, Item: ptr8728440016, Optional: false}

	ptr8728754016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728719408 := &system.Object{Context: ptr8728754016, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728753856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728719520 := &system.Object{Context: ptr8728753856, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728440096 := &system.Number_rule{Object: ptr8728719520, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728488432 := &system.Property{Object: ptr8728719408, Defaulter: false, Item: ptr8728440096, Optional: false}

	ptr8728754464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728719632 := &system.Object{Context: ptr8728754464, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728754432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728719744 := &system.Object{Context: ptr8728754432, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728552544 := &system.String_rule{Object: ptr8728719744, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488480 := &system.Property{Object: ptr8728719632, Defaulter: false, Item: ptr8728552544, Optional: false}

	ptr8728754912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728719856 := &system.Object{Context: ptr8728754912, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728754880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728719968 := &system.Object{Context: ptr8728754880, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728552704 := &system.String_rule{Object: ptr8728719968, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488576 := &system.Property{Object: ptr8728719856, Defaulter: false, Item: ptr8728552704, Optional: false}

	ptr8728755360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728720080 := &system.Object{Context: ptr8728755360, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728755328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728720192 := &system.Object{Context: ptr8728755328, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728553024 := &system.String_rule{Object: ptr8728720192, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728488624 := &system.Property{Object: ptr8728720080, Defaulter: false, Item: ptr8728553024, Optional: false}

	ptr8728755872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728720304 := &system.Object{Context: ptr8728755872, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728755840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728720640 := &system.Object{Context: ptr8728755840, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728264064 := &system.Bool_rule{Object: ptr8728720640, Default: system.Bool{Value: false, Exists: false}}

	ptr8728488672 := &system.Property{Object: ptr8728720304, Defaulter: false, Item: ptr8728264064, Optional: false}

	ptr8728574144 := &system.Type{Object: ptr8728719072, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"false": ptr8728489152, "int": ptr8728488384, "float": ptr8728488432, "string": ptr8728488480, "string2": ptr8728488576, "null": ptr8728488624, "true": ptr8728488672}, Rule: (*system.Type)(nil)}

	ptr8728342656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728720416 := &system.Object{Context: ptr8728342656, Description: "Automatically created basic rule for expr", Id: "@expr", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728572992 := &system.Type{Object: ptr8728720416, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/jsonselect:@basic", ptr8728572848)

	system.RegisterType("kego.io/jsonselect:@c", ptr8728573424)

	system.RegisterType("kego.io/jsonselect:@collision", ptr8728573856)

	system.RegisterType("kego.io/jsonselect:@expr", ptr8728572992)

	system.RegisterType("kego.io/jsonselect:@kid", ptr8728574288)

	system.RegisterType("kego.io/jsonselect:@polykids", ptr8728574720)

	system.RegisterType("kego.io/jsonselect:@sibling", ptr8728575584)

	system.RegisterType("kego.io/jsonselect:@typed", ptr8728649728)

	system.RegisterType("kego.io/jsonselect:basic", ptr8728575440)

	system.RegisterType("kego.io/jsonselect:c", ptr8728573280)

	system.RegisterType("kego.io/jsonselect:collision", ptr8728573712)

	system.RegisterType("kego.io/jsonselect:expr", ptr8728574144)

	system.RegisterType("kego.io/jsonselect:kid", ptr8728575296)

	system.RegisterType("kego.io/jsonselect:polykids", ptr8728574576)

	system.RegisterType("kego.io/jsonselect:sibling", ptr8728575008)

	system.RegisterType("kego.io/jsonselect:typed", ptr8728655920)

}
