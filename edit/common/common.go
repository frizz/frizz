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
	Data   map[string][]byte  // all .frizz.json data files in this package
	Source map[string]Package // go source of all local packages needed to compile this package
	Lib    []string           // list of all standard library packages needed to compile this package
}

type Package struct {
	Hash   uint64
	Source map[string][]byte
}

func NewBlob(in []byte) (*Blob, error) {
	blob := &Blob{}
	buf := bytes.NewBuffer(in)
	if err := gob.NewDecoder(buf).Decode(blob); err != nil {
		return nil, err
	}
	return blob, nil
}

func (b *Blob) ToBytes() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	if err := gob.NewEncoder(buf).Encode(b); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
