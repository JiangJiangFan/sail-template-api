package models

type Welcome struct {
	Model
	UserId    uint   `json:"u_id" gorm:"index:; not null;"`
	Keyword   string `json:"keyword" gorm:"index:keyword,unique; not null; default:''"`
	Content   string `json:"content" gorm:"not null;"`
	IsDefault bool   `json:"is_default" gorm:"not null; default:0"`
}
