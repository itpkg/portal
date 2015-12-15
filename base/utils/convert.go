package utils

import (
	"encoding/base64"
)

func Bytes2String(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func String2Bytes(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
