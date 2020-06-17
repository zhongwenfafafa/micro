package dao

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Id        int       `json:"id" gorm:"column:id"`
	Username  string    `json:"username" gorm:"column:username" description:"用户名"`
	Mobile    string    `json:"mobile" gorm:"column:mobile" description:"手机号"`
	Password  string    `json:"password" gorm:"column:password" description:"密码"`
	IsDeleted int       `json:"is_deleted" gorm:"column:is_deleted" description:"是否软删除"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at" description:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at" description:"更新时间"`
}

func (user *User) TableName() string {
	return "account"
}

func (u *User) Find(tx *gorm.DB, id int64) (*User, error) {
	user := &User{}
	err := tx.Table(u.TableName()).Where("id=?", id).First(user).Error
	if err != nil {
		return nil,err
	}

	return user, nil
}
