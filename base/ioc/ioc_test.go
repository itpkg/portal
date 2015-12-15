package ioc_test

import (
	"testing"
	"time"

	"github.com/itpkg/portal/base/ioc"
)

type S struct {
	Value   int        `inject:"value"`
	Version string     `inject:"version"`
	Now     *time.Time `inject:"now"`
}

func TestIoc(t *testing.T) {
	val := 1234
	now := time.Now()
	s := S{}
	if e := ioc.Use(map[string]interface{}{"value": val, "version": "v20130425", "now": &now}); e != nil {
		t.Errorf("bad in use: %v", e)
	}
	if e := ioc.Use(map[string]interface{}{"s": &s}); e != nil {
		t.Errorf("bad in use: %v", e)
	}
	if e := ioc.Init(); e != nil {
		t.Errorf("bad in init: %v", e)
	}
	if v := ioc.Get("value").(int); v != val {
		t.Errorf("want %d, get %d", val, v)
	}
}
