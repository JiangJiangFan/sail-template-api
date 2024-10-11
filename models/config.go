package models

type Config struct {
	Model
	ConfName  string `json:"conf_name" gorm:"not null; default:''"`
	ConfKey   string `json:"conf_key" gorm:"index:conf_key,unique; not null; default:''"`
	ConfValue string `json:"conf_value" gorm:"not null; default:''"`
}