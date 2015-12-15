package base

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itpkg/portal/base/engine"
	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `sql:"not null;default:current_timestamp"`
	UpdatedAt time.Time `sql:"not null"`
}

type Setting struct {
	Model
	Key  string `sql:"not null;unique"`
	Val  []byte `sql:"not null"`
	Flag bool   `sql:"not null;default:false"`
}

type Locale struct {
	Model
	Code    string `sql:"not null;index:idx_locales_code"`
	Lang    string `sql:"not null;default:'en';type:varchar(5);index:idx_locales_lang"`
	Message string `sql:"not null;type:text"`
}

//==============================================================================
type SiteEngine struct {
	Db *gorm.DB `inject:"db"`
}

func (p *SiteEngine) Mount(*gin.Engine) {
}

func (p *SiteEngine) Seed() {

}

func (p *SiteEngine) Migrate() {
	db := p.Db
	db.AutoMigrate(&Locale{}, &Setting{})
	db.Model(&Locale{}).AddUniqueIndex("idx_locales_code_lang", "code", "lang")
}

//==============================================================================
func init() {
	engine.Register(&SiteEngine{})
}
