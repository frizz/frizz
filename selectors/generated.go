package jsonselect

import (
	"reflect"

	"kego.io/json"

	"kego.io/system"
)

//***********************************************************
//*** @basic ***
//***********************************************************

// Automatically created basic rule for basic
type Basic_rule struct {
	*system.Object
}

//***********************************************************
//*** @c ***
//***********************************************************

// Automatically created basic rule for c
type C_rule struct {
	*system.Object
}

//***********************************************************
//*** @collision ***
//***********************************************************

// Automatically created basic rule for collision
type Collision_rule struct {
	*system.Object
}

//***********************************************************
//*** @expr ***
//***********************************************************

// Automatically created basic rule for expr
type Expr_rule struct {
	*system.Object
}

//***********************************************************
//*** @kid ***
//***********************************************************

// Automatically created basic rule for kid
type Kid_rule struct {
	*system.Object
}

//***********************************************************
//*** @polykids ***
//***********************************************************

// Automatically created basic rule for polykids
type Polykids_rule struct {
	*system.Object
}

//***********************************************************
//*** @sibling ***
//***********************************************************

// Automatically created basic rule for sibling
type Sibling_rule struct {
	*system.Object
}

//***********************************************************
//*** basic ***
//***********************************************************

type Basic struct {
	*system.Object

	DrinkPreference []system.String

	FavoriteColor system.String

	LanguagesSpoken []map[string]system.String

	Name map[string]system.String

	SeatingPreference []system.String

	Weight system.Number
}

//***********************************************************
//*** c ***
//***********************************************************

type C struct {
	*system.Object

	A system.Number

	B system.Number

	C map[string]system.Number
}

//***********************************************************
//*** collision ***
//***********************************************************

type Collision struct {
	*system.Object

	Number map[string]system.String
}

//***********************************************************
//*** expr ***
//***********************************************************

type Expr struct {
	*system.Object

	False system.Bool

	Float system.Number

	Int system.Number

	Null system.String

	String system.String

	String2 system.String

	True system.Bool
}

//***********************************************************
//*** kid ***
//***********************************************************

type Kid struct {
	*system.Object

	Language system.String

	Level system.String

	Preferred system.Bool
}

//***********************************************************
//*** polykids ***
//***********************************************************

type Polykids struct {
	*system.Object

	A []*Kid
}

//***********************************************************
//*** sibling ***
//***********************************************************

type Sibling struct {
	*system.Object

	A system.Number

	B system.Number

	C *C

	D map[string]system.Number

	E map[string]system.Number
}

func init() {

	json.RegisterType("kego.io/jsonselect:@basic", reflect.TypeOf(&Basic_rule{}))

	json.RegisterType("kego.io/jsonselect:@c", reflect.TypeOf(&C_rule{}))

	json.RegisterType("kego.io/jsonselect:@collision", reflect.TypeOf(&Collision_rule{}))

	json.RegisterType("kego.io/jsonselect:@expr", reflect.TypeOf(&Expr_rule{}))

	json.RegisterType("kego.io/jsonselect:@kid", reflect.TypeOf(&Kid_rule{}))

	json.RegisterType("kego.io/jsonselect:@polykids", reflect.TypeOf(&Polykids_rule{}))

	json.RegisterType("kego.io/jsonselect:@sibling", reflect.TypeOf(&Sibling_rule{}))

	json.RegisterType("kego.io/jsonselect:basic", reflect.TypeOf(&Basic{}))

	json.RegisterType("kego.io/jsonselect:c", reflect.TypeOf(&C{}))

	json.RegisterType("kego.io/jsonselect:collision", reflect.TypeOf(&Collision{}))

	json.RegisterType("kego.io/jsonselect:expr", reflect.TypeOf(&Expr{}))

	json.RegisterType("kego.io/jsonselect:kid", reflect.TypeOf(&Kid{}))

	json.RegisterType("kego.io/jsonselect:polykids", reflect.TypeOf(&Polykids{}))

	json.RegisterType("kego.io/jsonselect:sibling", reflect.TypeOf(&Sibling{}))

}
