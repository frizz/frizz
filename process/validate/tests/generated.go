// info:{"Path":"kego.io/process/validate/tests","Hash":6676302982221491352}
package tests

import (
	"reflect"

	"golang.org/x/net/context"
	"kego.io/context/jsonctx"
	"kego.io/system"
)

// Automatically created basic rule for a
type ARule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for b
type BRule struct {
	*system.Object
	*system.Rule
}
type CRule struct {
	*system.Object
	*system.Rule
	Fail *system.Bool `json:"fail"`
}

// Automatically created basic rule for d
type DRule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for e
type ERule struct {
	*system.Object
	*system.Rule
}

// Automatically created basic rule for f
type FRule struct {
	*system.Object
	*system.Rule
}

// A is a simple type containing a string B
type A struct {
	*system.Object
	B *system.String `json:"b"`
}
type AInterface interface {
	GetA(ctx context.Context) *A
}

func (o *A) GetA(ctx context.Context) *A {
	return o
}

// A is a type containing a string interface B
type B struct {
	*system.Object
	C system.StringInterface `json:"c"`
}
type BInterface interface {
	GetB(ctx context.Context) *B
}

func (o *B) GetB(ctx context.Context) *B {
	return o
}

// D is a type containing a field of type C
type D struct {
	*system.Object
	A C `json:"a"`
}
type DInterface interface {
	GetD(ctx context.Context) *D
}

func (o *D) GetD(ctx context.Context) *D {
	return o
}

// E is a type containing an array of strings, and a map of strings
type E struct {
	*system.Object
	A []*system.String          `json:"a"`
	B map[string]*system.String `json:"b"`
}
type EInterface interface {
	GetE(ctx context.Context) *E
}

func (o *E) GetE(ctx context.Context) *E {
	return o
}

// F is a type with an extra rule attached to the field
type F struct {
	*system.Object
	A *A            `json:"a"`
	B []*A          `json:"b"`
	C map[string]*A `json:"c"`
}
type FInterface interface {
	GetF(ctx context.Context) *F
}

func (o *F) GetF(ctx context.Context) *F {
	return o
}
func init() {
	pkg := jsonctx.InitPackage("kego.io/process/validate/tests", 6676302982221491352)
	pkg.InitType("a", reflect.TypeOf((*A)(nil)), reflect.TypeOf((*ARule)(nil)), reflect.TypeOf((*AInterface)(nil)).Elem())
	pkg.InitType("b", reflect.TypeOf((*B)(nil)), reflect.TypeOf((*BRule)(nil)), reflect.TypeOf((*BInterface)(nil)).Elem())
	pkg.InitType("c", reflect.TypeOf((*C)(nil)).Elem(), reflect.TypeOf((*CRule)(nil)), nil)
	pkg.InitType("d", reflect.TypeOf((*D)(nil)), reflect.TypeOf((*DRule)(nil)), reflect.TypeOf((*DInterface)(nil)).Elem())
	pkg.InitType("e", reflect.TypeOf((*E)(nil)), reflect.TypeOf((*ERule)(nil)), reflect.TypeOf((*EInterface)(nil)).Elem())
	pkg.InitType("f", reflect.TypeOf((*F)(nil)), reflect.TypeOf((*FRule)(nil)), reflect.TypeOf((*FInterface)(nil)).Elem())
}
