package types

import (
	system "kego.io/system"
)

func init() {
	ptr8728923392 := &system.Base{Description: "This is the native json bool data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}}
	ptr8728923648 := &system.Base{Description: "Restriction rules for bools", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr8728923776 := &system.Base{Description: "Default value if this is missing or null", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8729023584 := &system.RuleBase{Optional: true}
	ptr8729023520 := &system.Bool_rule{Base: ptr8728923776, RuleBase: ptr8729023584, Default: system.Bool{}}
	ptr8728923520 := &system.Type{Base: ptr8728923648, Fields: map[string]system.Rule{"default": ptr8729023520}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728923264 := &system.Type{Base: ptr8728923392, Fields: map[string]system.Rule(nil), Rule: ptr8728923520, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}}
	ptr8728923136 := &system.Base{Description: "Automatically created basic rule for base", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}}
	ptr8728923008 := &system.Type{Base: ptr8728923136, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728926976 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}}
	ptr8728927232 := &system.Base{Description: "Restriction rules for references", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr8728927360 := &system.Base{Description: "Default value of this is missing or null", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8729023360 := &system.RuleBase{Optional: true}
	ptr8728571072 := &system.Reference_rule{Base: ptr8728927360, RuleBase: ptr8729023360, Default: system.Reference{}}
	ptr8728927104 := &system.Type{Base: ptr8728927232, Fields: map[string]system.Rule{"default": ptr8728571072}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728926848 := &system.Type{Base: ptr8728926976, Fields: map[string]system.Rule(nil), Rule: ptr8728927104, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}}
	ptr8728925056 := &system.Base{Description: "Restriction rules for integers", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr8728925184 := &system.Base{Description: "Default value if this property is omitted", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729018400 := &system.RuleBase{Optional: true}
	ptr8728969696 := &system.Int_rule{Base: ptr8728925184, RuleBase: ptr8729018400, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728920064 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729018624 := &system.RuleBase{Optional: true}
	ptr8728969776 := &system.Int_rule{Base: ptr8728920064, RuleBase: ptr8729018624, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728920192 := &system.Base{Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729018848 := &system.RuleBase{Optional: true}
	ptr8728969856 := &system.Int_rule{Base: ptr8728920192, RuleBase: ptr8729018848, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728920320 := &system.Base{Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729019072 := &system.RuleBase{Optional: true}
	ptr8728969936 := &system.Int_rule{Base: ptr8728920320, RuleBase: ptr8729019072, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728924928 := &system.Type{Base: ptr8728925056, Fields: map[string]system.Rule{"default": ptr8728969696, "multipleOf": ptr8728969776, "minimum": ptr8728969856, "maximum": ptr8728969936}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728924544 := &system.Base{Description: "Automatically created basic rule for imports", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}}
	ptr8728924416 := &system.Type{Base: ptr8728924544, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728925696 := &system.Base{Description: "This is the native json number data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}}
	ptr8728925952 := &system.Base{Description: "Restriction rules for numbers", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr8728926336 := &system.Base{Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729021600 := &system.RuleBase{Optional: true}
	ptr8728592672 := &system.Number_rule{Base: ptr8728926336, RuleBase: ptr8729021600, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728926464 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8729021920 := &system.RuleBase{Optional: true}
	ptr8728996432 := &system.JsonBool_rule{Base: ptr8728926464, RuleBase: ptr8729021920}
	ptr8728926592 := &system.Base{Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729022176 := &system.RuleBase{Optional: true}
	ptr8728592864 := &system.Number_rule{Base: ptr8728926592, RuleBase: ptr8729022176, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728926720 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8729022560 := &system.RuleBase{Optional: true}
	ptr8728996864 := &system.JsonBool_rule{Base: ptr8728926720, RuleBase: ptr8729022560}
	ptr8728926080 := &system.Base{Description: "Default value if this property is omitted", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729021088 := &system.RuleBase{Optional: true}
	ptr8728592480 := &system.Number_rule{Base: ptr8728926080, RuleBase: ptr8729021088, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728926208 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729021344 := &system.RuleBase{Optional: true}
	ptr8728593152 := &system.Number_rule{Base: ptr8728926208, RuleBase: ptr8729021344, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728925824 := &system.Type{Base: ptr8728925952, Fields: map[string]system.Rule{"minimum": ptr8728592672, "exclusiveMinimum": ptr8728996432, "maximum": ptr8728592864, "exclusiveMaximum": ptr8728996864, "default": ptr8728592480, "multipleOf": ptr8728593152}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728925568 := &system.Type{Base: ptr8728925696, Fields: map[string]system.Rule(nil), Rule: ptr8728925824, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}}
	ptr8728920832 := &system.Base{Description: "Restriction rules for maps", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr8728920960 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8728994992 := &system.Rule_rule{Base: ptr8728920960, RuleBase: (*system.RuleBase)(nil)}
	ptr8728925312 := &system.Base{Description: "This is the minimum number of items alowed in the array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729020032 := &system.RuleBase{Optional: true}
	ptr8728970016 := &system.Int_rule{Base: ptr8728925312, RuleBase: ptr8729020032, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728925440 := &system.Base{Description: "This is the maximum number of items alowed in the array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729020288 := &system.RuleBase{Optional: true}
	ptr8728970096 := &system.Int_rule{Base: ptr8728925440, RuleBase: ptr8729020288, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728920704 := &system.Type{Base: ptr8728920832, Fields: map[string]system.Rule{"items": ptr8728994992, "minItems": ptr8728970016, "maxItems": ptr8728970096}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728588160 := &system.Base{Description: "This is the most basic type.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr8728589568 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Id: system.Reference{}}
	ptr8728398272 := &system.RuleBase{Optional: true}
	ptr8728995504 := &system.Type_rule{Base: ptr8728589568, RuleBase: ptr8728398272}
	ptr8728588288 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728396992 := &system.RuleBase{Optional: true}
	ptr8729000320 := &system.JsonBool_rule{Base: ptr8728588288, RuleBase: ptr8728396992}
	ptr8728588416 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728397408 := &system.RuleBase{Optional: true}
	ptr8728588544 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8728571968 := &system.Reference_rule{Base: ptr8728588544, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728970816 := &system.Array_rule{Base: ptr8728588416, RuleBase: ptr8728397408, Items: ptr8728571968, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728588672 := &system.Base{Description: "Array of interface types that this type should support", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728396128 := &system.RuleBase{Optional: true}
	ptr8728588800 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8728568320 := &system.Reference_rule{Base: ptr8728588800, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728970896 := &system.Array_rule{Base: ptr8728588672, RuleBase: ptr8728396128, Items: ptr8728568320, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728588928 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728396832 := &system.RuleBase{Optional: true}
	ptr8728739840 := &system.String_rule{Base: ptr8728588928, RuleBase: ptr8728396832, Pattern: system.String{}, Format: system.String{}, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728589056 := &system.Base{Description: "Is this type an interface?", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728397344 := &system.RuleBase{Optional: true}
	ptr8728994656 := &system.JsonBool_rule{Base: ptr8728589056, RuleBase: ptr8728397344}
	ptr8728589184 := &system.Base{Description: "Exclude this type from the generated json?", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728403904 := &system.RuleBase{Optional: true}
	ptr8728994864 := &system.JsonBool_rule{Base: ptr8728589184, RuleBase: ptr8728403904}
	ptr8728589312 := &system.Base{Description: "Each field is listed with it's type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728398048 := &system.RuleBase{Optional: true}
	ptr8728589440 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8728995296 := &system.Rule_rule{Base: ptr8728589440, RuleBase: (*system.RuleBase)(nil)}
	ptr8728568512 := &system.Map_rule{Base: ptr8728589312, RuleBase: ptr8728398048, Items: ptr8728995296, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728588032 := &system.Type{Base: ptr8728588160, Fields: map[string]system.Rule{"rule": ptr8728995504, "basic": ptr8729000320, "embed": ptr8728970816, "is": ptr8728970896, "native": ptr8728739840, "interface": ptr8728994656, "exclude": ptr8728994864, "fields": ptr8728568512}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728584320 := &system.Base{Description: "This is the native json string data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}}
	ptr8728586880 := &system.Base{Description: "Restriction rules for strings", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr8728587904 := &system.Base{Description: "This restricts the value to one of several built-in formats", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728395968 := &system.RuleBase{Optional: true}
	ptr8728740720 := &system.String_rule{Base: ptr8728587904, RuleBase: ptr8728395968, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728587008 := &system.Base{Description: "Default value of this is missing or null", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8729025376 := &system.RuleBase{Optional: true}
	ptr8728740016 := &system.String_rule{Base: ptr8728587008, RuleBase: ptr8729025376, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728587136 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8729025696 := &system.RuleBase{Optional: true}
	ptr8728587264 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728999264 := &system.JsonString_rule{Base: ptr8728587264, RuleBase: (*system.RuleBase)(nil)}
	ptr8728970336 := &system.Array_rule{Base: ptr8728587136, RuleBase: ptr8729025696, Items: ptr8728999264, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728587392 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729025920 := &system.RuleBase{Optional: true}
	ptr8728970496 := &system.Int_rule{Base: ptr8728587392, RuleBase: ptr8729025920, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728587520 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729026144 := &system.RuleBase{Optional: true}
	ptr8728970576 := &system.Int_rule{Base: ptr8728587520, RuleBase: ptr8729026144, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728587648 := &system.Base{Description: "This is a string that the value must match", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8729026368 := &system.RuleBase{Optional: true}
	ptr8728740368 := &system.String_rule{Base: ptr8728587648, RuleBase: ptr8729026368, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728587776 := &system.Base{Description: "This is a regex to match the value to", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728402880 := &system.RuleBase{Optional: true}
	ptr8728740544 := &system.String_rule{Base: ptr8728587776, RuleBase: ptr8728402880, Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728586752 := &system.Type{Base: ptr8728586880, Fields: map[string]system.Rule{"format": ptr8728740720, "default": ptr8728740016, "enum": ptr8728970336, "minLength": ptr8728970496, "maxLength": ptr8728970576, "equal": ptr8728740368, "pattern": ptr8728740544}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728592256 := &system.Type{Base: ptr8728584320, Fields: map[string]system.Rule(nil), Rule: ptr8728586752, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}}
	ptr8728592128 := &system.Base{Description: "Automatically created basic rule for ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}}
	ptr8728592000 := &system.Type{Base: ptr8728592128, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728922240 := &system.Base{Description: "This is the most basic type.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}}
	ptr8728922368 := &system.Base{Description: "Type of the object.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8728911488 := &system.Reference_rule{Base: ptr8728922368, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728922496 := &system.Base{Description: "All global objects should have an id.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8729022208 := &system.RuleBase{Optional: true}
	ptr8728911552 := &system.Reference_rule{Base: ptr8728922496, RuleBase: ptr8729022208, Default: system.Reference{}}
	ptr8728922624 := &system.Base{Description: "Description for the developer", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8729022528 := &system.RuleBase{Optional: true}
	ptr8728997424 := &system.JsonString_rule{Base: ptr8728922624, RuleBase: ptr8729022528}
	ptr8728922752 := &system.Base{Description: "Extra validation rules for this object or descendants", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8729022816 := &system.RuleBase{Optional: true}
	ptr8728922880 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8728997648 := &system.Rule_rule{Base: ptr8728922880, RuleBase: (*system.RuleBase)(nil)}
	ptr8728969616 := &system.Array_rule{Base: ptr8728922752, RuleBase: ptr8729022816, Items: ptr8728997648, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728922112 := &system.Type{Base: ptr8728922240, Fields: map[string]system.Rule{"type": ptr8728911488, "id": ptr8728911552, "description": ptr8728997424, "rules": ptr8728969616}, Rule: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728921472 := &system.Base{Description: "Restriction rules for arrays", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr8728921600 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8728996336 := &system.Rule_rule{Base: ptr8728921600, RuleBase: (*system.RuleBase)(nil)}
	ptr8728921728 := &system.Base{Description: "This is the minimum number of items allowed in the array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729020928 := &system.RuleBase{Optional: true}
	ptr8728969456 := &system.Int_rule{Base: ptr8728921728, RuleBase: ptr8729020928, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728921856 := &system.Base{Description: "This is the maximum number of items allowed in the array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729021184 := &system.RuleBase{Optional: true}
	ptr8728969536 := &system.Int_rule{Base: ptr8728921856, RuleBase: ptr8729021184, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728921984 := &system.Base{Description: "If this is true, each item must be unique", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8729021472 := &system.RuleBase{Optional: true}
	ptr8728996768 := &system.JsonBool_rule{Base: ptr8728921984, RuleBase: ptr8729021472}
	ptr8728921344 := &system.Type{Base: ptr8728921472, Fields: map[string]system.Rule{"items": ptr8728996336, "minItems": ptr8728969456, "maxItems": ptr8728969536, "uniqueItems": ptr8728996768}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728927872 := &system.Base{Description: "Automatically created basic rule for rule", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr8728927744 := &system.Type{Base: ptr8728927872, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728921216 := &system.Base{Description: "This is the native json array data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}}
	ptr8728921088 := &system.Type{Base: ptr8728921216, Fields: map[string]system.Rule(nil), Rule: ptr8728921344, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}}
	ptr8728589824 := &system.Base{Description: "Automatically created basic rule for type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr8728589696 := &system.Type{Base: ptr8728589824, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728924032 := &system.Base{Description: "Lists imports used in this package.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}}
	ptr8728924160 := &system.Base{Description: "Map of path to alias.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728924288 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728740192 := &system.String_rule{Base: ptr8728924288, RuleBase: (*system.RuleBase)(nil), Pattern: system.String{}, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}}
	ptr8728571200 := &system.Map_rule{Base: ptr8728924160, RuleBase: (*system.RuleBase)(nil), Items: ptr8728740192, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728923904 := &system.Type{Base: ptr8728924032, Fields: map[string]system.Rule{"imports": ptr8728571200}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728920576 := &system.Base{Description: "This is the native json object data type.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}}
	ptr8728920448 := &system.Type{Base: ptr8728920576, Fields: map[string]system.Rule(nil), Rule: ptr8728920704, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}}
	ptr8728924800 := &system.Base{Description: "This is the int data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}}
	ptr8728924672 := &system.Type{Base: ptr8728924800, Fields: map[string]system.Rule(nil), Rule: ptr8728924928, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}}
	ptr8728928128 := &system.Base{Description: "All rules should embed this type.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}
	ptr8728591744 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8729024192 := &system.RuleBase{Optional: true}
	ptr8728998192 := &system.JsonBool_rule{Base: ptr8728591744, RuleBase: ptr8729024192}
	ptr8728591872 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8729024608 := &system.RuleBase{Optional: true}
	ptr8728998368 := &system.JsonString_rule{Base: ptr8728591872, RuleBase: ptr8729024608}
	ptr8728928000 := &system.Type{Base: ptr8728928128, Fields: map[string]system.Rule{"optional": ptr8728998192, "selector": ptr8728998368}, Rule: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728927616 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}
	ptr8728927488 := &system.Type{Base: ptr8728927616, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true}
	system.RegisterType("kego.io/system", "@map", ptr8728920704)
	system.RegisterType("kego.io/system", "@number", ptr8728925824)
	system.RegisterType("kego.io/system", "string", ptr8728592256)
	system.RegisterType("kego.io/system", "base", ptr8728922112)
	system.RegisterType("kego.io/system", "@rule", ptr8728927744)
	system.RegisterType("kego.io/system", "ruleBase", ptr8728928000)
	system.RegisterType("kego.io/system", "reference", ptr8728926848)
	system.RegisterType("kego.io/system", "@imports", ptr8728924416)
	system.RegisterType("kego.io/system", "number", ptr8728925568)
	system.RegisterType("kego.io/system", "@bool", ptr8728923520)
	system.RegisterType("kego.io/system", "@ruleBase", ptr8728592000)
	system.RegisterType("kego.io/system", "array", ptr8728921088)
	system.RegisterType("kego.io/system", "rule", ptr8728927488)
	system.RegisterType("kego.io/system", "@int", ptr8728924928)
	system.RegisterType("kego.io/system", "@string", ptr8728586752)
	system.RegisterType("kego.io/system", "imports", ptr8728923904)
	system.RegisterType("kego.io/system", "@reference", ptr8728927104)
	system.RegisterType("kego.io/system", "int", ptr8728924672)
	system.RegisterType("kego.io/system", "bool", ptr8728923264)
	system.RegisterType("kego.io/system", "@base", ptr8728923008)
	system.RegisterType("kego.io/system", "type", ptr8728588032)
	system.RegisterType("kego.io/system", "@array", ptr8728921344)
	system.RegisterType("kego.io/system", "@type", ptr8728589696)
	system.RegisterType("kego.io/system", "map", ptr8728920448)
}
