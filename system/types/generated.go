package types

import "kego.io/system"

func init() {

	ptr8728352512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728672160 := &system.Object{Context: ptr8728352512, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728694000 := &system.Type{Object: ptr8728672160, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728653568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728708752 := &system.Object{Context: ptr8728653568, Description: "Restriction rules for maps", Id: "@map", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728652736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728708864 := &system.Object{Context: ptr8728652736, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728652704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728709088 := &system.Object{Context: ptr8728652704, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728437312 := &system.Rule_rule{Object: ptr8728709088}

	ptr8728504048 := &system.Property{Object: ptr8728708864, Defaulter: false, Item: ptr8728437312, Optional: false}

	ptr8728653088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728710544 := &system.Object{Context: ptr8728653088, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728653056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728710656 := &system.Object{Context: ptr8728653056, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728446928 := &system.Int_rule{Object: ptr8728710656, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728504144 := &system.Property{Object: ptr8728710544, Defaulter: false, Item: ptr8728446928, Optional: true}

	ptr8728653440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728710768 := &system.Object{Context: ptr8728653440, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728653408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728710880 := &system.Object{Context: ptr8728653408, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728447008 := &system.Int_rule{Object: ptr8728710880, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728504240 := &system.Property{Object: ptr8728710768, Defaulter: false, Item: ptr8728447008, Optional: true}

	ptr8728691696 := &system.Type{Object: ptr8728708752, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"items": ptr8728504048, "minItems": ptr8728504144, "maxItems": ptr8728504240}, Rule: (*system.Type)(nil)}

	ptr8728353248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728707072 := &system.Object{Context: ptr8728353248, Description: "This is the native json array data type", Id: "array", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728352928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728707184 := &system.Object{Context: ptr8728352928, Description: "Restriction rules for arrays", Id: "@array", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728351616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728707520 := &system.Object{Context: ptr8728351616, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728351584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728707856 := &system.Object{Context: ptr8728351584, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728437136 := &system.Rule_rule{Object: ptr8728707856}

	ptr8728503760 := &system.Property{Object: ptr8728707520, Defaulter: false, Item: ptr8728437136, Optional: false}

	ptr8728352032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728707968 := &system.Object{Context: ptr8728352032, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728352000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728708304 := &system.Object{Context: ptr8728352000, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728446528 := &system.Int_rule{Object: ptr8728708304, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728504336 := &system.Property{Object: ptr8728707968, Defaulter: false, Item: ptr8728446528, Optional: true}

	ptr8728352384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728708416 := &system.Object{Context: ptr8728352384, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728352352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728708528 := &system.Object{Context: ptr8728352352, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728446688 := &system.Int_rule{Object: ptr8728708528, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728504720 := &system.Property{Object: ptr8728708416, Defaulter: false, Item: ptr8728446688, Optional: true}

	ptr8728352800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728708640 := &system.Object{Context: ptr8728352800, Description: "If this is true, each item must be unique", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728352768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728708976 := &system.Object{Context: ptr8728352768, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728268704 := &system.Bool_rule{Object: ptr8728708976, Default: system.Bool{Value: false, Exists: true}}

	ptr8728504912 := &system.Property{Object: ptr8728708640, Defaulter: false, Item: ptr8728268704, Optional: true}

	ptr8728690688 := &system.Type{Object: ptr8728707184, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"items": ptr8728503760, "minItems": ptr8728504336, "maxItems": ptr8728504720, "uniqueItems": ptr8728504912}, Rule: (*system.Type)(nil)}

	ptr8728698464 := &system.Type{Object: ptr8728707072, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728690688}

	ptr8728653472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728670368 := &system.Object{Context: ptr8728653472, Description: "This is the most basic type.", Id: "type", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728652960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728666448 := &system.Object{Context: ptr8728652960, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728652928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728666784 := &system.Object{Context: ptr8728652928, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@type", Package: "kego.io/system", Type: "@type", Exists: true}}

	ptr8728436784 := &system.Type_rule{Object: ptr8728666784}

	ptr8728502416 := &system.Property{Object: ptr8728666448, Defaulter: false, Item: ptr8728436784, Optional: true}

	ptr8728652288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728670480 := &system.Object{Context: ptr8728652288, Description: "Type which this should extend", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728652224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728670592 := &system.Object{Context: ptr8728652224, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728626496 := &system.Reference_rule{Object: ptr8728670592, Default: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}}

	ptr8728507792 := &system.Property{Object: ptr8728670480, Defaulter: false, Item: ptr8728626496, Optional: true}

	ptr8728653504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728671264 := &system.Object{Context: ptr8728653504, Description: "Array of interface types that this type should support", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728653312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728672384 := &system.Object{Context: ptr8728653312, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728653280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728672496 := &system.Object{Context: ptr8728653280, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728626752 := &system.Reference_rule{Object: ptr8728672496, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728626624 := &system.Array_rule{Object: ptr8728672384, Items: ptr8728626752, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728507840 := &system.Property{Object: ptr8728671264, Defaulter: false, Item: ptr8728626624, Optional: true}

	ptr8728655008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728672608 := &system.Object{Context: ptr8728655008, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728654816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728672720 := &system.Object{Context: ptr8728654816, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728602656 := &system.String_rule{Object: ptr8728672720, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728508080 := &system.Property{Object: ptr8728672608, Defaulter: false, Item: ptr8728602656, Optional: true}

	ptr8728655936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728672832 := &system.Object{Context: ptr8728655936, Description: "Is this type an interface?", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728655840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728672944 := &system.Object{Context: ptr8728655840, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728438144 := &system.JsonBool_rule{Object: ptr8728672944}

	ptr8728508128 := &system.Property{Object: ptr8728672832, Defaulter: false, Item: ptr8728438144, Optional: true}

	ptr8728656960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728673056 := &system.Object{Context: ptr8728656960, Description: "Exclude this type from the generated json?", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728656928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728673168 := &system.Object{Context: ptr8728656928, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728438168 := &system.JsonBool_rule{Object: ptr8728673168}

	ptr8728508320 := &system.Property{Object: ptr8728673056, Defaulter: false, Item: ptr8728438168, Optional: true}

	ptr8728651360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728673280 := &system.Object{Context: ptr8728651360, Description: "Each field is listed with it's type", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728651104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728673392 := &system.Object{Context: ptr8728651104, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728650944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728673728 := &system.Object{Context: ptr8728650944, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@property", Package: "kego.io/system", Type: "@property", Exists: true}}

	ptr8728438200 := &system.Property_rule{Object: ptr8728673728}

	ptr8728626880 := &system.Map_rule{Object: ptr8728673392, Items: ptr8728438200, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728508416 := &system.Property{Object: ptr8728673280, Defaulter: false, Item: ptr8728626880, Optional: true}

	ptr8728693856 := &system.Type{Object: ptr8728670368, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"rule": ptr8728502416, "extends": ptr8728507792, "is": ptr8728507840, "native": ptr8728508080, "interface": ptr8728508128, "exclude": ptr8728508320, "properties": ptr8728508416}, Rule: (*system.Type)(nil)}

	ptr8728351200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728671824 := &system.Object{Context: ptr8728351200, Description: "Restriction rules for references", Id: "@reference", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728351040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728671936 := &system.Object{Context: ptr8728351040, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728351008 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728672048 := &system.Object{Context: ptr8728351008, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728627136 := &system.Reference_rule{Object: ptr8728672048, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728507936 := &system.Property{Object: ptr8728671936, Defaulter: true, Item: ptr8728627136, Optional: true}

	ptr8728693712 := &system.Type{Object: ptr8728671824, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8728507936}, Rule: (*system.Type)(nil)}

	ptr8728650304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728709872 := &system.Object{Context: ptr8728650304, Description: "Automatically created basic rule for context", Id: "@context", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728691552 := &system.Type{Object: ptr8728709872, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728650688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728711104 := &system.Object{Context: ptr8728650688, Description: "Restriction rules for numbers", Id: "@number", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728655296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728711216 := &system.Object{Context: ptr8728655296, Description: "Default value if this property is omitted", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728655136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728711552 := &system.Object{Context: ptr8728655136, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728447168 := &system.Number_rule{Object: ptr8728711552, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728504576 := &system.Property{Object: ptr8728711216, Defaulter: true, Item: ptr8728447168, Optional: true}

	ptr8728655744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728711664 := &system.Object{Context: ptr8728655744, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728655584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728711776 := &system.Object{Context: ptr8728655584, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728447248 := &system.Number_rule{Object: ptr8728711776, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728505584 := &system.Property{Object: ptr8728711664, Defaulter: false, Item: ptr8728447248, Optional: true}

	ptr8728656192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728711888 := &system.Object{Context: ptr8728656192, Description: "This provides a lower bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728656032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728712000 := &system.Object{Context: ptr8728656032, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728452128 := &system.Number_rule{Object: ptr8728712000, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728505872 := &system.Property{Object: ptr8728711888, Defaulter: false, Item: ptr8728452128, Optional: true}

	ptr8728656544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728712112 := &system.Object{Context: ptr8728656544, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728656512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728712224 := &system.Object{Context: ptr8728656512, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728271760 := &system.Bool_rule{Object: ptr8728712224, Default: system.Bool{Value: false, Exists: true}}

	ptr8728506496 := &system.Property{Object: ptr8728712112, Defaulter: false, Item: ptr8728271760, Optional: true}

	ptr8728656992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728712336 := &system.Object{Context: ptr8728656992, Description: "This provides an upper bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728656832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728712448 := &system.Object{Context: ptr8728656832, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8728447488 := &system.Number_rule{Object: ptr8728712448, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8728506688 := &system.Property{Object: ptr8728712336, Defaulter: false, Item: ptr8728447488, Optional: true}

	ptr8728650432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728712560 := &system.Object{Context: ptr8728650432, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728650400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728712672 := &system.Object{Context: ptr8728650400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728272176 := &system.Bool_rule{Object: ptr8728712672, Default: system.Bool{Value: false, Exists: true}}

	ptr8728506832 := &system.Property{Object: ptr8728712560, Defaulter: false, Item: ptr8728272176, Optional: true}

	ptr8728692416 := &system.Type{Object: ptr8728711104, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8728504576, "multipleOf": ptr8728505584, "minimum": ptr8728505872, "exclusiveMinimum": ptr8728506496, "maximum": ptr8728506688, "exclusiveMaximum": ptr8728506832}, Rule: (*system.Type)(nil)}

	ptr8728347328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728714912 := &system.Object{Context: ptr8728347328, Description: "This is a property of a type", Id: "property", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728657600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728670816 := &system.Object{Context: ptr8728657600, Description: "This specifies that the field is the default value for a rule", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728657568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728670928 := &system.Object{Context: ptr8728657568, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728437864 := &system.JsonBool_rule{Object: ptr8728670928}

	ptr8728507600 := &system.Property{Object: ptr8728670816, Defaulter: false, Item: ptr8728437864, Optional: true}

	ptr8728347200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728671040 := &system.Object{Context: ptr8728347200, Description: "This is a rule object, defining the type and restrictions on the value of the this property", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728347168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728671376 := &system.Object{Context: ptr8728347168, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728437888 := &system.Rule_rule{Object: ptr8728671376}

	ptr8728507648 := &system.Property{Object: ptr8728671040, Defaulter: false, Item: ptr8728437888, Optional: false}

	ptr8728657184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728715024 := &system.Object{Context: ptr8728657184, Description: "This specifies that the field is optional", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728657120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728670704 := &system.Object{Context: ptr8728657120, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8728437792 := &system.JsonBool_rule{Object: ptr8728670704}

	ptr8728507264 := &system.Property{Object: ptr8728715024, Defaulter: false, Item: ptr8728437792, Optional: true}

	ptr8728693136 := &system.Type{Object: ptr8728714912, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"defaulter": ptr8728507600, "item": ptr8728507648, "optional": ptr8728507264}, Rule: (*system.Type)(nil)}

	ptr8728651456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728710096 := &system.Object{Context: ptr8728651456, Description: "Restriction rules for integers", Id: "@int", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728651072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728710208 := &system.Object{Context: ptr8728651072, Description: "Default value if this property is omitted", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728651040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728710320 := &system.Object{Context: ptr8728651040, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728447408 := &system.Int_rule{Object: ptr8728710320, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728506208 := &system.Property{Object: ptr8728710208, Defaulter: true, Item: ptr8728447408, Optional: true}

	ptr8728650176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728710432 := &system.Object{Context: ptr8728650176, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728650144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728707296 := &system.Object{Context: ptr8728650144, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728445408 := &system.Int_rule{Object: ptr8728707296, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728506352 := &system.Property{Object: ptr8728710432, Defaulter: false, Item: ptr8728445408, Optional: true}

	ptr8728650880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728707408 := &system.Object{Context: ptr8728650880, Description: "This provides a lower bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728650816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728707632 := &system.Object{Context: ptr8728650816, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728445488 := &system.Int_rule{Object: ptr8728707632, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728502560 := &system.Property{Object: ptr8728707408, Defaulter: false, Item: ptr8728445488, Optional: true}

	ptr8728651328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728707744 := &system.Object{Context: ptr8728651328, Description: "This provides an upper bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728651296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728708080 := &system.Object{Context: ptr8728651296, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728446768 := &system.Int_rule{Object: ptr8728708080, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728502704 := &system.Property{Object: ptr8728707744, Defaulter: false, Item: ptr8728446768, Optional: true}

	ptr8728692272 := &system.Type{Object: ptr8728710096, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8728506208, "multipleOf": ptr8728506352, "minimum": ptr8728502560, "maximum": ptr8728502704}, Rule: (*system.Type)(nil)}

	ptr8728651776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728709984 := &system.Object{Context: ptr8728651776, Description: "This is the int data type", Id: "int", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728692128 := &system.Type{Object: ptr8728709984, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728692272}

	ptr8728354080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728666224 := &system.Object{Context: ptr8728354080, Description: "Restriction rules for strings", Id: "@string", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728352608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728669920 := &system.Object{Context: ptr8728352608, Description: "This is a regex to match the value to", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728352576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728670032 := &system.Object{Context: ptr8728352576, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728602336 := &system.String_rule{Object: ptr8728670032, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728506976 := &system.Property{Object: ptr8728669920, Defaulter: false, Item: ptr8728602336, Optional: true}

	ptr8728353888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728670144 := &system.Object{Context: ptr8728353888, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728353824 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728670256 := &system.Object{Context: ptr8728353824, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728602496 := &system.String_rule{Object: ptr8728670256, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728507360 := &system.Property{Object: ptr8728670144, Defaulter: false, Item: ptr8728602496, Optional: true}

	ptr8728348672 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728666336 := &system.Object{Context: ptr8728348672, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728348640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728668576 := &system.Object{Context: ptr8728348640, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728601536 := &system.String_rule{Object: ptr8728668576, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728502512 := &system.Property{Object: ptr8728666336, Defaulter: true, Item: ptr8728601536, Optional: true}

	ptr8728349856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728668688 := &system.Object{Context: ptr8728349856, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728349760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728669024 := &system.Object{Context: ptr8728349760, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728349728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728669136 := &system.Object{Context: ptr8728349728, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728602016 := &system.String_rule{Object: ptr8728669136, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728625472 := &system.Array_rule{Object: ptr8728669024, Items: ptr8728602016, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8728504384 := &system.Property{Object: ptr8728668688, Defaulter: false, Item: ptr8728625472, Optional: true}

	ptr8728350176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728669248 := &system.Object{Context: ptr8728350176, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728350144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728669360 := &system.Object{Context: ptr8728350144, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728445328 := &system.Int_rule{Object: ptr8728669360, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728506304 := &system.Property{Object: ptr8728669248, Defaulter: false, Item: ptr8728445328, Optional: true}

	ptr8728351072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728669472 := &system.Object{Context: ptr8728351072, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728350848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728669584 := &system.Object{Context: ptr8728350848, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728445648 := &system.Int_rule{Object: ptr8728669584, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8728506592 := &system.Property{Object: ptr8728669472, Defaulter: false, Item: ptr8728445648, Optional: true}

	ptr8728351840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728669696 := &system.Object{Context: ptr8728351840, Description: "This is a string that the value must match", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728351744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728669808 := &system.Object{Context: ptr8728351744, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8728602176 := &system.String_rule{Object: ptr8728669808, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8728506784 := &system.Property{Object: ptr8728669696, Defaulter: false, Item: ptr8728602176, Optional: true}

	ptr8728691840 := &system.Type{Object: ptr8728666224, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"pattern": ptr8728506976, "format": ptr8728507360, "default": ptr8728502512, "enum": ptr8728504384, "minLength": ptr8728506304, "maxLength": ptr8728506592, "equal": ptr8728506784}, Rule: (*system.Type)(nil)}

	ptr8728353120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728672272 := &system.Object{Context: ptr8728353120, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728694144 := &system.Type{Object: ptr8728672272, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728653888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728708192 := &system.Object{Context: ptr8728653888, Description: "This is the native json object data type.", Id: "map", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728691264 := &system.Type{Object: ptr8728708192, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728691696}

	ptr8728650112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728709760 := &system.Object{Context: ptr8728650112, Description: "Unmarshal context.", Id: "context", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728691408 := &system.Type{Object: ptr8728709760, Exclude: true, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728354336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728709312 := &system.Object{Context: ptr8728354336, Description: "Restriction rules for bools", Id: "@bool", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728354208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728709424 := &system.Object{Context: ptr8728354208, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728354176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728709648 := &system.Object{Context: ptr8728354176, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728269744 := &system.Bool_rule{Object: ptr8728709648, Default: system.Bool{Value: false, Exists: false}}

	ptr8728505728 := &system.Property{Object: ptr8728709424, Defaulter: true, Item: ptr8728269744, Optional: true}

	ptr8728691120 := &system.Type{Object: ptr8728709312, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8728505728}, Rule: (*system.Type)(nil)}

	ptr8728654656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728666896 := &system.Object{Context: ptr8728654656, Description: "Automatically created basic rule for type", Id: "@type", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728694576 := &system.Type{Object: ptr8728666896, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728346656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728671488 := &system.Object{Context: ptr8728346656, Description: "Automatically created basic rule for property", Id: "@property", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728693280 := &system.Type{Object: ptr8728671488, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728351808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728671600 := &system.Object{Context: ptr8728351808, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728693568 := &system.Type{Object: ptr8728671600, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728693712}

	ptr8728354656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728709200 := &system.Object{Context: ptr8728354656, Description: "This is the native json bool data type", Id: "bool", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728690976 := &system.Type{Object: ptr8728709200, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728691120}

	ptr8728651424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728710992 := &system.Object{Context: ptr8728651424, Description: "This is the native json number data type", Id: "number", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728698320 := &system.Type{Object: ptr8728710992, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728692416}

	ptr8728354592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728666112 := &system.Object{Context: ptr8728354592, Description: "This is the native json string data type", Id: "string", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728690832 := &system.Type{Object: ptr8728666112, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8728691840}

	ptr8728655872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728711328 := &system.Object{Context: ptr8728655872, Description: "This is the most basic type.", Id: "object", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728654336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728713568 := &system.Object{Context: ptr8728654336, Description: "Description for the developer", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728654304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728713680 := &system.Object{Context: ptr8728654304, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728437464 := &system.JsonString_rule{Object: ptr8728713680}

	ptr8728505296 := &system.Property{Object: ptr8728713568, Defaulter: false, Item: ptr8728437464, Optional: true}

	ptr8728654944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728713792 := &system.Object{Context: ptr8728654944, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728654912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728714128 := &system.Object{Context: ptr8728654912, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@context", Package: "kego.io/system", Type: "@context", Exists: true}}

	ptr8728437496 := &system.Context_rule{Object: ptr8728714128}

	ptr8728505344 := &system.Property{Object: ptr8728713792, Defaulter: false, Item: ptr8728437496, Optional: true}

	ptr8728655712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728714240 := &system.Object{Context: ptr8728655712, Description: "Extra validation rules for this object or descendants", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728655680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728714576 := &system.Object{Context: ptr8728655680, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728655648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728714688 := &system.Object{Context: ptr8728655648, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728437736 := &system.Rule_rule{Object: ptr8728714688}

	ptr8728626048 := &system.Map_rule{Object: ptr8728714576, Items: ptr8728437736, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8728506400 := &system.Property{Object: ptr8728714240, Defaulter: false, Item: ptr8728626048, Optional: true}

	ptr8728652864 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728711440 := &system.Object{Context: ptr8728652864, Description: "Type of the object.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728652832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728713008 := &system.Object{Context: ptr8728652832, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8728625728 := &system.Reference_rule{Object: ptr8728713008, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8728503856 := &system.Property{Object: ptr8728711440, Defaulter: false, Item: ptr8728625728, Optional: false}

	ptr8728653632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728713120 := &system.Object{Context: ptr8728653632, Description: "All global objects should have an id.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728653600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728713456 := &system.Object{Context: ptr8728653600, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8728437336 := &system.JsonString_rule{Object: ptr8728713456}

	ptr8728504768 := &system.Property{Object: ptr8728713120, Defaulter: false, Item: ptr8728437336, Optional: true}

	ptr8728691984 := &system.Type{Object: ptr8728711328, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"description": ptr8728505296, "context": ptr8728505344, "rules": ptr8728506400, "type": ptr8728503856, "id": ptr8728504768}, Rule: (*system.Type)(nil)}

	ptr8728656160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728714800 := &system.Object{Context: ptr8728656160, Description: "Automatically created basic rule for object", Id: "@object", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728692848 := &system.Type{Object: ptr8728714800, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system:@array", ptr8728690688)

	system.RegisterType("kego.io/system:@bool", ptr8728691120)

	system.RegisterType("kego.io/system:@context", ptr8728691552)

	system.RegisterType("kego.io/system:@int", ptr8728692272)

	system.RegisterType("kego.io/system:@map", ptr8728691696)

	system.RegisterType("kego.io/system:@number", ptr8728692416)

	system.RegisterType("kego.io/system:@object", ptr8728692848)

	system.RegisterType("kego.io/system:@property", ptr8728693280)

	system.RegisterType("kego.io/system:@reference", ptr8728693712)

	system.RegisterType("kego.io/system:@rule", ptr8728694144)

	system.RegisterType("kego.io/system:@string", ptr8728691840)

	system.RegisterType("kego.io/system:@type", ptr8728694576)

	system.RegisterType("kego.io/system:array", ptr8728698464)

	system.RegisterType("kego.io/system:bool", ptr8728690976)

	system.RegisterType("kego.io/system:context", ptr8728691408)

	system.RegisterType("kego.io/system:int", ptr8728692128)

	system.RegisterType("kego.io/system:map", ptr8728691264)

	system.RegisterType("kego.io/system:number", ptr8728698320)

	system.RegisterType("kego.io/system:object", ptr8728691984)

	system.RegisterType("kego.io/system:property", ptr8728693136)

	system.RegisterType("kego.io/system:reference", ptr8728693568)

	system.RegisterType("kego.io/system:rule", ptr8728694000)

	system.RegisterType("kego.io/system:string", ptr8728690832)

	system.RegisterType("kego.io/system:type", ptr8728693856)

}
