package base

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/itpkg/portal/base/cmd"
)

func init() {
	cmd.Register(
		cli.Command{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the web server",
			Flags:   []cli.Flag{cmd.ENV},
			Action: cmd.Action(func(env string) error {
				//todo
				return nil
			}),
		},
		cli.Command{
			Name:    "cache",
			Aliases: []string{"c"},
			Usage:   "cache operations",
			Subcommands: []cli.Command{
				{
					Name:    "clear",
					Aliases: []string{"c"},
					Usage:   "clear cache items",
					Flags:   []cli.Flag{cmd.ENV},
					Action: cmd.Action(func(env string) error {
						//todo
						log.Println("db m env ", env)
						return nil
					}),
				},
			},
		},
		cli.Command{
			Name:    "database",
			Aliases: []string{"db"},
			Usage:   "database operations",
			Subcommands: []cli.Command{
				{
					Name:    "create",
					Aliases: []string{"n"},
					Usage:   "creates the database",
					Flags:   []cli.Flag{cmd.ENV},
					Action: cmd.Action(func(env string) error {
						//todo
						log.Println("db m env ", env)
						return nil
					}),
				},
				{
					Name:    "console",
					Aliases: []string{"c"},
					Usage:   "start a console for the database",
					Flags:   []cli.Flag{cmd.ENV},
					Action: cmd.Action(func(env string) error {
						//todo
						log.Println("db m env ", env)
						return nil
					}),
				},
				{
					Name:    "migrate",
					Aliases: []string{"m"},
					Usage:   "migrate the database",
					Flags:   []cli.Flag{cmd.ENV},
					Action: cmd.Action(func(env string) error {
						//todo
						log.Println("db m env ", env)
						return nil
					}),
				},
				{
					Name:    "seed",
					Aliases: []string{"s"},
					Usage:   "load the seed data",
					Flags:   []cli.Flag{cmd.ENV},
					Action: cmd.Action(func(env string) error {
						//todo
						log.Println("db m env ", env)
						return nil
					}),
				},
				{
					Name:    "drop",
					Aliases: []string{"d"},
					Usage:   "drops the database",
					Flags:   []cli.Flag{cmd.ENV},
					Action: cmd.Action(func(env string) error {
						//todo
						log.Println("db m env ", env)
						return nil
					}),
				},
			},
		},
	)
}
