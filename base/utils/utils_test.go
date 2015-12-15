package utils_test

import (
	"testing"

	"github.com/itpkg/portal/base/utils"
)

func TestShell(t *testing.T) {
	if e := utils.Shell("telnet", "www.google.com", "80"); e != nil {
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
