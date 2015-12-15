package utils_test

import (
	"testing"

	"github.com/itpkg/portal/base/utils"
)

func TestAes(t *testing.T) {
	c, e := utils.NewAesCipher(key)
	if e != nil {
		t.Errorf("bad in new aes cipher: %v", e)
	}
	a := utils.Aes{Cip: c}
	dest1, _ := a.Encrypt([]byte(hello))
	dest2, _ := a.Encrypt([]byte(hello))
	t.Logf("AES1: %v", dest1)
	t.Logf("AES2: %v", dest2)

	src, _ := a.Decrypt(dest1)
	if string(src) != hello {
		t.Errorf("val == %x, want %x", src, hello)
	}

}
