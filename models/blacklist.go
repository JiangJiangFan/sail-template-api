package models

type Blacklist struct {
	Model
	IP     string `json:"ip" gorm:"index:ip,unique; not null; default:''"`
	UserId uint   `json:"u_id" gorm:"not null;"`
}