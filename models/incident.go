package models

type Incident struct {
	// 事件ID
	ID int `json:"id"`
	// 事件名称
	Name string `json:"name"`
	// 事件描述
	Description string `json:"description"`
	// 事件类型
	Type string `json:"type"`
	// 事件状态
	Status string `json:"status"`
	// 事件发生时间
	OccurTime string `json:"occur_time"`
	// 事件结束时间
	EndTime string `json:"end_time"`
	// 事件影响范围
	Impact string `json:"impact"`
	// 事件恢复时间
	RecoveryTime string `json:"recovery_time"`
	// 事件恢复描述
	RecoveryDescription string `json:"recovery_description"`
	// 事件恢复状态
	RecoveryStatus string `json:"recovery_status"`
	// 事件发生地址
	Address string `json:"address"`
}