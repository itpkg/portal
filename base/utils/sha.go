package utils

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func Sha512(p []byte) string {
	buf := sha512.Sum512(p)
	return hex.EncodeToString(buf[:])
}

func Ssha512(p []byte, l int) (string, error) {
	salt := make([]byte, l)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	return ssha512(p, salt), nil
}

func ssha512(d, s []byte) string {
	buf := sha512.Sum512(append(d, s...))
	return base64.StdEncoding.EncodeToString(append(buf[:], s...))
}

func Csha512(d string, p []byte) (bool, error) {
	buf, err := base64.StdEncoding.DecodeString(d)
	if err == nil {
		salt := buf[sha512.Size:]
		return ssha512(p, salt) == d, nil
	}
	return false, err
}
