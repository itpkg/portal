package utils

import (
	"crypto/rand"

	"github.com/pborman/uuid"
)

func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, e := rand.Read(b)
	return b, e
}

func Uuid() string {
	return uuid.New()
}

func AppendSalt(src, salt []byte) []byte {
	return append(src, salt...)
}

func ParseSalt(src []byte, length int) ([]byte, []byte) {
	size := len(src)
	return src[0 : size-length], src[size-length : size]
}

func Equal(src []byte, dst []byte) bool {
	if src == nil && dst == nil {
		return true
	}
	if len(src) == len(dst) {
		for i, b := range src {
			if b != dst[i] {
				return false
			}
		}
	}
	return false
}
