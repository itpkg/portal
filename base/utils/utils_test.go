package utils_test

import (
	"testing"

	"github.com/itpkg/portal/base/utils"
)

func TestRandAndBase64(t *testing.T) {
	b, e := utils.RandomBytes(8)
	if e != nil {
		t.Errorf("bad in random bytes: %v", e)
	}
	s := utils.Bytes2String(b)
	t.Logf("base64: %s", s)
	if _, e = utils.String2Bytes(s); e != nil {
		t.Errorf("decode base64 error: %v", e)
	}
}
