package utils_test

import (
	"testing"

	"github.com/itpkg/portal/base/utils"
)

func TestSha(t *testing.T) {
	s := []byte("123456")
	if h, e := utils.Ssha512(s, 32); e == nil {
		t.Logf("ssha512(%s) = %s", s, h)
		if r, e := utils.Csha512(h, s); e == nil {
			if !r {
				t.Errorf("bad in check ssha512")
			}
		} else {
			t.Errorf("bad in csha512: %v", e)
		}
	} else {
		t.Errorf("bad in ssha512: %v", e)
	}
}

func TestMd5(t *testing.T) {
	s := "123456"
	m := "e10adc3949ba59abbe56e057f20f883e"
	if r := utils.Md5([]byte(s)); r != m {
		t.Errorf("md5(%s) want %s, get %s", s, m, r)
	}
}

func TestShell(t *testing.T) {
	if e := utils.Shell("uname", "-a"); e != nil {
		t.Errorf("bad in shell: %v", e)
	}
}

func TestRandAndBase64(t *testing.T) {
	b, e := utils.RandomBytes(8)
	if e != nil {
		t.Errorf("bad in random bytes: %v", e)
	}
	s := utils.ToBase64(b)
	t.Logf("base64: %s", s)
	if _, e = utils.FromBase64(s); e != nil {
		t.Errorf("decode base64 error: %v", e)
	}
}
