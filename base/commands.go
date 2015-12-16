package base

import (
	"fmt"

	"github.com/codegangsta/cli"
	re "github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/itpkg/portal/base/cdn"
	"github.com/itpkg/portal/base/cfg"
	"github.com/itpkg/portal/base/cmd"
	"github.com/itpkg/portal/base/engine"
	"github.com/itpkg/portal/base/ioc"
	"github.com/itpkg/portal/base/utils"
	"github.com/jinzhu/gorm"
)

func Init(env string) error {
	var db *gorm.DB
	if dbc, err := Database(env); err == nil {
		if db, err = dbc.Open(); err != nil {
			return err
		}

	} else {
		return err
	}
	if env != "production" {
		db.LogMode(true)
	}

	var redis *re.Pool
	if rec, err := Redis(env); err == nil {
		redis = rec.Open()
	} else {
		return err
	}

	var err error
	var http *cfg.Http
	if http, err = Http(env); err != nil {
		return err
	}

	if err = ioc.In(db, redis, &cdn.LocalProvider{Root: "public"}); err != nil {
		return err
	}
	if err = ioc.Use(map[string]interface{}{"http": http}); err != nil {
		return err
	}

	if err = engine.Loop(func(en engine.Engine) error {
		return ioc.In(en)
	}); err != nil {
		return err
	}

	return ioc.Init()
}

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
				if err := Init(env); err != nil {
					return err
				}
				if env == "production" {
					gin.SetMode(gin.ReleaseMode)
				}
				route := gin.Default()
				if err := engine.Loop(func(en engine.Engine) error {
					en.Mount(route)
					return nil
				}); err != nil {
					return err
				}
				http := ioc.Get("http").(*cfg.Http)
				return route.Run(fmt.Sprintf(":%d", http.Port))
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
