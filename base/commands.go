package base

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/itpkg/portal/base/cfg"
	"github.com/itpkg/portal/base/cmd"
	"github.com/itpkg/portal/base/utils"
)

func Database(env string) (*cfg.Database, error) {
	db := make(map[string]*cfg.Database)
	if err := utils.FromToml("config/database.toml", db); err != nil {
		return nil, err
	} else {
		return db[env], err
	}
}

func Redis(env string) (*cfg.Redis, error) {
	r := make(map[string]*cfg.Redis)
	if err := utils.FromToml("config/redis.toml", r); err != nil {
		return nil, err
	} else {
		return r[env], err
	}
}

func Http(env string) (*cfg.Http, error) {
	h := make(map[string]*cfg.Http)
	if err := utils.FromToml("config/http.toml", h); err != nil {
		return nil, err
	} else {
		return h[env], err
	}
}

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
						db, e := Database(env)
						if e != nil {
							return e
						}
						c, a := db.Execute(fmt.Sprintf("CREATE DATABASE %s WITH ENCODING='UTF8'", db.Name))
						return utils.Shell(c, a...)
					}),
				},
				{
					Name:    "console",
					Aliases: []string{"c"},
					Usage:   "start a console for the database",
					Flags:   []cli.Flag{cmd.ENV},
					Action: cmd.Action(func(env string) error {

						db, e := Database(env)
						if e != nil {
							return e
						}

						c, a := db.Console()
						return utils.Shell(c, a...)
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
						return nil
					}),
				},
				{
					Name:    "drop",
					Aliases: []string{"d"},
					Usage:   "drops the database",
					Flags:   []cli.Flag{cmd.ENV},
					Action: cmd.Action(func(env string) error {

						db, e := Database(env)
						if e != nil {
							return e
						}

						c, a := db.Execute(fmt.Sprintf("DROP DATABASE %s", db.Name))
						return utils.Shell(c, a...)
					}),
				},
			},
		},
	)
}
