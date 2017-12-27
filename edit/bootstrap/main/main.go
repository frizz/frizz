package main

import (
	"fmt"

	"frizz.io/edit/bootstrap"
	"frizz.io/edit/common"
)

func main() {
	c, err := common.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	if err := bootstrap.Start(c); err != nil {
		c.Log.Println(err)
	}
}
