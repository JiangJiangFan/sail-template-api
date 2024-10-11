package models

import (
	"time"

	"gorm.io/gorm"
)

type Sect struct {
	Model
	Name string `json:"name" gorm:"uniqueIndex; not null; comment:管理员名称"`
}

// CreateSect 创建用户
func CreateSect(d *gorm.DB, name string) (uint, error) {
	sect := Sect{
		Name: name,
	}
	sect.CreatedAt = LocalTime{time.Now()}
	err := d.Create(&sect).Error
	return sect.Id, err
}