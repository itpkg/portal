package cms

import (
	"github.com/itpkg/portal/base"
)

type Article struct {
	base.Model

	User   *base.User
	UserID uint `sql:"not null"`

	Title string `sql:"not null"`
	Body  string `sql:"not null;type:text"`
	Lang  string `sql:"not null;type:char(5);default:'en-US'"`

	Tags     []Tag `gorm:"many2many:articles_tags;"`
	Comments []Comment
}

type Tag struct {
	base.Model
	Name     string    `sql:"not null;unique"`
	Articles []Article `gorm:"many2many:articles_tags;"`
}

type Comment struct {
	User   base.User
	UserID uint

	Article   Article
	ArticleID uint   `sql:"not null"`
	Content   string `sql:"not null;type:text"`
}
