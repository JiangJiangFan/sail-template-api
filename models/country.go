package models

import (
	"sail-chat/global"
	"sail-chat/types"
	"sail-chat/utils"
)

type Country struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	PID   uint   `json:"pid" gorm:"not null; comment: 父级id"`
	Name  string `json:"name" gorm:"not null; comment: 国家/城市名称"`
	Level uint   `json:"level" gorm:"not null; comment: 等级"`
}

type CountryList struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null; comment: 国家/城市名称"`
	PName string `json:"p_name" gorm:"not null; comment: 父级名称"`
	PID   uint   `json:"pid,omitempty" gorm:"not null; comment: 父级id"`
	Level uint   `json:"level,omitempty" gorm:"not null; comment: 等级"`
}

// CreateCountry 创建国家/城市
func CreateCountry(country *Country) (uint, error) {
	// 创建国家/城市
	err := global.App.DB.Create(&country)
	return country.ID, err.Error
}

// SelectCountryByPID 根据父级id查询国家/城市
func SelectCountryByPID(pid uint) ([]CountryList, error) {
	var countryList []CountryList
	err := global.App.DB.Where("pid = ?", pid).Find(&countryList).Error
	return countryList, err
}

// SelectCountryByParam 根据参数查询国家/城市
func SelectCountryByParam(c types.Country, m *types.Meta) ([]CountryList, error) {
	var countryList []CountryList
	var count int64
	var db = global.App.DB.Model(&Country{})
	if c.PID != 0 {
		db.Where("pid = ?", c.PID)
	}
	if c.Name != "" {
		db.Where("name LIKE ?", "%"+c.Name+"%")
	}
	db.Count(&count)
	err := db.Scopes(utils.Paginate(m)).Find(&countryList).Error
	return countryList, err
}

// IsCountryExist 根据参数查询国家/城市是否存在
func IsCountryExist(c *types.Country) (bool, error) {
	var country Country
	err := global.App.DB.Where("pid = ? and name = ?", c.PID, c.Name).First(&country).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// SelectCountryByName 根据用户名称查询父级是否存在该城市
func SelectCountryByName(c types.Country) (Country, error) {
	var country Country
	err := global.App.DB.Where("pid = ? and name = ?", c.PID, c.Name).First(&country).Error
	return country, err
}

// UpdateCountry 更新国家/城市
func UpdateCountry(country *Country) (uint, error) {
	err := global.App.DB.Save(&country).Error
	return country.ID, err
}

// SelectCountryCustom 自定义查询
func SelectCountryCustom(m *types.Meta) ([]CountryList, error) {
	var countryList []CountryList
	// offset, limit := utils.PaginateString(m)
	// const str = `select name , (select name from countries where id = (c.p_id))
	// 	 as p_name from countries c where name like ?`
	var db = global.App.DB
	// err := db.Raw(str, "%"+m.Value+"%").Scopes(utils.Paginate(m)).Scan(&countryList).Error
	err := db.Table("countries c").Model(&Country{}).Select("id, name , (select name from countries where id = (c.p_id)) as p_name, p_id, level").Where("name like ?", "%"+m.Value+"%").Scopes(utils.Paginate(m)).Find(&countryList).Error
	return countryList, err
}

// SelectCountryCustomByValue  自定义查询
func SelectCountryCustomByValue(m *types.Meta) ([]CountryList, error) {
	var countryList []CountryList
	var db = global.App.DB
	err := db.Table("countries c").Model(&Country{}).Select("id, name , (select name from countries where id = (c.p_id)) as p_name, level").Where("name like ?", "%"+m.Value+"%").Find(&countryList).Error
	return countryList, err
}
