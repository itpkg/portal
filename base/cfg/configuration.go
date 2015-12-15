package cfg

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/BurntSushi/toml"
)

func Load(file, env string) (interface{}, error) {
	items := make(map[string]interface{})
	if _, err := toml.DecodeFile(file, items); err == nil {
		return items[env], nil
	} else {
		return nil, err
	}
}

func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, e := rand.Read(b)
	return b, e
}

func Bytes2String(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func String2Bytes(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
