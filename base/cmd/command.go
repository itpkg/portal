package cmd

import (
	"log"

	"github.com/codegangsta/cli"
)

type Command func(env string) error

var commands = make([]cli.Command, 0)

func Register(name, alias, usage string, command Command) {
	commands = append(commands, cli.Command{
		Name:    name,
		Aliases: []string{alias},
		Usage:   usage,
		Flags:   []cli.Flag{ENV},
		Action: func(c *cli.Context) {
			if e := command(c.String("environment")); e == nil {
				log.Println("Done!!!")
			} else {
				log.Fatalln(e)
			}
		},
	})
}
