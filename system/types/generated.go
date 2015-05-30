package types

import "kego.io/system"

func init() {

	ptr8728451904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728737696 := &system.Object{Context: ptr8728451904, Description: "This is the native json array data type", Id: "array", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728451584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728737808 := &system.Object{Context: ptr8728451584, Description: "Restriction rules for arrays", Id: "@array", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728451040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728739040 := &system.Object{Context: ptr8728451040, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728451008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728739152 := &system.Object{Context: ptr8728451008, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728544992 := &system.Int_rule{Object: ptr8728739152, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728603936 := &system.Property{Object: ptr8728739040, Defaulter: false, Item: ptr8728544992, Optional: true}

	ptr8728451456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728739264 := &system.Object{Context: ptr8728451456, Description: "If this is true, each item must be unique", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728451424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728739600 := &system.Object{Context: ptr8728451424, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728367232 := &system.Bool_rule{Object: ptr8728739600, Default: system.Bool{Value: false, Exists: true}}

	ptr8728604032 := &system.Property{Object: ptr8728739264, Defaulter: false, Item: ptr8728367232, Optional: true}

	ptr8728450272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728738144 := &system.Object{Context: ptr8728450272, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728450240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728738480 := &system.Object{Context: ptr8728450240, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728535448 := &system.Rule_rule{Object: ptr8728738480}

	ptr8728603024 := &system.Property{Object: ptr8728738144, Defaulter: false, Item: ptr8728535448, Optional: false}

	ptr8728450688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728738592 := &system.Object{Context: ptr8728450688, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728450656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728738928 := &system.Object{Context: ptr8728450656, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728544832 := &system.Int_rule{Object: ptr8728738928, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728603600 := &system.Property{Object: ptr8728738592, Defaulter: false, Item: ptr8728544832, Optional: true}

	ptr8728687824 := &system.Type{Object: ptr8728737808, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"maxItems": ptr8728603936, "uniqueItems": ptr8728604032, "items": ptr8728603024, "minItems": ptr8728603600}, Rule: (*system.Type)(nil)}

	ptr8728687536 := &system.Type{Object: ptr8728737696, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728687824}

	ptr8728449856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728897392 := &system.Object{Context: ptr8728449856, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728449472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728897840 := &system.Object{Context: ptr8728449472, Description: "Restriction rules for references", Id: "@reference", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728449120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728898400 := &system.Object{Context: ptr8728449120, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728449088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728898512 := &system.Object{Context: ptr8728449088, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728528768 := &system.Reference_rule{Object: ptr8728898512, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728603264 := &system.Property{Object: ptr8728898400, Defaulter: true, Item: ptr8728528768, Optional: true}

	ptr8728685664 := &system.Type{Object: ptr8728897840, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8728603264}, Rule: (*system.Type)(nil)}

	ptr8728685088 := &system.Type{Object: ptr8728897392, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728685664}

	ptr8728909120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728899744 := &system.Object{Context: ptr8728909120, Description: "Restriction rules for strings", Id: "@string", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728452416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728900304 := &system.Object{Context: ptr8728452416, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728452352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728900640 := &system.Object{Context: ptr8728452352, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728667072 := &system.String_rule{Object: ptr8728900640, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728604656 := &system.Property{Object: ptr8728900304, Defaulter: true, Item: ptr8728667072, Optional: true}

	ptr8728904352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728900752 := &system.Object{Context: ptr8728904352, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728904096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728901088 := &system.Object{Context: ptr8728904096, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728904032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728901200 := &system.Object{Context: ptr8728904032, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728667552 := &system.String_rule{Object: ptr8728901200, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728530048 := &system.Array_rule{Object: ptr8728901088, Items: ptr8728667552, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728605328 := &system.Property{Object: ptr8728900752, Defaulter: false, Item: ptr8728530048, Optional: true}

	ptr8728905216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728901312 := &system.Object{Context: ptr8728905216, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728905152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728901424 := &system.Object{Context: ptr8728905152, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728546032 := &system.Int_rule{Object: ptr8728901424, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728606000 := &system.Property{Object: ptr8728901312, Defaulter: false, Item: ptr8728546032, Optional: true}

	ptr8728905952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728901536 := &system.Object{Context: ptr8728905952, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728905920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728901648 := &system.Object{Context: ptr8728905920, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728544432 := &system.Int_rule{Object: ptr8728901648, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728606384 := &system.Property{Object: ptr8728901536, Defaulter: false, Item: ptr8728544432, Optional: true}

	ptr8728906752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728901760 := &system.Object{Context: ptr8728906752, Description: "This is a string that the value must match", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728906624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728901872 := &system.Object{Context: ptr8728906624, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728667712 := &system.String_rule{Object: ptr8728901872, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728606480 := &system.Property{Object: ptr8728901760, Defaulter: false, Item: ptr8728667712, Optional: true}

	ptr8728907456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728901984 := &system.Object{Context: ptr8728907456, Description: "This is a regex to match the value to", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728907328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728902096 := &system.Object{Context: ptr8728907328, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728667872 := &system.String_rule{Object: ptr8728902096, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728606720 := &system.Property{Object: ptr8728901984, Defaulter: false, Item: ptr8728667872, Optional: true}

	ptr8728908896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728902208 := &system.Object{Context: ptr8728908896, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728908704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728902320 := &system.Object{Context: ptr8728908704, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728668032 := &system.String_rule{Object: ptr8728902320, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728606864 := &system.Property{Object: ptr8728902208, Defaulter: false, Item: ptr8728668032, Optional: true}

	ptr8728688688 := &system.Type{Object: ptr8728899744, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8728604656, "enum": ptr8728605328, "minLength": ptr8728606000, "maxLength": ptr8728606384, "equal": ptr8728606480, "pattern": ptr8728606720, "format": ptr8728606864}, Rule: (*system.Type)(nil)}

	ptr8728906496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728732992 := &system.Object{Context: ptr8728906496, Description: "Restriction rules for integers", Id: "@int", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728905408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728733104 := &system.Object{Context: ptr8728905408, Description: "Default value if this property is omitted", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728905376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728733216 := &system.Object{Context: ptr8728905376, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728545472 := &system.Int_rule{Object: ptr8728733216, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728602400 := &system.Property{Object: ptr8728733104, Defaulter: true, Item: ptr8728545472, Optional: true}

	ptr8728905728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728733328 := &system.Object{Context: ptr8728905728, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728905696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728733440 := &system.Object{Context: ptr8728905696, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728545552 := &system.Int_rule{Object: ptr8728733440, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728602640 := &system.Property{Object: ptr8728733328, Defaulter: false, Item: ptr8728545552, Optional: true}

	ptr8728906048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728733552 := &system.Object{Context: ptr8728906048, Description: "This provides a lower bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728906016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728733664 := &system.Object{Context: ptr8728906016, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728545632 := &system.Int_rule{Object: ptr8728733664, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728602736 := &system.Property{Object: ptr8728733552, Defaulter: false, Item: ptr8728545632, Optional: true}

	ptr8728906368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728737024 := &system.Object{Context: ptr8728906368, Description: "This provides an upper bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728906336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728737136 := &system.Object{Context: ptr8728906336, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728545712 := &system.Int_rule{Object: ptr8728737136, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728603072 := &system.Property{Object: ptr8728737024, Defaulter: false, Item: ptr8728545712, Optional: true}

	ptr8728689264 := &system.Type{Object: ptr8728732992, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8728602400, "multipleOf": ptr8728602640, "minimum": ptr8728602736, "maximum": ptr8728603072}, Rule: (*system.Type)(nil)}

	ptr8728903904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728739712 := &system.Object{Context: ptr8728903904, Description: "This is the native json bool data type", Id: "bool", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728452992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728387696 := &system.Object{Context: ptr8728452992, Description: "Restriction rules for bools", Id: "@bool", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728452864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728895488 := &system.Object{Context: ptr8728452864, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728452832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728895600 := &system.Object{Context: ptr8728452832, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728368272 := &system.Bool_rule{Object: ptr8728895600, Default: system.Bool{Value: false, Exists: false}}

	ptr8728604848 := &system.Property{Object: ptr8728895488, Defaulter: true, Item: ptr8728368272, Optional: true}

	ptr8728688544 := &system.Type{Object: ptr8728387696, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8728604848}, Rule: (*system.Type)(nil)}

	ptr8728688112 := &system.Type{Object: ptr8728739712, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728688544}

	ptr8728445600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728899296 := &system.Object{Context: ptr8728445600, Description: "Automatically created basic rule for object", Id: "@object", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728685808 := &system.Type{Object: ptr8728899296, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728450528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728898848 := &system.Object{Context: ptr8728450528, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728689840 := &system.Type{Object: ptr8728898848, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728906816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728732880 := &system.Object{Context: ptr8728906816, Description: "This is the int data type", Id: "int", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728689120 := &system.Type{Object: ptr8728732880, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728689264}

	ptr8728446560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728900416 := &system.Object{Context: ptr8728446560, Description: "This is the most basic type.", Id: "type", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728906240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728900528 := &system.Object{Context: ptr8728906240, Description: "Type which this should extend", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728906176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728900864 := &system.Object{Context: ptr8728906176, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728527040 := &system.Reference_rule{Object: ptr8728900864, Default: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}}

	ptr8728600768 := &system.Property{Object: ptr8728900528, Defaulter: false, Item: ptr8728527040, Optional: true}

	ptr8728908320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728900976 := &system.Object{Context: ptr8728908320, Description: "Array of interface types that this type should support", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728908000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728902432 := &system.Object{Context: ptr8728908000, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728907968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728902544 := &system.Object{Context: ptr8728907968, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728527936 := &system.Reference_rule{Object: ptr8728902544, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728527808 := &system.Array_rule{Object: ptr8728902432, Items: ptr8728527936, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728600816 := &system.Property{Object: ptr8728900976, Defaulter: false, Item: ptr8728527808, Optional: true}

	ptr8728910144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728902656 := &system.Object{Context: ptr8728910144, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728910048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728902768 := &system.Object{Context: ptr8728910048, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728666272 := &system.String_rule{Object: ptr8728902768, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728600960 := &system.Property{Object: ptr8728902656, Defaulter: false, Item: ptr8728666272, Optional: true}

	ptr8728910944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728902880 := &system.Object{Context: ptr8728910944, Description: "Is this type an interface?", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728910912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728902992 := &system.Object{Context: ptr8728910912, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728535648 := &system.JsonBool_rule{Object: ptr8728902992}

	ptr8728601056 := &system.Property{Object: ptr8728902880, Defaulter: false, Item: ptr8728535648, Optional: true}

	ptr8728911520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728903104 := &system.Object{Context: ptr8728911520, Description: "Exclude this type from the generated json?", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728911488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728903216 := &system.Object{Context: ptr8728911488, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728535720 := &system.JsonBool_rule{Object: ptr8728903216}

	ptr8728601104 := &system.Property{Object: ptr8728903104, Defaulter: false, Item: ptr8728535720, Optional: true}

	ptr8728445984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728903328 := &system.Object{Context: ptr8728445984, Description: "Each field is listed with it's type", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728445952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728903440 := &system.Object{Context: ptr8728445952, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728445888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728731872 := &system.Object{Context: ptr8728445888, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@property", Package: "kego.io/system", Type: "@property", Exists: true}}

	ptr8728535752 := &system.Property_rule{Object: ptr8728731872}

	ptr8728528064 := &system.Map_rule{Object: ptr8728903440, Items: ptr8728535752, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728601776 := &system.Property{Object: ptr8728903328, Defaulter: false, Item: ptr8728528064, Optional: true}

	ptr8728446432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728735792 := &system.Object{Context: ptr8728446432, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728446400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728736128 := &system.Object{Context: ptr8728446400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@type", Package: "kego.io/system", Type: "@type", Exists: true}}

	ptr8728536256 := &system.Type_rule{Object: ptr8728736128}

	ptr8728604944 := &system.Property{Object: ptr8728735792, Defaulter: false, Item: ptr8728536256, Optional: true}

	ptr8728684224 := &system.Type{Object: ptr8728900416, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"extends": ptr8728600768, "is": ptr8728600816, "native": ptr8728600960, "interface": ptr8728601056, "exclude": ptr8728601104, "properties": ptr8728601776, "rule": ptr8728604944}, Rule: (*system.Type)(nil)}

	ptr8728907904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728739376 := &system.Object{Context: ptr8728907904, Description: "Restriction rules for numbers", Id: "@number", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728905120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728895936 := &system.Object{Context: ptr8728905120, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728904832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728896048 := &system.Object{Context: ptr8728904832, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728543472 := &system.Number_rule{Object: ptr8728896048, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728601920 := &system.Property{Object: ptr8728895936, Defaulter: false, Item: ptr8728543472, Optional: true}

	ptr8728905856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728896160 := &system.Object{Context: ptr8728905856, Description: "This provides a lower bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728905600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728896272 := &system.Object{Context: ptr8728905600, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728543552 := &system.Number_rule{Object: ptr8728896272, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728602208 := &system.Property{Object: ptr8728896160, Defaulter: false, Item: ptr8728543552, Optional: true}

	ptr8728906464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728896384 := &system.Object{Context: ptr8728906464, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728906432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728896496 := &system.Object{Context: ptr8728906432, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728367040 := &system.Bool_rule{Object: ptr8728896496, Default: system.Bool{Value: false, Exists: true}}

	ptr8728602688 := &system.Property{Object: ptr8728896384, Defaulter: false, Item: ptr8728367040, Optional: true}

	ptr8728907168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728896608 := &system.Object{Context: ptr8728907168, Description: "This provides an upper bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728906880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728896720 := &system.Object{Context: ptr8728906880, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728543632 := &system.Number_rule{Object: ptr8728896720, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728603168 := &system.Property{Object: ptr8728896608, Defaulter: false, Item: ptr8728543632, Optional: true}

	ptr8728907680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728896832 := &system.Object{Context: ptr8728907680, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728907648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728896944 := &system.Object{Context: ptr8728907648, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728367824 := &system.Bool_rule{Object: ptr8728896944, Default: system.Bool{Value: false, Exists: true}}

	ptr8728603504 := &system.Property{Object: ptr8728896832, Defaulter: false, Item: ptr8728367824, Optional: true}

	ptr8728904256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728739488 := &system.Object{Context: ptr8728904256, Description: "Default value if this property is omitted", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728903968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728895824 := &system.Object{Context: ptr8728903968, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728546112 := &system.Number_rule{Object: ptr8728895824, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728604128 := &system.Property{Object: ptr8728739488, Defaulter: true, Item: ptr8728546112, Optional: true}

	ptr8728690128 := &system.Type{Object: ptr8728739376, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"multipleOf": ptr8728601920, "minimum": ptr8728602208, "exclusiveMinimum": ptr8728602688, "maximum": ptr8728603168, "exclusiveMaximum": ptr8728603504, "default": ptr8728604128}, Rule: (*system.Type)(nil)}

	ptr8728451168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728898960 := &system.Object{Context: ptr8728451168, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728686960 := &system.Type{Object: ptr8728898960, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728909920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728899632 := &system.Object{Context: ptr8728909920, Description: "This is the native json string data type", Id: "string", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728687968 := &system.Type{Object: ptr8728899632, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728688688}

	ptr8728447200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728736240 := &system.Object{Context: ptr8728447200, Description: "Automatically created basic rule for type", Id: "@type", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728684656 := &system.Type{Object: ptr8728736240, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728445408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728897056 := &system.Object{Context: ptr8728445408, Description: "This is the most basic type.", Id: "object", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728910880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728898064 := &system.Object{Context: ptr8728910880, Description: "Description for the developer", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728910848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728898176 := &system.Object{Context: ptr8728910848, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728536016 := &system.JsonString_rule{Object: ptr8728898176}

	ptr8728605376 := &system.Property{Object: ptr8728898064, Defaulter: false, Item: ptr8728536016, Optional: true}

	ptr8728911296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728898288 := &system.Object{Context: ptr8728911296, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728911264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728898624 := &system.Object{Context: ptr8728911264, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@context", Package: "kego.io/system", Type: "@context", Exists: true}}

	ptr8728536040 := &system.Context_rule{Object: ptr8728898624}

	ptr8728605424 := &system.Property{Object: ptr8728898288, Defaulter: false, Item: ptr8728536040, Optional: true}

	ptr8728911840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728898736 := &system.Object{Context: ptr8728911840, Description: "Extra validation rules for this object or descendants", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728911808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728899072 := &system.Object{Context: ptr8728911808, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728911776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728899184 := &system.Object{Context: ptr8728911776, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728536192 := &system.Rule_rule{Object: ptr8728899184}

	ptr8728529088 := &system.Map_rule{Object: ptr8728899072, Items: ptr8728536192, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728605712 := &system.Property{Object: ptr8728898736, Defaulter: false, Item: ptr8728529088, Optional: true}

	ptr8728909856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728897168 := &system.Object{Context: ptr8728909856, Description: "Type of the object.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728909824 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728897504 := &system.Object{Context: ptr8728909824, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728528832 := &system.Reference_rule{Object: ptr8728897504, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728604464 := &system.Property{Object: ptr8728897168, Defaulter: false, Item: ptr8728528832, Optional: false}

	ptr8728910432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728897616 := &system.Object{Context: ptr8728910432, Description: "All global objects should have an id.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728910400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728897952 := &system.Object{Context: ptr8728910400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728535912 := &system.JsonString_rule{Object: ptr8728897952}

	ptr8728605040 := &system.Property{Object: ptr8728897616, Defaulter: false, Item: ptr8728535912, Optional: true}

	ptr8728685520 := &system.Type{Object: ptr8728897056, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"description": ptr8728605376, "context": ptr8728605424, "rules": ptr8728605712, "type": ptr8728604464, "id": ptr8728605040}, Rule: (*system.Type)(nil)}

	ptr8728445920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728899408 := &system.Object{Context: ptr8728445920, Description: "This is a property of a type", Id: "property", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728447680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728899520 := &system.Object{Context: ptr8728447680, Description: "This specifies that the field is optional", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728447648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728899856 := &system.Object{Context: ptr8728447648, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728536248 := &system.JsonBool_rule{Object: ptr8728899856}

	ptr8728606288 := &system.Property{Object: ptr8728899520, Defaulter: false, Item: ptr8728536248, Optional: true}

	ptr8728449248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728899968 := &system.Object{Context: ptr8728449248, Description: "This specifies that the field is the default value for a rule", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728449216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728900080 := &system.Object{Context: ptr8728449216, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728536320 := &system.JsonBool_rule{Object: ptr8728900080}

	ptr8728606624 := &system.Property{Object: ptr8728899968, Defaulter: false, Item: ptr8728536320, Optional: true}

	ptr8728445792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728900192 := &system.Object{Context: ptr8728445792, Description: "This is a rule object, defining the type and restrictions on the value of the this property", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728445760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728895712 := &system.Object{Context: ptr8728445760, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728535168 := &system.Rule_rule{Object: ptr8728895712}

	ptr8728606672 := &system.Property{Object: ptr8728900192, Defaulter: false, Item: ptr8728535168, Optional: false}

	ptr8728686672 := &system.Type{Object: ptr8728899408, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"optional": ptr8728606288, "defaulter": ptr8728606624, "item": ptr8728606672}, Rule: (*system.Type)(nil)}

	ptr8728908416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728738816 := &system.Object{Context: ptr8728908416, Description: "This is the native json number data type", Id: "number", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728689984 := &system.Type{Object: ptr8728738816, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728690128}

	ptr8728904640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728732768 := &system.Object{Context: ptr8728904640, Description: "Automatically created basic rule for context", Id: "@context", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728688400 := &system.Type{Object: ptr8728732768, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728908928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728737248 := &system.Object{Context: ptr8728908928, Description: "This is the native json object data type.", Id: "map", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728908608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728737360 := &system.Object{Context: ptr8728908608, Description: "Restriction rules for maps", Id: "@map", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728907776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728737584 := &system.Object{Context: ptr8728907776, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728907744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728737920 := &system.Object{Context: ptr8728907744, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728535784 := &system.Rule_rule{Object: ptr8728737920}

	ptr8728603456 := &system.Property{Object: ptr8728737584, Defaulter: false, Item: ptr8728535784, Optional: false}

	ptr8728908128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728738032 := &system.Object{Context: ptr8728908128, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728908096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728738256 := &system.Object{Context: ptr8728908096, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728545872 := &system.Int_rule{Object: ptr8728738256, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728603552 := &system.Property{Object: ptr8728738032, Defaulter: false, Item: ptr8728545872, Optional: true}

	ptr8728908480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728738368 := &system.Object{Context: ptr8728908480, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728908448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728738704 := &system.Object{Context: ptr8728908448, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728545952 := &system.Int_rule{Object: ptr8728738704, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728603696 := &system.Property{Object: ptr8728738368, Defaulter: false, Item: ptr8728545952, Optional: true}

	ptr8728689696 := &system.Type{Object: ptr8728737360, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"items": ptr8728603456, "minItems": ptr8728603552, "maxItems": ptr8728603696}, Rule: (*system.Type)(nil)}

	ptr8728689552 := &system.Type{Object: ptr8728737248, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728689696}

	ptr8728904448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728732656 := &system.Object{Context: ptr8728904448, Description: "Unmarshal context.", Id: "context", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728686240 := &system.Type{Object: ptr8728732656, Exclude: true, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728446976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728897280 := &system.Object{Context: ptr8728446976, Description: "Automatically created basic rule for property", Id: "@property", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728684080 := &system.Type{Object: ptr8728897280, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system:@array", ptr8728687824)

	system.RegisterType("kego.io/system:@bool", ptr8728688544)

	system.RegisterType("kego.io/system:@context", ptr8728688400)

	system.RegisterType("kego.io/system:@int", ptr8728689264)

	system.RegisterType("kego.io/system:@map", ptr8728689696)

	system.RegisterType("kego.io/system:@number", ptr8728690128)

	system.RegisterType("kego.io/system:@object", ptr8728685808)

	system.RegisterType("kego.io/system:@property", ptr8728684080)

	system.RegisterType("kego.io/system:@reference", ptr8728685664)

	system.RegisterType("kego.io/system:@rule", ptr8728686960)

	system.RegisterType("kego.io/system:@string", ptr8728688688)

	system.RegisterType("kego.io/system:@type", ptr8728684656)

	system.RegisterType("kego.io/system:array", ptr8728687536)

	system.RegisterType("kego.io/system:bool", ptr8728688112)

	system.RegisterType("kego.io/system:context", ptr8728686240)

	system.RegisterType("kego.io/system:int", ptr8728689120)

	system.RegisterType("kego.io/system:map", ptr8728689552)

	system.RegisterType("kego.io/system:number", ptr8728689984)

	system.RegisterType("kego.io/system:object", ptr8728685520)

	system.RegisterType("kego.io/system:property", ptr8728686672)

	system.RegisterType("kego.io/system:reference", ptr8728685088)

	system.RegisterType("kego.io/system:rule", ptr8728689840)

	system.RegisterType("kego.io/system:string", ptr8728687968)

	system.RegisterType("kego.io/system:type", ptr8728684224)

}
