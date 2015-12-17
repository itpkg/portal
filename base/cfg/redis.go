package cfg

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
	Db   int    `toml:"db"`
}

func (p *Redis) Open() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, e := redis.Dial("tcp", fmt.Sprintf("%s:%d", p.Host, p.Port))
			if e != nil {
				return nil, e
			}
			if _, e = c.Do("SELECT", p.Db); e != nil {
				c.Close()
				return nil, e
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
