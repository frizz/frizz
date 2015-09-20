package types

import (
	"kego.io/system"
)

func init() {
	ptr0 := &system.Object_base{Description: "This is the base type for the object interface. All ke objects have this type embedded.", Id: system.Reference{Package: "kego.io/system", Name: "$object", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Object_base{Description: "Description for the developer", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr2 := &system.Rule_base{Optional: true}
	ptr3 := &system.JsonString_rule{Object_base: ptr1, Rule_base: ptr2}
	ptr4 := &system.Object_base{Description: "All global objects should have an id.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr5 := &system.Rule_base{Optional: true}
	ptr6 := &system.Reference_rule{Object_base: ptr4, Rule_base: ptr5, Default: system.Reference{}}
	ptr7 := &system.Object_base{Description: "Extra validation rules for this object or descendants", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr8 := &system.Rule_base{Optional: true}
	ptr9 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr10 := &system.Rule_base{}
	ptr11 := &system.Rule_rule{Object_base: ptr9, Rule_base: ptr10}
	ptr12 := &system.Array_rule{Object_base: ptr7, Rule_base: ptr8, Items: ptr11, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr13 := &system.Object_base{Description: "Type of the object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr14 := &system.Rule_base{}
	ptr15 := &system.Reference_rule{Object_base: ptr13, Rule_base: ptr14, Default: system.Reference{}}
	ptr16 := &system.Type{Object_base: ptr0, Base: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"description": ptr3, "id": ptr6, "rules": ptr12, "type": ptr15}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr17 := &system.Object_base{Description: "All rules will have this embedded in them.", Id: system.Reference{Package: "kego.io/system", Name: "$rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr18 := &system.Object_base{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr19 := &system.Rule_base{Optional: true}
	ptr20 := &system.JsonBool_rule{Object_base: ptr18, Rule_base: ptr19}
	ptr21 := &system.Object_base{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr22 := &system.Rule_base{Optional: true}
	ptr23 := &system.JsonString_rule{Object_base: ptr21, Rule_base: ptr22}
	ptr24 := &system.Type{Object_base: ptr17, Base: (*system.Type)(nil), Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.Rule{"optional": ptr20, "selector": ptr23}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr25 := &system.Object_base{Description: "Automatically created basic rule for aliases", Id: system.Reference{Package: "kego.io/system", Name: "@aliases", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr26 := &system.Type{Object_base: ptr25, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr27 := &system.Object_base{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr28 := &system.Object_base{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr29 := &system.Rule_base{}
	ptr30 := &system.Rule_rule{Object_base: ptr28, Rule_base: ptr29}
	ptr31 := &system.Object_base{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr32 := &system.Rule_base{Optional: true}
	ptr33 := &system.Int_rule{Object_base: ptr31, Rule_base: ptr32, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr34 := &system.Object_base{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr35 := &system.Rule_base{Optional: true}
	ptr36 := &system.Int_rule{Object_base: ptr34, Rule_base: ptr35, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr37 := &system.Object_base{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr38 := &system.Rule_base{Optional: true}
	ptr39 := &system.JsonBool_rule{Object_base: ptr37, Rule_base: ptr38}
	ptr40 := &system.Type{Object_base: ptr27, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"items": ptr30, "maxItems": ptr33, "minItems": ptr36, "uniqueItems": ptr39}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr41 := &system.Object_base{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr42 := &system.Object_base{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr43 := &system.Rule_base{Optional: true}
	ptr44 := &system.Bool_rule{Object_base: ptr42, Rule_base: ptr43, Default: system.Bool{}}
	ptr45 := &system.Type{Object_base: ptr41, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr44}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr46 := &system.Object_base{Description: "Automatically created basic rule for imports", Id: system.Reference{Package: "kego.io/system", Name: "@imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr47 := &system.Type{Object_base: ptr46, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr48 := &system.Object_base{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr49 := &system.Object_base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr50 := &system.Rule_base{Optional: true}
	ptr51 := &system.Int_rule{Object_base: ptr49, Rule_base: ptr50, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr52 := &system.Object_base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr53 := &system.Rule_base{Optional: true}
	ptr54 := &system.Int_rule{Object_base: ptr52, Rule_base: ptr53, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr55 := &system.Object_base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr56 := &system.Rule_base{Optional: true}
	ptr57 := &system.Int_rule{Object_base: ptr55, Rule_base: ptr56, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr58 := &system.Object_base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr59 := &system.Rule_base{Optional: true}
	ptr60 := &system.Int_rule{Object_base: ptr58, Rule_base: ptr59, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr61 := &system.Type{Object_base: ptr48, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr51, "maximum": ptr54, "minimum": ptr57, "multipleOf": ptr60}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr62 := &system.Object_base{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr63 := &system.Object_base{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr64 := &system.Rule_base{}
	ptr65 := &system.Rule_rule{Object_base: ptr63, Rule_base: ptr64}
	ptr66 := &system.Object_base{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr67 := &system.Rule_base{Optional: true}
	ptr68 := &system.Int_rule{Object_base: ptr66, Rule_base: ptr67, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr69 := &system.Object_base{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr70 := &system.Rule_base{Optional: true}
	ptr71 := &system.Int_rule{Object_base: ptr69, Rule_base: ptr70, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr72 := &system.Type{Object_base: ptr62, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"items": ptr65, "maxItems": ptr68, "minItems": ptr71}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr73 := &system.Object_base{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr74 := &system.Object_base{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr75 := &system.Rule_base{Optional: true}
	ptr76 := &system.Number_rule{Object_base: ptr74, Rule_base: ptr75, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr77 := &system.Object_base{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr78 := &system.Rule_base{Optional: true}
	ptr79 := &system.JsonBool_rule{Object_base: ptr77, Rule_base: ptr78}
	ptr80 := &system.Object_base{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr81 := &system.Rule_base{Optional: true}
	ptr82 := &system.JsonBool_rule{Object_base: ptr80, Rule_base: ptr81}
	ptr83 := &system.Object_base{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr84 := &system.Rule_base{Optional: true}
	ptr85 := &system.Number_rule{Object_base: ptr83, Rule_base: ptr84, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr86 := &system.Object_base{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr87 := &system.Rule_base{Optional: true}
	ptr88 := &system.Number_rule{Object_base: ptr86, Rule_base: ptr87, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr89 := &system.Object_base{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr90 := &system.Rule_base{Optional: true}
	ptr91 := &system.Number_rule{Object_base: ptr89, Rule_base: ptr90, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr92 := &system.Type{Object_base: ptr73, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr76, "exclusiveMaximum": ptr79, "exclusiveMinimum": ptr82, "maximum": ptr85, "minimum": ptr88, "multipleOf": ptr91}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr93 := &system.Object_base{Description: "Automatically created basic rule for object", Id: system.Reference{Package: "kego.io/system", Name: "@object", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr94 := &system.Type{Object_base: ptr93, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr95 := &system.Object_base{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr96 := &system.Object_base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr97 := &system.Rule_base{Optional: true}
	ptr98 := &system.Reference_rule{Object_base: ptr96, Rule_base: ptr97, Default: system.Reference{}}
	ptr99 := &system.Type{Object_base: ptr95, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr98}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr100 := &system.Object_base{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr101 := &system.Type{Object_base: ptr100, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr102 := &system.Object_base{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr103 := &system.Object_base{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr104 := &system.Rule_base{Optional: true}
	ptr105 := &system.String_rule{Object_base: ptr103, Rule_base: ptr104, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr106 := &system.Object_base{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr107 := &system.Rule_base{Optional: true}
	ptr108 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr109 := &system.Rule_base{}
	ptr110 := &system.JsonString_rule{Object_base: ptr108, Rule_base: ptr109}
	ptr111 := &system.Array_rule{Object_base: ptr106, Rule_base: ptr107, Items: ptr110, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr112 := &system.Object_base{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr113 := &system.Rule_base{Optional: true}
	ptr114 := &system.String_rule{Object_base: ptr112, Rule_base: ptr113, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr115 := &system.Object_base{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr116 := &system.Rule_base{Optional: true}
	ptr117 := &system.String_rule{Object_base: ptr115, Rule_base: ptr116, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr118 := &system.Object_base{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr119 := &system.Rule_base{Optional: true}
	ptr120 := &system.Int_rule{Object_base: ptr118, Rule_base: ptr119, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr121 := &system.Object_base{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr122 := &system.Rule_base{Optional: true}
	ptr123 := &system.Int_rule{Object_base: ptr121, Rule_base: ptr122, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr124 := &system.Object_base{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr125 := &system.Rule_base{Optional: true}
	ptr126 := &system.String_rule{Object_base: ptr124, Rule_base: ptr125, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr127 := &system.Type{Object_base: ptr102, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"default": ptr105, "enum": ptr111, "equal": ptr114, "format": ptr117, "maxLength": ptr120, "minLength": ptr123, "pattern": ptr126}, Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr128 := &system.Object_base{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr129 := &system.Type{Object_base: ptr128, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr130 := &system.Object_base{Description: "Lists aliases used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "aliases", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr131 := &system.Object_base{Description: "Map of path to alias.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr132 := &system.Rule_base{}
	ptr133 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr134 := &system.Rule_base{}
	ptr135 := &system.JsonString_rule{Object_base: ptr133, Rule_base: ptr134}
	ptr136 := &system.Map_rule{Object_base: ptr131, Rule_base: ptr132, Items: ptr135, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr137 := &system.Type{Object_base: ptr130, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"aliases": ptr136}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr138 := &system.Object_base{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr139 := &system.Type{Object_base: ptr138, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr40}
	ptr140 := &system.Object_base{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr141 := &system.Type{Object_base: ptr140, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr45}
	ptr142 := &system.Object_base{Description: "Lists imports used in this package.", Id: system.Reference{Package: "kego.io/system", Name: "imports", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr143 := &system.Object_base{Description: "Map of path to alias.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr144 := &system.Rule_base{}
	ptr145 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr146 := &system.Rule_base{}
	ptr147 := &system.JsonString_rule{Object_base: ptr145, Rule_base: ptr146}
	ptr148 := &system.Map_rule{Object_base: ptr143, Rule_base: ptr144, Items: ptr147, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr149 := &system.Type{Object_base: ptr142, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"imports": ptr148}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr150 := &system.Object_base{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr151 := &system.Type{Object_base: ptr150, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr61}
	ptr152 := &system.Object_base{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr153 := &system.Type{Object_base: ptr152, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr72}
	ptr154 := &system.Object_base{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr155 := &system.Type{Object_base: ptr154, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr92}
	ptr156 := &system.Object_base{Description: "This is the interface that all ke types should implement", Id: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr157 := &system.Type{Object_base: ptr156, Base: ptr16, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr158 := &system.Object_base{Description: "This is a reference to another object, of the form: [local id] or [alias]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr159 := &system.Type{Object_base: ptr158, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr99}
	ptr160 := &system.Object_base{Description: "This is one of several rule types, derived from the rules property of other types", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr161 := &system.Type{Object_base: ptr160, Base: ptr24, Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr162 := &system.Object_base{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr163 := &system.Type{Object_base: ptr162, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr127}
	ptr164 := &system.Object_base{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr165 := &system.Object_base{Description: "Optionally this is a type which is embedded in each object that implements this interface.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr166 := &system.Rule_base{Optional: true}
	ptr167 := &system.Type_rule{Object_base: ptr165, Rule_base: ptr166}
	ptr168 := &system.Object_base{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr169 := &system.Rule_base{Optional: true}
	ptr170 := &system.JsonBool_rule{Object_base: ptr168, Rule_base: ptr169}
	ptr171 := &system.Object_base{Description: "Types which this should embed - system:object is always added unless base = true.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr172 := &system.Rule_base{Optional: true}
	ptr173 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr174 := &system.Rule_base{}
	ptr175 := &system.Reference_rule{Object_base: ptr173, Rule_base: ptr174, Default: system.Reference{}}
	ptr176 := &system.Array_rule{Object_base: ptr171, Rule_base: ptr172, Items: ptr175, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr177 := &system.Object_base{Description: "Exclude this type from the generated json?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr178 := &system.Rule_base{Optional: true}
	ptr179 := &system.JsonBool_rule{Object_base: ptr177, Rule_base: ptr178}
	ptr180 := &system.Object_base{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr181 := &system.Rule_base{Optional: true}
	ptr182 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr183 := &system.Rule_base{}
	ptr184 := &system.Rule_rule{Object_base: ptr182, Rule_base: ptr183}
	ptr185 := &system.Map_rule{Object_base: ptr180, Rule_base: ptr181, Items: ptr184, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr186 := &system.Object_base{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr187 := &system.Rule_base{Optional: true}
	ptr188 := &system.JsonBool_rule{Object_base: ptr186, Rule_base: ptr187}
	ptr189 := &system.Object_base{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr190 := &system.Rule_base{Optional: true}
	ptr191 := &system.Object_base{Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr192 := &system.Rule_base{}
	ptr193 := &system.Reference_rule{Object_base: ptr191, Rule_base: ptr192, Default: system.Reference{}}
	ptr194 := &system.Array_rule{Object_base: ptr189, Rule_base: ptr190, Items: ptr193, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr195 := &system.Object_base{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr196 := &system.Rule_base{Optional: true}
	ptr197 := &system.String_rule{Object_base: ptr195, Rule_base: ptr196, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr198 := &system.Object_base{Description: "Type that defines restriction rules for this type.", Id: system.Reference{}, Rules: []system.Rule(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr199 := &system.Rule_base{Optional: true}
	ptr200 := &system.Type_rule{Object_base: ptr198, Rule_base: ptr199}
	ptr201 := &system.Type{Object_base: ptr164, Base: (*system.Type)(nil), Embed: []system.Reference(nil), Fields: map[string]system.Rule{"base": ptr167, "basic": ptr170, "embed": ptr176, "exclude": ptr179, "fields": ptr185, "interface": ptr188, "is": ptr194, "native": ptr197, "rule": ptr200}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/system", "$object", ptr16, 0xa80f42e03150e43e)
	system.Register("kego.io/system", "$rule", ptr24, 0x61f37939f1737cbf)
	system.Register("kego.io/system", "@aliases", ptr26, 0xa201259ad19e56d5)
	system.Register("kego.io/system", "@array", ptr40, 0xf6f09a20ac87e96f)
	system.Register("kego.io/system", "@bool", ptr45, 0x849d95e096ea903a)
	system.Register("kego.io/system", "@imports", ptr47, 0x9bb5dc657899f549)
	system.Register("kego.io/system", "@int", ptr61, 0x9f6cffbf2042e2aa)
	system.Register("kego.io/system", "@map", ptr72, 0xb95cdb828f1494cc)
	system.Register("kego.io/system", "@number", ptr92, 0x14f986941edf71e9)
	system.Register("kego.io/system", "@object", ptr94, 0xa80f42e03150e43e)
	system.Register("kego.io/system", "@reference", ptr99, 0x67e9d97dde75d10f)
	system.Register("kego.io/system", "@rule", ptr101, 0x61f37939f1737cbf)
	system.Register("kego.io/system", "@string", ptr127, 0xe1e0d90cd0a489ca)
	system.Register("kego.io/system", "@type", ptr129, 0xc3b35155074abe72)
	system.Register("kego.io/system", "aliases", ptr137, 0xa201259ad19e56d5)
	system.Register("kego.io/system", "array", ptr139, 0xf6f09a20ac87e96f)
	system.Register("kego.io/system", "bool", ptr141, 0x849d95e096ea903a)
	system.Register("kego.io/system", "imports", ptr149, 0x9bb5dc657899f549)
	system.Register("kego.io/system", "int", ptr151, 0x9f6cffbf2042e2aa)
	system.Register("kego.io/system", "map", ptr153, 0xb95cdb828f1494cc)
	system.Register("kego.io/system", "number", ptr155, 0x14f986941edf71e9)
	system.Register("kego.io/system", "object", ptr157, 0xa80f42e03150e43e)
	system.Register("kego.io/system", "reference", ptr159, 0x67e9d97dde75d10f)
	system.Register("kego.io/system", "rule", ptr161, 0x61f37939f1737cbf)
	system.Register("kego.io/system", "string", ptr163, 0xe1e0d90cd0a489ca)
	system.Register("kego.io/system", "type", ptr201, 0xc3b35155074abe72)
}
