package cfg

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Database struct {
	Adapter  string            `toml:"adapter"`
	Host     string            `toml:"host"`
	Port     int               `toml:"port"`
	Name     string            `toml:"name"`
	User     string            `toml:"user"`
	Password string            `toml:"password"`
	Extra    map[string]string `toml:"extra"`
}

func (p *Database) Execute(sql string) (string, []string) {
	switch p.Adapter {
	case "postgres":
		return "psql", []string{"-U", p.User, "-h", p.Host, "-p", strconv.Itoa(p.Port), "-c", sql}
	default:
		return "echo", []string{fmt.Sprintf("Unsupported database adapter %s", p.Adapter)}
	}
}

func (p *Database) Console() (string, []string) {
	switch p.Adapter {
	case "postgres":
		return "psql", []string{"-U", p.User, "-h", p.Host, "-p", strconv.Itoa(p.Port), p.Name}
	default:
		return "echo", []string{fmt.Sprintf("Unsupported database adapter %s", p.Adapter)}
	}
}

func (p *Database) Open() (*gorm.DB, error) {
	var db gorm.DB
	var err error
	switch p.Adapter {
	case "postgres":
		if db, err = gorm.Open("postgres", fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s", p.User, p.Password, p.Host, p.Port, p.Name, p.Extra["sslmode"])); err != nil {
			return nil, err
		}

	default:
		return nil, errors.New(fmt.Sprintf("Unsupported adapter %s", p.Adapter))
	}
	if err := db.DB().Ping(); err == nil {
		return &db, nil
	} else {
		return nil, err
	}
}
