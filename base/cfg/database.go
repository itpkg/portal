package cfg

import (
	"errors"
	"fmt"

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

func (p *Database) Open() (*gorm.DB, error) {
	var db gorm.DB
	var err error
	switch p.Adapter {
	case "postgres":
		if db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", p.Host, p.Port, p.Name, p.User, p.Password, p.Extra["sslmode"])); err != nil {
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
