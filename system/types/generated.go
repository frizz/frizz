package types

import (
	system "kego.io/system"
)

func init() {
	ptr8728719232 := &system.Base{Description: "Restriction rules for bools", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr8728719360 := &system.Base{Description: "Default value if this is missing or null", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8729212256 := &system.RuleBase{Optional: true}
	ptr8729212192 := &system.Bool_rule{Base: ptr8728719360, RuleBase: ptr8729212256, Default: system.Bool{}}
	ptr8728962368 := &system.Type{Base: ptr8728719232, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8729212192}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728718208 := &system.Base{Description: "This is the most basic type.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}}
	ptr8728718336 := &system.Base{Description: "Type of the object.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8729083520 := &system.Reference_rule{Base: ptr8728718336, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728718464 := &system.Base{Description: "All global objects should have an id.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8729210880 := &system.RuleBase{Optional: true}
	ptr8729083584 := &system.Reference_rule{Base: ptr8728718464, RuleBase: ptr8729210880, Default: system.Reference{}}
	ptr8728718592 := &system.Base{Description: "Description for the developer", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8729211200 := &system.RuleBase{Optional: true}
	ptr8729185952 := &system.JsonString_rule{Base: ptr8728718592, RuleBase: ptr8729211200}
	ptr8728718720 := &system.Base{Description: "Extra validation rules for this object or descendants", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8729211488 := &system.RuleBase{Optional: true}
	ptr8728718848 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8729186176 := &system.Rule_rule{Base: ptr8728718848, RuleBase: (*system.RuleBase)(nil)}
	ptr8728895888 := &system.Array_rule{Base: ptr8728718720, RuleBase: ptr8729211488, Items: ptr8729186176, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728961808 := &system.Type{Base: ptr8728718208, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"type": ptr8729083520, "id": ptr8729083584, "description": ptr8729185952, "rules": ptr8728895888}, Rule: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728717184 := &system.Base{Description: "This is the native json object data type.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}}
	ptr8728717312 := &system.Base{Description: "Restriction rules for maps", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr8728720384 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8729183856 := &system.Rule_rule{Base: ptr8728720384, RuleBase: (*system.RuleBase)(nil)}
	ptr8728720512 := &system.Base{Description: "This is the minimum number of items alowed in the array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729208896 := &system.RuleBase{Optional: true}
	ptr8728896288 := &system.Int_rule{Base: ptr8728720512, RuleBase: ptr8729208896, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728720640 := &system.Base{Description: "This is the maximum number of items alowed in the array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729209184 := &system.RuleBase{Optional: true}
	ptr8728896368 := &system.Int_rule{Base: ptr8728720640, RuleBase: ptr8729209184, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728962928 := &system.Type{Base: ptr8728717312, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729183856, "minItems": ptr8728896288, "maxItems": ptr8728896368}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728961248 := &system.Type{Base: ptr8728717184, Native: system.String{Value: "map", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728962928, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728722816 := &system.Base{Description: "Restriction rules for strings", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr8729100416 := &system.Base{Description: "This is a string that the value must match", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728527008 := &system.RuleBase{Optional: true}
	ptr8729035280 := &system.String_rule{Base: ptr8729100416, RuleBase: ptr8728527008, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729100544 := &system.Base{Description: "This is a regex to match the value to", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728527616 := &system.RuleBase{Optional: true}
	ptr8729035456 := &system.String_rule{Base: ptr8729100544, RuleBase: ptr8728527616, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729100672 := &system.Base{Description: "This restricts the value to one of several built-in formats", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728528000 := &system.RuleBase{Optional: true}
	ptr8729035632 := &system.String_rule{Base: ptr8729100672, RuleBase: ptr8728528000, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728722944 := &system.Base{Description: "Default value of this is missing or null", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8729214144 := &system.RuleBase{Optional: true}
	ptr8729034928 := &system.String_rule{Base: ptr8728722944, RuleBase: ptr8729214144, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728723072 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8729214464 := &system.RuleBase{Optional: true}
	ptr8728723200 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8729187888 := &system.JsonString_rule{Base: ptr8728723200, RuleBase: (*system.RuleBase)(nil)}
	ptr8728896608 := &system.Array_rule{Base: ptr8728723072, RuleBase: ptr8729214464, Items: ptr8729187888, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728723328 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729214688 := &system.RuleBase{Optional: true}
	ptr8728896768 := &system.Int_rule{Base: ptr8728723328, RuleBase: ptr8729214688, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8729100288 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729214912 := &system.RuleBase{Optional: true}
	ptr8728896848 := &system.Int_rule{Base: ptr8729100288, RuleBase: ptr8729214912, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728967184 := &system.Type{Base: ptr8728722816, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"equal": ptr8729035280, "pattern": ptr8729035456, "format": ptr8729035632, "default": ptr8729034928, "enum": ptr8728896608, "minLength": ptr8728896768, "maxLength": ptr8728896848}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728720000 := &system.Base{Description: "This is the int data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}}
	ptr8728715264 := &system.Base{Description: "Restriction rules for integers", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr8728716544 := &system.Base{Description: "Default value if this property is omitted", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729207264 := &system.RuleBase{Optional: true}
	ptr8728895968 := &system.Int_rule{Base: ptr8728716544, RuleBase: ptr8729207264, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728716672 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729207488 := &system.RuleBase{Optional: true}
	ptr8728896048 := &system.Int_rule{Base: ptr8728716672, RuleBase: ptr8729207488, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728716928 := &system.Base{Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729207712 := &system.RuleBase{Optional: true}
	ptr8728896128 := &system.Int_rule{Base: ptr8728716928, RuleBase: ptr8729207712, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728717056 := &system.Base{Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729207936 := &system.RuleBase{Optional: true}
	ptr8728896208 := &system.Int_rule{Base: ptr8728717056, RuleBase: ptr8729207936, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728961024 := &system.Type{Base: ptr8728715264, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728895968, "multipleOf": ptr8728896048, "minimum": ptr8728896128, "maximum": ptr8728896208}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728962704 := &system.Type{Base: ptr8728720000, Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728961024, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728717568 := &system.Base{Description: "Restriction rules for arrays", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr8728717696 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8729184864 := &system.Rule_rule{Base: ptr8728717696, RuleBase: (*system.RuleBase)(nil)}
	ptr8728717824 := &system.Base{Description: "This is the minimum number of items allowed in the array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729209600 := &system.RuleBase{Optional: true}
	ptr8728895728 := &system.Int_rule{Base: ptr8728717824, RuleBase: ptr8729209600, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728717952 := &system.Base{Description: "This is the maximum number of items allowed in the array", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}}
	ptr8729209856 := &system.RuleBase{Optional: true}
	ptr8728895808 := &system.Int_rule{Base: ptr8728717952, RuleBase: ptr8729209856, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728718080 := &system.Base{Description: "If this is true, each item must be unique", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8729210144 := &system.RuleBase{Optional: true}
	ptr8729185296 := &system.JsonBool_rule{Base: ptr8728718080, RuleBase: ptr8729210144}
	ptr8728961584 := &system.Type{Base: ptr8728717568, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"items": ptr8729184864, "minItems": ptr8728895728, "maxItems": ptr8728895808, "uniqueItems": ptr8729185296}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728722560 := &system.Base{Description: "Automatically created basic rule for ruleBase", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}}
	ptr8728963936 := &system.Type{Base: ptr8728722560, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728722176 := &system.Base{Description: "All rules should embed this type.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}
	ptr8728722432 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8729213408 := &system.RuleBase{Optional: true}
	ptr8729187088 := &system.JsonString_rule{Base: ptr8728722432, RuleBase: ptr8729213408}
	ptr8728722304 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8729213120 := &system.RuleBase{Optional: true}
	ptr8729186896 := &system.JsonBool_rule{Base: ptr8728722304, RuleBase: ptr8729213120}
	ptr8728963712 := &system.Type{Base: ptr8728722176, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"selector": ptr8729187088, "optional": ptr8729186896}, Rule: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8729102336 := &system.Base{Description: "Automatically created basic rule for type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr8728965168 := &system.Type{Base: ptr8729102336, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728720768 := &system.Base{Description: "This is the native json number data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}}
	ptr8728720896 := &system.Base{Description: "Restriction rules for numbers", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr8728721408 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8729211424 := &system.RuleBase{Optional: true}
	ptr8729185552 := &system.JsonBool_rule{Base: ptr8728721408, RuleBase: ptr8729211424}
	ptr8728720128 := &system.Base{Description: "Default value if this property is omitted", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729209984 := &system.RuleBase{Optional: true}
	ptr8728723552 := &system.Number_rule{Base: ptr8728720128, RuleBase: ptr8729209984, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728720256 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729210240 := &system.RuleBase{Optional: true}
	ptr8728724224 := &system.Number_rule{Base: ptr8728720256, RuleBase: ptr8729210240, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728721024 := &system.Base{Description: "This provides a lower bound for the restriction", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729210496 := &system.RuleBase{Optional: true}
	ptr8728723744 := &system.Number_rule{Base: ptr8728721024, RuleBase: ptr8729210496, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728721152 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8729210816 := &system.RuleBase{Optional: true}
	ptr8729185136 := &system.JsonBool_rule{Base: ptr8728721152, RuleBase: ptr8729210816}
	ptr8728721280 := &system.Base{Description: "This provides an upper bound for the restriction", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}}
	ptr8729211072 := &system.RuleBase{Optional: true}
	ptr8728723936 := &system.Number_rule{Base: ptr8728721280, RuleBase: ptr8729211072, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728964944 := &system.Type{Base: ptr8728720896, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"exclusiveMaximum": ptr8729185552, "default": ptr8728723552, "multipleOf": ptr8728724224, "minimum": ptr8728723744, "exclusiveMinimum": ptr8729185136, "maximum": ptr8728723936}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728964272 := &system.Type{Base: ptr8728720768, Native: system.String{Value: "number", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728964944, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728718976 := &system.Base{Description: "Automatically created basic rule for base", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}}
	ptr8728962032 := &system.Type{Base: ptr8728718976, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728719872 := &system.Base{Description: "Automatically created basic rule for imports", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}}
	ptr8728962592 := &system.Type{Base: ptr8728719872, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728721920 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}
	ptr8728963488 := &system.Type{Base: ptr8728721920, Native: system.String{Value: "object", Exists: true}, Interface: true, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728719488 := &system.Base{Description: "Lists imports used in this package.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}}
	ptr8728719616 := &system.Base{Description: "Map of path to alias.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728719744 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8729035104 := &system.String_rule{Base: ptr8728719744, RuleBase: (*system.RuleBase)(nil), Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8728710528 := &system.Map_rule{Base: ptr8728719616, RuleBase: (*system.RuleBase)(nil), Items: ptr8729035104, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728962480 := &system.Type{Base: ptr8728719488, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"imports": ptr8728710528}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728721536 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}}
	ptr8728721664 := &system.Base{Description: "Restriction rules for references", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr8728721792 := &system.Base{Description: "Default value of this is missing or null", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8729212288 := &system.RuleBase{Optional: true}
	ptr8728710400 := &system.Reference_rule{Base: ptr8728721792, RuleBase: ptr8729212288, Default: system.Reference{}}
	ptr8728963376 := &system.Type{Base: ptr8728721664, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"default": ptr8728710400}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728963040 := &system.Type{Base: ptr8728721536, Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728963376, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728719104 := &system.Base{Description: "This is the native json bool data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}}
	ptr8728962256 := &system.Type{Base: ptr8728719104, Native: system.String{Value: "bool", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728962368, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728722048 := &system.Base{Description: "Automatically created basic rule for rule", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr8728963600 := &system.Type{Base: ptr8728722048, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}}
	ptr8728722688 := &system.Base{Description: "This is the native json string data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}}
	ptr8728964048 := &system.Type{Base: ptr8728722688, Native: system.String{Value: "string", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728967184, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8729100800 := &system.Base{Description: "This is the most basic type.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr8729102208 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Id: system.Reference{}}
	ptr8728534816 := &system.RuleBase{Optional: true}
	ptr8729183824 := &system.Type_rule{Base: ptr8729102208, RuleBase: ptr8728534816}
	ptr8729100928 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728528864 := &system.RuleBase{Optional: true}
	ptr8729188944 := &system.JsonBool_rule{Base: ptr8729100928, RuleBase: ptr8728528864}
	ptr8729101056 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728529184 := &system.RuleBase{Optional: true}
	ptr8729101184 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8728711296 := &system.Reference_rule{Base: ptr8729101184, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728897088 := &system.Array_rule{Base: ptr8729101056, RuleBase: ptr8728529184, Items: ptr8728711296, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729101312 := &system.Base{Description: "Array of interface types that this type should support", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}}
	ptr8728526976 := &system.RuleBase{Optional: true}
	ptr8729101440 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8728711552 := &system.Reference_rule{Base: ptr8729101440, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728897168 := &system.Array_rule{Base: ptr8729101312, RuleBase: ptr8728526976, Items: ptr8728711552, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729101568 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728527776 := &system.RuleBase{Optional: true}
	ptr8729034752 := &system.String_rule{Base: ptr8729101568, RuleBase: ptr8728527776, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}, Format: system.String{}}
	ptr8729101696 := &system.Base{Description: "Is this type an interface?", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728528352 := &system.RuleBase{Optional: true}
	ptr8729182656 := &system.JsonBool_rule{Base: ptr8729101696, RuleBase: ptr8728528352}
	ptr8729101824 := &system.Base{Description: "Exclude this type from the generated json?", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}}
	ptr8728528896 := &system.RuleBase{Optional: true}
	ptr8729182848 := &system.JsonBool_rule{Base: ptr8729101824, RuleBase: ptr8728528896}
	ptr8729101952 := &system.Base{Description: "Each field is listed with it's type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}}
	ptr8728529248 := &system.RuleBase{Optional: true}
	ptr8729102080 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8729183472 := &system.Rule_rule{Base: ptr8729102080, RuleBase: (*system.RuleBase)(nil)}
	ptr8728707648 := &system.Map_rule{Base: ptr8729101952, RuleBase: ptr8728529248, Items: ptr8729183472, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728967744 := &system.Type{Base: ptr8729100800, Native: system.String{Value: "object", Exists: true}, Fields: map[string]system.Rule{"rule": ptr8729183824, "basic": ptr8729188944, "embed": ptr8728897088, "is": ptr8728897168, "native": ptr8729034752, "interface": ptr8729182656, "exclude": ptr8729182848, "fields": ptr8728707648}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	ptr8728717440 := &system.Base{Description: "This is the native json array data type", Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}}
	ptr8728961360 := &system.Type{Base: ptr8728717440, Native: system.String{Value: "array", Exists: true}, Fields: map[string]system.Rule(nil), Rule: ptr8728961584, Embed: []system.Reference(nil), Is: []system.Reference(nil)}
	system.RegisterType("kego.io/system", "@imports", ptr8728962592)
	system.RegisterType("kego.io/system", "@rule", ptr8728963600)
	system.RegisterType("kego.io/system", "int", ptr8728962704)
	system.RegisterType("kego.io/system", "@map", ptr8728962928)
	system.RegisterType("kego.io/system", "@base", ptr8728962032)
	system.RegisterType("kego.io/system", "bool", ptr8728962256)
	system.RegisterType("kego.io/system", "@number", ptr8728964944)
	system.RegisterType("kego.io/system", "string", ptr8728964048)
	system.RegisterType("kego.io/system", "array", ptr8728961360)
	system.RegisterType("kego.io/system", "map", ptr8728961248)
	system.RegisterType("kego.io/system", "@array", ptr8728961584)
	system.RegisterType("kego.io/system", "imports", ptr8728962480)
	system.RegisterType("kego.io/system", "reference", ptr8728963040)
	system.RegisterType("kego.io/system", "@bool", ptr8728962368)
	system.RegisterType("kego.io/system", "@ruleBase", ptr8728963936)
	system.RegisterType("kego.io/system", "@int", ptr8728961024)
	system.RegisterType("kego.io/system", "@type", ptr8728965168)
	system.RegisterType("kego.io/system", "number", ptr8728964272)
	system.RegisterType("kego.io/system", "rule", ptr8728963488)
	system.RegisterType("kego.io/system", "@reference", ptr8728963376)
	system.RegisterType("kego.io/system", "type", ptr8728967744)
	system.RegisterType("kego.io/system", "base", ptr8728961808)
	system.RegisterType("kego.io/system", "@string", ptr8728967184)
	system.RegisterType("kego.io/system", "ruleBase", ptr8728963712)
}
