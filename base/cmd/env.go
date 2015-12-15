package cmd

import (
	"github.com/codegangsta/cli"
)

var ENV = cli.StringFlag{
	Name:   "environment, e",
	Value:  "development",
	Usage:  "Specifies the environment to run this server under (test/development/production).",
	EnvVar: "ITPKG_ENV",
}
