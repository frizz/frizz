package main

import (
	"fmt"

	"github.com/davelondon/kerr"
)

func foo() {
	err := fmt.Errorf("foo")
	_ = kerr.New("EQBCISEPSL", "A", "B")
	_ = kerr.Wrap("XIJAQTCQBQ", err)
}
