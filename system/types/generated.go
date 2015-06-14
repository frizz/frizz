package types

import (
	"kego.io/system"
)

func init() {

	ptr8729410208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570304 := &system.Object{Context: ptr8729410208, Description: "Automatically created basic rule for selector", Id: "@selector", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729210384 := &system.Type{Object: ptr8729570304, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729407808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569792 := &system.Object{Context: ptr8729407808, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729209952 := &system.Type{Object: ptr8729569792, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728925312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568000 := &system.Object{Context: ptr8728925312, Description: "Automatically created basic rule for object", Id: "@object", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729206784 := &system.Type{Object: ptr8729568000, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729407232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569664 := &system.Object{Context: ptr8729407232, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729209808 := &system.Type{Object: ptr8729569664, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729404192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729162368 := &system.Object{Context: ptr8729404192, Description: "Restriction rules for maps", Id: "@map", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729404896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729162496 := &system.Object{Context: ptr8729404896, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729404800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729162624 := &system.Object{Context: ptr8729404800, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729404704 := &system.Selector{Object: ptr8729162624, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729011096 := &system.Rule_rule{Selector: ptr8729404704}

	ptr8729040944 := &system.Property{Object: ptr8729162496, Defaulter: false, Item: ptr8729011096, Optional: false}

	ptr8729405344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729162752 := &system.Object{Context: ptr8729405344, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729405248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729162880 := &system.Object{Context: ptr8729405248, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729405120 := &system.Selector{Object: ptr8729162880, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729019808 := &system.Int_rule{Selector: ptr8729405120, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729041472 := &system.Property{Object: ptr8729162752, Defaulter: false, Item: ptr8729019808, Optional: true}

	ptr8729403904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729163008 := &system.Object{Context: ptr8729403904, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729403808 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729163136 := &system.Object{Context: ptr8729403808, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729403616 := &system.Selector{Object: ptr8729163136, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729018608 := &system.Int_rule{Selector: ptr8729403616, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729041136 := &system.Property{Object: ptr8729163008, Defaulter: false, Item: ptr8729018608, Optional: true}

	ptr8729208656 := &system.Type{Object: ptr8729162368, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"items": ptr8729040944, "minItems": ptr8729041472, "maxItems": ptr8729041136}, Rule: (*system.Type)(nil)}

	ptr8729409280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729163392 := &system.Object{Context: ptr8729409280, Description: "Restriction rules for numbers", Id: "@number", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729406624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729163520 := &system.Object{Context: ptr8729406624, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729406400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729163648 := &system.Object{Context: ptr8729406400, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729406304 := &system.Selector{Object: ptr8729163648, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729018848 := &system.Number_rule{Selector: ptr8729406304, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729037776 := &system.Property{Object: ptr8729163520, Defaulter: true, Item: ptr8729018848, Optional: true}

	ptr8729407168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729163776 := &system.Object{Context: ptr8729407168, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729406944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729163904 := &system.Object{Context: ptr8729406944, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729406848 := &system.Selector{Object: ptr8729163904, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729018928 := &system.Number_rule{Selector: ptr8729406848, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729040560 := &system.Property{Object: ptr8729163776, Defaulter: false, Item: ptr8729018928, Optional: true}

	ptr8729407712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729164032 := &system.Object{Context: ptr8729407712, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729407488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729164160 := &system.Object{Context: ptr8729407488, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729407392 := &system.Selector{Object: ptr8729164160, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729019008 := &system.Number_rule{Selector: ptr8729407392, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729040704 := &system.Property{Object: ptr8729164032, Defaulter: false, Item: ptr8729019008, Optional: true}

	ptr8729408160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729164288 := &system.Object{Context: ptr8729408160, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729408064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729164416 := &system.Object{Context: ptr8729408064, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729407936 := &system.Selector{Object: ptr8729164416, Selector: system.String{Value: ":root", Exists: true}}

	ptr8728844112 := &system.Bool_rule{Selector: ptr8729407936, Default: system.Bool{Value: false, Exists: true}}

	ptr8729040848 := &system.Property{Object: ptr8729164288, Defaulter: false, Item: ptr8728844112, Optional: true}

	ptr8729408704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729164544 := &system.Object{Context: ptr8729408704, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729408480 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729164672 := &system.Object{Context: ptr8729408480, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729408384 := &system.Selector{Object: ptr8729164672, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729019088 := &system.Number_rule{Selector: ptr8729408384, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729041184 := &system.Property{Object: ptr8729164544, Defaulter: false, Item: ptr8729019088, Optional: true}

	ptr8729409152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729164800 := &system.Object{Context: ptr8729409152, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729409056 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729164928 := &system.Object{Context: ptr8729409056, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729408928 := &system.Selector{Object: ptr8729164928, Selector: system.String{Value: ":root", Exists: true}}

	ptr8728844656 := &system.Bool_rule{Selector: ptr8729408928, Default: system.Bool{Value: false, Exists: true}}

	ptr8729041328 := &system.Property{Object: ptr8729164800, Defaulter: false, Item: ptr8728844656, Optional: true}

	ptr8729207648 := &system.Type{Object: ptr8729163392, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729037776, "multipleOf": ptr8729040560, "minimum": ptr8729040704, "exclusiveMinimum": ptr8729040848, "maximum": ptr8729041184, "exclusiveMaximum": ptr8729041328}, Rule: (*system.Type)(nil)}

	ptr8728926048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729160704 := &system.Object{Context: ptr8728926048, Description: "Unmarshal context.", Id: "context", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729207360 := &system.Type{Object: ptr8729160704, Exclude: true, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729406560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569152 := &system.Object{Context: ptr8729406560, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729406464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569280 := &system.Object{Context: ptr8729406464, Description: "Restriction rules for references", Id: "@reference", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729406240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569408 := &system.Object{Context: ptr8729406240, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729405952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569536 := &system.Object{Context: ptr8729405952, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729405824 := &system.Selector{Object: ptr8729569536, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729380416 := &system.Reference_rule{Selector: ptr8729405824, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729577056 := &system.Property{Object: ptr8729569408, Defaulter: true, Item: ptr8729380416, Optional: true}

	ptr8729209520 := &system.Type{Object: ptr8729569280, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729577056}, Rule: (*system.Type)(nil)}

	ptr8729209376 := &system.Type{Object: ptr8729569152, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729209520}

	ptr8729403776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729160960 := &system.Object{Context: ptr8729403776, Description: "This is the int data type", Id: "int", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729403680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729161088 := &system.Object{Context: ptr8729403680, Description: "Restriction rules for integers", Id: "@int", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728927584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729161472 := &system.Object{Context: ptr8728927584, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728927488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729161600 := &system.Object{Context: ptr8728927488, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728927392 := &system.Selector{Object: ptr8729161600, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729019408 := &system.Int_rule{Selector: ptr8728927392, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729040080 := &system.Property{Object: ptr8729161472, Defaulter: false, Item: ptr8729019408, Optional: true}

	ptr8728928000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729161728 := &system.Object{Context: ptr8728928000, Description: "This provides a lower bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728927904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729161856 := &system.Object{Context: ptr8728927904, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728927808 := &system.Selector{Object: ptr8729161856, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729019488 := &system.Int_rule{Selector: ptr8728927808, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729040176 := &system.Property{Object: ptr8729161728, Defaulter: false, Item: ptr8729019488, Optional: true}

	ptr8729403552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729161984 := &system.Object{Context: ptr8729403552, Description: "This provides an upper bound for the restriction", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729403456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729162112 := &system.Object{Context: ptr8729403456, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728928224 := &system.Selector{Object: ptr8729162112, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729019568 := &system.Int_rule{Selector: ptr8728928224, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729040320 := &system.Property{Object: ptr8729161984, Defaulter: false, Item: ptr8729019568, Optional: true}

	ptr8728927168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729161216 := &system.Object{Context: ptr8728927168, Description: "Default value if this property is omitted", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728927072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729161344 := &system.Object{Context: ptr8728927072, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728926976 := &system.Selector{Object: ptr8729161344, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729019328 := &system.Int_rule{Selector: ptr8728926976, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729039936 := &system.Property{Object: ptr8729161216, Defaulter: true, Item: ptr8729019328, Optional: true}

	ptr8729208224 := &system.Type{Object: ptr8729161088, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"multipleOf": ptr8729040080, "minimum": ptr8729040176, "maximum": ptr8729040320, "default": ptr8729039936}, Rule: (*system.Type)(nil)}

	ptr8729208080 := &system.Type{Object: ptr8729160960, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729208224}

	ptr8728924160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729157632 := &system.Object{Context: ptr8728924160, Description: "This is the native json array data type", Id: "array", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728924064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729158912 := &system.Object{Context: ptr8728924064, Description: "Restriction rules for arrays", Id: "@array", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728922720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729159680 := &system.Object{Context: ptr8728922720, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728922592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729159808 := &system.Object{Context: ptr8728922592, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728922400 := &system.Selector{Object: ptr8729159808, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729018528 := &system.Int_rule{Selector: ptr8728922400, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729036528 := &system.Property{Object: ptr8729159680, Defaulter: false, Item: ptr8729018528, Optional: true}

	ptr8728923936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729159936 := &system.Object{Context: ptr8728923936, Description: "If this is true, each item must be unique", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728923840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729160064 := &system.Object{Context: ptr8728923840, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728923712 := &system.Selector{Object: ptr8729160064, Selector: system.String{Value: ":root", Exists: true}}

	ptr8728841600 := &system.Bool_rule{Selector: ptr8728923712, Default: system.Bool{Value: false, Exists: true}}

	ptr8729036912 := &system.Property{Object: ptr8729159936, Defaulter: false, Item: ptr8728841600, Optional: true}

	ptr8728923264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729159168 := &system.Object{Context: ptr8728923264, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728923168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729159296 := &system.Object{Context: ptr8728923168, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728923072 := &system.Selector{Object: ptr8729159296, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729010520 := &system.Rule_rule{Selector: ptr8728923072}

	ptr8729039744 := &system.Property{Object: ptr8729159168, Defaulter: false, Item: ptr8729010520, Optional: false}

	ptr8728921696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729159424 := &system.Object{Context: ptr8728921696, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728921408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729159552 := &system.Object{Context: ptr8728921408, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8728921088 := &system.Selector{Object: ptr8729159552, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729023728 := &system.Int_rule{Selector: ptr8728921088, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729040512 := &system.Property{Object: ptr8729159424, Defaulter: false, Item: ptr8729023728, Optional: true}

	ptr8729213120 := &system.Type{Object: ptr8729158912, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"maxItems": ptr8729036528, "uniqueItems": ptr8729036912, "items": ptr8729039744, "minItems": ptr8729040512}, Rule: (*system.Type)(nil)}

	ptr8729212976 := &system.Type{Object: ptr8729157632, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729213120}

	ptr8729405600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572608 := &system.Object{Context: ptr8729405600, Description: "This is the most basic type.", Id: "type", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729499520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573632 := &system.Object{Context: ptr8729499520, Description: "Is this type an interface?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729499488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573760 := &system.Object{Context: ptr8729499488, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729011520 := &system.JsonBool_rule{Object: ptr8729573760}

	ptr8729577152 := &system.Property{Object: ptr8729573632, Defaulter: false, Item: ptr8729011520, Optional: true}

	ptr8729499936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573888 := &system.Object{Context: ptr8729499936, Description: "Exclude this type from the generated json?", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729499904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574016 := &system.Object{Context: ptr8729499904, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729011552 := &system.JsonBool_rule{Object: ptr8729574016}

	ptr8729577200 := &system.Property{Object: ptr8729573888, Defaulter: false, Item: ptr8729011552, Optional: true}

	ptr8729501312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574144 := &system.Object{Context: ptr8729501312, Description: "Each field is listed with it's type", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729501216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574272 := &system.Object{Context: ptr8729501216, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729500544 := &system.Selector{Object: ptr8729574272, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729501120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574400 := &system.Object{Context: ptr8729501120, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@property", Package: "kego.io/system", Type: "@property", Exists: true}}

	ptr8729501024 := &system.Selector{Object: ptr8729574400, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729011688 := &system.Property_rule{Selector: ptr8729501024}

	ptr8729380480 := &system.Map_rule{Selector: ptr8729500544, Items: ptr8729011688, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729577392 := &system.Property{Object: ptr8729574144, Defaulter: false, Item: ptr8729380480, Optional: true}

	ptr8729405216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574528 := &system.Object{Context: ptr8729405216, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729404928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574656 := &system.Object{Context: ptr8729404928, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@type", Package: "kego.io/system", Type: "@type", Exists: true}}

	ptr8729404480 := &system.Selector{Object: ptr8729574656, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729011792 := &system.Type_rule{Selector: ptr8729404480}

	ptr8729578928 := &system.Property{Object: ptr8729574528, Defaulter: false, Item: ptr8729011792, Optional: true}

	ptr8729497600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572736 := &system.Object{Context: ptr8729497600, Description: "Type which this should extend", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729497504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572864 := &system.Object{Context: ptr8729497504, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729497280 := &system.Selector{Object: ptr8729572864, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729379712 := &system.Reference_rule{Selector: ptr8729497280, Default: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}}

	ptr8729576624 := &system.Property{Object: ptr8729572736, Defaulter: false, Item: ptr8729379712, Optional: true}

	ptr8729498336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572992 := &system.Object{Context: ptr8729498336, Description: "Array of interface types that this type should support", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729498176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573120 := &system.Object{Context: ptr8729498176, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729497856 := &system.Selector{Object: ptr8729573120, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729498080 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573248 := &system.Object{Context: ptr8729498080, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729497984 := &system.Selector{Object: ptr8729573248, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729380288 := &system.Reference_rule{Selector: ptr8729497984, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729380160 := &system.Array_rule{Selector: ptr8729497856, Items: ptr8729380288, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729576672 := &system.Property{Object: ptr8729572992, Defaulter: false, Item: ptr8729380160, Optional: true}

	ptr8729499072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573376 := &system.Object{Context: ptr8729499072, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729498976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729573504 := &system.Object{Context: ptr8729498976, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729498592 := &system.Selector{Object: ptr8729573504, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729145888 := &system.String_rule{Selector: ptr8729498592, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729576960 := &system.Property{Object: ptr8729573376, Defaulter: false, Item: ptr8729145888, Optional: true}

	ptr8729209664 := &system.Type{Object: ptr8729572608, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"interface": ptr8729577152, "exclude": ptr8729577200, "properties": ptr8729577392, "rule": ptr8729578928, "extends": ptr8729576624, "is": ptr8729576672, "native": ptr8729576960}, Rule: (*system.Type)(nil)}

	ptr8729409504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569920 := &system.Object{Context: ptr8729409504, Description: "All rules should extend this type.", Id: "selector", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729409344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570048 := &system.Object{Context: ptr8729409344, Description: "Json selector defining what nodes this rule should be applied to.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729409216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570176 := &system.Object{Context: ptr8729409216, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729409024 := &system.Selector{Object: ptr8729570176, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729143488 := &system.String_rule{Selector: ptr8729409024, Default: system.String{Value: ":root", Exists: true}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729577584 := &system.Property{Object: ptr8729570048, Defaulter: false, Item: ptr8729143488, Optional: false}

	ptr8729210240 := &system.Type{Object: ptr8729569920, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"selector": ptr8729577584}, Rule: (*system.Type)(nil)}

	ptr8729404352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729162240 := &system.Object{Context: ptr8729404352, Description: "This is the native json object data type.", Id: "map", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729208512 := &system.Type{Object: ptr8729162240, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729208656}

	ptr8728926240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729160832 := &system.Object{Context: ptr8728926240, Description: "Automatically created basic rule for context", Id: "@context", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729207504 := &system.Type{Object: ptr8729160832, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728925408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729160320 := &system.Object{Context: ptr8728925408, Description: "Restriction rules for bools", Id: "@bool", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728925280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729160448 := &system.Object{Context: ptr8728925280, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728925184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729160576 := &system.Object{Context: ptr8728925184, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728925088 := &system.Selector{Object: ptr8729160576, Selector: system.String{Value: ":root", Exists: true}}

	ptr8728843280 := &system.Bool_rule{Selector: ptr8728925088, Default: system.Bool{Value: false, Exists: false}}

	ptr8729039168 := &system.Property{Object: ptr8729160448, Defaulter: true, Item: ptr8728843280, Optional: true}

	ptr8729207072 := &system.Type{Object: ptr8729160320, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729039168}, Rule: (*system.Type)(nil)}

	ptr8729496448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570560 := &system.Object{Context: ptr8729496448, Description: "Restriction rules for strings", Id: "@string", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729493760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570944 := &system.Object{Context: ptr8729493760, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729493536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571072 := &system.Object{Context: ptr8729493536, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729493792 := &system.Selector{Object: ptr8729571072, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729494016 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571200 := &system.Object{Context: ptr8729494016, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729493920 := &system.Selector{Object: ptr8729571200, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729144928 := &system.String_rule{Selector: ptr8729493920, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729381824 := &system.Array_rule{Selector: ptr8729493792, Items: ptr8729144928, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729578448 := &system.Property{Object: ptr8729570944, Defaulter: false, Item: ptr8729381824, Optional: true}

	ptr8729494400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571328 := &system.Object{Context: ptr8729494400, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729494304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571456 := &system.Object{Context: ptr8729494304, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729494208 := &system.Selector{Object: ptr8729571456, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729018768 := &system.Int_rule{Selector: ptr8729494208, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729575568 := &system.Property{Object: ptr8729571328, Defaulter: false, Item: ptr8729018768, Optional: true}

	ptr8729494816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571584 := &system.Object{Context: ptr8729494816, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729494720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571712 := &system.Object{Context: ptr8729494720, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729494624 := &system.Selector{Object: ptr8729571712, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729019248 := &system.Int_rule{Selector: ptr8729494624, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729575664 := &system.Property{Object: ptr8729571584, Defaulter: false, Item: ptr8729019248, Optional: true}

	ptr8729495232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571840 := &system.Object{Context: ptr8729495232, Description: "This is a string that the value must match", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729495136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729571968 := &system.Object{Context: ptr8729495136, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729495040 := &system.Selector{Object: ptr8729571968, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729142208 := &system.String_rule{Selector: ptr8729495040, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729575760 := &system.Property{Object: ptr8729571840, Defaulter: false, Item: ptr8729142208, Optional: true}

	ptr8729495648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572096 := &system.Object{Context: ptr8729495648, Description: "This is a regex to match the value to", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729495552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572224 := &system.Object{Context: ptr8729495552, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729495456 := &system.Selector{Object: ptr8729572224, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729142848 := &system.String_rule{Selector: ptr8729495456, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729576000 := &system.Property{Object: ptr8729572096, Defaulter: false, Item: ptr8729142848, Optional: true}

	ptr8729496320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572352 := &system.Object{Context: ptr8729496320, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729496224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729572480 := &system.Object{Context: ptr8729496224, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729495872 := &system.Selector{Object: ptr8729572480, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729145568 := &system.String_rule{Selector: ptr8729495872, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729576240 := &system.Property{Object: ptr8729572352, Defaulter: false, Item: ptr8729145568, Optional: true}

	ptr8729493568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570688 := &system.Object{Context: ptr8729493568, Description: "Default value of this is missing or null", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729411520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570816 := &system.Object{Context: ptr8729411520, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729411424 := &system.Selector{Object: ptr8729570816, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729144128 := &system.String_rule{Selector: ptr8729411424, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729578304 := &system.Property{Object: ptr8729570688, Defaulter: true, Item: ptr8729144128, Optional: true}

	ptr8729210816 := &system.Type{Object: ptr8729570560, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"enum": ptr8729578448, "minLength": ptr8729575568, "maxLength": ptr8729575664, "equal": ptr8729575760, "pattern": ptr8729576000, "format": ptr8729576240, "default": ptr8729578304}, Rule: (*system.Type)(nil)}

	ptr8729406688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729574784 := &system.Object{Context: ptr8729406688, Description: "Automatically created basic rule for type", Id: "@type", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729210960 := &system.Type{Object: ptr8729574784, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728925504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729160192 := &system.Object{Context: ptr8728925504, Description: "This is the native json bool data type", Id: "bool", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729206928 := &system.Type{Object: ptr8729160192, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729207072}

	ptr8728928192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568128 := &system.Object{Context: ptr8728928192, Description: "This is a property of a type", Id: "property", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728928032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568768 := &system.Object{Context: ptr8728928032, Description: "This is a rule object, defining the type and restrictions on the value of the this property", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728927840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568896 := &system.Object{Context: ptr8728927840, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728927680 := &system.Selector{Object: ptr8729568896, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729010984 := &system.Rule_rule{Selector: ptr8728927680}

	ptr8729576768 := &system.Property{Object: ptr8729568768, Defaulter: false, Item: ptr8729010984, Optional: false}

	ptr8728926560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568256 := &system.Object{Context: ptr8728926560, Description: "This specifies that the field is optional", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728926496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568384 := &system.Object{Context: ptr8728926496, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729010888 := &system.JsonBool_rule{Object: ptr8729568384}

	ptr8729576384 := &system.Property{Object: ptr8729568256, Defaulter: false, Item: ptr8729010888, Optional: true}

	ptr8728927264 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568512 := &system.Object{Context: ptr8728927264, Description: "This specifies that the field is the default value for a rule", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728927232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729568640 := &system.Object{Context: ptr8728927232, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729010960 := &system.JsonBool_rule{Object: ptr8729568640}

	ptr8729576720 := &system.Property{Object: ptr8729568512, Defaulter: false, Item: ptr8729010960, Optional: true}

	ptr8729208944 := &system.Type{Object: ptr8729568128, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"item": ptr8729576768, "optional": ptr8729576384, "defaulter": ptr8729576720}, Rule: (*system.Type)(nil)}

	ptr8729404288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729569024 := &system.Object{Context: ptr8729404288, Description: "Automatically created basic rule for property", Id: "@property", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729209088 := &system.Type{Object: ptr8729569024, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729496544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729570432 := &system.Object{Context: ptr8729496544, Description: "This is the native json string data type", Id: "string", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729210672 := &system.Type{Object: ptr8729570432, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729210816}

	ptr8729409376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729163264 := &system.Object{Context: ptr8729409376, Description: "This is the native json number data type", Id: "number", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729207216 := &system.Type{Object: ptr8729163264, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729207648}

	ptr8728924928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729165056 := &system.Object{Context: ptr8728924928, Description: "This is the most basic type.", Id: "object", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728923552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729567360 := &system.Object{Context: ptr8728923552, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728923296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729567488 := &system.Object{Context: ptr8728923296, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@context", Package: "kego.io/system", Type: "@context", Exists: true}}

	ptr8728923040 := &system.Selector{Object: ptr8729567488, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729011312 := &system.Context_rule{Selector: ptr8728923040}

	ptr8729042480 := &system.Property{Object: ptr8729567360, Defaulter: false, Item: ptr8729011312, Optional: true}

	ptr8728924736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729567616 := &system.Object{Context: ptr8728924736, Description: "Extra validation rules for this object or descendants", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728924512 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729567744 := &system.Object{Context: ptr8728924512, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728924000 := &system.Selector{Object: ptr8729567744, Selector: system.String{Value: ":root", Exists: true}}

	ptr8728924288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729567872 := &system.Object{Context: ptr8728924288, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8728924192 := &system.Selector{Object: ptr8729567872, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729010712 := &system.Rule_rule{Selector: ptr8728924192}

	ptr8729379072 := &system.Array_rule{Selector: ptr8728924000, Items: ptr8729010712, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729575424 := &system.Property{Object: ptr8729567616, Defaulter: false, Item: ptr8729379072, Optional: true}

	ptr8729410688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729165184 := &system.Object{Context: ptr8729410688, Description: "Type of the object.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729410592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729165312 := &system.Object{Context: ptr8729410592, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729410496 := &system.Selector{Object: ptr8729165312, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729379968 := &system.Reference_rule{Selector: ptr8729410496, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729041616 := &system.Property{Object: ptr8729165184, Defaulter: false, Item: ptr8729379968, Optional: false}

	ptr8729411168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729165440 := &system.Object{Context: ptr8729411168, Description: "All global objects should have an id.", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729411136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729165568 := &system.Object{Context: ptr8729411136, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8729011216 := &system.JsonString_rule{Object: ptr8729165568}

	ptr8729042096 := &system.Property{Object: ptr8729165440, Defaulter: false, Item: ptr8729011216, Optional: true}

	ptr8728920960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729165696 := &system.Object{Context: ptr8728920960, Description: "Description for the developer", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728920896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729567232 := &system.Object{Context: ptr8728920896, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8729011288 := &system.JsonString_rule{Object: ptr8729567232}

	ptr8729042432 := &system.Property{Object: ptr8729165696, Defaulter: false, Item: ptr8729011288, Optional: true}

	ptr8729208800 := &system.Type{Object: ptr8729165056, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"context": ptr8729042480, "rules": ptr8729575424, "type": ptr8729041616, "id": ptr8729042096, "description": ptr8729042432}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system:@array", ptr8729213120)

	system.RegisterType("kego.io/system:@bool", ptr8729207072)

	system.RegisterType("kego.io/system:@context", ptr8729207504)

	system.RegisterType("kego.io/system:@int", ptr8729208224)

	system.RegisterType("kego.io/system:@map", ptr8729208656)

	system.RegisterType("kego.io/system:@number", ptr8729207648)

	system.RegisterType("kego.io/system:@object", ptr8729206784)

	system.RegisterType("kego.io/system:@property", ptr8729209088)

	system.RegisterType("kego.io/system:@reference", ptr8729209520)

	system.RegisterType("kego.io/system:@rule", ptr8729209952)

	system.RegisterType("kego.io/system:@selector", ptr8729210384)

	system.RegisterType("kego.io/system:@string", ptr8729210816)

	system.RegisterType("kego.io/system:@type", ptr8729210960)

	system.RegisterType("kego.io/system:array", ptr8729212976)

	system.RegisterType("kego.io/system:bool", ptr8729206928)

	system.RegisterType("kego.io/system:context", ptr8729207360)

	system.RegisterType("kego.io/system:int", ptr8729208080)

	system.RegisterType("kego.io/system:map", ptr8729208512)

	system.RegisterType("kego.io/system:number", ptr8729207216)

	system.RegisterType("kego.io/system:object", ptr8729208800)

	system.RegisterType("kego.io/system:property", ptr8729208944)

	system.RegisterType("kego.io/system:reference", ptr8729209376)

	system.RegisterType("kego.io/system:rule", ptr8729209808)

	system.RegisterType("kego.io/system:selector", ptr8729210240)

	system.RegisterType("kego.io/system:string", ptr8729210672)

	system.RegisterType("kego.io/system:type", ptr8729209664)

}
