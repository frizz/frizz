package main

import (
	"fmt"

	"kego.io/kerr"
)

func foo() {
	err := fmt.Errorf("foo")
	_ = kerr.New1("EQBCISEPSL", "A", "B")
	_ = kerr.Wrap("XIJAQTCQBQ", err)
}
