package models

type MesType string
type MesStatus string

const (
	MEMBER  MesType = "member"
	VISITOR MesType = "visitor"
)
const (
	READ   MesStatus = "read"
	UNREAD MesStatus = "unread"
)

// Message 属性
type Message struct {
	Model
	UserId    uint      `json:"u_id" gorm:"index:user_id,unique; not null; comment:用户对应ID"`
	VisitorID string    `json:"visitor_Id" gorm:"index:visitor_id,unique; not null; comment:访客对应"`
	Content   string    `json:"content" gorm:"not null; comment:消息内容"`
	MesType   MesType   `json:"mes_type" gorm:"not null; default:'visitor'; comment:消息类型"`
	Status    MesStatus `json:"status" gorm:"not null; default 'unread'; comment:消息状态"`
}