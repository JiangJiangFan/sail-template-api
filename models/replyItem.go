package models

type ReplyItem struct {
	Model
	Content string `json:"content" gorm:"not null; default:''"`
	GID     uint   `json:"g_id" gorm:"index:g_id,unique; not null; comment:组ID"`
	UserId  uint   `json:"u_id" gorm:"index:user_id,unique; not null;"`
	Details string `json:"details" gorm:"not null; default:''"`
}