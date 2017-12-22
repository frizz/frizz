package main

import (
	"fmt"

	"frizz.io/edit/common"

	"strings"

	"net/http"

	"io/ioutil"

	"net/url"

	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func main() {
	if err := start(); err != nil {
		js.Global.Get("document").Call("write", err.Error())
	}
}

func start() error {
	info, err := common.NewRequestInfo(getRawInfo())
	if err != nil {
		return err
	}
	info.Id = []byte{0}
	path := strings.TrimPrefix(strings.TrimSuffix(js.Global.Get("window").Get("location").Get("pathname").String(), "/"), "/")
	blobResponse, err := http.Get("/" + path + "/blob.bin?info=" + url.QueryEscape(getRawInfo()))
	if err != nil {
		return err
	}
	blobBytes, err := ioutil.ReadAll(blobResponse.Body)
	if err != nil {
		return err
	}
	blob, err := common.NewBlob(blobBytes)
	for n, f := range blob.Files {
		js.Global.Get("document").Call("write", fmt.Sprintf("%s:%d<br>", n, len(f)))
	}

	return nil
}

func getRawInfo() string {
	return dom.
		GetWindow().
		Document().(dom.HTMLDocument).
		GetElementByID("body").(*dom.HTMLBodyElement).
		GetAttribute("info")
}
