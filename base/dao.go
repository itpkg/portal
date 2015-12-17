package base

import (
	"fmt"
	"time"

	"github.com/itpkg/portal/base/utils"
	"github.com/jinzhu/gorm"
)

type Dao struct {
	Db  *gorm.DB   `inject:""`
	Aes *utils.Aes `inject:""`
}

func (*Dao) site_key(key, lang string) string {
	if lang == "" {
		return fmt.Sprintf("site://%s", key)
	} else {
		return fmt.Sprintf("site://%s/%s", lang, key)
	}

}

func (p *Dao) SetSiteInfo(tx *gorm.DB, key, lang string, val interface{}, flag bool) error {
	return p.Set(tx, p.site_key(key, lang), val, flag)
}

func (p *Dao) GetSiteInfo(key, lang string) string {
	var val string
	p.Get(p.site_key(key, lang), &val)
	return val
}

func (p *Dao) GetUserByEmail(email string) (*User, error) {
	var u User
	if e := p.Db.Where("email = ?", email).First(&u).Error; e == nil {
		return &u, e
	} else {
		return nil, e
	}
}

func (p *Dao) Set(tx *gorm.DB, key string, val interface{}, flag bool) error {
	buf, err := utils.ToBits(val)
	if err != nil {
		return err
	}

	s := Setting{Key: key, Flag: flag}
	if flag {
		if v, e := p.Aes.Encrypt(buf); e == nil {
			s.Val = v
		} else {
			return e
		}
	} else {
		s.Val = buf
	}

	var c int
	tx.Where("key = ?", key).Count(&c)
	if c == 0 {
		return tx.Create(&s).Error
	} else {
		return tx.Model(&Setting{}).Where("key = ?", key).UpdateColumn(val, s.Val).Error
	}

}

func (p *Dao) Get(key string, val interface{}) error {
	var s Setting
	err := p.Db.Where("key = ?", key).First(&s).Error
	if err != nil {
		return err
	}

	var buf []byte
	if s.Flag {
		if buf, err = p.Aes.Decrypt(s.Val); err != nil {
			return err
		}
	} else {
		buf = s.Val
	}
	return utils.FromBits(buf, val)
}

func (*Dao) Log(tx *gorm.DB, user uint, message string) error {
	return tx.Create(&Log{UserID: user, Message: message}).Error
}

func (*Dao) ConfirmUser(tx *gorm.DB, id uint) error {
	return tx.Model(&User{}).Where("id = ?", id).UpdateColumn("confirmed_at", time.Now()).Error
}

func (*Dao) NewEmailUser(tx *gorm.DB, name, email, password string) (*User, error) {
	passwd, err := utils.Ssha512([]byte(password), 8)
	if err != nil {
		return nil, err
	}

	u := User{
		Name:       name,
		Email:      email,
		Password:   passwd,
		Uid:        utils.Uuid(),
		ProviderId: email,
	}
	if err = tx.Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (Dao) NewRole(tx *gorm.DB, name string, resource_type string, resource_id uint) (*Role, error) {
	r := Role{
		Name:         name,
		ResourceType: resource_type,
		ResourceId:   resource_id,
	}
	if err := tx.Create(&r).Error; err != nil {
		return nil, err
	}
	return &r, nil
}

func (Dao) Apply(tx *gorm.DB, role uint, user uint, dur time.Duration) error {
	begin := time.Now()
	end := begin.Add(dur)
	var count int
	tx.Model(Permission{}).Where("role_id = ? AND user_id = ?", role, user).Count(&count)
	if count == 0 {
		return tx.Create(&Permission{
			UserID: user,
			RoleID: role,
			Begin:  begin,
			End:    end,
		}).Error
	} else {
		return tx.Model(&Permission{}).Where("role_id = ? AND user_id = ?", role, user).Updates(map[string]interface{}{"begin": begin, "end": end}).Error
	}
}
