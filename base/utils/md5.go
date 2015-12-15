package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(p []byte) string {
	buf := md5.Sum([]byte(p))
	return hex.EncodeToString(buf[:])
}
