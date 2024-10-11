package models

import (
	"sail-chat/global"
	"sail-chat/types"
	"sail-chat/utils"
)

type Abilities struct {
	Model
	Name string `json:"name" gorm:"not null; comment:名称"`
	User []User `gorm:"many2many:user_abilities;"`
}

func CreateAbilities(name string) (uint, error) {
	abilities := Abilities{Name: name}
	err := global.App.DB.Create(&abilities).Error
	return abilities.Id, err
}

func DeleteAbilities(id uint) error {
	abilities := Abilities{Model: Model{Id: id}}
	err := global.App.DB.Model(&abilities).Updates(map[string]interface{}{"is_deleted": true}).Error
	return err
}

func UpdateAbilities(id uint, name string) error {
	abilities := Abilities{Model: Model{Id: id}}
	err := global.App.DB.Model(&abilities).Updates(map[string]interface{}{"name": name}).Error
	return err
}

func GetAbilities(id uint) (Abilities, error) {
	abilities := Abilities{}
	err := global.App.DB.Where("id = ?", id).First(&abilities).Error
	return abilities, err
}

func GetAbilitiesList(meta *types.Meta) ([]Abilities, int, error) {
	var db = global.App.DB.Model(&Abilities{})
	var count int64
	var abilities []Abilities
	var err error
	db.Where("name like ?", "%"+meta.Value+"%").Count(&count)
	if count < meta.Size {
		err = db.Having("name like ?", "%"+meta.Value+"%").Find(&abilities).Error
	} else {
		err = db.Having("name like ?", "%"+meta.Value+"%").Scopes(utils.Paginate(meta)).Find(&abilities).Error
	}
	return abilities, int(count), err
}
