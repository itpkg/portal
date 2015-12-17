package base

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itpkg/portal/base/cfg"
	"github.com/itpkg/portal/base/engine"
	"github.com/jinzhu/gorm"
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

func Form(fm interface{}, fn func(*gin.Context, interface{}) (interface{}, error)) func(c *gin.Context) {
	return func(c *gin.Context) {
		if err := c.Bind(fm); err == nil {
			if data, err := fn(c, fm); err == nil {
				c.JSON(http.StatusOK, gin.H{"ok": true, "data": data})
			} else {
				c.JSON(http.StatusOK, gin.H{"ok": false, "errors": []string{err.Error()}})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"ok": false, "errors": strings.Split(err.Error(), "\n")})
		}
	}
}

//==============================================================================
type UsersEngine struct {
	Db   *gorm.DB  `inject:""`
	Dao  *Dao      `inject:""`
	Http *cfg.Http `inject:""`
	I18n *I18n     `inject:""`
}

func (p *UsersEngine) Build(dir string) error {
	return nil
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
		tx := p.Db.Begin()
		u, e := p.Dao.NewEmailUser(tx, fm.Username, fm.Email, fm.Password)
		if e != nil {
			tx.Rollback()
			return nil, e
		}
		p.Dao.Log(tx, u.ID, p.I18n.T(c, "log.user.sign_up"))
		tx.Commit()
		return nil, nil
	}))
	r.GET("/users/confirm", func(c *gin.Context) {
		//todo
	})
	r.POST("/users/confirm", func(c *gin.Context) {
		//todo
	})
	r.POST("/users/forgot_password", func(c *gin.Context) {
		//todo
	})
	r.POST("/users/change_password", func(c *gin.Context) {
		//todo
	})
	r.GET("/users/unlock", func(c *gin.Context) {
		//todo
	})
	r.POST("/users/unlock", func(c *gin.Context) {
		//todo
	})
	r.GET("/users/profile", func(c *gin.Context) {
		//todo
	})
	r.POST("/users/profile", func(c *gin.Context) {
		//todo
	})
}

func (p *UsersEngine) Seed() error {
	tx := p.Db.Begin()
	var count int
	tx.Model(User{}).Count(&count)
	if count == 0 {
		var root *User
		var adminR *Role
		var rootR *Role
		var err error
		if root, err = p.Dao.NewEmailUser(tx, "root", fmt.Sprintf("root@%s", p.Http.Domain), "changeme"); err != nil {
			tx.Rollback()
			return err
		}

		dur := 24 * 365 * 10 * time.Hour

		if err = p.Dao.ConfirmUser(tx, root.ID); err != nil {
			tx.Rollback()
			return err
		}
		if rootR, err = p.Dao.NewRole(tx, "root", "-", 0); err != nil {
			tx.Rollback()
			return err
		}
		if err = p.Dao.Apply(tx, rootR.ID, root.ID, dur); err != nil {
			tx.Rollback()
			return err
		}
		if adminR, err = p.Dao.NewRole(tx, "admin", "-", 0); err != nil {
			tx.Rollback()
			return err
		}
		if err = p.Dao.Apply(tx, adminR.ID, root.ID, dur); err != nil {
			tx.Rollback()
			return err
		}

	}
	tx.Commit()
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
