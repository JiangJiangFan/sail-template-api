package models

import (
	g "sail-chat/global"
	"time"

	"gorm.io/gorm"
)

type UserRole struct {
	Model
	// 用户ID
	UserId uint `json:"user_id" gorm:"primaryKey;index"`
	RoleId uint `json:"role_id,omitempty" gorm:"primaryKey;index"`
	// User        User `json:"user" gorm:"foreignKey:UserID"`
	// 用户权限
	// Role 				Role `json:"role,omitempty" gorm:"foreignKey:RoleID"`
}

// CreateUserRole 用户关联角色
func CreateUserRole(d *gorm.DB, userId uint, roleId uint) error {
	uRole := UserRole{
		UserId: userId,
		RoleId: roleId,
	}
	uRole.CreatedAt = LocalTime{time.Now()}
	err := d.Create(&uRole).Error
	return err
}

func UpdateUserRoleByUserId(d *gorm.DB,user_id uint, role_id uint) error {
	err := d.Model(&UserRole{}).Select("role_id","updated_at").Where("user_id = ?",user_id).Updates(map[string]interface{}{
		"role_id": role_id,
		"updated_at": LocalTime{time.Now()},
	}).Error
	return err
}

func DelRoleByUserId(userId uint) error {
	omit := []string{"create_at", "deleted_at", "id","user_id"}
	err := g.App.DB.Model(&UserRole{}).Omit(omit...).Where("user_id = ?",userId).Updates(map[string]interface{}{
		"updated_at": time.Now(),
		"is_deleted": true,
	}).Error
	return err
}

func GetRoleByUserId(userId uint) (UserRole, error){
	var uRole UserRole
	err := g.App.DB.Where("user_id = ?",userId).First(&uRole).Error
	return uRole, err
}
