package types

import (
	"kego.io/system"
)

func init() {
	ptr0 := &system.Object{Description: "Restriction rules for arrays", Id: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr1 := &system.Object{Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr2 := &system.Rule{Interface: true}
	ptr3 := &system.RuleRule{Object: ptr1, Rule: ptr2}
	ptr4 := &system.Object{Description: "This is the maximum number of items allowed in the array", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr5 := &system.Rule{Optional: true}
	ptr6 := &system.IntRule{Object: ptr4, Rule: ptr5, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr7 := &system.Object{Description: "This is the minimum number of items allowed in the array", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr8 := &system.Rule{Optional: true}
	ptr9 := &system.IntRule{Object: ptr7, Rule: ptr8, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr10 := &system.Object{Description: "If this is true, each item must be unique", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr11 := &system.Rule{Optional: true}
	ptr12 := &system.JsonBoolRule{Object: ptr10, Rule: ptr11}
	ptr13 := &system.Type{Object: ptr0, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface{"items": ptr3, "maxItems": ptr6, "minItems": ptr9, "uniqueItems": ptr12}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr14 := &system.Object{Description: "Restriction rules for bools", Id: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr15 := &system.Object{Description: "Default value if this is missing or null", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@bool", Exists: true}}
	ptr16 := &system.Rule{Optional: true}
	ptr17 := &system.BoolRule{Object: ptr15, Rule: ptr16, Default: system.Bool{}}
	ptr18 := &system.Type{Object: ptr14, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface{"default": ptr17}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr19 := &system.Object{Description: "Restriction rules for integers", Id: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr20 := &system.Object{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr21 := &system.Rule{Optional: true}
	ptr22 := &system.IntRule{Object: ptr20, Rule: ptr21, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr23 := &system.Object{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr24 := &system.Rule{Optional: true}
	ptr25 := &system.IntRule{Object: ptr23, Rule: ptr24, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr26 := &system.Object{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr27 := &system.Rule{Optional: true}
	ptr28 := &system.IntRule{Object: ptr26, Rule: ptr27, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr29 := &system.Object{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr30 := &system.Rule{Optional: true}
	ptr31 := &system.IntRule{Object: ptr29, Rule: ptr30, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr32 := &system.Type{Object: ptr19, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface{"default": ptr22, "maximum": ptr25, "minimum": ptr28, "multipleOf": ptr31}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr33 := &system.Object{Description: "Restriction rules for maps", Id: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr34 := &system.Object{Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr35 := &system.Rule{Interface: true}
	ptr36 := &system.RuleRule{Object: ptr34, Rule: ptr35}
	ptr37 := &system.Object{Description: "This is the maximum number of items alowed in the array", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr38 := &system.Rule{Optional: true}
	ptr39 := &system.IntRule{Object: ptr37, Rule: ptr38, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr40 := &system.Object{Description: "This is the minimum number of items alowed in the array", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr41 := &system.Rule{Optional: true}
	ptr42 := &system.IntRule{Object: ptr40, Rule: ptr41, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0}}
	ptr43 := &system.Type{Object: ptr33, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface{"items": ptr36, "maxItems": ptr39, "minItems": ptr42}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr44 := &system.Object{Description: "Restriction rules for numbers", Id: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr45 := &system.Object{Description: "Default value if this property is omitted", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr46 := &system.Rule{Optional: true}
	ptr47 := &system.NumberRule{Object: ptr45, Rule: ptr46, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr48 := &system.Object{Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr49 := &system.Rule{Optional: true}
	ptr50 := &system.JsonBoolRule{Object: ptr48, Rule: ptr49}
	ptr51 := &system.Object{Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr52 := &system.Rule{Optional: true}
	ptr53 := &system.JsonBoolRule{Object: ptr51, Rule: ptr52}
	ptr54 := &system.Object{Description: "This provides an upper bound for the restriction", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr55 := &system.Rule{Optional: true}
	ptr56 := &system.NumberRule{Object: ptr54, Rule: ptr55, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr57 := &system.Object{Description: "This provides a lower bound for the restriction", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr58 := &system.Rule{Optional: true}
	ptr59 := &system.NumberRule{Object: ptr57, Rule: ptr58, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr60 := &system.Object{Description: "This restricts the number to be a multiple of the given number", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@number", Exists: true}}
	ptr61 := &system.Rule{Optional: true}
	ptr62 := &system.NumberRule{Object: ptr60, Rule: ptr61, Default: system.Number{}, Maximum: system.Number{}, Minimum: system.Number{}, MultipleOf: system.Number{}}
	ptr63 := &system.Type{Object: ptr44, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface{"default": ptr47, "exclusiveMaximum": ptr50, "exclusiveMinimum": ptr53, "maximum": ptr56, "minimum": ptr59, "multipleOf": ptr62}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr64 := &system.Object{Description: "Automatically created basic rule for object", Id: system.Reference{Package: "kego.io/system", Name: "@object", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr65 := &system.Type{Object: ptr64, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr66 := &system.Object{Description: "Automatically created basic rule for package", Id: system.Reference{Package: "kego.io/system", Name: "@package", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr67 := &system.Type{Object: ptr66, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr68 := &system.Object{Description: "Restriction rules for references", Id: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr69 := &system.Object{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr70 := &system.Rule{Optional: true}
	ptr71 := &system.ReferenceRule{Object: ptr69, Rule: ptr70, Default: system.Reference{}}
	ptr72 := &system.Type{Object: ptr68, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface{"default": ptr71}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr73 := &system.Object{Description: "Automatically created basic rule for rule", Id: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr74 := &system.Type{Object: ptr73, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr75 := &system.Object{Description: "Restriction rules for strings", Id: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr76 := &system.Object{Description: "Default value of this is missing or null", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr77 := &system.Rule{Optional: true}
	ptr78 := &system.StringRule{Object: ptr76, Rule: ptr77, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr79 := &system.Object{Description: "The value of this string is restricted to one of the provided values", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr80 := &system.Rule{Optional: true}
	ptr81 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr82 := &system.Rule{}
	ptr83 := &system.JsonStringRule{Object: ptr81, Rule: ptr82}
	ptr84 := &system.ArrayRule{Object: ptr79, Rule: ptr80, Items: ptr83, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr85 := &system.Object{Description: "This is a string that the value must match", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr86 := &system.Rule{Optional: true}
	ptr87 := &system.StringRule{Object: ptr85, Rule: ptr86, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr88 := &system.Object{Description: "This restricts the value to one of several built-in formats", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr89 := &system.Rule{Optional: true}
	ptr90 := &system.StringRule{Object: ptr88, Rule: ptr89, Default: system.String{}, Enum: []string{"date-time", "email", "hostname", "ipv4", "ipv6", "uri"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr91 := &system.Object{Description: "The value must be shorter or equal to the provided maximum length", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr92 := &system.Rule{Optional: true}
	ptr93 := &system.IntRule{Object: ptr91, Rule: ptr92, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr94 := &system.Object{Description: "The value must be longer or equal to the provided minimum length", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@int", Exists: true}}
	ptr95 := &system.Rule{Optional: true}
	ptr96 := &system.IntRule{Object: ptr94, Rule: ptr95, Default: system.Int{Value: 0}, Maximum: system.Int{Value: 0}, Minimum: system.Int{Value: 0}, MultipleOf: system.Int{Value: 0}}
	ptr97 := &system.Object{Description: "This is a regex to match the value to", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr98 := &system.Rule{Optional: true}
	ptr99 := &system.StringRule{Object: ptr97, Rule: ptr98, Default: system.String{}, Enum: []string(nil), Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr100 := &system.Type{Object: ptr75, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface{"default": ptr78, "enum": ptr84, "equal": ptr87, "format": ptr90, "maxLength": ptr93, "minLength": ptr96, "pattern": ptr99}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr101 := &system.Object{Description: "Automatically created basic rule for type", Id: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr102 := &system.Type{Object: ptr101, Embed: []system.Reference{system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}}, Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr103 := &system.Object{Description: "This is the native json array data type", Id: system.Reference{Package: "kego.io/system", Name: "array", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr104 := &system.Type{Object: ptr103, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Rule: ptr13}
	ptr105 := &system.Object{Description: "This is the native json bool data type", Id: system.Reference{Package: "kego.io/system", Name: "bool", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr106 := &system.Type{Object: ptr105, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Rule: ptr18}
	ptr107 := &system.Object{Description: "This is the int data type", Id: system.Reference{Package: "kego.io/system", Name: "int", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr108 := &system.Type{Object: ptr107, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr32}
	ptr109 := &system.Object{Description: "This is the native json object data type.", Id: system.Reference{Package: "kego.io/system", Name: "map", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr110 := &system.Type{Object: ptr109, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Rule: ptr43}
	ptr111 := &system.Object{Description: "This is the native json number data type", Id: system.Reference{Package: "kego.io/system", Name: "number", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr112 := &system.Type{Object: ptr111, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Rule: ptr63}
	ptr113 := &system.Object{Description: "This is the base type for the object interface. All ke objects have this type embedded.", Id: system.Reference{Package: "kego.io/system", Name: "object", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr114 := &system.Object{Description: "Description for the developer", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr115 := &system.Rule{Optional: true}
	ptr116 := &system.JsonStringRule{Object: ptr114, Rule: ptr115}
	ptr117 := &system.Object{Description: "All global objects should have an id.", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr118 := &system.Rule{Optional: true}
	ptr119 := &system.ReferenceRule{Object: ptr117, Rule: ptr118, Default: system.Reference{}}
	ptr120 := &system.Object{Description: "Extra validation rules for this object or descendants", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr121 := &system.Rule{Optional: true}
	ptr122 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr123 := &system.Rule{Interface: true}
	ptr124 := &system.RuleRule{Object: ptr122, Rule: ptr123}
	ptr125 := &system.ArrayRule{Object: ptr120, Rule: ptr121, Items: ptr124, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr126 := &system.Object{Description: "Type of the object.", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr127 := &system.Rule{}
	ptr128 := &system.ReferenceRule{Object: ptr126, Rule: ptr127, Default: system.Reference{}}
	ptr129 := &system.Type{Object: ptr113, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"description": ptr116, "id": ptr119, "rules": ptr125, "type": ptr128}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr130 := &system.Object{Description: "Package info - forms the root node of the package", Id: system.Reference{Package: "kego.io/system", Name: "package", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr131 := &system.Object{Description: "Map of import aliases used in this package: key = package path, value = alias.", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr132 := &system.Rule{Optional: true}
	ptr133 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr134 := &system.Rule{}
	ptr135 := &system.JsonStringRule{Object: ptr133, Rule: ptr134}
	ptr136 := &system.MapRule{Object: ptr131, Rule: ptr132, Items: ptr135, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr137 := &system.Type{Object: ptr130, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"aliases": ptr136}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr138 := &system.Object{Description: "This is a reference to another object, of the form: [local id] or [alias]:[id]", Id: system.Reference{Package: "kego.io/system", Name: "reference", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr139 := &system.Type{Object: ptr138, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr72}
	ptr140 := &system.Object{Description: "All rules will have this embedded in them.", Id: system.Reference{Package: "kego.io/system", Name: "rule", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr141 := &system.Object{Description: "Special field that should be excluded when marshaling - e.g. package.global", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr142 := &system.Rule{Optional: true}
	ptr143 := &system.JsonBoolRule{Object: ptr141, Rule: ptr142}
	ptr144 := &system.Object{Description: "Use the single method getter interface for this type", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr145 := &system.Rule{Optional: true}
	ptr146 := &system.JsonBoolRule{Object: ptr144, Rule: ptr145}
	ptr147 := &system.Object{Description: "If this rule is a field, this specifies that the field is optional", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr148 := &system.Rule{Optional: true}
	ptr149 := &system.JsonBoolRule{Object: ptr147, Rule: ptr148}
	ptr150 := &system.Object{Description: "Json selector defining what nodes this rule should be applied to.", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/json", Name: "@string", Exists: true}}
	ptr151 := &system.Rule{Optional: true}
	ptr152 := &system.JsonStringRule{Object: ptr150, Rule: ptr151}
	ptr153 := &system.Type{Object: ptr140, Basic: true, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"exclude": ptr143, "interface": ptr146, "optional": ptr149, "selector": ptr152}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	ptr154 := &system.Object{Description: "This is the native json string data type", Id: system.Reference{Package: "kego.io/system", Name: "string", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr155 := &system.Type{Object: ptr154, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface(nil), Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Rule: ptr100}
	ptr156 := &system.Object{Description: "This is the most basic type.", Id: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "type", Exists: true}}
	ptr157 := &system.Object{Description: "Basic types don't have system:object added by default to the embedded types.", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr158 := &system.Rule{Optional: true}
	ptr159 := &system.JsonBoolRule{Object: ptr157, Rule: ptr158}
	ptr160 := &system.Object{Description: "Types which this should embed - system:object is always added unless basic = true.", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr161 := &system.Rule{Optional: true}
	ptr162 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr163 := &system.Rule{}
	ptr164 := &system.ReferenceRule{Object: ptr162, Rule: ptr163, Default: system.Reference{}}
	ptr165 := &system.ArrayRule{Object: ptr160, Rule: ptr161, Items: ptr164, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr166 := &system.Object{Description: "Each field is listed with it's type", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@map", Exists: true}}
	ptr167 := &system.Rule{Optional: true}
	ptr168 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@rule", Exists: true}}
	ptr169 := &system.Rule{Interface: true}
	ptr170 := &system.RuleRule{Object: ptr168, Rule: ptr169}
	ptr171 := &system.MapRule{Object: ptr166, Rule: ptr167, Items: ptr170, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr172 := &system.Object{Description: "Is this type an interface?", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/json", Name: "@bool", Exists: true}}
	ptr173 := &system.Rule{Optional: true}
	ptr174 := &system.JsonBoolRule{Object: ptr172, Rule: ptr173}
	ptr175 := &system.Object{Description: "Array of interface types that this type should support", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@array", Exists: true}}
	ptr176 := &system.Rule{Optional: true}
	ptr177 := &system.Object{Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@reference", Exists: true}}
	ptr178 := &system.Rule{}
	ptr179 := &system.ReferenceRule{Object: ptr177, Rule: ptr178, Default: system.Reference{}}
	ptr180 := &system.ArrayRule{Object: ptr175, Rule: ptr176, Items: ptr179, MaxItems: system.Int{Value: 0}, MinItems: system.Int{Value: 0}}
	ptr181 := &system.Object{Description: "This is the native json type that represents this type. If omitted, default is object.", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@string", Exists: true}}
	ptr182 := &system.Rule{Optional: true}
	ptr183 := &system.StringRule{Object: ptr181, Rule: ptr182, Default: system.String{Value: "object", Exists: true}, Enum: []string{"string", "number", "bool", "array", "object", "map"}, Equal: system.String{}, Format: system.String{}, MaxLength: system.Int{Value: 0}, MinLength: system.Int{Value: 0}, Pattern: system.String{}}
	ptr184 := &system.Object{Description: "Type that defines restriction rules for this type.", Id: system.Reference{}, Rules: []system.RuleInterface(nil), Type: system.Reference{Package: "kego.io/system", Name: "@type", Exists: true}}
	ptr185 := &system.Rule{Optional: true}
	ptr186 := &system.TypeRule{Object: ptr184, Rule: ptr185}
	ptr187 := &system.Type{Object: ptr156, Embed: []system.Reference(nil), Fields: map[string]system.RuleInterface{"basic": ptr159, "embed": ptr165, "fields": ptr171, "interface": ptr174, "is": ptr180, "native": ptr183, "rule": ptr186}, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Rule: (*system.Type)(nil)}
	system.Register("kego.io/system", "@array", ptr13, 0x6eab209bb7c67261)
	system.Register("kego.io/system", "@bool", ptr18, 0x99f09d4d8377c2f6)
	system.Register("kego.io/system", "@int", ptr32, 0xcfdd76bad69c7b35)
	system.Register("kego.io/system", "@map", ptr43, 0x81bbbdddd4851344)
	system.Register("kego.io/system", "@number", ptr63, 0x392342d28da9f896)
	system.Register("kego.io/system", "@object", ptr65, 0x88b237630565f154)
	system.Register("kego.io/system", "@package", ptr67, 0x5168e36553335f0b)
	system.Register("kego.io/system", "@reference", ptr72, 0xecda4ba48ecb2d27)
	system.Register("kego.io/system", "@rule", ptr74, 0xb1194b97d6d9cd62)
	system.Register("kego.io/system", "@string", ptr100, 0xaeb6afcd8077ac5c)
	system.Register("kego.io/system", "@type", ptr102, 0x9647539a0b5c4b68)
	system.Register("kego.io/system", "array", ptr104, 0x6eab209bb7c67261)
	system.Register("kego.io/system", "bool", ptr106, 0x99f09d4d8377c2f6)
	system.Register("kego.io/system", "int", ptr108, 0xcfdd76bad69c7b35)
	system.Register("kego.io/system", "map", ptr110, 0x81bbbdddd4851344)
	system.Register("kego.io/system", "number", ptr112, 0x392342d28da9f896)
	system.Register("kego.io/system", "object", ptr129, 0x88b237630565f154)
	system.Register("kego.io/system", "package", ptr137, 0x5168e36553335f0b)
	system.Register("kego.io/system", "reference", ptr139, 0xecda4ba48ecb2d27)
	system.Register("kego.io/system", "rule", ptr153, 0xb1194b97d6d9cd62)
	system.Register("kego.io/system", "string", ptr155, 0xaeb6afcd8077ac5c)
	system.Register("kego.io/system", "type", ptr187, 0x9647539a0b5c4b68)
}
