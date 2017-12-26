package common

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
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
