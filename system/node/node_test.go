package node

import (
	"testing"

	"context"

	"github.com/davelondon/ktest/assert"
	"github.com/davelondon/ktest/require"
	"kego.io/packer"
	"kego.io/process/parser"
	"kego.io/system"
	"kego.io/tests"
)

func TestNode_SetValueZero3(t *testing.T) {

	n := NewNode()
	err := n.setZero(context.Background(), false, true)
	assert.IsError(t, err, "NYQULBBBHO")

	n = NewNode()
	err = n.setZero(context.Background(), true, true)
	assert.IsError(t, err, "XRYLQWRNPH")

	n = NewNode()
	err = n.setZero(context.Background(), true, false)
	assert.IsError(t, err, "ABXFQOYCBA")

	n = NewNode()
	n.Type = &system.Type{Native: system.NewString("map")}
	err = n.setZero(context.Background(), false, false)
	assert.IsError(t, err, "VGKTIRMDTJ")

}

func TestNode_setCorrectTypeField(t *testing.T) {
	n := NewNode()
	err := n.setCorrectTypeField(context.Background())
	assert.IsError(t, err, "DQKGYKFQKJ")
}

func TestNode_extractType(t *testing.T) {
	r := &system.RuleWrapper{
		Struct: &system.Rule{Interface: true},
		Parent: &system.Type{Interface: true},
	}
	_, err := extractType(context.Background(), packer.Pack(nil), r)
	assert.IsError(t, err, "TDXTPGVFAK")

	ty, err := extractType(context.Background(), nil, nil)
	require.NoError(t, err)
	assert.Nil(t, ty)

	r = &system.RuleWrapper{
		Parent: &system.Type{Interface: true},
	}
	_, err = extractType(context.Background(), packer.Pack(""), r)
	assert.IsError(t, err, "DLSQRFLINL")

	r = &system.RuleWrapper{
		Struct: &system.Rule{Interface: true},
	}
	ty, err = extractType(context.Background(), nil, r)
	require.NoError(t, err)
	assert.Nil(t, ty)

	r = &system.RuleWrapper{
		Struct: &system.Rule{Interface: true},
	}
	_, err = extractType(context.Background(), packer.PackString(`[""]`), r)
	assert.IsError(t, err, "SNYLGBJYTM")

	_, err = extractType(context.Background(), packer.PackString(`{}`), nil)
	assert.IsError(t, err, "HBJVDKAKBJ")

	r = &system.RuleWrapper{
		Parent: &system.Type{Interface: true},
	}
	_, err = extractType(context.Background(), packer.PackString(`{}`), r)
	assert.IsError(t, err, "HBJVDKAKBJ")

	cb := tests.Context("a.b/c").Sempty()
	_, err = extractType(cb.Ctx(), packer.PackString(`{"type": "a"}`), nil)
	assert.IsError(t, err, "IJFMJJWVCA")

}

func TestNode_extractFields(t *testing.T) {
	cb := tests.Context("a.b/c").Sempty()

	ty := &system.Type{}
	f := map[string]*system.Field{}
	err := extractFields(cb.Ctx(), f, ty)
	assert.IsError(t, err, "YRFWOTIGFT")

	cb.Ssystem(parser.Parse)

	ty = &system.Type{Embed: []*system.Reference{system.NewReference("a.b/c", "d")}}
	err = extractFields(cb.Ctx(), f, ty)
	assert.IsError(t, err, "SLIRILCARQ")

	f = map[string]*system.Field{"a": nil}
	ty = &system.Type{Fields: map[string]system.RuleInterface{"a": nil}}
	err = extractFields(cb.Ctx(), f, ty)
	assert.IsError(t, err, "BARXPFXQNB")
}
