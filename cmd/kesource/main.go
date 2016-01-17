package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"golang.org/x/tools/imports"
	"kego.io/kerr"
	"kego.io/kerr/kerrsource"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	flag.Parse()
	if err := processFile(flag.Arg(0)); err != nil {
		panic(err.Error())
	}
}

func processFile(filename string) error {

	if filename == "" {
		return kerr.New("CKBKFTNVXI", "File not found")
	}

	options := &imports.Options{
		TabWidth:  4,
		TabIndent: true,
		Comments:  true,
		Fragment:  true,
	}

	f, err := os.Open(filename)
	if err != nil {
		return kerr.Wrap("WCQEVUWMAI", err)
	}
	defer f.Close()

	src, err := ioutil.ReadAll(f)
	if err != nil {
		return kerr.Wrap("DWTHNAVSVY", err)
	}

	var res []byte

	res, err = imports.Process(filename, src, options)
	if err != nil {
		return kerr.Wrap("XFILQRWHSI", err)
	}

	res, err = kerrsource.Process(filename, res)
	if err != nil {
		return kerr.Wrap("UFVVOEXSYE", err)
	}

	if !bytes.Equal(src, res) {
		// formatting has changed
		err = ioutil.WriteFile(filename, res, 0)
		if err != nil {
			return kerr.Wrap("NFFAOHJQLG", err)
		}
	}

	return nil

}
