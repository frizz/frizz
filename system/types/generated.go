package types

import (
	system "kego.io/system"
)

func init() {
	ptr8728590848 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Description: "This is one of several rule types, derived from the rules property of other types"}
	ptr8728824224 := &system.Type{Base: ptr8728590848, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Interface: true}
	ptr8728586112 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Description: "This is the native json object data type."}
	ptr8728586240 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Description: "Restriction rules for maps"}
	ptr8728586368 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Description: "This is a rule object, defining the type and restrictions on the value of the items."}
	ptr8729020016 := &system.Rule_rule{Base: ptr8728586368, RuleBase: (*system.RuleBase)(nil)}
	ptr8728589312 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the minimum number of items alowed in the array"}
	ptr8728954944 := &system.RuleBase{Optional: true}
	ptr8728896288 := &system.Int_rule{Base: ptr8728589312, RuleBase: ptr8728954944, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728589440 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the maximum number of items alowed in the array"}
	ptr8728955232 := &system.RuleBase{Optional: true}
	ptr8728896368 := &system.Int_rule{Base: ptr8728589440, RuleBase: ptr8728955232, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728823664 := &system.Type{Base: ptr8728586240, Fields: map[string]system.Rule{"items": ptr8729020016, "minItems": ptr8728896288, "maxItems": ptr8728896368}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728821984 := &system.Type{Base: ptr8728586112, Fields: map[string]system.Rule(nil), Rule: ptr8728823664, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}}
	ptr8728590592 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Description: "Restriction rules for references"}
	ptr8728590720 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "Default value of this is missing or null"}
	ptr8728958336 := &system.RuleBase{Optional: true}
	ptr8728571072 := &system.Reference_rule{Base: ptr8728590720, RuleBase: ptr8728958336, Default: system.Reference{}}
	ptr8728824112 := &system.Type{Base: ptr8728590592, Fields: map[string]system.Rule{"default": ptr8728571072}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728585344 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Description: "Restriction rules for integers"}
	ptr8728585472 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "Default value if this property is omitted"}
	ptr8728953312 := &system.RuleBase{Optional: true}
	ptr8728895968 := &system.Int_rule{Base: ptr8728585472, RuleBase: ptr8728953312, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728585728 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This restricts the number to be a multiple of the given number"}
	ptr8728953536 := &system.RuleBase{Optional: true}
	ptr8728896048 := &system.Int_rule{Base: ptr8728585728, RuleBase: ptr8728953536, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728585856 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This provides a lower bound for the restriction"}
	ptr8728953760 := &system.RuleBase{Optional: true}
	ptr8728896128 := &system.Int_rule{Base: ptr8728585856, RuleBase: ptr8728953760, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728585984 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This provides an upper bound for the restriction"}
	ptr8728953984 := &system.RuleBase{Optional: true}
	ptr8728896208 := &system.Int_rule{Base: ptr8728585984, RuleBase: ptr8728953984, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728821760 := &system.Type{Base: ptr8728585344, Fields: map[string]system.Rule{"default": ptr8728895968, "multipleOf": ptr8728896048, "minimum": ptr8728896128, "maximum": ptr8728896208}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728589696 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Description: "Restriction rules for numbers"}
	ptr8728589184 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "Default value if this property is omitted"}
	ptr8728956032 := &system.RuleBase{Optional: true}
	ptr8728576096 := &system.Number_rule{Base: ptr8728589184, RuleBase: ptr8728956032, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728589824 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This restricts the number to be a multiple of the given number"}
	ptr8728956288 := &system.RuleBase{Optional: true}
	ptr8728576768 := &system.Number_rule{Base: ptr8728589824, RuleBase: ptr8728956288, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728589952 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This provides a lower bound for the restriction"}
	ptr8728956544 := &system.RuleBase{Optional: true}
	ptr8728576288 := &system.Number_rule{Base: ptr8728589952, RuleBase: ptr8728956544, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728590080 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum."}
	ptr8728956864 := &system.RuleBase{Optional: true}
	ptr8729021296 := &system.JsonBool_rule{Base: ptr8728590080, RuleBase: ptr8728956864}
	ptr8728590208 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Id: system.Reference{}, Description: "This provides an upper bound for the restriction"}
	ptr8728957152 := &system.RuleBase{Optional: true}
	ptr8728576480 := &system.Number_rule{Base: ptr8728590208, RuleBase: ptr8728957152, Maximum: system.Number{}, Default: system.Number{}, MultipleOf: system.Number{}, Minimum: system.Number{}}
	ptr8728590336 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum."}
	ptr8728957472 := &system.RuleBase{Optional: true}
	ptr8729021696 := &system.JsonBool_rule{Base: ptr8728590336, RuleBase: ptr8728957472}
	ptr8728825680 := &system.Type{Base: ptr8728589696, Fields: map[string]system.Rule{"default": ptr8728576096, "multipleOf": ptr8728576768, "minimum": ptr8728576288, "exclusiveMinimum": ptr8729021296, "maximum": ptr8728576480, "exclusiveMaximum": ptr8729021696}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728586624 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Description: "Restriction rules for arrays"}
	ptr8728586880 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the minimum number of items allowed in the array"}
	ptr8728955616 := &system.RuleBase{Optional: true}
	ptr8728895728 := &system.Int_rule{Base: ptr8728586880, RuleBase: ptr8728955616, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728587008 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "This is the maximum number of items allowed in the array"}
	ptr8728955872 := &system.RuleBase{Optional: true}
	ptr8728895808 := &system.Int_rule{Base: ptr8728587008, RuleBase: ptr8728955872, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, Maximum: system.Int{Value: 0}}
	ptr8728587136 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this is true, each item must be unique"}
	ptr8728956160 := &system.RuleBase{Optional: true}
	ptr8729021520 := &system.JsonBool_rule{Base: ptr8728587136, RuleBase: ptr8728956160}
	ptr8728586752 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}, Description: "This is a rule object, defining the type and restrictions on the value of the items"}
	ptr8729021088 := &system.Rule_rule{Base: ptr8728586752, RuleBase: (*system.RuleBase)(nil)}
	ptr8728822320 := &system.Type{Base: ptr8728586624, Fields: map[string]system.Rule{"minItems": ptr8728895728, "maxItems": ptr8728895808, "uniqueItems": ptr8729021520, "items": ptr8729021088}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728588544 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Description: "Lists imports used in this package."}
	ptr8728588672 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "Map of import name to path."}
	ptr8728588800 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8728756576 := &system.String_rule{Base: ptr8728588800, RuleBase: (*system.RuleBase)(nil), Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728571200 := &system.Map_rule{Base: ptr8728588672, RuleBase: (*system.RuleBase)(nil), Items: ptr8728756576, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728823216 := &system.Type{Base: ptr8728588544, Fields: map[string]system.Rule{"imports": ptr8728571200}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728591744 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Description: "Restriction rules for strings"}
	ptr8729075840 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is a string that the value must match"}
	ptr8728395648 := &system.RuleBase{Optional: true}
	ptr8728756752 := &system.String_rule{Base: ptr8729075840, RuleBase: ptr8728395648, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8729075968 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is a regex to match the value to"}
	ptr8728388000 := &system.RuleBase{Optional: true}
	ptr8728756928 := &system.String_rule{Base: ptr8729075968, RuleBase: ptr8728388000, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8729076096 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This restricts the value to one of several built-in formats"}
	ptr8728388576 := &system.RuleBase{Optional: true}
	ptr8728757104 := &system.String_rule{Base: ptr8729076096, RuleBase: ptr8728388576, Format: system.String{}, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728591872 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Default value of this is missing or null"}
	ptr8728960192 := &system.RuleBase{Optional: true}
	ptr8728756400 := &system.String_rule{Base: ptr8728591872, RuleBase: ptr8728960192, Format: system.String{}, Default: system.String{}, Enum: []string(nil), MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8728592000 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "The value of this string is restricted to one of the provided values"}
	ptr8728960512 := &system.RuleBase{Optional: true}
	ptr8728592128 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}}
	ptr8729024048 := &system.JsonString_rule{Base: ptr8728592128, RuleBase: (*system.RuleBase)(nil)}
	ptr8728896608 := &system.Array_rule{Base: ptr8728592000, RuleBase: ptr8728960512, Items: ptr8729024048, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728592256 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "The value must be longer or equal to the provided minimum length"}
	ptr8728960736 := &system.RuleBase{Optional: true}
	ptr8728896768 := &system.Int_rule{Base: ptr8728592256, RuleBase: ptr8728960736, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8729075712 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Id: system.Reference{}, Description: "The value must be shorter or equal to the provided maximum length"}
	ptr8728960960 := &system.RuleBase{Optional: true}
	ptr8728896848 := &system.Int_rule{Base: ptr8729075712, RuleBase: ptr8728960960, Default: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, Maximum: system.Int{Value: 0}}
	ptr8728827920 := &system.Type{Base: ptr8728591744, Fields: map[string]system.Rule{"equal": ptr8728756752, "pattern": ptr8728756928, "format": ptr8728757104, "default": ptr8728756400, "enum": ptr8728896608, "minLength": ptr8728896768, "maxLength": ptr8728896848}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728588928 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Description: "Automatically created basic rule for imports"}
	ptr8728823328 := &system.Type{Base: ptr8728588928, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728589568 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Description: "This is the native json number data type"}
	ptr8728825008 := &system.Type{Base: ptr8728589568, Fields: map[string]system.Rule(nil), Rule: ptr8728825680, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}}
	ptr8728591616 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Description: "This is the native json string data type"}
	ptr8728824784 := &system.Type{Base: ptr8728591616, Fields: map[string]system.Rule(nil), Rule: ptr8728827920, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}}
	ptr8728590976 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Description: "Automatically created basic rule for rule"}
	ptr8728824336 := &system.Type{Base: ptr8728590976, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8729076224 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Description: "This is the most basic type."}
	ptr8729077632 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Id: system.Reference{}, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character."}
	ptr8728390304 := &system.RuleBase{Optional: true}
	ptr8729020272 := &system.Type_rule{Base: ptr8729077632, RuleBase: ptr8728390304}
	ptr8729076352 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Basic types don't have system:object added by default to the embedded types."}
	ptr8728389408 := &system.RuleBase{Optional: true}
	ptr8729025104 := &system.JsonBool_rule{Base: ptr8729076352, RuleBase: ptr8728389408}
	ptr8729076480 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Types which this should embed - system:object is always added unless base = true."}
	ptr8728387744 := &system.RuleBase{Optional: true}
	ptr8729076608 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8728568320 := &system.Reference_rule{Base: ptr8729076608, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728897088 := &system.Array_rule{Base: ptr8729076480, RuleBase: ptr8728387744, Items: ptr8728568320, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729076736 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Array of interface types that this type should support"}
	ptr8728388480 := &system.RuleBase{Optional: true}
	ptr8729076864 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}}
	ptr8728568448 := &system.Reference_rule{Base: ptr8729076864, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728897168 := &system.Array_rule{Base: ptr8729076736, RuleBase: ptr8728388480, Items: ptr8728568448, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8729076992 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "This is the native json type that represents this type. If omitted, default is object."}
	ptr8728389056 := &system.RuleBase{Optional: true}
	ptr8728756224 := &system.String_rule{Base: ptr8729076992, RuleBase: ptr8728389056, Format: system.String{}, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, MinLength: system.Int{Value: 0}, MaxLength: system.Int{Value: 0}, Equal: system.String{}, Pattern: system.String{}}
	ptr8729077120 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Is this type an interface?"}
	ptr8728389632 := &system.RuleBase{Optional: true}
	ptr8729019456 := &system.JsonBool_rule{Base: ptr8729077120, RuleBase: ptr8728389632}
	ptr8729077248 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Exclude this type from the generated json?"}
	ptr8728389792 := &system.RuleBase{Optional: true}
	ptr8729019648 := &system.JsonBool_rule{Base: ptr8729077248, RuleBase: ptr8728389792}
	ptr8729077376 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Id: system.Reference{}, Description: "Each field is listed with it's type"}
	ptr8728390080 := &system.RuleBase{Optional: true}
	ptr8729077504 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8729020096 := &system.Rule_rule{Base: ptr8729077504, RuleBase: (*system.RuleBase)(nil)}
	ptr8728569024 := &system.Map_rule{Base: ptr8729077376, RuleBase: ptr8728390080, Items: ptr8729020096, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728828480 := &system.Type{Base: ptr8729076224, Fields: map[string]system.Rule{"rule": ptr8729020272, "basic": ptr8729025104, "embed": ptr8728897088, "is": ptr8728897168, "native": ptr8728756224, "interface": ptr8729019456, "exclude": ptr8729019648, "fields": ptr8728569024}, Rule: (*system.Type)(nil), Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8729077760 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Description: "Automatically created basic rule for type"}
	ptr8728825904 := &system.Type{Base: ptr8729077760, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728588288 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Description: "Restriction rules for bools"}
	ptr8728588416 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "Default value if this is missing or null"}
	ptr8728958272 := &system.RuleBase{Optional: true}
	ptr8728958208 := &system.Bool_rule{Base: ptr8728588416, RuleBase: ptr8728958272, Default: system.Bool{}}
	ptr8728823104 := &system.Type{Base: ptr8728588288, Fields: map[string]system.Rule{"default": ptr8728958208}, Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728587264 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Description: "This is the most basic type."}
	ptr8728587776 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Id: system.Reference{}, Description: "Extra validation rules for this object or descendants"}
	ptr8728957504 := &system.RuleBase{Optional: true}
	ptr8728587904 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Id: system.Reference{}}
	ptr8729022400 := &system.Rule_rule{Base: ptr8728587904, RuleBase: (*system.RuleBase)(nil)}
	ptr8728895888 := &system.Array_rule{Base: ptr8728587776, RuleBase: ptr8728957504, Items: ptr8729022400, MinItems: system.Int{Value: 0}, MaxItems: system.Int{Value: 0}}
	ptr8728587392 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "Type of the object."}
	ptr8728936064 := &system.Reference_rule{Base: ptr8728587392, RuleBase: (*system.RuleBase)(nil), Default: system.Reference{}}
	ptr8728587520 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Id: system.Reference{}, Description: "All global objects should have an id."}
	ptr8728956896 := &system.RuleBase{Optional: true}
	ptr8728936128 := &system.Reference_rule{Base: ptr8728587520, RuleBase: ptr8728956896, Default: system.Reference{}}
	ptr8728587648 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Description for the developer"}
	ptr8728957216 := &system.RuleBase{Optional: true}
	ptr8729022176 := &system.JsonString_rule{Base: ptr8728587648, RuleBase: ptr8728957216}
	ptr8728822544 := &system.Type{Base: ptr8728587264, Fields: map[string]system.Rule{"rules": ptr8728895888, "type": ptr8728936064, "id": ptr8728936128, "description": ptr8729022176}, Rule: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728591104 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Description: "All rules should embed this type."}
	ptr8728591232 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}, Id: system.Reference{}, Description: "If this rule is a field, this specifies that the field is optional"}
	ptr8728959168 := &system.RuleBase{Optional: true}
	ptr8729023040 := &system.JsonBool_rule{Base: ptr8728591232, RuleBase: ptr8728959168}
	ptr8728591360 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}, Id: system.Reference{}, Description: "Json selector defining what nodes this rule should be applied to."}
	ptr8728959456 := &system.RuleBase{Optional: true}
	ptr8729023216 := &system.JsonString_rule{Base: ptr8728591360, RuleBase: ptr8728959456}
	ptr8728824448 := &system.Type{Base: ptr8728591104, Fields: map[string]system.Rule{"optional": ptr8729023040, "selector": ptr8729023216}, Rule: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}}
	ptr8728588160 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Description: "This is the native json bool data type"}
	ptr8728822992 := &system.Type{Base: ptr8728588160, Fields: map[string]system.Rule(nil), Rule: ptr8728823104, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}}
	ptr8728586496 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Description: "This is the native json array data type"}
	ptr8728822096 := &system.Type{Base: ptr8728586496, Fields: map[string]system.Rule(nil), Rule: ptr8728822320, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}}
	ptr8728591488 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Description: "Automatically created basic rule for ruleBase"}
	ptr8728824672 := &system.Type{Base: ptr8728591488, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728588032 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Description: "Automatically created basic rule for base"}
	ptr8728822768 := &system.Type{Base: ptr8728588032, Fields: map[string]system.Rule(nil), Rule: (*system.Type)(nil), Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}}
	ptr8728590464 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]"}
	ptr8728823776 := &system.Type{Base: ptr8728590464, Fields: map[string]system.Rule(nil), Rule: ptr8728824112, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}}
	ptr8728589056 := &system.Base{Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Description: "This is the int data type"}
	ptr8728823440 := &system.Type{Base: ptr8728589056, Fields: map[string]system.Rule(nil), Rule: ptr8728821760, Embed: []system.Reference(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}}
	system.RegisterType("kego.io/system", "bool", ptr8728822992)
	system.RegisterType("kego.io/system", "imports", ptr8728823216)
	system.RegisterType("kego.io/system", "@string", ptr8728827920)
	system.RegisterType("kego.io/system", "@map", ptr8728823664)
	system.RegisterType("kego.io/system", "@ruleBase", ptr8728824672)
	system.RegisterType("kego.io/system", "reference", ptr8728823776)
	system.RegisterType("kego.io/system", "int", ptr8728823440)
	system.RegisterType("kego.io/system", "@number", ptr8728825680)
	system.RegisterType("kego.io/system", "string", ptr8728824784)
	system.RegisterType("kego.io/system", "ruleBase", ptr8728824448)
	system.RegisterType("kego.io/system", "@imports", ptr8728823328)
	system.RegisterType("kego.io/system", "@rule", ptr8728824336)
	system.RegisterType("kego.io/system", "@type", ptr8728825904)
	system.RegisterType("kego.io/system", "@base", ptr8728822768)
	system.RegisterType("kego.io/system", "rule", ptr8728824224)
	system.RegisterType("kego.io/system", "map", ptr8728821984)
	system.RegisterType("kego.io/system", "@int", ptr8728821760)
	system.RegisterType("kego.io/system", "type", ptr8728828480)
	system.RegisterType("kego.io/system", "@bool", ptr8728823104)
	system.RegisterType("kego.io/system", "base", ptr8728822544)
	system.RegisterType("kego.io/system", "array", ptr8728822096)
	system.RegisterType("kego.io/system", "@reference", ptr8728824112)
	system.RegisterType("kego.io/system", "@array", ptr8728822320)
	system.RegisterType("kego.io/system", "number", ptr8728825008)
}
