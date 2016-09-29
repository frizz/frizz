package tests

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
}

func (p *Person) Say() string {
	return p.Name
}

type ArrAlias []*Person

type foo struct {
	p ArrAlias
}

func TestFoo(t *testing.T) {
	f := &foo{
		p: []*Person{&Person{Name: "dave"}},
	}
	fmt.Println(f.p[0].Say())

}
