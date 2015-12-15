package base

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/itpkg/portal/base/engine"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model

	Uid          string    `sql:"not null;unique;type:char(36)"`
	Logo         string    `sql:"not null"`
	Name         string    `sql:"not null"`
	ProviderType string    `sql:"not null;default:'email';index:idx_users_provider_type"`
	ProviderId   string    `sql:"not null;index:idx_users_provider_id"`
	LastSignIn   time.Time `sql:"not null"`

	Roles []Role `gorm:"many2many:users_roles;"`
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

	Begin time.Time `sql:"not null;default:current_date;type:date"`
	End   time.Time `sql:"not null;default:'1000-1-1';type:date"`
}

//==============================================================================
type UsersEngine struct {
	Db *gorm.DB `inject:"db"`
}

func (p *UsersEngine) Mount(*gin.Engine) {
}

func (p *UsersEngine) Seed() {

}

func (p *UsersEngine) Migrate() {
	db := p.Db
	db.AutoMigrate(&User{}, &Contact{}, &Role{}, &Log{})
	db.Model(&User{}).AddUniqueIndex("idx_user_provider_type_id", "provider_type", "provider_id")
	db.Model(&Role{}).AddUniqueIndex("idx_roles_name_resource_type", "name", "resource_type")
}

//==============================================================================
func init() {
	engine.Register(&UsersEngine{})
}
