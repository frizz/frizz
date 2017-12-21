package auther

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
)

type Auther struct {
	secret []byte
}

func New() Auther {
	a := Auther{
		secret: make([]byte, 256),
	}
	if _, err := rand.Read(a.secret); err != nil {
		panic(err.Error())
	}
	return a
}

func (a Auther) Auth(public []byte, hash []byte) bool {
	return bytes.Equal(hash, a.Sign(public))
}

func (a Auther) Sign(public []byte) []byte {
	// Secret is fixed langth, so no deliminator is needed.
	return sha256.New().Sum(append(a.secret, public...))
}
