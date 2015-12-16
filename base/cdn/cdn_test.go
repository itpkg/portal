package cdn_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/itpkg/portal/base/cdn"
)

var name = "hello"
var content = "Hi, this is it-package."

func TestLocal(t *testing.T) {
	var p cdn.Provider
	p = &cdn.LocalProvider{Root: "tmp"}
	for i := 0; i < 5; i++ {
		if err := p.Write(fmt.Sprintf("hello/%d", i), fmt.Sprintf("hello_%d", i), func(wrt io.Writer) error {
			_, err := wrt.Write([]byte(content))
			return err
		}); err != nil {
			t.Errorf("bad in local: %v", err)
		}
	}
}
