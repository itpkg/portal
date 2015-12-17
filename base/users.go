package base

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itpkg/portal/base/cfg"
	"github.com/itpkg/portal/base/email"
	"github.com/itpkg/portal/base/engine"
	"github.com/itpkg/portal/base/token"
	"github.com/jinzhu/gorm"
	"github.com/op/go-logging"
)

type User struct {
	Model

	Email        string `sql:"not null;unique"`
	Password     string
	Uid          string `sql:"not null;unique;type:char(36)"`
	Logo         string `sql:"not null"`
	Name         string `sql:"not null"`
	ProviderType string `sql:"not null;default:'email';index:idx_users_provider_type"`
	ProviderId   string
	LastSignIn   time.Time `sql:"not null"`
	SignInCount  uint      `sql:"not null;default:0"`

	ConfirmedAt *time.Time
	LockedAt    *time.Time
}

type Contact struct {
	Model
	UserID uint `sql:"not null"`
	User   User

	Email    string
	Qq       string
	Wechat   string
	Weibo    string
	Facebook string
	Blog     string

	Tel     string
	Fax     string
	Address string

	Profile string `sql:"type:text"`
}

type Log struct {
	ID        uint `gorm:"primary_key"`
	UserID    uint `sql:"not null"`
	User      User
	Message   string    `sql:"not null"`
	CreatedAt time.Time `sql:"not null;default:current_timestamp"`
}

type Role struct {
	ID           uint   `gorm:"primary_key"`
	Name         string `sql:"not null;index:idx_roles_name"`
	ResourceType string `sql:"not null;default:'-';index:idx_roles_resource_type"`
	ResourceId   uint   `sql:"not null;default:0"`
}

type Permission struct {
	ID     uint `gorm:"primary_key"`
	User   User
	UserID uint `sql:"not null"`
	Role   Role
	RoleID uint      `sql:"not null"`
	Begin  time.Time `sql:"not null;default:current_date;type:date"`
	End    time.Time `sql:"not null;default:'1000-1-1';type:date"`
}

//==============================================================================
type SignInForm struct {
	RememberMe bool   `form:"remember_me"`
	Email      string `from:"email" binding:"email"`
	Password   string `form:"password" binding:"min=8"`
}
type SignUpForm struct {
	Username             string `form:"username" binding:"min=2,max=20"`
	Email                string `form:"email" binding:"email"`
	Password             string `form:"password" binding:"min=8"`
	PasswordConfirmation string `form:"password_confirmation" binding:"eqfield=Password"`
}

type EmailForm struct {
	Email string `form:"email" binding:"email"`
}

//==============================================================================
type UsersEngine struct {
	Db     *gorm.DB        `inject:""`
	Dao    *Dao            `inject:""`
	Http   *cfg.Http       `inject:"http"`
	I18n   *I18n           `inject:""`
	Smtp   email.Provider  `inject:""`
	Logger *logging.Logger `inject:""`
	Token  token.Provider  `inject:""`
}

func (p *UsersEngine) Build(dir string) error {
	return nil
}

func (p *UsersEngine) sendMail(ctx *gin.Context, action string, user *User) {
	type Model struct {
		Username string
		Url      string
	}

	switch action {
	case "confirm":
	case "unlock":
	case "reset_password":
	default:
		return
	}

	tkn, err := p.Token.New(map[string]interface{}{"action": action, "user": user.Uid}, 60)
	if err != nil {
		p.Logger.Error("bad in generate token: %v", err)
		return
	}

	var buf bytes.Buffer
	if tpl, err := template.New("").Parse(p.I18n.T(ctx, fmt.Sprintf("email.user.%s.body", action))); err == nil {
		if err := tpl.Execute(&buf, &Model{Username: user.Name, Url: fmt.Sprintf("%s/users/%s?token=%s", p.Http.Home(), action, tkn)}); err != nil {
			p.Logger.Error("bad in parse email template email.user.%s.body: %v", action, err)
			return
		}
	} else {
		p.Logger.Error("bad template email.user.%s.body: %v", action, err)
		return
	}
	p.Smtp.Send(
		[]string{user.Email},
		p.I18n.T(ctx, fmt.Sprintf("email.user.%s.subject", action)),
		buf.String(),
	)

}

func (p *UsersEngine) token(action string, fn func(*gin.Context, *User) (bool, string)) func(*gin.Context) {
	return func(c *gin.Context) {
		tkn, err := p.Token.Parse(c.DefaultQuery("token", ""))
		if err != nil {
			Message(c, false, err.Error())
			return
		}
		if tkn["action"].(string) != action {
			Message(c, false, p.I18n.T(c, "bad_action"))
			return
		}
		user, err := p.Dao.GetUserByUid(tkn["user"].(string))
		if err != nil {
			Message(c, false, err.Error())
			return
		}
		ok, msg := fn(c, user)
		Message(c, ok, msg)
	}
}

func (p *UsersEngine) Mount(r *gin.Engine) {
	r.POST("/users/sign_in", func(c *gin.Context) {
		//todo
	})
	r.POST("/users/sign_up", Form(&SignUpForm{}, func(c *gin.Context, f interface{}) (interface{}, error) {
		fm := f.(*SignUpForm)
		if u, _ := p.Dao.GetUserByEmail(fm.Email); u != nil {
			return nil, errors.New(p.I18n.T(c, "valid.user.email.exists", fm.Email))
		}

		u, e := p.Dao.NewEmailUser(fm.Username, fm.Email, fm.Password)
		if e != nil {
			return nil, e
		}
		p.Dao.Log(u.ID, p.I18n.T(c, "log.user.sign_up"))
		p.sendMail(c, "confirm", u)
		return []string{p.I18n.T(c, "log.user.messages.send_confirm_instructions")}, nil
	}))
	r.GET("/users/confirm", p.token("confirm", func(c *gin.Context, user *User) (bool, string) {

		if user.ConfirmedAt != nil {
			return false, p.I18n.T(c, "valid.user.already_confirmed")
		}

		err := p.Dao.ConfirmUser(user.ID)
		if err != nil {
			return false, err.Error()
		}
		p.Dao.Log(user.ID, p.I18n.T(c, "log.user.confirm"))
		return true, p.I18n.T(c, "log.user.messages.confirmed")

	}))
	r.POST("/users/confirm", Form(&EmailForm{}, func(c *gin.Context, f interface{}) (interface{}, error) {
		fm := f.(*EmailForm)
		u, e := p.Dao.GetUserByEmail(fm.Email)
		if e != nil {
			return nil, errors.New(p.I18n.T(c, "valid.user.email.not_exists", fm.Email))
		}
		if u.ConfirmedAt != nil {
			return nil, errors.New(p.I18n.T(c, "valid.user.already_confirmed"))
		}
		p.sendMail(c, "confirm", u)
		return []string{p.I18n.T(c, "log.user.messages.send_confirm_instructions")}, nil
	}))
	r.POST("/users/forgot_password", Form(&EmailForm{}, func(c *gin.Context, f interface{}) (interface{}, error) {
		fm := f.(*EmailForm)
		u, e := p.Dao.GetUserByEmail(fm.Email)
		if e != nil {
			return nil, errors.New(p.I18n.T(c, "valid.user.email.not_exists", fm.Email))
		}
		p.sendMail(c, "reset_password", u)
		return []string{p.I18n.T(c, "log.user.messages.send_reset_password_instructions")}, nil
	}))
	r.POST("/users/reset_password", func(c *gin.Context) {
		//todo
	})
	r.GET("/users/unlock", func(c *gin.Context) {
		//todo
	})
	r.POST("/users/unlock", Form(&EmailForm{}, func(c *gin.Context, f interface{}) (interface{}, error) {
		fm := f.(*EmailForm)
		u, e := p.Dao.GetUserByEmail(fm.Email)
		if e != nil {
			return nil, errors.New(p.I18n.T(c, "valid.user.email.not_exists", fm.Email))
		}
		if u.LockedAt == nil {
			return nil, errors.New(p.I18n.T(c, "valid.user.not_locked"))
		}
		p.sendMail(c, "unlock", u)
		return []string{p.I18n.T(c, "log.user.messages.send_unlock_instructions")}, nil
	}))
	r.GET("/users/profile", func(c *gin.Context) {
		//todo
	})
	r.POST("/users/profile", func(c *gin.Context) {
		//todo
	})
}

func (p *UsersEngine) Seed() error {
	var count int
	p.Db.Model(User{}).Count(&count)
	if count == 0 {
		var root *User
		var adminR *Role
		var rootR *Role
		var err error
		if root, err = p.Dao.NewEmailUser("root", fmt.Sprintf("root@%s", p.Http.Domain), "changeme"); err != nil {
			return err
		}

		dur := 24 * 365 * 10 * time.Hour

		if err = p.Dao.ConfirmUser(root.ID); err != nil {
			return err
		}
		if rootR, err = p.Dao.NewRole("root", "-", 0); err != nil {
			return err
		}
		if err = p.Dao.Apply(rootR.ID, root.ID, dur); err != nil {
			return err
		}
		if adminR, err = p.Dao.NewRole("admin", "-", 0); err != nil {
			return err
		}
		if err = p.Dao.Apply(adminR.ID, root.ID, dur); err != nil {
			return err
		}

	}
	return nil
}

func (p *UsersEngine) Migrate() {
	db := p.Db
	db.AutoMigrate(&User{}, &Contact{}, &Role{}, &Permission{}, &Log{})
	db.Model(&User{}).AddUniqueIndex("idx_user_provider_type_id", "provider_type", "provider_id")
	db.Model(&Role{}).AddUniqueIndex("idx_roles_name_resource_type_id", "name", "resource_type", "resource_id")
	db.Model(&Permission{}).AddUniqueIndex("idx_permissions_user_role", "user_id", "role_id")
}

//==============================================================================
func init() {
	engine.Register(&UsersEngine{})
}
