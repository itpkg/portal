package base

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type I18n struct {
	Db *gorm.DB `inject:""`
}

func (p *I18n) T(c *gin.Context, code string, arg ...interface{}) string {
	l := Locale{}
	if e := p.Db.Select("message").Where("code = ? AND lang = ?", code, ParseLocale(c)).First(&l).Error; e == nil {
		return fmt.Sprintf(l.Message, arg...)
	} else {
		return code
	}
}
