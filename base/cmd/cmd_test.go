package cmd_test

import (
	"fmt"
	"testing"

	"github.com/itpkg/portal/base/cmd"
)

func TestCmd(t *testing.T) {
	cmd.Register("hello", "h", "echo env", func(name string) error {
		fmt.Println("env: %s", name)
		return nil
	})

}
