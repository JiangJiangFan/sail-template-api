package types

type User struct {
	// 定义User结构体
	Id       uint   `form:"id,omitempty" json:"id,omitempty"`
	Username string `form:"username,omitempty" json:"username" binding:"required" label:"用户名"`
	Password string `form:"password,omitempty" json:"password"`
	NewPass  string `form:"new_pass,omitempty" json:"new_pass,omitempty"`
	Nickname string `form:"nickname,omitempty" json:"nickname,omitempty"`
	Avator   string `form:"avator,omitempty"   json:"avator,omitempty"`
	RoleId   uint   `form:"role_id,omitempty"  json:"role_id"`
}
