package base

import (
	"fmt"
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
type UsersEngine struct {
	Db   *gorm.DB  `inject:""`
	Dao  *Dao      `inject:""`
	Http *cfg.Http `inject:""`
}

func (p *UsersEngine) Build(dir string) error {
	return nil
}

func (p *UsersEngine) Mount(*gin.Engine) {
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