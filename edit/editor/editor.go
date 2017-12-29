package editor

import (
	"fmt"

	"encoding/gob"

	"frizz.io/edit/common"
	"frizz.io/edit/editor/views"
	"frizz.io/edit/util"
	"frizz.io/global"
	"github.com/gopherjs/vecty"
	"honnef.co/go/js/dom"
)

func Edit(p global.Package, hash uint64) {
	c, err := common.NewClient()
	if err != nil {
		fmt.Println(err)
	}
	e := &Editor{Client: c, Package: p, Hash: hash}
	go func() {
		if err := e.Start(); err != nil {
			c.Log.Println(err)
		}
	}()
}

type Editor struct {
	Client  *common.Client
	Hash    uint64
	Package global.Package
	Div     *dom.HTMLDivElement
}

func (e *Editor) Start() error {
	e.Client.Log.Println("Loading data:", e.Package.Path())
	r, err := util.GetReader(fmt.Sprintf("/data/%s.bin?hash=%d", e.Package.Path(), e.Hash))
	if err != nil {
		return err
	}
	bundle := &common.Bundle{}
	if err := gob.NewDecoder(r).Decode(bundle); err != nil {
		return err
	}
	e.Client.Log.Printf("Got %d data files", len(bundle.Files))

	vecty.RenderBody(views.NewPage())

	return nil
}
