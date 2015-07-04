package types

import (
	system "kego.io/system"
)

func init() {
	ptr8728254720 := &system.Base{Description: "Restriction rules for references", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr8728254848 := &system.Base{Description: "Default value of this is missing or null", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8728754208 := &system.RuleBase{Optional: true}
	ptr8728270144 := &system.Reference_rule{Base: ptr8728254848, RuleBase: ptr8728754208, Default: system.Reference{}}
	ptr8728480272 := &system.Type{Base: ptr8728254720, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728270144}, Rule: (*system.Type)(nil)}
	ptr8728255616 := &system.Base{Description: "Automatically created basic rule for ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}}
	ptr8728480832 := &system.Type{Base: ptr8728255616, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728255744 := &system.Base{Description: "This is the native json string data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}}
	ptr8728255872 := &system.Base{Description: "Restriction rules for strings", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr8728567936 := &system.Base{Description: "This is a string that the value must match", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728749184 := &system.RuleBase{Optional: true}
	ptr8728420528 := &system.String_rule{Base: ptr8728567936, RuleBase: ptr8728749184, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728568064 := &system.Base{Description: "This is a regex to match the value to", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728749440 := &system.RuleBase{Optional: true}
	ptr8728420880 := &system.String_rule{Base: ptr8728568064, RuleBase: ptr8728749440, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728568192 := &system.Base{Description: "This restricts the value to one of several built-in formats", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728749760 := &system.RuleBase{Optional: true}
	ptr8728421056 := &system.String_rule{Base: ptr8728568192, RuleBase: ptr8728749760, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728256000 := &system.Base{Description: "Default value of this is missing or null", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728748064 := &system.RuleBase{Optional: true}
	ptr8728420704 := &system.String_rule{Base: ptr8728256000, RuleBase: ptr8728748064, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728256128 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728748448 := &system.RuleBase{Optional: true}
	ptr8728256256 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728699120 := &system.JsonString_rule{Base: ptr8728256256, RuleBase: (*system.RuleBase)(nil)}
	ptr8728626672 := &system.Array_rule{Base: ptr8728256128, RuleBase: ptr8728748448, Items: ptr8728699120, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728256384 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8728748672 := &system.RuleBase{Optional: true}
	ptr8728626832 := &system.Int_rule{Base: ptr8728256384, RuleBase: ptr8728748672, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728567808 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8728748928 := &system.RuleBase{Optional: true}
	ptr8728626912 := &system.Int_rule{Base: ptr8728567808, RuleBase: ptr8728748928, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728481056 := &system.Type{Base: ptr8728255872, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"equal": ptr8728420528, "pattern": ptr8728420880, "format": ptr8728421056, "default": ptr8728420704, "enum": ptr8728626672, "minLength": ptr8728626832, "maxLength": ptr8728626912}, Rule: (*system.Type)(nil)}
	ptr8728480944 := &system.Type{Base: ptr8728255744, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728481056}
	ptr8728255104 := &system.Base{Description: "Automatically created basic rule for rule", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr8728480608 := &system.Type{Base: ptr8728255104, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728569856 := &system.Base{Description: "Automatically created basic rule for type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr8728484080 := &system.Type{Base: ptr8728569856, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728251648 := &system.Base{Description: "This is the native json number data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}}
	ptr8728251776 := &system.Base{Description: "Restriction rules for numbers", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr8728251904 := &system.Base{Description: "Default value if this property is omitted", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728752224 := &system.RuleBase{Optional: true}
	ptr8728272992 := &system.Number_rule{Base: ptr8728251904, RuleBase: ptr8728752224, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728252032 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728752448 := &system.RuleBase{Optional: true}
	ptr8728273088 := &system.Number_rule{Base: ptr8728252032, RuleBase: ptr8728752448, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728252160 := &system.Base{Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728752672 := &system.RuleBase{Optional: true}
	ptr8728273280 := &system.Number_rule{Base: ptr8728252160, RuleBase: ptr8728752672, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728252288 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728752960 := &system.RuleBase{Optional: true}
	ptr8728702432 := &system.JsonBool_rule{Base: ptr8728252288, RuleBase: ptr8728752960}
	ptr8728254336 := &system.Base{Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8728753184 := &system.RuleBase{Optional: true}
	ptr8728273472 := &system.Number_rule{Base: ptr8728254336, RuleBase: ptr8728753184, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728254464 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728753472 := &system.RuleBase{Optional: true}
	ptr8728702752 := &system.JsonBool_rule{Base: ptr8728254464, RuleBase: ptr8728753472}
	ptr8728479936 := &system.Type{Base: ptr8728251776, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728272992, "multipleOf": ptr8728273088, "minimum": ptr8728273280, "exclusiveMinimum": ptr8728702432, "maximum": ptr8728273472, "exclusiveMaximum": ptr8728702752}, Rule: (*system.Type)(nil)}
	ptr8728479712 := &system.Type{Base: ptr8728251648, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728479936}
	ptr8728255232 := &system.Base{Description: "All rules should embed this type.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}
	ptr8728255360 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728754944 := &system.RuleBase{Optional: true}
	ptr8728704144 := &system.JsonBool_rule{Base: ptr8728255360, RuleBase: ptr8728754944}
	ptr8728255488 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728755232 := &system.RuleBase{Optional: true}
	ptr8728704368 := &system.JsonString_rule{Base: ptr8728255488, RuleBase: ptr8728755232}
	ptr8728480720 := &system.Type{Base: ptr8728255232, Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8728704144, "selector": ptr8728704368}, Rule: (*system.Type)(nil)}
	ptr8728252544 := &system.Base{Description: "Restriction rules for arrays", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr8728252672 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8728702960 := &system.Rule_rule{Base: ptr8728252672, RuleBase: (*system.RuleBase)(nil)}
	ptr8728252800 := &system.Base{Description: "This is the minimum number of items allowed in the array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8728081984 := &system.RuleBase{Optional: true}
	ptr8728625792 := &system.Int_rule{Base: ptr8728252800, RuleBase: ptr8728081984, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728252928 := &system.Base{Description: "This is the maximum number of items allowed in the array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8728082240 := &system.RuleBase{Optional: true}
	ptr8728625872 := &system.Int_rule{Base: ptr8728252928, RuleBase: ptr8728082240, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728253056 := &system.Base{Description: "If this is true, each item must be unique", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728082528 := &system.RuleBase{Optional: true}
	ptr8728703392 := &system.JsonBool_rule{Base: ptr8728253056, RuleBase: ptr8728082528}
	ptr8728478368 := &system.Type{Base: ptr8728252544, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8728702960, "minItems": ptr8728625792, "maxItems": ptr8728625872, "uniqueItems": ptr8728703392}, Rule: (*system.Type)(nil)}
	ptr8728250368 := &system.Base{Description: "Restriction rules for integers", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr8728250624 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8728749824 := &system.RuleBase{Optional: true}
	ptr8728626112 := &system.Int_rule{Base: ptr8728250624, RuleBase: ptr8728749824, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728250752 := &system.Base{Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8728750048 := &system.RuleBase{Optional: true}
	ptr8728626192 := &system.Int_rule{Base: ptr8728250752, RuleBase: ptr8728750048, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728250880 := &system.Base{Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8728750272 := &system.RuleBase{Optional: true}
	ptr8728626272 := &system.Int_rule{Base: ptr8728250880, RuleBase: ptr8728750272, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728250496 := &system.Base{Description: "Default value if this property is omitted", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8728749600 := &system.RuleBase{Optional: true}
	ptr8728626032 := &system.Int_rule{Base: ptr8728250496, RuleBase: ptr8728749600, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728479152 := &system.Type{Base: ptr8728250368, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"multipleOf": ptr8728626112, "minimum": ptr8728626192, "maximum": ptr8728626272, "default": ptr8728626032}, Rule: (*system.Type)(nil)}
	ptr8728254976 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}
	ptr8728480384 := &system.Type{Base: ptr8728254976, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728253184 := &system.Base{Description: "This is the most basic type.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}}
	ptr8728253312 := &system.Base{Description: "Type of the object.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8728267648 := &system.Reference_rule{Base: ptr8728253312, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728253440 := &system.Base{Description: "All global objects should have an id.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8728083264 := &system.RuleBase{Optional: true}
	ptr8728267776 := &system.Reference_rule{Base: ptr8728253440, RuleBase: ptr8728083264, Default: system.Reference{}}
	ptr8728253568 := &system.Base{Description: "Description for the developer", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728083584 := &system.RuleBase{Optional: true}
	ptr8728704048 := &system.JsonString_rule{Base: ptr8728253568, RuleBase: ptr8728083584}
	ptr8728253696 := &system.Base{Description: "Extra validation rules for this object or descendants", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728083968 := &system.RuleBase{Optional: true}
	ptr8728253824 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8728704272 := &system.Rule_rule{Base: ptr8728253824, RuleBase: (*system.RuleBase)(nil)}
	ptr8728625952 := &system.Array_rule{Base: ptr8728253696, RuleBase: ptr8728083968, Items: ptr8728704272, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728478592 := &system.Type{Base: ptr8728253184, Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"type": ptr8728267648, "id": ptr8728267776, "description": ptr8728704048, "rules": ptr8728625952}, Rule: (*system.Type)(nil)}
	ptr8728251008 := &system.Base{Description: "This is the native json object data type.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}}
	ptr8728251136 := &system.Base{Description: "Restriction rules for maps", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr8728251520 := &system.Base{Description: "This is the maximum number of items alowed in the array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8728751488 := &system.RuleBase{Optional: true}
	ptr8728626432 := &system.Int_rule{Base: ptr8728251520, RuleBase: ptr8728751488, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728251264 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8728701088 := &system.Rule_rule{Base: ptr8728251264, RuleBase: (*system.RuleBase)(nil)}
	ptr8728251392 := &system.Base{Description: "This is the minimum number of items alowed in the array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8728751232 := &system.RuleBase{Optional: true}
	ptr8728626352 := &system.Int_rule{Base: ptr8728251392, RuleBase: ptr8728751232, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728479600 := &system.Type{Base: ptr8728251136, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"maxItems": ptr8728626432, "items": ptr8728701088, "minItems": ptr8728626352}, Rule: (*system.Type)(nil)}
	ptr8728479376 := &system.Type{Base: ptr8728251008, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728479600}
	ptr8728254592 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}}
	ptr8728480048 := &system.Type{Base: ptr8728254592, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728480272}
	ptr8728254080 := &system.Base{Description: "This is the native json bool data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}}
	ptr8728254208 := &system.Base{Description: "Restriction rules for bools", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr8728249600 := &system.Base{Description: "Default value if this is missing or null", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728748288 := &system.RuleBase{Optional: true}
	ptr8728748224 := &system.Bool_rule{Base: ptr8728249600, RuleBase: ptr8728748288, Default: system.Bool{}}
	ptr8728479040 := &system.Type{Base: ptr8728254208, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728748224}, Rule: (*system.Type)(nil)}
	ptr8728478928 := &system.Type{Base: ptr8728254080, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728479040}
	ptr8728250240 := &system.Base{Description: "This is the int data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}}
	ptr8728478032 := &system.Type{Base: ptr8728250240, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728479152}
	ptr8728568320 := &system.Base{Description: "This is the most basic type.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr8728568448 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728750432 := &system.RuleBase{Optional: true}
	ptr8728700768 := &system.JsonBool_rule{Base: ptr8728568448, RuleBase: ptr8728750432}
	ptr8728568576 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728750752 := &system.RuleBase{Optional: true}
	ptr8728568704 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8728267968 := &system.Reference_rule{Base: ptr8728568704, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728627152 := &system.Array_rule{Base: ptr8728568576, RuleBase: ptr8728750752, Items: ptr8728267968, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728568832 := &system.Base{Description: "Array of interface types that this type should support", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728751040 := &system.RuleBase{Optional: true}
	ptr8728568960 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8728268096 := &system.Reference_rule{Base: ptr8728568960, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728627232 := &system.Array_rule{Base: ptr8728568832, RuleBase: ptr8728751040, Items: ptr8728268096, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728569088 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728751392 := &system.RuleBase{Optional: true}
	ptr8728421232 := &system.String_rule{Base: ptr8728569088, RuleBase: ptr8728751392, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728569216 := &system.Base{Description: "Is this type an interface?", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728751744 := &system.RuleBase{Optional: true}
	ptr8728701872 := &system.JsonBool_rule{Base: ptr8728569216, RuleBase: ptr8728751744}
	ptr8728569344 := &system.Base{Description: "Exclude this type from the generated json?", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728752064 := &system.RuleBase{Optional: true}
	ptr8728702080 := &system.JsonBool_rule{Base: ptr8728569344, RuleBase: ptr8728752064}
	ptr8728569472 := &system.Base{Description: "Each field is listed with it's type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728752384 := &system.RuleBase{Optional: true}
	ptr8728569600 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8728702352 := &system.Rule_rule{Base: ptr8728569600, RuleBase: (*system.RuleBase)(nil)}
	ptr8728268224 := &system.Map_rule{Base: ptr8728569472, RuleBase: ptr8728752384, MaxItems: system.Int{Value: 0}, Items: ptr8728702352, MinItems: system.Int{Value: 0}}
	ptr8728569728 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Id: system.Reference{}}
	ptr8728752640 := &system.RuleBase{Optional: true}
	ptr8728702544 := &system.Type_rule{Base: ptr8728569728, RuleBase: ptr8728752640}
	ptr8728481168 := &system.Type{Base: ptr8728568320, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"basic": ptr8728700768, "embed": ptr8728627152, "is": ptr8728627232, "native": ptr8728421232, "interface": ptr8728701872, "exclude": ptr8728702080, "fields": ptr8728268224, "rule": ptr8728702544}, Rule: (*system.Type)(nil)}
	ptr8728253952 := &system.Base{Description: "Automatically created basic rule for base", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}}
	ptr8728478816 := &system.Type{Base: ptr8728253952, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728249728 := &system.Base{Description: "Lists imports used in this package.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}}
	ptr8728249856 := &system.Base{Description: "Map of import name to path.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728249984 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728420352 := &system.String_rule{Base: ptr8728249984, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728267200 := &system.Map_rule{Base: ptr8728249856, RuleBase: (*system.RuleBase)(nil), MaxItems: system.Int{Value: 0}, Items: ptr8728420352, MinItems: system.Int{Value: 0}}
	ptr8728477696 := &system.Type{Base: ptr8728249728, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8728267200}, Rule: (*system.Type)(nil)}
	ptr8728250112 := &system.Base{Description: "Automatically created basic rule for imports", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}}
	ptr8728477920 := &system.Type{Base: ptr8728250112, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil)}
	ptr8728252416 := &system.Base{Description: "This is the native json array data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}}
	ptr8728478144 := &system.Type{Base: ptr8728252416, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728478368}
	system.RegisterType("kego.io/system", "map", ptr8728479376)
	system.RegisterType("kego.io/system", "int", ptr8728478032)
	system.RegisterType("kego.io/system", "type", ptr8728481168)
	system.RegisterType("kego.io/system", "@number", ptr8728479936)
	system.RegisterType("kego.io/system", "imports", ptr8728477696)
	system.RegisterType("kego.io/system", "@imports", ptr8728477920)
	system.RegisterType("kego.io/system", "@bool", ptr8728479040)
	system.RegisterType("kego.io/system", "@map", ptr8728479600)
	system.RegisterType("kego.io/system", "@reference", ptr8728480272)
	system.RegisterType("kego.io/system", "@rule", ptr8728480608)
	system.RegisterType("kego.io/system", "bool", ptr8728478928)
	system.RegisterType("kego.io/system", "@int", ptr8728479152)
	system.RegisterType("kego.io/system", "rule", ptr8728480384)
	system.RegisterType("kego.io/system", "reference", ptr8728480048)
	system.RegisterType("kego.io/system", "@base", ptr8728478816)
	system.RegisterType("kego.io/system", "array", ptr8728478144)
	system.RegisterType("kego.io/system", "string", ptr8728480944)
	system.RegisterType("kego.io/system", "ruleBase", ptr8728480720)
	system.RegisterType("kego.io/system", "@array", ptr8728478368)
	system.RegisterType("kego.io/system", "base", ptr8728478592)
	system.RegisterType("kego.io/system", "@string", ptr8728481056)
	system.RegisterType("kego.io/system", "@ruleBase", ptr8728480832)
	system.RegisterType("kego.io/system", "@type", ptr8728484080)
	system.RegisterType("kego.io/system", "number", ptr8728479712)
}
