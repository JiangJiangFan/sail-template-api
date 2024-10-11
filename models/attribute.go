package models

// Attribute 属性
type Attribute struct {
	// 属性ID
	ID int64 `json:"id" gorm:"not null" `
	// 属性名称
	Name string `json:"name" gorm:"not null; comment:名称"`
	// 属性值
	Value string `json:"value" gorm:"not null; comment:属性值"`
	// 属性类型
	Type string `json:"type" gorm:"not null; comment:属性类型"`
	// 属性描述
	Description string `json:"description" gorm:"not null; comment:属性描述"`
}