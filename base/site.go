package base

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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
	ID      uint   `gorm:"primary_key"`
	Code    string `sql:"not null;index:idx_locales_code"`
	Lang    string `sql:"not null;default:'en';type:varchar(5);index:idx_locales_lang"`
	Message string `sql:"not null;type:text"`
}

//==============================================================================
type SiteEngine struct {
	Db *gorm.DB `inject:"db"`
}

func (p *SiteEngine) Mount(router *gin.Engine) {
	router.GET("/locales/:lang", func(c *gin.Context) {
		lang := c.Param("lang")
		items := make([]Locale, 0)
		p.Db.Select("code, message").Where("code LIKE ? AND lang = ?", "web.%", lang).Order("code").Find(&items)

		rt := make(map[string]interface{})
		for _, item := range items {
			codes := strings.Split(item.Code[4:], ".")
			tmp := rt
			for i, c := range codes {
				if i+1 == len(codes) {
					tmp[c] = item.Message
				} else {
					if tmp[c] == nil {
						tmp[c] = make(map[string]interface{})
					}
					tmp = tmp[c].(map[string]interface{})
				}
			}
		}
		c.JSON(http.StatusOK, rt)
	})
}

func (p *SiteEngine) Seed() error {
	root := "locales"
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}
	for _, file := range files {
		log.Printf("Find file %s/%s", root, file.Name())
		ss := strings.Split(file.Name(), ".")
		if len(ss) < 3 {
			return errors.New(fmt.Sprintf("bad filename %s", file.Name()))
		}
		lang := ss[1]

		var fd *os.File

		if fd, err = os.Open(fmt.Sprintf("%s/%s", root, file.Name())); err != nil {
			return err
		}
		scanner := bufio.NewScanner(fd)

		for scanner.Scan() {
			line := scanner.Text()
			si := strings.Split(strings.TrimSpace(line), "=")
			if len(si) < 2 {
				log.Printf("ingnore line %s", line)
			}

			code := fmt.Sprintf("%s.%s", ss[0], si[0])
			var c int
			p.Db.Model(&Locale{}).Where("code = ? AND lang = ?", code, lang).Count(&c)
			if c == 0 {
				if err = p.Db.Create(&Locale{Code: code, Lang: lang, Message: strings.Join(si[1:], "=")}).Error; err != nil {
					return err
				}
			}
		}

	}

	return nil
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
