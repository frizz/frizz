package types

import (
	"kego.io/system"

	_ "kego.io/system/types"

	"kego.io/jsonselect"
)

func init() {

	ptr8729583232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729763952 := &system.Object{Context: ptr8729583232, Description: "", Id: "collision", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729583104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729764064 := &system.Object{Context: ptr8729583104, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729583072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729764176 := &system.Object{Context: ptr8729583072, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729583040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729764288 := &system.Object{Context: ptr8729583040, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729502176 := &system.String_rule{Object: ptr8729764288, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729673984 := &system.Map_rule{Object: ptr8729764176, Items: ptr8729502176, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729496144 := &system.Property{Object: ptr8729764064, Defaulter: false, Item: ptr8729673984, Optional: false}

	ptr8729510752 := &system.Type{Object: ptr8729763952, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"number": ptr8729496144}, Rule: (*system.Type)(nil)}

	ptr8729754688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729768544 := &system.Object{Context: ptr8729754688, Description: "Automatically created basic rule for basic", Id: "@basic", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729514640 := &system.Type{Object: ptr8729768544, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729576896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729768656 := &system.Object{Context: ptr8729576896, Description: "", Id: "c", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729575968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729768992 := &system.Object{Context: ptr8729575968, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729575808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729769104 := &system.Object{Context: ptr8729575808, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729817632 := &system.Number_rule{Object: ptr8729769104, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729499600 := &system.Property{Object: ptr8729768992, Defaulter: false, Item: ptr8729817632, Optional: false}

	ptr8729576736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729769216 := &system.Object{Context: ptr8729576736, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729576704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729769328 := &system.Object{Context: ptr8729576704, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729576544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729769440 := &system.Object{Context: ptr8729576544, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729817712 := &system.Number_rule{Object: ptr8729769440, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729675136 := &system.Map_rule{Object: ptr8729769328, Items: ptr8729817712, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729499648 := &system.Property{Object: ptr8729769216, Defaulter: false, Item: ptr8729675136, Optional: false}

	ptr8729755488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729768768 := &system.Object{Context: ptr8729755488, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729755328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729768880 := &system.Object{Context: ptr8729755328, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729817552 := &system.Number_rule{Object: ptr8729768880, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729499552 := &system.Property{Object: ptr8729768768, Defaulter: false, Item: ptr8729817552, Optional: false}

	ptr8729514928 := &system.Type{Object: ptr8729768656, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"b": ptr8729499600, "c": ptr8729499648, "a": ptr8729499552}, Rule: (*system.Type)(nil)}

	ptr8729576992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729764400 := &system.Object{Context: ptr8729576992, Description: "Automatically created basic rule for collision", Id: "@collision", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729510896 := &system.Type{Object: ptr8729764400, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729805216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729371520 := &system.Object{Context: ptr8729805216, Description: "", Id: "polykids", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729805088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729371632 := &system.Object{Context: ptr8729805088, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729804992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729371744 := &system.Object{Context: ptr8729804992, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729804960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729372080 := &system.Object{Context: ptr8729804960, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@kid", Package: "kego.io/jsonselect", Type: "@kid", Exists: true}}

	ptr8729241776 := &jsonselect.Kid_rule{Object: ptr8729372080}

	ptr8729675392 := &system.Array_rule{Object: ptr8729371744, Items: ptr8729241776, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729500032 := &system.Property{Object: ptr8729371632, Defaulter: false, Item: ptr8729675392, Optional: false}

	ptr8729513776 := &system.Type{Object: ptr8729371520, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8729500032}, Rule: (*system.Type)(nil)}

	ptr8729754272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729765408 := &system.Object{Context: ptr8729754272, Description: "", Id: "basic", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729753248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729767760 := &system.Object{Context: ptr8729753248, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729753152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729767872 := &system.Object{Context: ptr8729753152, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729753120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729767984 := &system.Object{Context: ptr8729753120, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729503776 := &system.String_rule{Object: ptr8729767984, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729674816 := &system.Array_rule{Object: ptr8729767872, Items: ptr8729503776, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729498736 := &system.Property{Object: ptr8729767760, Defaulter: false, Item: ptr8729674816, Optional: false}

	ptr8729754144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729768096 := &system.Object{Context: ptr8729754144, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729753984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729768432 := &system.Object{Context: ptr8729753984, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729817072 := &system.Number_rule{Object: ptr8729768432, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729498832 := &system.Property{Object: ptr8729768096, Defaulter: false, Item: ptr8729817072, Optional: false}

	ptr8729750272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729765744 := &system.Object{Context: ptr8729750272, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729750240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729766080 := &system.Object{Context: ptr8729750240, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729750208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729766416 := &system.Object{Context: ptr8729750208, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729502816 := &system.String_rule{Object: ptr8729766416, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729674368 := &system.Map_rule{Object: ptr8729766080, Items: ptr8729502816, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729496864 := &system.Property{Object: ptr8729765744, Defaulter: false, Item: ptr8729674368, Optional: false}

	ptr8729750720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729766528 := &system.Object{Context: ptr8729750720, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729750688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729766640 := &system.Object{Context: ptr8729750688, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729503136 := &system.String_rule{Object: ptr8729766640, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729497680 := &system.Property{Object: ptr8729766528, Defaulter: false, Item: ptr8729503136, Optional: false}

	ptr8729751808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729766752 := &system.Object{Context: ptr8729751808, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729751712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729767088 := &system.Object{Context: ptr8729751712, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729751680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729767200 := &system.Object{Context: ptr8729751680, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729751648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729767312 := &system.Object{Context: ptr8729751648, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729503456 := &system.String_rule{Object: ptr8729767312, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729674688 := &system.Map_rule{Object: ptr8729767200, Items: ptr8729503456, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729674496 := &system.Array_rule{Object: ptr8729767088, Items: ptr8729674688, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729497728 := &system.Property{Object: ptr8729766752, Defaulter: false, Item: ptr8729674496, Optional: false}

	ptr8729752512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729767424 := &system.Object{Context: ptr8729752512, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729752416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729767536 := &system.Object{Context: ptr8729752416, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729752384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729767648 := &system.Object{Context: ptr8729752384, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729503616 := &system.String_rule{Object: ptr8729767648, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729674752 := &system.Array_rule{Object: ptr8729767536, Items: ptr8729503616, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729498304 := &system.Property{Object: ptr8729767424, Defaulter: false, Item: ptr8729674752, Optional: false}

	ptr8729514496 := &system.Type{Object: ptr8729765408, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"drinkPreference": ptr8729498736, "weight": ptr8729498832, "name": ptr8729496864, "favoriteColor": ptr8729497680, "languagesSpoken": ptr8729497728, "seatingPreference": ptr8729498304}, Rule: (*system.Type)(nil)}

	ptr8729812288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729375104 := &system.Object{Context: ptr8729812288, Description: "", Id: "typed", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729809472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729375216 := &system.Object{Context: ptr8729809472, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729809440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729375328 := &system.Object{Context: ptr8729809440, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729809408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729375440 := &system.Object{Context: ptr8729809408, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729501856 := &system.String_rule{Object: ptr8729375440, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729675456 := &system.Map_rule{Object: ptr8729375328, Items: ptr8729501856, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729494320 := &system.Property{Object: ptr8729375216, Defaulter: false, Item: ptr8729675456, Optional: false}

	ptr8729809920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729375552 := &system.Object{Context: ptr8729809920, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729809888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729375664 := &system.Object{Context: ptr8729809888, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729502016 := &system.String_rule{Object: ptr8729375664, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729494416 := &system.Property{Object: ptr8729375552, Defaulter: false, Item: ptr8729502016, Optional: false}

	ptr8729810496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729375776 := &system.Object{Context: ptr8729810496, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729810464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729375888 := &system.Object{Context: ptr8729810464, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729810432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729376000 := &system.Object{Context: ptr8729810432, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@kid", Package: "kego.io/jsonselect", Type: "@kid", Exists: true}}

	ptr8729241160 := &jsonselect.Kid_rule{Object: ptr8729376000}

	ptr8729675712 := &system.Map_rule{Object: ptr8729375888, Items: ptr8729241160, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729494464 := &system.Property{Object: ptr8729375776, Defaulter: false, Item: ptr8729675712, Optional: false}

	ptr8729810880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729376112 := &system.Object{Context: ptr8729810880, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729810848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729376224 := &system.Object{Context: ptr8729810848, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@image", Package: "kego.io/jsonselect", Type: "@image", Exists: true}}

	ptr8729241184 := &jsonselect.Image_rule{Object: ptr8729376224}

	ptr8729494560 := &system.Property{Object: ptr8729376112, Defaulter: false, Item: ptr8729241184, Optional: false}

	ptr8729811584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729376336 := &system.Object{Context: ptr8729811584, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729811488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729376448 := &system.Object{Context: ptr8729811488, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729811456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729376560 := &system.Object{Context: ptr8729811456, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729503296 := &system.String_rule{Object: ptr8729376560, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729675840 := &system.Array_rule{Object: ptr8729376448, Items: ptr8729503296, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729494608 := &system.Property{Object: ptr8729376336, Defaulter: false, Item: ptr8729675840, Optional: false}

	ptr8729812160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729377008 := &system.Object{Context: ptr8729812160, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729812000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729377120 := &system.Object{Context: ptr8729812000, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729813632 := &system.Number_rule{Object: ptr8729377120, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729494704 := &system.Property{Object: ptr8729377008, Defaulter: false, Item: ptr8729813632, Optional: false}

	ptr8729510176 := &system.Type{Object: ptr8729375104, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"name": ptr8729494320, "favoriteColor": ptr8729494416, "kids": ptr8729494464, "avatar": ptr8729494560, "drinkPreference": ptr8729494608, "weight": ptr8729494704}, Rule: (*system.Type)(nil)}

	ptr8729582144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729764960 := &system.Object{Context: ptr8729582144, Description: "", Id: "expr", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729581056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729766304 := &system.Object{Context: ptr8729581056, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729581024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729766864 := &system.Object{Context: ptr8729581024, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729502976 := &system.String_rule{Object: ptr8729766864, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729497104 := &system.Property{Object: ptr8729766304, Defaulter: false, Item: ptr8729502976, Optional: false}

	ptr8729581568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729766976 := &system.Object{Context: ptr8729581568, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729581536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729769552 := &system.Object{Context: ptr8729581536, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729761968 := &system.Bool_rule{Object: ptr8729769552, Default: system.Bool{Value: false, Exists: false}}

	ptr8729497152 := &system.Property{Object: ptr8729766976, Defaulter: false, Item: ptr8729761968, Optional: false}

	ptr8729582016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729769664 := &system.Object{Context: ptr8729582016, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729581984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729769888 := &system.Object{Context: ptr8729581984, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729762512 := &system.Bool_rule{Object: ptr8729769888, Default: system.Bool{Value: false, Exists: false}}

	ptr8729497440 := &system.Property{Object: ptr8729769664, Defaulter: false, Item: ptr8729762512, Optional: false}

	ptr8729579136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729765072 := &system.Object{Context: ptr8729579136, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729578944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729765184 := &system.Object{Context: ptr8729578944, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729814192 := &system.Number_rule{Object: ptr8729765184, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729496912 := &system.Property{Object: ptr8729765072, Defaulter: false, Item: ptr8729814192, Optional: false}

	ptr8729579712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729765296 := &system.Object{Context: ptr8729579712, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729579552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729765520 := &system.Object{Context: ptr8729579552, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729814272 := &system.Number_rule{Object: ptr8729765520, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729496960 := &system.Property{Object: ptr8729765296, Defaulter: false, Item: ptr8729814272, Optional: false}

	ptr8729580160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729765632 := &system.Object{Context: ptr8729580160, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729580128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729765856 := &system.Object{Context: ptr8729580128, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729502496 := &system.String_rule{Object: ptr8729765856, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729497008 := &system.Property{Object: ptr8729765632, Defaulter: false, Item: ptr8729502496, Optional: false}

	ptr8729580608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729765968 := &system.Object{Context: ptr8729580608, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729580576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729766192 := &system.Object{Context: ptr8729580576, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729502656 := &system.String_rule{Object: ptr8729766192, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729497056 := &system.Property{Object: ptr8729765968, Defaulter: false, Item: ptr8729502656, Optional: false}

	ptr8729511616 := &system.Type{Object: ptr8729764960, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"null": ptr8729497104, "true": ptr8729497152, "false": ptr8729497440, "int": ptr8729496912, "float": ptr8729496960, "string": ptr8729497008, "string2": ptr8729497056}, Rule: (*system.Type)(nil)}

	ptr8729748000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729770784 := &system.Object{Context: ptr8729748000, Description: "Automatically created basic rule for gallery", Id: "@gallery", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729512192 := &system.Type{Object: ptr8729770784, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729751392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729771904 := &system.Object{Context: ptr8729751392, Description: "Automatically created basic rule for kid", Id: "@kid", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729513056 := &system.Type{Object: ptr8729771904, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729748704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729771008 := &system.Object{Context: ptr8729748704, Description: "Automatically created basic rule for image", Id: "@image", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729512624 := &system.Type{Object: ptr8729771008, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729577888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729764512 := &system.Object{Context: ptr8729577888, Description: "", Id: "diagram", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729577760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729764624 := &system.Object{Context: ptr8729577760, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729577728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729764736 := &system.Object{Context: ptr8729577728, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729502336 := &system.String_rule{Object: ptr8729764736, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729496480 := &system.Property{Object: ptr8729764624, Defaulter: false, Item: ptr8729502336, Optional: false}

	ptr8729511184 := &system.Type{Object: ptr8729764512, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/jsonselect:image", Package: "kego.io/jsonselect", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"url": ptr8729496480}, Rule: (*system.Type)(nil)}

	ptr8729805632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729372192 := &system.Object{Context: ptr8729805632, Description: "Automatically created basic rule for polykids", Id: "@polykids", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729513920 := &system.Type{Object: ptr8729372192, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729747584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729770112 := &system.Object{Context: ptr8729747584, Description: "This represents a gallery - it's just a list of images", Id: "gallery", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729747456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729770224 := &system.Object{Context: ptr8729747456, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729156352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729770336 := &system.Object{Context: ptr8729156352, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729156320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729770672 := &system.Object{Context: ptr8729156320, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@image", Package: "kego.io/jsonselect", Type: "@image", Exists: true}}

	ptr8729241416 := &jsonselect.Image_rule{Object: ptr8729770672}

	ptr8729674624 := &system.Map_rule{Object: ptr8729770336, Items: ptr8729241416, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729497824 := &system.Property{Object: ptr8729770224, Defaulter: false, Item: ptr8729674624, Optional: false}

	ptr8729512048 := &system.Type{Object: ptr8729770112, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"images": ptr8729497824}, Rule: (*system.Type)(nil)}

	ptr8729582560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729770000 := &system.Object{Context: ptr8729582560, Description: "Automatically created basic rule for expr", Id: "@expr", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729511760 := &system.Type{Object: ptr8729770000, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729808192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729376672 := &system.Object{Context: ptr8729808192, Description: "", Id: "sibling", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729804832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729376784 := &system.Object{Context: ptr8729804832, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729806272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729376896 := &system.Object{Context: ptr8729806272, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729816992 := &system.Number_rule{Object: ptr8729376896, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729500656 := &system.Property{Object: ptr8729376784, Defaulter: false, Item: ptr8729816992, Optional: false}

	ptr8729805888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729371856 := &system.Object{Context: ptr8729805888, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729805504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729371968 := &system.Object{Context: ptr8729805504, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729812992 := &system.Number_rule{Object: ptr8729371968, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729493552 := &system.Property{Object: ptr8729371856, Defaulter: false, Item: ptr8729812992, Optional: false}

	ptr8729806528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729372304 := &system.Object{Context: ptr8729806528, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729806496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729374096 := &system.Object{Context: ptr8729806496, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@c", Package: "kego.io/jsonselect", Type: "@c", Exists: true}}

	ptr8729240152 := &jsonselect.C_rule{Object: ptr8729374096}

	ptr8729493600 := &system.Property{Object: ptr8729372304, Defaulter: false, Item: ptr8729240152, Optional: false}

	ptr8729807296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729374208 := &system.Object{Context: ptr8729807296, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729807264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729374320 := &system.Object{Context: ptr8729807264, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729807104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729374544 := &system.Object{Context: ptr8729807104, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729813072 := &system.Number_rule{Object: ptr8729374544, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729673792 := &system.Map_rule{Object: ptr8729374320, Items: ptr8729813072, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729493936 := &system.Property{Object: ptr8729374208, Defaulter: false, Item: ptr8729673792, Optional: false}

	ptr8729808064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729374656 := &system.Object{Context: ptr8729808064, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729808032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729374768 := &system.Object{Context: ptr8729808032, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729807872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729374880 := &system.Object{Context: ptr8729807872, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729813152 := &system.Number_rule{Object: ptr8729374880, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729674112 := &system.Map_rule{Object: ptr8729374768, Items: ptr8729813152, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729494032 := &system.Property{Object: ptr8729374656, Defaulter: false, Item: ptr8729674112, Optional: false}

	ptr8729514352 := &system.Type{Object: ptr8729376672, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8729500656, "b": ptr8729493552, "c": ptr8729493600, "d": ptr8729493936, "e": ptr8729494032}, Rule: (*system.Type)(nil)}

	ptr8729754848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729371408 := &system.Object{Context: ptr8729754848, Description: "Automatically created basic rule for photo", Id: "@photo", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729513488 := &system.Type{Object: ptr8729371408, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729748288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729770896 := &system.Object{Context: ptr8729748288, Description: "", Id: "image", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729512480 := &system.Type{Object: ptr8729770896, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729754080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729370624 := &system.Object{Context: ptr8729754080, Description: "This represents an image, and contains path, server and protocol separately", Id: "photo", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729753280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729370960 := &system.Object{Context: ptr8729753280, Description: "The server for the url - e.g. www.google.com", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729753216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729371072 := &system.Object{Context: ptr8729753216, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729504576 := &system.String_rule{Object: ptr8729371072, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729499312 := &system.Property{Object: ptr8729370960, Defaulter: false, Item: ptr8729504576, Optional: false}

	ptr8729753888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729371184 := &system.Object{Context: ptr8729753888, Description: "The path for the url - e.g. /foo/bar.jpg", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729753856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729371296 := &system.Object{Context: ptr8729753856, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729504736 := &system.String_rule{Object: ptr8729371296, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "^/.*$", Exists: true}}

	ptr8729499744 := &system.Property{Object: ptr8729371184, Defaulter: false, Item: ptr8729504736, Optional: false}

	ptr8729752640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729370736 := &system.Object{Context: ptr8729752640, Description: "The protocol for the url - e.g. http or https", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729752608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729370848 := &system.Object{Context: ptr8729752608, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729504416 := &system.String_rule{Object: ptr8729370848, Default: system.String{Value: "http", Exists: true}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729499120 := &system.Property{Object: ptr8729370736, Defaulter: false, Item: ptr8729504416, Optional: true}

	ptr8729513344 := &system.Type{Object: ptr8729370624, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/jsonselect:image", Package: "kego.io/jsonselect", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"server": ptr8729499312, "path": ptr8729499744, "protocol": ptr8729499120}, Rule: (*system.Type)(nil)}

	ptr8729576064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729763840 := &system.Object{Context: ptr8729576064, Description: "Automatically created basic rule for c", Id: "@c", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729510320 := &system.Type{Object: ptr8729763840, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729812704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729377232 := &system.Object{Context: ptr8729812704, Description: "Automatically created basic rule for typed", Id: "@typed", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729510608 := &system.Type{Object: ptr8729377232, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729750880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729771120 := &system.Object{Context: ptr8729750880, Description: "", Id: "kid", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729749536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729771232 := &system.Object{Context: ptr8729749536, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729749472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729771344 := &system.Object{Context: ptr8729749472, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729504096 := &system.String_rule{Object: ptr8729771344, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729498592 := &system.Property{Object: ptr8729771232, Defaulter: false, Item: ptr8729504096, Optional: false}

	ptr8729750048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729771456 := &system.Object{Context: ptr8729750048, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729750016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729771568 := &system.Object{Context: ptr8729750016, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729504256 := &system.String_rule{Object: ptr8729771568, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729498640 := &system.Property{Object: ptr8729771456, Defaulter: false, Item: ptr8729504256, Optional: false}

	ptr8729750752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729771680 := &system.Object{Context: ptr8729750752, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729750624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729771792 := &system.Object{Context: ptr8729750624, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729072528 := &system.Bool_rule{Object: ptr8729771792, Default: system.Bool{Value: false, Exists: false}}

	ptr8729498784 := &system.Property{Object: ptr8729771680, Defaulter: false, Item: ptr8729072528, Optional: true}

	ptr8729512912 := &system.Type{Object: ptr8729771120, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"language": ptr8729498592, "level": ptr8729498640, "preferred": ptr8729498784}, Rule: (*system.Type)(nil)}

	ptr8729808608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729374992 := &system.Object{Context: ptr8729808608, Description: "Automatically created basic rule for sibling", Id: "@sibling", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729509888 := &system.Type{Object: ptr8729374992, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729578304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729764848 := &system.Object{Context: ptr8729578304, Description: "Automatically created basic rule for diagram", Id: "@diagram", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729511328 := &system.Type{Object: ptr8729764848, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/jsonselect:@basic", ptr8729514640)

	system.RegisterType("kego.io/jsonselect:@c", ptr8729510320)

	system.RegisterType("kego.io/jsonselect:@collision", ptr8729510896)

	system.RegisterType("kego.io/jsonselect:@diagram", ptr8729511328)

	system.RegisterType("kego.io/jsonselect:@expr", ptr8729511760)

	system.RegisterType("kego.io/jsonselect:@gallery", ptr8729512192)

	system.RegisterType("kego.io/jsonselect:@image", ptr8729512624)

	system.RegisterType("kego.io/jsonselect:@kid", ptr8729513056)

	system.RegisterType("kego.io/jsonselect:@photo", ptr8729513488)

	system.RegisterType("kego.io/jsonselect:@polykids", ptr8729513920)

	system.RegisterType("kego.io/jsonselect:@sibling", ptr8729509888)

	system.RegisterType("kego.io/jsonselect:@typed", ptr8729510608)

	system.RegisterType("kego.io/jsonselect:basic", ptr8729514496)

	system.RegisterType("kego.io/jsonselect:c", ptr8729514928)

	system.RegisterType("kego.io/jsonselect:collision", ptr8729510752)

	system.RegisterType("kego.io/jsonselect:diagram", ptr8729511184)

	system.RegisterType("kego.io/jsonselect:expr", ptr8729511616)

	system.RegisterType("kego.io/jsonselect:gallery", ptr8729512048)

	system.RegisterType("kego.io/jsonselect:image", ptr8729512480)

	system.RegisterType("kego.io/jsonselect:kid", ptr8729512912)

	system.RegisterType("kego.io/jsonselect:photo", ptr8729513344)

	system.RegisterType("kego.io/jsonselect:polykids", ptr8729513776)

	system.RegisterType("kego.io/jsonselect:sibling", ptr8729514352)

	system.RegisterType("kego.io/jsonselect:typed", ptr8729510176)

}
