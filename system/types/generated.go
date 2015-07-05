package types

import (
	system "kego.io/system"
)

func init() {
	ptr8729085952 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Description: "Automatically created basic rule for type", Rules: []system.Rule(nil)}
	ptr8728800880 := &system.Type{Base: ptr8729085952, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728597760 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Description: "This is the native json number data type", Rules: []system.Rule(nil)}
	ptr8728597888 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Description: "Restriction rules for numbers", Rules: []system.Rule(nil)}
	ptr8728598016 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "Default value if this property is omitted", Rules: []system.Rule(nil)}
	ptr8728922784 := &system.RuleBase{Optional: true}
	ptr8728567904 := &system.Number_rule{Base: ptr8728598016, RuleBase: ptr8728922784, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728598144 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil)}
	ptr8728923040 := &system.RuleBase{Optional: true}
	ptr8728568576 := &system.Number_rule{Base: ptr8728598144, RuleBase: ptr8728923040, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728597376 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil)}
	ptr8728923296 := &system.RuleBase{Optional: true}
	ptr8728568096 := &system.Number_rule{Base: ptr8728597376, RuleBase: ptr8728923296, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728598272 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Rules: []system.Rule(nil)}
	ptr8728923616 := &system.RuleBase{Optional: true}
	ptr8728996432 := &system.JsonBool_rule{Base: ptr8728598272, RuleBase: ptr8728923616}
	ptr8728598400 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil)}
	ptr8728923872 := &system.RuleBase{Optional: true}
	ptr8728568288 := &system.Number_rule{Base: ptr8728598400, RuleBase: ptr8728923872, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728598528 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Rules: []system.Rule(nil)}
	ptr8728924256 := &system.RuleBase{Optional: true}
	ptr8728996864 := &system.JsonBool_rule{Base: ptr8728598528, RuleBase: ptr8728924256}
	ptr8728800656 := &system.Type{Base: ptr8728597888, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728567904, "multipleOf": ptr8728568576, "minimum": ptr8728568096, "exclusiveMinimum": ptr8728996432, "maximum": ptr8728568288, "exclusiveMaximum": ptr8728996864}, Rule: (*system.Type)(nil)}
	ptr8728799088 := &system.Type{Base: ptr8728597760, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728800656}
	ptr8729084416 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Description: "This is the most basic type.", Rules: []system.Rule(nil)}
	ptr8729085568 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "Each field is listed with it's type", Rules: []system.Rule(nil)}
	ptr8728398048 := &system.RuleBase{Optional: true}
	ptr8729085696 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728995296 := &system.Rule_rule{Base: ptr8729085696, RuleBase: (*system.RuleBase)(nil)}
	ptr8728576704 := &system.Map_rule{Base: ptr8729085568, RuleBase: ptr8728398048, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8728995296}
	ptr8729085824 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Id: system.Reference{}, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Rules: []system.Rule(nil)}
	ptr8728398272 := &system.RuleBase{Optional: true}
	ptr8728995504 := &system.Type_rule{Base: ptr8729085824, RuleBase: ptr8728398272}
	ptr8729084544 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Basic types don't have system:object added by default to the embedded types.", Rules: []system.Rule(nil)}
	ptr8728396992 := &system.RuleBase{Optional: true}
	ptr8729000320 := &system.JsonBool_rule{Base: ptr8729084544, RuleBase: ptr8728396992}
	ptr8729084672 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Types which this should embed - system:object is always added unless base = true.", Rules: []system.Rule(nil)}
	ptr8728397408 := &system.RuleBase{Optional: true}
	ptr8729084800 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728580160 := &system.Reference_rule{Base: ptr8729084800, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728725056 := &system.Array_rule{Base: ptr8729084672, RuleBase: ptr8728397408, Items: ptr8728580160, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729084928 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Array of interface types that this type should support", Rules: []system.Rule(nil)}
	ptr8728396128 := &system.RuleBase{Optional: true}
	ptr8729085056 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728576512 := &system.Reference_rule{Base: ptr8729085056, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728725136 := &system.Array_rule{Base: ptr8729084928, RuleBase: ptr8728396128, Items: ptr8728576512, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729085184 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is the native json type that represents this type. If omitted, default is object.", Rules: []system.Rule(nil)}
	ptr8728396832 := &system.RuleBase{Optional: true}
	ptr8728739840 := &system.String_rule{Base: ptr8729085184, RuleBase: ptr8728396832, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729085312 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Is this type an interface?", Rules: []system.Rule(nil)}
	ptr8728397344 := &system.RuleBase{Optional: true}
	ptr8728994656 := &system.JsonBool_rule{Base: ptr8729085312, RuleBase: ptr8728397344}
	ptr8729085440 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Exclude this type from the generated json?", Rules: []system.Rule(nil)}
	ptr8728403904 := &system.RuleBase{Optional: true}
	ptr8728994864 := &system.JsonBool_rule{Base: ptr8729085440, RuleBase: ptr8728403904}
	ptr8728803120 := &system.Type{Base: ptr8729084416, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"fields": ptr8728576704, "rule": ptr8728995504, "basic": ptr8729000320, "embed": ptr8728725056, "is": ptr8728725136, "native": ptr8728739840, "interface": ptr8728994656, "exclude": ptr8728994864}, Rule: (*system.Type)(nil)}
	ptr8728598784 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Description: "Restriction rules for references", Rules: []system.Rule(nil)}
	ptr8728598912 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "Default value of this is missing or null", Rules: []system.Rule(nil)}
	ptr8728925056 := &system.RuleBase{Optional: true}
	ptr8728579264 := &system.Reference_rule{Base: ptr8728598912, RuleBase: ptr8728925056, Default: system.Reference{}}
	ptr8728799536 := &system.Type{Base: ptr8728598784, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728579264}, Rule: (*system.Type)(nil)}
	ptr8728596224 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Description: "Restriction rules for bools", Rules: []system.Rule(nil)}
	ptr8728596352 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Default value if this is missing or null", Rules: []system.Rule(nil)}
	ptr8728925280 := &system.RuleBase{Optional: true}
	ptr8728925216 := &system.Bool_rule{Base: ptr8728596352, RuleBase: ptr8728925280, Default: system.Bool{}}
	ptr8728798416 := &system.Type{Base: ptr8728596224, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728925216}, Rule: (*system.Type)(nil)}
	ptr8728599680 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Description: "Automatically created basic rule for ruleBase", Rules: []system.Rule(nil)}
	ptr8728800096 := &system.Type{Base: ptr8728599680, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728597120 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Description: "Restriction rules for integers", Rules: []system.Rule(nil)}
	ptr8728597248 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "Default value if this property is omitted", Rules: []system.Rule(nil)}
	ptr8728920096 := &system.RuleBase{Optional: true}
	ptr8728723936 := &system.Int_rule{Base: ptr8728597248, RuleBase: ptr8728920096, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}}
	ptr8728593664 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil)}
	ptr8728920320 := &system.RuleBase{Optional: true}
	ptr8728724016 := &system.Int_rule{Base: ptr8728593664, RuleBase: ptr8728920320, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}}
	ptr8728593792 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil)}
	ptr8728920544 := &system.RuleBase{Optional: true}
	ptr8728724096 := &system.Int_rule{Base: ptr8728593792, RuleBase: ptr8728920544, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}}
	ptr8728593920 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil)}
	ptr8728920768 := &system.RuleBase{Optional: true}
	ptr8728724176 := &system.Int_rule{Base: ptr8728593920, RuleBase: ptr8728920768, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}}
	ptr8728800320 := &system.Type{Base: ptr8728597120, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728723936, "multipleOf": ptr8728724016, "minimum": ptr8728724096, "maximum": ptr8728724176}, Rule: (*system.Type)(nil)}
	ptr8728599040 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Description: "This is one of several rule types, derived from the rules property of other types", Rules: []system.Rule(nil)}
	ptr8728799648 := &system.Type{Base: ptr8728599040, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728599296 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Description: "All rules should embed this type.", Rules: []system.Rule(nil)}
	ptr8728599424 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this rule is a field, this specifies that the field is optional", Rules: []system.Rule(nil)}
	ptr8728925888 := &system.RuleBase{Optional: true}
	ptr8728998192 := &system.JsonBool_rule{Base: ptr8728599424, RuleBase: ptr8728925888}
	ptr8728599552 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Json selector defining what nodes this rule should be applied to.", Rules: []system.Rule(nil)}
	ptr8728926304 := &system.RuleBase{Optional: true}
	ptr8728998368 := &system.JsonString_rule{Base: ptr8728599552, RuleBase: ptr8728926304}
	ptr8728799872 := &system.Type{Base: ptr8728599296, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8728998192, "selector": ptr8728998368}, Rule: (*system.Type)(nil), Basic: true}
	ptr8728594432 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Description: "This is the native json array data type", Rules: []system.Rule(nil)}
	ptr8728594560 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Description: "Restriction rules for arrays", Rules: []system.Rule(nil)}
	ptr8728594688 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Description: "This is a rule object, defining the type and restrictions on the value of the items", Rules: []system.Rule(nil)}
	ptr8728996336 := &system.Rule_rule{Base: ptr8728594688, RuleBase: (*system.RuleBase)(nil)}
	ptr8728594816 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the minimum number of items allowed in the array", Rules: []system.Rule(nil)}
	ptr8728922624 := &system.RuleBase{Optional: true}
	ptr8728723696 := &system.Int_rule{Base: ptr8728594816, RuleBase: ptr8728922624, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}}
	ptr8728594944 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the maximum number of items allowed in the array", Rules: []system.Rule(nil)}
	ptr8728922880 := &system.RuleBase{Optional: true}
	ptr8728723776 := &system.Int_rule{Base: ptr8728594944, RuleBase: ptr8728922880, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}}
	ptr8728595072 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, each item must be unique", Rules: []system.Rule(nil)}
	ptr8728923168 := &system.RuleBase{Optional: true}
	ptr8728996768 := &system.JsonBool_rule{Base: ptr8728595072, RuleBase: ptr8728923168}
	ptr8728797632 := &system.Type{Base: ptr8728594560, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8728996336, "minItems": ptr8728723696, "maxItems": ptr8728723776, "uniqueItems": ptr8728996768}, Rule: (*system.Type)(nil)}
	ptr8728797520 := &system.Type{Base: ptr8728594432, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728797632}
	ptr8728599168 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Description: "Automatically created basic rule for rule", Rules: []system.Rule(nil)}
	ptr8728799760 := &system.Type{Base: ptr8728599168, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728595968 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Description: "Automatically created basic rule for base", Rules: []system.Rule(nil)}
	ptr8728798080 := &system.Type{Base: ptr8728595968, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728596480 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Description: "Lists imports used in this package.", Rules: []system.Rule(nil)}
	ptr8728596608 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "Map of path to alias.", Rules: []system.Rule(nil)}
	ptr8728596736 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728740192 := &system.String_rule{Base: ptr8728596736, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728579392 := &system.Map_rule{Base: ptr8728596608, RuleBase: (*system.RuleBase)(nil), MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}, Items: ptr8728740192}
	ptr8728798528 := &system.Type{Base: ptr8728596480, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8728579392}, Rule: (*system.Type)(nil)}
	ptr8728596096 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Description: "This is the native json bool data type", Rules: []system.Rule(nil)}
	ptr8728798304 := &system.Type{Base: ptr8728596096, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728798416}
	ptr8728594048 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Description: "This is the native json object data type.", Rules: []system.Rule(nil)}
	ptr8728594176 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Description: "Restriction rules for maps", Rules: []system.Rule(nil)}
	ptr8728594304 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Rules: []system.Rule(nil)}
	ptr8728994992 := &system.Rule_rule{Base: ptr8728594304, RuleBase: (*system.RuleBase)(nil)}
	ptr8728597504 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the minimum number of items alowed in the array", Rules: []system.Rule(nil)}
	ptr8728921728 := &system.RuleBase{Optional: true}
	ptr8728724256 := &system.Int_rule{Base: ptr8728597504, RuleBase: ptr8728921728, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}}
	ptr8728597632 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the maximum number of items alowed in the array", Rules: []system.Rule(nil)}
	ptr8728921984 := &system.RuleBase{Optional: true}
	ptr8728724336 := &system.Int_rule{Base: ptr8728597632, RuleBase: ptr8728921984, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}}
	ptr8728797408 := &system.Type{Base: ptr8728594176, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8728994992, "minItems": ptr8728724256, "maxItems": ptr8728724336}, Rule: (*system.Type)(nil)}
	ptr8728797184 := &system.Type{Base: ptr8728594048, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728797408}
	ptr8728599808 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Description: "This is the native json string data type", Rules: []system.Rule(nil)}
	ptr8728599936 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Description: "Restriction rules for strings", Rules: []system.Rule(nil)}
	ptr8729084032 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is a string that the value must match", Rules: []system.Rule(nil)}
	ptr8728928064 := &system.RuleBase{Optional: true}
	ptr8728740368 := &system.String_rule{Base: ptr8729084032, RuleBase: ptr8728928064, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729084160 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is a regex to match the value to", Rules: []system.Rule(nil)}
	ptr8728402880 := &system.RuleBase{Optional: true}
	ptr8728740544 := &system.String_rule{Base: ptr8729084160, RuleBase: ptr8728402880, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729084288 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This restricts the value to one of several built-in formats", Rules: []system.Rule(nil)}
	ptr8728395968 := &system.RuleBase{Optional: true}
	ptr8728740720 := &system.String_rule{Base: ptr8729084288, RuleBase: ptr8728395968, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728600064 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Default value of this is missing or null", Rules: []system.Rule(nil)}
	ptr8728927072 := &system.RuleBase{Optional: true}
	ptr8728740016 := &system.String_rule{Base: ptr8728600064, RuleBase: ptr8728927072, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728600192 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "The value of this string is restricted to one of the provided values", Rules: []system.Rule(nil)}
	ptr8728927392 := &system.RuleBase{Optional: true}
	ptr8728600320 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728999264 := &system.JsonString_rule{Base: ptr8728600320, RuleBase: (*system.RuleBase)(nil)}
	ptr8728724576 := &system.Array_rule{Base: ptr8728600192, RuleBase: ptr8728927392, Items: ptr8728999264, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728600448 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "The value must be longer or equal to the provided minimum length", Rules: []system.Rule(nil)}
	ptr8728927616 := &system.RuleBase{Optional: true}
	ptr8728724736 := &system.Int_rule{Base: ptr8728600448, RuleBase: ptr8728927616, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}}
	ptr8729083904 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "The value must be shorter or equal to the provided maximum length", Rules: []system.Rule(nil)}
	ptr8728927840 := &system.RuleBase{Optional: true}
	ptr8728724816 := &system.Int_rule{Base: ptr8729083904, RuleBase: ptr8728927840, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}}
	ptr8728802448 := &system.Type{Base: ptr8728599936, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"equal": ptr8728740368, "pattern": ptr8728740544, "format": ptr8728740720, "default": ptr8728740016, "enum": ptr8728724576, "minLength": ptr8728724736, "maxLength": ptr8728724816}, Rule: (*system.Type)(nil)}
	ptr8728800208 := &system.Type{Base: ptr8728599808, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728802448}
	ptr8728596992 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Description: "This is the int data type", Rules: []system.Rule(nil)}
	ptr8728798864 := &system.Type{Base: ptr8728596992, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728800320}
	ptr8728596864 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Description: "Automatically created basic rule for imports", Rules: []system.Rule(nil)}
	ptr8728798640 := &system.Type{Base: ptr8728596864, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728595200 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Description: "This is the most basic type.", Rules: []system.Rule(nil)}
	ptr8728595328 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "Type of the object.", Rules: []system.Rule(nil)}
	ptr8728911488 := &system.Reference_rule{Base: ptr8728595328, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728595456 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "All global objects should have an id.", Rules: []system.Rule(nil)}
	ptr8728923904 := &system.RuleBase{Optional: true}
	ptr8728911552 := &system.Reference_rule{Base: ptr8728595456, RuleBase: ptr8728923904, Default: system.Reference{}}
	ptr8728595584 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Description for the developer", Rules: []system.Rule(nil)}
	ptr8728924224 := &system.RuleBase{Optional: true}
	ptr8728997424 := &system.JsonString_rule{Base: ptr8728595584, RuleBase: ptr8728924224}
	ptr8728595712 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Extra validation rules for this object or descendants", Rules: []system.Rule(nil)}
	ptr8728924512 := &system.RuleBase{Optional: true}
	ptr8728595840 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728997648 := &system.Rule_rule{Base: ptr8728595840, RuleBase: (*system.RuleBase)(nil)}
	ptr8728723856 := &system.Array_rule{Base: ptr8728595712, RuleBase: ptr8728924512, Items: ptr8728997648, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728797856 := &system.Type{Base: ptr8728595200, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"type": ptr8728911488, "id": ptr8728911552, "description": ptr8728997424, "rules": ptr8728723856}, Rule: (*system.Type)(nil), Basic: true}
	ptr8728598656 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Rules: []system.Rule(nil)}
	ptr8728799200 := &system.Type{Base: ptr8728598656, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728799536}
	system.RegisterType("kego.io/system", "@reference", ptr8728799536)
	system.RegisterType("kego.io/system", "@int", ptr8728800320)
	system.RegisterType("kego.io/system", "@array", ptr8728797632)
	system.RegisterType("kego.io/system", "base", ptr8728797856)
	system.RegisterType("kego.io/system", "bool", ptr8728798304)
	system.RegisterType("kego.io/system", "map", ptr8728797184)
	system.RegisterType("kego.io/system", "int", ptr8728798864)
	system.RegisterType("kego.io/system", "@type", ptr8728800880)
	system.RegisterType("kego.io/system", "number", ptr8728799088)
	system.RegisterType("kego.io/system", "rule", ptr8728799648)
	system.RegisterType("kego.io/system", "imports", ptr8728798528)
	system.RegisterType("kego.io/system", "@imports", ptr8728798640)
	system.RegisterType("kego.io/system", "reference", ptr8728799200)
	system.RegisterType("kego.io/system", "type", ptr8728803120)
	system.RegisterType("kego.io/system", "ruleBase", ptr8728799872)
	system.RegisterType("kego.io/system", "@string", ptr8728802448)
	system.RegisterType("kego.io/system", "@map", ptr8728797408)
	system.RegisterType("kego.io/system", "@base", ptr8728798080)
	system.RegisterType("kego.io/system", "string", ptr8728800208)
	system.RegisterType("kego.io/system", "@number", ptr8728800656)
	system.RegisterType("kego.io/system", "@bool", ptr8728798416)
	system.RegisterType("kego.io/system", "@ruleBase", ptr8728800096)
	system.RegisterType("kego.io/system", "array", ptr8728797520)
	system.RegisterType("kego.io/system", "@rule", ptr8728799760)
}
