package types

type Admin struct {
	// 定义User结构体
	Id       uint   `form:"id,omitempty" json:"id,omitempty"`
	Name     string `form:"name,omitempty" json:"name" binding:"required" label:"用户名"`
	Pass     string `form:"pass,omitempty" json:"pass"`
	NewPass  string `form:"new_pass,omitempty" json:"new_pass,omitempty"`
	Account  string `form:"account,omitempty" json:"account,omitempty"`
	Phone    string `form:"phone,omitempty"  json:"phone"`
	Avator   string `form:"avator,omitempty"   json:"avator,omitempty"`
	SectId   uint   `form:"sect_id,omitempty"  json:"sect_id"`
	SectName string `form:"sect_name,omitempty"  json:"sect_name"`
}