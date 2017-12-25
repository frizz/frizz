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

	Data     map[string][]byte
	Source   map[string]common.Package
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

	b.Data = make(map[string][]byte)
	b.Source = make(map[string]common.Package)
	b.Archives = make(map[string]*compiler.Archive)
	b.Packages = make(map[string]*types.Package)

	response, err := get("/" + b.Path + "/blob.bin?info=" + url.QueryEscape(b.RawInfo))
	if err != nil {
		return err
	}
	blob, err := common.NewBlob(response)
	if err != nil {
		return err
	}

	b.Data = blob.Data
	b.Source = blob.Source

	// Get the standard library package archives from the server in parallel
	fmt.Println("Loading standard library")
	archives, err := b.standard(blob.Lib)
	if err != nil {
		return err
	}
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

func (b *Bootstrap) standard(paths []string) (map[string][]byte, error) {
	parallel := NewParallel(20)
	mutex := &sync.Mutex{}
	archives := map[string][]byte{}
	paths = append(paths, "github.com/gopherjs/gopherjs/js", "github.com/gopherjs/gopherjs/nosync")
	for _, path := range paths {
		path := path
		parallel.In <- func() error {
			response, err := get(fmt.Sprintf("/data/pkg/%s.a", path))
			if err != nil {
				if os.IsNotExist(err) {
					return nil
				}
				return err
			}
			if len(response) == 0 {
				return nil
			}
			mutex.Lock()
			archives[path] = response
			mutex.Unlock()
			return nil
		}
	}
	if err := parallel.Finish(); err != nil {
		return nil, err
	}
	return archives, nil
}

func get(url string) ([]byte, error) {
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
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
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
