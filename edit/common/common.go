package common

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

type RequestInfo struct {
	Id   []byte
	Hash []byte
}

func NewRequestInfo(infoBase64 string) (*RequestInfo, error) {
	info := &RequestInfo{}
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

type Blob struct {
	Hash   uint64
	Source map[string]uint64 // package path and hash of all local (e.g. non standard library) packages needed to compile this package
	Lib    []string          // list of all standard library packages needed to compile this package
}

type Bundle struct {
	Hash  uint64
	Files map[string][]byte
}
