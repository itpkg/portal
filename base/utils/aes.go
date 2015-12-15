package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

type Aes struct {
	//16、24或者32位的[]byte，分别对应AES-128, AES-192或AES-256算法
	Cip cipher.Block `inject:"aes.cipher"`
}

func (p *Aes) Encrypt(pn []byte) ([]byte, error) {

	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(p.Cip, iv)
	ct := make([]byte, len(pn))
	cfb.XORKeyStream(ct, pn)

	return append(ct, iv...), nil

}

func (p *Aes) Decrypt(sr []byte) ([]byte, error) {
	bln := len(sr)
	cln := bln - aes.BlockSize
	ct := sr[0:cln]
	iv := sr[cln:bln]

	cfb := cipher.NewCFBDecrypter(p.Cip, iv)
	pt := make([]byte, cln)
	cfb.XORKeyStream(pt, ct)
	return pt, nil
}

//==============================================================================
func NewAesCipher(key []byte) (cipher.Block, error) {
	if len(key) != 32 {
		return nil, errors.New("bad length, must 32")
	}
	return aes.NewCipher(key)
}
