package base

import (
	"crypto/cipher"

	re "github.com/garyburd/redigo/redis"
	"github.com/itpkg/portal/base/cdn"
	"github.com/itpkg/portal/base/cfg"
	"github.com/itpkg/portal/base/email"
	"github.com/itpkg/portal/base/engine"
	"github.com/itpkg/portal/base/ioc"
	"github.com/itpkg/portal/base/token"
	"github.com/itpkg/portal/base/utils"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

const LAYOUT = "views/layout.html"

func Init(env string) error {
	var err error
	var http *cfg.Http
	if http, err = Http(env); err != nil {
		return err
	}
	//-----------------logger
	var logger = logging.MustGetLogger("portal")
	if http.IsProduction() {
		if bkd, err := logging.NewSyslogBackend("itpkg"); err == nil {
			logging.SetBackend(bkd)
		} else {
			return err
		}
	}

	//------------------email
	dao := &Dao{}
	var emailP email.Provider
	if http.IsProduction() {
		emailP = &email.SmtpProvider{
			Func: func() (*email.Smtp, error) {
				smtp := email.Smtp{}
				err := dao.Get("site.smtp", &smtp)
				return &smtp, err
			},
		}
	} else {
		emailP = &email.LocalProvider{}
	}

	//-----------------aes
	var aesKey []byte
	if aesKey, err = http.Key(60, 92); err != nil {
		return err
	}
	var cip cipher.Block
	if cip, err = utils.NewAesCipher(aesKey); err != nil {
		return err
	}

	//--------------database
	var db *gorm.DB
	if dbc, err := Database(env); err == nil {
		if db, err = dbc.Open(); err != nil {
			return err
		}

	} else {
		return err
	}
	if !http.IsProduction() {
		db.LogMode(true)
	}

	//------------redis
	var redis *re.Pool
	if rec, err := Redis(env); err == nil {
		redis = rec.Open()
	} else {
		return err
	}

	//--------------cdn
	var cdnP cdn.Provider
	if http.IsProduction() {
		cdnP = &cdn.LocalProvider{Root: "public"}
	} else {
		cdnP = &cdn.LocalProvider{Root: "assets"}
	}

	//---------------token
	var tokenP token.Provider
	tokenP = &token.RedisProvider{}

	//--------------
	if err = ioc.In(db, redis, logger,
		cdnP, dao,
		emailP, tokenP); err != nil {
		return err
	}
	if err = ioc.Use(map[string]interface{}{
		"http":       http,
		"aes.cipher": cip,
	}); err != nil {
		return err
	}

	if err = engine.Loop(func(en engine.Engine) error {
		return ioc.In(en)
	}); err != nil {
		return err
	}

	return ioc.Init()
}
