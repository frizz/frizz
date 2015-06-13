package types

import (
	"kego.io/system"
)

func init() {

	ptr8728993632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729150336 := &system.Object{Context: ptr8728993632, Description: "This is the native json array data type", Id: "array", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728993312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729150448 := &system.Object{Context: ptr8728993312, Description: "Restriction rules for arrays", Id: "@array", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728992000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729150784 := &system.Object{Context: ptr8728992000, Description: "This is a rule object, defining the type and restrictions on the value of the items", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728991968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729151120 := &system.Object{Context: ptr8728991968, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729076168 := &system.Rule_rule{Object: ptr8729151120}

	ptr8729145952 := &system.Property{Object: ptr8729150784, Defaulter: false, Item: ptr8729076168, Optional: false}

	ptr8728992416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729151232 := &system.Object{Context: ptr8728992416, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728992384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729151568 := &system.Object{Context: ptr8728992384, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729474608 := &system.Int_rule{Object: ptr8729151568, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729146528 := &system.Property{Object: ptr8729151232, Defaulter: false, Item: ptr8729474608, Optional: true}

	ptr8728992768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729151792 := &system.Object{Context: ptr8728992768, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728992736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729151904 := &system.Object{Context: ptr8728992736, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729474688 := &system.Int_rule{Object: ptr8729151904, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729146864 := &system.Property{Object: ptr8729151792, Defaulter: false, Item: ptr8729474688, Optional: true}

	ptr8728993184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729152016 := &system.Object{Context: ptr8728993184, Description: "If this is true, each item must be unique", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728993152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729152352 := &system.Object{Context: ptr8728993152, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728908976 := &system.Bool_rule{Object: ptr8729152352, Default: system.Bool{Value: false, Exists: true}}

	ptr8729146960 := &system.Property{Object: ptr8729152016, Defaulter: false, Item: ptr8728908976, Optional: true}

	ptr8729443216 := &system.Type{Object: ptr8729150448, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"items": ptr8729145952, "minItems": ptr8729146528, "maxItems": ptr8729146864, "uniqueItems": ptr8729146960}, Rule: (*system.Type)(nil)}

	ptr8729443072 := &system.Type{Object: ptr8729150336, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "array", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729443216}

	ptr8729581856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728929712 := &system.Object{Context: ptr8729581856, Description: "This is a reference to another object, of the form: [local id] or [import name]:[id] or [full package path]:[id]", Id: "reference", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729581440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728929936 := &system.Object{Context: ptr8729581440, Description: "Restriction rules for references", Id: "@reference", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729581216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728931392 := &system.Object{Context: ptr8729581216, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729581152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728931504 := &system.Object{Context: ptr8729581152, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729511552 := &system.Reference_rule{Object: ptr8728931504, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729148784 := &system.Property{Object: ptr8728931392, Defaulter: true, Item: ptr8729511552, Optional: true}

	ptr8729439040 := &system.Type{Object: ptr8728929936, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729148784}, Rule: (*system.Type)(nil)}

	ptr8729438896 := &system.Type{Object: ptr8728929712, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729439040}

	ptr8729615296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728931280 := &system.Object{Context: ptr8729615296, Description: "This is the most basic type.", Id: "type", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729612160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728932176 := &system.Object{Context: ptr8729612160, Description: "Type which this should extend", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729612128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728932288 := &system.Object{Context: ptr8729612128, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729510720 := &system.Reference_rule{Object: ptr8728932288, Default: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}}

	ptr8729102208 := &system.Property{Object: ptr8728932176, Defaulter: false, Item: ptr8729510720, Optional: true}

	ptr8729612704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728934304 := &system.Object{Context: ptr8729612704, Description: "Array of interface types that this type should support", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729612608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728934416 := &system.Object{Context: ptr8729612608, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729612576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728934528 := &system.Object{Context: ptr8729612576, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729511104 := &system.Reference_rule{Object: ptr8728934528, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729510848 := &system.Array_rule{Object: ptr8728934416, Items: ptr8729511104, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729102256 := &system.Property{Object: ptr8728934304, Defaulter: false, Item: ptr8729510848, Optional: true}

	ptr8729613344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728934640 := &system.Object{Context: ptr8729613344, Description: "This is the native json type that represents this type. If omitted, default is object.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729613312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728934752 := &system.Object{Context: ptr8729613312, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729243232 := &system.String_rule{Object: ptr8728934752, Default: system.String{Value: "object", Exists: true}, Enum: []system.String{system.String{Value: "string", Exists: true}, system.String{Value: "number", Exists: true}, system.String{Value: "bool", Exists: true}, system.String{Value: "array", Exists: true}, system.String{Value: "object", Exists: true}, system.String{Value: "map", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729102352 := &system.Property{Object: ptr8728934640, Defaulter: false, Item: ptr8729243232, Optional: true}

	ptr8729613792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728934864 := &system.Object{Context: ptr8729613792, Description: "Is this type an interface?", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729613760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728934976 := &system.Object{Context: ptr8729613760, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729076856 := &system.JsonBool_rule{Object: ptr8728934976}

	ptr8729102400 := &system.Property{Object: ptr8728934864, Defaulter: false, Item: ptr8729076856, Optional: true}

	ptr8729614208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728935088 := &system.Object{Context: ptr8729614208, Description: "Exclude this type from the generated json?", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729614176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728935200 := &system.Object{Context: ptr8729614176, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729076888 := &system.JsonBool_rule{Object: ptr8728935200}

	ptr8729102448 := &system.Property{Object: ptr8728935088, Defaulter: false, Item: ptr8729076888, Optional: true}

	ptr8729614752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728935312 := &system.Object{Context: ptr8729614752, Description: "Each field is listed with it's type", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729614720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728935760 := &system.Object{Context: ptr8729614720, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729614688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728936096 := &system.Object{Context: ptr8729614688, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@property", Package: "kego.io/system", Type: "@property", Exists: true}}

	ptr8729076920 := &system.Property_rule{Object: ptr8728936096}

	ptr8729511232 := &system.Map_rule{Object: ptr8728935760, Items: ptr8729076920, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729102544 := &system.Property{Object: ptr8728935312, Defaulter: false, Item: ptr8729511232, Optional: true}

	ptr8729615168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728936208 := &system.Object{Context: ptr8729615168, Description: "Embedded type that defines restriction rules for this type. By convention, the ID should be this type prefixed with the @ character.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729615136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729370496 := &system.Object{Context: ptr8729615136, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@type", Package: "kego.io/system", Type: "@type", Exists: true}}

	ptr8729076992 := &system.Type_rule{Object: ptr8729370496}

	ptr8729104608 := &system.Property{Object: ptr8728936208, Defaulter: false, Item: ptr8729076992, Optional: true}

	ptr8729438176 := &system.Type{Object: ptr8728931280, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"extends": ptr8729102208, "is": ptr8729102256, "native": ptr8729102352, "interface": ptr8729102400, "exclude": ptr8729102448, "properties": ptr8729102544, "rule": ptr8729104608}, Rule: (*system.Type)(nil)}

	ptr8729576640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729311248 := &system.Object{Context: ptr8729576640, Description: "Automatically created basic rule for object", Id: "@object", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729437168 := &system.Type{Object: ptr8729311248, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729583232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728931728 := &system.Object{Context: ptr8729583232, Description: "Automatically created basic rule for rule", Id: "@rule", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729439472 := &system.Type{Object: ptr8728931728, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729581408 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729153024 := &system.Object{Context: ptr8729581408, Description: "Restriction rules for maps", Id: "@map", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729580576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729153248 := &system.Object{Context: ptr8729580576, Description: "This is a rule object, defining the type and restrictions on the value of the items.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729580544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729153584 := &system.Object{Context: ptr8729580544, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729076544 := &system.Rule_rule{Object: ptr8729153584}

	ptr8729145808 := &system.Property{Object: ptr8729153248, Defaulter: false, Item: ptr8729076544, Optional: false}

	ptr8729580928 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729153696 := &system.Object{Context: ptr8729580928, Description: "This is the minimum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729580896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729154032 := &system.Object{Context: ptr8729580896, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469968 := &system.Int_rule{Object: ptr8729154032, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729146048 := &system.Property{Object: ptr8729153696, Defaulter: false, Item: ptr8729469968, Optional: true}

	ptr8729581280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729154144 := &system.Object{Context: ptr8729581280, Description: "This is the maximum number of items alowed in the array", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729581248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729154256 := &system.Object{Context: ptr8729581248, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729470048 := &system.Int_rule{Object: ptr8729154256, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: true}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729146144 := &system.Property{Object: ptr8729154144, Defaulter: false, Item: ptr8729470048, Optional: true}

	ptr8729438032 := &system.Type{Object: ptr8729153024, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"items": ptr8729145808, "minItems": ptr8729146048, "maxItems": ptr8729146144}, Rule: (*system.Type)(nil)}

	ptr8728988032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729154368 := &system.Object{Context: ptr8728988032, Description: "This is the native json number data type", Id: "number", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728986848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729154480 := &system.Object{Context: ptr8728986848, Description: "Restriction rules for numbers", Id: "@number", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728992128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729156384 := &system.Object{Context: ptr8728992128, Description: "This provides an upper bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728991872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729156608 := &system.Object{Context: ptr8728991872, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729470448 := &system.Number_rule{Object: ptr8729156608, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729147632 := &system.Property{Object: ptr8729156384, Defaulter: false, Item: ptr8729470448, Optional: true}

	ptr8728986720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729156720 := &system.Object{Context: ptr8728986720, Description: "If this is true, the value must be less than maximum. If false or not provided, the value must be less than or equal to the maximum.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728986688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729154816 := &system.Object{Context: ptr8728986688, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729460912 := &system.Bool_rule{Object: ptr8729154816, Default: system.Bool{Value: false, Exists: true}}

	ptr8729147968 := &system.Property{Object: ptr8729156720, Defaulter: false, Item: ptr8729460912, Optional: true}

	ptr8729583136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729154704 := &system.Object{Context: ptr8729583136, Description: "Default value if this property is omitted", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729582976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729155264 := &system.Object{Context: ptr8729582976, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729470208 := &system.Number_rule{Object: ptr8729155264, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729146432 := &system.Property{Object: ptr8729154704, Defaulter: true, Item: ptr8729470208, Optional: true}

	ptr8729583584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729155376 := &system.Object{Context: ptr8729583584, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729583424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729155488 := &system.Object{Context: ptr8729583424, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729470288 := &system.Number_rule{Object: ptr8729155488, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729147776 := &system.Property{Object: ptr8729155376, Defaulter: false, Item: ptr8729470288, Optional: true}

	ptr8728990464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729155600 := &system.Object{Context: ptr8728990464, Description: "This provides a lower bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728990112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729155824 := &system.Object{Context: ptr8728990112, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729470368 := &system.Number_rule{Object: ptr8729155824, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729147392 := &system.Property{Object: ptr8729155600, Defaulter: false, Item: ptr8729470368, Optional: true}

	ptr8728991168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729156048 := &system.Object{Context: ptr8728991168, Description: "If this is true, the value must be greater than minimum. If false or not provided, the value must be greater than or equal to the minimum.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728991104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729156160 := &system.Object{Context: ptr8728991104, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729461984 := &system.Bool_rule{Object: ptr8729156160, Default: system.Bool{Value: false, Exists: true}}

	ptr8729147536 := &system.Property{Object: ptr8729156048, Defaulter: false, Item: ptr8729461984, Optional: true}

	ptr8729438464 := &system.Type{Object: ptr8729154480, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"maximum": ptr8729147632, "exclusiveMaximum": ptr8729147968, "default": ptr8729146432, "multipleOf": ptr8729147776, "minimum": ptr8729147392, "exclusiveMinimum": ptr8729147536}, Rule: (*system.Type)(nil)}

	ptr8729438320 := &system.Type{Object: ptr8729154368, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729438464}

	ptr8729580032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728929600 := &system.Object{Context: ptr8729580032, Description: "Automatically created basic rule for property", Id: "@property", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729442928 := &system.Type{Object: ptr8728929600, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729611200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728931840 := &system.Object{Context: ptr8729611200, Description: "This is the native json string data type", Id: "string", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729610880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728931952 := &system.Object{Context: ptr8729610880, Description: "Restriction rules for strings", Id: "@string", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729608736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728932064 := &system.Object{Context: ptr8729608736, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729608704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728935424 := &system.Object{Context: ptr8729608704, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729243552 := &system.String_rule{Object: ptr8728935424, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729149312 := &system.Property{Object: ptr8728932064, Defaulter: true, Item: ptr8729243552, Optional: true}

	ptr8729608832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728935536 := &system.Object{Context: ptr8729608832, Description: "The value of this string is restricted to one of the provided values", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729608640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728929488 := &system.Object{Context: ptr8729608640, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729608608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728930048 := &system.Object{Context: ptr8729608608, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729240512 := &system.String_rule{Object: ptr8728930048, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729512512 := &system.Array_rule{Object: ptr8728929488, Items: ptr8729240512, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729102592 := &system.Property{Object: ptr8728935536, Defaulter: false, Item: ptr8729512512, Optional: true}

	ptr8729609216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728930160 := &system.Object{Context: ptr8729609216, Description: "The value must be longer or equal to the provided minimum length", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729609184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728930272 := &system.Object{Context: ptr8729609184, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469008 := &system.Int_rule{Object: ptr8728930272, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729101488 := &system.Property{Object: ptr8728930160, Defaulter: false, Item: ptr8729469008, Optional: true}

	ptr8729609536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728930384 := &system.Object{Context: ptr8729609536, Description: "The value must be shorter or equal to the provided maximum length", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729609504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728930496 := &system.Object{Context: ptr8729609504, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469168 := &system.Int_rule{Object: ptr8728930496, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729101584 := &system.Property{Object: ptr8728930384, Defaulter: false, Item: ptr8729469168, Optional: true}

	ptr8729609856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728930608 := &system.Object{Context: ptr8729609856, Description: "This is a string that the value must match", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729609824 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728930720 := &system.Object{Context: ptr8729609824, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729241792 := &system.String_rule{Object: ptr8728930720, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729101680 := &system.Property{Object: ptr8728930608, Defaulter: false, Item: ptr8729241792, Optional: true}

	ptr8729610176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728930832 := &system.Object{Context: ptr8729610176, Description: "This is a regex to match the value to", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729610144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728930944 := &system.Object{Context: ptr8729610144, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729242432 := &system.String_rule{Object: ptr8728930944, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729101824 := &system.Property{Object: ptr8728930832, Defaulter: false, Item: ptr8729242432, Optional: true}

	ptr8729610752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728931056 := &system.Object{Context: ptr8729610752, Description: "This restricts the value to one of several built-in formats", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729610720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728931168 := &system.Object{Context: ptr8729610720, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729242912 := &system.String_rule{Object: ptr8728931168, Default: system.String{Value: "", Exists: false}, Enum: []system.String{system.String{Value: "date-time", Exists: true}, system.String{Value: "email", Exists: true}, system.String{Value: "hostname", Exists: true}, system.String{Value: "ipv4", Exists: true}, system.String{Value: "ipv6", Exists: true}, system.String{Value: "uri", Exists: true}}, Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729101968 := &system.Property{Object: ptr8728931056, Defaulter: false, Item: ptr8729242912, Optional: true}

	ptr8729439904 := &system.Type{Object: ptr8728931952, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729149312, "enum": ptr8729102592, "minLength": ptr8729101488, "maxLength": ptr8729101584, "equal": ptr8729101680, "pattern": ptr8729101824, "format": ptr8729101968}, Rule: (*system.Type)(nil)}

	ptr8729439760 := &system.Type{Object: ptr8728931840, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "string", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729439904}

	ptr8729576704 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729149440 := &system.Object{Context: ptr8729576704, Description: "This is the native json bool data type", Id: "bool", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729576384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729149664 := &system.Object{Context: ptr8729576384, Description: "Restriction rules for bools", Id: "@bool", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729576256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729149776 := &system.Object{Context: ptr8729576256, Description: "Default value of this is missing or null", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729576224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729149888 := &system.Object{Context: ptr8729576224, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8728906704 := &system.Bool_rule{Object: ptr8729149888, Default: system.Bool{Value: false, Exists: false}}

	ptr8729142832 := &system.Property{Object: ptr8729149776, Defaulter: true, Item: ptr8728906704, Optional: true}

	ptr8729436448 := &system.Type{Object: ptr8729149664, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"default": ptr8729142832}, Rule: (*system.Type)(nil)}

	ptr8729436304 := &system.Type{Object: ptr8729149440, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "bool", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729436448}

	ptr8729579296 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729150896 := &system.Object{Context: ptr8729579296, Description: "Restriction rules for integers", Id: "@int", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729579168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729152688 := &system.Object{Context: ptr8729579168, Description: "This provides an upper bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729579136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729152800 := &system.Object{Context: ptr8729579136, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469808 := &system.Int_rule{Object: ptr8729152800, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729145328 := &system.Property{Object: ptr8729152688, Defaulter: false, Item: ptr8729469808, Optional: true}

	ptr8729578208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729151008 := &system.Object{Context: ptr8729578208, Description: "Default value if this property is omitted", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729578176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729151344 := &system.Object{Context: ptr8729578176, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469568 := &system.Int_rule{Object: ptr8729151344, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729144128 := &system.Property{Object: ptr8729151008, Defaulter: true, Item: ptr8729469568, Optional: true}

	ptr8729578528 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729151456 := &system.Object{Context: ptr8729578528, Description: "This restricts the number to be a multiple of the given number", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729578496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729152128 := &system.Object{Context: ptr8729578496, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469648 := &system.Int_rule{Object: ptr8729152128, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729145040 := &system.Property{Object: ptr8729151456, Defaulter: false, Item: ptr8729469648, Optional: true}

	ptr8729578848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729152240 := &system.Object{Context: ptr8729578848, Description: "This provides a lower bound for the restriction", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729578816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729152464 := &system.Object{Context: ptr8729578816, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@int", Package: "kego.io/system", Type: "@int", Exists: true}}

	ptr8729469728 := &system.Int_rule{Object: ptr8729152464, Default: system.Int{Value: 0, Exists: false}, Maximum: system.Int{Value: 0, Exists: false}, Minimum: system.Int{Value: 0, Exists: false}, MultipleOf: system.Int{Value: 0, Exists: false}}

	ptr8729145184 := &system.Property{Object: ptr8729152240, Defaulter: false, Item: ptr8729469728, Optional: true}

	ptr8729437600 := &system.Type{Object: ptr8729150896, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"maximum": ptr8729145328, "default": ptr8729144128, "multipleOf": ptr8729145040, "minimum": ptr8729145184}, Rule: (*system.Type)(nil)}

	ptr8729577440 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729150560 := &system.Object{Context: ptr8729577440, Description: "Automatically created basic rule for context", Id: "@context", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729436880 := &system.Type{Object: ptr8729150560, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729579616 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729150672 := &system.Object{Context: ptr8729579616, Description: "This is the int data type", Id: "int", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729437456 := &system.Type{Object: ptr8729150672, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "number", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729437600}

	ptr8729582656 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8728931616 := &system.Object{Context: ptr8729582656, Description: "This is one of several rule types, derived from the rules property of other types", Id: "rule", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729439328 := &system.Type{Object: ptr8728931616, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729581728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729152912 := &system.Object{Context: ptr8729581728, Description: "This is the native json object data type.", Id: "map", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729437888 := &system.Type{Object: ptr8729152912, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "map", Exists: true}, Properties: map[string]*system.Property(nil), Rule: ptr8729438032}

	ptr8729577248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729150112 := &system.Object{Context: ptr8729577248, Description: "Unmarshal context.", Id: "context", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729436736 := &system.Type{Object: ptr8729150112, Exclude: true, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729615712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729362432 := &system.Object{Context: ptr8729615712, Description: "Automatically created basic rule for type", Id: "@type", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729440048 := &system.Type{Object: ptr8729362432, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729579392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729311472 := &system.Object{Context: ptr8729579392, Description: "This is a property of a type", Id: "property", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729577920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729311696 := &system.Object{Context: ptr8729577920, Description: "This specifies that the field is optional", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729577888 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729312144 := &system.Object{Context: ptr8729577888, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729076688 := &system.JsonBool_rule{Object: ptr8729312144}

	ptr8729148208 := &system.Property{Object: ptr8729311696, Defaulter: false, Item: ptr8729076688, Optional: true}

	ptr8729578624 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729312368 := &system.Object{Context: ptr8729578624, Description: "This specifies that the field is the default value for a rule", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729578592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729312816 := &system.Object{Context: ptr8729578592, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@bool", Package: "kego.io/json", Type: "@bool", Exists: true}}

	ptr8729076848 := &system.JsonBool_rule{Object: ptr8729312816}

	ptr8729148448 := &system.Property{Object: ptr8729312368, Defaulter: false, Item: ptr8729076848, Optional: true}

	ptr8729579232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729312928 := &system.Object{Context: ptr8729579232, Description: "This is a rule object, defining the type and restrictions on the value of the this property", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729579200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729313152 := &system.Object{Context: ptr8729579200, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729076872 := &system.Rule_rule{Object: ptr8729313152}

	ptr8729148496 := &system.Property{Object: ptr8729312928, Defaulter: false, Item: ptr8729076872, Optional: false}

	ptr8729437744 := &system.Type{Object: ptr8729311472, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"optional": ptr8729148208, "defaulter": ptr8729148448, "item": ptr8729148496}, Rule: (*system.Type)(nil)}

	ptr8729576352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729155152 := &system.Object{Context: ptr8729576352, Description: "This is the most basic type.", Id: "object", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729576128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729308336 := &system.Object{Context: ptr8729576128, Description: "Extra validation rules for this object or descendants", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729576064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729310800 := &system.Object{Context: ptr8729576064, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729576000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729311024 := &system.Object{Context: ptr8729576000, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@rule", Package: "kego.io/system", Type: "@rule", Exists: true}}

	ptr8729076616 := &system.Rule_rule{Object: ptr8729311024}

	ptr8729510400 := &system.Map_rule{Object: ptr8729310800, Items: ptr8729076616, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729147056 := &system.Property{Object: ptr8729308336, Defaulter: false, Item: ptr8729510400, Optional: true}

	ptr8728991040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729156944 := &system.Object{Context: ptr8728991040, Description: "Type of the object.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728990944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729157392 := &system.Object{Context: ptr8728990944, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@reference", Package: "kego.io/system", Type: "@reference", Exists: true}}

	ptr8729510272 := &system.Reference_rule{Object: ptr8729157392, Default: system.Reference{Value: "", Package: "", Type: "", Exists: false}}

	ptr8729143744 := &system.Property{Object: ptr8729156944, Defaulter: false, Item: ptr8729510272, Optional: false}

	ptr8728993760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729305088 := &system.Object{Context: ptr8728993760, Description: "All global objects should have an id.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728993664 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729307552 := &system.Object{Context: ptr8728993664, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8729076232 := &system.JsonString_rule{Object: ptr8729307552}

	ptr8729145520 := &system.Property{Object: ptr8729305088, Defaulter: false, Item: ptr8729076232, Optional: true}

	ptr8728992960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729307664 := &system.Object{Context: ptr8728992960, Description: "Description for the developer", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728992896 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729307776 := &system.Object{Context: ptr8728992896, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/json:@string", Package: "kego.io/json", Type: "@string", Exists: true}}

	ptr8729076432 := &system.JsonString_rule{Object: ptr8729307776}

	ptr8729146576 := &system.Property{Object: ptr8729307664, Defaulter: false, Item: ptr8729076432, Optional: true}

	ptr8728993568 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729307888 := &system.Object{Context: ptr8728993568, Description: "Unmarshaling context. This should not be in the json - it's added by the unmarshaler.", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728993504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/system"}

	ptr8729308224 := &system.Object{Context: ptr8728993504, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@context", Package: "kego.io/system", Type: "@context", Exists: true}}

	ptr8729076456 := &system.Context_rule{Object: ptr8729308224}

	ptr8729146624 := &system.Property{Object: ptr8729307888, Defaulter: false, Item: ptr8729076456, Optional: true}

	ptr8729437024 := &system.Type{Object: ptr8729155152, Exclude: false, Extends: system.Reference{Value: "", Package: "", Type: "", Exists: false}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"rules": ptr8729147056, "type": ptr8729143744, "id": ptr8729145520, "description": ptr8729146576, "context": ptr8729146624}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/system:@array", ptr8729443216)

	system.RegisterType("kego.io/system:@bool", ptr8729436448)

	system.RegisterType("kego.io/system:@context", ptr8729436880)

	system.RegisterType("kego.io/system:@int", ptr8729437600)

	system.RegisterType("kego.io/system:@map", ptr8729438032)

	system.RegisterType("kego.io/system:@number", ptr8729438464)

	system.RegisterType("kego.io/system:@object", ptr8729437168)

	system.RegisterType("kego.io/system:@property", ptr8729442928)

	system.RegisterType("kego.io/system:@reference", ptr8729439040)

	system.RegisterType("kego.io/system:@rule", ptr8729439472)

	system.RegisterType("kego.io/system:@string", ptr8729439904)

	system.RegisterType("kego.io/system:@type", ptr8729440048)

	system.RegisterType("kego.io/system:array", ptr8729443072)

	system.RegisterType("kego.io/system:bool", ptr8729436304)

	system.RegisterType("kego.io/system:context", ptr8729436736)

	system.RegisterType("kego.io/system:int", ptr8729437456)

	system.RegisterType("kego.io/system:map", ptr8729437888)

	system.RegisterType("kego.io/system:number", ptr8729438320)

	system.RegisterType("kego.io/system:object", ptr8729437024)

	system.RegisterType("kego.io/system:property", ptr8729437744)

	system.RegisterType("kego.io/system:reference", ptr8729438896)

	system.RegisterType("kego.io/system:rule", ptr8729439328)

	system.RegisterType("kego.io/system:string", ptr8729439760)

	system.RegisterType("kego.io/system:type", ptr8729438176)

}
