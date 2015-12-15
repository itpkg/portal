package base

import (
	"github.com/itpkg/portal/base/cmd"
)

func init() {
	cmd.Register("server", "s", "start the server", func(env string) error {
		//todo
		return nil
	})
}
