package types

import (
	"kego.io/system"
)

func init() {
	ptr0 := &system.Base{Description: "Automatically created basic rule for aliases", Id: system.Reference{Package: "kego.io/system", Name: "@aliases", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Type{Base: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr2 := &system.Base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr3 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr4 := &system.RuleBase{}
	ptr5 := &system.Rule_rule{Base: ptr3, RuleBase: ptr4}
	ptr6 := &system.Base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr7 := &system.RuleBase{Optional: true}
	ptr8 := &system.Int_rule{Base: ptr6, RuleBase: ptr7, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr9 := &system.Base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr10 := &system.RuleBase{Optional: true}
	ptr11 := &system.Int_rule{Base: ptr9, RuleBase: ptr10, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr12 := &system.Base{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr13 := &system.RuleBase{Optional: true}
	ptr14 := &system.JsonBool_rule{Base: ptr12, RuleBase: ptr13}
	ptr15 := &system.Type{Base: ptr2, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"items": ptr5, "maxItems": ptr8, "minItems": ptr11, "uniqueItems": ptr14}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr16 := &system.Base{Description: "Automatically created basic rule for base", Id: system.Reference{Package: "kego.io/system", Name: "@base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr17 := &system.Type{Base: ptr16, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr18 := &system.Base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr19 := &system.Base{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr20 := &system.RuleBase{Optional: true}
	ptr21 := &system.Bool_rule{Base: ptr19, RuleBase: ptr20, Default: system.Bool{}}
	ptr22 := &system.Type{Base: ptr18, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr21}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr23 := &system.Base{Description: "Automatically created basic rule for imports", Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr24 := &system.Type{Base: ptr23, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr25 := &system.Base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr26 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr27 := &system.RuleBase{Optional: true}
	ptr28 := &system.Int_rule{Base: ptr26, RuleBase: ptr27, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr29 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr30 := &system.RuleBase{Optional: true}
	ptr31 := &system.Int_rule{Base: ptr29, RuleBase: ptr30, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr32 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr33 := &system.RuleBase{Optional: true}
	ptr34 := &system.Int_rule{Base: ptr32, RuleBase: ptr33, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr35 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr36 := &system.RuleBase{Optional: true}
	ptr37 := &system.Int_rule{Base: ptr35, RuleBase: ptr36, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr38 := &system.Type{Base: ptr25, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr28, "maximum": ptr31, "minimum": ptr34, "multipleOf": ptr37}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr39 := &system.Base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr40 := &system.Base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr41 := &system.RuleBase{}
	ptr42 := &system.Rule_rule{Base: ptr40, RuleBase: ptr41}
	ptr43 := &system.Base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr44 := &system.RuleBase{Optional: true}
	ptr45 := &system.Int_rule{Base: ptr43, RuleBase: ptr44, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr46 := &system.Base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr47 := &system.RuleBase{Optional: true}
	ptr48 := &system.Int_rule{Base: ptr46, RuleBase: ptr47, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr49 := &system.Type{Base: ptr39, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"items": ptr42, "maxItems": ptr45, "minItems": ptr48}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr50 := &system.Base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr51 := &system.Base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr52 := &system.RuleBase{Optional: true}
	ptr53 := &system.Number_rule{Base: ptr51, RuleBase: ptr52, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr54 := &system.Base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr55 := &system.RuleBase{Optional: true}
	ptr56 := &system.JsonBool_rule{Base: ptr54, RuleBase: ptr55}
	ptr57 := &system.Base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr58 := &system.RuleBase{Optional: true}
	ptr59 := &system.JsonBool_rule{Base: ptr57, RuleBase: ptr58}
	ptr60 := &system.Base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr61 := &system.RuleBase{Optional: true}
	ptr62 := &system.Number_rule{Base: ptr60, RuleBase: ptr61, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr63 := &system.Base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr64 := &system.RuleBase{Optional: true}
	ptr65 := &system.Number_rule{Base: ptr63, RuleBase: ptr64, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr66 := &system.Base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr67 := &system.RuleBase{Optional: true}
	ptr68 := &system.Number_rule{Base: ptr66, RuleBase: ptr67, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr69 := &system.Type{Base: ptr50, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr53, "exclusiveMaximum": ptr56, "exclusiveMinimum": ptr59, "maximum": ptr62, "minimum": ptr65, "multipleOf": ptr68}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr70 := &system.Base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr71 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr72 := &system.RuleBase{Optional: true}
	ptr73 := &system.Reference_rule{Base: ptr71, RuleBase: ptr72, Default: system.Reference{}}
	ptr74 := &system.Type{Base: ptr70, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr73}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr75 := &system.Base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr76 := &system.Type{Base: ptr75, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr77 := &system.Base{Description: "Automatically created basic rule for ruleBase", Id: system.Reference{Package: "kego.io/system", Name: "@ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr78 := &system.Type{Base: ptr77, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr79 := &system.Base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr80 := &system.Base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr81 := &system.RuleBase{Optional: true}
	ptr82 := &system.String_rule{Base: ptr80, RuleBase: ptr81, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr83 := &system.Base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr84 := &system.RuleBase{Optional: true}
	ptr85 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr86 := &system.RuleBase{}
	ptr87 := &system.JsonString_rule{Base: ptr85, RuleBase: ptr86}
	ptr88 := &system.Array_rule{Base: ptr83, RuleBase: ptr84, Items: ptr87, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr89 := &system.Base{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr90 := &system.RuleBase{Optional: true}
	ptr91 := &system.String_rule{Base: ptr89, RuleBase: ptr90, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr92 := &system.Base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr93 := &system.RuleBase{Optional: true}
	ptr94 := &system.String_rule{Base: ptr92, RuleBase: ptr93, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr95 := &system.Base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr96 := &system.RuleBase{Optional: true}
	ptr97 := &system.Int_rule{Base: ptr95, RuleBase: ptr96, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr98 := &system.Base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr99 := &system.RuleBase{Optional: true}
	ptr100 := &system.Int_rule{Base: ptr98, RuleBase: ptr99, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr101 := &system.Base{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr102 := &system.RuleBase{Optional: true}
	ptr103 := &system.String_rule{Base: ptr101, RuleBase: ptr102, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr104 := &system.Type{Base: ptr79, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule{"default": ptr82, "enum": ptr88, "equal": ptr91, "format": ptr94, "maxLength": ptr97, "minLength": ptr100, "pattern": ptr103}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr105 := &system.Base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr106 := &system.Type{Base: ptr105, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}}, Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr107 := &system.Base{Description: "Lists aliases used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "aliases", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr108 := &system.Base{Description: "Map of path to alias.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr109 := &system.RuleBase{}
	ptr110 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr111 := &system.RuleBase{}
	ptr112 := &system.JsonString_rule{Base: ptr110, RuleBase: ptr111}
	ptr113 := &system.Map_rule{Base: ptr108, RuleBase: ptr109, Items: ptr112, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr114 := &system.Type{Base: ptr107, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"aliases": ptr113}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr115 := &system.Base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr116 := &system.Type{Base: ptr115, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr15}
	ptr117 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "base", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr118 := &system.Base{Description: "Description for the developer", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr119 := &system.RuleBase{Optional: true}
	ptr120 := &system.JsonString_rule{Base: ptr118, RuleBase: ptr119}
	ptr121 := &system.Base{Description: "All global objects should have an id.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr122 := &system.RuleBase{Optional: true}
	ptr123 := &system.Reference_rule{Base: ptr121, RuleBase: ptr122, Default: system.Reference{}}
	ptr124 := &system.Base{Description: "Extra validation rules for this object or descendants", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr125 := &system.RuleBase{Optional: true}
	ptr126 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr127 := &system.RuleBase{}
	ptr128 := &system.Rule_rule{Base: ptr126, RuleBase: ptr127}
	ptr129 := &system.Array_rule{Base: ptr124, RuleBase: ptr125, Items: ptr128, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr130 := &system.Base{Description: "Type of the object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr131 := &system.RuleBase{}
	ptr132 := &system.Reference_rule{Base: ptr130, RuleBase: ptr131, Default: system.Reference{}}
	ptr133 := &system.Type{Base: ptr117, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"description": ptr120, "id": ptr123, "rules": ptr129, "type": ptr132}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr134 := &system.Base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr135 := &system.Type{Base: ptr134, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr22}
	ptr136 := &system.Base{Description: "Lists imports used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr137 := &system.Base{Description: "Map of path to alias.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr138 := &system.RuleBase{}
	ptr139 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr140 := &system.RuleBase{}
	ptr141 := &system.JsonString_rule{Base: ptr139, RuleBase: ptr140}
	ptr142 := &system.Map_rule{Base: ptr137, RuleBase: ptr138, Items: ptr141, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr143 := &system.Type{Base: ptr136, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"imports": ptr142}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr144 := &system.Base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr145 := &system.Type{Base: ptr144, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr38}
	ptr146 := &system.Base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr147 := &system.Type{Base: ptr146, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr49}
	ptr148 := &system.Base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr149 := &system.Type{Base: ptr148, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr69}
	ptr150 := &system.Base{Description: "This is a reference to another object, of the form: [local id] or [alias]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr151 := &system.Type{Base: ptr150, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr74}
	ptr152 := &system.Base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr153 := &system.Type{Base: ptr152, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr154 := &system.Base{Description: "All rules should embed this type.", Id: system.Reference{Package: "kego.io/system", Name: "ruleBase", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr155 := &system.Base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr156 := &system.RuleBase{Optional: true}
	ptr157 := &system.JsonBool_rule{Base: ptr155, RuleBase: ptr156}
	ptr158 := &system.Base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr159 := &system.RuleBase{Optional: true}
	ptr160 := &system.JsonString_rule{Base: ptr158, RuleBase: ptr159}
	ptr161 := &system.Type{Base: ptr154, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"optional": ptr157, "selector": ptr160}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr162 := &system.Base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr163 := &system.Type{Base: ptr162, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr104}
	ptr164 := &system.Base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr165 := &system.Base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr166 := &system.RuleBase{Optional: true}
	ptr167 := &system.JsonBool_rule{Base: ptr165, RuleBase: ptr166}
	ptr168 := &system.Base{Description: "Types which this should embed - system:object is always added unless base = true.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr169 := &system.RuleBase{Optional: true}
	ptr170 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr171 := &system.RuleBase{}
	ptr172 := &system.Reference_rule{Base: ptr170, RuleBase: ptr171, Default: system.Reference{}}
	ptr173 := &system.Array_rule{Base: ptr168, RuleBase: ptr169, Items: ptr172, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr174 := &system.Base{Description: "Exclude this type from the generated json?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr175 := &system.RuleBase{Optional: true}
	ptr176 := &system.JsonBool_rule{Base: ptr174, RuleBase: ptr175}
	ptr177 := &system.Base{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr178 := &system.RuleBase{Optional: true}
	ptr179 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr180 := &system.RuleBase{}
	ptr181 := &system.Rule_rule{Base: ptr179, RuleBase: ptr180}
	ptr182 := &system.Map_rule{Base: ptr177, RuleBase: ptr178, Items: ptr181, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr183 := &system.Base{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr184 := &system.RuleBase{Optional: true}
	ptr185 := &system.JsonBool_rule{Base: ptr183, RuleBase: ptr184}
	ptr186 := &system.Base{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr187 := &system.RuleBase{Optional: true}
	ptr188 := &system.Base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr189 := &system.RuleBase{}
	ptr190 := &system.Reference_rule{Base: ptr188, RuleBase: ptr189, Default: system.Reference{}}
	ptr191 := &system.Array_rule{Base: ptr186, RuleBase: ptr187, Items: ptr190, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr192 := &system.Base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr193 := &system.RuleBase{Optional: true}
	ptr194 := &system.String_rule{Base: ptr192, RuleBase: ptr193, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr195 := &system.Base{Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr196 := &system.RuleBase{Optional: true}
	ptr197 := &system.Type_rule{Base: ptr195, RuleBase: ptr196}
	ptr198 := &system.Type{Base: ptr164, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"basic": ptr167, "embed": ptr173, "exclude": ptr176, "fields": ptr182, "interface": ptr185, "is": ptr191, "native": ptr194, "rule": ptr197}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/system", "@aliases", ptr1, 0xa201259ad19e56d5)
	system.Register("kego.io/system", "@array", ptr15, 0x9396ab3adc9b26c2)
	system.Register("kego.io/system", "@base", ptr17, 0x55165bf6126129f)
	system.Register("kego.io/system", "@bool", ptr22, 0x18e72064215d8f71)
	system.Register("kego.io/system", "@imports", ptr24, 0x9bb5dc657899f549)
	system.Register("kego.io/system", "@int", ptr38, 0xe7b98d2f66bb5b1)
	system.Register("kego.io/system", "@map", ptr49, 0x68d11e9fab3ed4bb)
	system.Register("kego.io/system", "@number", ptr69, 0x8c2b0d23320b8aeb)
	system.Register("kego.io/system", "@reference", ptr74, 0x2a5ddf0f63b58876)
	system.Register("kego.io/system", "@rule", ptr76, 0x50ad825a1446eccd)
	system.Register("kego.io/system", "@ruleBase", ptr78, 0xab213a8c8218179a)
	system.Register("kego.io/system", "@string", ptr104, 0xae4613a6af2bb7d6)
	system.Register("kego.io/system", "@type", ptr106, 0xc1bba2c7da9518b7)
	system.Register("kego.io/system", "aliases", ptr114, 0xa201259ad19e56d5)
	system.Register("kego.io/system", "array", ptr116, 0x9396ab3adc9b26c2)
	system.Register("kego.io/system", "base", ptr133, 0x55165bf6126129f)
	system.Register("kego.io/system", "bool", ptr135, 0x18e72064215d8f71)
	system.Register("kego.io/system", "imports", ptr143, 0x9bb5dc657899f549)
	system.Register("kego.io/system", "int", ptr145, 0xe7b98d2f66bb5b1)
	system.Register("kego.io/system", "map", ptr147, 0x68d11e9fab3ed4bb)
	system.Register("kego.io/system", "number", ptr149, 0x8c2b0d23320b8aeb)
	system.Register("kego.io/system", "reference", ptr151, 0x2a5ddf0f63b58876)
	system.Register("kego.io/system", "rule", ptr153, 0x50ad825a1446eccd)
	system.Register("kego.io/system", "ruleBase", ptr161, 0xab213a8c8218179a)
	system.Register("kego.io/system", "string", ptr163, 0xae4613a6af2bb7d6)
	system.Register("kego.io/system", "type", ptr198, 0xc1bba2c7da9518b7)
}
