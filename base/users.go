package base

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itpkg/portal/base/cfg"
	"github.com/itpkg/portal/base/email"
	"github.com/itpkg/portal/base/engine"
	"github.com/itpkg/portal/base/token"
	"github.com/itpkg/portal/base/utils"
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

func (p *User) SetLogoByGravatar() {
	p.Logo = fmt.Sprintf("https://gravatar.com/avatar/%s.png", utils.Md5([]byte(strings.ToLower(p.Email))))
}

func (p *User) IsLocked() bool {
	return p.LockedAt != nil
}
func (p *User) IsConfirmed() bool {
	return p.ConfirmedAt != nil
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
	Email      string `form:"email" binding:"email"`
	Password   string `form:"password" binding:"min=8"`
}
type SignUpForm struct {
	Username             string `form:"username" binding:"min=2,max=20"`
	Email                string `form:"email" binding:"email"`
	Password             string `form:"password" binding:"min=8"`
	PasswordConfirmation string `form:"password_confirmation" binding:"eqfield=Password"`
}

type PasswordForm struct {
	Token                string `form:"token" binding:"required"`
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

	tkn, err := p.Token.New(map[string]interface{}{"action": action, "user": user.Uid}, 60)
	if err != nil {
		p.Logger.Error("bad in generate token: %v", err)
		return
	}

	url := fmt.Sprintf("%s/users/%s?token=%s", p.Http.Home(), action, tkn)
	switch action {
	case "confirm":
	case "unlock":
	case "reset_password":
		url = fmt.Sprintf("%s/?token=%s#/users/reset-password", p.Http.Home(), tkn)
	default:
		return
	}

	var buf bytes.Buffer
	if tpl, err := template.New("").Parse(p.I18n.T(ctx, fmt.Sprintf("email.user.%s.body", action))); err == nil {
		if err := tpl.Execute(&buf, &Model{Username: user.Name, Url: url}); err != nil {
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
	r.POST("/users/sign_in", Form(&SignInForm{}, func(c *gin.Context, f interface{}) (interface{}, error) {
		fm := f.(*SignInForm)
		u, e := p.Dao.GetUserByEmail(fm.Email)
		if e != nil {
			return nil, errors.New(p.I18n.T(c, "valid.user.email.not_exists", fm.Email))
		}
		if b, e := utils.Csha512(u.Password, []byte(fm.Password)); e != nil || !b {
			return nil, errors.New(p.I18n.T(c, "valid.user.email_password_not_match"))
		}
		if u.IsLocked() {
			return nil, errors.New(p.I18n.T(c, "valid.user.locked"))
		}
		if !u.IsConfirmed() {
			return nil, errors.New(p.I18n.T(c, "valid.user.not_confirmed"))
		}
		p.Dao.SetUserSignIn(u)
		p.Dao.Log(u.ID, p.I18n.T(c, "log.user.sign_in"))

		if tk, err := p.Token.New(
			map[string]interface{}{
				"uid": u.Uid,
			},
			60*24); err == nil {
			return map[string]interface{}{
				"token": tk,
				"name":  u.Name,
				"email": u.Email,
				"logo":  u.Logo,
			}, nil
		} else {
			return nil, err
		}

	}))
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

		if user.IsConfirmed() {
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
		if u.IsConfirmed() {
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

	r.POST("/users/reset_password", Form(&PasswordForm{}, func(c *gin.Context, f interface{}) (interface{}, error) {
		fm := f.(*PasswordForm)
		tkn, err := p.Token.Parse(fm.Token)
		if err != nil {
			return nil, err
		}
		if tkn["action"].(string) != "reset_password" {
			return nil, errors.New(p.I18n.T(c, "bad_action"))
		}
		user, err := p.Dao.GetUserByUid(tkn["user"].(string))
		if err != nil {
			return nil, err
		}

		if err := p.Dao.SetUserPassword(user.ID, fm.Password); err != nil {
			return nil, err
		}
		p.Dao.Log(user.ID, p.I18n.T(c, "log.user.reset_password"))

		return []string{p.I18n.T(c, "log.success")}, nil
	}))

	r.GET("/users/unlock", p.token("unlock", func(c *gin.Context, user *User) (bool, string) {
		if !user.IsLocked() {
			return false, p.I18n.T(c, "valid.user.not_locked")
		}
		err := p.Dao.LockUser(user.ID, false)
		if err != nil {
			return false, err.Error()
		}
		p.Dao.Log(user.ID, p.I18n.T(c, "log.user.unlock"))
		return true, p.I18n.T(c, "log.user.messages.unlocked")
	}))
	r.POST("/users/unlock", Form(&EmailForm{}, func(c *gin.Context, f interface{}) (interface{}, error) {
		fm := f.(*EmailForm)
		u, e := p.Dao.GetUserByEmail(fm.Email)
		if e != nil {
			return nil, errors.New(p.I18n.T(c, "valid.user.email.not_exists", fm.Email))
		}
		if !u.IsLocked() {
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
