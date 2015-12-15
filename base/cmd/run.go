package cmd

import (
	"os"

	"github.com/codegangsta/cli"
)

func Run() error {
	app := cli.NewApp()
	app.Name = "portal"
	app.Usage = "it-package portal system"
	app.Version = "v20151214"
	app.Commands = commands

	return app.Run(os.Args)
}
