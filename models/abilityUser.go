package models

import "sail-chat/global"

type AbilitiesUser struct {
	Model
	// 用户ID
	UserID      uint `json:"user_id" gorm:"primaryKey"`
	AbilitiesID uint `json:"ability_id,omitempty" gorm:"primaryKey"`
	User        User `json:"user" gorm:"foreignKey:UserID"`
	// 用户权限
	Abilities Abilities `json:"abilities,omitempty" gorm:"foreignKey:AbilitiesID"`
}

func (AbilitiesUser) TableName() string {
	return "abilities_user"
}

func (a AbilitiesUser) GetAbilitiesUser() (AbilitiesUser, error) {
	var abilitiesUser AbilitiesUser
	var db = global.App.DB
	err := db.First(&abilitiesUser, a).Error
	return abilitiesUser, err
}
