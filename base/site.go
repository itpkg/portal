package base

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/expvar"
	"github.com/gin-gonic/gin"
	"github.com/itpkg/portal/base/cdn"
	"github.com/itpkg/portal/base/cfg"
	"github.com/itpkg/portal/base/engine"
	"github.com/itpkg/portal/base/seo"
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
	Db   *gorm.DB     `inject:""`
	Cdn  cdn.Provider `inject:""`
	Dao  *Dao         `inject:""`
	Http *cfg.Http    `inject:""`
}

func (p *SiteEngine) Build(dir string) error {
	//------------------------- by lang ------------------------------------
	rows, err := p.Db.Model(&Locale{}).Select("DISTINCT(lang)").Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		//-------------------------locales------------------------------------
		var lang string
		rows.Scan(&lang)

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

		if err := p.Cdn.Write("locales", fmt.Sprintf("%s.json", lang), func(wrt io.Writer) error {
			end := json.NewEncoder(wrt)
			return end.Encode(rt)
		}); err != nil {
			return err
		}

		//------------------rss.atom--------------------------------------
		if err := p.Cdn.Write("", "rss.atom", func(wrt io.Writer) error {
			return seo.Rss(wrt, lang,
				p.Dao.GetSiteInfo("title", lang),
				fmt.Sprintf("https://www.%s", p.Http.Domain),
				p.Dao.GetSiteInfo("author.name", ""),
				p.Dao.GetSiteInfo("author.email", ""),
			)
		}); err != nil {
			return err
		}
	}

	//------------------sitemap.xml.gz--------------------------------
	p.Cdn.Write("", "sitemap.xml.gz", func(wrt io.Writer) error {
		return seo.Sitemap(wrt)
	})
	//------------------robots.txt------------------------------------
	rn, rb := seo.Robots(p.Dao.GetSiteInfo("robots.txt", ""))
	if err := p.Cdn.Write("", rn, func(wrt io.Writer) error {
		_, e := wrt.Write([]byte(rb))
		return e
	}); err != nil {
		return err
	}
	//------------------google----------------------------------------
	gn, gb := seo.GoogleVerify(p.Dao.GetSiteInfo("google.verify", ""))
	if err := p.Cdn.Write("", gn, func(wrt io.Writer) error {
		_, e := wrt.Write([]byte(gb))
		return e
	}); err != nil {
		return err
	}
	//------------------baidu----------------------------------------
	bn, bb := seo.BaiduVerify(p.Dao.GetSiteInfo("baidu.verify", ""))
	if err := p.Cdn.Write("", bn, func(wrt io.Writer) error {
		_, e := wrt.Write([]byte(bb))
		return e
	}); err != nil {
		return err
	}

	//-------------------------
	return nil
}

func (p *SiteEngine) Mount(router *gin.Engine) {
	router.GET("/debug/vars", expvar.Handler()) //todo only admin can
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
		zone := ss[0]
		lang := ss[1]

		var fd *os.File

		if fd, err = os.Open(fmt.Sprintf("%s/%s", root, file.Name())); err != nil {
			return err
		}
		scanner := bufio.NewScanner(fd)

		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if len(line) == 0 || line[0] == '#' {
				continue
			}
			idx := strings.Index(line, "=")
			var key, val string
			if idx == -1 {
				key = line
				val = " "
			} else {
				key = strings.TrimSpace(line[0:idx])
				val = strings.TrimSpace(line[idx+1:])
			}

			code := fmt.Sprintf("%s.%s", zone, key)
			var c int
			p.Db.Model(&Locale{}).Where("code = ? AND lang = ?", code, lang).Count(&c)
			if c == 0 {
				if err = p.Db.Create(&Locale{Code: code, Lang: lang, Message: val}).Error; err != nil {
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
