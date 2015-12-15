package main

import (
	_ "github.com/itpkg/portal/base"
	"github.com/itpkg/portal/base/cmd"
	_ "github.com/lib/pq"
)

func main() {
	cmd.Run()
}
