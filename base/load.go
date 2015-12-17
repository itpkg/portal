package base

import (
	"github.com/itpkg/portal/base/cfg"
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
		e := h[env]
		e.Env = env
		return e, err
	}
}
