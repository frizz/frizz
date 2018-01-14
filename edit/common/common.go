package common

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"

	"fmt"

	"frizz.io/edit/wcache"
	"honnef.co/go/js/dom"
)

type Auth struct {
	Id   []byte
	Hash []byte
}

func DecodeAuth(in string) (*Auth, error) {
	a := &Auth{}
	infoBytes, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(infoBytes)
	if err := gob.NewDecoder(buf).Decode(a); err != nil {
		return nil, err
	}
	return a, nil
}

func (a Auth) Encode() (string, error) {
	buf := &bytes.Buffer{}
	if err := gob.NewEncoder(buf).Encode(a); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

type Blob struct {
	Hash   uint64
	Source map[string]uint64 // package path and hash of all local (e.g. non standard library) packages needed to compile this package
	Lib    []string          // list of all standard library packages needed to compile this package
}

type Bundle struct {
	Hash  uint64
	Files map[string][]byte
}

func NewClient() (*Client, error) {
	c := &Client{}

	if err := c.Init(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) Init() error {
	csh, err := wcache.New("v1")
	if err != nil {
		return err
	}
	c.Cache = csh

	c.Doc = dom.GetWindow().Document().(dom.HTMLDocument)
	c.Body = c.Doc.GetElementByID("wrapper").(*dom.HTMLBodyElement)
	c.Log = NewLogger(c.Doc.GetElementByID("log").(*dom.HTMLSpanElement))

	c.AuthAttribute = c.Body.GetAttribute("auth")
	auth, err := DecodeAuth(c.AuthAttribute)
	if err != nil {
		return err
	}
	c.Auth = auth
	return nil
}

type Client struct {
	Doc           dom.HTMLDocument
	Body          *dom.HTMLBodyElement
	AuthAttribute string
	Auth          *Auth
	Cache         *wcache.Cache
	Log           Logger
}

func NewLogger(span *dom.HTMLSpanElement) Logger {
	return Logger{span: span}
}

type Logger struct {
	span *dom.HTMLSpanElement
}

func (l Logger) Printf(format string, a ...interface{}) (int, error) {
	l.span.SetInnerHTML(fmt.Sprintf(format, a...))
	return fmt.Printf(format, a...)
}

func (l Logger) Println(a ...interface{}) (int, error) {
	l.span.SetInnerHTML(fmt.Sprintln(a...))
	return fmt.Println(a...)
}

func (l Logger) Print(a ...interface{}) (int, error) {
	l.span.SetInnerHTML(fmt.Sprint(a...))
	return fmt.Print(a...)
}
