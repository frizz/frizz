package types

import (
	"kego.io/system"

	_ "kego.io/system/types"

	"kego.io/selectors"
)

func init() {

	ptr8728862272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729412608 := &system.Base{Context: ptr8728862272, Description: "", Id: "polykids", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728862080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729412736 := &system.Base{Context: ptr8728862080, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728862048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729412864 := &system.Base{Context: ptr8728862048, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@kid", Package: "kego.io/selectors", Type: "@kid", Exists: true}}

	ptr8729601632 := &selectors.Kid_rule{Base: ptr8729412864, RuleBase: (*system.RuleBase)(nil)}

	ptr8729559760 := &system.Array_rule{Base: ptr8729412736, RuleBase: (*system.RuleBase)(nil), Items: ptr8729601632, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729392464 := &system.Type{Base: ptr8729412608, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"a": ptr8729559760}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729376992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729304960 := &system.Base{Context: ptr8729376992, Description: "Automatically created basic rule for image", Id: "@image", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729390528 := &system.Type{Base: ptr8729304960, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729513568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729413120 := &system.Base{Context: ptr8729513568, Description: "", Id: "sibling", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729513440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729413888 := &system.Base{Context: ptr8729513440, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729513216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729414016 := &system.Base{Context: ptr8729513216, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729052096 := &system.Number_rule{Base: ptr8729414016, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729200576 := &system.Map_rule{Base: ptr8729413888, RuleBase: (*system.RuleBase)(nil), Items: ptr8729052096, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729510496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729413248 := &system.Base{Context: ptr8729510496, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729051808 := &system.Number_rule{Base: ptr8729413248, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729511104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729413376 := &system.Base{Context: ptr8729511104, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729051904 := &system.Number_rule{Base: ptr8729413376, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729511584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729413504 := &system.Base{Context: ptr8729511584, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@c", Package: "kego.io/selectors", Type: "@c", Exists: true}}

	ptr8729602640 := &selectors.C_rule{Base: ptr8729413504, RuleBase: (*system.RuleBase)(nil)}

	ptr8729512576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729413632 := &system.Base{Context: ptr8729512576, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729512288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729413760 := &system.Base{Context: ptr8729512288, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729052000 := &system.Number_rule{Base: ptr8729413760, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729200512 := &system.Map_rule{Base: ptr8729413632, RuleBase: (*system.RuleBase)(nil), Items: ptr8729052000, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729392816 := &system.Type{Base: ptr8729413120, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"e": ptr8729200576, "a": ptr8729051808, "b": ptr8729051904, "c": ptr8729602640, "d": ptr8729200512}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729374592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729303168 := &system.Base{Context: ptr8729374592, Description: "", Id: "expr", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729374464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729304064 := &system.Base{Context: ptr8729374464, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729374336 := &system.Bool_rule{Base: ptr8729304064, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729372160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729303296 := &system.Base{Context: ptr8729372160, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729051616 := &system.Number_rule{Base: ptr8729303296, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729372608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729303424 := &system.Base{Context: ptr8729372608, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729051712 := &system.Number_rule{Base: ptr8729303424, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729373056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729303552 := &system.Base{Context: ptr8729373056, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729389296 := &system.String_rule{Base: ptr8729303552, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729373376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729303680 := &system.Base{Context: ptr8729373376, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729389472 := &system.String_rule{Base: ptr8729303680, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729373696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729303808 := &system.Base{Context: ptr8729373696, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729389648 := &system.String_rule{Base: ptr8729303808, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729374112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729303936 := &system.Base{Context: ptr8729374112, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729373984 := &system.Bool_rule{Base: ptr8729303936, RuleBase: (*system.RuleBase)(nil), Default: system.Bool{Value: false, Exists: false}}

	ptr8729389120 := &system.Type{Base: ptr8729303168, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"false": ptr8729374336, "int": ptr8729051616, "float": ptr8729051712, "string": ptr8729389296, "string2": ptr8729389472, "null": ptr8729389648, "true": ptr8729373984}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729517888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729302656 := &system.Base{Context: ptr8729517888, Description: "Automatically created basic rule for collision", Id: "@collision", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729388416 := &system.Type{Base: ptr8729302656, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728862688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729412992 := &system.Base{Context: ptr8728862688, Description: "Automatically created basic rule for polykids", Id: "@polykids", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729392640 := &system.Type{Base: ptr8729412992, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729517376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729302272 := &system.Base{Context: ptr8729517376, Description: "", Id: "collision", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729517248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729302400 := &system.Base{Context: ptr8729517248, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729517216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729302528 := &system.Base{Context: ptr8729517216, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729388240 := &system.String_rule{Base: ptr8729302528, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729199424 := &system.Map_rule{Base: ptr8729302400, RuleBase: (*system.RuleBase)(nil), Items: ptr8729388240, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729388064 := &system.Type{Base: ptr8729302272, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"number": ptr8729199424}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728861376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729412480 := &system.Base{Context: ptr8728861376, Description: "Automatically created basic rule for photo", Id: "@photo", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729392288 := &system.Type{Base: ptr8729412480, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729375872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729304320 := &system.Base{Context: ptr8729375872, Description: "This represents a gallery - it's just a list of images", Id: "gallery", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729375744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729304448 := &system.Base{Context: ptr8729375744, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729375712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729304576 := &system.Base{Context: ptr8729375712, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@image", Package: "kego.io/selectors", Type: "@image", Exists: true}}

	ptr8729024256 := &selectors.Image_rule{Base: ptr8729304576, RuleBase: (*system.RuleBase)(nil)}

	ptr8729200000 := &system.Map_rule{Base: ptr8729304448, RuleBase: (*system.RuleBase)(nil), Items: ptr8729024256, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729390000 := &system.Type{Base: ptr8729304320, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"images": ptr8729200000}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729376576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729304832 := &system.Base{Context: ptr8729376576, Description: "", Id: "image", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729390352 := &system.Type{Base: ptr8729304832, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729516480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729302144 := &system.Base{Context: ptr8729516480, Description: "Automatically created basic rule for c", Id: "@c", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729387888 := &system.Type{Base: ptr8729302144, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729378784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729411840 := &system.Base{Context: ptr8729378784, Description: "Automatically created basic rule for kid", Id: "@kid", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729391232 := &system.Type{Base: ptr8729411840, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8728860960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729411968 := &system.Base{Context: ptr8728860960, Description: "This represents an image, and contains path, server and protocol separately", Id: "photo", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728860032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729412096 := &system.Base{Context: ptr8728860032, Description: "The protocol for the url - e.g. http or https", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728860000 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729391760 := &system.String_rule{Base: ptr8729412096, RuleBase: ptr8728860000, Default: system.String{Value: "http", Exists: true}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728860416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729412224 := &system.Base{Context: ptr8728860416, Description: "The server for the url - e.g. www.google.com", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729391936 := &system.String_rule{Base: ptr8729412224, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728860832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729412352 := &system.Base{Context: ptr8728860832, Description: "The path for the url - e.g. /foo/bar.jpg", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729392112 := &system.String_rule{Base: ptr8729412352, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "^/.*$", Exists: true}}

	ptr8729391408 := &system.Type{Base: ptr8729411968, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"protocol": ptr8729391760, "server": ptr8729391936, "path": ptr8729392112}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/selectors:image", Package: "kego.io/selectors", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729513696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729299840 := &system.Base{Context: ptr8729513696, Description: "", Id: "basic", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729510976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729300224 := &system.Base{Context: ptr8729510976, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729173744 := &system.String_rule{Base: ptr8729300224, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729511744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729300352 := &system.Base{Context: ptr8729511744, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729511712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729300480 := &system.Base{Context: ptr8729511712, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729511680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729300608 := &system.Base{Context: ptr8729511680, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729387008 := &system.String_rule{Base: ptr8729300608, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729198784 := &system.Map_rule{Base: ptr8729300480, RuleBase: (*system.RuleBase)(nil), Items: ptr8729387008, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729559520 := &system.Array_rule{Base: ptr8729300352, RuleBase: (*system.RuleBase)(nil), Items: ptr8729198784, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729512384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729300736 := &system.Base{Context: ptr8729512384, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729512352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729300864 := &system.Base{Context: ptr8729512352, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729387184 := &system.String_rule{Base: ptr8729300864, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729559600 := &system.Array_rule{Base: ptr8729300736, RuleBase: (*system.RuleBase)(nil), Items: ptr8729387184, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729513024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729300992 := &system.Base{Context: ptr8729513024, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729512992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729301120 := &system.Base{Context: ptr8729512992, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729387360 := &system.String_rule{Base: ptr8729301120, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729559680 := &system.Array_rule{Base: ptr8729300992, RuleBase: (*system.RuleBase)(nil), Items: ptr8729387360, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729513408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729301248 := &system.Base{Context: ptr8729513408, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729051232 := &system.Number_rule{Base: ptr8729301248, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729510656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729299968 := &system.Base{Context: ptr8729510656, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729510624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729300096 := &system.Base{Context: ptr8729510624, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729173568 := &system.String_rule{Base: ptr8729300096, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729198720 := &system.Map_rule{Base: ptr8729299968, RuleBase: (*system.RuleBase)(nil), Items: ptr8729173568, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729173392 := &system.Type{Base: ptr8729299840, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"favoriteColor": ptr8729173744, "languagesSpoken": ptr8729559520, "seatingPreference": ptr8729559600, "drinkPreference": ptr8729559680, "weight": ptr8729051232, "name": ptr8729198720}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729516064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729301504 := &system.Base{Context: ptr8729516064, Description: "", Id: "c", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729515936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729301888 := &system.Base{Context: ptr8729515936, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729515776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729302016 := &system.Base{Context: ptr8729515776, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729051520 := &system.Number_rule{Base: ptr8729302016, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729199232 := &system.Map_rule{Base: ptr8729301888, RuleBase: (*system.RuleBase)(nil), Items: ptr8729051520, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729514688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729301632 := &system.Base{Context: ptr8729514688, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729051328 := &system.Number_rule{Base: ptr8729301632, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729515136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729301760 := &system.Base{Context: ptr8729515136, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729051424 := &system.Number_rule{Base: ptr8729301760, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729387712 := &system.Type{Base: ptr8729301504, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"c": ptr8729199232, "a": ptr8729051328, "b": ptr8729051424}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729514336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729414144 := &system.Base{Context: ptr8729514336, Description: "Automatically created basic rule for sibling", Id: "@sibling", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729392992 := &system.Type{Base: ptr8729414144, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729371136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729414272 := &system.Base{Context: ptr8729371136, Description: "", Id: "typed", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729516736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729415040 := &system.Base{Context: ptr8729516736, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@image", Package: "kego.io/selectors", Type: "@image", Exists: true}}

	ptr8729604672 := &selectors.Image_rule{Base: ptr8729415040, RuleBase: (*system.RuleBase)(nil)}

	ptr8729517472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729415168 := &system.Base{Context: ptr8729517472, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729517440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729415296 := &system.Base{Context: ptr8729517440, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729393696 := &system.String_rule{Base: ptr8729415296, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729560640 := &system.Array_rule{Base: ptr8729415168, RuleBase: (*system.RuleBase)(nil), Items: ptr8729393696, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729370784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729415424 := &system.Base{Context: ptr8729370784, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729052192 := &system.Number_rule{Base: ptr8729415424, RuleBase: (*system.RuleBase)(nil), Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729515264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729414400 := &system.Base{Context: ptr8729515264, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729515232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729414528 := &system.Base{Context: ptr8729515232, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729393344 := &system.String_rule{Base: ptr8729414528, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729200768 := &system.Map_rule{Base: ptr8729414400, RuleBase: (*system.RuleBase)(nil), Items: ptr8729393344, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729515616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729414656 := &system.Base{Context: ptr8729515616, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729393520 := &system.String_rule{Base: ptr8729414656, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729516256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729414784 := &system.Base{Context: ptr8729516256, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729516192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729414912 := &system.Base{Context: ptr8729516192, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/selectors:@kid", Package: "kego.io/selectors", Type: "@kid", Exists: true}}

	ptr8729604512 := &selectors.Kid_rule{Base: ptr8729414912, RuleBase: (*system.RuleBase)(nil)}

	ptr8729200832 := &system.Map_rule{Base: ptr8729414784, RuleBase: (*system.RuleBase)(nil), Items: ptr8729604512, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729393168 := &system.Type{Base: ptr8729414272, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"avatar": ptr8729604672, "drinkPreference": ptr8729560640, "weight": ptr8729052192, "name": ptr8729200768, "favoriteColor": ptr8729393520, "kids": ptr8729200832}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729371616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729303040 := &system.Base{Context: ptr8729371616, Description: "Automatically created basic rule for diagram", Id: "@diagram", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729388944 := &system.Type{Base: ptr8729303040, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729378368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729182208 := &system.Base{Context: ptr8729378368, Description: "", Id: "kid", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729377536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729183488 := &system.Base{Context: ptr8729377536, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729390880 := &system.String_rule{Base: ptr8729183488, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729377856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729411584 := &system.Base{Context: ptr8729377856, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729391056 := &system.String_rule{Base: ptr8729411584, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729378240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729411712 := &system.Base{Context: ptr8729378240, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729378208 := &system.RuleBase{Optional: true, Selector: ""}

	ptr8729378080 := &system.Bool_rule{Base: ptr8729411712, RuleBase: ptr8729378208, Default: system.Bool{Value: false, Exists: false}}

	ptr8729390704 := &system.Type{Base: ptr8729182208, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"language": ptr8729390880, "level": ptr8729391056, "preferred": ptr8729378080}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729371840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729415552 := &system.Base{Context: ptr8729371840, Description: "Automatically created basic rule for typed", Id: "@typed", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729393872 := &system.Type{Base: ptr8729415552, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729375008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729304192 := &system.Base{Context: ptr8729375008, Description: "Automatically created basic rule for expr", Id: "@expr", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729389824 := &system.Type{Base: ptr8729304192, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729371200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729302784 := &system.Base{Context: ptr8729371200, Description: "", Id: "diagram", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729371072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729302912 := &system.Base{Context: ptr8729371072, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729388768 := &system.String_rule{Base: ptr8729302912, RuleBase: (*system.RuleBase)(nil), Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729388592 := &system.Type{Base: ptr8729302784, Basic: false, Embed: []system.Reference(nil), Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Fields: map[string]system.Rule{"url": ptr8729388768}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/selectors:image", Package: "kego.io/selectors", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729376288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729304704 := &system.Base{Context: ptr8729376288, Description: "Automatically created basic rule for gallery", Id: "@gallery", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729390176 := &system.Type{Base: ptr8729304704, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	ptr8729514144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/selectors"}

	ptr8729301376 := &system.Base{Context: ptr8729514144, Description: "Automatically created basic rule for basic", Id: "@basic", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729387536 := &system.Type{Base: ptr8729301376, Basic: false, Embed: []system.Reference{system.Reference{Value: "kego.io/system:ruleBase", Package: "kego.io/system", Type: "ruleBase", Exists: true}}, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Fields: map[string]system.Rule(nil), Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/selectors:@basic", ptr8729387536)

	system.RegisterType("kego.io/selectors:@c", ptr8729387888)

	system.RegisterType("kego.io/selectors:@collision", ptr8729388416)

	system.RegisterType("kego.io/selectors:@diagram", ptr8729388944)

	system.RegisterType("kego.io/selectors:@expr", ptr8729389824)

	system.RegisterType("kego.io/selectors:@gallery", ptr8729390176)

	system.RegisterType("kego.io/selectors:@image", ptr8729390528)

	system.RegisterType("kego.io/selectors:@kid", ptr8729391232)

	system.RegisterType("kego.io/selectors:@photo", ptr8729392288)

	system.RegisterType("kego.io/selectors:@polykids", ptr8729392640)

	system.RegisterType("kego.io/selectors:@sibling", ptr8729392992)

	system.RegisterType("kego.io/selectors:@typed", ptr8729393872)

	system.RegisterType("kego.io/selectors:basic", ptr8729173392)

	system.RegisterType("kego.io/selectors:c", ptr8729387712)

	system.RegisterType("kego.io/selectors:collision", ptr8729388064)

	system.RegisterType("kego.io/selectors:diagram", ptr8729388592)

	system.RegisterType("kego.io/selectors:expr", ptr8729389120)

	system.RegisterType("kego.io/selectors:gallery", ptr8729390000)

	system.RegisterType("kego.io/selectors:image", ptr8729390352)

	system.RegisterType("kego.io/selectors:kid", ptr8729390704)

	system.RegisterType("kego.io/selectors:photo", ptr8729391408)

	system.RegisterType("kego.io/selectors:polykids", ptr8729392464)

	system.RegisterType("kego.io/selectors:sibling", ptr8729392816)

	system.RegisterType("kego.io/selectors:typed", ptr8729393168)

	var _ selectors.Nothing

}
