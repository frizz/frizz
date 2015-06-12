package types

import "kego.io/system"

func init() {

	ptr8729505024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729247968 := &system.Object{Context: ptr8729505024, Description: "This is the int data type", Id: "int", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729504704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729248080 := &system.Object{Context: ptr8729504704, Description: "Restriction rules for integers", Id: "@int", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729503616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729248416 := &system.Object{Context: ptr8729503616, Description: "Default value if this property is omitted", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729503584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729248864 := &system.Object{Context: ptr8729503584, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729395760 := &system.Int_rule{Object: ptr8729248864, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729069392 := &system.Property{Object: ptr8729248416, Defaulter: true, Item: ptr8729395760, Optional: true}

	ptr8729503936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729249088 := &system.Object{Context: ptr8729503936, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729503904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729249312 := &system.Object{Context: ptr8729503904, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729395840 := &system.Int_rule{Object: ptr8729249312, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729069776 := &system.Property{Object: ptr8729249088, Defaulter: false, Item: ptr8729395840, Optional: true}

	ptr8729504256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729249424 := &system.Object{Context: ptr8729504256, Description: "This provides a lower bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729504224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729249760 := &system.Object{Context: ptr8729504224, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729395920 := &system.Int_rule{Object: ptr8729249760, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729070112 := &system.Property{Object: ptr8729249424, Defaulter: false, Item: ptr8729395920, Optional: true}

	ptr8729504576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729249984 := &system.Object{Context: ptr8729504576, Description: "This provides an upper bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729504544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729250544 := &system.Object{Context: ptr8729504544, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729396000 := &system.Int_rule{Object: ptr8729250544, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729070400 := &system.Property{Object: ptr8729249984, Defaulter: false, Item: ptr8729396000, Optional: true}

	ptr8729281520 := &system.Type{Object: ptr8729248080, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729069392, "multipleOf": ptr8729069776, "minimum": ptr8729070112, "maximum": ptr8729070400}, Rule: (*system.Type)(nil)}

	ptr8729281376 := &system.Type{Object: ptr8729247968, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729281520}

	ptr8728917856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729254352 := &system.Object{Context: ptr8728917856, Description: "This is the most basic type.", Id: "object", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728913920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729254576 := &system.Object{Context: ptr8728913920, Description: "Type of the object.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728912032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729252672 := &system.Object{Context: ptr8728912032, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729364800 := &system.Reference_rule{Object: ptr8729252672, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729074048 := &system.Property{Object: ptr8729254576, Defaulter: false, Item: ptr8729364800, Optional: false}

	ptr8728914400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729252784 := &system.Object{Context: ptr8728914400, Description: "All global objects should have an id.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728914368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729255024 := &system.Object{Context: ptr8728914368, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8729002056 := &system.JsonString_rule{Object: ptr8729255024}

	ptr8729067808 := &system.Property{Object: ptr8729252784, Defaulter: false, Item: ptr8729002056, Optional: true}

	ptr8728914848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729255136 := &system.Object{Context: ptr8728914848, Description: "Description for the developer", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728914816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729255248 := &system.Object{Context: ptr8728914816, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8729002320 := &system.JsonString_rule{Object: ptr8729255248}

	ptr8729069200 := &system.Property{Object: ptr8729255136, Defaulter: false, Item: ptr8729002320, Optional: true}

	ptr8728915904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729255360 := &system.Object{Context: ptr8728915904, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728915712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729255696 := &system.Object{Context: ptr8728915712, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@context", Package: "kego.io/system", Type: "@context", Exists: true}}

	ptr8729002344 := &system.Context_rule{Object: ptr8729255696}

	ptr8729069680 := &system.Property{Object: ptr8729255360, Defaulter: false, Item: ptr8729002344, Optional: true}

	ptr8728917600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729255808 := &system.Object{Context: ptr8728917600, Description: "Extra validation rules for this object or descendants", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728917472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729198816 := &system.Object{Context: ptr8728917472, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728917440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729198928 := &system.Object{Context: ptr8728917440, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729002744 := &system.Rule_rule{Object: ptr8729198928}

	ptr8729362560 := &system.Map_rule{Object: ptr8729198816, Items: ptr8729002744, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729071360 := &system.Property{Object: ptr8729255808, Defaulter: false, Item: ptr8729362560, Optional: true}

	ptr8729282960 := &system.Type{Object: ptr8729254352, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"type": ptr8729074048, "id": ptr8729067808, "description": ptr8729069200, "context": ptr8729069680, "rules": ptr8729071360}, Rule: (*system.Type)(nil)}

	ptr8729506848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729200832 := &system.Object{Context: ptr8729506848, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729283248 := &system.Type{Object: ptr8729200832, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729569024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729203296 := &system.Object{Context: ptr8729569024, Description: "Restriction rules for strings", Id: "@string", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729508608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729203408 := &system.Object{Context: ptr8729508608, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729508576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729201280 := &system.Object{Context: ptr8729508576, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729136736 := &system.String_rule{Object: ptr8729201280, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729074864 := &system.Property{Object: ptr8729203408, Defaulter: true, Item: ptr8729136736, Optional: true}

	ptr8729509568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729201392 := &system.Object{Context: ptr8729509568, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729509472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729201728 := &system.Object{Context: ptr8729509472, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729509408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729201840 := &system.Object{Context: ptr8729509408, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729137696 := &system.String_rule{Object: ptr8729201840, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729370240 := &system.Array_rule{Object: ptr8729201728, Items: ptr8729137696, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729075248 := &system.Property{Object: ptr8729201392, Defaulter: false, Item: ptr8729370240, Optional: true}

	ptr8729567360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729201952 := &system.Object{Context: ptr8729567360, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729567328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729198592 := &system.Object{Context: ptr8729567328, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729395200 := &system.Int_rule{Object: ptr8729198592, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729031264 := &system.Property{Object: ptr8729201952, Defaulter: false, Item: ptr8729395200, Optional: true}

	ptr8729567680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729198704 := &system.Object{Context: ptr8729567680, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729567648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729199376 := &system.Object{Context: ptr8729567648, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729395360 := &system.Int_rule{Object: ptr8729199376, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729027664 := &system.Property{Object: ptr8729198704, Defaulter: false, Item: ptr8729395360, Optional: true}

	ptr8729568000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729199488 := &system.Object{Context: ptr8729568000, Description: "This is a string that the value must match", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729567968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729201056 := &system.Object{Context: ptr8729567968, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729133056 := &system.String_rule{Object: ptr8729201056, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729027760 := &system.Property{Object: ptr8729199488, Defaulter: false, Item: ptr8729133056, Optional: true}

	ptr8729568320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729201168 := &system.Object{Context: ptr8729568320, Description: "This is a regex to match the value to", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729568288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729201504 := &system.Object{Context: ptr8729568288, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729135296 := &system.String_rule{Object: ptr8729201504, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729027904 := &system.Property{Object: ptr8729201168, Defaulter: false, Item: ptr8729135296, Optional: true}

	ptr8729568896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729201616 := &system.Object{Context: ptr8729568896, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729568864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729202176 := &system.Object{Context: ptr8729568864, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729135936 := &system.String_rule{Object: ptr8729202176, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729028048 := &system.Property{Object: ptr8729201616, Defaulter: false, Item: ptr8729135936, Optional: true}

	ptr8729284688 := &system.Type{Object: ptr8729203296, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729074864, "enum": ptr8729075248, "minLength": ptr8729031264, "maxLength": ptr8729027664, "equal": ptr8729027760, "pattern": ptr8729027904, "format": ptr8729028048}, Rule: (*system.Type)(nil)}

	ptr8729573856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729205424 := &system.Object{Context: ptr8729573856, Description: "Automatically created basic rule for type", Id: "@type", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729284976 := &system.Type{Object: ptr8729205424, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729502784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729247856 := &system.Object{Context: ptr8729502784, Description: "Automatically created basic rule for context", Id: "@context", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729280800 := &system.Type{Object: ptr8729247856, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729504160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729200160 := &system.Object{Context: ptr8729504160, Description: "Automatically created basic rule for property", Id: "@property", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729281232 := &system.Type{Object: ptr8729200160, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729573440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729202288 := &system.Object{Context: ptr8729573440, Description: "This is the most basic type.", Id: "type", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729570304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729202400 := &system.Object{Context: ptr8729570304, Description: "Type which this should extend", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729570272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729202512 := &system.Object{Context: ptr8729570272, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729363328 := &system.Reference_rule{Object: ptr8729202512, Default: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}}

	ptr8729028336 := &system.Property{Object: ptr8729202400, Defaulter: false, Item: ptr8729363328, Optional: true}

	ptr8729570848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729202624 := &system.Object{Context: ptr8729570848, Description: "Array of interface types that this type should support", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729570752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729202736 := &system.Object{Context: ptr8729570752, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729570720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729202848 := &system.Object{Context: ptr8729570720, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729363648 := &system.Reference_rule{Object: ptr8729202848, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729363520 := &system.Array_rule{Object: ptr8729202736, Items: ptr8729363648, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729028384 := &system.Property{Object: ptr8729202624, Defaulter: false, Item: ptr8729363520, Optional: true}

	ptr8729571488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729203632 := &system.Object{Context: ptr8729571488, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729571456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729203744 := &system.Object{Context: ptr8729571456, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729136416 := &system.String_rule{Object: ptr8729203744, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729028480 := &system.Property{Object: ptr8729203632, Defaulter: false, Item: ptr8729136416, Optional: true}

	ptr8729571936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729203856 := &system.Object{Context: ptr8729571936, Description: "Is this type an interface?", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729571904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729202064 := &system.Object{Context: ptr8729571904, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729003056 := &system.JsonBool_rule{Object: ptr8729202064}

	ptr8729028528 := &system.Property{Object: ptr8729203856, Defaulter: false, Item: ptr8729003056, Optional: true}

	ptr8729572352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729203072 := &system.Object{Context: ptr8729572352, Description: "Exclude this type from the generated json?", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729572320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729203968 := &system.Object{Context: ptr8729572320, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729003080 := &system.JsonBool_rule{Object: ptr8729203968}

	ptr8729028576 := &system.Property{Object: ptr8729203072, Defaulter: false, Item: ptr8729003080, Optional: true}

	ptr8729572896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729204192 := &system.Object{Context: ptr8729572896, Description: "Each field is listed with it's type", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729572864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729204304 := &system.Object{Context: ptr8729572864, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729572832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729204752 := &system.Object{Context: ptr8729572832, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@property", Package: "kego.io/system", Type: "@property", Exists: true}}

	ptr8729003112 := &system.Property_rule{Object: ptr8729204752}

	ptr8729363840 := &system.Map_rule{Object: ptr8729204304, Items: ptr8729003112, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729028672 := &system.Property{Object: ptr8729204192, Defaulter: false, Item: ptr8729363840, Optional: true}

	ptr8729573312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729204976 := &system.Object{Context: ptr8729573312, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729573280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729205312 := &system.Object{Context: ptr8729573280, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@type", Package: "kego.io/system", Type: "@type", Exists: true}}

	ptr8729003248 := &system.Type_rule{Object: ptr8729205312}

	ptr8729034480 := &system.Property{Object: ptr8729204976, Defaulter: false, Item: ptr8729003248, Optional: true}

	ptr8729282816 := &system.Type{Object: ptr8729202288, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"extends": ptr8729028336, "is": ptr8729028384, "native": ptr8729028480, "interface": ptr8729028528, "exclude": ptr8729028576, "properties": ptr8729028672, "rule": ptr8729034480}, Rule: (*system.Type)(nil)}

	ptr8728918656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729199040 := &system.Object{Context: ptr8728918656, Description: "Automatically created basic rule for object", Id: "@object", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729280512 := &system.Type{Object: ptr8729199040, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728919808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729248640 := &system.Object{Context: ptr8728919808, Description: "This is the native json array data type", Id: "array", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728919488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729248752 := &system.Object{Context: ptr8728919488, Description: "Restriction rules for arrays", Id: "@array", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728918592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729249648 := &system.Object{Context: ptr8728918592, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728918560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729250096 := &system.Object{Context: ptr8728918560, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729399920 := &system.Int_rule{Object: ptr8729250096, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729072656 := &system.Property{Object: ptr8729249648, Defaulter: false, Item: ptr8729399920, Optional: true}

	ptr8728918944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729250208 := &system.Object{Context: ptr8728918944, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728918912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729250320 := &system.Object{Context: ptr8728918912, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729400000 := &system.Int_rule{Object: ptr8729250320, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729072992 := &system.Property{Object: ptr8729250208, Defaulter: false, Item: ptr8729400000, Optional: true}

	ptr8728919360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729250432 := &system.Object{Context: ptr8728919360, Description: "If this is true, each item must be unique", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728919328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729250880 := &system.Object{Context: ptr8728919328, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728835104 := &system.Bool_rule{Object: ptr8729250880, Default: system.Bool{Value: false, Exists: true}}

	ptr8729073088 := &system.Property{Object: ptr8729250432, Defaulter: false, Item: ptr8728835104, Optional: true}

	ptr8728918176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729249200 := &system.Object{Context: ptr8728918176, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728918144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729249536 := &system.Object{Context: ptr8728918144, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729002432 := &system.Rule_rule{Object: ptr8729249536}

	ptr8729072080 := &system.Property{Object: ptr8729249200, Defaulter: false, Item: ptr8729002432, Optional: false}

	ptr8729284112 := &system.Type{Object: ptr8729248752, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"minItems": ptr8729072656, "maxItems": ptr8729072992, "uniqueItems": ptr8729073088, "items": ptr8729072080}, Rule: (*system.Type)(nil)}

	ptr8729283968 := &system.Type{Object: ptr8729248640, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729284112}

	ptr8729507488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729200944 := &system.Object{Context: ptr8729507488, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729283392 := &system.Type{Object: ptr8729200944, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729502848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729250992 := &system.Object{Context: ptr8729502848, Description: "This is the native json bool data type", Id: "bool", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729502528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729251104 := &system.Object{Context: ptr8729502528, Description: "Restriction rules for bools", Id: "@bool", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729502400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729251216 := &system.Object{Context: ptr8729502400, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729502368 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729251328 := &system.Object{Context: ptr8729502368, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728836144 := &system.Bool_rule{Object: ptr8729251328, Default: system.Bool{Value: false, Exists: false}}

	ptr8729073712 := &system.Property{Object: ptr8729251216, Defaulter: true, Item: ptr8728836144, Optional: true}

	ptr8729284544 := &system.Type{Object: ptr8729251104, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729073712}, Rule: (*system.Type)(nil)}

	ptr8729284400 := &system.Type{Object: ptr8729250992, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729284544}

	ptr8728917536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729252336 := &system.Object{Context: ptr8728917536, Description: "This is the native json number data type", Id: "number", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728916928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729252448 := &system.Object{Context: ptr8728916928, Description: "Restriction rules for numbers", Id: "@number", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729509440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729253344 := &system.Object{Context: ptr8729509440, Description: "This provides a lower bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729509280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729253456 := &system.Object{Context: ptr8729509280, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729396560 := &system.Number_rule{Object: ptr8729253456, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729073232 := &system.Property{Object: ptr8729253344, Defaulter: false, Item: ptr8729396560, Optional: true}

	ptr8729509792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729253568 := &system.Object{Context: ptr8729509792, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729509760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729253680 := &system.Object{Context: ptr8729509760, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729444816 := &system.Bool_rule{Object: ptr8729253680, Default: system.Bool{Value: false, Exists: true}}

	ptr8729073376 := &system.Property{Object: ptr8729253568, Defaulter: false, Item: ptr8729444816, Optional: true}

	ptr8728915776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729253792 := &system.Object{Context: ptr8728915776, Description: "This provides an upper bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728915104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729253904 := &system.Object{Context: ptr8728915104, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729396640 := &system.Number_rule{Object: ptr8729253904, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729073472 := &system.Property{Object: ptr8729253792, Defaulter: false, Item: ptr8729396640, Optional: true}

	ptr8728916608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729254016 := &system.Object{Context: ptr8728916608, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728916544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729254128 := &system.Object{Context: ptr8728916544, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729445232 := &system.Bool_rule{Object: ptr8729254128, Default: system.Bool{Value: false, Exists: true}}

	ptr8729074000 := &system.Property{Object: ptr8729254016, Defaulter: false, Item: ptr8729445232, Optional: true}

	ptr8729508544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729252560 := &system.Object{Context: ptr8729508544, Description: "Default value if this property is omitted", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729508384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729252896 := &system.Object{Context: ptr8729508384, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729396400 := &system.Number_rule{Object: ptr8729252896, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729072272 := &system.Property{Object: ptr8729252560, Defaulter: true, Item: ptr8729396400, Optional: true}

	ptr8729508992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729253120 := &system.Object{Context: ptr8729508992, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729508832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729253232 := &system.Object{Context: ptr8729508832, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729396480 := &system.Number_rule{Object: ptr8729253232, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729073040 := &system.Property{Object: ptr8729253120, Defaulter: false, Item: ptr8729396480, Optional: true}

	ptr8729282384 := &system.Type{Object: ptr8729252448, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"minimum": ptr8729073232, "exclusiveMinimum": ptr8729073376, "maximum": ptr8729073472, "exclusiveMaximum": ptr8729074000, "default": ptr8729072272, "multipleOf": ptr8729073040}, Rule: (*system.Type)(nil)}

	ptr8729282240 := &system.Type{Object: ptr8729252336, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729282384}

	ptr8729503488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729199152 := &system.Object{Context: ptr8729503488, Description: "This is a property of a type", Id: "property", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728919840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729199264 := &system.Object{Context: ptr8728919840, Description: "This specifies that the field is optional", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728919776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729199600 := &system.Object{Context: ptr8728919776, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729002800 := &system.JsonBool_rule{Object: ptr8729199600}

	ptr8729072800 := &system.Property{Object: ptr8729199264, Defaulter: false, Item: ptr8729002800, Optional: true}

	ptr8729502464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729199712 := &system.Object{Context: ptr8729502464, Description: "This specifies that the field is the default value for a rule", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729502304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729199824 := &system.Object{Context: ptr8729502304, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729002960 := &system.JsonBool_rule{Object: ptr8729199824}

	ptr8729073808 := &system.Property{Object: ptr8729199712, Defaulter: false, Item: ptr8729002960, Optional: true}

	ptr8729503296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729199936 := &system.Object{Context: ptr8729503296, Description: "This is a rule object, defining the type and restrictions on the value of the this property", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729503232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729200048 := &system.Object{Context: ptr8729503232, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729002984 := &system.Rule_rule{Object: ptr8729200048}

	ptr8729073856 := &system.Property{Object: ptr8729199936, Defaulter: false, Item: ptr8729002984, Optional: false}

	ptr8729281088 := &system.Type{Object: ptr8729199152, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"optional": ptr8729072800, "defaulter": ptr8729073808, "item": ptr8729073856}, Rule: (*system.Type)(nil)}

	ptr8729506080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729200272 := &system.Object{Context: ptr8729506080, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729505536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729200496 := &system.Object{Context: ptr8729505536, Description: "Restriction rules for references", Id: "@reference", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729505376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729200608 := &system.Object{Context: ptr8729505376, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729505344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729200720 := &system.Object{Context: ptr8729505344, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729363712 := &system.Reference_rule{Object: ptr8729200720, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729074336 := &system.Property{Object: ptr8729200608, Defaulter: true, Item: ptr8729363712, Optional: true}

	ptr8729284256 := &system.Type{Object: ptr8729200496, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729074336}, Rule: (*system.Type)(nil)}

	ptr8729282096 := &system.Type{Object: ptr8729200272, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729284256}

	ptr8729506816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729251440 := &system.Object{Context: ptr8729506816, Description: "Restriction rules for maps", Id: "@map", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729506336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729251776 := &system.Object{Context: ptr8729506336, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729506304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729252000 := &system.Object{Context: ptr8729506304, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729396160 := &system.Int_rule{Object: ptr8729252000, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729071696 := &system.Property{Object: ptr8729251776, Defaulter: false, Item: ptr8729396160, Optional: true}

	ptr8729506688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729252112 := &system.Object{Context: ptr8729506688, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729506656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729252224 := &system.Object{Context: ptr8729506656, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729396240 := &system.Int_rule{Object: ptr8729252224, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729071792 := &system.Property{Object: ptr8729252112, Defaulter: false, Item: ptr8729396240, Optional: true}

	ptr8729505984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729251552 := &system.Object{Context: ptr8729505984, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729505952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729251664 := &system.Object{Context: ptr8729505952, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729002704 := &system.Rule_rule{Object: ptr8729251664}

	ptr8729071600 := &system.Property{Object: ptr8729251552, Defaulter: false, Item: ptr8729002704, Optional: false}

	ptr8729281952 := &system.Type{Object: ptr8729251440, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"minItems": ptr8729071696, "maxItems": ptr8729071792, "items": ptr8729071600}, Rule: (*system.Type)(nil)}

	ptr8729507136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729250768 := &system.Object{Context: ptr8729507136, Description: "This is the native json object data type.", Id: "map", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729281808 := &system.Type{Object: ptr8729250768, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729281952}

	ptr8729502496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729247744 := &system.Object{Context: ptr8729502496, Description: "Unmarshal context.", Id: "context", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729280656 := &system.Type{Object: ptr8729247744, Exclude: true, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729569344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729203184 := &system.Object{Context: ptr8729569344, Description: "This is the native json string data type", Id: "string", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729283824 := &system.Type{Object: ptr8729203184, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729284688}

	system.RegisterType("kego.io/system:@array", ptr8729284112)

	system.RegisterType("kego.io/system:@bool", ptr8729284544)

	system.RegisterType("kego.io/system:@context", ptr8729280800)

	system.RegisterType("kego.io/system:@int", ptr8729281520)

	system.RegisterType("kego.io/system:@map", ptr8729281952)

	system.RegisterType("kego.io/system:@number", ptr8729282384)

	system.RegisterType("kego.io/system:@object", ptr8729280512)

	system.RegisterType("kego.io/system:@property", ptr8729281232)

	system.RegisterType("kego.io/system:@reference", ptr8729284256)

	system.RegisterType("kego.io/system:@rule", ptr8729283392)

	system.RegisterType("kego.io/system:@string", ptr8729284688)

	system.RegisterType("kego.io/system:@type", ptr8729284976)

	system.RegisterType("kego.io/system:array", ptr8729283968)

	system.RegisterType("kego.io/system:bool", ptr8729284400)

	system.RegisterType("kego.io/system:context", ptr8729280656)

	system.RegisterType("kego.io/system:int", ptr8729281376)

	system.RegisterType("kego.io/system:map", ptr8729281808)

	system.RegisterType("kego.io/system:number", ptr8729282240)

	system.RegisterType("kego.io/system:object", ptr8729282960)

	system.RegisterType("kego.io/system:property", ptr8729281088)

	system.RegisterType("kego.io/system:reference", ptr8729282096)

	system.RegisterType("kego.io/system:rule", ptr8729283248)

	system.RegisterType("kego.io/system:string", ptr8729283824)

	system.RegisterType("kego.io/system:type", ptr8729282816)

}
