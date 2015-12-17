package base

import (
	"fmt"
	"os"
	"text/template"

	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"github.com/itpkg/portal/base/cfg"
	"github.com/itpkg/portal/base/cmd"
	"github.com/itpkg/portal/base/engine"
	"github.com/itpkg/portal/base/ioc"
	"github.com/itpkg/portal/base/utils"
)

func init() {
	cmd.Register(
		cli.Command{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start the web server",
			Flags:   []cli.Flag{cmd.ENV},
			Action: cmd.Action(func(env string) error {
				if err := Init(env); err != nil {
					return err
				}
				http := ioc.Get("http").(*cfg.Http)

				if http.IsProduction() {
					gin.SetMode(gin.ReleaseMode)
				}
				router := gin.Default()
				if !http.IsProduction() {
					router.Static("/assets", "assets")
				}

				if err := engine.Loop(func(en engine.Engine) error {
					en.Mount(router)
					return nil
				}); err != nil {
					return err
				}

				return router.Run(fmt.Sprintf(":%d", http.Port))
			}),
		},
		cli.Command{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build static files",
			Flags:   []cli.Flag{cmd.ENV},
			Action: cmd.Action(func(env string) error {
				if err := Init(env); err != nil {
					return err
				}
				return engine.Loop(func(en engine.Engine) error {
					return en.Build("public")
				})
			}),
		},
		cli.Command{
			Name:    "nginx",
			Aliases: []string{"n"},
			Usage:   "generate nginx files",
			Flags:   []cli.Flag{cmd.ENV},
			Action: cmd.Action(func(env string) error {
				h, e1 := Http(env)
				if e1 != nil {
					return e1
				}
				t, e2 := template.ParseFiles("views/nginx.conf")
				if e2 != nil {
					return e2
				}
				f, e3 := os.OpenFile("config/nginx.conf", os.O_WRONLY|os.O_CREATE, 0600)
				if e3 != nil {
					return e3
				}
				return t.Execute(f, h)
			}),
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
						if err := Init(env); err != nil {
							return err
						}
						return engine.Loop(func(en engine.Engine) error {
							en.Migrate()
							return nil
						})
					}),
				},
				{
					Name:    "seed",
					Aliases: []string{"s"},
					Usage:   "load the seed data",
					Flags:   []cli.Flag{cmd.ENV},
					Action: cmd.Action(func(env string) error {
						if err := Init(env); err != nil {
							return err
						}
						return engine.Loop(func(en engine.Engine) error {
							return en.Seed()
						})
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
