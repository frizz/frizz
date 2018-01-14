package bootstrap

import (
	"fmt"
	"go/types"
	"net/url"
	"sync"

	"bytes"

	"os"

	"encoding/gob"

	"frizz.io/edit/common"
	"frizz.io/edit/util"
	"github.com/dave/jennifer/jen"
	"github.com/gopherjs/gopherjs/compiler"
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/dom"
)

func Start(c *common.Client) error {

	fail := make(chan error)
	go func() {
		for err := range fail {
			c.Log.Println(err.Error())
		}
	}()

	button := c.Doc.GetElementByID("load").(*dom.HTMLButtonElement)
	textbox := c.Doc.GetElementByID("path").(*dom.HTMLInputElement)
	js.Global.Set("bootstrap", func(path string) {
		go func() {
			if err := c.Init(); err != nil {
				fail <- err
				return
			}
			b := &Bootstrap{
				Client: c,
				Path:   path,
			}
			if err := b.Start(); err != nil {
				fail <- err
				return
			}
		}()
	})
	button.AddEventListener("click", false, func(event dom.Event) {
		event.PreventDefault()
		js.Global.Call("bootstrap", textbox.GetAttribute("value"))
	})

	c.Log.Println("Ready.")

	return nil
}

type Bootstrap struct {
	*common.Client
	Path string
	Hash uint64

	Source   map[string]common.Bundle
	Archives map[string]*compiler.Archive
	Packages map[string]*types.Package

	Code string
}

func (b *Bootstrap) Start() error {

	if err := b.Init(); err != nil {
		return err
	}

	if err := b.Compile(); err != nil {
		return err
	}

	if err := b.Open(); err != nil {
		return err
	}

	return nil
}

func (b *Bootstrap) Open() error {

	js.Global.Set("$checkForDeadlock", true)
	js.Global.Call("eval", js.InternalObject(b.Code))
	/*
		win := dom.GetWindow().Open("", b.Path, "")
		doc := win.Document()
		body := doc.GetElementsByTagName("body")[0]
		scr := doc.CreateElement("script")
		body.AppendChild(scr)
		scr.SetInnerHTML(b.Code)
	*/
	return nil
}

func (b *Bootstrap) Init() error {

	b.Source = make(map[string]common.Bundle)
	b.Archives = make(map[string]*compiler.Archive)
	b.Packages = make(map[string]*types.Package)

	b.Log.Println("Loading blob:", b.Path)

	response, err := util.Get("/blob/" + b.Path + ".bin?auth=" + url.QueryEscape(b.AuthAttribute))
	if err != nil {
		return err
	}

	blob := &common.Blob{}

	buf := bytes.NewBuffer(response)
	if err := gob.NewDecoder(buf).Decode(blob); err != nil {
		return err
	}

	b.Hash = blob.Hash

	archives, source, err := b.loadSource(blob.Lib, blob.Source)
	if err != nil {
		return err
	}

	b.Source = source

	b.Log.Println("Importing archives...")
	for path, raw := range archives {
		archive, err := compiler.ReadArchive(path+".a", path, bytes.NewReader(raw), b.Packages)
		if err != nil {
			return err
		}
		b.Archives[path] = archive
	}

	return nil
}

func (b *Bootstrap) loadSource(std []string, local map[string]uint64) (map[string][]byte, map[string]common.Bundle, error) {
	parallel := NewParallel(50)

	b.Log.Printf("Loading %d source packages and %d standard library archives\n", len(local), len(std))

	sourceM := &sync.Mutex{}
	source := map[string]common.Bundle{}
	for path, hash := range local {
		path := path
		hash := hash
		parallel.In <- func() error {
			body, err := util.GetReader(fmt.Sprintf("/src/%s.bin?hash=%d", path, hash))
			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}
				return err
			}
			var pack common.Bundle
			gob.NewDecoder(body).Decode(&pack)
			sourceM.Lock()
			source[path] = pack
			sourceM.Unlock()
			return nil
		}
	}

	std = append(std, "github.com/gopherjs/gopherjs/js", "github.com/gopherjs/gopherjs/nosync")

	archivesM := &sync.Mutex{}
	archives := map[string][]byte{}
	for _, path := range std {
		path := path
		parallel.In <- func() error {
			response, err := util.Get(fmt.Sprintf("/static/pkg/%s.a", path))
			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}
				return err
			}
			if len(response) == 0 {
				return nil
			}
			archivesM.Lock()
			archives[path] = response
			archivesM.Unlock()
			return nil
		}
	}

	if err := parallel.Finish(); err != nil {
		return nil, nil, err
	}

	return archives, source, nil
}

func (b *Bootstrap) GenStub() ([]byte, error) {
	f := jen.NewFile("main")
	/*
		func main() {
			editor.Edit(<pkg>.Package, <hash>)
		}
	*/
	f.Func().Id("main").Params().Block(
		jen.Qual("frizz.io/edit/editor", "Edit").Call(jen.Qual(b.Path, "Package"), jen.Lit(b.Hash)),
	)
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}
