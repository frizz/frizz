package system

import (
	"fmt"
	"testing"

	"kego.io/kerr/assert"
)

type Alerter interface {
	Alert() string
}

type Man struct {
	Name string
}

func (m Man) Alert() string { return fmt.Sprint("Hello ", m.Name) }

var _ Alerter = (*Man)(nil)

func TestPlainMethod(t *testing.T) {

	var mi interface{} = Man{"Dave"}
	m, ok := mi.(Alerter)
	assert.True(t, ok)
	a := m.Alert()
	assert.Equal(t, "Hello Dave", a)

	var mpi interface{} = &Man{"Mike"}
	mp, ok := mpi.(Alerter)
	assert.True(t, ok)
	ap := mp.Alert()
	assert.Equal(t, "Hello Mike", ap)

}

type Dog struct {
	Name string
}

func (d *Dog) Alert() string { return fmt.Sprint("Woof ", d.Name) }

var _ Alerter = (*Dog)(nil)

func TestPointerMethod(t *testing.T) {

	var di interface{} = Dog{"Spot"}
	_, ok := di.(Alerter)
	assert.False(t, ok)

	var dpi interface{} = &Dog{"Fido"}
	dp, ok := dpi.(Alerter)
	assert.True(t, ok)
	ap := dp.Alert()
	assert.Equal(t, "Woof Fido", ap)

}
