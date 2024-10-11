package models

// Part 角色
type Part struct {
	// 定义Part结构体
	Model
	Username string `json:"username" gorm:"not null; comment:角色名称"`
}