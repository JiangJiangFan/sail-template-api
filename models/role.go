package models

import (
	"sail-chat/global"
	"sail-chat/types"
	"sail-chat/utils"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	Model
	Name string `json:"name" gorm:"not null; comment:名称"`
	Method string `json:"method" gorm:"not null; default:''"`
	Path string `json:"path" gorm:"not null; default:''"`
	// User []User `gorm:"many2many:user_roles"`
}

func CreateRole(r Role) (uint, error) {
	role := Role{Name: r.Name}
	role.CreatedAt = LocalTime{time.Now()}
	err := global.App.DB.Create(&role).Error
	return role.Id, err
}

func DeleteRole(id uint) error {
	Role := Role{Model: Model{Id: id}}
	err := global.App.DB.Model(&Role).Updates(map[string]interface{}{"is_deleted": true,"updated_at": time.Now()}).Error
	return err
}

// UpdateRole 更新角色——不包括方法和路径
func UpdateRole(id uint, name string) error {
	role := Role{Model: Model{Id: id}}
	err := global.App.DB.Model(&role).Updates(map[string]interface{}{"name": name,"updated_at": time.Now()}).Error
	return err
}

func GetRole(tx *gorm.DB, id uint) (Role, error) {
	Role := Role{}
	err := tx.Where("id = ?", id).First(&Role).Error
	return Role, err
}

func GetRoleList(meta *types.Meta) ([]Role, int, error) {
	var db = global.App.DB.Model(&Role{})
	var count int64
	var Role []Role
	var err error
	db.Where("name like ?", "%"+meta.Value+"%").Count(&count)
	if count < meta.Size {
		err = db.Having("name like ?", "%"+meta.Value+"%").Find(&Role).Error
	} else {
		err = db.Having("name like ?", "%"+meta.Value+"%").Scopes(utils.Paginate(meta)).Find(&Role).Error
	}
	return Role, int(count), err
}
