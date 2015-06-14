package types

import (
	"kego.io/system"

	_ "kego.io/system/types"

	"kego.io/jsonselect"
)

func init() {

	ptr8729682496 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729697792 := &system.Object{Context: ptr8729682496, Description: "", Id: "kid", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729140768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729698176 := &system.Object{Context: ptr8729140768, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729140640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729493248 := &system.Object{Context: ptr8729140640, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729140480 := &system.Selector{Object: ptr8729493248, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729362240 := &system.String_rule{Selector: ptr8729140480, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729500992 := &system.Property{Object: ptr8729698176, Defaulter: false, Item: ptr8729362240, Optional: false}

	ptr8729682336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729493376 := &system.Object{Context: ptr8729682336, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729682208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729370624 := &system.Object{Context: ptr8729682208, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729681984 := &system.Selector{Object: ptr8729370624, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729778112 := &system.Bool_rule{Selector: ptr8729681984, Default: system.Bool{Value: false, Exists: false}}

	ptr8729501040 := &system.Property{Object: ptr8729493376, Defaulter: false, Item: ptr8729778112, Optional: true}

	ptr8729140032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729697920 := &system.Object{Context: ptr8729140032, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729139840 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729698048 := &system.Object{Context: ptr8729139840, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729139680 := &system.Selector{Object: ptr8729698048, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729362080 := &system.String_rule{Selector: ptr8729139680, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729501376 := &system.Property{Object: ptr8729697920, Defaulter: false, Item: ptr8729362080, Optional: false}

	ptr8729608624 := &system.Type{Object: ptr8729697792, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"level": ptr8729500992, "preferred": ptr8729501040, "language": ptr8729501376}, Rule: (*system.Type)(nil)}

	ptr8729686400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729640960 := &system.Object{Context: ptr8729686400, Description: "This represents an image, and contains path, server and protocol separately", Id: "photo", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729686240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729641600 := &system.Object{Context: ptr8729686240, Description: "The path for the url - e.g. /foo/bar.jpg", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729686112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729641728 := &system.Object{Context: ptr8729686112, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729685792 := &system.Selector{Object: ptr8729641728, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729649472 := &system.String_rule{Selector: ptr8729685792, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "^/.*$", Exists: true}}

	ptr8729501616 := &system.Property{Object: ptr8729641600, Defaulter: false, Item: ptr8729649472, Optional: false}

	ptr8729684544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729641088 := &system.Object{Context: ptr8729684544, Description: "The protocol for the url - e.g. http or https", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729684416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729641216 := &system.Object{Context: ptr8729684416, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729684224 := &system.Selector{Object: ptr8729641216, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729649152 := &system.String_rule{Selector: ptr8729684224, Default: system.String{Value: "http", Exists: true}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729501424 := &system.Property{Object: ptr8729641088, Defaulter: false, Item: ptr8729649152, Optional: true}

	ptr8729685344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729641344 := &system.Object{Context: ptr8729685344, Description: "The server for the url - e.g. www.google.com", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729685248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729641472 := &system.Object{Context: ptr8729685248, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729685024 := &system.Selector{Object: ptr8729641472, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729649312 := &system.String_rule{Selector: ptr8729685024, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729501520 := &system.Property{Object: ptr8729641344, Defaulter: false, Item: ptr8729649312, Optional: false}

	ptr8729609056 := &system.Type{Object: ptr8729640960, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/jsonselect:image", Package: "kego.io/jsonselect", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"path": ptr8729501616, "protocol": ptr8729501424, "server": ptr8729501520}, Rule: (*system.Type)(nil)}

	ptr8729140576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729692544 := &system.Object{Context: ptr8729140576, Description: "", Id: "c", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729689728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729692672 := &system.Object{Context: ptr8729689728, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729689472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729692800 := &system.Object{Context: ptr8729689472, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729689152 := &system.Selector{Object: ptr8729692800, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729625136 := &system.Number_rule{Selector: ptr8729689152, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729499504 := &system.Property{Object: ptr8729692672, Defaulter: false, Item: ptr8729625136, Optional: false}

	ptr8729139456 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729692928 := &system.Object{Context: ptr8729139456, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729139168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729693056 := &system.Object{Context: ptr8729139168, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729690048 := &system.Selector{Object: ptr8729693056, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729625296 := &system.Number_rule{Selector: ptr8729690048, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729499552 := &system.Property{Object: ptr8729692928, Defaulter: false, Item: ptr8729625296, Optional: false}

	ptr8729140448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729693184 := &system.Object{Context: ptr8729140448, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729140352 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729693312 := &system.Object{Context: ptr8729140352, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729139776 := &system.Selector{Object: ptr8729693312, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729140128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729693440 := &system.Object{Context: ptr8729140128, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729140000 := &system.Selector{Object: ptr8729693440, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729625376 := &system.Number_rule{Selector: ptr8729140000, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729633792 := &system.Map_rule{Selector: ptr8729139776, Items: ptr8729625376, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729499648 := &system.Property{Object: ptr8729693184, Defaulter: false, Item: ptr8729633792, Optional: false}

	ptr8729384576 := &system.Type{Object: ptr8729692544, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8729499504, "b": ptr8729499552, "c": ptr8729499648}, Rule: (*system.Type)(nil)}

	ptr8729820608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729697536 := &system.Object{Context: ptr8729820608, Description: "", Id: "image", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729608192 := &system.Type{Object: ptr8729697536, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729818400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729696768 := &system.Object{Context: ptr8729818400, Description: "Automatically created basic rule for expr", Id: "@expr", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729384432 := &system.Type{Object: ptr8729696768, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729683072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729371904 := &system.Object{Context: ptr8729683072, Description: "Automatically created basic rule for kid", Id: "@kid", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729608768 := &system.Type{Object: ptr8729371904, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729814336 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729694208 := &system.Object{Context: ptr8729814336, Description: "Automatically created basic rule for collision", Id: "@collision", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729385152 := &system.Type{Object: ptr8729694208, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729686912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729641856 := &system.Object{Context: ptr8729686912, Description: "Automatically created basic rule for photo", Id: "@photo", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729609200 := &system.Type{Object: ptr8729641856, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729140992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729693568 := &system.Object{Context: ptr8729140992, Description: "Automatically created basic rule for c", Id: "@c", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729384720 := &system.Type{Object: ptr8729693568, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729815744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729694720 := &system.Object{Context: ptr8729815744, Description: "Automatically created basic rule for diagram", Id: "@diagram", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729386304 := &system.Type{Object: ptr8729694720, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729815328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729694336 := &system.Object{Context: ptr8729815328, Description: "", Id: "diagram", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729815200 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729694464 := &system.Object{Context: ptr8729815200, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729815104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729694592 := &system.Object{Context: ptr8729815104, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729814976 := &system.Selector{Object: ptr8729694592, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729361440 := &system.String_rule{Selector: ptr8729814976, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729500368 := &system.Property{Object: ptr8729694464, Defaulter: false, Item: ptr8729361440, Optional: false}

	ptr8729386160 := &system.Type{Object: ptr8729694336, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/jsonselect:image", Package: "kego.io/jsonselect", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"url": ptr8729500368}, Rule: (*system.Type)(nil)}

	ptr8729821024 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729697664 := &system.Object{Context: ptr8729821024, Description: "Automatically created basic rule for image", Id: "@image", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729608336 := &system.Type{Object: ptr8729697664, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729819904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729696896 := &system.Object{Context: ptr8729819904, Description: "This represents a gallery - it's just a list of images", Id: "gallery", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729819776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729697024 := &system.Object{Context: ptr8729819776, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729819680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729697152 := &system.Object{Context: ptr8729819680, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729818976 := &system.Selector{Object: ptr8729697152, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729819584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729697280 := &system.Object{Context: ptr8729819584, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@image", Package: "kego.io/jsonselect", Type: "@image", Exists: true}}

	ptr8729819488 := &system.Selector{Object: ptr8729697280, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729224632 := &jsonselect.Image_rule{Selector: ptr8729819488}

	ptr8729633472 := &system.Map_rule{Selector: ptr8729818976, Items: ptr8729224632, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729498784 := &system.Property{Object: ptr8729697024, Defaulter: false, Item: ptr8729633472, Optional: false}

	ptr8729385296 := &system.Type{Object: ptr8729696896, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"images": ptr8729498784}, Rule: (*system.Type)(nil)}

	ptr8729864224 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729644288 := &system.Object{Context: ptr8729864224, Description: "Automatically created basic rule for sibling", Id: "@sibling", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729608480 := &system.Type{Object: ptr8729644288, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729688320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729692416 := &system.Object{Context: ptr8729688320, Description: "Automatically created basic rule for basic", Id: "@basic", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729384144 := &system.Type{Object: ptr8729692416, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729820320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729697408 := &system.Object{Context: ptr8729820320, Description: "Automatically created basic rule for gallery", Id: "@gallery", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729385872 := &system.Type{Object: ptr8729697408, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729869760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729646464 := &system.Object{Context: ptr8729869760, Description: "Automatically created basic rule for typed", Id: "@typed", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729610064 := &system.Type{Object: ptr8729646464, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729689952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729641984 := &system.Object{Context: ptr8729689952, Description: "", Id: "polykids", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729689792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729642112 := &system.Object{Context: ptr8729689792, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729689600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729642240 := &system.Object{Context: ptr8729689600, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729687808 := &system.Selector{Object: ptr8729642240, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729689216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729642368 := &system.Object{Context: ptr8729689216, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@kid", Package: "kego.io/jsonselect", Type: "@kid", Exists: true}}

	ptr8729689088 := &system.Selector{Object: ptr8729642368, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729225328 := &jsonselect.Kid_rule{Selector: ptr8729689088}

	ptr8729634368 := &system.Array_rule{Selector: ptr8729687808, Items: ptr8729225328, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729250048 := &system.Property{Object: ptr8729642112, Defaulter: false, Item: ptr8729634368, Optional: false}

	ptr8729609488 := &system.Type{Object: ptr8729641984, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8729250048}, Rule: (*system.Type)(nil)}

	ptr8729687712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729691392 := &system.Object{Context: ptr8729687712, Description: "", Id: "basic", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729682240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729691520 := &system.Object{Context: ptr8729682240, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729682144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729691648 := &system.Object{Context: ptr8729682144, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729688960 := &system.Selector{Object: ptr8729691648, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729682048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729691776 := &system.Object{Context: ptr8729682048, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729689504 := &system.Selector{Object: ptr8729691776, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729361120 := &system.String_rule{Selector: ptr8729689504, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729633664 := &system.Map_rule{Selector: ptr8729688960, Items: ptr8729361120, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729499840 := &system.Property{Object: ptr8729691520, Defaulter: false, Item: ptr8729633664, Optional: false}

	ptr8729682784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729690112 := &system.Object{Context: ptr8729682784, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729682688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729690240 := &system.Object{Context: ptr8729682688, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729682560 := &system.Selector{Object: ptr8729690240, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729358560 := &system.String_rule{Selector: ptr8729682560, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729497440 := &system.Property{Object: ptr8729690112, Defaulter: false, Item: ptr8729358560, Optional: false}

	ptr8729684448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729690368 := &system.Object{Context: ptr8729684448, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729684288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729690496 := &system.Object{Context: ptr8729684288, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729683584 := &system.Selector{Object: ptr8729690496, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729684192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729690624 := &system.Object{Context: ptr8729684192, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729683776 := &system.Selector{Object: ptr8729690624, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729684096 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729690752 := &system.Object{Context: ptr8729684096, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729683968 := &system.Selector{Object: ptr8729690752, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729359200 := &system.String_rule{Selector: ptr8729683968, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729633024 := &system.Map_rule{Selector: ptr8729683776, Items: ptr8729359200, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729632832 := &system.Array_rule{Selector: ptr8729683584, Items: ptr8729633024, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729497536 := &system.Property{Object: ptr8729690368, Defaulter: false, Item: ptr8729632832, Optional: false}

	ptr8729685376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729690880 := &system.Object{Context: ptr8729685376, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729685216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729691008 := &system.Object{Context: ptr8729685216, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729684768 := &system.Selector{Object: ptr8729691008, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729685120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729691136 := &system.Object{Context: ptr8729685120, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729684992 := &system.Selector{Object: ptr8729691136, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729359520 := &system.String_rule{Selector: ptr8729684992, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729633088 := &system.Array_rule{Selector: ptr8729684768, Items: ptr8729359520, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729498256 := &system.Property{Object: ptr8729690880, Defaulter: false, Item: ptr8729633088, Optional: false}

	ptr8729686304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729691264 := &system.Object{Context: ptr8729686304, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729686144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729691904 := &system.Object{Context: ptr8729686144, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729685728 := &system.Selector{Object: ptr8729691904, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729686048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729692032 := &system.Object{Context: ptr8729686048, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729685920 := &system.Selector{Object: ptr8729692032, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729360160 := &system.String_rule{Selector: ptr8729685920, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729633216 := &system.Array_rule{Selector: ptr8729685728, Items: ptr8729360160, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729498352 := &system.Property{Object: ptr8729691264, Defaulter: false, Item: ptr8729633216, Optional: false}

	ptr8729687584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729692160 := &system.Object{Context: ptr8729687584, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729687360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729692288 := &system.Object{Context: ptr8729687360, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729687232 := &system.Selector{Object: ptr8729692288, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729624656 := &system.Number_rule{Selector: ptr8729687232, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729498496 := &system.Property{Object: ptr8729692160, Defaulter: false, Item: ptr8729624656, Optional: false}

	ptr8729386016 := &system.Type{Object: ptr8729691392, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"name": ptr8729499840, "favoriteColor": ptr8729497440, "languagesSpoken": ptr8729497536, "seatingPreference": ptr8729498256, "drinkPreference": ptr8729498352, "weight": ptr8729498496}, Rule: (*system.Type)(nil)}

	ptr8729813920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729693696 := &system.Object{Context: ptr8729813920, Description: "", Id: "collision", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729813792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729693824 := &system.Object{Context: ptr8729813792, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729813696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729693952 := &system.Object{Context: ptr8729813696, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729813280 := &system.Selector{Object: ptr8729693952, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729813600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729694080 := &system.Object{Context: ptr8729813600, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729813472 := &system.Selector{Object: ptr8729694080, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729361280 := &system.String_rule{Selector: ptr8729813472, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729633920 := &system.Map_rule{Selector: ptr8729813280, Items: ptr8729361280, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729500080 := &system.Property{Object: ptr8729693824, Defaulter: false, Item: ptr8729633920, Optional: false}

	ptr8729385008 := &system.Type{Object: ptr8729693696, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"number": ptr8729500080}, Rule: (*system.Type)(nil)}

	ptr8729862432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729642496 := &system.Object{Context: ptr8729862432, Description: "Automatically created basic rule for polykids", Id: "@polykids", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729609632 := &system.Type{Object: ptr8729642496, Exclude: false, Extends: system.Reference{Value: "kego.io/system:selector", Package: "kego.io/system", Type: "selector", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729869344 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729644416 := &system.Object{Context: ptr8729869344, Description: "", Id: "typed", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729865856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729644544 := &system.Object{Context: ptr8729865856, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729865728 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729644672 := &system.Object{Context: ptr8729865728, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729865120 := &system.Selector{Object: ptr8729644672, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729865600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729644800 := &system.Object{Context: ptr8729865600, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729865408 := &system.Selector{Object: ptr8729644800, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729649632 := &system.String_rule{Selector: ptr8729865408, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729634240 := &system.Map_rule{Selector: ptr8729865120, Items: ptr8729649632, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729879392 := &system.Property{Object: ptr8729644544, Defaulter: false, Item: ptr8729634240, Optional: false}

	ptr8729866400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729644928 := &system.Object{Context: ptr8729866400, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729866304 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729645056 := &system.Object{Context: ptr8729866304, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729866176 := &system.Selector{Object: ptr8729645056, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729649792 := &system.String_rule{Selector: ptr8729866176, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729879488 := &system.Property{Object: ptr8729644928, Defaulter: false, Item: ptr8729649792, Optional: false}

	ptr8729867168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729645184 := &system.Object{Context: ptr8729867168, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729867072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729645312 := &system.Object{Context: ptr8729867072, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729866720 := &system.Selector{Object: ptr8729645312, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729866976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729645440 := &system.Object{Context: ptr8729866976, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@kid", Package: "kego.io/jsonselect", Type: "@kid", Exists: true}}

	ptr8729866880 := &system.Selector{Object: ptr8729645440, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729224264 := &jsonselect.Kid_rule{Selector: ptr8729866880}

	ptr8729634560 := &system.Map_rule{Selector: ptr8729866720, Items: ptr8729224264, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729879536 := &system.Property{Object: ptr8729645184, Defaulter: false, Item: ptr8729634560, Optional: false}

	ptr8729867648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729645568 := &system.Object{Context: ptr8729867648, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729867552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729645696 := &system.Object{Context: ptr8729867552, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@image", Package: "kego.io/jsonselect", Type: "@image", Exists: true}}

	ptr8729867456 := &system.Selector{Object: ptr8729645696, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729224624 := &jsonselect.Image_rule{Selector: ptr8729867456}

	ptr8729879632 := &system.Property{Object: ptr8729645568, Defaulter: false, Item: ptr8729224624, Optional: false}

	ptr8729868544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729645824 := &system.Object{Context: ptr8729868544, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729868384 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729645952 := &system.Object{Context: ptr8729868384, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729867968 := &system.Selector{Object: ptr8729645952, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729868288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729646080 := &system.Object{Context: ptr8729868288, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729868160 := &system.Selector{Object: ptr8729646080, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729649952 := &system.String_rule{Selector: ptr8729868160, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729634880 := &system.Array_rule{Selector: ptr8729867968, Items: ptr8729649952, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729879680 := &system.Property{Object: ptr8729645824, Defaulter: false, Item: ptr8729634880, Optional: false}

	ptr8729869216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729646208 := &system.Object{Context: ptr8729869216, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729868992 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729646336 := &system.Object{Context: ptr8729868992, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729868864 := &system.Selector{Object: ptr8729646336, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729625456 := &system.Number_rule{Selector: ptr8729868864, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729879776 := &system.Property{Object: ptr8729646208, Defaulter: false, Item: ptr8729625456, Optional: false}

	ptr8729609344 := &system.Type{Object: ptr8729644416, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"name": ptr8729879392, "favoriteColor": ptr8729879488, "kids": ptr8729879536, "avatar": ptr8729879632, "drinkPreference": ptr8729879680, "weight": ptr8729879776}, Rule: (*system.Type)(nil)}

	ptr8729863680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729642624 := &system.Object{Context: ptr8729863680, Description: "", Id: "sibling", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729863328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729642752 := &system.Object{Context: ptr8729863328, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729863104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729642880 := &system.Object{Context: ptr8729863104, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729862976 := &system.Selector{Object: ptr8729642880, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729628016 := &system.Number_rule{Selector: ptr8729862976, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729878720 := &system.Property{Object: ptr8729642752, Defaulter: false, Item: ptr8729628016, Optional: false}

	ptr8729864000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729643008 := &system.Object{Context: ptr8729864000, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729863776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729643136 := &system.Object{Context: ptr8729863776, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729863648 := &system.Selector{Object: ptr8729643136, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729628096 := &system.Number_rule{Selector: ptr8729863648, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729878768 := &system.Property{Object: ptr8729643008, Defaulter: false, Item: ptr8729628096, Optional: false}

	ptr8729864832 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729643264 := &system.Object{Context: ptr8729864832, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729864736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729643392 := &system.Object{Context: ptr8729864736, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@c", Package: "kego.io/jsonselect", Type: "@c", Exists: true}}

	ptr8729864640 := &system.Selector{Object: ptr8729643392, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729225560 := &jsonselect.C_rule{Selector: ptr8729864640}

	ptr8729878816 := &system.Property{Object: ptr8729643264, Defaulter: false, Item: ptr8729225560, Optional: false}

	ptr8729865792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729643520 := &system.Object{Context: ptr8729865792, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729865696 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729643648 := &system.Object{Context: ptr8729865696, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729865152 := &system.Selector{Object: ptr8729643648, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729865472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729643776 := &system.Object{Context: ptr8729865472, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729865344 := &system.Selector{Object: ptr8729643776, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729628176 := &system.Number_rule{Selector: ptr8729865344, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729634816 := &system.Map_rule{Selector: ptr8729865152, Items: ptr8729628176, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729879248 := &system.Property{Object: ptr8729643520, Defaulter: false, Item: ptr8729634816, Optional: false}

	ptr8729863488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729643904 := &system.Object{Context: ptr8729863488, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729863392 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729644032 := &system.Object{Context: ptr8729863392, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729862656 := &system.Selector{Object: ptr8729644032, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729863136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729644160 := &system.Object{Context: ptr8729863136, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729862912 := &system.Selector{Object: ptr8729644160, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729624576 := &system.Number_rule{Selector: ptr8729862912, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729632896 := &system.Map_rule{Selector: ptr8729862656, Items: ptr8729624576, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729879344 := &system.Property{Object: ptr8729643904, Defaulter: false, Item: ptr8729632896, Optional: false}

	ptr8729609920 := &system.Type{Object: ptr8729642624, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8729878720, "b": ptr8729878768, "c": ptr8729878816, "d": ptr8729879248, "e": ptr8729879344}, Rule: (*system.Type)(nil)}

	ptr8729817984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729694848 := &system.Object{Context: ptr8729817984, Description: "", Id: "expr", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729814464 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729695488 := &system.Object{Context: ptr8729814464, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729814176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729695616 := &system.Object{Context: ptr8729814176, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729813984 := &system.Selector{Object: ptr8729695616, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729358880 := &system.String_rule{Selector: ptr8729813984, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729497584 := &system.Property{Object: ptr8729695488, Defaulter: false, Item: ptr8729358880, Optional: false}

	ptr8729815232 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729695744 := &system.Object{Context: ptr8729815232, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729815040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729695872 := &system.Object{Context: ptr8729815040, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729814880 := &system.Selector{Object: ptr8729695872, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729360480 := &system.String_rule{Selector: ptr8729814880, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729497632 := &system.Property{Object: ptr8729695744, Defaulter: false, Item: ptr8729360480, Optional: false}

	ptr8729816192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729696000 := &system.Object{Context: ptr8729816192, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729816064 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729696128 := &system.Object{Context: ptr8729816064, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729815904 := &system.Selector{Object: ptr8729696128, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729361600 := &system.String_rule{Selector: ptr8729815904, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729497680 := &system.Property{Object: ptr8729696000, Defaulter: false, Item: ptr8729361600, Optional: false}

	ptr8729817312 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729696256 := &system.Object{Context: ptr8729817312, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729817216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729696384 := &system.Object{Context: ptr8729817216, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729817088 := &system.Selector{Object: ptr8729696384, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729773936 := &system.Bool_rule{Selector: ptr8729817088, Default: system.Bool{Value: false, Exists: false}}

	ptr8729497728 := &system.Property{Object: ptr8729696256, Defaulter: false, Item: ptr8729773936, Optional: false}

	ptr8729817856 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729696512 := &system.Object{Context: ptr8729817856, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729817760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729696640 := &system.Object{Context: ptr8729817760, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729817632 := &system.Selector{Object: ptr8729696640, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729774528 := &system.Bool_rule{Selector: ptr8729817632, Default: system.Bool{Value: false, Exists: false}}

	ptr8729498448 := &system.Property{Object: ptr8729696512, Defaulter: false, Item: ptr8729774528, Optional: false}

	ptr8729816640 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729694976 := &system.Object{Context: ptr8729816640, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729816416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729695104 := &system.Object{Context: ptr8729816416, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729816288 := &system.Selector{Object: ptr8729695104, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729626816 := &system.Number_rule{Selector: ptr8729816288, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729500704 := &system.Property{Object: ptr8729694976, Defaulter: false, Item: ptr8729626816, Optional: false}

	ptr8729813504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729695232 := &system.Object{Context: ptr8729813504, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729813184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729695360 := &system.Object{Context: ptr8729813184, Description: "", Id: "", Rules: []system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729816960 := &system.Selector{Object: ptr8729695360, Selector: system.String{Value: ":root", Exists: true}}

	ptr8729626896 := &system.Number_rule{Selector: ptr8729816960, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729500752 := &system.Property{Object: ptr8729695232, Defaulter: false, Item: ptr8729626896, Optional: false}

	ptr8729386592 := &system.Type{Object: ptr8729694848, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"string": ptr8729497584, "string2": ptr8729497632, "null": ptr8729497680, "true": ptr8729497728, "false": ptr8729498448, "int": ptr8729500704, "float": ptr8729500752}, Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/jsonselect:@basic", ptr8729384144)

	system.RegisterType("kego.io/jsonselect:@c", ptr8729384720)

	system.RegisterType("kego.io/jsonselect:@collision", ptr8729385152)

	system.RegisterType("kego.io/jsonselect:@diagram", ptr8729386304)

	system.RegisterType("kego.io/jsonselect:@expr", ptr8729384432)

	system.RegisterType("kego.io/jsonselect:@gallery", ptr8729385872)

	system.RegisterType("kego.io/jsonselect:@image", ptr8729608336)

	system.RegisterType("kego.io/jsonselect:@kid", ptr8729608768)

	system.RegisterType("kego.io/jsonselect:@photo", ptr8729609200)

	system.RegisterType("kego.io/jsonselect:@polykids", ptr8729609632)

	system.RegisterType("kego.io/jsonselect:@sibling", ptr8729608480)

	system.RegisterType("kego.io/jsonselect:@typed", ptr8729610064)

	system.RegisterType("kego.io/jsonselect:basic", ptr8729386016)

	system.RegisterType("kego.io/jsonselect:c", ptr8729384576)

	system.RegisterType("kego.io/jsonselect:collision", ptr8729385008)

	system.RegisterType("kego.io/jsonselect:diagram", ptr8729386160)

	system.RegisterType("kego.io/jsonselect:expr", ptr8729386592)

	system.RegisterType("kego.io/jsonselect:gallery", ptr8729385296)

	system.RegisterType("kego.io/jsonselect:image", ptr8729608192)

	system.RegisterType("kego.io/jsonselect:kid", ptr8729608624)

	system.RegisterType("kego.io/jsonselect:photo", ptr8729609056)

	system.RegisterType("kego.io/jsonselect:polykids", ptr8729609488)

	system.RegisterType("kego.io/jsonselect:sibling", ptr8729609920)

	system.RegisterType("kego.io/jsonselect:typed", ptr8729609344)

}
