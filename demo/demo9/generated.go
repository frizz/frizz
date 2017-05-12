// info:{"Path":"frizz.io/demo/demo9","Hash":17284658946093704279}
package demo9

import (
	context "context"
	fmt "fmt"
	reflect "reflect"

	jsonctx "frizz.io/context/jsonctx"
	system "frizz.io/system"
)

// notest

// Automatically created basic rule for person
type PersonRule struct {
	*system.Object
	*system.Rule
}

func (v *PersonRule) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/demo/demo9", "@person"); err != nil {
		return err
	}
	if v.Rule == nil {
		v.Rule = new(system.Rule)
	}
	if err := v.Rule.Unpack(ctx, in, false); err != nil {
		return err
	}
	return nil
}
func (v *PersonRule) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/demo/demo9", "@person", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Rule != nil {
		ob, _, _, _, err := v.Rule.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	return m, "frizz.io/demo/demo9", "@person", system.J_OBJECT, nil
}

type Person struct {
	*system.Object
	Age  *system.Int
	Name *system.String
}
type PersonInterface interface {
	GetPerson(ctx context.Context) *Person
}

func (o *Person) GetPerson(ctx context.Context) *Person {
	return o
}
func UnpackPersonInterface(ctx context.Context, in system.Packed) (PersonInterface, error) {
	switch in.Type() {
	case system.J_MAP:
		i, err := system.UnpackUnknownType(ctx, in, true, "frizz.io/demo/demo9", "person")
		if err != nil {
			return nil, err
		}
		ob, ok := i.(PersonInterface)
		if !ok {
			return nil, fmt.Errorf("%T does not implement PersonInterface", i)
		}
		return ob, nil
	default:
		return nil, fmt.Errorf("Unsupported json type %s when unpacking into PersonInterface.", in.Type())
	}
}
func (v *Person) Unpack(ctx context.Context, in system.Packed, iface bool) error {
	if in == nil || in.Type() == system.J_NULL {
		return nil
	}
	if v.Object == nil {
		v.Object = new(system.Object)
	}
	if err := v.Object.Unpack(ctx, in, false); err != nil {
		return err
	}
	if err := v.Object.InitializeType("frizz.io/demo/demo9", "person"); err != nil {
		return err
	}
	if field, ok := in.Map()["age"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.Int)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Age = ob0
	}
	if field, ok := in.Map()["name"]; ok && field.Type() != system.J_NULL {
		ob0 := new(system.String)
		if err := ob0.Unpack(ctx, field, false); err != nil {
			return err
		}
		v.Name = ob0
	}
	return nil
}
func (v *Person) Repack(ctx context.Context) (data interface{}, typePackage string, typeName string, jsonType system.JsonType, err error) {
	if v == nil {
		return nil, "frizz.io/demo/demo9", "person", system.J_NULL, nil
	}
	m := map[string]interface{}{}
	if v.Object != nil {
		ob, _, _, _, err := v.Object.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		for key, val := range ob.(map[string]interface{}) {
			m[key] = val
		}
	}
	if v.Age != nil {
		ob0, _, _, _, err := v.Age.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["age"] = ob0
	}
	if v.Name != nil {
		ob0, _, _, _, err := v.Name.Repack(ctx)
		if err != nil {
			return nil, "", "", "", err
		}
		m["name"] = ob0
	}
	return m, "frizz.io/demo/demo9", "person", system.J_OBJECT, nil
}
func init() {
	pkg := jsonctx.InitPackage("frizz.io/demo/demo9")
	pkg.SetHash(uint64(0xefdf71bbc1b16457))
	pkg.Init("person", func() interface{} {
		return new(Person)
	}, nil, func() interface{} {
		return new(PersonRule)
	}, func() reflect.Type {
		return reflect.TypeOf((*PersonInterface)(nil)).Elem()
	})
}
