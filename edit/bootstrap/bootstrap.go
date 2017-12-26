package bootstrap

import (
	"fmt"
	"go/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"bytes"

	"os"

	"encoding/gob"
	"io"

	"frizz.io/edit/common"
	"frizz.io/edit/wcache"
	"github.com/dave/jennifer/jen"
	"github.com/gopherjs/gopherjs/compiler"
	"honnef.co/go/js/dom"
)

type Bootstrap struct {
	Doc     dom.HTMLDocument
	Body    *dom.HTMLBodyElement
	Path    string
	RawInfo string
	Info    *common.RequestInfo
	Cache   *wcache.Cache

	Source   map[string]common.Bundle
	Archives map[string]*compiler.Archive
	Packages map[string]*types.Package
}

func (b *Bootstrap) Start() error {

	if err := b.Init(); err != nil {
		return err
	}

	if err := b.Compile(); err != nil {
		return err
	}

	//for name, pkg := range b.Source {
	//	for fname, contents := range pkg {
	//		b.Doc.Underlying().Call("write", fmt.Sprintf("%s - %s - %d<br>", name, fname, len(contents)))
	//	}
	//}

	return nil
}

func (b *Bootstrap) Init() error {

	csh, err := wcache.New("v1")
	if err != nil {
		return err
	}
	b.Cache = csh
	b.Doc = dom.GetWindow().Document().(dom.HTMLDocument)
	b.Body = b.Doc.GetElementByID("body").(*dom.HTMLBodyElement)
	b.Path = strings.TrimPrefix(strings.TrimSuffix(dom.GetWindow().Location().Pathname, "/"), "/")

	b.RawInfo = b.Body.GetAttribute("info")

	info, err := common.NewRequestInfo(b.RawInfo)
	if err != nil {
		return err
	}
	b.Info = info

	b.Source = make(map[string]common.Bundle)
	b.Archives = make(map[string]*compiler.Archive)
	b.Packages = make(map[string]*types.Package)

	response, err := get("/blob/" + b.Path + ".bin?info=" + url.QueryEscape(b.RawInfo))
	if err != nil {
		return err
	}

	blob := &common.Blob{}

	buf := bytes.NewBuffer(response)
	if err := gob.NewDecoder(buf).Decode(blob); err != nil {
		return err
	}

	// Get the standard library package archives from the server in parallel
	fmt.Println("Loading code")
	archives, source, err := b.standard(blob.Lib, blob.Source)
	if err != nil {
		return err
	}

	b.Source = source

	fmt.Println("Importing standard library")
	for path, raw := range archives {
		archive, err := compiler.ReadArchive(path+".a", path, bytes.NewReader(raw), b.Packages)
		if err != nil {
			return err
		}
		b.Archives[path] = archive
	}

	return nil
}

func (b *Bootstrap) standard(std []string, local map[string]uint64) (map[string][]byte, map[string]common.Bundle, error) {
	parallel := NewParallel(50)

	sourceM := &sync.Mutex{}
	source := map[string]common.Bundle{}
	for path, hash := range local {
		path := path
		hash := hash
		parallel.In <- func() error {
			body, err := getReader(fmt.Sprintf("/src/%s.bin?hash=%d", path, hash))
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
			response, err := get(fmt.Sprintf("/static/pkg/%s.a", path))
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

func get(url string) ([]byte, error) {
	reader, err := getReader(url)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func getReader(url string) (io.Reader, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 404 {
		return nil, os.ErrNotExist
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("error %s", response.Status)
	}
	return response.Body, nil
}

func (b *Bootstrap) GenStub() ([]byte, error) {
	f := jen.NewFile("main")
	/*
		func main() {
			editor.Edit(<pkg>.Package)
		}
	*/
	f.Func().Id("main").Params().Block(
		jen.Qual("frizz.io/edit/editor", "Edit").Call(jen.Qual(b.Path, "Package")),
	)
	buf := &bytes.Buffer{}
	if err := f.Render(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

}
