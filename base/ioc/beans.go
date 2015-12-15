package ioc

import (
	"github.com/facebookgo/inject"
)

var beans inject.Graph

func In(objects ...interface{}) error {
	for _, v := range objects {
		if e := beans.Provide(&inject.Object{Value: v}); e != nil {
			return e
		}
	}
	return nil
}

func Use(objects map[string]interface{}) error {
	for n, v := range objects {
		if e := beans.Provide(&inject.Object{Name: n, Value: v}); e != nil {
			return e
		}
	}
	return nil
}

func Get(name string) interface{} {
	for _, o := range beans.Objects() {
		if o.Name == name {
			return o.Value
		}
	}
	return nil
}

func Init() error {
	return beans.Populate()
}
