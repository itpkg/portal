package engine

import (
	"github.com/gin-gonic/gin"
)

type Engine interface {
	Mount(*gin.Engine)
	Seed() error
	Migrate()
}

var engines = make([]Engine, 0)

func Register(ens ...Engine) {
	engines = append(engines, ens...)
}

func Loop(fn func(Engine) error) error {
	for _, en := range engines {
		if err := fn(en); err != nil {
			return err
		}
	}
	return nil
}
