package utils

import (
	"crypto/hmac"
	"crypto/sha512"
	"hash"
)

type Hmac struct {
	Key []byte           `inject:"hmac.key"` //32 bits
	Fn  func() hash.Hash `inject:"hmac.fn"`
}

func (p *Hmac) Sum(src []byte) []byte {
	mac := hmac.New(p.Fn, p.Key)
	mac.Write(src)
	return mac.Sum(nil)
}

func (p *Hmac) Equal(src, dst []byte) bool {
	return hmac.Equal(src, dst)
}

//=========================================================================
func NewHmacHash() func() hash.Hash {
	return sha512.New
}
