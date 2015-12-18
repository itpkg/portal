package cache

type Provider interface {
	GetOrSet(key string, val interface{}, cb func(interface{}) (uint, error)) error
	Set(key string, val interface{}, minutes uint) error
	Get(key string, val interface{}) error
}
