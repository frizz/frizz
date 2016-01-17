package main

import (
	"fmt"

	"kego.io/kerr"
)

func foo() {
	err := fmt.Errorf("foo")
	_ = kerr.New("EQBCISEPSL", "A", "B")
	_ = kerr.Wrap("XIJAQTCQBQ", err)
}
