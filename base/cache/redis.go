package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type RedisProvider struct {
	Redis *redis.Pool `inject:""`
}

func (p *RedisProvider) key(k string) string {
	return fmt.Sprintf("cache://%s", k)
}

func (p *RedisProvider) GetOrSet(key string, val interface{}, cb func(interface{}) (uint, error)) error {
	if err := p.Get(key, val); err == nil {
		return nil
	}
	t, e := cb(val)
	if e == nil {
		return p.Set(key, val, t)
	} else {
		return e
	}
}

func (p *RedisProvider) Set(key string, val interface{}, minutes uint) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(val); err != nil {
		return err
	}
	rc := p.Redis.Get()
	defer rc.Close()
	_, err := rc.Do("SET", p.key(key), buf.Bytes(), "EX", minutes*60)
	return err
}

func (p *RedisProvider) Get(key string, val interface{}) error {
	rc := p.Redis.Get()
	defer rc.Close()
	bys, err := redis.Bytes(rc.Do("GET", p.key(key)))
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	buf.Write(bys)
	return dec.Decode(val)
}
