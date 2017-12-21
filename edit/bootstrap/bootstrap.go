package main

import (
	"fmt"

	"frizz.io/edit/common"

	"bytes"
	"encoding/base64"
	"encoding/gob"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func main() {
	if err := start(); err != nil {
		js.Global.Get("document").Call("write", err.Error())
	}
}

func start() error {
	info, err := getInfo(getRawInfo())
	if err != nil {
		return err
	}
	js.Global.Get("document").Call("write", fmt.Sprintf("%#v\n", info))
	return nil
}

func getInfo(infoBase64 string) (info *common.RequestInfo, err error) {
	info = &common.RequestInfo{}
	infoBytes, err := base64.StdEncoding.DecodeString(infoBase64)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(infoBytes)
	if err := gob.NewDecoder(buf).Decode(info); err != nil {
		return nil, err
	}
	return info, nil
}

func getRawInfo() string {
	return dom.
		GetWindow().
		Document().(dom.HTMLDocument).
		GetElementByID("body").(*dom.HTMLBodyElement).
		GetAttribute("info")
}
