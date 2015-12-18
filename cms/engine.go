package cms

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"fmt"
	"github.com/itpkg/portal/base"
	"github.com/itpkg/portal/base/cdn"
	"github.com/itpkg/portal/base/cfg"
	"github.com/itpkg/portal/base/engine"
	"github.com/itpkg/portal/base/tpl"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

type Engine struct {
	Db      *gorm.DB        `inject:""`
	Dao     *Dao            `inject:""`
	BaseDao *base.Dao       `inject:""`
	Logger  *logging.Logger `inject:""`
	Cdn     cdn.Provider    `inject:""`
	Http    *cfg.Http       `inject:"http"`
}

func (p *Engine) Build(string) error {
	if err := filepath.Walk("tmp/blogs", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		var buf []byte
		if buf, err = ioutil.ReadFile(path); err != nil {
			return err
		}

		ext := filepath.Ext(path)
		idx := strings.LastIndex(path, "/")
		dir := path[4:idx]
		name := path[idx+1:]
		lang := dir[6:11]

		switch {
		case ext == ".html" || ext == ".htm":
			p.Logger.Info("[%s] %s => %s/%s", lang, path, dir, name)
			return p.Cdn.Write(dir, name, func(wrt io.Writer) error {
				mod := p.BaseDao.GetSiteModel(lang)
				mod.Url = fmt.Sprintf("%s%s/%s", p.Http.Assets(), dir, name)
				st := FirstLine(path)
				mod.SubTitle = st[4 : len(st)-5]
				mod.SetBody(string(buf))
				return tpl.Dump(wrt, base.LAYOUT, mod)
			})
		case ext == ".md":

			name = name[:len(name)-3] + ".html"
			p.Logger.Info("[%s] %s => %s/%s", lang, path, dir, name)

			return p.Cdn.Write(dir, name, func(wrt io.Writer) error {
				mod := p.BaseDao.GetSiteModel(lang)
				mod.Url = fmt.Sprintf("%s%s/%s", p.Http.Assets(), dir, name)
				mod.SubTitle = FirstLine(path)
				mod.SetBody(string(Md2Hm([]byte(strings.Replace(string(buf), ".md)", ".html)", -1)))))
				return tpl.Dump(wrt, base.LAYOUT, mod)
			})
		default:
			p.Logger.Info("%s => %s/%s", path, dir, name)
			return p.Cdn.Write(dir, name, func(wrt io.Writer) error {
				_, e := wrt.Write(buf)
				return e
			})
		}

	}); err != nil {
		return err
	}

	//-----------------
	return nil

}

//-----------------------------------------------------------------------------
func init() {
	en := Engine{}
	engine.Register(&en)
}
