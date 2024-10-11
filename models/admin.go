package models

import (
	"sail-chat/types"
	"time"

	"gorm.io/gorm"
)

// Admin 管理员 【厂家 | 企业】
type Admin struct {
	Model
	Name string `json:"name" gorm:"not null; comment:管理员名称"`
	Account string `json:"account" gorm:"uniqueIndex; not null; comment:管理员账号"`
	Pass string `json:"pass" gorm:"not null; comment:管理员密码"`
	Phone string `json:"phone" gorm:"commnet:管理员手机号"`
}

type NotAdminPass struct {
	Model
	Name string `json:"name" gorm:"not null; comment:管理员名称"`
	Account string `json:"account" gorm:"uniqueIndex; not null; comment:管理员账号"`
	Phone string `json:"phone" gorm:"commnet:管理员手机号"`
	Sect string `json:"sect_name" sql:"-"`
}

func CreateAdmin(d *gorm.DB, a *types.Admin) (uint, error) {
	admin := Admin{
		Name: a.Name,
		Pass: a.Pass,
		Account: a.Account,
		Phone: a.Phone,
	}
	admin.CreatedAt = LocalTime{time.Now()}
	err := d.Create(&admin).Error
	return admin.Id, err
}