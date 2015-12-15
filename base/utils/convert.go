package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

func FromToml(f string, v interface{}) error {
	_, err := toml.DecodeFile(f, v)
	return err
}

func ToToml(f string, v interface{}) error {
	fi, err := os.Create(f)
	defer fi.Close()

	if err == nil {
		end := toml.NewEncoder(fi)
		err = end.Encode(v)
	}
	return err
}

func ToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func FromBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func FromHex(s string) ([]byte, error) {
	return hex.DecodeString(s)
}

func ToHex(bs []byte) string {
	return hex.EncodeToString(bs)
}

func ToBits(obj interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(obj)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func FromBits(data []byte, obj interface{}) error {
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(data)
	err := dec.Decode(obj)
	if err != nil {
		return err
	}
	return nil
}

func ToJson(o interface{}) ([]byte, error) {
	return json.Marshal(o)

}

func FromJson(j []byte, o interface{}) error {
	return json.Unmarshal(j, o)
}

func init() {
	gob.Register(time.Time{})
}
