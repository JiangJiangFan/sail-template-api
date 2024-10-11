package models

// Visitor 角色
type Visitor struct {
	// 定义 Visitor 结构体
	Model
	Username 	string `json:"username" gorm:"not null; comment:访客名称"`
	Avator 		string `json:"avator" gorm:"not null; default:''"`
	SourceIp 	string `json:"source_ip" gorm:"not null; default:''; comment:回复IP"`
	ToName 		  string `json:"to_name" gorm:"index:to_name,unique; not null; default:''; comment:对应客服"`
	VisitorId string `json:"visitor_id" gorm:"index:visitor_id,unique; not null; default:''; comment:访客UUID"`
	Status 		bool 	 `json:"status" gorm:"not null; default:0; comment:在线状态"`
	Refer 		string `json:"refer" gorm:"not null; default:'';comment:信息来源"`
	City 			string `json:"city" gorm:"not null; default:''; comment:城市"`
	ClientIp 	string `json:"client_ip" gorm:"not null; default:''; comment:客户IP"`
	Extra 		string `json:"extra" gorm:"not null; default:''; comment:备注"`
}