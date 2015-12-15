package utils

import (
	"crypto/rand"
)

func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, e := rand.Read(b)
	return b, e
}
