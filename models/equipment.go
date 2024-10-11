package models

// Equipment 装备
type Equipment struct {
	// 装备ID
	ID int `json:"id" gorm:"not null;primaryKey"`
	// 装备名称
	Name string `json:"name" gorm:"not null; comment: 装备名称"`
	// 装备类型
	Type string `json:"type" gorm:"not null; comment: 装备类型"`
	// 装备描述
	Description string `json:"description" gorm:"not null; comment: 装备描述"`
}
