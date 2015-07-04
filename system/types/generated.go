package types

import (
	system "kego.io/system"
)

func init() {
	ptr8728722688 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Description: "This is the native json string data type", Rules: []system.Rule(nil)}
	ptr8728722816 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Description: "Restriction rules for strings", Rules: []system.Rule(nil)}
	ptr8729239808 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is a regex to match the value to", Rules: []system.Rule(nil)}
	ptr8728527584 := &system.RuleBase{Optional: true}
	ptr8729002688 := &system.String_rule{Base: ptr8729239808, RuleBase: ptr8728527584, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729239936 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This restricts the value to one of several built-in formats", Rules: []system.Rule(nil)}
	ptr8728528000 := &system.RuleBase{Optional: true}
	ptr8729002864 := &system.String_rule{Base: ptr8729239936, RuleBase: ptr8728528000, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728722944 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Default value of this is missing or null", Rules: []system.Rule(nil)}
	ptr8729107648 := &system.RuleBase{Optional: true}
	ptr8729002160 := &system.String_rule{Base: ptr8728722944, RuleBase: ptr8729107648, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728723072 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "The value of this string is restricted to one of the provided values", Rules: []system.Rule(nil)}
	ptr8729107968 := &system.RuleBase{Optional: true}
	ptr8728723200 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729187888 := &system.JsonString_rule{Base: ptr8728723200, RuleBase: (*system.RuleBase)(nil)}
	ptr8728863840 := &system.Array_rule{Base: ptr8728723072, RuleBase: ptr8729107968, Items: ptr8729187888, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728723328 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "The value must be longer or equal to the provided minimum length", Rules: []system.Rule(nil)}
	ptr8729108192 := &system.RuleBase{Optional: true}
	ptr8728864000 := &system.Int_rule{Base: ptr8728723328, RuleBase: ptr8729108192, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr8729239552 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "The value must be shorter or equal to the provided maximum length", Rules: []system.Rule(nil)}
	ptr8729108416 := &system.RuleBase{Optional: true}
	ptr8728864080 := &system.Int_rule{Base: ptr8729239552, RuleBase: ptr8729108416, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr8729239680 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is a string that the value must match", Rules: []system.Rule(nil)}
	ptr8728526976 := &system.RuleBase{Optional: true}
	ptr8729002512 := &system.String_rule{Base: ptr8729239680, RuleBase: ptr8728526976, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728934416 := &system.Type{Base: ptr8728722816, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"pattern": ptr8729002688, "format": ptr8729002864, "default": ptr8729002160, "enum": ptr8728863840, "minLength": ptr8728864000, "maxLength": ptr8728864080, "equal": ptr8729002512}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728931392 := &system.Type{Base: ptr8728722688, Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728934416, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728719232 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Description: "Restriction rules for bools", Rules: []system.Rule(nil)}
	ptr8728719360 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Default value if this is missing or null", Rules: []system.Rule(nil)}
	ptr8729105760 := &system.RuleBase{Optional: true}
	ptr8729105696 := &system.Bool_rule{Base: ptr8728719360, RuleBase: ptr8729105760, Default: system.Bool{}}
	ptr8728929600 := &system.Type{Base: ptr8728719232, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729105696}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728717184 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Description: "This is the native json object data type.", Rules: []system.Rule(nil)}
	ptr8728717312 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Description: "Restriction rules for maps", Rules: []system.Rule(nil)}
	ptr8728720640 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the maximum number of items alowed in the array", Rules: []system.Rule(nil)}
	ptr8729102688 := &system.RuleBase{Optional: true}
	ptr8728863600 := &system.Int_rule{Base: ptr8728720640, RuleBase: ptr8729102688, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr8728720384 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Rules: []system.Rule(nil)}
	ptr8729183856 := &system.Rule_rule{Base: ptr8728720384, RuleBase: (*system.RuleBase)(nil)}
	ptr8728720512 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the minimum number of items alowed in the array", Rules: []system.Rule(nil)}
	ptr8729102400 := &system.RuleBase{Optional: true}
	ptr8728863520 := &system.Int_rule{Base: ptr8728720512, RuleBase: ptr8729102400, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr8728930160 := &system.Type{Base: ptr8728717312, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"maxItems": ptr8728863600, "items": ptr8729183856, "minItems": ptr8728863520}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728928480 := &system.Type{Base: ptr8728717184, Native: system.String{Value: "map", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728930160, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728720896 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Description: "Restriction rules for numbers", Rules: []system.Rule(nil)}
	ptr8728721280 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil)}
	ptr8729104576 := &system.RuleBase{Optional: true}
	ptr8728707552 := &system.Number_rule{Base: ptr8728721280, RuleBase: ptr8729104576, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728721408 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Rules: []system.Rule(nil)}
	ptr8729104928 := &system.RuleBase{Optional: true}
	ptr8729185552 := &system.JsonBool_rule{Base: ptr8728721408, RuleBase: ptr8729104928}
	ptr8728720128 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "Default value if this property is omitted", Rules: []system.Rule(nil)}
	ptr8729103488 := &system.RuleBase{Optional: true}
	ptr8728707168 := &system.Number_rule{Base: ptr8728720128, RuleBase: ptr8729103488, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728720256 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil)}
	ptr8729103744 := &system.RuleBase{Optional: true}
	ptr8728707840 := &system.Number_rule{Base: ptr8728720256, RuleBase: ptr8729103744, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728721024 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil)}
	ptr8729104000 := &system.RuleBase{Optional: true}
	ptr8728707360 := &system.Number_rule{Base: ptr8728721024, RuleBase: ptr8729104000, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}, Maximum: system.Number{}}
	ptr8728721152 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Rules: []system.Rule(nil)}
	ptr8729104320 := &system.RuleBase{Optional: true}
	ptr8729185136 := &system.JsonBool_rule{Base: ptr8728721152, RuleBase: ptr8729104320}
	ptr8728932176 := &system.Type{Base: ptr8728720896, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"maximum": ptr8728707552, "exclusiveMaximum": ptr8729185552, "default": ptr8728707168, "multipleOf": ptr8728707840, "minimum": ptr8728707360, "exclusiveMinimum": ptr8729185136}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728717440 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Description: "This is the native json array data type", Rules: []system.Rule(nil)}
	ptr8728717568 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Description: "Restriction rules for arrays", Rules: []system.Rule(nil)}
	ptr8728717696 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Description: "This is a rule object, defining the type and restrictions on the value of the items", Rules: []system.Rule(nil)}
	ptr8729184864 := &system.Rule_rule{Base: ptr8728717696, RuleBase: (*system.RuleBase)(nil)}
	ptr8728717824 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the minimum number of items allowed in the array", Rules: []system.Rule(nil)}
	ptr8729103104 := &system.RuleBase{Optional: true}
	ptr8728862960 := &system.Int_rule{Base: ptr8728717824, RuleBase: ptr8729103104, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr8728717952 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the maximum number of items allowed in the array", Rules: []system.Rule(nil)}
	ptr8729103360 := &system.RuleBase{Optional: true}
	ptr8728863040 := &system.Int_rule{Base: ptr8728717952, RuleBase: ptr8729103360, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr8728718080 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, each item must be unique", Rules: []system.Rule(nil)}
	ptr8729103648 := &system.RuleBase{Optional: true}
	ptr8729185296 := &system.JsonBool_rule{Base: ptr8728718080, RuleBase: ptr8729103648}
	ptr8728928816 := &system.Type{Base: ptr8728717568, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729184864, "minItems": ptr8728862960, "maxItems": ptr8728863040, "uniqueItems": ptr8729185296}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728928704 := &system.Type{Base: ptr8728717440, Native: system.String{Value: "array", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728928816, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728719104 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Description: "This is the native json bool data type", Rules: []system.Rule(nil)}
	ptr8728929488 := &system.Type{Base: ptr8728719104, Native: system.String{Value: "bool", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728929600, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728719488 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Description: "Lists imports used in this package.", Rules: []system.Rule(nil)}
	ptr8728719616 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "Map of path to alias.", Rules: []system.Rule(nil)}
	ptr8728719744 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729002336 := &system.String_rule{Base: ptr8728719744, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728702336 := &system.Map_rule{Base: ptr8728719616, RuleBase: (*system.RuleBase)(nil), Items: ptr8729002336, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728929712 := &system.Type{Base: ptr8728719488, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8728702336}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728715264 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Description: "Restriction rules for integers", Rules: []system.Rule(nil)}
	ptr8728716672 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil)}
	ptr8729100992 := &system.RuleBase{Optional: true}
	ptr8728863280 := &system.Int_rule{Base: ptr8728716672, RuleBase: ptr8729100992, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr8728716928 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil)}
	ptr8729101216 := &system.RuleBase{Optional: true}
	ptr8728863360 := &system.Int_rule{Base: ptr8728716928, RuleBase: ptr8729101216, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr8728717056 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil)}
	ptr8729101440 := &system.RuleBase{Optional: true}
	ptr8728863440 := &system.Int_rule{Base: ptr8728717056, RuleBase: ptr8729101440, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr8728716544 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "Default value if this property is omitted", Rules: []system.Rule(nil)}
	ptr8729100768 := &system.RuleBase{Optional: true}
	ptr8728863200 := &system.Int_rule{Base: ptr8728716544, RuleBase: ptr8729100768, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr8728928256 := &system.Type{Base: ptr8728715264, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"multipleOf": ptr8728863280, "minimum": ptr8728863360, "maximum": ptr8728863440, "default": ptr8728863200}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728722560 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Description: "Automatically created basic rule for ruleBase", Rules: []system.Rule(nil)}
	ptr8728931168 := &system.Type{Base: ptr8728722560, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728719872 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Description: "Automatically created basic rule for imports", Rules: []system.Rule(nil)}
	ptr8728929824 := &system.Type{Base: ptr8728719872, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728721536 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Rules: []system.Rule(nil)}
	ptr8728721664 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Description: "Restriction rules for references", Rules: []system.Rule(nil)}
	ptr8728721792 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "Default value of this is missing or null", Rules: []system.Rule(nil)}
	ptr8729105792 := &system.RuleBase{Optional: true}
	ptr8728702208 := &system.Reference_rule{Base: ptr8728721792, RuleBase: ptr8729105792, Default: system.Reference{}}
	ptr8728930608 := &system.Type{Base: ptr8728721664, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728702208}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728930272 := &system.Type{Base: ptr8728721536, Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728930608, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728722176 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Description: "All rules should embed this type.", Rules: []system.Rule(nil)}
	ptr8728722304 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this rule is a field, this specifies that the field is optional", Rules: []system.Rule(nil)}
	ptr8729106624 := &system.RuleBase{Optional: true}
	ptr8729186896 := &system.JsonBool_rule{Base: ptr8728722304, RuleBase: ptr8729106624}
	ptr8728722432 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Json selector defining what nodes this rule should be applied to.", Rules: []system.Rule(nil)}
	ptr8729106912 := &system.RuleBase{Optional: true}
	ptr8729187088 := &system.JsonString_rule{Base: ptr8728722432, RuleBase: ptr8729106912}
	ptr8728930944 := &system.Type{Base: ptr8728722176, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"optional": ptr8729186896, "selector": ptr8729187088}, Rule: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8729240064 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Description: "This is the most basic type.", Rules: []system.Rule(nil)}
	ptr8729241472 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Id: system.Reference{}, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Rules: []system.Rule(nil)}
	ptr8728534816 := &system.RuleBase{Optional: true}
	ptr8729183824 := &system.Type_rule{Base: ptr8729241472, RuleBase: ptr8728534816}
	ptr8729240192 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Basic types don't have system:object added by default to the embedded types.", Rules: []system.Rule(nil)}
	ptr8728528864 := &system.RuleBase{Optional: true}
	ptr8729188944 := &system.JsonBool_rule{Base: ptr8729240192, RuleBase: ptr8728528864}
	ptr8729240320 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Types which this should embed - system:object is always added unless base = true.", Rules: []system.Rule(nil)}
	ptr8728529184 := &system.RuleBase{Optional: true}
	ptr8729240448 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728703104 := &system.Reference_rule{Base: ptr8729240448, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728864320 := &system.Array_rule{Base: ptr8729240320, RuleBase: ptr8728529184, Items: ptr8728703104, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729240576 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Array of interface types that this type should support", Rules: []system.Rule(nil)}
	ptr8728526944 := &system.RuleBase{Optional: true}
	ptr8729240704 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8728703360 := &system.Reference_rule{Base: ptr8729240704, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728864400 := &system.Array_rule{Base: ptr8729240576, RuleBase: ptr8728526944, Items: ptr8728703360, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729240832 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is the native json type that represents this type. If omitted, default is object.", Rules: []system.Rule(nil)}
	ptr8728527776 := &system.RuleBase{Optional: true}
	ptr8729001984 := &system.String_rule{Base: ptr8729240832, RuleBase: ptr8728527776, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729240960 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Is this type an interface?", Rules: []system.Rule(nil)}
	ptr8728528352 := &system.RuleBase{Optional: true}
	ptr8729182656 := &system.JsonBool_rule{Base: ptr8729240960, RuleBase: ptr8728528352}
	ptr8729241088 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Exclude this type from the generated json?", Rules: []system.Rule(nil)}
	ptr8728528896 := &system.RuleBase{Optional: true}
	ptr8729182848 := &system.JsonBool_rule{Base: ptr8729241088, RuleBase: ptr8728528896}
	ptr8729241216 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "Each field is listed with it's type", Rules: []system.Rule(nil)}
	ptr8728529248 := &system.RuleBase{Optional: true}
	ptr8729241344 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729183472 := &system.Rule_rule{Base: ptr8729241344, RuleBase: (*system.RuleBase)(nil)}
	ptr8728699456 := &system.Map_rule{Base: ptr8729241216, RuleBase: ptr8728529248, Items: ptr8729183472, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728934976 := &system.Type{Base: ptr8729240064, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"rule": ptr8729183824, "basic": ptr8729188944, "embed": ptr8728864320, "is": ptr8728864400, "native": ptr8729001984, "interface": ptr8729182656, "exclude": ptr8729182848, "fields": ptr8728699456}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728720768 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Description: "This is the native json number data type", Rules: []system.Rule(nil)}
	ptr8728931616 := &system.Type{Base: ptr8728720768, Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728932176, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728720000 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Description: "This is the int data type", Rules: []system.Rule(nil)}
	ptr8728929936 := &system.Type{Base: ptr8728720000, Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728928256, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8729241600 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Description: "Automatically created basic rule for type", Rules: []system.Rule(nil)}
	ptr8728932400 := &system.Type{Base: ptr8729241600, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728718976 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Description: "Automatically created basic rule for base", Rules: []system.Rule(nil)}
	ptr8728929264 := &system.Type{Base: ptr8728718976, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728722048 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Description: "Automatically created basic rule for rule", Rules: []system.Rule(nil)}
	ptr8728930832 := &system.Type{Base: ptr8728722048, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728721920 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Description: "This is one of several rule types, derived from the rules property of other types", Rules: []system.Rule(nil)}
	ptr8728930720 := &system.Type{Base: ptr8728721920, Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728718208 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Description: "This is the most basic type.", Rules: []system.Rule(nil)}
	ptr8728718336 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "Type of the object.", Rules: []system.Rule(nil)}
	ptr8729083520 := &system.Reference_rule{Base: ptr8728718336, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728718464 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "All global objects should have an id.", Rules: []system.Rule(nil)}
	ptr8729104384 := &system.RuleBase{Optional: true}
	ptr8729083584 := &system.Reference_rule{Base: ptr8728718464, RuleBase: ptr8729104384, Default: system.Reference{}}
	ptr8728718592 := &system.Base{Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Description for the developer", Rules: []system.Rule(nil)}
	ptr8729104704 := &system.RuleBase{Optional: true}
	ptr8729185952 := &system.JsonString_rule{Base: ptr8728718592, RuleBase: ptr8729104704}
	ptr8728718720 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Extra validation rules for this object or descendants", Rules: []system.Rule(nil)}
	ptr8729104992 := &system.RuleBase{Optional: true}
	ptr8728718848 := &system.Base{Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Rules: []system.Rule(nil)}
	ptr8729186176 := &system.Rule_rule{Base: ptr8728718848, RuleBase: (*system.RuleBase)(nil)}
	ptr8728863120 := &system.Array_rule{Base: ptr8728718720, RuleBase: ptr8729104992, Items: ptr8729186176, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728929040 := &system.Type{Base: ptr8728718208, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"type": ptr8729083520, "id": ptr8729083584, "description": ptr8729185952, "rules": ptr8728863120}, Rule: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	system.RegisterType("kego.io/system", "@base", ptr8728929264)
	system.RegisterType("kego.io/system", "@map", ptr8728930160)
	system.RegisterType("kego.io/system", "@bool", ptr8728929600)
	system.RegisterType("kego.io/system", "reference", ptr8728930272)
	system.RegisterType("kego.io/system", "@imports", ptr8728929824)
	system.RegisterType("kego.io/system", "ruleBase", ptr8728930944)
	system.RegisterType("kego.io/system", "type", ptr8728934976)
	system.RegisterType("kego.io/system", "number", ptr8728931616)
	system.RegisterType("kego.io/system", "rule", ptr8728930720)
	system.RegisterType("kego.io/system", "@number", ptr8728932176)
	system.RegisterType("kego.io/system", "@int", ptr8728928256)
	system.RegisterType("kego.io/system", "@ruleBase", ptr8728931168)
	system.RegisterType("kego.io/system", "int", ptr8728929936)
	system.RegisterType("kego.io/system", "@rule", ptr8728930832)
	system.RegisterType("kego.io/system", "base", ptr8728929040)
	system.RegisterType("kego.io/system", "array", ptr8728928704)
	system.RegisterType("kego.io/system", "imports", ptr8728929712)
	system.RegisterType("kego.io/system", "bool", ptr8728929488)
	system.RegisterType("kego.io/system", "@string", ptr8728934416)
	system.RegisterType("kego.io/system", "@type", ptr8728932400)
	system.RegisterType("kego.io/system", "@array", ptr8728928816)
	system.RegisterType("kego.io/system", "@reference", ptr8728930608)
	system.RegisterType("kego.io/system", "string", ptr8728931392)
	system.RegisterType("kego.io/system", "map", ptr8728928480)
}
