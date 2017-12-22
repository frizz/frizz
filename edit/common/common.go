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
	Files map[string][]byte
}

func NewBlob(in []byte) (*Blob, error) {
	blob := &Blob{}
	buf := bytes.NewBuffer(in)
	if err := gob.NewDecoder(buf).Decode(blob); err != nil {
		return nil, err
	}
	return blob, nil
}
