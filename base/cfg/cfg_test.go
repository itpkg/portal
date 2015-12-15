package cfg_test

import (
	"os"
	"testing"

	"github.com/BurntSushi/toml"
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

	if e := write("http.toml", http); e != nil {
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

	if e := write("database.toml", database); e != nil {
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

	if e := write("redis.toml", redis); e != nil {
		t.Errorf("redis error: %v", e)
	}

}

func TestLoad(t *testing.T) {
	if r, e := cfg.Load("redis.toml", "test"); e == nil {
		t.Logf("redis test: %v", r)
	} else {
		t.Errorf("load error: %v", e)
	}
}

func write(name string, val map[string]interface{}) error {
	f, e := os.Create(name)
	if e != nil {
		return e
	}
	defer f.Close()
	return toml.NewEncoder(f).Encode(val)
}

func random() string {
	b, _ := utils.RandomBytes(512)
	return utils.Bytes2String(b)
}
