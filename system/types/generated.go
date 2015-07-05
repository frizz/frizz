package types

import (
	system "kego.io/system"
)

func init() {
	ptr8728586112 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Description: "Automatically created basic rule for imports", Rules: []system.Rule(nil)}
	ptr8728797520 := &system.Type{Base: ptr8728586112, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728586240 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Description: "This is the int data type", Rules: []system.Rule(nil)}
	ptr8728586368 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Description: "Restriction rules for integers", Rules: []system.Rule(nil)}
	ptr8728586496 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "Default value if this property is omitted", Rules: []system.Rule(nil)}
	ptr8729045440 := &system.RuleBase{Optional: true}
	ptr8729068400 := &system.Int_rule{Base: ptr8728586496, RuleBase: ptr8729045440, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728586624 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil)}
	ptr8729045760 := &system.RuleBase{Optional: true}
	ptr8729068480 := &system.Int_rule{Base: ptr8728586624, RuleBase: ptr8729045760, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728586752 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil)}
	ptr8729046080 := &system.RuleBase{Optional: true}
	ptr8729068560 := &system.Int_rule{Base: ptr8728586752, RuleBase: ptr8729046080, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728586880 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil)}
	ptr8729046400 := &system.RuleBase{Optional: true}
	ptr8729068640 := &system.Int_rule{Base: ptr8728586880, RuleBase: ptr8729046400, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728799200 := &system.Type{Base: ptr8728586368, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729068400, "multipleOf": ptr8729068480, "minimum": ptr8729068560, "maximum": ptr8729068640}, Rule: (*system.Type)(nil)}
	ptr8728799088 := &system.Type{Base: ptr8728586240, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728799200}
	ptr8728588928 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Description: "Restriction rules for maps", Rules: []system.Rule(nil)}
	ptr8728589056 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Rules: []system.Rule(nil)}
	ptr8729013696 := &system.Rule_rule{Base: ptr8728589056, RuleBase: (*system.RuleBase)(nil)}
	ptr8728589184 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the minimum number of items alowed in the array", Rules: []system.Rule(nil)}
	ptr8729047712 := &system.RuleBase{Optional: true}
	ptr8729068720 := &system.Int_rule{Base: ptr8728589184, RuleBase: ptr8729047712, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728589312 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the maximum number of items alowed in the array", Rules: []system.Rule(nil)}
	ptr8729048096 := &system.RuleBase{Optional: true}
	ptr8729068800 := &system.Int_rule{Base: ptr8728589312, RuleBase: ptr8729048096, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728798640 := &system.Type{Base: ptr8728588928, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729013696, "minItems": ptr8729068720, "maxItems": ptr8729068800}, Rule: (*system.Type)(nil)}
	ptr8728587008 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Description: "This is the native json array data type", Rules: []system.Rule(nil)}
	ptr8728587136 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Description: "Restriction rules for arrays", Rules: []system.Rule(nil)}
	ptr8728587264 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Description: "This is a rule object, defining the type and restrictions on the value of the items", Rules: []system.Rule(nil)}
	ptr8729014768 := &system.Rule_rule{Base: ptr8728587264, RuleBase: (*system.RuleBase)(nil)}
	ptr8728587392 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the minimum number of items allowed in the array", Rules: []system.Rule(nil)}
	ptr8728401920 := &system.RuleBase{Optional: true}
	ptr8729068160 := &system.Int_rule{Base: ptr8728587392, RuleBase: ptr8728401920, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728587520 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the maximum number of items allowed in the array", Rules: []system.Rule(nil)}
	ptr8728402400 := &system.RuleBase{Optional: true}
	ptr8729068240 := &system.Int_rule{Base: ptr8728587520, RuleBase: ptr8728402400, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728587648 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, each item must be unique", Rules: []system.Rule(nil)}
	ptr8728402816 := &system.RuleBase{Optional: true}
	ptr8729015504 := &system.JsonBool_rule{Base: ptr8728587648, RuleBase: ptr8728402816}
	ptr8728797856 := &system.Type{Base: ptr8728587136, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729014768, "minItems": ptr8729068160, "maxItems": ptr8729068240, "uniqueItems": ptr8729015504}, Rule: (*system.Type)(nil)}
	ptr8728797632 := &system.Type{Base: ptr8728587008, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728797856}
	ptr8728587776 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Description: "This is the most basic type.", Rules: []system.Rule(nil)}
	ptr8728587904 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "Type of the object.", Rules: []system.Rule(nil)}
	ptr8728595264 := &system.Reference_rule{Base: ptr8728587904, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728588032 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "All global objects should have an id.", Rules: []system.Rule(nil)}
	ptr8729043008 := &system.RuleBase{Optional: true}
	ptr8728595392 := &system.Reference_rule{Base: ptr8728588032, RuleBase: ptr8729043008, Default: system.Reference{}}
	ptr8728588160 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Description for the developer", Rules: []system.Rule(nil)}
	ptr8729043424 := &system.RuleBase{Optional: true}
	ptr8729016688 := &system.JsonString_rule{Base: ptr8728588160, RuleBase: ptr8729043424}
	ptr8728588288 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Extra validation rules for this object or descendants", Rules: []system.Rule(nil)}
	ptr8729043872 := &system.RuleBase{Optional: true}
	ptr8728588416 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729017056 := &system.Rule_rule{Base: ptr8728588416, RuleBase: (*system.RuleBase)(nil)}
	ptr8729068320 := &system.Array_rule{Base: ptr8728588288, RuleBase: ptr8729043872, MaxItems: system.Int{Value: 0}, Items: ptr8729017056, MinItems: system.Int{Value: 0}}
	ptr8728798080 := &system.Type{Base: ptr8728587776, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"type": ptr8728595264, "id": ptr8728595392, "description": ptr8729016688, "rules": ptr8729068320}, Rule: (*system.Type)(nil), Basic: true}
	ptr8728590976 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Description: "Automatically created basic rule for rule", Rules: []system.Rule(nil)}
	ptr8728799760 := &system.Type{Base: ptr8728590976, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728591616 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Description: "This is the native json string data type", Rules: []system.Rule(nil)}
	ptr8728591744 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Description: "Restriction rules for strings", Rules: []system.Rule(nil)}
	ptr8728592000 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "The value of this string is restricted to one of the provided values", Rules: []system.Rule(nil)}
	ptr8728396800 := &system.RuleBase{Optional: true}
	ptr8728592128 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728561040 := &system.JsonString_rule{Base: ptr8728592128, RuleBase: (*system.RuleBase)(nil)}
	ptr8729069040 := &system.Array_rule{Base: ptr8728592000, RuleBase: ptr8728396800, MaxItems: system.Int{Value: 0}, Items: ptr8728561040, MinItems: system.Int{Value: 0}}
	ptr8728592256 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "The value must be longer or equal to the provided minimum length", Rules: []system.Rule(nil)}
	ptr8728397280 := &system.RuleBase{Optional: true}
	ptr8729069200 := &system.Int_rule{Base: ptr8728592256, RuleBase: ptr8728397280, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728821760 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "The value must be shorter or equal to the provided maximum length", Rules: []system.Rule(nil)}
	ptr8728397792 := &system.RuleBase{Optional: true}
	ptr8729069280 := &system.Int_rule{Base: ptr8728821760, RuleBase: ptr8728397792, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728821888 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is a string that the value must match", Rules: []system.Rule(nil)}
	ptr8728398144 := &system.RuleBase{Optional: true}
	ptr8728993968 := &system.String_rule{Base: ptr8728821888, RuleBase: ptr8728398144, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728822016 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is a regex to match the value to", Rules: []system.Rule(nil)}
	ptr8728398496 := &system.RuleBase{Optional: true}
	ptr8728994320 := &system.String_rule{Base: ptr8728822016, RuleBase: ptr8728398496, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728822144 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This restricts the value to one of several built-in formats", Rules: []system.Rule(nil)}
	ptr8728398944 := &system.RuleBase{Optional: true}
	ptr8728994496 := &system.String_rule{Base: ptr8728822144, RuleBase: ptr8728398944, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728591872 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Default value of this is missing or null", Rules: []system.Rule(nil)}
	ptr8728395904 := &system.RuleBase{Optional: true}
	ptr8728994144 := &system.String_rule{Base: ptr8728591872, RuleBase: ptr8728395904, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728800320 := &system.Type{Base: ptr8728591744, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"enum": ptr8729069040, "minLength": ptr8729069200, "maxLength": ptr8729069280, "equal": ptr8728993968, "pattern": ptr8728994320, "format": ptr8728994496, "default": ptr8728994144}, Rule: (*system.Type)(nil)}
	ptr8728803120 := &system.Type{Base: ptr8728591616, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728800320}
	ptr8728585472 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Description: "Restriction rules for bools", Rules: []system.Rule(nil)}
	ptr8728585600 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Default value if this is missing or null", Rules: []system.Rule(nil)}
	ptr8729043616 := &system.RuleBase{Optional: true}
	ptr8729043552 := &system.Bool_rule{Base: ptr8728585600, RuleBase: ptr8729043616, Default: system.Bool{}}
	ptr8728797184 := &system.Type{Base: ptr8728585472, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729043552}, Rule: (*system.Type)(nil)}
	ptr8728823808 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Description: "Automatically created basic rule for type", Rules: []system.Rule(nil)}
	ptr8728800880 := &system.Type{Base: ptr8728823808, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728590848 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Description: "This is one of several rule types, derived from the rules property of other types", Rules: []system.Rule(nil)}
	ptr8728799648 := &system.Type{Base: ptr8728590848, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728588544 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Description: "Automatically created basic rule for base", Rules: []system.Rule(nil)}
	ptr8728798304 := &system.Type{Base: ptr8728588544, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728589440 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Description: "This is the native json number data type", Rules: []system.Rule(nil)}
	ptr8728589568 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Description: "Restriction rules for numbers", Rules: []system.Rule(nil)}
	ptr8728589952 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil)}
	ptr8729049728 := &system.RuleBase{Optional: true}
	ptr8728568096 := &system.Number_rule{Base: ptr8728589952, RuleBase: ptr8729049728, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728590080 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Rules: []system.Rule(nil)}
	ptr8729050144 := &system.RuleBase{Optional: true}
	ptr8729015792 := &system.JsonBool_rule{Base: ptr8728590080, RuleBase: ptr8729050144}
	ptr8728590208 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil)}
	ptr8729050496 := &system.RuleBase{Optional: true}
	ptr8728568384 := &system.Number_rule{Base: ptr8728590208, RuleBase: ptr8729050496, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728590336 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Rules: []system.Rule(nil)}
	ptr8729050912 := &system.RuleBase{Optional: true}
	ptr8729016304 := &system.JsonBool_rule{Base: ptr8728590336, RuleBase: ptr8729050912}
	ptr8728589696 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "Default value if this property is omitted", Rules: []system.Rule(nil)}
	ptr8729049088 := &system.RuleBase{Optional: true}
	ptr8728567904 := &system.Number_rule{Base: ptr8728589696, RuleBase: ptr8729049088, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728589824 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil)}
	ptr8729049408 := &system.RuleBase{Optional: true}
	ptr8728568000 := &system.Number_rule{Base: ptr8728589824, RuleBase: ptr8729049408, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728800096 := &system.Type{Base: ptr8728589568, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"minimum": ptr8728568096, "exclusiveMinimum": ptr8729015792, "maximum": ptr8728568384, "exclusiveMaximum": ptr8729016304, "default": ptr8728567904, "multipleOf": ptr8728568000}, Rule: (*system.Type)(nil)}
	ptr8728798864 := &system.Type{Base: ptr8728589440, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728800096}
	ptr8728822272 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Description: "This is the most basic type.", Rules: []system.Rule(nil)}
	ptr8728823168 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Is this type an interface?", Rules: []system.Rule(nil)}
	ptr8728401664 := &system.RuleBase{Optional: true}
	ptr8728566496 := &system.JsonBool_rule{Base: ptr8728823168, RuleBase: ptr8728401664}
	ptr8728823296 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Exclude this type from the generated json?", Rules: []system.Rule(nil)}
	ptr8728402080 := &system.RuleBase{Optional: true}
	ptr8728567024 := &system.JsonBool_rule{Base: ptr8728823296, RuleBase: ptr8728402080}
	ptr8728823424 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "Each field is listed with it's type", Rules: []system.Rule(nil)}
	ptr8728402560 := &system.RuleBase{Optional: true}
	ptr8728823552 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728567744 := &system.Rule_rule{Base: ptr8728823552, RuleBase: (*system.RuleBase)(nil)}
	ptr8728595904 := &system.Map_rule{Base: ptr8728823424, RuleBase: ptr8728402560, Items: ptr8728567744, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728823680 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Id: system.Reference{}, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Rules: []system.Rule(nil)}
	ptr8728402912 := &system.RuleBase{Optional: true}
	ptr8729010384 := &system.Type_rule{Base: ptr8728823680, RuleBase: ptr8728402912}
	ptr8728822400 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Basic types don't have system:object added by default to the embedded types.", Rules: []system.Rule(nil)}
	ptr8728399840 := &system.RuleBase{Optional: true}
	ptr8728564752 := &system.JsonBool_rule{Base: ptr8728822400, RuleBase: ptr8728399840}
	ptr8728822528 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Types which this should embed - system:object is always added unless base = true.", Rules: []system.Rule(nil)}
	ptr8728400288 := &system.RuleBase{Optional: true}
	ptr8728822656 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728595648 := &system.Reference_rule{Base: ptr8728822656, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8729069520 := &system.Array_rule{Base: ptr8728822528, RuleBase: ptr8728400288, MaxItems: system.Int{Value: 0}, Items: ptr8728595648, MinItems: system.Int{Value: 0}}
	ptr8728822784 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Array of interface types that this type should support", Rules: []system.Rule(nil)}
	ptr8728400736 := &system.RuleBase{Optional: true}
	ptr8728822912 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728595776 := &system.Reference_rule{Base: ptr8728822912, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8729069600 := &system.Array_rule{Base: ptr8728822784, RuleBase: ptr8728400736, MaxItems: system.Int{Value: 0}, Items: ptr8728595776, MinItems: system.Int{Value: 0}}
	ptr8728823040 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is the native json type that represents this type. If omitted, default is object.", Rules: []system.Rule(nil)}
	ptr8728401248 := &system.RuleBase{Optional: true}
	ptr8728994672 := &system.String_rule{Base: ptr8728823040, RuleBase: ptr8728401248, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728800656 := &system.Type{Base: ptr8728822272, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"interface": ptr8728566496, "exclude": ptr8728567024, "fields": ptr8728595904, "rule": ptr8729010384, "basic": ptr8728564752, "embed": ptr8729069520, "is": ptr8729069600, "native": ptr8728994672}, Rule: (*system.Type)(nil)}
	ptr8728590592 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Description: "Restriction rules for references", Rules: []system.Rule(nil)}
	ptr8728590720 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "Default value of this is missing or null", Rules: []system.Rule(nil)}
	ptr8728397152 := &system.RuleBase{Optional: true}
	ptr8728597888 := &system.Reference_rule{Base: ptr8728590720, RuleBase: ptr8728397152, Default: system.Reference{}}
	ptr8728799536 := &system.Type{Base: ptr8728590592, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728597888}, Rule: (*system.Type)(nil)}
	ptr8728588672 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Description: "This is the native json bool data type", Rules: []system.Rule(nil)}
	ptr8728798416 := &system.Type{Base: ptr8728588672, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728797184}
	ptr8728588800 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Description: "This is the native json object data type.", Rules: []system.Rule(nil)}
	ptr8728798528 := &system.Type{Base: ptr8728588800, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728798640}
	ptr8728591488 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Description: "Automatically created basic rule for ruleBase", Rules: []system.Rule(nil)}
	ptr8728802448 := &system.Type{Base: ptr8728591488, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728591104 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Description: "All rules should embed this type.", Rules: []system.Rule(nil)}
	ptr8728591232 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this rule is a field, this specifies that the field is optional", Rules: []system.Rule(nil)}
	ptr8728398464 := &system.RuleBase{Optional: true}
	ptr8729018144 := &system.JsonBool_rule{Base: ptr8728591232, RuleBase: ptr8728398464}
	ptr8728591360 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Json selector defining what nodes this rule should be applied to.", Rules: []system.Rule(nil)}
	ptr8728398848 := &system.RuleBase{Optional: true}
	ptr8728560144 := &system.JsonString_rule{Base: ptr8728591360, RuleBase: ptr8728398848}
	ptr8728799872 := &system.Type{Base: ptr8728591104, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8729018144, "selector": ptr8728560144}, Rule: (*system.Type)(nil), Basic: true}
	ptr8728585728 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Description: "Lists imports used in this package.", Rules: []system.Rule(nil)}
	ptr8728585856 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "Map of path to alias.", Rules: []system.Rule(nil)}
	ptr8728585984 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728993792 := &system.String_rule{Base: ptr8728585984, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728594880 := &system.Map_rule{Base: ptr8728585856, RuleBase: (*system.RuleBase)(nil), Items: ptr8728993792, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728797408 := &system.Type{Base: ptr8728585728, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8728594880}, Rule: (*system.Type)(nil)}
	ptr8728590464 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Rules: []system.Rule(nil)}
	ptr8728800208 := &system.Type{Base: ptr8728590464, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728799536}
	system.RegisterType("kego.io/system", "@string", ptr8728800320)
	system.RegisterType("kego.io/system", "rule", ptr8728799648)
	system.RegisterType("kego.io/system", "number", ptr8728798864)
	system.RegisterType("kego.io/system", "@reference", ptr8728799536)
	system.RegisterType("kego.io/system", "@ruleBase", ptr8728802448)
	system.RegisterType("kego.io/system", "@imports", ptr8728797520)
	system.RegisterType("kego.io/system", "int", ptr8728799088)
	system.RegisterType("kego.io/system", "@rule", ptr8728799760)
	system.RegisterType("kego.io/system", "@array", ptr8728797856)
	system.RegisterType("kego.io/system", "@bool", ptr8728797184)
	system.RegisterType("kego.io/system", "@base", ptr8728798304)
	system.RegisterType("kego.io/system", "type", ptr8728800656)
	system.RegisterType("kego.io/system", "map", ptr8728798528)
	system.RegisterType("kego.io/system", "reference", ptr8728800208)
	system.RegisterType("kego.io/system", "ruleBase", ptr8728799872)
	system.RegisterType("kego.io/system", "@map", ptr8728798640)
	system.RegisterType("kego.io/system", "array", ptr8728797632)
	system.RegisterType("kego.io/system", "base", ptr8728798080)
	system.RegisterType("kego.io/system", "string", ptr8728803120)
	system.RegisterType("kego.io/system", "@type", ptr8728800880)
	system.RegisterType("kego.io/system", "bool", ptr8728798416)
	system.RegisterType("kego.io/system", "@int", ptr8728799200)
	system.RegisterType("kego.io/system", "@number", ptr8728800096)
	system.RegisterType("kego.io/system", "imports", ptr8728797408)
}
