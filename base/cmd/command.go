package cmd

import (
	"log"

	"github.com/codegangsta/cli"
)

func Action(act func(env string) error) func(c *cli.Context) {
	return func(c *cli.Context) {
		if e := act(c.String("environment")); e == nil {
			log.Println("Done!!!")
		} else {
			log.Fatalln(e)
		}
	}
}

var commands = make([]cli.Command, 0)

func Register(args ...cli.Command) {
	commands = append(commands, args...)
}
