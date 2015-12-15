package base

import (
	"time"

	"github.com/itpkg/portal/base/utils"
	"github.com/jinzhu/gorm"
)

type Dao struct {
}

func (*Dao) NewEmailUser(tx *gorm.DB, name, email, password string) (*User, error) {
	passwd, err := utils.Ssha512([]byte(password), 8)
	if err != nil {
		return nil, err
	}

	u := User{
		Name:     name,
		Email:    email,
		Password: passwd,
		Uid:      utils.Uuid(),
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
