package main

import (
	"fmt"

	"frizz.io/edit/bootstrap"
)

func main() {
	b := &bootstrap.Bootstrap{}
	if err := b.Start(); err != nil {
		fmt.Println("Error...")
		fmt.Println(err.Error())
	}
}
