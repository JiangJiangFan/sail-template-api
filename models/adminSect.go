package models

import (
	g "sail-chat/global"
	"time"

	"gorm.io/gorm"
)

type AdminSect struct {
	Model
	// 用户ID
	AdminId uint `json:"admin_id" gorm:"primaryKey;index"`
	SectId  uint `json:"sect_id,omitempty" gorm:"primaryKey;index"`
}

func CreateAdminSect(d *gorm.DB, admin uint, sect uint) error {
	aSect := AdminSect{
		AdminId: admin,
		SectId: sect,
	}
	aSect.CreatedAt = LocalTime{time.Now()}
	err := d.Create(&aSect).Error
	return err
}

func DelAdminSect(admin uint) error {
	omit := []string{"create_at", "deleted_at", "id","admin_id"}
	err := g.App.DB.Model(&AdminSect{}).Omit(omit...).Where("admin_id = ? ",admin).Updates(map[string]interface{}{
		"updated_at": time.Now(),
		"is_deleted": true,
	}).Error
	return err
}

func GetSectByAdmin(admin uint) (AdminSect, error){
	var aSect AdminSect
	err := g.App.DB.Where("admin_id = ?",admin).First(&aSect).Error
	return aSect, err
}