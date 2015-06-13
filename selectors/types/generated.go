package types

import (
	"kego.io/system"

	_ "kego.io/system/types"

	"kego.io/jsonselect"
)

func init() {

	ptr8729431680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728932064 := &system.Object{Context: ptr8729431680, Description: "", Id: "sibling", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729433760 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728932176 := &system.Object{Context: ptr8729433760, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729433600 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728932512 := &system.Object{Context: ptr8729433600, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729089744 := &system.Number_rule{Object: ptr8728932512, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729102256 := &system.Property{Object: ptr8728932176, Defaulter: false, Item: ptr8729089744, Optional: false}

	ptr8729428448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728932624 := &system.Object{Context: ptr8729428448, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729428256 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728931056 := &system.Object{Context: ptr8729428256, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729084704 := &system.Number_rule{Object: ptr8728931056, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729102448 := &system.Property{Object: ptr8728932624, Defaulter: false, Item: ptr8729084704, Optional: false}

	ptr8729429280 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728931392 := &system.Object{Context: ptr8729429280, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729429248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728933184 := &system.Object{Context: ptr8729429248, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@c", Package: "kego.io/jsonselect", Type: "@c", Exists: true}}

	ptr8729076312 := &jsonselect.C_rule{Object: ptr8728933184}

	ptr8729102736 := &system.Property{Object: ptr8728931392, Defaulter: false, Item: ptr8729076312, Optional: false}

	ptr8729430272 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728933408 := &system.Object{Context: ptr8729430272, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729430240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728933520 := &system.Object{Context: ptr8729430240, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729430048 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728933632 := &system.Object{Context: ptr8729430048, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729084864 := &system.Number_rule{Object: ptr8728933632, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729239616 := &system.Map_rule{Object: ptr8728933520, Items: ptr8729084864, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729103360 := &system.Property{Object: ptr8728933408, Defaulter: false, Item: ptr8729239616, Optional: false}

	ptr8729431360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728933744 := &system.Object{Context: ptr8729431360, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729431328 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728933856 := &system.Object{Context: ptr8729431328, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729431104 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728934080 := &system.Object{Context: ptr8729431104, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729084944 := &system.Number_rule{Object: ptr8728934080, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729240512 := &system.Map_rule{Object: ptr8728933856, Items: ptr8729084944, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729103504 := &system.Property{Object: ptr8728933744, Defaulter: false, Item: ptr8729240512, Optional: false}

	ptr8729481440 := &system.Type{Object: ptr8728932064, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8729102256, "b": ptr8729102448, "c": ptr8729102736, "d": ptr8729103360, "e": ptr8729103504}, Rule: (*system.Type)(nil)}

	ptr8729519136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728934416 := &system.Object{Context: ptr8729519136, Description: "", Id: "typed", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729435872 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729310800 := &system.Object{Context: ptr8729435872, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729435776 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729310912 := &system.Object{Context: ptr8729435776, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729435744 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729311024 := &system.Object{Context: ptr8729435744, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729486112 := &system.String_rule{Object: ptr8729311024, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729241664 := &system.Array_rule{Object: ptr8729310912, Items: ptr8729486112, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729326848 := &system.Property{Object: ptr8729310800, Defaulter: false, Item: ptr8729241664, Optional: false}

	ptr8729518912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729311136 := &system.Object{Context: ptr8729518912, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729518592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729311248 := &system.Object{Context: ptr8729518592, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729086464 := &system.Number_rule{Object: ptr8729311248, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729326992 := &system.Property{Object: ptr8729311136, Defaulter: false, Item: ptr8729086464, Optional: false}

	ptr8729433664 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728934640 := &system.Object{Context: ptr8729433664, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729433632 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729309680 := &system.Object{Context: ptr8729433632, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729433536 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729309904 := &system.Object{Context: ptr8729433536, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729485792 := &system.String_rule{Object: ptr8729309904, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729241408 := &system.Map_rule{Object: ptr8729309680, Items: ptr8729485792, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729326176 := &system.Property{Object: ptr8728934640, Defaulter: false, Item: ptr8729241408, Optional: false}

	ptr8729434208 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729310016 := &system.Object{Context: ptr8729434208, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729434176 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729310128 := &system.Object{Context: ptr8729434176, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729485952 := &system.String_rule{Object: ptr8729310128, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729326272 := &system.Property{Object: ptr8729310016, Defaulter: false, Item: ptr8729485952, Optional: false}

	ptr8729434784 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729310240 := &system.Object{Context: ptr8729434784, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729434752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729310352 := &system.Object{Context: ptr8729434752, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729434720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729310464 := &system.Object{Context: ptr8729434720, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@kid", Package: "kego.io/jsonselect", Type: "@kid", Exists: true}}

	ptr8729077424 := &jsonselect.Kid_rule{Object: ptr8729310464}

	ptr8729241536 := &system.Map_rule{Object: ptr8729310352, Items: ptr8729077424, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729326320 := &system.Property{Object: ptr8729310240, Defaulter: false, Item: ptr8729241536, Optional: false}

	ptr8729435168 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729310576 := &system.Object{Context: ptr8729435168, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729435136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729310688 := &system.Object{Context: ptr8729435136, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@image", Package: "kego.io/jsonselect", Type: "@image", Exists: true}}

	ptr8729077456 := &jsonselect.Image_rule{Object: ptr8729310688}

	ptr8729326704 := &system.Property{Object: ptr8729310576, Defaulter: false, Item: ptr8729077456, Optional: false}

	ptr8729478560 := &system.Type{Object: ptr8728934416, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"drinkPreference": ptr8729326848, "weight": ptr8729326992, "name": ptr8729326176, "favoriteColor": ptr8729326272, "kids": ptr8729326320, "avatar": ptr8729326704}, Rule: (*system.Type)(nil)}

	ptr8728991040 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729270400 := &system.Object{Context: ptr8728991040, Description: "This represents a gallery - it's just a list of images", Id: "gallery", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8728990880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729270512 := &system.Object{Context: ptr8728990880, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728990848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729270624 := &system.Object{Context: ptr8728990848, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728990816 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729270960 := &system.Object{Context: ptr8728990816, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@image", Package: "kego.io/jsonselect", Type: "@image", Exists: true}}

	ptr8729077448 := &jsonselect.Image_rule{Object: ptr8729270960}

	ptr8729240320 := &system.Map_rule{Object: ptr8729270624, Items: ptr8729077448, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729326800 := &system.Property{Object: ptr8729270512, Defaulter: false, Item: ptr8729240320, Optional: false}

	ptr8729479136 := &system.Type{Object: ptr8729270400, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"images": ptr8729326800}, Rule: (*system.Type)(nil)}

	ptr8728992544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729271296 := &system.Object{Context: ptr8728992544, Description: "Automatically created basic rule for image", Id: "@image", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729479712 := &system.Type{Object: ptr8729271296, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729520160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729264800 := &system.Object{Context: ptr8729520160, Description: "Automatically created basic rule for diagram", Id: "@diagram", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729478416 := &system.Type{Object: ptr8729264800, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729432544 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728930608 := &system.Object{Context: ptr8729432544, Description: "", Id: "polykids", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729432416 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728930720 := &system.Object{Context: ptr8729432416, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729432320 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728930832 := &system.Object{Context: ptr8729432320, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729432288 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728931728 := &system.Object{Context: ptr8729432288, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/jsonselect:@kid", Package: "kego.io/jsonselect", Type: "@kid", Exists: true}}

	ptr8729077824 := &jsonselect.Kid_rule{Object: ptr8728931728}

	ptr8729241088 := &system.Array_rule{Object: ptr8728930832, Items: ptr8729077824, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729329248 := &system.Property{Object: ptr8728930720, Defaulter: false, Item: ptr8729241088, Optional: false}

	ptr8729480864 := &system.Type{Object: ptr8728930608, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"a": ptr8729329248}, Rule: (*system.Type)(nil)}

	ptr8729431072 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728928368 := &system.Object{Context: ptr8729431072, Description: "This represents an image, and contains path, server and protocol separately", Id: "photo", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729429984 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728929600 := &system.Object{Context: ptr8729429984, Description: "The protocol for the url - e.g. http or https", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729429952 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728929712 := &system.Object{Context: ptr8729429952, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729214784 := &system.String_rule{Object: ptr8728929712, Default: system.String{Value: "http", Exists: true}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729328768 := &system.Property{Object: ptr8728929600, Defaulter: false, Item: ptr8729214784, Optional: true}

	ptr8729430432 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728929936 := &system.Object{Context: ptr8729430432, Description: "The server for the url - e.g. www.google.com", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729430400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728930048 := &system.Object{Context: ptr8729430400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729485312 := &system.String_rule{Object: ptr8728930048, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729328864 := &system.Property{Object: ptr8728929936, Defaulter: false, Item: ptr8729485312, Optional: false}

	ptr8729430944 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728930160 := &system.Object{Context: ptr8729430944, Description: "The path for the url - e.g. /foo/bar.jpg", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729430912 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728930272 := &system.Object{Context: ptr8729430912, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729485472 := &system.String_rule{Object: ptr8728930272, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 1, Exists: true}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "^/.*$", Exists: true}}

	ptr8729328960 := &system.Property{Object: ptr8728930160, Defaulter: false, Item: ptr8729485472, Optional: false}

	ptr8729480432 := &system.Type{Object: ptr8728928368, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/jsonselect:image", Package: "kego.io/jsonselect", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"protocol": ptr8729328768, "server": ptr8729328864, "path": ptr8729328960}, Rule: (*system.Type)(nil)}

	ptr8728992000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729271184 := &system.Object{Context: ptr8728992000, Description: "", Id: "image", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729479568 := &system.Type{Object: ptr8729271184, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: true, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729522720 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729267936 := &system.Object{Context: ptr8729522720, Description: "", Id: "c", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729521824 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729268272 := &system.Object{Context: ptr8729521824, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729521664 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729268384 := &system.Object{Context: ptr8729521664, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729086784 := &system.Number_rule{Object: ptr8729268384, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729328240 := &system.Property{Object: ptr8729268272, Defaulter: false, Item: ptr8729086784, Optional: false}

	ptr8729522592 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729268496 := &system.Object{Context: ptr8729522592, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729522560 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729268608 := &system.Object{Context: ptr8729522560, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729522400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729268720 := &system.Object{Context: ptr8729522400, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729086864 := &system.Number_rule{Object: ptr8729268720, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729240576 := &system.Map_rule{Object: ptr8729268608, Items: ptr8729086864, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729328288 := &system.Property{Object: ptr8729268496, Defaulter: false, Item: ptr8729240576, Optional: false}

	ptr8729521248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729268048 := &system.Object{Context: ptr8729521248, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729521088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729268160 := &system.Object{Context: ptr8729521088, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729086704 := &system.Number_rule{Object: ptr8729268160, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729328192 := &system.Property{Object: ptr8729268048, Defaulter: false, Item: ptr8729086704, Optional: false}

	ptr8729477696 := &system.Type{Object: ptr8729267936, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"b": ptr8729328240, "c": ptr8729328288, "a": ptr8729328192}, Rule: (*system.Type)(nil)}

	ptr8729525376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729264912 := &system.Object{Context: ptr8729525376, Description: "", Id: "expr", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729522752 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729266144 := &system.Object{Context: ptr8729522752, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729522688 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729266256 := &system.Object{Context: ptr8729522688, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729208704 := &system.String_rule{Object: ptr8729266256, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729325936 := &system.Property{Object: ptr8729266144, Defaulter: false, Item: ptr8729208704, Optional: false}

	ptr8729523520 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729267488 := &system.Object{Context: ptr8729523520, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729523488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729267600 := &system.Object{Context: ptr8729523488, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729209344 := &system.String_rule{Object: ptr8729267600, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729325984 := &system.Property{Object: ptr8729267488, Defaulter: false, Item: ptr8729209344, Optional: false}

	ptr8729524192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729269392 := &system.Object{Context: ptr8729524192, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729524160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729269504 := &system.Object{Context: ptr8729524160, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729210624 := &system.String_rule{Object: ptr8729269504, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729326032 := &system.Property{Object: ptr8729269392, Defaulter: false, Item: ptr8729210624, Optional: false}

	ptr8729524800 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729269616 := &system.Object{Context: ptr8729524800, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729524768 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729269952 := &system.Object{Context: ptr8729524768, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729588192 := &system.Bool_rule{Object: ptr8729269952, Default: system.Bool{Value: false, Exists: false}}

	ptr8729326080 := &system.Property{Object: ptr8729269616, Defaulter: false, Item: ptr8729588192, Optional: false}

	ptr8729525248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729270064 := &system.Object{Context: ptr8729525248, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729525216 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729270176 := &system.Object{Context: ptr8729525216, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729588720 := &system.Bool_rule{Object: ptr8729270176, Default: system.Bool{Value: false, Exists: false}}

	ptr8729326416 := &system.Property{Object: ptr8729270064, Defaulter: false, Item: ptr8729588720, Optional: false}

	ptr8729521376 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729265136 := &system.Object{Context: ptr8729521376, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729521184 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729265248 := &system.Object{Context: ptr8729521184, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729086624 := &system.Number_rule{Object: ptr8729265248, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729325840 := &system.Property{Object: ptr8729265136, Defaulter: false, Item: ptr8729086624, Optional: false}

	ptr8729522112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729265472 := &system.Object{Context: ptr8729522112, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729521920 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729265584 := &system.Object{Context: ptr8729521920, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729086944 := &system.Number_rule{Object: ptr8729265584, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729325888 := &system.Property{Object: ptr8729265472, Defaulter: false, Item: ptr8729086944, Optional: false}

	ptr8729478704 := &system.Type{Object: ptr8729264912, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"string": ptr8729325936, "string2": ptr8729325984, "null": ptr8729326032, "true": ptr8729326080, "false": ptr8729326416, "int": ptr8729325840, "float": ptr8729325888}, Rule: (*system.Type)(nil)}

	ptr8729524128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729268944 := &system.Object{Context: ptr8729524128, Description: "", Id: "collision", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729524000 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729269056 := &system.Object{Context: ptr8729524000, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729523968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729269168 := &system.Object{Context: ptr8729523968, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8729523936 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729269280 := &system.Object{Context: ptr8729523936, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729214304 := &system.String_rule{Object: ptr8729269280, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729240768 := &system.Map_rule{Object: ptr8729269168, Items: ptr8729214304, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729328672 := &system.Property{Object: ptr8729269056, Defaulter: false, Item: ptr8729240768, Optional: false}

	ptr8729478128 := &system.Type{Object: ptr8729268944, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"number": ptr8729328672}, Rule: (*system.Type)(nil)}

	ptr8729520032 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729264576 := &system.Object{Context: ptr8729520032, Description: "", Id: "basic", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729518976 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729267040 := &system.Object{Context: ptr8729518976, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729518880 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729267152 := &system.Object{Context: ptr8729518880, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729518848 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729267264 := &system.Object{Context: ptr8729518848, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729212864 := &system.String_rule{Object: ptr8729267264, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729240192 := &system.Array_rule{Object: ptr8729267152, Items: ptr8729212864, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729327280 := &system.Property{Object: ptr8729267040, Defaulter: false, Item: ptr8729240192, Optional: false}

	ptr8729519904 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729267376 := &system.Object{Context: ptr8729519904, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729519712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729267712 := &system.Object{Context: ptr8729519712, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@number", Package: "kego.io/system", Type: "@number", Exists: true}}

	ptr8729085424 := &system.Number_rule{Object: ptr8729267712, Default: system.Number{Value: 0, Exists: false}, ExclusiveMaximum: system.Bool{Value: false, Exists: true}, ExclusiveMinimum: system.Bool{Value: false, Exists: true}, Maximum: system.Number{Value: 0, Exists: false}, Minimum: system.Number{Value: 0, Exists: false}, MultipleOf: system.Number{Value: 0, Exists: false}}

	ptr8729327424 := &system.Property{Object: ptr8729267376, Defaulter: false, Item: ptr8729085424, Optional: false}

	ptr8728991712 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729265024 := &system.Object{Context: ptr8728991712, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728991680 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729265360 := &system.Object{Context: ptr8728991680, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728991648 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729265696 := &system.Object{Context: ptr8728991648, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729209984 := &system.String_rule{Object: ptr8729265696, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729239744 := &system.Map_rule{Object: ptr8729265360, Items: ptr8729209984, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729325648 := &system.Property{Object: ptr8729265024, Defaulter: false, Item: ptr8729239744, Optional: false}

	ptr8728992160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729265808 := &system.Object{Context: ptr8728992160, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728992128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729265920 := &system.Object{Context: ptr8728992128, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729211264 := &system.String_rule{Object: ptr8729265920, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729326560 := &system.Property{Object: ptr8729265808, Defaulter: false, Item: ptr8729211264, Optional: false}

	ptr8728993248 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729266032 := &system.Object{Context: ptr8728993248, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728993152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729266368 := &system.Object{Context: ptr8728993152, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8728993120 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729266480 := &system.Object{Context: ptr8728993120, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@map", Package: "kego.io/system", Type: "@map", Exists: true}}

	ptr8728993088 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729266592 := &system.Object{Context: ptr8728993088, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729211904 := &system.String_rule{Object: ptr8729266592, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729240000 := &system.Map_rule{Object: ptr8729266480, Items: ptr8729211904, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}}

	ptr8729239872 := &system.Array_rule{Object: ptr8729266368, Items: ptr8729240000, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729326608 := &system.Property{Object: ptr8729266032, Defaulter: false, Item: ptr8729239872, Optional: false}

	ptr8729518240 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729266704 := &system.Object{Context: ptr8729518240, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729518144 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729266816 := &system.Object{Context: ptr8729518144, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@array", Package: "kego.io/system", Type: "@array", Exists: true}}

	ptr8729518112 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729266928 := &system.Object{Context: ptr8729518112, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729212224 := &system.String_rule{Object: ptr8729266928, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729240064 := &system.Array_rule{Object: ptr8729266816, Items: ptr8729212224, MaxItems: system.Int{Value: 0, Exists: false}, MinItems: system.Int{Value: 0, Exists: false}, UniqueItems: system.Bool{Value: false, Exists: true}}

	ptr8729327184 := &system.Property{Object: ptr8729266704, Defaulter: false, Item: ptr8729240064, Optional: false}

	ptr8729477264 := &system.Type{Object: ptr8729264576, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"drinkPreference": ptr8729327280, "weight": ptr8729327424, "name": ptr8729325648, "favoriteColor": ptr8729326560, "languagesSpoken": ptr8729326608, "seatingPreference": ptr8729327184}, Rule: (*system.Type)(nil)}

	ptr8729519968 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729311360 := &system.Object{Context: ptr8729519968, Description: "Automatically created basic rule for typed", Id: "@typed", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729478992 := &system.Type{Object: ptr8729311360, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729432192 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728934192 := &system.Object{Context: ptr8729432192, Description: "Automatically created basic rule for sibling", Id: "@sibling", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729477552 := &system.Type{Object: ptr8728934192, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729431488 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728930496 := &system.Object{Context: ptr8729431488, Description: "Automatically created basic rule for photo", Id: "@photo", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729480576 := &system.Type{Object: ptr8728930496, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729518400 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729264128 := &system.Object{Context: ptr8729518400, Description: "Automatically created basic rule for collision", Id: "@collision", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729477120 := &system.Type{Object: ptr8729264128, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729432960 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8728931840 := &system.Object{Context: ptr8729432960, Description: "Automatically created basic rule for polykids", Id: "@polykids", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729481008 := &system.Type{Object: ptr8728931840, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729525792 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729270288 := &system.Object{Context: ptr8729525792, Description: "Automatically created basic rule for expr", Id: "@expr", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729478848 := &system.Type{Object: ptr8729270288, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729519584 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729264240 := &system.Object{Context: ptr8729519584, Description: "", Id: "diagram", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729519424 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729264352 := &system.Object{Context: ptr8729519424, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729519360 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729264464 := &system.Object{Context: ptr8729519360, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729208064 := &system.String_rule{Object: ptr8729264464, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729325312 := &system.Property{Object: ptr8729264352, Defaulter: false, Item: ptr8729208064, Optional: false}

	ptr8729478272 := &system.Type{Object: ptr8729264240, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/jsonselect:image", Package: "kego.io/jsonselect", Type: "image", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"url": ptr8729325312}, Rule: (*system.Type)(nil)}

	ptr8729428736 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729271408 := &system.Object{Context: ptr8729428736, Description: "", Id: "kid", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729428160 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729271744 := &system.Object{Context: ptr8729428160, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729428128 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729271856 := &system.Object{Context: ptr8729428128, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729214624 := &system.String_rule{Object: ptr8729271856, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729327760 := &system.Property{Object: ptr8729271744, Defaulter: false, Item: ptr8729214624, Optional: false}

	ptr8729428608 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729271968 := &system.Object{Context: ptr8729428608, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8729428576 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729272080 := &system.Object{Context: ptr8729428576, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@bool", Package: "kego.io/system", Type: "@bool", Exists: true}}

	ptr8729591680 := &system.Bool_rule{Object: ptr8729272080, Default: system.Bool{Value: false, Exists: false}}

	ptr8729327808 := &system.Property{Object: ptr8729271968, Defaulter: false, Item: ptr8729591680, Optional: true}

	ptr8728993504 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729271520 := &system.Object{Context: ptr8728993504, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:property", Package: "kego.io/system", Type: "property", Exists: true}}

	ptr8728993472 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729271632 := &system.Object{Context: ptr8728993472, Description: "", Id: "", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:@string", Package: "kego.io/system", Type: "@string", Exists: true}}

	ptr8729214464 := &system.String_rule{Object: ptr8729271632, Default: system.String{Value: "", Exists: false}, Enum: []system.String(nil), Equal: system.String{Value: "", Exists: false}, Format: system.String{Value: "", Exists: false}, MaxLength: system.Int{Value: 0, Exists: false}, MinLength: system.Int{Value: 0, Exists: false}, Pattern: system.String{Value: "", Exists: false}}

	ptr8729327712 := &system.Property{Object: ptr8729271520, Defaulter: false, Item: ptr8729214464, Optional: false}

	ptr8729480000 := &system.Type{Object: ptr8729271408, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference(nil), Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property{"level": ptr8729327760, "preferred": ptr8729327808, "language": ptr8729327712}, Rule: (*system.Type)(nil)}

	ptr8729520448 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729267824 := &system.Object{Context: ptr8729520448, Description: "Automatically created basic rule for basic", Id: "@basic", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729477408 := &system.Type{Object: ptr8729267824, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8728991552 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729271072 := &system.Object{Context: ptr8728991552, Description: "Automatically created basic rule for gallery", Id: "@gallery", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729479280 := &system.Type{Object: ptr8729271072, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729429152 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729272192 := &system.Object{Context: ptr8729429152, Description: "Automatically created basic rule for kid", Id: "@kid", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729480144 := &system.Type{Object: ptr8729272192, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	ptr8729523136 := &system.Context{Imports: map[string]string{}, Package: "kego.io/jsonselect"}

	ptr8729268832 := &system.Object{Context: ptr8729523136, Description: "Automatically created basic rule for c", Id: "@c", Rules: map[string]system.Rule(nil), Type: system.Reference{Value: "kego.io/system:type", Package: "kego.io/system", Type: "type", Exists: true}}

	ptr8729477840 := &system.Type{Object: ptr8729268832, Exclude: false, Extends: system.Reference{Value: "kego.io/system:object", Package: "kego.io/system", Type: "object", Exists: true}, Interface: false, Is: []system.Reference{system.Reference{Value: "kego.io/system:rule", Package: "kego.io/system", Type: "rule", Exists: true}}, Native: system.String{Value: "object", Exists: true}, Properties: map[string]*system.Property(nil), Rule: (*system.Type)(nil)}

	system.RegisterType("kego.io/jsonselect:@basic", ptr8729477408)

	system.RegisterType("kego.io/jsonselect:@c", ptr8729477840)

	system.RegisterType("kego.io/jsonselect:@collision", ptr8729477120)

	system.RegisterType("kego.io/jsonselect:@diagram", ptr8729478416)

	system.RegisterType("kego.io/jsonselect:@expr", ptr8729478848)

	system.RegisterType("kego.io/jsonselect:@gallery", ptr8729479280)

	system.RegisterType("kego.io/jsonselect:@image", ptr8729479712)

	system.RegisterType("kego.io/jsonselect:@kid", ptr8729480144)

	system.RegisterType("kego.io/jsonselect:@photo", ptr8729480576)

	system.RegisterType("kego.io/jsonselect:@polykids", ptr8729481008)

	system.RegisterType("kego.io/jsonselect:@sibling", ptr8729477552)

	system.RegisterType("kego.io/jsonselect:@typed", ptr8729478992)

	system.RegisterType("kego.io/jsonselect:basic", ptr8729477264)

	system.RegisterType("kego.io/jsonselect:c", ptr8729477696)

	system.RegisterType("kego.io/jsonselect:collision", ptr8729478128)

	system.RegisterType("kego.io/jsonselect:diagram", ptr8729478272)

	system.RegisterType("kego.io/jsonselect:expr", ptr8729478704)

	system.RegisterType("kego.io/jsonselect:gallery", ptr8729479136)

	system.RegisterType("kego.io/jsonselect:image", ptr8729479568)

	system.RegisterType("kego.io/jsonselect:kid", ptr8729480000)

	system.RegisterType("kego.io/jsonselect:photo", ptr8729480432)

	system.RegisterType("kego.io/jsonselect:polykids", ptr8729480864)

	system.RegisterType("kego.io/jsonselect:sibling", ptr8729481440)

	system.RegisterType("kego.io/jsonselect:typed", ptr8729478560)

}
