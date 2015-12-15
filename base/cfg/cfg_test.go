package cfg_test

import (
	"testing"

	"github.com/itpkg/portal/base/cfg"
	"github.com/itpkg/portal/base/utils"
)

func TestHttp(t *testing.T) {
	http := map[string]interface{}{
		"development": &cfg.Http{
			Domain:  "localhost.localdomain",
			Port:    3000,
			Secrets: random(),
		},
		"production": &cfg.Http{
			Domain:  "change-me",
			Port:    3000,
			Secrets: random(),
		},
		"test": &cfg.Http{
			Domain:  "localhost.localdomain",
			Port:    3000,
			Secrets: random(),
		},
	}

	if e := utils.ToToml("http.toml", http); e != nil {
		t.Errorf("http error: %v", e)
	}

}

func TestDatabase(t *testing.T) {
	database := map[string]interface{}{
		"development": &cfg.Database{
			Adapter: "postgres",
			Host:    "localhost",
			Port:    5432,
			Name:    "itpkg_portal_dev",
			User:    "postgres",
			Extra:   map[string]string{"sslmode": "disable"},
		},
		"production": &cfg.Database{
			Adapter: "postgres",
			Host:    "localhost",
			Port:    5432,
			Name:    "itpkg_portal_prod",
			User:    "portal",
			Extra:   map[string]string{"sslmode": "disable"},
		},
		"test": &cfg.Database{
			Adapter: "postgres",
			Host:    "localhost",
			Port:    5432,
			Name:    "itpkg_portal_test",
			User:    "postgres",
			Extra:   map[string]string{"sslmode": "disable"},
		},
	}

	if e := utils.ToToml("database.toml", database); e != nil {
		t.Errorf("database error: %v", e)
	}

}

func TestRedis(t *testing.T) {
	redis := map[string]interface{}{
		"development": &cfg.Redis{
			Host: "localhost",
			Port: 6379,
			Db:   1,
		},
		"production": &cfg.Redis{
			Host: "localhost",
			Port: 6379,
			Db:   2,
		},
		"test": &cfg.Redis{
			Host: "localhost",
			Port: 6379,
			Db:   2,
		},
	}

	if e := utils.ToToml("redis.toml", redis); e != nil {
		t.Errorf("redis error: %v", e)
	}

}

func TestLoad(t *testing.T) {
	r := make(map[string]*cfg.Redis)
	if e := utils.FromToml("redis.toml", r); e == nil {
		t.Logf("redis test: %v", r["test"])
	} else {
		t.Errorf("load error: %v", e)
	}
}

func random() string {
	b, _ := utils.RandomBytes(512)
	return utils.ToBase64(b)
}
