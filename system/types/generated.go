package types

import (
	system "kego.io/system"
)

func init() {
	ptr8728566528 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Description: "All rules should embed this type.", Rules: []system.Rule(nil)}
	ptr8728566656 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this rule is a field, this specifies that the field is optional", Rules: []system.Rule(nil)}
	ptr8728925888 := &system.RuleBase{Optional: true}
	ptr8728990000 := &system.JsonBool_rule{Base: ptr8728566656, RuleBase: ptr8728925888}
	ptr8728566784 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Json selector defining what nodes this rule should be applied to.", Rules: []system.Rule(nil)}
	ptr8728926304 := &system.RuleBase{Optional: true}
	ptr8728990176 := &system.JsonString_rule{Base: ptr8728566784, RuleBase: ptr8728926304}
	ptr8728791680 := &system.Type{Base: ptr8728566528, Rule: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8728990000, "selector": ptr8728990176}}
	ptr8728566912 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Description: "Automatically created basic rule for ruleBase", Rules: []system.Rule(nil)}
	ptr8728791904 := &system.Type{Base: ptr8728566912, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728561664 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Description: "This is the native json array data type", Rules: []system.Rule(nil)}
	ptr8728561792 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Description: "Restriction rules for arrays", Rules: []system.Rule(nil)}
	ptr8728561920 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Description: "This is a rule object, defining the type and restrictions on the value of the items", Rules: []system.Rule(nil)}
	ptr8728988144 := &system.Rule_rule{Base: ptr8728561920, RuleBase: (*system.RuleBase)(nil)}
	ptr8728562048 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the minimum number of items allowed in the array", Rules: []system.Rule(nil)}
	ptr8728922624 := &system.RuleBase{Optional: true}
	ptr8728854768 := &system.Int_rule{Base: ptr8728562048, RuleBase: ptr8728922624, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728562176 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the maximum number of items allowed in the array", Rules: []system.Rule(nil)}
	ptr8728922880 := &system.RuleBase{Optional: true}
	ptr8728854848 := &system.Int_rule{Base: ptr8728562176, RuleBase: ptr8728922880, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728562304 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, each item must be unique", Rules: []system.Rule(nil)}
	ptr8728923168 := &system.RuleBase{Optional: true}
	ptr8728988576 := &system.JsonBool_rule{Base: ptr8728562304, RuleBase: ptr8728923168}
	ptr8728789440 := &system.Type{Base: ptr8728561792, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8728988144, "minItems": ptr8728854768, "maxItems": ptr8728854848, "uniqueItems": ptr8728988576}}
	ptr8728789328 := &system.Type{Base: ptr8728561664, Rule: ptr8728789440, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728564224 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Description: "This is the int data type", Rules: []system.Rule(nil)}
	ptr8728564352 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Description: "Restriction rules for integers", Rules: []system.Rule(nil)}
	ptr8728564480 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "Default value if this property is omitted", Rules: []system.Rule(nil)}
	ptr8728920096 := &system.RuleBase{Optional: true}
	ptr8728855008 := &system.Int_rule{Base: ptr8728564480, RuleBase: ptr8728920096, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728560896 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil)}
	ptr8728920320 := &system.RuleBase{Optional: true}
	ptr8728855088 := &system.Int_rule{Base: ptr8728560896, RuleBase: ptr8728920320, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728561024 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil)}
	ptr8728920544 := &system.RuleBase{Optional: true}
	ptr8728855168 := &system.Int_rule{Base: ptr8728561024, RuleBase: ptr8728920544, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728561152 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil)}
	ptr8728920768 := &system.RuleBase{Optional: true}
	ptr8728855248 := &system.Int_rule{Base: ptr8728561152, RuleBase: ptr8728920768, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728792128 := &system.Type{Base: ptr8728564352, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728855008, "multipleOf": ptr8728855088, "minimum": ptr8728855168, "maximum": ptr8728855248}}
	ptr8728790672 := &system.Type{Base: ptr8728564224, Rule: ptr8728792128, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728565120 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Description: "Restriction rules for numbers", Rules: []system.Rule(nil)}
	ptr8728565504 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Rules: []system.Rule(nil)}
	ptr8728923616 := &system.RuleBase{Optional: true}
	ptr8728988240 := &system.JsonBool_rule{Base: ptr8728565504, RuleBase: ptr8728923616}
	ptr8728565632 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil)}
	ptr8728923872 := &system.RuleBase{Optional: true}
	ptr8728584672 := &system.Number_rule{Base: ptr8728565632, RuleBase: ptr8728923872, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}}
	ptr8728565760 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Rules: []system.Rule(nil)}
	ptr8728924256 := &system.RuleBase{Optional: true}
	ptr8728988672 := &system.JsonBool_rule{Base: ptr8728565760, RuleBase: ptr8728924256}
	ptr8728565248 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "Default value if this property is omitted", Rules: []system.Rule(nil)}
	ptr8728922784 := &system.RuleBase{Optional: true}
	ptr8728584288 := &system.Number_rule{Base: ptr8728565248, RuleBase: ptr8728922784, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}}
	ptr8728565376 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil)}
	ptr8728923040 := &system.RuleBase{Optional: true}
	ptr8728584960 := &system.Number_rule{Base: ptr8728565376, RuleBase: ptr8728923040, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}}
	ptr8728564608 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil)}
	ptr8728923296 := &system.RuleBase{Optional: true}
	ptr8728584480 := &system.Number_rule{Base: ptr8728564608, RuleBase: ptr8728923296, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}, Default: system.Number{}}
	ptr8728792464 := &system.Type{Base: ptr8728565120, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"exclusiveMinimum": ptr8728988240, "maximum": ptr8728584672, "exclusiveMaximum": ptr8728988672, "default": ptr8728584288, "multipleOf": ptr8728584960, "minimum": ptr8728584480}}
	ptr8728563200 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Description: "Automatically created basic rule for base", Rules: []system.Rule(nil)}
	ptr8728789888 := &system.Type{Base: ptr8728563200, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728848384 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Description: "Automatically created basic rule for type", Rules: []system.Rule(nil)}
	ptr8728792688 := &system.Type{Base: ptr8728848384, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728566016 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Description: "Restriction rules for references", Rules: []system.Rule(nil)}
	ptr8728566144 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "Default value of this is missing or null", Rules: []system.Rule(nil)}
	ptr8728925056 := &system.RuleBase{Optional: true}
	ptr8728571072 := &system.Reference_rule{Base: ptr8728566144, RuleBase: ptr8728925056, Default: system.Reference{}}
	ptr8728791344 := &system.Type{Base: ptr8728566016, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728571072}}
	ptr8728561408 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Description: "Restriction rules for maps", Rules: []system.Rule(nil)}
	ptr8728561536 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Rules: []system.Rule(nil)}
	ptr8728986800 := &system.Rule_rule{Base: ptr8728561536, RuleBase: (*system.RuleBase)(nil)}
	ptr8728564736 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the minimum number of items alowed in the array", Rules: []system.Rule(nil)}
	ptr8728921728 := &system.RuleBase{Optional: true}
	ptr8728855328 := &system.Int_rule{Base: ptr8728564736, RuleBase: ptr8728921728, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728564864 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the maximum number of items alowed in the array", Rules: []system.Rule(nil)}
	ptr8728921984 := &system.RuleBase{Optional: true}
	ptr8728855408 := &system.Int_rule{Base: ptr8728564864, RuleBase: ptr8728921984, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728789216 := &system.Type{Base: ptr8728561408, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8728986800, "minItems": ptr8728855328, "maxItems": ptr8728855408}}
	ptr8728565888 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Rules: []system.Rule(nil)}
	ptr8728791008 := &system.Type{Base: ptr8728565888, Rule: ptr8728791344, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728562432 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Description: "This is the most basic type.", Rules: []system.Rule(nil)}
	ptr8728562560 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "Type of the object.", Rules: []system.Rule(nil)}
	ptr8728903296 := &system.Reference_rule{Base: ptr8728562560, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728562688 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "All global objects should have an id.", Rules: []system.Rule(nil)}
	ptr8728923904 := &system.RuleBase{Optional: true}
	ptr8728903360 := &system.Reference_rule{Base: ptr8728562688, RuleBase: ptr8728923904, Default: system.Reference{}}
	ptr8728562816 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Description for the developer", Rules: []system.Rule(nil)}
	ptr8728924224 := &system.RuleBase{Optional: true}
	ptr8728989232 := &system.JsonString_rule{Base: ptr8728562816, RuleBase: ptr8728924224}
	ptr8728562944 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Extra validation rules for this object or descendants", Rules: []system.Rule(nil)}
	ptr8728924512 := &system.RuleBase{Optional: true}
	ptr8728563072 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728989456 := &system.Rule_rule{Base: ptr8728563072, RuleBase: (*system.RuleBase)(nil)}
	ptr8728854928 := &system.Array_rule{Base: ptr8728562944, RuleBase: ptr8728924512, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8728989456}
	ptr8728789664 := &system.Type{Base: ptr8728562432, Rule: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"type": ptr8728903296, "id": ptr8728903360, "description": ptr8728989232, "rules": ptr8728854928}}
	ptr8728846848 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Description: "This is the most basic type.", Rules: []system.Rule(nil)}
	ptr8728848256 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Id: system.Reference{}, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Rules: []system.Rule(nil)}
	ptr8728390080 := &system.RuleBase{Optional: true}
	ptr8728987312 := &system.Type_rule{Base: ptr8728848256, RuleBase: ptr8728390080}
	ptr8728846976 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Basic types don't have system:object added by default to the embedded types.", Rules: []system.Rule(nil)}
	ptr8728388800 := &system.RuleBase{Optional: true}
	ptr8728992128 := &system.JsonBool_rule{Base: ptr8728846976, RuleBase: ptr8728388800}
	ptr8728847104 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Types which this should embed - system:object is always added unless base = true.", Rules: []system.Rule(nil)}
	ptr8728389216 := &system.RuleBase{Optional: true}
	ptr8728847232 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728571968 := &system.Reference_rule{Base: ptr8728847232, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728856128 := &system.Array_rule{Base: ptr8728847104, RuleBase: ptr8728389216, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8728571968}
	ptr8728847360 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Array of interface types that this type should support", Rules: []system.Rule(nil)}
	ptr8728387936 := &system.RuleBase{Optional: true}
	ptr8728847488 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728568320 := &system.Reference_rule{Base: ptr8728847488, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728856208 := &system.Array_rule{Base: ptr8728847360, RuleBase: ptr8728387936, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8728568320}
	ptr8728847616 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is the native json type that represents this type. If omitted, default is object.", Rules: []system.Rule(nil)}
	ptr8728388640 := &system.RuleBase{Optional: true}
	ptr8728723456 := &system.String_rule{Base: ptr8728847616, RuleBase: ptr8728388640, Enum: []string{"string", "number", "bool", "array", "object", "map"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{Value: "object", Exists: true}}
	ptr8728847744 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Is this type an interface?", Rules: []system.Rule(nil)}
	ptr8728389152 := &system.RuleBase{Optional: true}
	ptr8728986464 := &system.JsonBool_rule{Base: ptr8728847744, RuleBase: ptr8728389152}
	ptr8728847872 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Exclude this type from the generated json?", Rules: []system.Rule(nil)}
	ptr8728395712 := &system.RuleBase{Optional: true}
	ptr8728986672 := &system.JsonBool_rule{Base: ptr8728847872, RuleBase: ptr8728395712}
	ptr8728848000 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "Each field is listed with it's type", Rules: []system.Rule(nil)}
	ptr8728389856 := &system.RuleBase{Optional: true}
	ptr8728848128 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728987104 := &system.Rule_rule{Base: ptr8728848128, RuleBase: (*system.RuleBase)(nil)}
	ptr8728568512 := &system.Map_rule{Base: ptr8728848000, RuleBase: ptr8728389856, Items: ptr8728987104, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728794928 := &system.Type{Base: ptr8728846848, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"rule": ptr8728987312, "basic": ptr8728992128, "embed": ptr8728856128, "is": ptr8728856208, "native": ptr8728723456, "interface": ptr8728986464, "exclude": ptr8728986672, "fields": ptr8728568512}}
	ptr8728566272 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Description: "This is one of several rule types, derived from the rules property of other types", Rules: []system.Rule(nil)}
	ptr8728791456 := &system.Type{Base: ptr8728566272, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil)}
	ptr8728567168 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Description: "Restriction rules for strings", Rules: []system.Rule(nil)}
	ptr8728567424 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "The value of this string is restricted to one of the provided values", Rules: []system.Rule(nil)}
	ptr8728927392 := &system.RuleBase{Optional: true}
	ptr8728567552 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728991072 := &system.JsonString_rule{Base: ptr8728567552, RuleBase: (*system.RuleBase)(nil)}
	ptr8728855648 := &system.Array_rule{Base: ptr8728567424, RuleBase: ptr8728927392, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8728991072}
	ptr8728567680 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "The value must be longer or equal to the provided minimum length", Rules: []system.Rule(nil)}
	ptr8728927616 := &system.RuleBase{Optional: true}
	ptr8728855808 := &system.Int_rule{Base: ptr8728567680, RuleBase: ptr8728927616, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728846336 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "The value must be shorter or equal to the provided maximum length", Rules: []system.Rule(nil)}
	ptr8728927840 := &system.RuleBase{Optional: true}
	ptr8728855888 := &system.Int_rule{Base: ptr8728846336, RuleBase: ptr8728927840, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728846464 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is a string that the value must match", Rules: []system.Rule(nil)}
	ptr8728928064 := &system.RuleBase{Optional: true}
	ptr8728723984 := &system.String_rule{Base: ptr8728846464, RuleBase: ptr8728928064, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}}
	ptr8728846592 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is a regex to match the value to", Rules: []system.Rule(nil)}
	ptr8728394688 := &system.RuleBase{Optional: true}
	ptr8728724160 := &system.String_rule{Base: ptr8728846592, RuleBase: ptr8728394688, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}}
	ptr8728846720 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This restricts the value to one of several built-in formats", Rules: []system.Rule(nil)}
	ptr8728387776 := &system.RuleBase{Optional: true}
	ptr8728724336 := &system.String_rule{Base: ptr8728846720, RuleBase: ptr8728387776, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}}
	ptr8728567296 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Default value of this is missing or null", Rules: []system.Rule(nil)}
	ptr8728927072 := &system.RuleBase{Optional: true}
	ptr8728723632 := &system.String_rule{Base: ptr8728567296, RuleBase: ptr8728927072, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}}
	ptr8728794256 := &system.Type{Base: ptr8728567168, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"enum": ptr8728855648, "minLength": ptr8728855808, "maxLength": ptr8728855888, "equal": ptr8728723984, "pattern": ptr8728724160, "format": ptr8728724336, "default": ptr8728723632}}
	ptr8728567040 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Description: "This is the native json string data type", Rules: []system.Rule(nil)}
	ptr8728792016 := &system.Type{Base: ptr8728567040, Rule: ptr8728794256, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728563456 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Description: "Restriction rules for bools", Rules: []system.Rule(nil)}
	ptr8728563584 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Default value if this is missing or null", Rules: []system.Rule(nil)}
	ptr8728925280 := &system.RuleBase{Optional: true}
	ptr8728925216 := &system.Bool_rule{Base: ptr8728563584, RuleBase: ptr8728925280, Default: system.Bool{}}
	ptr8728790224 := &system.Type{Base: ptr8728563456, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728925216}}
	ptr8728561280 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Description: "This is the native json object data type.", Rules: []system.Rule(nil)}
	ptr8728788992 := &system.Type{Base: ptr8728561280, Rule: ptr8728789216, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728564992 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Description: "This is the native json number data type", Rules: []system.Rule(nil)}
	ptr8728790896 := &system.Type{Base: ptr8728564992, Rule: ptr8728792464, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728566400 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Description: "Automatically created basic rule for rule", Rules: []system.Rule(nil)}
	ptr8728791568 := &system.Type{Base: ptr8728566400, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728564096 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Description: "Automatically created basic rule for imports", Rules: []system.Rule(nil)}
	ptr8728790448 := &system.Type{Base: ptr8728564096, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728563328 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Description: "This is the native json bool data type", Rules: []system.Rule(nil)}
	ptr8728790112 := &system.Type{Base: ptr8728563328, Rule: ptr8728790224, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Fields: map[string]system.Rule(nil)}
	ptr8728563712 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Description: "Lists imports used in this package.", Rules: []system.Rule(nil)}
	ptr8728563840 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "Map of import name to path.", Rules: []system.Rule(nil)}
	ptr8728563968 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728723808 := &system.String_rule{Base: ptr8728563968, RuleBase: (*system.RuleBase)(nil), Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}, Default: system.String{}}
	ptr8728571200 := &system.Map_rule{Base: ptr8728563840, RuleBase: (*system.RuleBase)(nil), Items: ptr8728723808, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728790336 := &system.Type{Base: ptr8728563712, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8728571200}}
	system.RegisterType("kego.io/system", "@type", ptr8728792688)
	system.RegisterType("kego.io/system", "@map", ptr8728789216)
	system.RegisterType("kego.io/system", "@string", ptr8728794256)
	system.RegisterType("kego.io/system", "string", ptr8728792016)
	system.RegisterType("kego.io/system", "@array", ptr8728789440)
	system.RegisterType("kego.io/system", "ruleBase", ptr8728791680)
	system.RegisterType("kego.io/system", "@ruleBase", ptr8728791904)
	system.RegisterType("kego.io/system", "@number", ptr8728792464)
	system.RegisterType("kego.io/system", "@rule", ptr8728791568)
	system.RegisterType("kego.io/system", "bool", ptr8728790112)
	system.RegisterType("kego.io/system", "type", ptr8728794928)
	system.RegisterType("kego.io/system", "imports", ptr8728790336)
	system.RegisterType("kego.io/system", "@base", ptr8728789888)
	system.RegisterType("kego.io/system", "@reference", ptr8728791344)
	system.RegisterType("kego.io/system", "reference", ptr8728791008)
	system.RegisterType("kego.io/system", "map", ptr8728788992)
	system.RegisterType("kego.io/system", "number", ptr8728790896)
	system.RegisterType("kego.io/system", "@int", ptr8728792128)
	system.RegisterType("kego.io/system", "int", ptr8728790672)
	system.RegisterType("kego.io/system", "base", ptr8728789664)
	system.RegisterType("kego.io/system", "rule", ptr8728791456)
	system.RegisterType("kego.io/system", "array", ptr8728789328)
	system.RegisterType("kego.io/system", "@bool", ptr8728790224)
	system.RegisterType("kego.io/system", "@imports", ptr8728790448)
}
