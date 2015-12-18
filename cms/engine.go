package cms

import (
	"github.com/itpkg/portal/base/engine"
	"github.com/jinzhu/gorm"
)

type Engine struct {
	Db *gorm.DB `inject:""`
}

func (p *Engine) Build(string) error {
	//todo
	return nil
}

//-----------------------------------------------------------------------------
func init() {
	engine.Register(&Engine{})
}
