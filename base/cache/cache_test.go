package cache_test

import (
	"testing"

	"github.com/itpkg/portal/base/cache"
	"github.com/itpkg/portal/base/cfg"
)

type S struct {
	Val int
}

func TestRedis(t *testing.T) {
	rc := cfg.Redis{Host: "localhost", Port: 6379}
	var cp cache.Provider
	cp = &cache.RedisProvider{Redis: rc.Open()}
	val := 111
	s := S{}
	if err := cp.GetOrSet("kkk", &s, func(o interface{}) (uint, error) {
		s1 := o.(*S)
		s1.Val = val
		return 60, nil
	}); err != nil {
		t.Errorf("Bad in get or set: %v", err)
	}
	if s.Val != val {
		t.Errorf("Wang %d, get %d", val, s.Val)
	}
}
